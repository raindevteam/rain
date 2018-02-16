// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
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
	Connect                 = "Connect"
	Disconnect              = "Disconnect"
	RateLimit               = "RateLimit"
	Event                   = "Event"
	Ready                   = "Ready"
	ChannelCreate           = "ChannelCreate"
	ChannelUpdate           = "ChannelUpdate"
	ChannelDelete           = "ChannelDelete"
	ChannelPinsUpdate       = "ChannelPinsUpdate"
	GuildCreate             = "GuildCreate"
	GuildUpdate             = "GuildUpdate"
	GuildDelete             = "GuildDelete"
	GuildBanAdd             = "GuildBanAdd"
	GuildBanRemove          = "GuildBanRemove"
	GuildMemberAdd          = "GuildMemberAdd"
	GuildMemberUpdate       = "GuildMemberUpdate"
	GuildMemberRemove       = "GuildMemberRemove"
	GuildRoleCreate         = "GuildRoleCreate"
	GuildRoleUpdate         = "GuildRoleUpdate"
	GuildRoleDelete         = "GuildRoleDelete"
	GuildEmojisUpdate       = "GuildEmojisUpdate"
	GuildMembersChunk       = "GuildMembersChunk"
	GuildIntegrationsUpdate = "GuildIntegrationsUpdate"
	MessageAck              = "MessageAck"
	MessageCreate           = "MessageCreate"
	MessageUpdate           = "MessageUpdate"
	MessageDelete           = "MessageDelete"
	MessageReactionAdd      = "MessageReactionAdd"
	MessageReactionRemove   = "MessageReactionRemove"
	PresencesReplace        = "PresencesReplace"
	PresenceUpdate          = "PresenceUpdate"
	Resumed                 = "Resumed"
	RelationshipAdd         = "RelationshipAdd"
	RelationshipRemove      = "RelationshipRemove"
	TypingStart             = "TypingStart"
	UserUpdate              = "UserUpdate"
	UserSettingsUpdate      = "UserSettingsUpdate"
	UserGuildSettingsUpdate = "UserGuildSettingsUpdate"
	VoiceServerUpdate       = "VoiceServerUpdate"
	VoiceStateUpdate        = "VoiceStateUpdate"
)

// The Registry holds all listeners registered with the bot. They are grouped by
// droplet, however each contains an entry of listeners belonging to the bot.
// This entry is identified with the "__INTERNAL__" string constant.
type Registry struct {
	Listeners map[string]map[string][]Listener
	Commands  map[string]map[string]Command
}

// ConnectHandler is the handler for Connect Listeners.
type ConnectHandler func(*discordgo.Connect)

// Do runs the underlying function for the handled Listener.
func (eh ConnectHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Connect); ok {
		eh(e)
	}
}

// DisconnectHandler is the handler for Disconnect Listeners.
type DisconnectHandler func(*discordgo.Disconnect)

// Do runs the underlying function for the handled Listener.
func (eh DisconnectHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Disconnect); ok {
		eh(e)
	}
}

// RateLimitHandler is the handler for RateLimit Listeners.
type RateLimitHandler func(*discordgo.RateLimit)

// Do runs the underlying function for the handled Listener.
func (eh RateLimitHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.RateLimit); ok {
		eh(e)
	}
}

// EventHandler is the handler for Event Listeners.
type EventHandler func(*discordgo.Event)

// Do runs the underlying function for the handled Listener.
func (eh EventHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Event); ok {
		eh(e)
	}
}

// ReadyHandler is the handler for Ready Listeners.
type ReadyHandler func(*discordgo.Ready)

// Do runs the underlying function for the handled Listener.
func (eh ReadyHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Ready); ok {
		eh(e)
	}
}

// ChannelCreateHandler is the handler for ChannelCreate Listeners.
type ChannelCreateHandler func(*discordgo.ChannelCreate)

// Do runs the underlying function for the handled Listener.
func (eh ChannelCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelCreate); ok {
		eh(e)
	}
}

// ChannelUpdateHandler is the handler for ChannelUpdate Listeners.
type ChannelUpdateHandler func(*discordgo.ChannelUpdate)

// Do runs the underlying function for the handled Listener.
func (eh ChannelUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelUpdate); ok {
		eh(e)
	}
}

// ChannelDeleteHandler is the handler for ChannelDelete Listeners.
type ChannelDeleteHandler func(*discordgo.ChannelDelete)

// Do runs the underlying function for the handled Listener.
func (eh ChannelDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelDelete); ok {
		eh(e)
	}
}

// ChannelPinsUpdateHandler is the handler for ChannelPinsUpdate Listeners.
type ChannelPinsUpdateHandler func(*discordgo.ChannelPinsUpdate)

// Do runs the underlying function for the handled Listener.
func (eh ChannelPinsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.ChannelPinsUpdate); ok {
		eh(e)
	}
}

// GuildCreateHandler is the handler for GuildCreate Listeners.
type GuildCreateHandler func(*discordgo.GuildCreate)

// Do runs the underlying function for the handled Listener.
func (eh GuildCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildCreate); ok {
		eh(e)
	}
}

// GuildUpdateHandler is the handler for GuildUpdate Listeners.
type GuildUpdateHandler func(*discordgo.GuildUpdate)

// Do runs the underlying function for the handled Listener.
func (eh GuildUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildUpdate); ok {
		eh(e)
	}
}

// GuildDeleteHandler is the handler for GuildDelete Listeners.
type GuildDeleteHandler func(*discordgo.GuildDelete)

// Do runs the underlying function for the handled Listener.
func (eh GuildDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildDelete); ok {
		eh(e)
	}
}

// GuildBanAddHandler is the handler for GuildBanAdd Listeners.
type GuildBanAddHandler func(*discordgo.GuildBanAdd)

// Do runs the underlying function for the handled Listener.
func (eh GuildBanAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildBanAdd); ok {
		eh(e)
	}
}

// GuildBanRemoveHandler is the handler for GuildBanRemove Listeners.
type GuildBanRemoveHandler func(*discordgo.GuildBanRemove)

// Do runs the underlying function for the handled Listener.
func (eh GuildBanRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildBanRemove); ok {
		eh(e)
	}
}

// GuildMemberAddHandler is the handler for GuildMemberAdd Listeners.
type GuildMemberAddHandler func(*discordgo.GuildMemberAdd)

// Do runs the underlying function for the handled Listener.
func (eh GuildMemberAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMemberAdd); ok {
		eh(e)
	}
}

// GuildMemberUpdateHandler is the handler for GuildMemberUpdate Listeners.
type GuildMemberUpdateHandler func(*discordgo.GuildMemberUpdate)

// Do runs the underlying function for the handled Listener.
func (eh GuildMemberUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMemberUpdate); ok {
		eh(e)
	}
}

// GuildMemberRemoveHandler is the handler for GuildMemberRemove Listeners.
type GuildMemberRemoveHandler func(*discordgo.GuildMemberRemove)

// Do runs the underlying function for the handled Listener.
func (eh GuildMemberRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMemberRemove); ok {
		eh(e)
	}
}

// GuildRoleCreateHandler is the handler for GuildRoleCreate Listeners.
type GuildRoleCreateHandler func(*discordgo.GuildRoleCreate)

// Do runs the underlying function for the handled Listener.
func (eh GuildRoleCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildRoleCreate); ok {
		eh(e)
	}
}

// GuildRoleUpdateHandler is the handler for GuildRoleUpdate Listeners.
type GuildRoleUpdateHandler func(*discordgo.GuildRoleUpdate)

// Do runs the underlying function for the handled Listener.
func (eh GuildRoleUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildRoleUpdate); ok {
		eh(e)
	}
}

// GuildRoleDeleteHandler is the handler for GuildRoleDelete Listeners.
type GuildRoleDeleteHandler func(*discordgo.GuildRoleDelete)

// Do runs the underlying function for the handled Listener.
func (eh GuildRoleDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildRoleDelete); ok {
		eh(e)
	}
}

// GuildEmojisUpdateHandler is the handler for GuildEmojisUpdate Listeners.
type GuildEmojisUpdateHandler func(*discordgo.GuildEmojisUpdate)

// Do runs the underlying function for the handled Listener.
func (eh GuildEmojisUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildEmojisUpdate); ok {
		eh(e)
	}
}

// GuildMembersChunkHandler is the handler for GuildMembersChunk Listeners.
type GuildMembersChunkHandler func(*discordgo.GuildMembersChunk)

// Do runs the underlying function for the handled Listener.
func (eh GuildMembersChunkHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildMembersChunk); ok {
		eh(e)
	}
}

// GuildIntegrationsUpdateHandler is the handler for GuildIntegrationsUpdate Listeners.
type GuildIntegrationsUpdateHandler func(*discordgo.GuildIntegrationsUpdate)

// Do runs the underlying function for the handled Listener.
func (eh GuildIntegrationsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.GuildIntegrationsUpdate); ok {
		eh(e)
	}
}

// MessageAckHandler is the handler for MessageAck Listeners.
type MessageAckHandler func(*discordgo.MessageAck)

// Do runs the underlying function for the handled Listener.
func (eh MessageAckHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageAck); ok {
		eh(e)
	}
}

// MessageCreateHandler is the handler for MessageCreate Listeners.
type MessageCreateHandler func(*discordgo.MessageCreate)

// Do runs the underlying function for the handled Listener.
func (eh MessageCreateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageCreate); ok {
		eh(e)
	}
}

// MessageUpdateHandler is the handler for MessageUpdate Listeners.
type MessageUpdateHandler func(*discordgo.MessageUpdate)

// Do runs the underlying function for the handled Listener.
func (eh MessageUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageUpdate); ok {
		eh(e)
	}
}

// MessageDeleteHandler is the handler for MessageDelete Listeners.
type MessageDeleteHandler func(*discordgo.MessageDelete)

// Do runs the underlying function for the handled Listener.
func (eh MessageDeleteHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageDelete); ok {
		eh(e)
	}
}

// MessageReactionAddHandler is the handler for MessageReactionAdd Listeners.
type MessageReactionAddHandler func(*discordgo.MessageReactionAdd)

// Do runs the underlying function for the handled Listener.
func (eh MessageReactionAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageReactionAdd); ok {
		eh(e)
	}
}

// MessageReactionRemoveHandler is the handler for MessageReactionRemove Listeners.
type MessageReactionRemoveHandler func(*discordgo.MessageReactionRemove)

// Do runs the underlying function for the handled Listener.
func (eh MessageReactionRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.MessageReactionRemove); ok {
		eh(e)
	}
}

// PresencesReplaceHandler is the handler for PresencesReplace Listeners.
type PresencesReplaceHandler func(*discordgo.PresencesReplace)

// Do runs the underlying function for the handled Listener.
func (eh PresencesReplaceHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.PresencesReplace); ok {
		eh(e)
	}
}

// PresenceUpdateHandler is the handler for PresenceUpdate Listeners.
type PresenceUpdateHandler func(*discordgo.PresenceUpdate)

// Do runs the underlying function for the handled Listener.
func (eh PresenceUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.PresenceUpdate); ok {
		eh(e)
	}
}

// ResumedHandler is the handler for Resumed Listeners.
type ResumedHandler func(*discordgo.Resumed)

// Do runs the underlying function for the handled Listener.
func (eh ResumedHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.Resumed); ok {
		eh(e)
	}
}

// RelationshipAddHandler is the handler for RelationshipAdd Listeners.
type RelationshipAddHandler func(*discordgo.RelationshipAdd)

// Do runs the underlying function for the handled Listener.
func (eh RelationshipAddHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.RelationshipAdd); ok {
		eh(e)
	}
}

// RelationshipRemoveHandler is the handler for RelationshipRemove Listeners.
type RelationshipRemoveHandler func(*discordgo.RelationshipRemove)

// Do runs the underlying function for the handled Listener.
func (eh RelationshipRemoveHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.RelationshipRemove); ok {
		eh(e)
	}
}

// TypingStartHandler is the handler for TypingStart Listeners.
type TypingStartHandler func(*discordgo.TypingStart)

// Do runs the underlying function for the handled Listener.
func (eh TypingStartHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.TypingStart); ok {
		eh(e)
	}
}

// UserUpdateHandler is the handler for UserUpdate Listeners.
type UserUpdateHandler func(*discordgo.UserUpdate)

// Do runs the underlying function for the handled Listener.
func (eh UserUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.UserUpdate); ok {
		eh(e)
	}
}

// UserSettingsUpdateHandler is the handler for UserSettingsUpdate Listeners.
type UserSettingsUpdateHandler func(*discordgo.UserSettingsUpdate)

// Do runs the underlying function for the handled Listener.
func (eh UserSettingsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.UserSettingsUpdate); ok {
		eh(e)
	}
}

// UserGuildSettingsUpdateHandler is the handler for UserGuildSettingsUpdate Listeners.
type UserGuildSettingsUpdateHandler func(*discordgo.UserGuildSettingsUpdate)

// Do runs the underlying function for the handled Listener.
func (eh UserGuildSettingsUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.UserGuildSettingsUpdate); ok {
		eh(e)
	}
}

// VoiceServerUpdateHandler is the handler for VoiceServerUpdate Listeners.
type VoiceServerUpdateHandler func(*discordgo.VoiceServerUpdate)

// Do runs the underlying function for the handled Listener.
func (eh VoiceServerUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.VoiceServerUpdate); ok {
		eh(e)
	}
}

// VoiceStateUpdateHandler is the handler for VoiceStateUpdate Listeners.
type VoiceStateUpdateHandler func(*discordgo.VoiceStateUpdate)

// Do runs the underlying function for the handled Listener.
func (eh VoiceStateUpdateHandler) Do(v interface{}) {
	if e, ok := v.(*discordgo.VoiceStateUpdate); ok {
		eh(e)
	}
}

// Initialize will initialize all maps in the registry.
func (r *Registry) Initialize() {
	r.Listeners = make(map[string]map[string][]Listener)
	r.Commands = make(map[string]map[string]Command)
	r.Listeners[Connect] = make(map[string][]Listener)
	r.Listeners[Disconnect] = make(map[string][]Listener)
	r.Listeners[RateLimit] = make(map[string][]Listener)
	r.Listeners[Event] = make(map[string][]Listener)
	r.Listeners[Ready] = make(map[string][]Listener)
	r.Listeners[ChannelCreate] = make(map[string][]Listener)
	r.Listeners[ChannelUpdate] = make(map[string][]Listener)
	r.Listeners[ChannelDelete] = make(map[string][]Listener)
	r.Listeners[ChannelPinsUpdate] = make(map[string][]Listener)
	r.Listeners[GuildCreate] = make(map[string][]Listener)
	r.Listeners[GuildUpdate] = make(map[string][]Listener)
	r.Listeners[GuildDelete] = make(map[string][]Listener)
	r.Listeners[GuildBanAdd] = make(map[string][]Listener)
	r.Listeners[GuildBanRemove] = make(map[string][]Listener)
	r.Listeners[GuildMemberAdd] = make(map[string][]Listener)
	r.Listeners[GuildMemberUpdate] = make(map[string][]Listener)
	r.Listeners[GuildMemberRemove] = make(map[string][]Listener)
	r.Listeners[GuildRoleCreate] = make(map[string][]Listener)
	r.Listeners[GuildRoleUpdate] = make(map[string][]Listener)
	r.Listeners[GuildRoleDelete] = make(map[string][]Listener)
	r.Listeners[GuildEmojisUpdate] = make(map[string][]Listener)
	r.Listeners[GuildMembersChunk] = make(map[string][]Listener)
	r.Listeners[GuildIntegrationsUpdate] = make(map[string][]Listener)
	r.Listeners[MessageAck] = make(map[string][]Listener)
	r.Listeners[MessageCreate] = make(map[string][]Listener)
	r.Listeners[MessageUpdate] = make(map[string][]Listener)
	r.Listeners[MessageDelete] = make(map[string][]Listener)
	r.Listeners[MessageReactionAdd] = make(map[string][]Listener)
	r.Listeners[MessageReactionRemove] = make(map[string][]Listener)
	r.Listeners[PresencesReplace] = make(map[string][]Listener)
	r.Listeners[PresenceUpdate] = make(map[string][]Listener)
	r.Listeners[Resumed] = make(map[string][]Listener)
	r.Listeners[RelationshipAdd] = make(map[string][]Listener)
	r.Listeners[RelationshipRemove] = make(map[string][]Listener)
	r.Listeners[TypingStart] = make(map[string][]Listener)
	r.Listeners[UserUpdate] = make(map[string][]Listener)
	r.Listeners[UserSettingsUpdate] = make(map[string][]Listener)
	r.Listeners[UserGuildSettingsUpdate] = make(map[string][]Listener)
	r.Listeners[VoiceServerUpdate] = make(map[string][]Listener)
	r.Listeners[VoiceStateUpdate] = make(map[string][]Listener)
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

func (h *Handler) dispatchConnect(s *discordgo.Session, e *discordgo.Connect) {
	h.runListeners(h.registry.Listeners[Connect], e)
}

func (h *Handler) dispatchDisconnect(s *discordgo.Session, e *discordgo.Disconnect) {
	h.runListeners(h.registry.Listeners[Disconnect], e)
}

func (h *Handler) dispatchRateLimit(s *discordgo.Session, e *discordgo.RateLimit) {
	h.runListeners(h.registry.Listeners[RateLimit], e)
}

func (h *Handler) dispatchEvent(s *discordgo.Session, e *discordgo.Event) {
	h.runListeners(h.registry.Listeners[Event], e)
}

func (h *Handler) dispatchReady(s *discordgo.Session, e *discordgo.Ready) {
	h.runListeners(h.registry.Listeners[Ready], e)
}

func (h *Handler) dispatchChannelCreate(s *discordgo.Session, e *discordgo.ChannelCreate) {
	h.runListeners(h.registry.Listeners[ChannelCreate], e)
}

func (h *Handler) dispatchChannelUpdate(s *discordgo.Session, e *discordgo.ChannelUpdate) {
	h.runListeners(h.registry.Listeners[ChannelUpdate], e)
}

func (h *Handler) dispatchChannelDelete(s *discordgo.Session, e *discordgo.ChannelDelete) {
	h.runListeners(h.registry.Listeners[ChannelDelete], e)
}

func (h *Handler) dispatchChannelPinsUpdate(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
	h.runListeners(h.registry.Listeners[ChannelPinsUpdate], e)
}

func (h *Handler) dispatchGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	h.runListeners(h.registry.Listeners[GuildCreate], e)
}

func (h *Handler) dispatchGuildUpdate(s *discordgo.Session, e *discordgo.GuildUpdate) {
	h.runListeners(h.registry.Listeners[GuildUpdate], e)
}

func (h *Handler) dispatchGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	h.runListeners(h.registry.Listeners[GuildDelete], e)
}

func (h *Handler) dispatchGuildBanAdd(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	h.runListeners(h.registry.Listeners[GuildBanAdd], e)
}

func (h *Handler) dispatchGuildBanRemove(s *discordgo.Session, e *discordgo.GuildBanRemove) {
	h.runListeners(h.registry.Listeners[GuildBanRemove], e)
}

func (h *Handler) dispatchGuildMemberAdd(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	h.runListeners(h.registry.Listeners[GuildMemberAdd], e)
}

func (h *Handler) dispatchGuildMemberUpdate(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
	h.runListeners(h.registry.Listeners[GuildMemberUpdate], e)
}

func (h *Handler) dispatchGuildMemberRemove(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	h.runListeners(h.registry.Listeners[GuildMemberRemove], e)
}

func (h *Handler) dispatchGuildRoleCreate(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
	h.runListeners(h.registry.Listeners[GuildRoleCreate], e)
}

func (h *Handler) dispatchGuildRoleUpdate(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	h.runListeners(h.registry.Listeners[GuildRoleUpdate], e)
}

func (h *Handler) dispatchGuildRoleDelete(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
	h.runListeners(h.registry.Listeners[GuildRoleDelete], e)
}

func (h *Handler) dispatchGuildEmojisUpdate(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
	h.runListeners(h.registry.Listeners[GuildEmojisUpdate], e)
}

func (h *Handler) dispatchGuildMembersChunk(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
	h.runListeners(h.registry.Listeners[GuildMembersChunk], e)
}

func (h *Handler) dispatchGuildIntegrationsUpdate(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
	h.runListeners(h.registry.Listeners[GuildIntegrationsUpdate], e)
}

func (h *Handler) dispatchMessageAck(s *discordgo.Session, e *discordgo.MessageAck) {
	h.runListeners(h.registry.Listeners[MessageAck], e)
}

func (h *Handler) dispatchMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	h.runListeners(h.registry.Listeners[MessageCreate], e)
}

func (h *Handler) dispatchMessageUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	h.runListeners(h.registry.Listeners[MessageUpdate], e)
}

func (h *Handler) dispatchMessageDelete(s *discordgo.Session, e *discordgo.MessageDelete) {
	h.runListeners(h.registry.Listeners[MessageDelete], e)
}

func (h *Handler) dispatchMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	h.runListeners(h.registry.Listeners[MessageReactionAdd], e)
}

func (h *Handler) dispatchMessageReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	h.runListeners(h.registry.Listeners[MessageReactionRemove], e)
}

func (h *Handler) dispatchPresencesReplace(s *discordgo.Session, e *discordgo.PresencesReplace) {
	h.runListeners(h.registry.Listeners[PresencesReplace], e)
}

func (h *Handler) dispatchPresenceUpdate(s *discordgo.Session, e *discordgo.PresenceUpdate) {
	h.runListeners(h.registry.Listeners[PresenceUpdate], e)
}

func (h *Handler) dispatchResumed(s *discordgo.Session, e *discordgo.Resumed) {
	h.runListeners(h.registry.Listeners[Resumed], e)
}

func (h *Handler) dispatchRelationshipAdd(s *discordgo.Session, e *discordgo.RelationshipAdd) {
	h.runListeners(h.registry.Listeners[RelationshipAdd], e)
}

func (h *Handler) dispatchRelationshipRemove(s *discordgo.Session, e *discordgo.RelationshipRemove) {
	h.runListeners(h.registry.Listeners[RelationshipRemove], e)
}

func (h *Handler) dispatchTypingStart(s *discordgo.Session, e *discordgo.TypingStart) {
	h.runListeners(h.registry.Listeners[TypingStart], e)
}

func (h *Handler) dispatchUserUpdate(s *discordgo.Session, e *discordgo.UserUpdate) {
	h.runListeners(h.registry.Listeners[UserUpdate], e)
}

func (h *Handler) dispatchUserSettingsUpdate(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
	h.runListeners(h.registry.Listeners[UserSettingsUpdate], e)
}

func (h *Handler) dispatchUserGuildSettingsUpdate(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
	h.runListeners(h.registry.Listeners[UserGuildSettingsUpdate], e)
}

func (h *Handler) dispatchVoiceServerUpdate(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
	h.runListeners(h.registry.Listeners[VoiceServerUpdate], e)
}

func (h *Handler) dispatchVoiceStateUpdate(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	h.runListeners(h.registry.Listeners[VoiceStateUpdate], e)
}

// CreateListener creates a new listener from a given function.
func (r *Registry) CreateListener(v interface{}, isInternal bool) Listener {
	var l Listener
	if isInternal {
		l = &InternalListener{}
	} else {
		l = &DropletListener{}
	}

	switch f := v.(type) {
	case func(*discordgo.Connect):
		l.SetActionHandler(ConnectHandler(f))
		r.Listeners[Connect][l.Owner()] = append(r.Listeners[Connect][l.Owner()], l)
	case func(*discordgo.Disconnect):
		l.SetActionHandler(DisconnectHandler(f))
		r.Listeners[Disconnect][l.Owner()] = append(r.Listeners[Disconnect][l.Owner()], l)
	case func(*discordgo.RateLimit):
		l.SetActionHandler(RateLimitHandler(f))
		r.Listeners[RateLimit][l.Owner()] = append(r.Listeners[RateLimit][l.Owner()], l)
	case func(*discordgo.Event):
		l.SetActionHandler(EventHandler(f))
		r.Listeners[Event][l.Owner()] = append(r.Listeners[Event][l.Owner()], l)
	case func(*discordgo.Ready):
		l.SetActionHandler(ReadyHandler(f))
		r.Listeners[Ready][l.Owner()] = append(r.Listeners[Ready][l.Owner()], l)
	case func(*discordgo.ChannelCreate):
		l.SetActionHandler(ChannelCreateHandler(f))
		r.Listeners[ChannelCreate][l.Owner()] = append(r.Listeners[ChannelCreate][l.Owner()], l)
	case func(*discordgo.ChannelUpdate):
		l.SetActionHandler(ChannelUpdateHandler(f))
		r.Listeners[ChannelUpdate][l.Owner()] = append(r.Listeners[ChannelUpdate][l.Owner()], l)
	case func(*discordgo.ChannelDelete):
		l.SetActionHandler(ChannelDeleteHandler(f))
		r.Listeners[ChannelDelete][l.Owner()] = append(r.Listeners[ChannelDelete][l.Owner()], l)
	case func(*discordgo.ChannelPinsUpdate):
		l.SetActionHandler(ChannelPinsUpdateHandler(f))
		r.Listeners[ChannelPinsUpdate][l.Owner()] = append(r.Listeners[ChannelPinsUpdate][l.Owner()], l)
	case func(*discordgo.GuildCreate):
		l.SetActionHandler(GuildCreateHandler(f))
		r.Listeners[GuildCreate][l.Owner()] = append(r.Listeners[GuildCreate][l.Owner()], l)
	case func(*discordgo.GuildUpdate):
		l.SetActionHandler(GuildUpdateHandler(f))
		r.Listeners[GuildUpdate][l.Owner()] = append(r.Listeners[GuildUpdate][l.Owner()], l)
	case func(*discordgo.GuildDelete):
		l.SetActionHandler(GuildDeleteHandler(f))
		r.Listeners[GuildDelete][l.Owner()] = append(r.Listeners[GuildDelete][l.Owner()], l)
	case func(*discordgo.GuildBanAdd):
		l.SetActionHandler(GuildBanAddHandler(f))
		r.Listeners[GuildBanAdd][l.Owner()] = append(r.Listeners[GuildBanAdd][l.Owner()], l)
	case func(*discordgo.GuildBanRemove):
		l.SetActionHandler(GuildBanRemoveHandler(f))
		r.Listeners[GuildBanRemove][l.Owner()] = append(r.Listeners[GuildBanRemove][l.Owner()], l)
	case func(*discordgo.GuildMemberAdd):
		l.SetActionHandler(GuildMemberAddHandler(f))
		r.Listeners[GuildMemberAdd][l.Owner()] = append(r.Listeners[GuildMemberAdd][l.Owner()], l)
	case func(*discordgo.GuildMemberUpdate):
		l.SetActionHandler(GuildMemberUpdateHandler(f))
		r.Listeners[GuildMemberUpdate][l.Owner()] = append(r.Listeners[GuildMemberUpdate][l.Owner()], l)
	case func(*discordgo.GuildMemberRemove):
		l.SetActionHandler(GuildMemberRemoveHandler(f))
		r.Listeners[GuildMemberRemove][l.Owner()] = append(r.Listeners[GuildMemberRemove][l.Owner()], l)
	case func(*discordgo.GuildRoleCreate):
		l.SetActionHandler(GuildRoleCreateHandler(f))
		r.Listeners[GuildRoleCreate][l.Owner()] = append(r.Listeners[GuildRoleCreate][l.Owner()], l)
	case func(*discordgo.GuildRoleUpdate):
		l.SetActionHandler(GuildRoleUpdateHandler(f))
		r.Listeners[GuildRoleUpdate][l.Owner()] = append(r.Listeners[GuildRoleUpdate][l.Owner()], l)
	case func(*discordgo.GuildRoleDelete):
		l.SetActionHandler(GuildRoleDeleteHandler(f))
		r.Listeners[GuildRoleDelete][l.Owner()] = append(r.Listeners[GuildRoleDelete][l.Owner()], l)
	case func(*discordgo.GuildEmojisUpdate):
		l.SetActionHandler(GuildEmojisUpdateHandler(f))
		r.Listeners[GuildEmojisUpdate][l.Owner()] = append(r.Listeners[GuildEmojisUpdate][l.Owner()], l)
	case func(*discordgo.GuildMembersChunk):
		l.SetActionHandler(GuildMembersChunkHandler(f))
		r.Listeners[GuildMembersChunk][l.Owner()] = append(r.Listeners[GuildMembersChunk][l.Owner()], l)
	case func(*discordgo.GuildIntegrationsUpdate):
		l.SetActionHandler(GuildIntegrationsUpdateHandler(f))
		r.Listeners[GuildIntegrationsUpdate][l.Owner()] = append(r.Listeners[GuildIntegrationsUpdate][l.Owner()], l)
	case func(*discordgo.MessageAck):
		l.SetActionHandler(MessageAckHandler(f))
		r.Listeners[MessageAck][l.Owner()] = append(r.Listeners[MessageAck][l.Owner()], l)
	case func(*discordgo.MessageCreate):
		l.SetActionHandler(MessageCreateHandler(f))
		r.Listeners[MessageCreate][l.Owner()] = append(r.Listeners[MessageCreate][l.Owner()], l)
	case func(*discordgo.MessageUpdate):
		l.SetActionHandler(MessageUpdateHandler(f))
		r.Listeners[MessageUpdate][l.Owner()] = append(r.Listeners[MessageUpdate][l.Owner()], l)
	case func(*discordgo.MessageDelete):
		l.SetActionHandler(MessageDeleteHandler(f))
		r.Listeners[MessageDelete][l.Owner()] = append(r.Listeners[MessageDelete][l.Owner()], l)
	case func(*discordgo.MessageReactionAdd):
		l.SetActionHandler(MessageReactionAddHandler(f))
		r.Listeners[MessageReactionAdd][l.Owner()] = append(r.Listeners[MessageReactionAdd][l.Owner()], l)
	case func(*discordgo.MessageReactionRemove):
		l.SetActionHandler(MessageReactionRemoveHandler(f))
		r.Listeners[MessageReactionRemove][l.Owner()] = append(r.Listeners[MessageReactionRemove][l.Owner()], l)
	case func(*discordgo.PresencesReplace):
		l.SetActionHandler(PresencesReplaceHandler(f))
		r.Listeners[PresencesReplace][l.Owner()] = append(r.Listeners[PresencesReplace][l.Owner()], l)
	case func(*discordgo.PresenceUpdate):
		l.SetActionHandler(PresenceUpdateHandler(f))
		r.Listeners[PresenceUpdate][l.Owner()] = append(r.Listeners[PresenceUpdate][l.Owner()], l)
	case func(*discordgo.Resumed):
		l.SetActionHandler(ResumedHandler(f))
		r.Listeners[Resumed][l.Owner()] = append(r.Listeners[Resumed][l.Owner()], l)
	case func(*discordgo.RelationshipAdd):
		l.SetActionHandler(RelationshipAddHandler(f))
		r.Listeners[RelationshipAdd][l.Owner()] = append(r.Listeners[RelationshipAdd][l.Owner()], l)
	case func(*discordgo.RelationshipRemove):
		l.SetActionHandler(RelationshipRemoveHandler(f))
		r.Listeners[RelationshipRemove][l.Owner()] = append(r.Listeners[RelationshipRemove][l.Owner()], l)
	case func(*discordgo.TypingStart):
		l.SetActionHandler(TypingStartHandler(f))
		r.Listeners[TypingStart][l.Owner()] = append(r.Listeners[TypingStart][l.Owner()], l)
	case func(*discordgo.UserUpdate):
		l.SetActionHandler(UserUpdateHandler(f))
		r.Listeners[UserUpdate][l.Owner()] = append(r.Listeners[UserUpdate][l.Owner()], l)
	case func(*discordgo.UserSettingsUpdate):
		l.SetActionHandler(UserSettingsUpdateHandler(f))
		r.Listeners[UserSettingsUpdate][l.Owner()] = append(r.Listeners[UserSettingsUpdate][l.Owner()], l)
	case func(*discordgo.UserGuildSettingsUpdate):
		l.SetActionHandler(UserGuildSettingsUpdateHandler(f))
		r.Listeners[UserGuildSettingsUpdate][l.Owner()] = append(r.Listeners[UserGuildSettingsUpdate][l.Owner()], l)
	case func(*discordgo.VoiceServerUpdate):
		l.SetActionHandler(VoiceServerUpdateHandler(f))
		r.Listeners[VoiceServerUpdate][l.Owner()] = append(r.Listeners[VoiceServerUpdate][l.Owner()], l)
	case func(*discordgo.VoiceStateUpdate):
		l.SetActionHandler(VoiceStateUpdateHandler(f))
		r.Listeners[VoiceStateUpdate][l.Owner()] = append(r.Listeners[VoiceStateUpdate][l.Owner()], l)
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

	s.AddHandler(h.dispatchConnect)
	s.AddHandler(h.dispatchDisconnect)
	s.AddHandler(h.dispatchRateLimit)
	s.AddHandler(h.dispatchEvent)
	s.AddHandler(h.dispatchReady)
	s.AddHandler(h.dispatchChannelCreate)
	s.AddHandler(h.dispatchChannelUpdate)
	s.AddHandler(h.dispatchChannelDelete)
	s.AddHandler(h.dispatchChannelPinsUpdate)
	s.AddHandler(h.dispatchGuildCreate)
	s.AddHandler(h.dispatchGuildUpdate)
	s.AddHandler(h.dispatchGuildDelete)
	s.AddHandler(h.dispatchGuildBanAdd)
	s.AddHandler(h.dispatchGuildBanRemove)
	s.AddHandler(h.dispatchGuildMemberAdd)
	s.AddHandler(h.dispatchGuildMemberUpdate)
	s.AddHandler(h.dispatchGuildMemberRemove)
	s.AddHandler(h.dispatchGuildRoleCreate)
	s.AddHandler(h.dispatchGuildRoleUpdate)
	s.AddHandler(h.dispatchGuildRoleDelete)
	s.AddHandler(h.dispatchGuildEmojisUpdate)
	s.AddHandler(h.dispatchGuildMembersChunk)
	s.AddHandler(h.dispatchGuildIntegrationsUpdate)
	s.AddHandler(h.dispatchMessageAck)
	s.AddHandler(h.dispatchMessageCreate)
	s.AddHandler(h.dispatchMessageUpdate)
	s.AddHandler(h.dispatchMessageDelete)
	s.AddHandler(h.dispatchMessageReactionAdd)
	s.AddHandler(h.dispatchMessageReactionRemove)
	s.AddHandler(h.dispatchPresencesReplace)
	s.AddHandler(h.dispatchPresenceUpdate)
	s.AddHandler(h.dispatchResumed)
	s.AddHandler(h.dispatchRelationshipAdd)
	s.AddHandler(h.dispatchRelationshipRemove)
	s.AddHandler(h.dispatchTypingStart)
	s.AddHandler(h.dispatchUserUpdate)
	s.AddHandler(h.dispatchUserSettingsUpdate)
	s.AddHandler(h.dispatchUserGuildSettingsUpdate)
	s.AddHandler(h.dispatchVoiceServerUpdate)
	s.AddHandler(h.dispatchVoiceStateUpdate)
}
