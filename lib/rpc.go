package rainbot

import (
	"fmt"
	"io"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RpcCodecClient() rpc.ClientCodec {
	conn, err := net.Dial("tcp", "localhost:5555")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonrpc.NewClientCodec(conn)
}

func RpcCodecClientWithPort(port string) rpc.ClientCodec {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonrpc.NewClientCodec(conn)
}

func RpcCodecServer(conn io.ReadWriteCloser) rpc.ServerCodec {
	return jsonrpc.NewServerCodec(conn)
}
