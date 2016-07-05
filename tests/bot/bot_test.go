package Tbot

import (
	"strings"
	"testing"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rbot"
	"github.com/raindevteam/rain/tests/_helpers"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type BotSuite struct {
	suite.Suite
	bot  *rbot.Bot
	done chan bool
}

func (s *BotSuite) SetupTest() {
	rcon, err := rbot.ReadConfig("../_helpers/tconfig.json")

	if err != nil {
		s.FailNow("Couldn't read config file")
	}

	s.bot = Thelpers.NewMockBot(rcon)
	s.done = make(chan bool)

	s.bot.AddListener(nimbus.JOIN, func(msg *nimbus.Message) {
		s.bot.Mu.Lock()
		defer s.bot.Mu.Unlock()

		who, _ := s.bot.Parser.ParsePrefix(msg.Prefix)
		where := msg.Args[0]

		if who == s.bot.GetNick() {
			channel := rbot.NewChannel(where)
			s.bot.Channels[strings.ToLower(where)] = channel
			s.done <- true
			return
		}

		channel := s.bot.Channels[strings.ToLower(where)]
		channel.Users[who] = ""
		s.done <- true
	})
}

func (s *BotSuite) send(raw string) {
	s.bot.Send(raw)
	<-s.done
}

func (s *BotSuite) sendBatch(raws []string) {
	toSend := len(raws)
	sent := 0

	allDone := make(chan bool)

	for _, raw := range raws {
		go func(raw string) {
			s.bot.Send(raw)
			<-s.done
			sent++
			if sent == toSend {
				allDone <- true
			}
		}(raw)
	}

	<-allDone
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *BotSuite) TestNewChannel() {
	channel := rbot.NewChannel("#rainbottest")
	s.NotNil(channel)
	s.Equal("#rainbottest", channel.Name, "Name should be #rainbottest")
	s.Equal("", channel.Topic, "Topic should not be set")
	s.Empty("", channel.Users, "Users should be empty")
}

// TODO: TestNewModule (First test commander)
// TODO: TestSetupModules(First test commander and modules)
// TODO: TestLoadModules(First test commander and modules)
// TODO: TestLoadModules(First test commander and modules)

func (s *BotSuite) TestConnect() {
	s.bot.Connect()
}

func (s *BotSuite) TestJoinListener() {
	s.sendBatch([]string{"JOIN #rainbot", "JOIN #snowybottest"})

	chanSnow := s.bot.Channels["#snowybottest"]
	s.NotNil(chanSnow)
	s.Equal("#snowybottest", chanSnow.Name)

	chanRain := s.bot.Channels["#rainbot"]
	s.NotNil(chanRain)
	s.Equal("#rainbot", chanRain.Name)
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestBotSuite(t *testing.T) {
	suite.Run(t, new(BotSuite))
}

/**************************************************************************************************/
