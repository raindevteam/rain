package rbot

import (
	"reflect"
	"testing"
)

var modules = map[string]map[string]string{
	"js": {
		"umbrella": "modulesOverHere/js:npm noload",
		"mlpshow":  ".",
	},
	"py": {
		"echo":  "../andSomeMoreOverHere/py:pip",
		"ltest": ".",
	},
	"go": {
		"raincore": "github.com/raindevteam",
	},
}

var default

func checkServerInfo(config *Config) (check bool) {
	check = config.Server.Host == "irc.examplehost.net"
	check = config.Server.Port == "6667" && check
	check = config.Server.Channels[0] == "#Your" && check
	check = config.Server.Channels[1] == "#Channels" && check
	return
}

func checkUserInfo(config *Config) (check bool) {
	check = config.User.Nick == "MyBot"
	check = config.User.RealName == "Wolfstein Jr. II" && check
	check = config.User.UserName == "wolfstein" && check
	check = config.User.Modes == "+B" && check
	return
}

func checkModuleInfo(config *Config) (check bool) {
	check = reflect.DeepEqual(config.Module.Modules, modules)
	//check =
	return
}

func checkConfig(config *Config) (check bool) {
	check = checkServerInfo(config)
	check = checkUserInfo(config) && check
	check = checkModuleInfo(config) && check
	return
}

func TestReadConfig(t *testing.T) {
	config, _ := ReadConfig(ExampleConfig)
	if checkConfig(config) != true {
		t.Error("Config did not parse correctly")
	}
}
