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
	ConnectListeners                 map[string][]Listener
	DisconnectListeners              map[string][]Listener
	RateLimitListeners               map[string][]Listener
	EventListeners                   map[string][]Listener
	ReadyListeners                   map[string][]Listener
	ChannelCreateListeners           map[string][]Listener
	ChannelUpdateListeners           map[string][]Listener
	ChannelDeleteListeners           map[string][]Listener
	ChannelPinsUpdateListeners       map[string][]Listener
	GuildCreateListeners             map[string][]Listener
	GuildUpdateListeners             map[string][]Listener
	GuildDeleteListeners             map[string][]Listener
	GuildBanAddListeners             map[string][]Listener
	GuildBanRemoveListeners          map[string][]Listener
	GuildMemberAddListeners          map[string][]Listener
	GuildMemberUpdateListeners       map[string][]Listener
	GuildMemberRemoveListeners       map[string][]Listener
	GuildRoleCreateListeners         map[string][]Listener
	GuildRoleUpdateListeners         map[string][]Listener
	GuildRoleDeleteListeners         map[string][]Listener
	GuildEmojisUpdateListeners       map[string][]Listener
	GuildMembersChunkListeners       map[string][]Listener
	GuildIntegrationsUpdateListeners map[string][]Listener
	MessageAckListeners              map[string][]Listener
	MessageCreateListeners           map[string][]Listener
	MessageUpdateListeners           map[string][]Listener
	MessageDeleteListeners           map[string][]Listener
	MessageReactionAddListeners      map[string][]Listener
	MessageReactionRemoveListeners   map[string][]Listener
	PresencesReplaceListeners        map[string][]Listener
	PresenceUpdateListeners          map[string][]Listener
	ResumedListeners                 map[string][]Listener
	RelationshipAddListeners         map[string][]Listener
	RelationshipRemoveListeners      map[string][]Listener
	TypingStartListeners             map[string][]Listener
	UserUpdateListeners              map[string][]Listener
	UserSettingsUpdateListeners      map[string][]Listener
	UserGuildSettingsUpdateListeners map[string][]Listener
	VoiceServerUpdateListeners       map[string][]Listener
	VoiceStateUpdateListeners        map[string][]Listener
}

type ConnectHandler func(*discordgo.Connect)

func (eh ConnectHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Connect); ok {
		eh(e)
	}
}

type DisconnectHandler func(*discordgo.Disconnect)

func (eh DisconnectHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Disconnect); ok {
		eh(e)
	}
}

type RateLimitHandler func(*discordgo.RateLimit)

func (eh RateLimitHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.RateLimit); ok {
		eh(e)
	}
}

type EventHandler func(*discordgo.Event)

func (eh EventHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Event); ok {
		eh(e)
	}
}

type ReadyHandler func(*discordgo.Ready)

func (eh ReadyHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Ready); ok {
		eh(e)
	}
}

type ChannelCreateHandler func(*discordgo.ChannelCreate)

func (eh ChannelCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelCreate); ok {
		eh(e)
	}
}

type ChannelUpdateHandler func(*discordgo.ChannelUpdate)

func (eh ChannelUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelUpdate); ok {
		eh(e)
	}
}

type ChannelDeleteHandler func(*discordgo.ChannelDelete)

func (eh ChannelDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelDelete); ok {
		eh(e)
	}
}

type ChannelPinsUpdateHandler func(*discordgo.ChannelPinsUpdate)

func (eh ChannelPinsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelPinsUpdate); ok {
		eh(e)
	}
}

type GuildCreateHandler func(*discordgo.GuildCreate)

func (eh GuildCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildCreate); ok {
		eh(e)
	}
}

type GuildUpdateHandler func(*discordgo.GuildUpdate)

func (eh GuildUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildUpdate); ok {
		eh(e)
	}
}

type GuildDeleteHandler func(*discordgo.GuildDelete)

func (eh GuildDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildDelete); ok {
		eh(e)
	}
}

type GuildBanAddHandler func(*discordgo.GuildBanAdd)

func (eh GuildBanAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildBanAdd); ok {
		eh(e)
	}
}

type GuildBanRemoveHandler func(*discordgo.GuildBanRemove)

func (eh GuildBanRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildBanRemove); ok {
		eh(e)
	}
}

type GuildMemberAddHandler func(*discordgo.GuildMemberAdd)

func (eh GuildMemberAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMemberAdd); ok {
		eh(e)
	}
}

type GuildMemberUpdateHandler func(*discordgo.GuildMemberUpdate)

func (eh GuildMemberUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMemberUpdate); ok {
		eh(e)
	}
}

type GuildMemberRemoveHandler func(*discordgo.GuildMemberRemove)

func (eh GuildMemberRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMemberRemove); ok {
		eh(e)
	}
}

type GuildRoleCreateHandler func(*discordgo.GuildRoleCreate)

func (eh GuildRoleCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildRoleCreate); ok {
		eh(e)
	}
}

type GuildRoleUpdateHandler func(*discordgo.GuildRoleUpdate)

func (eh GuildRoleUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildRoleUpdate); ok {
		eh(e)
	}
}

type GuildRoleDeleteHandler func(*discordgo.GuildRoleDelete)

func (eh GuildRoleDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildRoleDelete); ok {
		eh(e)
	}
}

type GuildEmojisUpdateHandler func(*discordgo.GuildEmojisUpdate)

func (eh GuildEmojisUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildEmojisUpdate); ok {
		eh(e)
	}
}

type GuildMembersChunkHandler func(*discordgo.GuildMembersChunk)

func (eh GuildMembersChunkHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMembersChunk); ok {
		eh(e)
	}
}

type GuildIntegrationsUpdateHandler func(*discordgo.GuildIntegrationsUpdate)

func (eh GuildIntegrationsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildIntegrationsUpdate); ok {
		eh(e)
	}
}

type MessageAckHandler func(*discordgo.MessageAck)

func (eh MessageAckHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageAck); ok {
		eh(e)
	}
}

type MessageCreateHandler func(*discordgo.MessageCreate)

func (eh MessageCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageCreate); ok {
		eh(e)
	}
}

type MessageUpdateHandler func(*discordgo.MessageUpdate)

func (eh MessageUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageUpdate); ok {
		eh(e)
	}
}

type MessageDeleteHandler func(*discordgo.MessageDelete)

func (eh MessageDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageDelete); ok {
		eh(e)
	}
}

type MessageReactionAddHandler func(*discordgo.MessageReactionAdd)

func (eh MessageReactionAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageReactionAdd); ok {
		eh(e)
	}
}

type MessageReactionRemoveHandler func(*discordgo.MessageReactionRemove)

func (eh MessageReactionRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageReactionRemove); ok {
		eh(e)
	}
}

type PresencesReplaceHandler func(*discordgo.PresencesReplace)

func (eh PresencesReplaceHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.PresencesReplace); ok {
		eh(e)
	}
}

type PresenceUpdateHandler func(*discordgo.PresenceUpdate)

func (eh PresenceUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.PresenceUpdate); ok {
		eh(e)
	}
}

type ResumedHandler func(*discordgo.Resumed)

func (eh ResumedHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Resumed); ok {
		eh(e)
	}
}

type RelationshipAddHandler func(*discordgo.RelationshipAdd)

func (eh RelationshipAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.RelationshipAdd); ok {
		eh(e)
	}
}

type RelationshipRemoveHandler func(*discordgo.RelationshipRemove)

func (eh RelationshipRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.RelationshipRemove); ok {
		eh(e)
	}
}

type TypingStartHandler func(*discordgo.TypingStart)

func (eh TypingStartHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.TypingStart); ok {
		eh(e)
	}
}

type UserUpdateHandler func(*discordgo.UserUpdate)

func (eh UserUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.UserUpdate); ok {
		eh(e)
	}
}

type UserSettingsUpdateHandler func(*discordgo.UserSettingsUpdate)

func (eh UserSettingsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.UserSettingsUpdate); ok {
		eh(e)
	}
}

type UserGuildSettingsUpdateHandler func(*discordgo.UserGuildSettingsUpdate)

func (eh UserGuildSettingsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.UserGuildSettingsUpdate); ok {
		eh(e)
	}
}

type VoiceServerUpdateHandler func(*discordgo.VoiceServerUpdate)

func (eh VoiceServerUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.VoiceServerUpdate); ok {
		eh(e)
	}
}

type VoiceStateUpdateHandler func(*discordgo.VoiceStateUpdate)

func (eh VoiceStateUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.VoiceStateUpdate); ok {
		eh(e)
	}
}

func (r Registry) Initilize() {
	r.ConnectListeners = make(map[string][]Listener)
	r.DisconnectListeners = make(map[string][]Listener)
	r.RateLimitListeners = make(map[string][]Listener)
	r.EventListeners = make(map[string][]Listener)
	r.ReadyListeners = make(map[string][]Listener)
	r.ChannelCreateListeners = make(map[string][]Listener)
	r.ChannelUpdateListeners = make(map[string][]Listener)
	r.ChannelDeleteListeners = make(map[string][]Listener)
	r.ChannelPinsUpdateListeners = make(map[string][]Listener)
	r.GuildCreateListeners = make(map[string][]Listener)
	r.GuildUpdateListeners = make(map[string][]Listener)
	r.GuildDeleteListeners = make(map[string][]Listener)
	r.GuildBanAddListeners = make(map[string][]Listener)
	r.GuildBanRemoveListeners = make(map[string][]Listener)
	r.GuildMemberAddListeners = make(map[string][]Listener)
	r.GuildMemberUpdateListeners = make(map[string][]Listener)
	r.GuildMemberRemoveListeners = make(map[string][]Listener)
	r.GuildRoleCreateListeners = make(map[string][]Listener)
	r.GuildRoleUpdateListeners = make(map[string][]Listener)
	r.GuildRoleDeleteListeners = make(map[string][]Listener)
	r.GuildEmojisUpdateListeners = make(map[string][]Listener)
	r.GuildMembersChunkListeners = make(map[string][]Listener)
	r.GuildIntegrationsUpdateListeners = make(map[string][]Listener)
	r.MessageAckListeners = make(map[string][]Listener)
	r.MessageCreateListeners = make(map[string][]Listener)
	r.MessageUpdateListeners = make(map[string][]Listener)
	r.MessageDeleteListeners = make(map[string][]Listener)
	r.MessageReactionAddListeners = make(map[string][]Listener)
	r.MessageReactionRemoveListeners = make(map[string][]Listener)
	r.PresencesReplaceListeners = make(map[string][]Listener)
	r.PresenceUpdateListeners = make(map[string][]Listener)
	r.ResumedListeners = make(map[string][]Listener)
	r.RelationshipAddListeners = make(map[string][]Listener)
	r.RelationshipRemoveListeners = make(map[string][]Listener)
	r.TypingStartListeners = make(map[string][]Listener)
	r.UserUpdateListeners = make(map[string][]Listener)
	r.UserSettingsUpdateListeners = make(map[string][]Listener)
	r.UserGuildSettingsUpdateListeners = make(map[string][]Listener)
	r.VoiceServerUpdateListeners = make(map[string][]Listener)
	r.VoiceStateUpdateListeners = make(map[string][]Listener)
}

func (h Handler) runListeners(ls map[string][]Listener, v interface{}) {
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

func (h Handler) runEventConnect(s *discordgo.Session, e *discordgo.Connect) {
	h.runListeners(h.registry.ConnectListeners, e)
}

func (h Handler) runEventDisconnect(s *discordgo.Session, e *discordgo.Disconnect) {
	h.runListeners(h.registry.DisconnectListeners, e)
}

func (h Handler) runEventRateLimit(s *discordgo.Session, e *discordgo.RateLimit) {
	h.runListeners(h.registry.RateLimitListeners, e)
}

func (h Handler) runEventEvent(s *discordgo.Session, e *discordgo.Event) {
	h.runListeners(h.registry.EventListeners, e)
}

func (h Handler) runEventReady(s *discordgo.Session, e *discordgo.Ready) {
	h.runListeners(h.registry.ReadyListeners, e)
}

func (h Handler) runEventChannelCreate(s *discordgo.Session, e *discordgo.ChannelCreate) {
	h.runListeners(h.registry.ChannelCreateListeners, e)
}

func (h Handler) runEventChannelUpdate(s *discordgo.Session, e *discordgo.ChannelUpdate) {
	h.runListeners(h.registry.ChannelUpdateListeners, e)
}

func (h Handler) runEventChannelDelete(s *discordgo.Session, e *discordgo.ChannelDelete) {
	h.runListeners(h.registry.ChannelDeleteListeners, e)
}

func (h Handler) runEventChannelPinsUpdate(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
	h.runListeners(h.registry.ChannelPinsUpdateListeners, e)
}

func (h Handler) runEventGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	h.runListeners(h.registry.GuildCreateListeners, e)
}

func (h Handler) runEventGuildUpdate(s *discordgo.Session, e *discordgo.GuildUpdate) {
	h.runListeners(h.registry.GuildUpdateListeners, e)
}

func (h Handler) runEventGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	h.runListeners(h.registry.GuildDeleteListeners, e)
}

func (h Handler) runEventGuildBanAdd(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	h.runListeners(h.registry.GuildBanAddListeners, e)
}

func (h Handler) runEventGuildBanRemove(s *discordgo.Session, e *discordgo.GuildBanRemove) {
	h.runListeners(h.registry.GuildBanRemoveListeners, e)
}

func (h Handler) runEventGuildMemberAdd(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	h.runListeners(h.registry.GuildMemberAddListeners, e)
}

func (h Handler) runEventGuildMemberUpdate(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
	h.runListeners(h.registry.GuildMemberUpdateListeners, e)
}

func (h Handler) runEventGuildMemberRemove(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	h.runListeners(h.registry.GuildMemberRemoveListeners, e)
}

func (h Handler) runEventGuildRoleCreate(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
	h.runListeners(h.registry.GuildRoleCreateListeners, e)
}

func (h Handler) runEventGuildRoleUpdate(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	h.runListeners(h.registry.GuildRoleUpdateListeners, e)
}

func (h Handler) runEventGuildRoleDelete(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
	h.runListeners(h.registry.GuildRoleDeleteListeners, e)
}

func (h Handler) runEventGuildEmojisUpdate(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
	h.runListeners(h.registry.GuildEmojisUpdateListeners, e)
}

func (h Handler) runEventGuildMembersChunk(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
	h.runListeners(h.registry.GuildMembersChunkListeners, e)
}

func (h Handler) runEventGuildIntegrationsUpdate(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
	h.runListeners(h.registry.GuildIntegrationsUpdateListeners, e)
}

func (h Handler) runEventMessageAck(s *discordgo.Session, e *discordgo.MessageAck) {
	h.runListeners(h.registry.MessageAckListeners, e)
}

func (h Handler) runEventMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	h.runListeners(h.registry.MessageCreateListeners, e)
}

func (h Handler) runEventMessageUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	h.runListeners(h.registry.MessageUpdateListeners, e)
}

func (h Handler) runEventMessageDelete(s *discordgo.Session, e *discordgo.MessageDelete) {
	h.runListeners(h.registry.MessageDeleteListeners, e)
}

func (h Handler) runEventMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	h.runListeners(h.registry.MessageReactionAddListeners, e)
}

func (h Handler) runEventMessageReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	h.runListeners(h.registry.MessageReactionRemoveListeners, e)
}

func (h Handler) runEventPresencesReplace(s *discordgo.Session, e *discordgo.PresencesReplace) {
	h.runListeners(h.registry.PresencesReplaceListeners, e)
}

func (h Handler) runEventPresenceUpdate(s *discordgo.Session, e *discordgo.PresenceUpdate) {
	h.runListeners(h.registry.PresenceUpdateListeners, e)
}

func (h Handler) runEventResumed(s *discordgo.Session, e *discordgo.Resumed) {
	h.runListeners(h.registry.ResumedListeners, e)
}

func (h Handler) runEventRelationshipAdd(s *discordgo.Session, e *discordgo.RelationshipAdd) {
	h.runListeners(h.registry.RelationshipAddListeners, e)
}

func (h Handler) runEventRelationshipRemove(s *discordgo.Session, e *discordgo.RelationshipRemove) {
	h.runListeners(h.registry.RelationshipRemoveListeners, e)
}

func (h Handler) runEventTypingStart(s *discordgo.Session, e *discordgo.TypingStart) {
	h.runListeners(h.registry.TypingStartListeners, e)
}

func (h Handler) runEventUserUpdate(s *discordgo.Session, e *discordgo.UserUpdate) {
	h.runListeners(h.registry.UserUpdateListeners, e)
}

func (h Handler) runEventUserSettingsUpdate(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
	h.runListeners(h.registry.UserSettingsUpdateListeners, e)
}

func (h Handler) runEventUserGuildSettingsUpdate(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
	h.runListeners(h.registry.UserGuildSettingsUpdateListeners, e)
}

func (h Handler) runEventVoiceServerUpdate(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
	h.runListeners(h.registry.VoiceServerUpdateListeners, e)
}

func (h Handler) runEventVoiceStateUpdate(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	h.runListeners(h.registry.VoiceStateUpdateListeners, e)
}

func (r Registry) CreateListener(v interface{}, isInternal bool) Listener {
	var l Listener
	if isInternal {
		l = &InternalListener{}
	} else {
		l = &DropletListener{}
	}

	switch f := v.(type) {
	case func(*discordgo.Connect):
		l.SetActionHandler(ConnectHandler(f))
		r.ConnectListeners[l.Owner()] = append(r.ConnectListeners[l.Owner()], l)
	case func(*discordgo.Disconnect):
		l.SetActionHandler(DisconnectHandler(f))
		r.DisconnectListeners[l.Owner()] = append(r.DisconnectListeners[l.Owner()], l)
	case func(*discordgo.RateLimit):
		l.SetActionHandler(RateLimitHandler(f))
		r.RateLimitListeners[l.Owner()] = append(r.RateLimitListeners[l.Owner()], l)
	case func(*discordgo.Event):
		l.SetActionHandler(EventHandler(f))
		r.EventListeners[l.Owner()] = append(r.EventListeners[l.Owner()], l)
	case func(*discordgo.Ready):
		l.SetActionHandler(ReadyHandler(f))
		r.ReadyListeners[l.Owner()] = append(r.ReadyListeners[l.Owner()], l)
	case func(*discordgo.ChannelCreate):
		l.SetActionHandler(ChannelCreateHandler(f))
		r.ChannelCreateListeners[l.Owner()] = append(r.ChannelCreateListeners[l.Owner()], l)
	case func(*discordgo.ChannelUpdate):
		l.SetActionHandler(ChannelUpdateHandler(f))
		r.ChannelUpdateListeners[l.Owner()] = append(r.ChannelUpdateListeners[l.Owner()], l)
	case func(*discordgo.ChannelDelete):
		l.SetActionHandler(ChannelDeleteHandler(f))
		r.ChannelDeleteListeners[l.Owner()] = append(r.ChannelDeleteListeners[l.Owner()], l)
	case func(*discordgo.ChannelPinsUpdate):
		l.SetActionHandler(ChannelPinsUpdateHandler(f))
		r.ChannelPinsUpdateListeners[l.Owner()] = append(r.ChannelPinsUpdateListeners[l.Owner()], l)
	case func(*discordgo.GuildCreate):
		l.SetActionHandler(GuildCreateHandler(f))
		r.GuildCreateListeners[l.Owner()] = append(r.GuildCreateListeners[l.Owner()], l)
	case func(*discordgo.GuildUpdate):
		l.SetActionHandler(GuildUpdateHandler(f))
		r.GuildUpdateListeners[l.Owner()] = append(r.GuildUpdateListeners[l.Owner()], l)
	case func(*discordgo.GuildDelete):
		l.SetActionHandler(GuildDeleteHandler(f))
		r.GuildDeleteListeners[l.Owner()] = append(r.GuildDeleteListeners[l.Owner()], l)
	case func(*discordgo.GuildBanAdd):
		l.SetActionHandler(GuildBanAddHandler(f))
		r.GuildBanAddListeners[l.Owner()] = append(r.GuildBanAddListeners[l.Owner()], l)
	case func(*discordgo.GuildBanRemove):
		l.SetActionHandler(GuildBanRemoveHandler(f))
		r.GuildBanRemoveListeners[l.Owner()] = append(r.GuildBanRemoveListeners[l.Owner()], l)
	case func(*discordgo.GuildMemberAdd):
		l.SetActionHandler(GuildMemberAddHandler(f))
		r.GuildMemberAddListeners[l.Owner()] = append(r.GuildMemberAddListeners[l.Owner()], l)
	case func(*discordgo.GuildMemberUpdate):
		l.SetActionHandler(GuildMemberUpdateHandler(f))
		r.GuildMemberUpdateListeners[l.Owner()] = append(r.GuildMemberUpdateListeners[l.Owner()], l)
	case func(*discordgo.GuildMemberRemove):
		l.SetActionHandler(GuildMemberRemoveHandler(f))
		r.GuildMemberRemoveListeners[l.Owner()] = append(r.GuildMemberRemoveListeners[l.Owner()], l)
	case func(*discordgo.GuildRoleCreate):
		l.SetActionHandler(GuildRoleCreateHandler(f))
		r.GuildRoleCreateListeners[l.Owner()] = append(r.GuildRoleCreateListeners[l.Owner()], l)
	case func(*discordgo.GuildRoleUpdate):
		l.SetActionHandler(GuildRoleUpdateHandler(f))
		r.GuildRoleUpdateListeners[l.Owner()] = append(r.GuildRoleUpdateListeners[l.Owner()], l)
	case func(*discordgo.GuildRoleDelete):
		l.SetActionHandler(GuildRoleDeleteHandler(f))
		r.GuildRoleDeleteListeners[l.Owner()] = append(r.GuildRoleDeleteListeners[l.Owner()], l)
	case func(*discordgo.GuildEmojisUpdate):
		l.SetActionHandler(GuildEmojisUpdateHandler(f))
		r.GuildEmojisUpdateListeners[l.Owner()] = append(r.GuildEmojisUpdateListeners[l.Owner()], l)
	case func(*discordgo.GuildMembersChunk):
		l.SetActionHandler(GuildMembersChunkHandler(f))
		r.GuildMembersChunkListeners[l.Owner()] = append(r.GuildMembersChunkListeners[l.Owner()], l)
	case func(*discordgo.GuildIntegrationsUpdate):
		l.SetActionHandler(GuildIntegrationsUpdateHandler(f))
		r.GuildIntegrationsUpdateListeners[l.Owner()] = append(r.GuildIntegrationsUpdateListeners[l.Owner()], l)
	case func(*discordgo.MessageAck):
		l.SetActionHandler(MessageAckHandler(f))
		r.MessageAckListeners[l.Owner()] = append(r.MessageAckListeners[l.Owner()], l)
	case func(*discordgo.MessageCreate):
		l.SetActionHandler(MessageCreateHandler(f))
		r.MessageCreateListeners[l.Owner()] = append(r.MessageCreateListeners[l.Owner()], l)
	case func(*discordgo.MessageUpdate):
		l.SetActionHandler(MessageUpdateHandler(f))
		r.MessageUpdateListeners[l.Owner()] = append(r.MessageUpdateListeners[l.Owner()], l)
	case func(*discordgo.MessageDelete):
		l.SetActionHandler(MessageDeleteHandler(f))
		r.MessageDeleteListeners[l.Owner()] = append(r.MessageDeleteListeners[l.Owner()], l)
	case func(*discordgo.MessageReactionAdd):
		l.SetActionHandler(MessageReactionAddHandler(f))
		r.MessageReactionAddListeners[l.Owner()] = append(r.MessageReactionAddListeners[l.Owner()], l)
	case func(*discordgo.MessageReactionRemove):
		l.SetActionHandler(MessageReactionRemoveHandler(f))
		r.MessageReactionRemoveListeners[l.Owner()] = append(r.MessageReactionRemoveListeners[l.Owner()], l)
	case func(*discordgo.PresencesReplace):
		l.SetActionHandler(PresencesReplaceHandler(f))
		r.PresencesReplaceListeners[l.Owner()] = append(r.PresencesReplaceListeners[l.Owner()], l)
	case func(*discordgo.PresenceUpdate):
		l.SetActionHandler(PresenceUpdateHandler(f))
		r.PresenceUpdateListeners[l.Owner()] = append(r.PresenceUpdateListeners[l.Owner()], l)
	case func(*discordgo.Resumed):
		l.SetActionHandler(ResumedHandler(f))
		r.ResumedListeners[l.Owner()] = append(r.ResumedListeners[l.Owner()], l)
	case func(*discordgo.RelationshipAdd):
		l.SetActionHandler(RelationshipAddHandler(f))
		r.RelationshipAddListeners[l.Owner()] = append(r.RelationshipAddListeners[l.Owner()], l)
	case func(*discordgo.RelationshipRemove):
		l.SetActionHandler(RelationshipRemoveHandler(f))
		r.RelationshipRemoveListeners[l.Owner()] = append(r.RelationshipRemoveListeners[l.Owner()], l)
	case func(*discordgo.TypingStart):
		l.SetActionHandler(TypingStartHandler(f))
		r.TypingStartListeners[l.Owner()] = append(r.TypingStartListeners[l.Owner()], l)
	case func(*discordgo.UserUpdate):
		l.SetActionHandler(UserUpdateHandler(f))
		r.UserUpdateListeners[l.Owner()] = append(r.UserUpdateListeners[l.Owner()], l)
	case func(*discordgo.UserSettingsUpdate):
		l.SetActionHandler(UserSettingsUpdateHandler(f))
		r.UserSettingsUpdateListeners[l.Owner()] = append(r.UserSettingsUpdateListeners[l.Owner()], l)
	case func(*discordgo.UserGuildSettingsUpdate):
		l.SetActionHandler(UserGuildSettingsUpdateHandler(f))
		r.UserGuildSettingsUpdateListeners[l.Owner()] = append(r.UserGuildSettingsUpdateListeners[l.Owner()], l)
	case func(*discordgo.VoiceServerUpdate):
		l.SetActionHandler(VoiceServerUpdateHandler(f))
		r.VoiceServerUpdateListeners[l.Owner()] = append(r.VoiceServerUpdateListeners[l.Owner()], l)
	case func(*discordgo.VoiceStateUpdate):
		l.SetActionHandler(VoiceStateUpdateHandler(f))
		r.VoiceStateUpdateListeners[l.Owner()] = append(r.VoiceStateUpdateListeners[l.Owner()], l)
	}

	return l
}

func (h Handler) AttachHandlers(b *bot.Bot) {
	b.Session.AddHandler(h.runEventConnect)
	b.Session.AddHandler(h.runEventDisconnect)
	b.Session.AddHandler(h.runEventRateLimit)
	b.Session.AddHandler(h.runEventEvent)
	b.Session.AddHandler(h.runEventReady)
	b.Session.AddHandler(h.runEventChannelCreate)
	b.Session.AddHandler(h.runEventChannelUpdate)
	b.Session.AddHandler(h.runEventChannelDelete)
	b.Session.AddHandler(h.runEventChannelPinsUpdate)
	b.Session.AddHandler(h.runEventGuildCreate)
	b.Session.AddHandler(h.runEventGuildUpdate)
	b.Session.AddHandler(h.runEventGuildDelete)
	b.Session.AddHandler(h.runEventGuildBanAdd)
	b.Session.AddHandler(h.runEventGuildBanRemove)
	b.Session.AddHandler(h.runEventGuildMemberAdd)
	b.Session.AddHandler(h.runEventGuildMemberUpdate)
	b.Session.AddHandler(h.runEventGuildMemberRemove)
	b.Session.AddHandler(h.runEventGuildRoleCreate)
	b.Session.AddHandler(h.runEventGuildRoleUpdate)
	b.Session.AddHandler(h.runEventGuildRoleDelete)
	b.Session.AddHandler(h.runEventGuildEmojisUpdate)
	b.Session.AddHandler(h.runEventGuildMembersChunk)
	b.Session.AddHandler(h.runEventGuildIntegrationsUpdate)
	b.Session.AddHandler(h.runEventMessageAck)
	b.Session.AddHandler(h.runEventMessageCreate)
	b.Session.AddHandler(h.runEventMessageUpdate)
	b.Session.AddHandler(h.runEventMessageDelete)
	b.Session.AddHandler(h.runEventMessageReactionAdd)
	b.Session.AddHandler(h.runEventMessageReactionRemove)
	b.Session.AddHandler(h.runEventPresencesReplace)
	b.Session.AddHandler(h.runEventPresenceUpdate)
	b.Session.AddHandler(h.runEventResumed)
	b.Session.AddHandler(h.runEventRelationshipAdd)
	b.Session.AddHandler(h.runEventRelationshipRemove)
	b.Session.AddHandler(h.runEventTypingStart)
	b.Session.AddHandler(h.runEventUserUpdate)
	b.Session.AddHandler(h.runEventUserSettingsUpdate)
	b.Session.AddHandler(h.runEventUserGuildSettingsUpdate)
	b.Session.AddHandler(h.runEventVoiceServerUpdate)
	b.Session.AddHandler(h.runEventVoiceStateUpdate)
}
