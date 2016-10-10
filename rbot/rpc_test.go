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
	"net/rpc"
	"testing"

	"github.com/raindevteam/rain/tests"
)

// Check if we can create and close an RPC server successfully using the RPC server codec.
func TestRpcCodecServer(t *testing.T) {
	r := Thelpers.NewRPCHandler()
	err := r.CreateRPCServer(RpcCodecServer)

	if err != nil {
		t.Fatalf("Could not create rpc server! %s:", err.Error())
	}

	err = r.CloseRPCServer()
	if err != nil {
		t.Fatalf("Could not close RPC server! %s:", err.Error())
	}
}

// Check if we can create a client without erroring.
func TestRpcCodecClientWithPort(t *testing.T) {
	r := Thelpers.NewRPCHandler()
	r.CreateRPCServer(RpcCodecServer)
	defer r.CloseRPCServer()

	RPCClient, err := RpcCodecClientWithPort(r.TestPort)
	if RPCClient != nil {
		defer RPCClient.Close()
	}

	if err != nil {
		t.Fatalf("Could not create an RPC client! %s:", err.Error())
	}

}

// Let's double check and make sure our server and client can speak to each other
func TestRpcCodecServerClientComm(t *testing.T) {
	r := Thelpers.NewRPCHandler()
	r.CreateRPCServer(RpcCodecServer)
	defer r.CloseRPCServer()

	RPCCodec, err := RpcCodecClientWithPort(r.TestPort)
	if err != nil {
		t.Fatalf("Error while trying to create RPCClientCodec: %s", err.Error())
	}

	RPCClient := rpc.NewClientWithCodec(RPCCodec)
	defer RPCClient.Close()

	var result string
	err = RPCClient.Call("TestMaster.Send", "Got RPC?", &result)
	if err != nil {
		t.Fatalf("Error while trying to send RPC message: %s", err.Error())
	}
}
