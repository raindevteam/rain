package main

import (
	"github.com/raindevteam/gorml"
	"gopkg.in/sorcix/irc.v1"
)

type GoMod struct{ *module.Module }

func (m *GoMod) Greet(msg *irc.Message, args []string) {
	m.Say(msg.Params[0], "Hello there!")
}

func main() {
	m := &GoMod{module.NewModule("GoMod", "Your module's short description")}

	m.AddCommand("greet", &module.Command{
		Help: "It greets you!",
		Fun:  m.Greet,
	})

	m.Register()
}
