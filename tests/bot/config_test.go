package Tbot

import (
	"testing"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/bot"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type ConfigSuite struct {
	suite.Suite
	testConfig string
	trcon      *rbot.Config
	tncon      *nimbus.Config
}

func (s *ConfigSuite) SetupTest() {
	s.testConfig = "../_helpers/tconfig.json"

	s.trcon = &rbot.Config{
		Host:    "irc.canternet.org",
		Port:    "6667",
		Channel: []string{"#snowybottest"},

		Nick:     "SnowBot",
		RealName: "RainBotGo",
		UserName: "RainBotGo",
		Modes:    "+B",

		CmdPrefix: ".",

		GoModules: map[string]string{
			"rcore": "github.com/wolfchase",
		},
	}

	s.tncon = &nimbus.Config{
		Port:     "6667",
		Channels: []string{"#snowybottest"},
		RealName: "RainBotGo",
		UserName: "RainBotGo",
		Password: "",
		Modes:    "+B",
	}
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *ConfigSuite) TestReadConfig() {
	config, err := rbot.ReadConfig(s.testConfig)

	s.Nil(err)
	s.Exactly(s.trcon, config, "Parsed tconfig.json")
}

/*
func TestGetConfigs(t *testing.T) {
	ncon, rcon, err := rainbot.GetConfigs(TestConfig)

	if err != nil {
		t.Fatalf("Error while reading config file: ", err)
	}

	if !reflect.DeepEqual(rcon, TestRainBotConfig) {
		t.Error("Could not parse config properly: ", TestConfig)
		t.Logf("Actual: %#v", rcon)
		t.Logf("Expected: %#v", TestRainBotConfig)
	}

	if !reflect.DeepEqual(ncon, TestNimbusConfig) {
		t.Error("Could not parse config properly: ", TestConfig)
		t.Logf("Actual: %#v", ncon)
		t.Logf("Expected: %#v", TestNimbusConfig)
	}
}
*/

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

/**************************************************************************************************/
