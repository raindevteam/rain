package main

import (
    "strings"
    "fmt"
    "strconv"

    "github.com/RyanPrintup/nimbus"
    "github.com/sorcix/irc"
    "github.com/Wolfchase/rainbot"
)

func main() {
    config := nimbus.Config{
        Port: "6667",
        Channels: []string{"#snowybottest"},
        RealName: "RainBotGo",
        UserName: "RainBotGo",
        Password: "",
    }

    bot := &rainbot.Bot{
        Client: nimbus.New("irc.canternet.org", "HailBot", config),
        ModuleNames: []string{"rainbot-core"},
        Handler: rainbot.NewHandler(),
    }

    bot.Client.Connect(func(e error) {
        if e != nil {
            fmt.Println(e)
            return
        }

        bot.LoadModules()

        bot.Client.AddListener(irc.PRIVMSG, func(msg *irc.Message) {

        })

        bot.Client.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
        })

        bot.Client.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
            if string(msg.Trailing[0]) == ";" {
                splitMessage := strings.Split(string(msg.Trailing[1:]), " ")
                command, args := splitMessage[0], splitMessage[1:]
                bot.Handler.Invoke(msg, rainbot.CommandName(command), args)
            }
        })

        bot.Client.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
            text := msg.Trailing
            if text == "Hello, " + bot.Client.Nick {
                bot.Client.Say(msg.Params[0], "Hello there!")
            }
        })

        bot.Client.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
            text := msg.Trailing
            if text == "listeners" {
                bot.Client.Say(msg.Params[0], strconv.Itoa(len(bot.Client.Listeners[irc.PRIVMSG])))
            }
        })

        ch := make(chan error)
        go bot.Client.Listen(ch)
        err := <- ch

        if err != nil {
            fmt.Println("boom")
            fmt.Println(err)
        }
    })
}
