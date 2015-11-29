package rainbot

import (
    "net/rpc"
    "net/rpc/jsonrpc"
    "fmt"

    "github.com/RyanPrintup/nimbus"
    "github.com/natefinch/pie"
    "github.com/sorcix/irc"
)

type ModuleITF interface {
	Init()
    GetName() string
    GetBot() *nimbus.Client
    AddCommand()
}

type Module struct {
	Name string
	Desc string
    Rpc  *rpc.Client
    Bot  *nimbus.Client
	Listeners []func(*irc.Message)
}

func (m *Module) Init() {
    m.Rpc = pie.NewConsumerCodec(jsonrpc.NewClientCodec)
    fmt.Println("made thing")
}

func (m *Module) GetName() string {
    return m.Name
}

func (m *Module) AddListener(l func(*irc.Message)) {
    fmt.Println("adding listener")
    m.Listeners = append(m.Listeners, l)
}

func (m *Module) Load() (string, error) {
    m.Init()
    if m.Rpc == nil {
        fmt.Println("borken")
    }
    result := ""
    err := m.Rpc.Call("Host.Register", m, &result)
    if err != nil {
        fmt.Println(err)
    }
    return result, err
}
