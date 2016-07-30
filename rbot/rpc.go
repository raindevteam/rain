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
	"io"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RpcCodecClientWithPort(port string) (rpc.ClientCodec, error) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClientCodec(conn), nil
}

func RpcCodecServer(conn io.ReadWriteCloser) rpc.ServerCodec {
	return jsonrpc.NewServerCodec(conn)
}
