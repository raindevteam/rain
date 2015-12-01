package rainbot

import (
    "fmt"
    "net"
    //"log"
    "net/rpc"
    //"io/ioutil"

    "github.com/RyanPrintup/nimbus"
    //"github.com/ugorji/go/codec"
    "github.com/sorcix/irc"
)

var bt *Bot

type Ticket struct {
    Port string
    Name ModuleName
}

type CommandRequest struct {
    Name   CommandName
    Module ModuleName
}

type Bot struct {
    Client      *nimbus.Client
    ModuleNames []string
    Handler     *Handler
}

func MakeModule(name string, desc string) *Module {
    m := &Module{
        Name: name,
        Desc: desc,
        Listeners: make(map[Event][]Listener),
        Commands: make(map[CommandName]*Command),
        Master: rpc.NewClientWithCodec(rpcCodecClient()), // Connect to master
    }
    // Start Provider server
    m.startRpcServer()
    return m
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
func (b* Bot) moduleRun(name string) {

}

func (b *Bot) LoadModules() {
    b.startRpcServer()
    bt = b
    for _, moduleName := range b.ModuleNames {
        fmt.Println(moduleName)
        b.moduleRun(moduleName)
    }
}

type BotApi struct {
    B *Bot
}

func (bpi BotApi) Send(raw string, result *string) error {
    bpi.B.Client.Send(irc.PRIVMSG, raw)
    return nil
}

func (bpi BotApi) RegisterCommand(cr CommandRequest, result *string) error {
    bpi.B.Handler.AddCommand(cr.Name, cr.Module)
    return nil
}

func (bpi BotApi) Loop(n string, result *string) error {
    c := make(chan bool)
    func(ch chan bool) {

    }(c)
    <- c
    return nil
}

func (bpi BotApi) Reg(t Ticket, result *string) error {
    fmt.Println("============================ Registering: " + t.Name)

    module := rpc.NewClientWithCodec(rpcCodecClientWithPort(t.Port))
    bpi.B.Handler.AddModule(t.Name, module)
    return nil
}
