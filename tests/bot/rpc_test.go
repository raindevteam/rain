package Tbot

import (
	"errors"
	"net"
	"net/rpc"
	"testing"

	"github.com/wolfchase/rainbot/bot"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type RpcSuite struct {
	suite.Suite
	RpcServer  net.Listener
	RpcClient  *rpc.Client
	TestPort   string
	GotMessage bool
}

func (s *RpcSuite) SetupTest() {
	s.GotMessage = false
	s.TestPort = "5556"
}

type TestAPI struct {
	s *RpcSuite
}

func (r TestAPI) Send(msg string, result *string) error {
	if msg == "Got RPC?" {
		r.s.GotMessage = true
		return nil
	}
	return errors.New("Didn't receive right message")
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *RpcSuite) Test1RpcCodecServerInit() {
	rpc.RegisterName("TestMaster", TestAPI{s})

	var err error
	s.RpcServer, err = net.Listen("tcp", ":"+s.TestPort)

	if err != nil {
		s.FailNow("Failed to create RPC server")
	}

	go func() {
		for {
			conn, err := s.RpcServer.Accept()
			if err != nil {
				break
			}
			go rpc.ServeCodec(rbot.RpcCodecServer(conn))
		}
	}()
}

func (s *RpcSuite) Test2RpcCodecClientConnect() {
	s.RpcClient = rpc.NewClientWithCodec(rbot.RpcCodecClientWithPort(s.TestPort))
	s.NotNil(s.RpcClient, "Client Created")
}

func (s *RpcSuite) Test3RpcCodecServerClientComm() {
	var result string
	err := s.RpcClient.Call("TestMaster.Send", "Got RPC?", &result)
	s.Nil(err)
	s.True(s.GotMessage, "Communication between rpc server and client successful")
}

func (s *RpcSuite) TearDownSuite() {
	s.RpcServer.Close()
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestRpcSuite(t *testing.T) {
	suite.Run(t, new(RpcSuite))
}

/**************************************************************************************************/
