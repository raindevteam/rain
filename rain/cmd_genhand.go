// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
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
	"Connect",
	"Disconnect",
	"RateLimit",
	"Event",
	"Ready",
	"ChannelCreate",
	"ChannelUpdate",
	"ChannelDelete",
	"ChannelPinsUpdate",
	"GuildCreate",
	"GuildUpdate",
	"GuildDelete",
	"GuildBanAdd",
	"GuildBanRemove",
	"GuildMemberAdd",
	"GuildMemberUpdate",
	"GuildMemberRemove",
	"GuildRoleCreate",
	"GuildRoleUpdate",
	"GuildRoleDelete",
	"GuildEmojisUpdate",
	"GuildMembersChunk",
	"GuildIntegrationsUpdate",
	"MessageAck",
	"MessageCreate",
	"MessageUpdate",
	"MessageDelete",
	"MessageReactionAdd",
	"MessageReactionRemove",
	"MessageReactionRemoveAll",
	"PresencesReplace",
	"PresenceUpdate",
	"Resumed",
	"RelationshipAdd",
	"RelationshipRemove",
	"TypingStart",
	"UserUpdate",
	"UserSettingsUpdate",
	"UserGuildSettingsUpdate",
	"UserNoteUpdate",
	"VoiceServerUpdate",
	"VoiceStateUpdate",
	"MessageDeleteBulk",
}

var handlersTemplate = `// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

// DO NOT EDIT; This code is automatically generated.
// See godoc package information for more details.

package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/rbot"
)

const Internal = "__INTERNAL__"

type Registry struct {
	{{ range $value := .}}
	{{ $value }}Listeners         map[string][]Listener
	{{ end }}
}

func (r Registry) Initilize() {
	{{ range $value := . }}
	r.{{ $value }}Listeners = make(map[string][]Listener)
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

func (r Registry) addListener(l Listener) {
	switch l.Action().(type) {
	{{ range $value := . }}
	case *discordgo.{{ $value }}:
		r.{{ $value }}Listeners[l.Owner()] = append(r.{{ $value }}Listeners[l.Owner()], l)
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
	b.Session.AddHandler(h.runEvent{{ $value }}) {{ end }}
}
`

// Gencode generates files needed for the handler subpackage.
func Gencode(cli *cli.Context) error {
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

// NewCmdGencode appends the Gencode function wrapped in a
// cli.Command struct.
func NewCmdGencode(app *cli.App) {
	Command := cli.Command{
		Name:      "gencode",
		Aliases:   []string{"gc"},
		Usage:     "Generate handler files (dispathers.go, registry.go)",
		ArgsUsage: "",
		Action:    Gencode,
	}
	app.Commands = append(app.Commands, Command)
}
