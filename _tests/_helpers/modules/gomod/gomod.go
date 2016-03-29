package main

import (
	"github.com/RyanPrintup/nimbus"
	"github.com/wolfchase/gorml"
)

type GoMod struct{ *module.Module }

func (m *GoMod) Greet(msg *nimbus.Message, args []string) {
	m.Say(msg.Args[0], "Hello there!")
}

func main() {
	m := &GoMod{module.MakeModule("GoMod", "Your module's short description")}

	m.CommandHook("greet", &module.Command{
		Help: "It greets you!",
		Fun:  m.Greet,
	})

	m.Register()
}
