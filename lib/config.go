package rainbot

import (
	"encoding/json"
	"os"
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
