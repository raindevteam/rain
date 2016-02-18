package rainbot_testing

import (
	"testing"

	"github.com/wolfchase/rainbot/bot"
)

func TestReadConfig(t *testing.T) {
	config, err := rainbot.ReadConfig("../tests/config.json")

	if err != nil || config.Host != "irc.canternet.org" {
		t.Fatalf("Couldn't read config file: ", err)
	}
}
