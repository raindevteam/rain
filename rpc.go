package rainbot

import (
    "net/rpc"
    "fmt"
    "net"
    "strconv"
    "io"

    "github.com/ugorji/go/codec"
)

var mh, mhs codec.MsgpackHandle

func getOpenPort() string {
    for i := 65535; i >= 49152; i-- {
        conn, err := net.Listen("tcp", ":" + strconv.Itoa(i))
        if err == nil {
            conn.Close()
            return strconv.Itoa(i)
        }
    }
    return ""
}

func rpcCodecClient() rpc.ClientCodec {
    conn, err := net.Dial("tcp", "localhost:5555")
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return codec.MsgpackSpecRpc.ClientCodec(conn, &mh)
}

func rpcCodecClientWithPort(port string) rpc.ClientCodec {
    conn, err := net.Dial("tcp", "localhost:" + port)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return codec.MsgpackSpecRpc.ClientCodec(conn, &mh)
}

func rpcCodecServer(conn io.ReadWriteCloser) rpc.ServerCodec {
    return codec.MsgpackSpecRpc.ServerCodec(conn, &mhs)
}
