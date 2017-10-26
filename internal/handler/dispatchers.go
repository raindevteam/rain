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
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/internal/hail"
	"github.com/raindevteam/rain/internal/rbot"
)

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
func (r *Registry) Initialize() {
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

func dispatchConnect(s *discordgo.Session, e *discordgo.Connect) {
	runListeners(H.registry.ConnectListeners, e)
}

func dispatchDisconnect(s *discordgo.Session, e *discordgo.Disconnect) {
	runListeners(H.registry.DisconnectListeners, e)
}

func dispatchRateLimit(s *discordgo.Session, e *discordgo.RateLimit) {
	runListeners(H.registry.RateLimitListeners, e)
}

func dispatchEvent(s *discordgo.Session, e *discordgo.Event) {
	runListeners(H.registry.EventListeners, e)
}

func dispatchReady(s *discordgo.Session, e *discordgo.Ready) {
	runListeners(H.registry.ReadyListeners, e)
}

func dispatchChannelCreate(s *discordgo.Session, e *discordgo.ChannelCreate) {
	runListeners(H.registry.ChannelCreateListeners, e)
}

func dispatchChannelUpdate(s *discordgo.Session, e *discordgo.ChannelUpdate) {
	runListeners(H.registry.ChannelUpdateListeners, e)
}

func dispatchChannelDelete(s *discordgo.Session, e *discordgo.ChannelDelete) {
	runListeners(H.registry.ChannelDeleteListeners, e)
}

func dispatchChannelPinsUpdate(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
	runListeners(H.registry.ChannelPinsUpdateListeners, e)
}

func dispatchGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	runListeners(H.registry.GuildCreateListeners, e)
}

func dispatchGuildUpdate(s *discordgo.Session, e *discordgo.GuildUpdate) {
	runListeners(H.registry.GuildUpdateListeners, e)
}

func dispatchGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	runListeners(H.registry.GuildDeleteListeners, e)
}

func dispatchGuildBanAdd(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	runListeners(H.registry.GuildBanAddListeners, e)
}

func dispatchGuildBanRemove(s *discordgo.Session, e *discordgo.GuildBanRemove) {
	runListeners(H.registry.GuildBanRemoveListeners, e)
}

func dispatchGuildMemberAdd(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	runListeners(H.registry.GuildMemberAddListeners, e)
}

func dispatchGuildMemberUpdate(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
	runListeners(H.registry.GuildMemberUpdateListeners, e)
}

func dispatchGuildMemberRemove(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	runListeners(H.registry.GuildMemberRemoveListeners, e)
}

func dispatchGuildRoleCreate(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
	runListeners(H.registry.GuildRoleCreateListeners, e)
}

func dispatchGuildRoleUpdate(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	runListeners(H.registry.GuildRoleUpdateListeners, e)
}

func dispatchGuildRoleDelete(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
	runListeners(H.registry.GuildRoleDeleteListeners, e)
}

func dispatchGuildEmojisUpdate(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
	runListeners(H.registry.GuildEmojisUpdateListeners, e)
}

func dispatchGuildMembersChunk(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
	runListeners(H.registry.GuildMembersChunkListeners, e)
}

func dispatchGuildIntegrationsUpdate(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
	runListeners(H.registry.GuildIntegrationsUpdateListeners, e)
}

func dispatchMessageAck(s *discordgo.Session, e *discordgo.MessageAck) {
	runListeners(H.registry.MessageAckListeners, e)
}

func dispatchMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	runListeners(H.registry.MessageCreateListeners, e)
}

func dispatchMessageUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	runListeners(H.registry.MessageUpdateListeners, e)
}

func dispatchMessageDelete(s *discordgo.Session, e *discordgo.MessageDelete) {
	runListeners(H.registry.MessageDeleteListeners, e)
}

func dispatchMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	runListeners(H.registry.MessageReactionAddListeners, e)
}

func dispatchMessageReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	runListeners(H.registry.MessageReactionRemoveListeners, e)
}

func dispatchPresencesReplace(s *discordgo.Session, e *discordgo.PresencesReplace) {
	runListeners(H.registry.PresencesReplaceListeners, e)
}

func dispatchPresenceUpdate(s *discordgo.Session, e *discordgo.PresenceUpdate) {
	runListeners(H.registry.PresenceUpdateListeners, e)
}

func dispatchResumed(s *discordgo.Session, e *discordgo.Resumed) {
	runListeners(H.registry.ResumedListeners, e)
}

func dispatchRelationshipAdd(s *discordgo.Session, e *discordgo.RelationshipAdd) {
	runListeners(H.registry.RelationshipAddListeners, e)
}

func dispatchRelationshipRemove(s *discordgo.Session, e *discordgo.RelationshipRemove) {
	runListeners(H.registry.RelationshipRemoveListeners, e)
}

func dispatchTypingStart(s *discordgo.Session, e *discordgo.TypingStart) {
	runListeners(H.registry.TypingStartListeners, e)
}

func dispatchUserUpdate(s *discordgo.Session, e *discordgo.UserUpdate) {
	runListeners(H.registry.UserUpdateListeners, e)
}

func dispatchUserSettingsUpdate(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
	runListeners(H.registry.UserSettingsUpdateListeners, e)
}

func dispatchUserGuildSettingsUpdate(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
	runListeners(H.registry.UserGuildSettingsUpdateListeners, e)
}

func dispatchVoiceServerUpdate(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
	runListeners(H.registry.VoiceServerUpdateListeners, e)
}

func dispatchVoiceStateUpdate(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	runListeners(H.registry.VoiceStateUpdateListeners, e)
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

// Attach will add all dispatch functions to the discord session for
// each supported discord event. Supported events are those from the discordgo
// library.
func Attach(b *rbot.Bot) {
	if H == nil {
		hail.Crit(hail.Fhandler,
			"H has not been created yet, cannot attach listeners! Exiting...")
		os.Exit(1)
	}
	b.Session.AddHandler(dispatchConnect)
	b.Session.AddHandler(dispatchDisconnect)
	b.Session.AddHandler(dispatchRateLimit)
	b.Session.AddHandler(dispatchEvent)
	b.Session.AddHandler(dispatchReady)
	b.Session.AddHandler(dispatchChannelCreate)
	b.Session.AddHandler(dispatchChannelUpdate)
	b.Session.AddHandler(dispatchChannelDelete)
	b.Session.AddHandler(dispatchChannelPinsUpdate)
	b.Session.AddHandler(dispatchGuildCreate)
	b.Session.AddHandler(dispatchGuildUpdate)
	b.Session.AddHandler(dispatchGuildDelete)
	b.Session.AddHandler(dispatchGuildBanAdd)
	b.Session.AddHandler(dispatchGuildBanRemove)
	b.Session.AddHandler(dispatchGuildMemberAdd)
	b.Session.AddHandler(dispatchGuildMemberUpdate)
	b.Session.AddHandler(dispatchGuildMemberRemove)
	b.Session.AddHandler(dispatchGuildRoleCreate)
	b.Session.AddHandler(dispatchGuildRoleUpdate)
	b.Session.AddHandler(dispatchGuildRoleDelete)
	b.Session.AddHandler(dispatchGuildEmojisUpdate)
	b.Session.AddHandler(dispatchGuildMembersChunk)
	b.Session.AddHandler(dispatchGuildIntegrationsUpdate)
	b.Session.AddHandler(dispatchMessageAck)
	b.Session.AddHandler(dispatchMessageCreate)
	b.Session.AddHandler(dispatchMessageUpdate)
	b.Session.AddHandler(dispatchMessageDelete)
	b.Session.AddHandler(dispatchMessageReactionAdd)
	b.Session.AddHandler(dispatchMessageReactionRemove)
	b.Session.AddHandler(dispatchPresencesReplace)
	b.Session.AddHandler(dispatchPresenceUpdate)
	b.Session.AddHandler(dispatchResumed)
	b.Session.AddHandler(dispatchRelationshipAdd)
	b.Session.AddHandler(dispatchRelationshipRemove)
	b.Session.AddHandler(dispatchTypingStart)
	b.Session.AddHandler(dispatchUserUpdate)
	b.Session.AddHandler(dispatchUserSettingsUpdate)
	b.Session.AddHandler(dispatchUserGuildSettingsUpdate)
	b.Session.AddHandler(dispatchVoiceServerUpdate)
	b.Session.AddHandler(dispatchVoiceStateUpdate)
}
