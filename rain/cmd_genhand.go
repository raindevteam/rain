// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package rain

import (
	"fmt"
	"html/template"
	"os"

	"gopkg.in/urfave/cli.v1"
)

const filename = "dispatchers.go"

var discordEvents = []string{
	"Ready",
	"Message",
}

var handlersTemplate = `// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
//
// Contact Rodolfo at rcvallada@gmail.com for any inquires of this file.

// DO NOT EDIT; This code is automatically generated.
// See godoc package information for more details.

package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/bot"
)

const Internal = "__INTERNAL__"

type Registry struct {
	{{ range $value := .}}
	{{ $value }}Listeners         map[string][]*Listener
	{{ end }}
}

{{ range $value := . }}
func (h Handler) runEvent{{ $value }}(s *discordgo.Session, e *discordgo.{{ $value }}) {
	for _, listeners := range h.registry.{{ $value }}Listeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}
{{ end }}

func (r Registry) addListener(l EventListener) {
	var listener EventListener
	switch l.(type) {
	case Listener:
		listener = l.(Listener)
    case InternalListener:
		listener = l.(InternalListener)
	}

	switch listener.act.(type) {
	{{ range $value := . }}
	case *discordgo.{{ $value }}:
		r.{{ $value }}InternalListeners = append(r.{{ $value }}InternalListeners, il)
    {{ end }}
	}
}

func runAction(v interface{}, act Action) {
	switch e := v.(type) {
    {{ range $value := . }}
	case *discordgo.{{ $value }}:
		f := act.(func(e *discordgo.{{ $value }}))
		f(e)
	{{ end }}	
	}
}

func (h Handler) AttachHandlers(b *bot.Bot) { {{ range $value := . }}
	b.Ds.AddHandler(h.RunEvent{{ $value }}) {{ end }}
}
`

// Dir will set the GOLNS_DIR environment variable to the new argument-passed subdirectory.
func Gdispatch(cli *cli.Context) error {
	if cli.NArg() != 0 {
		fmt.Println(cli.Command.ArgsUsage)
	}

	t := template.New(filename)
	t, err := t.Parse(handlersTemplate)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("./handler/" + filename)
	if err != nil {
		return err
	}
	t.Execute(f, discordEvents)
	f.Close()
	return nil
}

// NewCommandDir appends the Dir function wrapped in a cli.Command struct.
func NewCmdGenhand(app *cli.App) {
	Command := cli.Command{
		Name:      "gdispatch",
		Aliases:   []string{"gh"},
		Usage:     "Generate the handlers file",
		ArgsUsage: "",
		Action:    Gdispatch,
	}
	app.Commands = append(app.Commands, Command)
}
