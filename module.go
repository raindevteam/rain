package rainbot

import (
	"net"
	"net/rpc"
	"os"
	"path/filepath"
	"strings"

	"github.com/sorcix/irc"
)

type CommandFun func(*irc.Message, []string)
type TriggerFun func(*irc.Message) bool
type Listener func(*irc.Message)

type Event string

type Module struct {
	Name      string
	Desc      string
	Master    *rpc.Client
	RpcPort   string
	Listeners map[Event][]Listener
	Commands  map[CommandName]*Command
}

func MakeModule(name string, desc string) *Module {
	m := &Module{
		Name:      name,
		Desc:      desc,
		Listeners: make(map[Event][]Listener),
		Commands:  make(map[CommandName]*Command),
		Master:    rpc.NewClientWithCodec(rpcCodecClient()), // Connect to master
	}
	// Start Provider server
	m.startRpcServer()
	return m
}

func execName() ModuleName {
	return ModuleName(strings.TrimSuffix(filepath.Base(os.Args[0]),
		filepath.Ext(filepath.Base(os.Args[0]))))
}

func (m *Module) startRpcServer() {
	port := getOpenPort()
	if port == "" {
		return // Handle
	}
	rpc.RegisterName(string(execName()), ModuleApi{m})
	provider, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return // Handle
	}
	m.RpcPort = port
	go func() {
		for {
			conn, _ := provider.Accept()
			rpc.ServeCodec(rpcCodecServer(conn))
		}
	}()
}

func (m *Module) GetName() string {
	return m.Name
}

func (m *Module) Say(ch string, text string) {
	result := ""
	m.Master.Call("Master.Send", ch+" :"+text, &result)
}

func (m *Module) RawListener(event Event, l func(*irc.Message)) bool {
	m.Listeners[event] = append(m.Listeners[event], l)
	return true
}

func (m *Module) CommandHook(name CommandName, c *Command) {
	result := ""
	err := m.Master.Call("Master.RegisterCommand", CommandRequest{name, execName()}, &result)
	if err != nil {
		return
	}

	m.Commands[name] = c
}

func (m *Module) Register() (result string, err error) {
	m.Master.Call("Master.Reg", Ticket{m.RpcPort, execName()}, &result)
	m.Master.Call("Master.Loop", "", &result)
	return result, nil
}

type ModuleApi struct {
	M *Module
}

func (mpi ModuleApi) InvokeCommand(d *CommandData, result *string) error {
	mpi.M.Commands[d.Name].Fun(d.Msg, d.Args)
	return nil
}

func (mpi ModuleApi) Dispatch(d *IrcData, result *string) error {
	for _, listener := range mpi.M.Listeners[d.Event] {
		go listener(d.Msg)
	}
	return nil
}
