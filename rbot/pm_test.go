package rbot

import (
	"net"
	"net/rpc"
	"os"
	"testing"
)

var (
	RPCServer  net.Listener
	registered = make(chan bool)
)

type MasterAPI struct {
	reg chan bool
}

// The following two rpc procedures are only defined so that the test modules used here have
// something to delegate their calls to. We plan on fully testing rpc procedures with bapi.go

func (r MasterAPI) RegisterCommand(cr interface{}, result *string) error {
	return nil
}

func (r MasterAPI) Register(t interface{}, result *string) error {
	r.reg <- true
	return nil
}

func SetupTest() error {
	srv := rpc.NewServer()
	srv.RegisterName("Master", MasterAPI{registered})

	var err error
	RPCServer, err = net.Listen("tcp", ":5555")

	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := RPCServer.Accept()
			if err != nil {
				break
			}
			go srv.ServeCodec(RpcCodecServer(conn))
		}
	}()
	return nil
}

func TestMain(m *testing.M) {
	err := SetupTest()
	if err != nil {
		os.Exit(1)
	}
	ret := m.Run()
	err = tearDown()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(ret)
}

func TestProcessManagers(t *testing.T) {
	pms := []*ProcessManager{
		NewProcessManager("gomod", "go", "github.com/raindevteam/rain/tests/modules"),
		NewProcessManager("pymod", "py", "../tests/modules"),
		NewProcessManager("jsmod", "js", "../tests/modules"),
	}

	for _, pm := range pms {
		res := pm.Recompile()
		if res != nil && res.Err != nil {
			t.Fatalf("Failed to recompile %s: %s", pm.Name, res.Output)
		}

		cmd := pm.Start()
		select {
		case res := <-cmd:
			if res.Err != nil {
				t.Fatalf("Module exited prematurely\n ::::\n%s", res.Output)
			} else {
				t.Fatal("Module exited but there was no error, perhaps it is broken")
			}
		case <-registered:
		}

		pm.Kill()
	}
}

func tearDown() error {
	return RPCServer.Close()
}
