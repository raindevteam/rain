package main

import (
	"fmt"
	"os"

	"gopkg.in/sorcix/irc.v1"

	"github.com/raindevteam/rain/rain"
	"github.com/raindevteam/rain/rbot"
	"github.com/raindevteam/rain/setup"
)

var config = `
host     : irc.canternet.org
port     : 6667
channels : ["#snowybottest"]
nick     : SnowyBot
realname : Rain
username : Rain
modes    : +B
`

func main() {
	conf, err := rbot.ReadConfig(config)
	if err != nil {
		panic(err)
	}

	bot := rbot.NewBot(rain.Version(), conf)

	setup.Default(bot)
	bot.DefaultConnect()

	// Logs privmsgs from connected channels
	bot.AddListener("PRIVMSG", func(msg *irc.Message) {
		fmt.Println(msg.Trailing)
	})

	reason, err := bot.WaitForQuit()

	if err != nil {
		fmt.Printf("Quit Error: %s\n", reason)
		os.Exit(1)
	}

	fmt.Printf("Quit: %s\n", reason)
}
