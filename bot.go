package rainbot

import (
	"fmt"
	"net"
	"net/rpc"
	"os/exec"

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

// The Bot struct holds the internal nimbus.Client, used to register
// listeners for irc. ModuleNames is used to look up which plugins to start.
// The Handler provides management of commands, listeners and triggers.
type Bot struct {
	Client      *nimbus.Client
	ModuleNames []string
	Handler     *Handler
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
			rpc.ServeCodec(rpcCodecServer(conn))
		}
	}()
}

// moduleRun starts a plugin as a provider service. This allows
// the bot to dispatch events outbound to the module. The module
// can then communicate back via the master rpc server.
func (b *Bot) moduleRun(name string) {
	out, err := exec.Command(name).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
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
	module := rpc.NewClientWithCodec(rpcCodecClientWithPort(t.Port))
	bpi.B.Handler.AddModule(t.Name, module)
	return nil
}
