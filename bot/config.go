package bot

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config contains parsed information from a supplied YAML file. This file is
// provided via command line arguments.
type Config struct {
	Name    string
	testing bool
}

// NewConfigFromFile returns an initialized config from a given file path.
func NewConfigFromFile(filepath string) (*Config, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return NewConfigFromString(string(file))
}

// NewConfigFromString creates a config from a string using a YAML parser.
func NewConfigFromString(confstring string) (*Config, error) {
	conf := &Config{}
	err := yaml.Unmarshal([]byte(confstring), &conf)
	if err != nil {
		return nil, err
	}
	if os.Getenv("RAIN_TESTING") != "" {
		conf.testing = true
	} else {
		conf.testing = false
	}
	return conf, nil
}
