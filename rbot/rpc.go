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
