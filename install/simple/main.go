package main

import (
	"fmt"

	"gopkg.in/sorcix/irc.v1"

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

	bot := rbot.NewBot("0.6.0-alpha.2", conf)

	setup.Default(bot)
	bot.DefaultConnect()

	// Logs privmsgs from connected channels
	bot.AddListener("PRIVMSG", func(msg *irc.Message) {
		fmt.Println(msg.Trailing)
	})

	result := bot.Quit()
	if result != nil {
		fmt.Println(result.Error())
	}
}
