package Tbot

import (
	"testing"

	"github.com/wolfchase/rainbot/bot"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type CommanderSuite struct {
	suite.Suite
	GoModCmder *rainbot.Commander
}

func (s *CommanderSuite) SetupTest() {
	s.GoModCmder = rainbot.NewCommander("gomod", "go",
		"github.com/wolfchase/rainbot/_tests/_helpers/modules")
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

	s.True(s.Registered, "The module successfully registered upon starting")
	s.Registered = false

	err = s.GoModCmder.Kill()
	s.Nil(err, "No error from killing the module")
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestCommanderSuite(t *testing.T) {
	suite.Run(t, new(CommanderSuite))
}

/**************************************************************************************************/
