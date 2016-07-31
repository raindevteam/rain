// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

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
		// NewProcessManager("pymod", "py", "../tests/modules"),
		NewProcessManager("jsmod", "js", "../tests/modules"),
	}

	for _, pm := range pms {
		res := pm.Recompile()
		if res != nil && res.Err != nil {
			t.Fatalf("Failed to recompile %s: %s", pm.Name, res.Err)
		}

		done := make(chan *Result)
		go func(pm *ProcessManager) {
			done <- pm.Start("5555")
		}(pm)

		select {
		case res := <-done:
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
