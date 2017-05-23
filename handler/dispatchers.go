// DO NOT EDIT; This code is automatically generated.
// See godoc package information for more details.

package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/bot"
)

const Internal = "__INTERNAL__"

type Registry struct {
	ReadyListeners map[string][]*Listener

	MessageListeners map[string][]*Listener
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

func (r Registry) addListener(l EventListener) {
	var listener interface{}
	switch l.(type) {
	case Listener:
		listener = l.(Listener)
	case InternalListener:
		listener = l.(InternalListener)
	}

	switch listener.act.(type) {

	case *discordgo.Ready:
		r.ReadyInternalListeners = append(r.ReadyInternalListeners, il)

	case *discordgo.Message:
		r.MessageInternalListeners = append(r.MessageInternalListeners, il)

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
	b.Ds.AddHandler(h.RunEventReady)
	b.Ds.AddHandler(h.RunEventMessage)
}
