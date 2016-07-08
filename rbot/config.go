package rbot

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rlog"
)

var ExampleConfig = `# Rain config
# 0.6.0-alpha.1

# Server information
host : irc.examplehost.net
port : 6667

# Channel Information
# Channels listed here will be automatically joined upon server connect
channels :
  - "#Your"
  - "#Channels"

# User information
nick     : MyBot
realname : Wolfstein Jr. II
username : wolfstein
modes    : +B

# Command Information
# The command prefix is used to denote commands from IRC users
commandprefix : .

# Modules
# These are all optional. A '.' as a module path means its located in the default path. You may 
# set a specific path if you wish. You may specify extra options for the module as a comma separated
# list by putting a ':' after its route. We currently only support a few options.
modules :
  # JavaScript Modules
  js :
    umbrella : modulesOverHere/js:npm noload
    mlpshow  : .

  # Python Modules
  py :
    echo  : ../andSomeMoreOverHere/py:pip
    ltest : .

  # Go Modules
  go :
    raincore : github.com/raindevteam

# These are optional
# We don't recommend setting a default go route as you may have packages installed in different
# places such as github.com and gopkg.in. You also can't move go packages outside you GOPATH.
default-routes :
  js : modules/js
  py : modules/py

# Global options define options to be used for all modules that fall under their respective type
global-options :
  js : npm
  py : noload
  go : noload
`

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
		rlog.Fatal("Bot", "Could not correctly parse the configuration, check your syntax :::")
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

// GetNimbusConfig takes an rbot.Config and uses to create a nimbus.Config
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
