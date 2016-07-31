// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rlog"
)

// The ServerInfo struct holds information about the IRC server.
type ServerInfo struct {
	Host     string
	Port     string
	Channels []string
}

// UserInfo holds information pertaining to the bot.
type UserInfo struct {
	Nick     string
	RealName string
	UserName string
	Modes    string
}

// CommandInfo holds information needed for parsing commands
type CommandInfo struct {
	Prefix string `yaml:"commandprefix"`
}

// ModuleInfo holds information for managing modules
type ModuleInfo struct {
	//      Modules -> Type -> Module
	// i.e. Modules["js"]["umbrella"] = "/modules/js"
	Modules       map[string]map[string]string
	DefaultRoutes map[string]string `yaml:"default-routes"`
	GlobalOptions map[string]string `yaml:"global-options"`
}

// Config keeps together all modular components of configuration
type Config struct {
	Server  ServerInfo
	User    UserInfo
	Command CommandInfo
	Module  ModuleInfo
}

// ReadConfigFile takes a path to a config file and parses it.
func ReadConfigFile(path string) (*Config, error) {
	file, _ := ioutil.ReadFile(path)
	return ReadConfig(string(file))
}

func check(err error) {
	if err != nil {
		rlog.Error("Config", "Could not correctly parse the configuration, check your syntax :::")
		panic(err)
	}
}

// ReadConfig parses the passed configuration string
func ReadConfig(configstr string) (*Config, error) {
	si := ServerInfo{}
	ui := UserInfo{}
	ci := CommandInfo{}
	mi := ModuleInfo{}

	err := yaml.Unmarshal([]byte(configstr), &si)
	check(err)
	err = yaml.Unmarshal([]byte(configstr), &ui)
	check(err)
	err = yaml.Unmarshal([]byte(configstr), &ci)
	check(err)
	err = yaml.Unmarshal([]byte(configstr), &mi)
	check(err)

	config := &Config{
		si,
		ui,
		ci,
		mi,
	}

	return config, nil
}

// GetNimbusConfig takes an rbot.Config and uses it to create a nimbus.Config
func GetNimbusConfig(rconf *Config) (nconf *nimbus.Config) {
	nconf = &nimbus.Config{
		Port:     rconf.Server.Port,
		Channels: rconf.Server.Channels,
		RealName: rconf.User.RealName,
		UserName: rconf.User.UserName,
		Password: "",
		Modes:    rconf.User.Modes,
	}

	return nconf
}
