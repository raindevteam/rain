package main

import (
    "github.com/RyanPrintup/nimbus"
    "github.com/sorcix/irc"
    //"strings"
    "fmt"
)

func main() {
    config := nimbus.Config{
        Port: "6667",
        Channels: []string{"#RainBot", "#snowybottest"},
        RealName: "RainBotGo",
        UserName: "RainBotGo",
        Password: "",
    }

    bot := nimbus.New("irc.canternet.org", "HailBot", config)

    bot.Connect(func(e error) {
        if e != nil {
            fmt.Println(e)
            return
        }

        bot.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
            text := msg.Trailing
            if text == "Hello, " + bot.Nick {
                bot.Say(msg.Params[0], "Hello there!")
            }
        })

        ch := make(chan error)
        go bot.Listen(ch)
        err := <- ch

        if err != nil {
            fmt.Println("boom")
            fmt.Println(err)
        }
    })
}
