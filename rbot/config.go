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
