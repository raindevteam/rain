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
	"net"
	"net/rpc"

	"github.com/raindevteam/rain/rbot"
)

// An RPCHandler is used to help manage an RPC connection while testing.
type RPCHandler struct {
	RPCServer  net.Listener
	conn       rpc.ServerCodec
	TestPort   string
	registered chan bool
}

// NewRPCHandler returns an intialized RPCHandler.
func (r *RPCHandler) NewRPCHandler() *RPCHandler {
	rh := &RPCHandler{
		registered: make(chan bool),
	}

	return rh
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
				break
			}
			r.conn = rbot.RpcCodecServer(conn)
			srv.ServeCodec(r.conn)
		}
	}()

	return nil
}

func (r *RPCHandler) CloseRPCServer() error {
	err := r.RPCServer.Close()
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

func (tapi TestAPI) Register(t interface{}, result *string) error {
	return nil
}
