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
	
	ConnectListeners         map[string][]Listener
	
	DisconnectListeners         map[string][]Listener
	
	RateLimitListeners         map[string][]Listener
	
	EventListeners         map[string][]Listener
	
	ReadyListeners         map[string][]Listener
	
	ChannelCreateListeners         map[string][]Listener
	
	ChannelUpdateListeners         map[string][]Listener
	
	ChannelDeleteListeners         map[string][]Listener
	
	ChannelPinsUpdateListeners         map[string][]Listener
	
	GuildCreateListeners         map[string][]Listener
	
	GuildUpdateListeners         map[string][]Listener
	
	GuildDeleteListeners         map[string][]Listener
	
	GuildBanAddListeners         map[string][]Listener
	
	GuildBanRemoveListeners         map[string][]Listener
	
	GuildMemberAddListeners         map[string][]Listener
	
	GuildMemberUpdateListeners         map[string][]Listener
	
	GuildMemberRemoveListeners         map[string][]Listener
	
	GuildRoleCreateListeners         map[string][]Listener
	
	GuildRoleUpdateListeners         map[string][]Listener
	
	GuildRoleDeleteListeners         map[string][]Listener
	
	GuildEmojisUpdateListeners         map[string][]Listener
	
	GuildMembersChunkListeners         map[string][]Listener
	
	GuildIntegrationsUpdateListeners         map[string][]Listener
	
	MessageAckListeners         map[string][]Listener
	
	MessageCreateListeners         map[string][]Listener
	
	MessageUpdateListeners         map[string][]Listener
	
	MessageDeleteListeners         map[string][]Listener
	
	MessageReactionAddListeners         map[string][]Listener
	
	MessageReactionRemoveListeners         map[string][]Listener
	
	MessageReactionRemoveAllListeners         map[string][]Listener
	
	PresencesReplaceListeners         map[string][]Listener
	
	PresenceUpdateListeners         map[string][]Listener
	
	ResumedListeners         map[string][]Listener
	
	RelationshipAddListeners         map[string][]Listener
	
	RelationshipRemoveListeners         map[string][]Listener
	
	TypingStartListeners         map[string][]Listener
	
	UserUpdateListeners         map[string][]Listener
	
	UserSettingsUpdateListeners         map[string][]Listener
	
	UserGuildSettingsUpdateListeners         map[string][]Listener
	
	UserNoteUpdateListeners         map[string][]Listener
	
	VoiceServerUpdateListeners         map[string][]Listener
	
	VoiceStateUpdateListeners         map[string][]Listener
	
	MessageDeleteBulkListeners         map[string][]Listener
	
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
	
	r.MessageReactionRemoveAllListeners = make(map[string][]Listener)
	
	r.PresencesReplaceListeners = make(map[string][]Listener)
	
	r.PresenceUpdateListeners = make(map[string][]Listener)
	
	r.ResumedListeners = make(map[string][]Listener)
	
	r.RelationshipAddListeners = make(map[string][]Listener)
	
	r.RelationshipRemoveListeners = make(map[string][]Listener)
	
	r.TypingStartListeners = make(map[string][]Listener)
	
	r.UserUpdateListeners = make(map[string][]Listener)
	
	r.UserSettingsUpdateListeners = make(map[string][]Listener)
	
	r.UserGuildSettingsUpdateListeners = make(map[string][]Listener)
	
	r.UserNoteUpdateListeners = make(map[string][]Listener)
	
	r.VoiceServerUpdateListeners = make(map[string][]Listener)
	
	r.VoiceStateUpdateListeners = make(map[string][]Listener)
	
	r.MessageDeleteBulkListeners = make(map[string][]Listener)
	
}


func (h Handler) runEventConnect(s *discordgo.Session, e *discordgo.Connect) {
	for _, listeners := range h.registry.ConnectListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventDisconnect(s *discordgo.Session, e *discordgo.Disconnect) {
	for _, listeners := range h.registry.DisconnectListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventRateLimit(s *discordgo.Session, e *discordgo.RateLimit) {
	for _, listeners := range h.registry.RateLimitListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventEvent(s *discordgo.Session, e *discordgo.Event) {
	for _, listeners := range h.registry.EventListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
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

func (h Handler) runEventChannelCreate(s *discordgo.Session, e *discordgo.ChannelCreate) {
	for _, listeners := range h.registry.ChannelCreateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventChannelUpdate(s *discordgo.Session, e *discordgo.ChannelUpdate) {
	for _, listeners := range h.registry.ChannelUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventChannelDelete(s *discordgo.Session, e *discordgo.ChannelDelete) {
	for _, listeners := range h.registry.ChannelDeleteListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventChannelPinsUpdate(s *discordgo.Session, e *discordgo.ChannelPinsUpdate) {
	for _, listeners := range h.registry.ChannelPinsUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	for _, listeners := range h.registry.GuildCreateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildUpdate(s *discordgo.Session, e *discordgo.GuildUpdate) {
	for _, listeners := range h.registry.GuildUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	for _, listeners := range h.registry.GuildDeleteListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildBanAdd(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	for _, listeners := range h.registry.GuildBanAddListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildBanRemove(s *discordgo.Session, e *discordgo.GuildBanRemove) {
	for _, listeners := range h.registry.GuildBanRemoveListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildMemberAdd(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	for _, listeners := range h.registry.GuildMemberAddListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildMemberUpdate(s *discordgo.Session, e *discordgo.GuildMemberUpdate) {
	for _, listeners := range h.registry.GuildMemberUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildMemberRemove(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	for _, listeners := range h.registry.GuildMemberRemoveListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildRoleCreate(s *discordgo.Session, e *discordgo.GuildRoleCreate) {
	for _, listeners := range h.registry.GuildRoleCreateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildRoleUpdate(s *discordgo.Session, e *discordgo.GuildRoleUpdate) {
	for _, listeners := range h.registry.GuildRoleUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildRoleDelete(s *discordgo.Session, e *discordgo.GuildRoleDelete) {
	for _, listeners := range h.registry.GuildRoleDeleteListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildEmojisUpdate(s *discordgo.Session, e *discordgo.GuildEmojisUpdate) {
	for _, listeners := range h.registry.GuildEmojisUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildMembersChunk(s *discordgo.Session, e *discordgo.GuildMembersChunk) {
	for _, listeners := range h.registry.GuildMembersChunkListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventGuildIntegrationsUpdate(s *discordgo.Session, e *discordgo.GuildIntegrationsUpdate) {
	for _, listeners := range h.registry.GuildIntegrationsUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageAck(s *discordgo.Session, e *discordgo.MessageAck) {
	for _, listeners := range h.registry.MessageAckListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	for _, listeners := range h.registry.MessageCreateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	for _, listeners := range h.registry.MessageUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageDelete(s *discordgo.Session, e *discordgo.MessageDelete) {
	for _, listeners := range h.registry.MessageDeleteListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	for _, listeners := range h.registry.MessageReactionAddListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	for _, listeners := range h.registry.MessageReactionRemoveListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageReactionRemoveAll(s *discordgo.Session, e *discordgo.MessageReactionRemoveAll) {
	for _, listeners := range h.registry.MessageReactionRemoveAllListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventPresencesReplace(s *discordgo.Session, e *discordgo.PresencesReplace) {
	for _, listeners := range h.registry.PresencesReplaceListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventPresenceUpdate(s *discordgo.Session, e *discordgo.PresenceUpdate) {
	for _, listeners := range h.registry.PresenceUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventResumed(s *discordgo.Session, e *discordgo.Resumed) {
	for _, listeners := range h.registry.ResumedListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventRelationshipAdd(s *discordgo.Session, e *discordgo.RelationshipAdd) {
	for _, listeners := range h.registry.RelationshipAddListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventRelationshipRemove(s *discordgo.Session, e *discordgo.RelationshipRemove) {
	for _, listeners := range h.registry.RelationshipRemoveListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventTypingStart(s *discordgo.Session, e *discordgo.TypingStart) {
	for _, listeners := range h.registry.TypingStartListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventUserUpdate(s *discordgo.Session, e *discordgo.UserUpdate) {
	for _, listeners := range h.registry.UserUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventUserSettingsUpdate(s *discordgo.Session, e *discordgo.UserSettingsUpdate) {
	for _, listeners := range h.registry.UserSettingsUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventUserGuildSettingsUpdate(s *discordgo.Session, e *discordgo.UserGuildSettingsUpdate) {
	for _, listeners := range h.registry.UserGuildSettingsUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventUserNoteUpdate(s *discordgo.Session, e *discordgo.UserNoteUpdate) {
	for _, listeners := range h.registry.UserNoteUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventVoiceServerUpdate(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
	for _, listeners := range h.registry.VoiceServerUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventVoiceStateUpdate(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	for _, listeners := range h.registry.VoiceStateUpdateListeners {
		for _, l := range listeners {
			if ok := h.Status; ok {
				l.Run(e)
			} else {
				return
			}
		}
	}
}

func (h Handler) runEventMessageDeleteBulk(s *discordgo.Session, e *discordgo.MessageDeleteBulk) {
	for _, listeners := range h.registry.MessageDeleteBulkListeners {
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
	
	case *discordgo.Connect:
		r.ConnectListeners[l.Owner()] = append(r.ConnectListeners[l.Owner()], l)
    
	case *discordgo.Disconnect:
		r.DisconnectListeners[l.Owner()] = append(r.DisconnectListeners[l.Owner()], l)
    
	case *discordgo.RateLimit:
		r.RateLimitListeners[l.Owner()] = append(r.RateLimitListeners[l.Owner()], l)
    
	case *discordgo.Event:
		r.EventListeners[l.Owner()] = append(r.EventListeners[l.Owner()], l)
    
	case *discordgo.Ready:
		r.ReadyListeners[l.Owner()] = append(r.ReadyListeners[l.Owner()], l)
    
	case *discordgo.ChannelCreate:
		r.ChannelCreateListeners[l.Owner()] = append(r.ChannelCreateListeners[l.Owner()], l)
    
	case *discordgo.ChannelUpdate:
		r.ChannelUpdateListeners[l.Owner()] = append(r.ChannelUpdateListeners[l.Owner()], l)
    
	case *discordgo.ChannelDelete:
		r.ChannelDeleteListeners[l.Owner()] = append(r.ChannelDeleteListeners[l.Owner()], l)
    
	case *discordgo.ChannelPinsUpdate:
		r.ChannelPinsUpdateListeners[l.Owner()] = append(r.ChannelPinsUpdateListeners[l.Owner()], l)
    
	case *discordgo.GuildCreate:
		r.GuildCreateListeners[l.Owner()] = append(r.GuildCreateListeners[l.Owner()], l)
    
	case *discordgo.GuildUpdate:
		r.GuildUpdateListeners[l.Owner()] = append(r.GuildUpdateListeners[l.Owner()], l)
    
	case *discordgo.GuildDelete:
		r.GuildDeleteListeners[l.Owner()] = append(r.GuildDeleteListeners[l.Owner()], l)
    
	case *discordgo.GuildBanAdd:
		r.GuildBanAddListeners[l.Owner()] = append(r.GuildBanAddListeners[l.Owner()], l)
    
	case *discordgo.GuildBanRemove:
		r.GuildBanRemoveListeners[l.Owner()] = append(r.GuildBanRemoveListeners[l.Owner()], l)
    
	case *discordgo.GuildMemberAdd:
		r.GuildMemberAddListeners[l.Owner()] = append(r.GuildMemberAddListeners[l.Owner()], l)
    
	case *discordgo.GuildMemberUpdate:
		r.GuildMemberUpdateListeners[l.Owner()] = append(r.GuildMemberUpdateListeners[l.Owner()], l)
    
	case *discordgo.GuildMemberRemove:
		r.GuildMemberRemoveListeners[l.Owner()] = append(r.GuildMemberRemoveListeners[l.Owner()], l)
    
	case *discordgo.GuildRoleCreate:
		r.GuildRoleCreateListeners[l.Owner()] = append(r.GuildRoleCreateListeners[l.Owner()], l)
    
	case *discordgo.GuildRoleUpdate:
		r.GuildRoleUpdateListeners[l.Owner()] = append(r.GuildRoleUpdateListeners[l.Owner()], l)
    
	case *discordgo.GuildRoleDelete:
		r.GuildRoleDeleteListeners[l.Owner()] = append(r.GuildRoleDeleteListeners[l.Owner()], l)
    
	case *discordgo.GuildEmojisUpdate:
		r.GuildEmojisUpdateListeners[l.Owner()] = append(r.GuildEmojisUpdateListeners[l.Owner()], l)
    
	case *discordgo.GuildMembersChunk:
		r.GuildMembersChunkListeners[l.Owner()] = append(r.GuildMembersChunkListeners[l.Owner()], l)
    
	case *discordgo.GuildIntegrationsUpdate:
		r.GuildIntegrationsUpdateListeners[l.Owner()] = append(r.GuildIntegrationsUpdateListeners[l.Owner()], l)
    
	case *discordgo.MessageAck:
		r.MessageAckListeners[l.Owner()] = append(r.MessageAckListeners[l.Owner()], l)
    
	case *discordgo.MessageCreate:
		r.MessageCreateListeners[l.Owner()] = append(r.MessageCreateListeners[l.Owner()], l)
    
	case *discordgo.MessageUpdate:
		r.MessageUpdateListeners[l.Owner()] = append(r.MessageUpdateListeners[l.Owner()], l)
    
	case *discordgo.MessageDelete:
		r.MessageDeleteListeners[l.Owner()] = append(r.MessageDeleteListeners[l.Owner()], l)
    
	case *discordgo.MessageReactionAdd:
		r.MessageReactionAddListeners[l.Owner()] = append(r.MessageReactionAddListeners[l.Owner()], l)
    
	case *discordgo.MessageReactionRemove:
		r.MessageReactionRemoveListeners[l.Owner()] = append(r.MessageReactionRemoveListeners[l.Owner()], l)
    
	case *discordgo.MessageReactionRemoveAll:
		r.MessageReactionRemoveAllListeners[l.Owner()] = append(r.MessageReactionRemoveAllListeners[l.Owner()], l)
    
	case *discordgo.PresencesReplace:
		r.PresencesReplaceListeners[l.Owner()] = append(r.PresencesReplaceListeners[l.Owner()], l)
    
	case *discordgo.PresenceUpdate:
		r.PresenceUpdateListeners[l.Owner()] = append(r.PresenceUpdateListeners[l.Owner()], l)
    
	case *discordgo.Resumed:
		r.ResumedListeners[l.Owner()] = append(r.ResumedListeners[l.Owner()], l)
    
	case *discordgo.RelationshipAdd:
		r.RelationshipAddListeners[l.Owner()] = append(r.RelationshipAddListeners[l.Owner()], l)
    
	case *discordgo.RelationshipRemove:
		r.RelationshipRemoveListeners[l.Owner()] = append(r.RelationshipRemoveListeners[l.Owner()], l)
    
	case *discordgo.TypingStart:
		r.TypingStartListeners[l.Owner()] = append(r.TypingStartListeners[l.Owner()], l)
    
	case *discordgo.UserUpdate:
		r.UserUpdateListeners[l.Owner()] = append(r.UserUpdateListeners[l.Owner()], l)
    
	case *discordgo.UserSettingsUpdate:
		r.UserSettingsUpdateListeners[l.Owner()] = append(r.UserSettingsUpdateListeners[l.Owner()], l)
    
	case *discordgo.UserGuildSettingsUpdate:
		r.UserGuildSettingsUpdateListeners[l.Owner()] = append(r.UserGuildSettingsUpdateListeners[l.Owner()], l)
    
	case *discordgo.UserNoteUpdate:
		r.UserNoteUpdateListeners[l.Owner()] = append(r.UserNoteUpdateListeners[l.Owner()], l)
    
	case *discordgo.VoiceServerUpdate:
		r.VoiceServerUpdateListeners[l.Owner()] = append(r.VoiceServerUpdateListeners[l.Owner()], l)
    
	case *discordgo.VoiceStateUpdate:
		r.VoiceStateUpdateListeners[l.Owner()] = append(r.VoiceStateUpdateListeners[l.Owner()], l)
    
	case *discordgo.MessageDeleteBulk:
		r.MessageDeleteBulkListeners[l.Owner()] = append(r.MessageDeleteBulkListeners[l.Owner()], l)
    
	}
}

func runAction(v interface{}, act Action) {
	switch e := v.(type) {
    
	case *discordgo.Connect:
		f := act.(func(e *discordgo.Connect))
		f(e)
	
	case *discordgo.Disconnect:
		f := act.(func(e *discordgo.Disconnect))
		f(e)
	
	case *discordgo.RateLimit:
		f := act.(func(e *discordgo.RateLimit))
		f(e)
	
	case *discordgo.Event:
		f := act.(func(e *discordgo.Event))
		f(e)
	
	case *discordgo.Ready:
		f := act.(func(e *discordgo.Ready))
		f(e)
	
	case *discordgo.ChannelCreate:
		f := act.(func(e *discordgo.ChannelCreate))
		f(e)
	
	case *discordgo.ChannelUpdate:
		f := act.(func(e *discordgo.ChannelUpdate))
		f(e)
	
	case *discordgo.ChannelDelete:
		f := act.(func(e *discordgo.ChannelDelete))
		f(e)
	
	case *discordgo.ChannelPinsUpdate:
		f := act.(func(e *discordgo.ChannelPinsUpdate))
		f(e)
	
	case *discordgo.GuildCreate:
		f := act.(func(e *discordgo.GuildCreate))
		f(e)
	
	case *discordgo.GuildUpdate:
		f := act.(func(e *discordgo.GuildUpdate))
		f(e)
	
	case *discordgo.GuildDelete:
		f := act.(func(e *discordgo.GuildDelete))
		f(e)
	
	case *discordgo.GuildBanAdd:
		f := act.(func(e *discordgo.GuildBanAdd))
		f(e)
	
	case *discordgo.GuildBanRemove:
		f := act.(func(e *discordgo.GuildBanRemove))
		f(e)
	
	case *discordgo.GuildMemberAdd:
		f := act.(func(e *discordgo.GuildMemberAdd))
		f(e)
	
	case *discordgo.GuildMemberUpdate:
		f := act.(func(e *discordgo.GuildMemberUpdate))
		f(e)
	
	case *discordgo.GuildMemberRemove:
		f := act.(func(e *discordgo.GuildMemberRemove))
		f(e)
	
	case *discordgo.GuildRoleCreate:
		f := act.(func(e *discordgo.GuildRoleCreate))
		f(e)
	
	case *discordgo.GuildRoleUpdate:
		f := act.(func(e *discordgo.GuildRoleUpdate))
		f(e)
	
	case *discordgo.GuildRoleDelete:
		f := act.(func(e *discordgo.GuildRoleDelete))
		f(e)
	
	case *discordgo.GuildEmojisUpdate:
		f := act.(func(e *discordgo.GuildEmojisUpdate))
		f(e)
	
	case *discordgo.GuildMembersChunk:
		f := act.(func(e *discordgo.GuildMembersChunk))
		f(e)
	
	case *discordgo.GuildIntegrationsUpdate:
		f := act.(func(e *discordgo.GuildIntegrationsUpdate))
		f(e)
	
	case *discordgo.MessageAck:
		f := act.(func(e *discordgo.MessageAck))
		f(e)
	
	case *discordgo.MessageCreate:
		f := act.(func(e *discordgo.MessageCreate))
		f(e)
	
	case *discordgo.MessageUpdate:
		f := act.(func(e *discordgo.MessageUpdate))
		f(e)
	
	case *discordgo.MessageDelete:
		f := act.(func(e *discordgo.MessageDelete))
		f(e)
	
	case *discordgo.MessageReactionAdd:
		f := act.(func(e *discordgo.MessageReactionAdd))
		f(e)
	
	case *discordgo.MessageReactionRemove:
		f := act.(func(e *discordgo.MessageReactionRemove))
		f(e)
	
	case *discordgo.MessageReactionRemoveAll:
		f := act.(func(e *discordgo.MessageReactionRemoveAll))
		f(e)
	
	case *discordgo.PresencesReplace:
		f := act.(func(e *discordgo.PresencesReplace))
		f(e)
	
	case *discordgo.PresenceUpdate:
		f := act.(func(e *discordgo.PresenceUpdate))
		f(e)
	
	case *discordgo.Resumed:
		f := act.(func(e *discordgo.Resumed))
		f(e)
	
	case *discordgo.RelationshipAdd:
		f := act.(func(e *discordgo.RelationshipAdd))
		f(e)
	
	case *discordgo.RelationshipRemove:
		f := act.(func(e *discordgo.RelationshipRemove))
		f(e)
	
	case *discordgo.TypingStart:
		f := act.(func(e *discordgo.TypingStart))
		f(e)
	
	case *discordgo.UserUpdate:
		f := act.(func(e *discordgo.UserUpdate))
		f(e)
	
	case *discordgo.UserSettingsUpdate:
		f := act.(func(e *discordgo.UserSettingsUpdate))
		f(e)
	
	case *discordgo.UserGuildSettingsUpdate:
		f := act.(func(e *discordgo.UserGuildSettingsUpdate))
		f(e)
	
	case *discordgo.UserNoteUpdate:
		f := act.(func(e *discordgo.UserNoteUpdate))
		f(e)
	
	case *discordgo.VoiceServerUpdate:
		f := act.(func(e *discordgo.VoiceServerUpdate))
		f(e)
	
	case *discordgo.VoiceStateUpdate:
		f := act.(func(e *discordgo.VoiceStateUpdate))
		f(e)
	
	case *discordgo.MessageDeleteBulk:
		f := act.(func(e *discordgo.MessageDeleteBulk))
		f(e)
		
	}
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
	b.Session.AddHandler(h.runEventMessageReactionRemoveAll) 
	b.Session.AddHandler(h.runEventPresencesReplace) 
	b.Session.AddHandler(h.runEventPresenceUpdate) 
	b.Session.AddHandler(h.runEventResumed) 
	b.Session.AddHandler(h.runEventRelationshipAdd) 
	b.Session.AddHandler(h.runEventRelationshipRemove) 
	b.Session.AddHandler(h.runEventTypingStart) 
	b.Session.AddHandler(h.runEventUserUpdate) 
	b.Session.AddHandler(h.runEventUserSettingsUpdate) 
	b.Session.AddHandler(h.runEventUserGuildSettingsUpdate) 
	b.Session.AddHandler(h.runEventUserNoteUpdate) 
	b.Session.AddHandler(h.runEventVoiceServerUpdate) 
	b.Session.AddHandler(h.runEventVoiceStateUpdate) 
	b.Session.AddHandler(h.runEventMessageDeleteBulk) 
}
