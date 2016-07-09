package main

import (
	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/gorml"
)

type GoMod struct{ *module.Module }

func (m *GoMod) Greet(msg *nimbus.Message, args []string) {
	m.Say(msg.Args[0], "Hello there!")
}

func main() {
	m := &GoMod{module.NewModule("GoMod", "Your module's short description")}

	m.AddCommand("greet", &module.Command{
		Help: "It greets you!",
		Fun:  m.Greet,
	})

	m.Register()
}
