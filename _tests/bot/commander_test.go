package RBTbot

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
	GoModCmder := rainbot.NewCommander("GoMod", "go",
		"github.com/wolfchase/rainbot/_tests/_helpers/modules")
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *CommanderSuite) TestRecompile() {
	err := s.GoModCmder.Recompile()
	s.Nil(err)
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestCommanderSuite(t *testing.T) {
	suite.Run(t, new(CommanderSuite))
}

/**************************************************************************************************/
