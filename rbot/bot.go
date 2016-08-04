// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

import (
	"errors"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/parser"
	"github.com/raindevteam/rain/rain"
	"github.com/raindevteam/rain/rlog"
	"gopkg.in/sorcix/irc.v1"
)

const (
	// DefaultModulesRoute tells the bot where to look for modules
	DefaultModulesRoute = "modules/"
)

//////////////////////////////          Bot Internals         //////////////////////////////////////

// The Channel struct is used to store information about connected channels. Channels should be
// added when the bot receives the JOIN event from the IRC server. Topic should be set to an empty
// string and filled when the bot receives the RPL_TOPIC event.
type Channel struct {
	Name  string
	Topic string
	Users map[string]string
	Modes string
}

// NewChannel creates a new channel. Note that topic is set to an empty string in hopes that it
// will be filled upon receiving the RPL_TOPIC event from the IRC server.
func NewChannel(name string) *Channel {
	channel := &Channel{
		Name:  name,
		Topic: "",
		Users: make(map[string]string),
	}
	return channel
}

// The Module struct holds in place all information pertaining to a module. It also contains its
// own process manager.
type Module struct {
	Name       string
	Route      string
	Options    map[string]bool
	PM         *ProcessManager
	Registered bool
}

// NewModule returns a module struct with its given information.
func NewModule(name string, path string, opts map[string]bool, cmdtype string) *Module {
	module := &Module{
		Name:       name,
		Route:      path,
		Options:    opts,
		PM:         NewProcessManager(name, cmdtype, path),
		Registered: false,
	}
	return module
}

// The Client interface allows the bot to have a systematic plugabble nature to its IRC
// capabilities. Any IRC library that complies to this Client interface can be popped in and in
// theory, the bot should operate as normal. The Nimbus IRC library has been implemented to
// fulfill this interface and should work appropriately. This modular nature also allows the bot
// to implement different types of clients meant for different purposes, such as command line
// interfacing and mockable testing. It is currently being decided whether this should be part
// of the public API avaible to the user. There are significant obstacles in the way that
// currently undermine this idea which include a generalized IRC message struct and IRC config.
type Client interface {
	GetNick() string
	SetNick(nick string)
	GetChannels() []string
	Connect() error
	Listen()
	Quit() chan error
	Send(raw ...string)
	Say(to string, text string)
	AddListener(event string, l nimbus.Listener)
	GetListeners(event string) []nimbus.Listener
	Emit(event string, msg *irc.Message)
}

// The Bot struct holds all core information for handling IRC. It has the internal IRC client used
// to make and handle an IRC connection. It holds information about modules, which is used to start
// module processes. It has a channels array to hold information about IRC channels if a user wishes
// to do so. The ToJoinChs map is used to validate whether the bot successfully managed to join
// a channel. It also has a parser and handler for commands and listeners. The bot struct also
// includes two limiters, one inbound for commands and another outbound for IRC sends.
type Bot struct {
	Client
	Version    string
	Modules    map[string]*Module
	Channels   map[string]*Channel
	ToJoinChs  map[string]string
	Parser     *parser.Parser
	Handler    *Handler
	Inlim      *rate.Limiter
	Config     *Config
	ListenPort string
	Mu         sync.Mutex
}

// NewBot initializes a number of things for proper operation. It will set appropriate flags
// for rlog and then creates a Nimbus config to pass to the internal nimbus IRC client. This
// client is embedded into an instance of Bot and returned. It has its fields initialized.
func NewBot(version string, rconf *Config) *Bot {
	rlog.SetFlags(rlog.Linfo | rlog.Lwarn | rlog.Lerror | rlog.Ldebug)
	rlog.SetLogFlags(0)

	nconf := GetNimbusConfig(rconf)

	bot := &Bot{
		/* Client     */ nimbus.NewClient(rconf.Server.Host, rconf.Server.Port,
			rconf.User.Nick, *nconf),
		/* Version    */ version,
		/* Modules    */ make(map[string]*Module),
		/* Channels   */ make(map[string]*Channel),
		/* ToJoinChs  */ make(map[string]string),
		/* Parser     */ parser.NewParser(rconf.Command.Prefix),
		/* Handler    */ NewHandler(),
		/* Limiter    */ rate.NewLimiter(3/5, 3),
		/* Config     */ rconf,
		/* ListenPort */ "0",
		/* Mutex      */ sync.Mutex{},
	}

	return bot
}

// InboundLimiter should be used for rate limiting in bound messages, usually commands. The chan
// value should also be checked as false will indicate that the bot has quit and should not accept
// anymore more inbound messages.
func (b *Bot) InboundLimiter() chan bool {
	ch := make(chan bool, 1)

	go func(ch chan bool) {
		r := b.Inlim.Reserve()
		if !r.OK() {
			rlog.Warn("Bot", "Inbound limiter not able to act")
		}

		// If in case the delay wins over the quit, the handler will have appropriated measures.
		select {
		case <-time.After(r.Delay()):
			ch <- true
		case <-b.Quit():
			ch <- false
		}
	}(ch)

	return ch
}

// RainVersion returns the library version
func (b *Bot) RainVersion() string {
	return rain.Version()
}

// DefaultConnect will connect to IRC and start listening. It will not log anything.
func (b *Bot) DefaultConnect() {
	err := b.Connect()
	if err != nil {
		rlog.Errorf("Bot", "Could not connect to irc server: %s\n ----\n The bot will now exit\n")
		os.Exit(1)
	}
	b.Listen()
}

// DefaultConnectWithMsg is similar to DefaultConnect, except a user may pass a pre-connect and
// post-connect message.
func (b *Bot) DefaultConnectWithMsg(pre string, post string) {
	if pre != "" {
		rlog.Info("Bot", pre)
	}

	b.Connect()

	if post != "" {
		rlog.Info("Bot", post)
	}

	b.Listen()
}

// AddCommand let's a user insert an internal command to the handler
func (b *Bot) AddCommand(name string, command *Command) {
	b.Handler.AddInternalCommand(CommandName(name), command)
}

/////////////////////////          Module Specific Methods         /////////////////////////////////

// EnableModules goes through every module list found in the config and sets them up appropriately.
func (b *Bot) EnableModules() {
	rlog.Info("Bot", "Enabling Modules...")

	b.Handler.AddRemoveCallback(func(name ModuleName) {
		b.Modules[strings.ToLower(string(name))].Registered = false
	})

	for modtype, modules := range b.Config.Module.Modules {
		for name, value := range modules {
			route, optionsArray := b.Parser.ParseModuleValue(value)

			optionsMap := make(map[string]bool)
			for _, opt := range optionsArray {
				optionsMap[opt] = true
			}

			if route == "." {
				route = DefaultModulesRoute + modtype
			}

			b.Modules[strings.ToLower(name)] = NewModule(name, route, optionsMap, modtype)
			rlog.Debug("Bot", "Module "+name+" Created")
		}
	}
	b.loadModules()
}

// LoadModules starts the bot's rpc master server and then calls moduleRun() on all modules.
func (b *Bot) loadModules() {
	b.startRPCServer()
	for name, module := range b.Modules {
		if _, ok := module.Options["noload"]; ok {
			rlog.Info("Bot", "Module "+name+" has option noload, not loading")
			continue
		}

		go func(name string, pm *ProcessManager) {
			result := pm.Start(b.ListenPort)
			b.moduleExited(name, result)
		}(name, module.PM)

		rlog.Info("Bot", "Module "+name+" loaded")
	}
}

// moduleExited is called when a module process has exited. If an error was returned when exiting
// it will be logged. The output (Stdout and Stderr) is also logged. If the module exists in the
// handler it will be removed. The module is also unregistered from the bot.
func (b *Bot) moduleExited(name string, result *Result) {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	if result.Err != nil {
		rlog.Error("Bot", name+" [Module Process] has exited with: "+result.Err.Error())
	} else {
		rlog.Info("Bot", "Module "+name+" [Module Process] has exited")
	}

	if result.Output != "" {
		rlog.Info(name, "Process output:"+"\n ----\n"+result.Output+"\n ----\n")
	}

	if b.Handler.ModuleExists(name) {
		// If the module process is dead we can rest assured signalling it to clean up will fail
		// so we won't even try. It's the sanest thing to do!
		b.Handler.RemoveModule(ModuleName(name))
	}

	if !b.Modules[name].Registered {
		rlog.Error("Bot", name+" [Module Client] did not manage to register")
	} else {
		b.Modules[name].Registered = false
		rlog.Info("Bot", name+" [Module Client] has been unregistered from the bot")
	}
}

// ModuleReload will reload a module by telling the handler to signal a kill cleanup to the module
// being reloaded. If the module refuses to be killed for whatever reason, reloading of the module
// will be aborted. If the module complies, the module's corresponding process will be killed. The
// module will then be recompiled if need be and will be restarted after.
func (b *Bot) ModuleReload(name string) (err error) {
	module, ok := b.Modules[strings.ToLower(name)]

	if !ok {
		return errors.New("Module is unknown to the bot")
	}

	if module.PM.IsRunning() {
		err = b.Handler.SignalCleanup(ModuleName(name))
		if err != nil {
			rlog.Error("Bot", "Error while cleaning up module "+name+": "+err.Error())
			return errors.New("Module did not cleanup, aborting reload")
		}

		module.PM.Kill()
		module.PM.Wait()
	}

	res := module.PM.Recompile()
	if res != nil && res.Err != nil {
		rlog.Errorf("Bot", "Could not recompile module %s\n ----\n%s\n ----\n\n",
			module.Name, res.Output)
		return errors.New("Could not recompile module")
	}

	b.moduleStart(name)
	return nil
}

// ModuleLoad will start a module process by calling moduleStart. It will not recompile,
// ModuleReload should be used for that instead. If the module is not found, an error is returned
// instead.
func (b *Bot) ModuleLoad(name string) error {
	if _, ok := b.Modules[strings.ToLower(name)]; !ok {
		return errors.New("Module is unknown to the bot")
	}

	b.moduleStart(strings.ToLower(name))
	return nil
}

// ModuleInfo gathers all relevant information about a module and formats into a returned string. If
// the module is not found within the bot and error is returned.
func (b *Bot) ModuleInfo(name string) (string, error) {
	module, ok := b.Modules[strings.ToLower(name)]

	if !ok {
		return "", errors.New("Module is unknown to the bot")
	}

	var info string

	if module.Registered {
		info = info + "+" + module.Name
	} else {
		info = info + "-" + module.Name
	}

	return info + " | " + module.PM.Type, nil
}

// moduleStart will start a module's process manager. It will then run a goroutine to check for when
// module process exits.
func (b *Bot) moduleStart(name string) {
	pm := b.Modules[name].PM

	if !pm.IsRunning() {
		go func(name string, pm *ProcessManager) {
			result := pm.Start(b.ListenPort)
			b.moduleExited(name, result)
		}(name, pm)
	}
}

///////////////////////////          IRC Specific Methods         //////////////////////////////////

// RemoveUser will delete a user entry from the given channel. If the user is the bot itself, the
// bot will isntead remove the channel from the bot's channel list.
func (b *Bot) RemoveUser(nick string, channel string) {
	if nick == b.GetNick() {
		delete(b.Channels, strings.ToLower(channel))
		return
	}

	delete(b.Channels[strings.ToLower(channel)].Users, nick)
}

///////////////////////////          RPC Specific Methods         //////////////////////////////////

// startRPCServer registers the master consumer for plugins. The master consumer allows plugins to
// communicate with the bot, allowing access to connected channels, users and registered modules.
// Conventionally, it uses a json codec to serve.
func (b *Bot) startRPCServer() {
	rpc.RegisterName("Master", BotAPI{b})
	master, err := net.Listen("tcp", ":0")
	b.ListenPort = strconv.Itoa(master.Addr().(*net.TCPAddr).Port)
	rlog.Info("Bot", "Listening on port: "+b.ListenPort)

	if err != nil {
		rlog.Error("Bot", err.Error())
	}

	// Start accepting connections
	go func() {
		for {
			conn, _ := master.Accept()
			go rpc.ServeCodec(RpcCodecServer(conn))
		}
	}()
}
