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

func (b *Bot) DefaultConnect() {
	b.Connect()
	b.Listen()
	result := <-b.Quit()
	rlog.Info("Bot", "Quitting Now: "+result.Error())
}

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

// SetupModules goes through every module list found in the config and sets them up appropriately.
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

/////////////////////////          Bot Master Consumer API         /////////////////////////////////

// BotAPI is the exposed api served via the bot's master consumer connection
type BotAPI struct {
	bot *Bot
}

// A Ticket is used to connect to a provider connection via rpc.
type Ticket struct {
	Port string
	Name ModuleName
}

// A CommandRequest is used to register commands via the handler.
type CommandRequest struct {
	Name   CommandName
	Module ModuleName
}

// A TriggerRequest is used to register triggers via the handler.
type TriggerRequest struct {
	Name  ModuleName
	Event Event
}

// Send transmits a message over irc as a PRIVMSG
func (b BotAPI) Send(raw string, result *string) error {

	b.bot.Send(nimbus.PRIVMSG, raw)
	return nil
}

// RegisterCommand registers a command from a module with the handler. The command request
// holds the command's name and the module it belongs to (used to signal the module to fire the
// command).
func (b BotAPI) RegisterCommand(cr CommandRequest, result *string) error {
	b.bot.Handler.AddCommand(cr.Name, cr.Module)
	rlog.Debug("Bot", "Added: "+string(cr.Name)+" for module: "+string(cr.Module))
	return nil
}

// RegisterTrigger will register a trigger from a module with the bot handler. If this the first
// trigger for its corresponding event, the bot will add a new listener that handles trigger firing
// for this event.
func (b BotAPI) RegisterTrigger(tr TriggerRequest, result *string) error {
	listeners := b.bot.GetListeners(string(tr.Event))
	if len(listeners) == 0 {
		b.bot.AddListener(string(tr.Event), func(msg *nimbus.Message) {
			b.bot.Handler.Fire(msg, tr.Event)
		})
	}
	b.bot.Handler.AddTrigger(tr.Name, tr.Event)
	return nil
}

// GetVersion will return the bot's current version.
func (b BotAPI) GetVersion(mName string, result *string) error {
	*result = b.bot.Version
	return nil
}

// GetConnectedUsers will return a user map (where every user has an IRC rank as a value).
func (b BotAPI) GetConnectedUsers(channel string, result *map[string]string) error {
	*result = b.bot.Channels[strings.ToLower(channel)].Users
	return nil
}

// GetTopic returns the channel's topic.
func (b BotAPI) GetTopic(channel string, result *string) error {
	if _, ok := b.bot.Channels[strings.ToLower(channel)]; !ok {
		*result = ""
		return errors.New("Channel doesn't exist")
	}

	*result = b.bot.Channels[strings.ToLower(channel)].Topic
	return nil
}

func (bpi BotAPI) Loop(n string, result *string) error {
	c := make(chan bool)
	func(ch chan bool) {

	}(c)
	<-c
	return nil
}

// Register registers a module with the bot. With the given port number in the Ticket, the bot
// creates a new rpc provider client connection to the module. The module is kept in the handler
// for event dispatching and module management.
func (b BotAPI) Register(t Ticket, result *string) error {
	module := rpc.NewClientWithCodec(RpcCodecClientWithPort(t.Port))
	if module == nil {
		rlog.Warn("Bot", "Could not register:"+string(t.Name))
		return errors.New("Failed to regsiter module")
	}
	b.bot.Handler.AddModule(ModuleName(strings.ToLower(string(t.Name))), module)
	rlog.Debug("Bot", "Registered "+string(t.Name)+" on port "+t.Port)
	return nil
}
