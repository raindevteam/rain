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

package Thelpers

import (
	"errors"
	"io"
	"net"
	"net/rpc"
)

// An RPCHandler is used to help manage an RPC connection while testing.
type RPCHandler struct {
	RPCServer  net.Listener
	conn       rpc.ServerCodec
	TestPort   string
	registered chan bool
}

// NewRPCHandler returns an intialized RPCHandler.
func NewRPCHandler() *RPCHandler {
	rh := &RPCHandler{
		registered: make(chan bool),
		TestPort:   "23524", // Make this random in future tests
	}

	return rh
}

// CreateRPCServer creates a new RPC server and makes it listen within a goroutine loop.
func (r *RPCHandler) CreateRPCServer(CodecServer func(io.ReadWriteCloser) rpc.ServerCodec) error {
	srv := rpc.NewServer()
	srv.RegisterName("TestMaster", TestAPI{r})

	var err error
	r.RPCServer, err = net.Listen("tcp", ":"+r.TestPort)

	if err != nil {
		return err
	}

	go func(srv *rpc.Server, r *RPCHandler) {
		for {
			conn, err := r.RPCServer.Accept()
			if err != nil {
				break
			}
			r.conn = CodecServer(conn)
			srv.ServeCodec(r.conn)
		}
	}(srv, r)

	return nil
}

// CloseRPCServer will close the RPCHandler's internal server.
func (r *RPCHandler) CloseRPCServer() error {
	err := r.RPCServer.Close()
	return err
}

// A TestAPI is used as a test RPC API for testing.
type TestAPI struct {
	t *RPCHandler
}

// Send should be called as an RPC procedure to ensure the RPC server can successfully receive
// messages.
func (tapi TestAPI) Send(msg string, result *string) error {
	if msg == "Got RPC?" {
		return nil
	}
	return errors.New("Didn't receive right message")
}

// Register will call on the registered chan of an RPCHandler when a module has registered with
// the test master RPC server.
func (tapi TestAPI) Register(t interface{}, result *string) error {
	tapi.t.registered <- true
	return nil
}
