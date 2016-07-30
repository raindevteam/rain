// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

import (
	"errors"
	"net"
	"net/rpc"
	"testing"
)

type RPCHandler struct {
	RPCServer net.Listener
	conn      rpc.ServerCodec
	done      chan bool
	TestPort  string
	stop      bool
	GotRPC    bool
}

func (r *RPCHandler) SetupTest() {
	r.stop = false
	r.GotRPC = false
	r.done = make(chan bool)
	r.TestPort = "5556"
}

// TODO: Create separate function to handle erroring
func (r *RPCHandler) CreateRPCServer() error {
	srv := rpc.NewServer()
	srv.RegisterName("TestMaster", TestAPI{r})

	var err error
	r.RPCServer, err = net.Listen("tcp", ":"+r.TestPort)

	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := r.RPCServer.Accept()
			if err != nil {
				r.done <- true
				return
			}
			r.conn = RpcCodecServer(conn)
			srv.ServeCodec(r.conn)
		}
	}()
	return nil
}

func (r *RPCHandler) CloseRPCServer() error {
	r.stop = true
	err := r.RPCServer.Close()
	<-r.done
	return err
}

type TestAPI struct {
	t *RPCHandler
}

func (tapi TestAPI) Send(msg string, result *string) error {
	if msg == "Got RPC?" {
		tapi.t.GotRPC = true
		return nil
	}
	return errors.New("Didn't receive right message")
}

// Check if we can create and close an RPC server successfully using the RPC server codec.
func TestRpcCodecServer(t *testing.T) {
	r := RPCHandler{}
	r.SetupTest()

	err := r.CreateRPCServer()
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
	r := RPCHandler{}
	r.SetupTest()
	r.CreateRPCServer()
	defer r.CloseRPCServer()

	RPCClient, err := RpcCodecClientWithPort(r.TestPort)
	defer RPCClient.Close()
	if err != nil {
		t.Fatalf("Could not create an RPC client! %s:", err.Error())
	}
}

// Let's double check and make sure our server and client can speak to each other
func TestRpcCodecServerClientComm(t *testing.T) {
	r := RPCHandler{}
	r.SetupTest()
	r.CreateRPCServer()
	defer r.CloseRPCServer()

	RPCCodec, _ := RpcCodecClientWithPort(r.TestPort)
	RPCClient := rpc.NewClientWithCodec(RPCCodec)
	defer RPCClient.Close()

	var result string
	err := RPCClient.Call("TestMaster.Send", "Got RPC?", &result)
	if err != nil {
		t.Fatalf("Error while trying to send RPC message: %s", err.Error())
	}

	if !r.GotRPC {
		t.Fatalf("Could not send correct message over RPC")
	}
}
