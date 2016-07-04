package main

import (
	"fmt"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/bot"
	"github.com/raindevteam/rain/setup"
)

func main() {
	bot := rbot.NewBot("Alpha 0.5.0", &rbot.Config{
		Host:     "irc.canternet.org",
		Port:     "6667",
		Channel:  []string{"#snowybottest"},
		Nick:     "SnowBot",
		RealName: "Rain",
		UserName: "Rain",
		Modes:    "+B",
	})

	setup.Default(bot)
	bot.DefaultConnect()

	// Logs privmsgs from connected channels
	bot.AddListener("PRIVMSG", func(msg *nimbus.Message) {
		fmt.Println(msg.Args[0], msg.Trailing)
	})
}
