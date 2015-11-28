package rainbot

import (
    "testing"
)

func TestNada(t *testing.T) {
    t.Log("We fine")
}

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("testConfig.json")

	if err != nil || config.Host != "irc.freenode.org" {
		t.Fatalf("Couldn't read config file: ",  err)
	}
}
