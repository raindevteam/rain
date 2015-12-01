package main

import (
    "github.com/Wolfchase/rainbot-go"
    "github.com/sorcix/irc"
)

type Core struct { *rainbot.Module }

func (m *Core) About(msg *irc.Message, args []string) {
    m.Say(msg.Params[0], "Maintained by Wolf and Ryan, Ver: 0.1.0")
}

func (m *Core) Info(msg *irc.Message, args []string) {
    m.Say(msg.Params[0], "See github.com/Wolfchase/rainbot-go/tree/develop")
}

func (m *Core) Ver(msg *irc.Message, args []string) {
    if len(args) > 0 && args[0] == "long" {
        m.Say(msg.Params[0], "DEV pre-0.1.0 Stability: Highly Unstable, Tests Passing: Unknown, Tests Failing: Unknown")
    } else {
        m.Say(msg.Params[0], "DEV pre-0.1.0")
    }
}

func main() {
    m := &Core{ rainbot.MakeModule("core", "core module") }

    m.CommandHook("about", &rainbot.Command{
        Help: "Tells you a little something about the bot",
        Fun: m.About,
    })

    m.CommandHook("info", &rainbot.Command{
        Help: "Gives you the repo to source",
        Fun: m.Info,
    })

    m.CommandHook("ver", &rainbot.Command{
        Help: "Shows version",
        Fun: m.Ver,
    })

    m.Register()
}
