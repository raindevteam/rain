package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RyanPrintup/nimbus"
	"github.com/Wolfchase/rainbot"
	"github.com/sorcix/irc"
)

func main() {
	var rainConfig *rainbot.Config
	var nimConfig *nimbus.Config

	var err error

	if len(os.Args) > 1 {
		rainConfig, err = rainbot.ReadConfig(os.Args[1])

		if err != nil {
			fmt.Println(err)
		}

		nimConfig = &nimbus.Config{
			Port:     rainConfig.Port,
			Channels: rainConfig.Channel,
			RealName: rainConfig.RealName,
			UserName: rainConfig.UserName,
			Password: "",
		}
	} else {
		fmt.Println("Usage: rainbot <config_file>")
		os.Exit(1)
	}

	bot := &rainbot.Bot{
		Client:      nimbus.NewClient(rainConfig.Host, rainConfig.Nick, *nimConfig),
		ModuleNames: []string{"rcore"},
		Handler:     rainbot.NewHandler(),
	}

	fmt.Print("Connecting... ")

	bot.Client.Connect(func(e error) {
		if e != nil {
			fmt.Println(e)
			return
		}

		fmt.Println("Done")

		bot.LoadModules()

		bot.Client.AddListener(irc.PRIVMSG, func(msg *nimbus.Message) {

		})

		bot.Client.AddListener(irc.PRIVMSG, func(msg *nimbus.Message) {
		})

		bot.Client.AddListener(irc.PRIVMSG, func(msg *nimbus.Message) {
			if string(msg.Trailing[0]) == ";" {
				splitMessage := strings.Split(string(msg.Trailing[1:]), " ")
				command, args := splitMessage[0], splitMessage[1:]
				bot.Handler.Invoke(msg, rainbot.CommandName(command), args)
			}
		})

		bot.Client.AddListener(irc.PRIVMSG, func(msg *nimbus.Message) {
			text := msg.Trailing
			if text == "Hello, "+bot.Client.Nick {
				bot.Client.Say(msg.Args[0], "Hello there!")
			}
		})

		ch := make(chan error)
		go bot.Client.Listen(ch)
		err := <-ch

		if err != nil {
			fmt.Println("boom")
			fmt.Println(err)
		}
	})
}
