package main

import (
	"github.com/RyanPrintup/nimbus"
	"github.com/wolfchase/rainbot"
)

type Core struct{ *rainbot.Module }

func (m *Core) About(msg *nimbus.Message, args []string) {
	m.Say(msg.Args[0], "Maintained by Wolf and Ryan, Ver: 0.1.0")
}

func (m *Core) Info(msg *nimbus.Message, args []string) {
	m.Say(msg.Args[0], "See github.com/Wolfchase/rainbot-go/tree/develop")
}

func (m *Core) Ver(msg *nimbus.Message, args []string) {
	ver := m.GetBotVersion()
	m.Say(msg.Args[0], ver)
}

func main() {
	m := &Core{rainbot.MakeModule("rCore", "The official core module for RainBot")}

	m.CommandHook("about", &rainbot.Command{
		Help: "Tells you a little something about the bot",
		Fun:  m.About,
	})

	m.CommandHook("info", &rainbot.Command{
		Help: "Gives you the repo to source",
		Fun:  m.Info,
	})

	m.CommandHook("ver", &rainbot.Command{
		Help: "Shows version",
		Fun:  m.Ver,
	})

	m.Register()
}
