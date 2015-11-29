package rainbot

import (
    "net/rpc"
    //"net/rpc/jsonrpc"
    "fmt"
    "io"

    "github.com/ugorji/go/codec"
    "github.com/RyanPrintup/nimbus"
    "github.com/natefinch/pie"
    "github.com/sorcix/irc"
)

var (
    mh codec.MsgpackHandle
    bh codec.BincHandle
    cb codec.CborHandle
)

func msgpackRpcCodec(pipe io.ReadWriteCloser) rpc.ClientCodec {
    return codec.MsgpackSpecRpc.ClientCodec(pipe, &mh)
}

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
    Bot  *Bot
	Listeners []func(*irc.Message)
}

func (m *Module) Init() {
    m.Rpc = pie.NewConsumerCodec(msgpackRpcCodec)
    fmt.Println("made thing")
}

func (m *Module) GetName() string {
    return m.Name
}

func (m *Module) AddListener(l func(*irc.Message)) {
    fmt.Println("adding listener")
    m.Listeners = append(m.Listeners, l)
}

func (m *Module) Load() (result string, err error) {
    m.Init()
    err = m.Rpc.Call("Host.Register", m, &result)
    if err != nil {
        fmt.Println(err)
        fmt.Println("print the error")
    } else {
        fmt.Println("all went well")
    }
    return "good", err
}
