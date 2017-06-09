package main

import (
	"fmt"
	"html/template"
	"os"
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
	"PresencesReplace",
	"PresenceUpdate",
	"Resumed",
	"RelationshipAdd",
	"RelationshipRemove",
	"TypingStart",
	"UserUpdate",
	"UserSettingsUpdate",
	"UserGuildSettingsUpdate",
	"VoiceServerUpdate",
	"VoiceStateUpdate",
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
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/hail"
	"github.com/raindevteam/rain/rbot"
)

// Internal is the string constant identifier for bot listeners.
const Internal = "__INTERNAL__"

// The Registry holds all listeners registered with the bot. They are grouped by
// droplet, however each contains an entry of listeners belonging to the bot.
// This entry is identified with the "__INTERNAL__" string constant. 
type Registry struct { {{ range $value := .}}
	{{ $value }}Listeners         map[string][]Listener {{ end }}
}

{{ range . }}
// {{ . }}Handler is the handler for {{ . }} Listeners.
type {{ . }}Handler func(*discordgo.{{ . }})

// Do runs the underlying function for the handled Listener.
func (eh {{ . }}Handler) Do(v interface{}) {
	if e, ok := v.(*discordgo.{{ . }}); ok {
		eh(e)
	}
}
{{ end }}

// Initialize will initialize all maps in the registry.
func (r *Registry) Initialize() { {{ range $value := . }}
	r.{{ $value }}Listeners = make(map[string][]Listener) {{ end }}
}

func runListeners(ls map[string][]Listener, v interface{}) {
	for _, listeners := range ls {
		for _, l := range listeners {
			if ok := H.Status; ok {
				l.SetEvent(v)
				l.Run()
			} else {
				return
			}
		}
	}
}

{{ range $value := . }}
func dispatch{{ $value }}(s *discordgo.Session, e *discordgo.{{ $value }}) {
	runListeners(H.registry.{{ $value }}Listeners, e)
}
{{ end }}

// CreateListener creates a new listener from a given function.
func (r *Registry) CreateListener(v interface{}, isInternal bool) Listener {
	var l Listener
	if isInternal {
		l = &InternalListener{}
	} else {
		l = &DropletListener{}
	}

	switch f := v.(type) { {{ range $value := . }}
	case func(*discordgo.{{ $value }}):
	    l.SetActionHandler({{ $value }}Handler(f))
		r.{{ $value }}Listeners[l.Owner()] = append(r.{{ $value }}Listeners[l.Owner()], l){{ end }}
	}

	return l
}

// Attach will add all dispatch functions to the discord session for
// each supported discord event. Supported events are those from the discordgo
// library.
func Attach(b *rbot.Bot) { 
	if H == nil {
		hail.Crit(hail.Feventhand,
			"H has not been created yet, cannot attach listeners! Exiting...")
		os.Exit(1)
	}{{ range $value := . }}
	b.Session.AddHandler(dispatch{{ $value }}) {{ end }}
}
`

// Gencode generates files needed for the handler subpackage.
func main() {
	t := template.New(filename)
	t, err := t.Parse(handlersTemplate)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("../handler/" + filename)
	if err != nil {
		os.Exit(1)
	}
	t.Execute(f, discordEvents)
	f.Close()
}
