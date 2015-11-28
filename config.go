package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host string
	Port string
	Channel string

	Nick string
	RealName string
	UserName string

	CmdPrefix string
}

func ReadConfig(string path) (&Config, err) {
	config := Config{}

	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	
	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
