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
	"fmt"
	"net/rpc"
	"os"
	"testing"

	"github.com/raindevteam/rain/tests"
)

const (
	mname = "gomod"
)

var (
	pm = NewProcessManager(mname, "go", "github.com/raindevteam/rain/tests/modules")
	r  = Thelpers.NewRPCHandler()
)

func HandleModuleExit(name string, result *Result) {
	fmt.Printf("%s has exited, test failing.\n::::\n%s\n", mname, result.Output)
	os.Exit(1)
}

func HandlerWithModule(r *Thelpers.RPCHandler) (*Handler, *rpc.Client) {
	h := NewHandler()

	<-r.Registered
	client, err := RpcCodecClientWithPort(r.ModPort)
	if err != nil {
		fmt.Printf("Could not establish an RPC client: " + err.Error())
		os.Exit(1)
	}

	module := rpc.NewClientWithCodec(client)
	if module == nil {
		fmt.Printf("Could not register:" + mname)
		os.Exit(1)
	}

	h.AddModule(mname, module)

	return h, module
}

func TestMain(m *testing.M) {
	err := r.CreateRPCServer(RpcCodecServer)
	defer r.CloseRPCServer()

	if err != nil {
		os.Exit(1)
	}

	go func(name string, pm *ProcessManager) {
		result := pm.Start(r.TestPort)
		HandleModuleExit(mname, result)
	}(mname, pm)

	ret := m.Run()
	os.Exit(ret)
}

func TestAddModule(t *testing.T) {
	h, _ := HandlerWithModule(r)
	//defer c.Close()

	_, ok := h.Modules["gomod"]
	if !ok {
		t.Fatalf("Module wasn't added to handler")
	}
}
