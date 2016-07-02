package Tbot

import (
	"testing"

	"github.com/wolfchase/rainbot/bot"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type ParserSuite struct {
	suite.Suite
	Parser *rainbot.Parser
}

func (s *ParserSuite) SetupTest() {
	s.Parser = rainbot.NewParser(".")
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *ParserSuite) TestNewParser() {
	s.NotNil(s.Parser, "Parser created")
}

func (s *ParserSuite) TestIsCommand() {
	s.True(s.Parser.IsCommand(".yes"))
	s.False(s.Parser.IsCommand("no, not a command"))
	s.True(s.Parser.IsCommand(".yes I am a command"))
	s.False(s.Parser.IsCommand(";nope, I am not a command either"))
}

func (s *ParserSuite) TestParseCommand() {
	var name string
	var args []string

	name, args = s.Parser.ParseCommand(".say")
	s.Equal("say", name)
	s.Empty(args)

	name, args = s.Parser.ParseCommand(".echo I am broot")
	s.Equal("echo", name)
	s.Exactly([]string{"I", "am", "broot"}, args)
}

func (s *ParserSuite) TestParsePrefix() {
	var who, host string

	who, host = s.Parser.ParsePrefix("NimBot")
	s.Equal("NimBot", who)
	s.Empty(host)

	who, host = s.Parser.ParsePrefix("NimBot!this.host")
	s.Equal("NimBot", who)
	s.Equal("this.host", host)
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

/**************************************************************************************************/
