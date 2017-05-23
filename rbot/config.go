// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

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
