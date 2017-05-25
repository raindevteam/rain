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

// Internal is the string constant identifier for bot listeners.
const Internal = "__INTERNAL__"

// The Registry holds all listeners registered with the bot. They are grouped by
// droplet, however each contains an entry of listeners belonging to the bot.
// This entry is identified with the "__INTERNAL__" string constant.
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
func (r Registry) Initialize() {
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

func (h Handler) dispatchConnect(s *discordgo.Session, e *discordgo.Connect) {
	h.runListeners(h.registry.ConnectListeners, e)
}

func (h Handler) dispatchDisconnect(s *discordgo.Session, e *discordgo.Disconnect) {
	h.runListeners(h.registry.DisconnectListeners, e)
}

func (h Handler) dispatchRateLimit(s *discordgo.Session, e *discordgo.RateLimit) {
	h.runListeners(h.registry.RateLimitListeners, e)
}

func (h Handler) dispatchEvent(s *discordgo.Session, e *discordgo.Event) {
	h.runListeners(h.registry.EventListeners, e)
}

func (h Handler) dispatchReady(s *discordgo.Session, e *discordgo.Ready) {
	h.runListeners(h.registry.ReadyListeners, e)
}

func (h Handler) dispatchChannelCreate(s *discordgo.Session, e *discordgo.ChannelCreate) {
	h.runListeners(h.registry.ChannelCreateListeners, e)
}

func (h Handler) dispatchChannelUpdate(s *discordgo.Session, e *discordgo.ChannelUpdate) {
	h.runListeners(h.registry.ChannelUpdateListeners, e)
}

func (h Handler) dispatchChannelDelete(s *discordgo.Session, e *discordgo.ChannelDelete) {
	h.runListeners(h.registry.ChannelDeleteListeners, e)
}

func (h Handler) dispatchChannelPinsUpdate(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
	h.runListeners(h.registry.ChannelPinsUpdateListeners, e)
}

func (h Handler) dispatchGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	h.runListeners(h.registry.GuildCreateListeners, e)
}

func (h Handler) dispatchGuildUpdate(s *discordgo.Session, e *discordgo.GuildUpdate) {
	h.runListeners(h.registry.GuildUpdateListeners, e)
}

func (h Handler) dispatchGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	h.runListeners(h.registry.GuildDeleteListeners, e)
}

func (h Handler) dispatchGuildBanAdd(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	h.runListeners(h.registry.GuildBanAddListeners, e)
}

func (h Handler) dispatchGuildBanRemove(s *discordgo.Session, e *discordgo.GuildBanRemove) {
	h.runListeners(h.registry.GuildBanRemoveListeners, e)
}

func (h Handler) dispatchGuildMemberAdd(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	h.runListeners(h.registry.GuildMemberAddListeners, e)
}

func (h Handler) dispatchGuildMemberUpdate(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
	h.runListeners(h.registry.GuildMemberUpdateListeners, e)
}

func (h Handler) dispatchGuildMemberRemove(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	h.runListeners(h.registry.GuildMemberRemoveListeners, e)
}

func (h Handler) dispatchGuildRoleCreate(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
	h.runListeners(h.registry.GuildRoleCreateListeners, e)
}

func (h Handler) dispatchGuildRoleUpdate(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	h.runListeners(h.registry.GuildRoleUpdateListeners, e)
}

func (h Handler) dispatchGuildRoleDelete(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
	h.runListeners(h.registry.GuildRoleDeleteListeners, e)
}

func (h Handler) dispatchGuildEmojisUpdate(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
	h.runListeners(h.registry.GuildEmojisUpdateListeners, e)
}

func (h Handler) dispatchGuildMembersChunk(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
	h.runListeners(h.registry.GuildMembersChunkListeners, e)
}

func (h Handler) dispatchGuildIntegrationsUpdate(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
	h.runListeners(h.registry.GuildIntegrationsUpdateListeners, e)
}

func (h Handler) dispatchMessageAck(s *discordgo.Session, e *discordgo.MessageAck) {
	h.runListeners(h.registry.MessageAckListeners, e)
}

func (h Handler) dispatchMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	h.runListeners(h.registry.MessageCreateListeners, e)
}

func (h Handler) dispatchMessageUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	h.runListeners(h.registry.MessageUpdateListeners, e)
}

func (h Handler) dispatchMessageDelete(s *discordgo.Session, e *discordgo.MessageDelete) {
	h.runListeners(h.registry.MessageDeleteListeners, e)
}

func (h Handler) dispatchMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	h.runListeners(h.registry.MessageReactionAddListeners, e)
}

func (h Handler) dispatchMessageReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	h.runListeners(h.registry.MessageReactionRemoveListeners, e)
}

func (h Handler) dispatchPresencesReplace(s *discordgo.Session, e *discordgo.PresencesReplace) {
	h.runListeners(h.registry.PresencesReplaceListeners, e)
}

func (h Handler) dispatchPresenceUpdate(s *discordgo.Session, e *discordgo.PresenceUpdate) {
	h.runListeners(h.registry.PresenceUpdateListeners, e)
}

func (h Handler) dispatchResumed(s *discordgo.Session, e *discordgo.Resumed) {
	h.runListeners(h.registry.ResumedListeners, e)
}

func (h Handler) dispatchRelationshipAdd(s *discordgo.Session, e *discordgo.RelationshipAdd) {
	h.runListeners(h.registry.RelationshipAddListeners, e)
}

func (h Handler) dispatchRelationshipRemove(s *discordgo.Session, e *discordgo.RelationshipRemove) {
	h.runListeners(h.registry.RelationshipRemoveListeners, e)
}

func (h Handler) dispatchTypingStart(s *discordgo.Session, e *discordgo.TypingStart) {
	h.runListeners(h.registry.TypingStartListeners, e)
}

func (h Handler) dispatchUserUpdate(s *discordgo.Session, e *discordgo.UserUpdate) {
	h.runListeners(h.registry.UserUpdateListeners, e)
}

func (h Handler) dispatchUserSettingsUpdate(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
	h.runListeners(h.registry.UserSettingsUpdateListeners, e)
}

func (h Handler) dispatchUserGuildSettingsUpdate(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
	h.runListeners(h.registry.UserGuildSettingsUpdateListeners, e)
}

func (h Handler) dispatchVoiceServerUpdate(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
	h.runListeners(h.registry.VoiceServerUpdateListeners, e)
}

func (h Handler) dispatchVoiceStateUpdate(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	h.runListeners(h.registry.VoiceStateUpdateListeners, e)
}

// CreateListener creates a new listener from a given function.
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

// AttachDispatchers will add all dispatch functions to the discord session for
// each supported discord event. Supported events are those from the discordgo
// library.
func (h Handler) AttachDispatchers(b *bot.Bot) {
	b.Session.AddHandler(h.dispatchConnect)
	b.Session.AddHandler(h.dispatchDisconnect)
	b.Session.AddHandler(h.dispatchRateLimit)
	b.Session.AddHandler(h.dispatchEvent)
	b.Session.AddHandler(h.dispatchReady)
	b.Session.AddHandler(h.dispatchChannelCreate)
	b.Session.AddHandler(h.dispatchChannelUpdate)
	b.Session.AddHandler(h.dispatchChannelDelete)
	b.Session.AddHandler(h.dispatchChannelPinsUpdate)
	b.Session.AddHandler(h.dispatchGuildCreate)
	b.Session.AddHandler(h.dispatchGuildUpdate)
	b.Session.AddHandler(h.dispatchGuildDelete)
	b.Session.AddHandler(h.dispatchGuildBanAdd)
	b.Session.AddHandler(h.dispatchGuildBanRemove)
	b.Session.AddHandler(h.dispatchGuildMemberAdd)
	b.Session.AddHandler(h.dispatchGuildMemberUpdate)
	b.Session.AddHandler(h.dispatchGuildMemberRemove)
	b.Session.AddHandler(h.dispatchGuildRoleCreate)
	b.Session.AddHandler(h.dispatchGuildRoleUpdate)
	b.Session.AddHandler(h.dispatchGuildRoleDelete)
	b.Session.AddHandler(h.dispatchGuildEmojisUpdate)
	b.Session.AddHandler(h.dispatchGuildMembersChunk)
	b.Session.AddHandler(h.dispatchGuildIntegrationsUpdate)
	b.Session.AddHandler(h.dispatchMessageAck)
	b.Session.AddHandler(h.dispatchMessageCreate)
	b.Session.AddHandler(h.dispatchMessageUpdate)
	b.Session.AddHandler(h.dispatchMessageDelete)
	b.Session.AddHandler(h.dispatchMessageReactionAdd)
	b.Session.AddHandler(h.dispatchMessageReactionRemove)
	b.Session.AddHandler(h.dispatchPresencesReplace)
	b.Session.AddHandler(h.dispatchPresenceUpdate)
	b.Session.AddHandler(h.dispatchResumed)
	b.Session.AddHandler(h.dispatchRelationshipAdd)
	b.Session.AddHandler(h.dispatchRelationshipRemove)
	b.Session.AddHandler(h.dispatchTypingStart)
	b.Session.AddHandler(h.dispatchUserUpdate)
	b.Session.AddHandler(h.dispatchUserSettingsUpdate)
	b.Session.AddHandler(h.dispatchUserGuildSettingsUpdate)
	b.Session.AddHandler(h.dispatchVoiceServerUpdate)
	b.Session.AddHandler(h.dispatchVoiceStateUpdate)
}
