package rainbot

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

	CmdPrefix string

	GoModules []string
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

func GetConfigs(file string) (ncon *nimbus.Config, rcon *Config, err error) {
	rcon, err = ReadConfig(file)

	if err != nil {
		return nil, nil, err
	}

	ncon = &nimbus.Config{
		Port:     rcon.Port,
		Channels: rcon.Channel,
		RealName: rcon.RealName,
		UserName: rcon.UserName,
		Password: "",
	}

	return ncon, rcon, nil
}
