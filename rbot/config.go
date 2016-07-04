package rbot

import (
	"encoding/json"
	"os"

	"github.com/RyanPrintup/nimbus"
)

type Config struct {
	Host    string
	Port    string
	Channel []string

	Nick     string
	RealName string
	UserName string
	Modes    string

	CmdPrefix string

	GoModules map[string]string
	JsModules map[string]string
}

func ReadConfig(path string) (*Config, error) {
	config := &Config{}

	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetNimbusConfig(rconf *Config) (nconf *nimbus.Config) {
	nconf = &nimbus.Config{
		Port:     rconf.Port,
		Channels: rconf.Channel,
		RealName: rconf.RealName,
		UserName: rconf.UserName,
		Password: "",
		Modes:    rconf.Modes,
	}

	return nconf
}
