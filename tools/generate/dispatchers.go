package main

import (
	"fmt"
	"html/template"
	"os"
)

const filename = "dispatchers.go"
const filenametest = "dispatchers_test.go"

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

package handler

import (
	"github.com/bwmarrin/discordgo"
)

const (
{{ range . }}
	{{ . }} = "{{ . }}" {{ end }}
)

// The Registry holds all listeners registered with the bot. They are grouped by
// droplet, however each contains an entry of listeners belonging to the bot.
// This entry is identified with the "__INTERNAL__" string constant. 
type Registry struct {
	Listeners         map[string]map[string][]Listener
	Commands          map[string]map[string]Command
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
func (r *Registry) Initialize() {
	r.Listeners = make(map[string]map[string][]Listener)
	r.Commands = make(map[string]map[string]Command) {{ range . }}
	r.Listeners[{{ . }}] = make(map[string][]Listener) {{ end }}
}

func (h *Handler) runListeners(ls map[string][]Listener, v interface{}) {
	for _, listeners := range ls {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.SetEvent(v)
				l.Run()
			} else {
				return
			}
		}
	}
}

{{ range . }}
func (h *Handler) dispatch{{ . }}(s *discordgo.Session, e *discordgo.{{ . }}) {
	h.runListeners(h.registry.Listeners[{{ . }}], e)
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

	switch f := v.(type) { {{ range . }}
	case func(*discordgo.{{ . }}):
	    l.SetActionHandler({{ . }}Handler(f))
		r.Listeners[{{ . }}][l.Owner()] = append(r.Listeners[{{ . }}][l.Owner()], l){{ end }}
	}

	return l
}

// AddCommand will add a command to the registry.
func (r *Registry) AddCommand(c Command) {
	if r.Commands[c.Owner()] == nil {
		r.Commands[c.Owner()] = make(map[string]Command)
	}
	r.Commands[c.Owner()][c.GetName()] = c
}

// Attach will add all dispatch functions to the discord session for
// each supported discord event. Supported events are those from the discordgo
// library.
func (h *Handler) Attach(s *discordgo.Session) { 
	{{ range . }}
	s.AddHandler(h.dispatch{{ . }}) {{ end }}
}
`
var handlersTestTemplate = `// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
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
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
)

var confirmDo = make(chan bool, 1)

{{ range . }}
func Test{{ . }}Handler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   {{ . }}Handler
		args args
	}{
		{
			name: "run do",
			eh: {{ . }}Handler(func(e *discordgo.{{ . }}) {
				confirmDo {{ noescape "<" }}- true
			}),
			args: args{&discordgo.{{ . }}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case {{ noescape "<" }}-confirmDo:
			case {{ noescape "<" }}-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}
{{ end }}

{{ range . }}
func Test_dispatch{{ . }}(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.{{ . }}
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch {{ . }} listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.{{ . }}{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch {{ . }} listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.{{ . }}{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch {{ . }} listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.{{ . }}{},
			},
			listenersToRegister: 2,
		},
	}
	// Manually set up handler.
	H = &Handler{}
	H.registry = &Registry{}
	H.Status = true
	H.registry.Initialize()
	for _, tt := range tests {
		switch tt.listenersToRegister {
		case 0:
			t.Run(tt.name, func(t *testing.T) {
				dispatch{{ . }}(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler({{ . }}Handler(func(e *discordgo.{{ . }}) {
				l1chan {{ noescape "<" }}- true
			}))
			H.registry.{{ . }}Listeners[l.Owner()] = append(H.registry.{{ . }}Listeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatch{{ . }}(tt.args.s, tt.args.e)
			})
			select {
			case {{ noescape "<" }}-l1chan:
			case {{ noescape "<" }}-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler({{ . }}Handler(func(e *discordgo.{{ . }}) {
				l1chan {{ noescape "<" }}- true
			}))
			H.registry.{{ . }}Listeners[l1.Owner()] = append(H.registry.{{ . }}Listeners[l1.Owner()], l1)
			l2.SetActionHandler({{ . }}Handler(func(e *discordgo.{{ . }}) {
				l1chan {{ noescape "<" }}- true
			}))
			H.registry.{{ . }}Listeners[l2.Owner()] = append(H.registry.{{ . }}Listeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatch{{ . }}(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i {{ noescape "<" }} 2; i++ {
				select {
				case {{ noescape "<" }}-l1chan:
				case {{ noescape "<" }}-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}
{{ end }}

`

func noescape(str string) template.HTML {
	return template.HTML(str)
}

var fns = template.FuncMap{
	"noescape": noescape,
}

func genCode(file string, tstr string) {
	t := template.New(file)
	t.Funcs(fns)
	t, err := t.Parse(tstr)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("../../internal/handler/" + file)
	if err != nil {
		os.Exit(1)
	}
	t.Execute(f, discordEvents)
	f.Close()
}

// Gencode generates files needed for the handler subpackage.
func main() {
	genCode(filename, handlersTemplate)
	genCode(filenametest, handlersTestTemplate)
}
