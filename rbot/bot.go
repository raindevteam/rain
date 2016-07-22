package rbot

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"strings"
	"sync"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/parser"
	"github.com/raindevteam/rain/rlog"
)

var DefaultModulesRoute = "modules/"

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

// The Module struct serves as a container for a module's commander. The struct holds its
// module's corresponding name to make it easier for reference.
type Module struct {
	Name    string
	Route   string
	Options map[string]bool
	PM      *ProcessManager
}

// NewModule returns a module struct with an assigned commander appropriated to the module's type.
func NewModule(name string, path string, opts map[string]bool, cmdtype string) *Module {
	module := &Module{
		Name:    name,
		Route:   path,
		Options: opts,
		PM:      NewProcessManager(name, cmdtype, path),
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
	GetChannels() []string
	Connect() error
	Listen()
	Quit() chan error
	Send(raw ...string)
	Say(channel string, text string)
	AddListener(event string, l nimbus.Listener)
	GetListeners(event string) []nimbus.Listener
	Emit(event string, msg *nimbus.Message)
}

// The Bot struct holds the internal Client, used to register listeners for IRC. ModuleNames
// is used to look up which plugins to start and are eventually passed to the handler. The Handler
// provides management of commands, listeners and triggers. Since listeners act independently of
// each other, a mutex is used to keep bot writes (such as channel updates) synchronised. The bot
// struct can be realized as a core to glue all modular components together, whilst keeping them
// separate in operation.
type Bot struct {
	Client
	Version  string
	Modules  map[string]*Module
	Channels map[string]*Channel
	Parser   *parser.Parser
	Handler  *Handler
	Config   *Config
	Mu       sync.Mutex
}

// NewBot initializes a number of things for proper operation. It will set appropriate flags
// for rlog. It then creates a Nimbus config to pass to the internal nimbus IRC client. This
// client is embedded into an instance of Bot and returned. It has its fields initialized.
func NewBot(version string, rconf *Config) *Bot {
	rlog.SetFlags(rlog.Linfo | rlog.Lwarn | rlog.Lerror | rlog.Ldebug)
	rlog.SetLogFlags(0)

	nconf := GetNimbusConfig(rconf)

	bot := &Bot{
		/* Client   */ nimbus.NewClient(rconf.Server.Host, rconf.User.Nick, *nconf),
		/* Version  */ version,
		/* Modules  */ make(map[string]*Module),
		/* Channels */ make(map[string]*Channel),
		/* Parser   */ parser.NewParser(rconf.Command.Prefix),
		/* Handler  */ NewHandler(),
		/* Config   */ rconf,
		/* Mutex    */ sync.Mutex{},
	}

	return bot
}

// DefaultConnect will connect to IRC and start listening. It will not log anything.
func (b *Bot) DefaultConnect() {
	b.Connect()
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

/////////////////////////          Module Specific Methods         /////////////////////////////////

// EnableModules goes through every module list found in the config and sets them up appropriately.
func (b *Bot) EnableModules() {
	rlog.Info("Bot", "Enabling Modules...")
	fmt.Println(b.Config.Module.Modules)
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
			rlog.Debug("Bot", "Module"+name+" Created")
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

		cmd := module.PM.Start()

		go func(name string, cmd chan *Result) {
			result := <-cmd
			b.moduleExited(name, result)
		}(name, cmd)

		rlog.Info("Bot", "Module "+name+" loaded")
	}
}

func (b *Bot) moduleExited(name string, result *Result) {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	if result.Err != nil {
		rlog.Error("Bot", name+"[Module Process] has exited with error\n ::::\n"+result.Err.Error())
	} else {
		rlog.Info("Bot", "Module "+name+"[Module Process] has exited")
	}

	if b.Handler.ModuleExists(name) {
		// If the module process is dead we can rest assured signalling it to clean up will fail
		// so we won't even try. It's the sanest thing to do!
		b.Handler.RemoveModule(ModuleName(name))
	} else {
		// We consider a module not starting before registering as an unrecoverable state.
		rlog.Error("Bot", name+"[Module Client] did manage to register, aborting startup")
		os.Exit(1)
	}

	delete(b.Modules, name)
}

// ModuleReload will reload a module by telling the handler to signal a kill cleanup to the module
// being reloaded. If the module refuses to be killed for whatever reason, reloading of the module
// will be aborted. If the module complies, the module's corresponding process will be killed. The
// module will then be recompiled if need be and will be restarted after.
func (b *Bot) ModuleReload(name string) (err error) {
	if !b.Handler.ModuleExists(name) {
		return errors.New("Module does not exist")
	}

	pm := b.Modules[strings.ToLower(name)].PM

	err = b.Handler.SignalCleanup(ModuleName(name))
	if err != nil {
		rlog.Error("Bot", err.Error())
		return errors.New("Module refused to stop, aborting reload")
	}

	err = pm.Kill()
	if err != nil {
		return errors.New("Could not kill module, aborting reload")
	}

	err = pm.Recompile()
	if err != nil {
		return errors.New("Could not recompile module")
	}

	b.moduleRun(name)
	return nil
}

// moduleRun starts a plugin as a provider service. This allows
// the bot to dispatch events outbound to the module. The module
// can then communicate back via the master rpc server.
func (b *Bot) moduleRun(name string) {
	pm := b.Modules[name].PM
	pm.Start()
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
	master, err := net.Listen("tcp", ":5555")

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
