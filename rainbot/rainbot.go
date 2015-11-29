package rainbot

import (
    "runtime"
    "fmt"
    "os"
    "log"
    //"net/rpc/jsonrpc"

    "github.com/RyanPrintup/nimbus"
    //"github.com/sorcix/irc"
    "github.com/natefinch/pie"
)

type Bot struct {
    C *nimbus.Client
    Modules []string
}

func MakeModule(name string, desc string) *Module {
    m := &Module{
        Name: name,
        Desc: desc,
    }
    return m
}

func (b *Bot) LoadModules() {
    for _, moduleName := range b.Modules {
        fmt.Println(moduleName)
        moduleRPC(moduleName)
    }
}

func moduleRPC(name string) {
    path := name
    if runtime.GOOS == "windows" {
        path = path + ".exe"
    }

    s, err := pie.StartConsumer(os.Stdout, path)
    if err != nil {
		log.Fatalf("failed to start consumer: %s", err)
	}

    if err := s.RegisterName("Host", ModuleApi{}); err != nil {
		log.Fatalf("failed to register Host: %s", err)
	}
    s.Serve()
}

type ModuleApi struct {
}

func (ModuleApi) Register(m *Module, result *string) error {
    listeners := m.Listeners
    if len(listeners) > 0 {
        //for _, listener := range listeners {
        //    mpi.B.C.AddListener(irc.PRIVMSG, listener) //yuck
        //}
        fmt.Println("found listener")
    }
    *result = "good"
    return nil
}
