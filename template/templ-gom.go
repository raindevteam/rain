package tmpl

const GOMTemplate = `package main

import (
	"github.com/RyanPrintup/nimbus"
	"github.com/wolfchase/gorml"
)

type {{.Name}} struct{ *module.Module }

func (m *{{.Name}}) Greet(msg *nimbus.Message, args []string) {
	m.Say(msg.Args[0], "Hello there!")
}

func main() {
	m := &{{.Name}}{module.MakeModule("{{.Name}}", "Your module's short description")}

	m.CommandHook("greet", &module.Command{
		Help: "It greets you!",
		Fun:  m.Info,
	})

	m.Register()
}
`
