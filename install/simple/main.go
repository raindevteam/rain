package main

import (
	"github.com/wolfchase/rainbot/bot"
	"github.com/RyanPrintup/nimbus"
)

func main() {
	bot := rbot.NewBot("Alpha 0.4.0 (Monterey Jack)", &rbot.Config{
		Host:    "irc.canternet.org",
		Port:    "6667",
		Channel: []string{"#snowybottest"},
		Nick:     "SnowBot",
		RealName: "RainBotGo",
		UserName: "RainBotGo",
		Modes:    "+B",
	})

	bot.DefaultSetup()
	bot.DefaultConnect()

	// Echo
	bot.AddListener("PRIVMSG", func (msg *nimbus.Message) {
		if len(msg.Args) > 0 {
			bot.Say(msg.Args[0], msg.Trailing)
		}
	})
}