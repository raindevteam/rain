package rbot

import (
	"net"
	"net/rpc"
	"os"
	"testing"
)

var RPCServer net.Listener

type MasterAPI struct {
}

// The following two rpc procedures are only defined so that the test modules used here have
// something to delegate their calls to. We plan on fully testing rpc procedures with bapi.go

func (r MasterAPI) RegisterCommand(cr interface{}, result *string) error {
	return nil
}

func (r MasterAPI) Register(t interface{}, result *string) error {
	return nil
}

func SetupTest() error {
	srv := rpc.NewServer()
	srv.RegisterName("Master", MasterAPI{})

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
		err := pm.Recompile()
		// In the future, implement a last command status struct
		if err != nil {
			t.Fatalf("Failed to recompile %s: %s", pm.Name, err.Error())
		}

		err = pm.Start()
		if err != nil {
			t.Fatalf("Failed to start %s: %s", pm.Name, err.Error())
		}

		err = pm.Kill()
		if err != nil {
			t.Fatalf("Failed to kill %s: %s", pm.Name, err.Error())
		}
	}
}

func tearDown() error {
	return RPCServer.Close()
}
