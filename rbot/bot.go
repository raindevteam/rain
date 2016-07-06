package rbot

import (
	"errors"
	"net"
	"net/rpc"
	"strings"
	"sync"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rlog"
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

// The Module struct serves as a container for a module's commander. The struct holds its
// module's corresponding name to make it easier for reference.
type Module struct {
	Name  string
	Cmder *Commander
}

// NewModule returns a module struct with an assigned commander appropriated to the module's type.
func NewModule(name string, path string, cmdtype string) *Module {
	module := &Module{
		Name:  name,
		Cmder: NewCommander(name, cmdtype, path),
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
	Parser   *Parser
	Handler  *Handler
	Mu       sync.Mutex
}

// NewBot initializes a number of things for proper operation. It will set appropriate flags
// for rlog. It then creates a Nimbus config to pass to the internal nimbus IRC client. This
// client is embedded into an instance of Bot and returned. It has its fields initialized.
func NewBot(version string, rconf *Config) *Bot {
	rlog.SetFlags(rlog.Linfo | rlog.Lwarn | rlog.Lerror)
	rlog.SetLogFlags(0)

	nconf := GetNimbusConfig(rconf)

	bot := &Bot{
		/* Client   */ nimbus.NewClient(rconf.Host, rconf.Nick, *nconf),
		/* Version  */ version,
		/* Modules  */ make(map[string]*Module),
		/* Channels */ make(map[string]*Channel),
		/* Parser   */ NewParser(rconf.CmdPrefix),
		/* Handler  */ NewHandler(),
		/* Mutex    */ sync.Mutex{},
	}

	return bot
}

// DefaultConnect will connect to IRC and start listening. It will not log anything. It also
// waits on Client.Quit(), and when the chan is served, it logs the quit result.
func (b *Bot) DefaultConnect() {
	b.Connect()
	b.Listen()
	result := <-b.Quit()
	rlog.Info("Bot", "Quitting Now: "+result.Error())
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
	result := <-b.Quit()
	rlog.Info("Bot", "Quitting Now: "+result.Error())
}

/////////////////////////          Module Specific Methods         /////////////////////////////////

// EnableModules goes through every module list found in the config and sets them up appropriately.
func (b *Bot) EnableModules(rainConfig *Config) {
	var name, path string

	for name, path = range rainConfig.GoModules {
		lowerName := strings.ToLower(name)
		b.Modules[lowerName] = NewModule(lowerName, path, "go")
	}

	for name, path = range rainConfig.JsModules {
		lowerName := strings.ToLower(name)
		b.Modules[lowerName] = NewModule(lowerName, path, "js")
	}

	b.loadModules()
}

// LoadModules starts the bot's rpc master server and then calls moduleRun() on all modules.
func (b *Bot) loadModules() {
	b.startRPCServer()
	for name := range b.Modules {
		b.moduleRun(name)
	}
}

// ModuleReload will reload a module by telling the handler to signal a kill cleanup to the module
// being reloaded. If the module refuses to be killed for whatever reason, reloading of the module
// will be aborted. If the module complies, the module's corresponding process will be killed. The
// module will then be recompiled if need be and will be restarted after.
func (b *Bot) ModuleReload(name string) (err error) {
	name = strings.ToLower(name)

	if !b.Handler.ModuleExists(name) {
		return errors.New("Module does not exist")
	}

	c := b.Modules[name].Cmder

	err = b.Handler.SignalKill(name)

	if err != nil {
		rlog.Error("Bot", err.Error())
		return errors.New("Module refused to stop, aborting reload")
	}

	err = c.Kill()

	if err != nil {
		return errors.New("Could not kill module, aborting reload")
	}

	err = c.Recompile()

	if err != nil {
		return errors.New("Could not recompile")
	}

	b.moduleRun(name)
	return nil
}

// moduleRun starts a plugin as a provider service. This allows
// the bot to dispatch events outbound to the module. The module
// can then communicate back via the master rpc server.
func (b *Bot) moduleRun(name string) {
	c := b.Modules[name].Cmder
	err := c.Start()
	if err != nil {
		rlog.Warn("Bot", "Could not start: "+name+"("+err.Error()+")")
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
