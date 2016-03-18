package rainbot

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strings"
	"sync"

	"github.com/RyanPrintup/nimbus"
)

/*** Bot internals ***/

// The Channel struct is used to store information about connected channels.
type Channel struct {
	Name  string
	Topic string
	Users map[string]string
	Modes string
}

// NewChannel creates a new channel.
func NewChannel(name string) *Channel {
	channel := &Channel{
		Name:  name,
		Topic: "",
		Users: make(map[string]string),
	}
	return channel
}

// The Module struct holds the name of the module and its corresponding commander helper.
type Module struct {
	Name  string
	Cmder *Commander
}

// NewModule creates a new module.
func NewModule(name string, path string, cmdtype string) *Module {
	m := &Module{
		Name:  name,
		Cmder: NewCommander(name, cmdtype, path),
	}
	return m
}

type Client interface {
	GetNick() string
	GetChannels() []string
	Connect(callback func(error)) error
	Listen()
	Quit() chan error
	Send(raw ...string)
	Say(channel string, text string)
	AddListener(event string, l nimbus.Listener)
	GetListeners(event string) []nimbus.Listener
	Emit(event string, msg *nimbus.Message)
}

// The Bot struct holds the internal nimbus.Client, used to register listeners for irc. ModuleNames
// is used to look up which plugins to start and are eventually passed to the handler. The Handler
// provides management of commands, listeners and triggers. Since listeners act independently of
// each other, a mutex is used to keep bot writes (such as channel updates) synchronised.
type Bot struct {
	Client
	Version  string
	Modules  map[string]*Module
	Channels map[string]*Channel
	Parser   *Parser
	Handler  *Handler
	Mu       sync.Mutex
}

/*** Module specific methods ***/

// SetupModules goes through every module list found in the config and sets them up appropriately.
func (b *Bot) SetupModules(rainConfig *Config) {
	var name, path string

	for name, path = range rainConfig.GoModules {
		b.Modules[name] = NewModule(name, path, "go")
	}
}

// LoadModules starts the bot's rpc master server and then calls moduleRun() on all modules.
func (b *Bot) LoadModules() {
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
	if !b.Handler.ModuleExists(name) {
		return errors.New("Module does not exist")
	}

	c := b.Modules[name].Cmder

	err = b.Handler.SignalKill(name)

	if err != nil {
		s := err.Error()
		fmt.Println(s)
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
		fmt.Println("Could not start: " + name + "(" + err.Error() + ")")
	}
}

/*** IRC specific methods ***/

// RemoveUser will delete a user entry from the given channel. If the user is the bot itself, the
// bot will isntead remove the channel from the bot's channel list.
func (b *Bot) RemoveUser(nick string, channel string) {
	if nick == b.GetNick() {
		delete(b.Channels, strings.ToLower(channel))
		return
	}

	delete(b.Channels[strings.ToLower(channel)].Users, nick)
}

/*** RPC specifc methods ***/

// startRPCServer registers the master consumer for plugins. The master consumer allows plugins to
// communicate with the bot, allowing access to connected channels, users and registered modules.
// Conventionally, it uses a json codec to serve.
func (b *Bot) startRPCServer() {
	rpc.RegisterName("Master", BotAPI{b})
	master, err := net.Listen("tcp", ":5555")

	if err != nil {
		fmt.Println(err)
	}

	// Start accepting connections
	go func() {
		for {
			conn, _ := master.Accept()
			go rpc.ServeCodec(RpcCodecServer(conn))
		}
	}()
}

/*** Bot's master consumer API ***/

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
	fmt.Println("Got request to send " + raw)
	b.bot.Send(nimbus.PRIVMSG, raw)
	return nil
}

// RegisterCommand registers a command from a module with the handler. The command request
// holds the command's name and the module it belongs to (used to signal the module to fire the
// command).
func (b BotAPI) RegisterCommand(cr CommandRequest, result *string) error {
	b.bot.Handler.AddCommand(cr.Name, cr.Module)
	fmt.Println("Added: " + string(cr.Name) + " for module: " + string(cr.Module))
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

// GetConnectedUsers will return a user map (where every user has a IRC rank as a value).
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
		fmt.Println("Could not register:" + t.Name)
		return errors.New("Failed to regsiter module")
	}
	b.bot.Handler.AddModule(t.Name, module)
	fmt.Println("[Registered] " + t.Name)
	fmt.Println("[Port] " + t.Port)
	return nil
}
