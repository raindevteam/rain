package main

import (
	"fmt"
	"os"

	"github.com/RyanPrintup/nimbus"
	"github.com/wolfchase/rainbot"
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
		Version:     "Alpha 0.1.0 (Second Wind)",
		Client:      nimbus.NewClient(rainConfig.Host, rainConfig.Nick, *nimConfig),
		ModuleNames: rainConfig.GoModules,
		Parser:      rainbot.NewParser(rainConfig.CmdPrefix),
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

		bot.Client.AddListener(nimbus.PRIVMSG, func(msg *nimbus.Message) {
			if bot.Parser.IsCommand(msg.Trailing) {
				command, args := bot.Parser.ParseCommand(msg.Trailing)
				bot.Handler.Invoke(msg, rainbot.CommandName(command), args)
			}
		})

		bot.Client.AddListener(nimbus.PRIVMSG, func(msg *nimbus.Message) {
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
