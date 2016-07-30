// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU Affero General Public License as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License along with this program.
// If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

/*
 * TODO: We need to implement better error handling for config.go. Also the tests here are
 * pretty useless if the configuration parsing fails since we don't ever say what really goes
 * wrong.
 */

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rain"
)

var nconfig = &nimbus.Config{
	Port:     "6667",
	Channels: []string{"#Your", "#Channels"},
	RealName: "Wolfstein Jr. II",
	UserName: "wolfstein",
	Password: "",
	Modes:    "+B",
}

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

var defaultRoutes = map[string]string{
	"js": "modules/js",
	"py": "modules/py",
}

var globalOptions = map[string]string{
	"js": "npm",
	"py": "noload",
	"go": "noload",
}

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

func checkCommandInfo(config *Config) (check bool) {
	check = config.Command.Prefix == "."
	return
}

func checkModuleInfo(config *Config) (check bool) {
	check = reflect.DeepEqual(config.Module.Modules, modules)
	check = reflect.DeepEqual(config.Module.DefaultRoutes, defaultRoutes) && check
	check = reflect.DeepEqual(config.Module.GlobalOptions, globalOptions) && check
	return
}

func checkConfig(config *Config) (check bool) {
	check = checkServerInfo(config)
	check = checkUserInfo(config) && check
	check = checkCommandInfo(config) && check
	check = checkModuleInfo(config) && check
	return
}

func TestReadConfigFile(t *testing.T) {
	conf, _ := ReadConfigFile("../tests/test_config.yaml")
	checks := checkConfig(conf)
	if checks != true {
		fmt.Println(conf)
		t.Error("Config did not parse correctly")
	}
}

func TestReadConfig(t *testing.T) {
	config, _ := ReadConfig(rain.ExampleConfig)
	if checkConfig(config) != true {
		t.Error("Config did not parse correctly")
	}
}

func TestGetNimbusConfig(t *testing.T) {
	rconfig, _ := ReadConfig(rain.ExampleConfig)
	config := GetNimbusConfig(rconfig)
	checks := reflect.DeepEqual(config, nconfig)
	if checks != true {
		t.Error("Could create a proper nimbus config")
	}
}
