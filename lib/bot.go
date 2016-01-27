package rainbot

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"sync"
	"strings"

	"github.com/RyanPrintup/nimbus"
	"github.com/sorcix/irc"
)

// Tickets are used to connect to a provider connection via
// module registration.
type Ticket struct {
	Port string
	Name ModuleName
}

// A CommandRequest is used to register commands via the handler.
type CommandRequest struct {
	Name   CommandName
	Module ModuleName
}

type TriggerRequest struct {
	Name   ModuleName
	Event  Event
}

type Channel struct {
	Name  string
	Topic string
	Users map[string]string
	Modes string
}

func NewChannel(name string) *Channel {
	channel := &Channel{
		Name: name,
		Topic: "",
		Users: make(map[string]string),
	}
	return channel
}

// The Bot struct holds the internal nimbus.Client, used to register
// listeners for irc. ModuleNames is used to look up which plugins to start.
// The Handler provides management of commands, listeners and triggers.
type Bot struct {
	*nimbus.Client
	Version     string
	ModuleNames []string
	Channels    map[string]*Channel
	Parser      *Parser
	Handler     *Handler
	Mu          sync.Mutex
}

func (b *Bot) RemoveUser(nick string, channel string) {
	if nick == b.Nick {
		delete(b.Channels, strings.ToLower(channel))
		return
	}

	delete(b.Channels[strings.ToLower(channel)].Users, nick)
}

// startRpcServer registers the master consumer for plugins.
// The master consumer allows plugins to communicate with the
// bot, allowing access to connected channels, users and registered
// modules. Conventionally, it uses a json codec to serve.
func (b *Bot) startRpcServer() {
	rpc.RegisterName("Master", BotApi{b})
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

// moduleRun starts a plugin as a provider service. This allows
// the bot to dispatch events outbound to the module. The module
// can then communicate back via the master rpc server.
func (b *Bot) moduleRun(name string) {
	err := exec.Command(name).Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// It firstly starts the master consumer server for the bot.
// It then starts every module via exec.
func (b *Bot) LoadModules() {
	b.startRpcServer()
	for _, moduleName := range b.ModuleNames {
		b.moduleRun(moduleName)
	}
}

// Api served via the bot's master consumer connection
type BotApi struct {
	B *Bot
}

// Sends a raw irc message via the nimbus client.
func (bpi BotApi) Send(raw string, result *string) error {
	bpi.B.Client.Send(irc.PRIVMSG, raw)
	return nil
}

// Register a command from a module using a command request. The command request
func (bpi BotApi) RegisterCommand(cr CommandRequest, result *string) error {
	bpi.B.Handler.AddCommand(cr.Name, cr.Module)
	fmt.Println(string(cr.Name) + " " + string(cr.Module))
	return nil
}

func (bpi BotApi) RegisterTrigger(tr TriggerRequest, result *string) error {
	listeners := bpi.B.GetListeners(string(tr.Event))
	if len(listeners) == 0 {
		bpi.B.AddListener(string(tr.Event), func (msg *nimbus.Message) {
			bpi.B.Handler.Fire(msg, tr.Event)
		})
	}
	bpi.B.Handler.AddTrigger(tr.Name, tr.Event)
	return nil
}

func (bpi BotApi) GetVersion(mName string, result *string) error {
	*result = bpi.B.Version
	return nil
}

// Keeps a module alive
func (bpi BotApi) Loop(n string, result *string) error {
	c := make(chan bool)
	func(ch chan bool) {

	}(c)
	<-c
	return nil
}

// Registers a module with the bot. With the given port number in
// in the Ticket, the bot creates a new rpc provider client connection
// to the module. The name is kept in the handler for event dispatching
// and module management.
func (bpi BotApi) Reg(t Ticket, result *string) error {
	fmt.Println("================== Registering: " + t.Name)
	module := rpc.NewClientWithCodec(RpcCodecClientWithPort(t.Port))
	bpi.B.Handler.AddModule(t.Name, module)
	return nil
}
