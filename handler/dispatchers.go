// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
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
	
	ReadyListeners         map[string][]Listener
	
	MessageListeners         map[string][]Listener
	
}

func (r Registry) Initilize() {
	
	r.ReadyListeners = make(map[string][]Listener)
	
	r.MessageListeners = make(map[string][]Listener)
	
}


func (h Handler) runEventReady(s *discordgo.Session, e *discordgo.Ready) {
	for _, listeners := range h.registry.ReadyListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessage(s *discordgo.Session, e *discordgo.Message) {
	for _, listeners := range h.registry.MessageListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}


func (r Registry) addListener(l Listener) {
	switch l.Action().(type) {
	
	case *discordgo.Ready:
		r.ReadyListeners[l.Owner()] = append(r.ReadyListeners[l.Owner()], l)
    
	case *discordgo.Message:
		r.MessageListeners[l.Owner()] = append(r.MessageListeners[l.Owner()], l)
    
	}
}

func runAction(v interface{}, act Action) {
	switch e := v.(type) {
    
	case *discordgo.Ready:
		f := act.(func(e *discordgo.Ready))
		f(e)
	
	case *discordgo.Message:
		f := act.(func(e *discordgo.Message))
		f(e)
		
	}
}

func (h Handler) AttachHandlers(b *bot.Bot) { 
	b.Session.AddHandler(h.runEventReady) 
	b.Session.AddHandler(h.runEventMessage) 
}
