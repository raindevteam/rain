package main

import (
    "github.com/RyanPrintup/nimbus"
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

    err := bot.Connect()
    if err == nil {
        for {}
    }
}
