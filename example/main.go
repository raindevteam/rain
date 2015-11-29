package main

import (
    "github.com/RyanPrintup/nimbus"
    "github.com/sorcix/irc"
    //"github.com/Wolfchase/rainbot-go/module"
    "github.com/Wolfchase/rainbot-go/rainbot"
    //"strings"
    "fmt"
    //"strconv"
)

func main() {
    config := nimbus.Config{
        Port: "6667",
        Channels: []string{"#RainBot", "#snowybottest"},
        RealName: "RainBotGo",
        UserName: "RainBotGo",
        Password: "",
    }

    bot := &rainbot.Bot{}
    bot.C = nimbus.New("irc.canternet.org", "HailBot", config)
    bot.Modules = []string{
        "rainbot-core",
    }

    bot.C.Connect(func(e error) {
        if e != nil {
            fmt.Println(e)
            return
        }

        bot.LoadModules()

        bot.C.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
            text := msg.Trailing
            if text == "Hello, " + bot.C.Nick {
                bot.C.Say(msg.Params[0], "Hello there!")
            }
        })

        bot.C.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
            text := msg.Trailing
            if text == "listeners" {
                bot.LoadModules()
            }
        })

        ch := make(chan error)
        go bot.C.Listen(ch)
        err := <- ch

        if err != nil {
            fmt.Println("boom")
            fmt.Println(err)
        }
    })
}
