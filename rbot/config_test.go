package bot

import "testing"

var configString = `
name: RainBot
`

func TestNewConfigFromFile(t *testing.T) {
	SetupTesting(t)

	conf, err := NewConfigFromFile("../exconf/exconf.yaml")
	if err != nil {
		t.Fatalf("Error while parsing config file: %s\n", err)
	}

	if conf.Name != "RainBot" {
		t.Errorf("Name set incorrectly, got %s instead\n", conf.Name)
	}
	if !conf.testing {
		t.Errorf("Testing set incorrectly, got %t instead\n", conf.testing)
	}
}

func TestNewConfigFromString(t *testing.T) {
	SetupTesting(t)

	conf, err := NewConfigFromString(configString)
	if err != nil {
		t.Fatalf("Error while parsing config file: %s\n", err)
	}

	if conf.Name != "RainBot" {
		t.Errorf("Name set incorrectly, got %s instead\n", conf.Name)
	}
	if !conf.testing {
		t.Errorf("Testing set incorrectly, got %t instead\n", conf.testing)
	}
}
