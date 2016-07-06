package main

import (
	"fmt"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rbot"
	"github.com/raindevteam/rain/setup"
)

var config = `
host     : irc.canternet.org
port     : 6667
channels : ["#snowybottest"]
nick     : SnowBot
realname : Rain
username : Rain
modes    : +B
`

func main() {
	bot := rbot.NewBot("0.5.0", rbot.ReadConfig(config))

	setup.Default(bot)
	bot.DefaultConnect()

	// Logs privmsgs from connected channels
	bot.AddListener("PRIVMSG", func(msg *nimbus.Message) {
		fmt.Println(msg.Trailing)
	})
}
