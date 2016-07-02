package Tbot

import (
	"net"
	"net/rpc"
	"testing"

	"github.com/wolfchase/rainbot/bot"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type CommanderSuite struct {
	suite.Suite
	GoModCmder *rainbot.Commander
	PyModCmder *rainbot.Commander
	JsModCmder *rainbot.Commander
	RpcServer  net.Listener
}

type MasterAPI struct {
	s *CommanderSuite
}

func (r MasterAPI) RegisterCommand(cr rainbot.CommandRequest, result *string) error {
	return nil
}

func (r MasterAPI) Register(t rainbot.Ticket, result *string) error {
	return nil
}

func (s *CommanderSuite) SetupSuite() {
	s.GoModCmder = rainbot.NewCommander("gomod", "go",
		"github.com/wolfchase/rainbot/_tests/_helpers/modules")
	s.PyModCmder = rainbot.NewCommander("pymod", "py", "../_helpers/modules")
	s.JsModCmder = rainbot.NewCommander("jsmod", "js", "../_helpers/modules")

	rpc.RegisterName("Master", MasterAPI{s})

	var err error
	s.RpcServer, err = net.Listen("tcp", ":5555")

	if err != nil {
		s.FailNow("Failed to create RPC server")
	}
	go func() {
		for {
			conn, err := s.RpcServer.Accept()
			if err != nil {
				break
			}
			go rpc.ServeCodec(rainbot.RpcCodecServer(conn))
		}
	}()
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *CommanderSuite) TestGoCommander() {
	var err error

	err = s.GoModCmder.Recompile()
	// In the future, implement a last command status struct
	s.Nil(err, "No error while recompiling")

	err = s.GoModCmder.Start()
	s.Nil(err, "No error from starting the module")

	err = s.GoModCmder.Kill()
	s.Nil(err, "No error from killing the module")
}

func (s *CommanderSuite) TestPyCommander() {
	var err error

	err = s.PyModCmder.Recompile()
	// In the future, implement a last command status struct
	s.Nil(err, "No error while recompiling")

	err = s.PyModCmder.Start()
	s.Nil(err, "No error from starting the module")

	err = s.PyModCmder.Kill()
	s.Nil(err, "No error from killing the module")
}

func (s *CommanderSuite) TestJsCommander() {
	var err error

	err = s.JsModCmder.Recompile()
	// In the future, implement a last command status struct
	s.Nil(err, "No error while recompiling")

	err = s.JsModCmder.Start()
	s.Nil(err, "No error from starting the module")

	err = s.JsModCmder.Kill()
	s.Nil(err, "No error from killing the module")
}

func (s *CommanderSuite) TearDownSuite() {
	s.RpcServer.Close()
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestCommanderSuite(t *testing.T) {
	suite.Run(t, new(CommanderSuite))
}

/**************************************************************************************************/