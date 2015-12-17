package rainbot

import "testing"

//"github.com/Wolfchase/rainbot"

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("../tests/config.json")

	if err != nil || config.Host != "irc.canternet.org" {
		t.Fatalf("Couldn't read config file: ", err)
	}
}
