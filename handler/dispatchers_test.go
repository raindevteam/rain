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
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
)

var confirmDo = make(chan bool, 1)


func TestConnectHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ConnectHandler
		args args
	}{
		{
			name: "run do",
			eh: ConnectHandler(func(e *discordgo.Connect) {
				confirmDo <- true
			}),
			args: args{&discordgo.Connect{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestDisconnectHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   DisconnectHandler
		args args
	}{
		{
			name: "run do",
			eh: DisconnectHandler(func(e *discordgo.Disconnect) {
				confirmDo <- true
			}),
			args: args{&discordgo.Disconnect{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestRateLimitHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   RateLimitHandler
		args args
	}{
		{
			name: "run do",
			eh: RateLimitHandler(func(e *discordgo.RateLimit) {
				confirmDo <- true
			}),
			args: args{&discordgo.RateLimit{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestEventHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   EventHandler
		args args
	}{
		{
			name: "run do",
			eh: EventHandler(func(e *discordgo.Event) {
				confirmDo <- true
			}),
			args: args{&discordgo.Event{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestReadyHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ReadyHandler
		args args
	}{
		{
			name: "run do",
			eh: ReadyHandler(func(e *discordgo.Ready) {
				confirmDo <- true
			}),
			args: args{&discordgo.Ready{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestChannelCreateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ChannelCreateHandler
		args args
	}{
		{
			name: "run do",
			eh: ChannelCreateHandler(func(e *discordgo.ChannelCreate) {
				confirmDo <- true
			}),
			args: args{&discordgo.ChannelCreate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestChannelUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ChannelUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: ChannelUpdateHandler(func(e *discordgo.ChannelUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.ChannelUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestChannelDeleteHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ChannelDeleteHandler
		args args
	}{
		{
			name: "run do",
			eh: ChannelDeleteHandler(func(e *discordgo.ChannelDelete) {
				confirmDo <- true
			}),
			args: args{&discordgo.ChannelDelete{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestChannelPinsUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ChannelPinsUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: ChannelPinsUpdateHandler(func(e *discordgo.ChannelPinsUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.ChannelPinsUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildCreateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildCreateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildCreateHandler(func(e *discordgo.GuildCreate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildCreate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildUpdateHandler(func(e *discordgo.GuildUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildDeleteHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildDeleteHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildDeleteHandler(func(e *discordgo.GuildDelete) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildDelete{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildBanAddHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildBanAddHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildBanAddHandler(func(e *discordgo.GuildBanAdd) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildBanAdd{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildBanRemoveHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildBanRemoveHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildBanRemoveHandler(func(e *discordgo.GuildBanRemove) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildBanRemove{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildMemberAddHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildMemberAddHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildMemberAddHandler(func(e *discordgo.GuildMemberAdd) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildMemberAdd{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildMemberUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildMemberUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildMemberUpdateHandler(func(e *discordgo.GuildMemberUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildMemberUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildMemberRemoveHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildMemberRemoveHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildMemberRemoveHandler(func(e *discordgo.GuildMemberRemove) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildMemberRemove{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildRoleCreateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildRoleCreateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildRoleCreateHandler(func(e *discordgo.GuildRoleCreate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildRoleCreate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildRoleUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildRoleUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildRoleUpdateHandler(func(e *discordgo.GuildRoleUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildRoleUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildRoleDeleteHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildRoleDeleteHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildRoleDeleteHandler(func(e *discordgo.GuildRoleDelete) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildRoleDelete{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildEmojisUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildEmojisUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildEmojisUpdateHandler(func(e *discordgo.GuildEmojisUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildEmojisUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildMembersChunkHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildMembersChunkHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildMembersChunkHandler(func(e *discordgo.GuildMembersChunk) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildMembersChunk{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestGuildIntegrationsUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   GuildIntegrationsUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: GuildIntegrationsUpdateHandler(func(e *discordgo.GuildIntegrationsUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.GuildIntegrationsUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestMessageAckHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   MessageAckHandler
		args args
	}{
		{
			name: "run do",
			eh: MessageAckHandler(func(e *discordgo.MessageAck) {
				confirmDo <- true
			}),
			args: args{&discordgo.MessageAck{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestMessageCreateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   MessageCreateHandler
		args args
	}{
		{
			name: "run do",
			eh: MessageCreateHandler(func(e *discordgo.MessageCreate) {
				confirmDo <- true
			}),
			args: args{&discordgo.MessageCreate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestMessageUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   MessageUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: MessageUpdateHandler(func(e *discordgo.MessageUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.MessageUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestMessageDeleteHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   MessageDeleteHandler
		args args
	}{
		{
			name: "run do",
			eh: MessageDeleteHandler(func(e *discordgo.MessageDelete) {
				confirmDo <- true
			}),
			args: args{&discordgo.MessageDelete{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestMessageReactionAddHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   MessageReactionAddHandler
		args args
	}{
		{
			name: "run do",
			eh: MessageReactionAddHandler(func(e *discordgo.MessageReactionAdd) {
				confirmDo <- true
			}),
			args: args{&discordgo.MessageReactionAdd{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestMessageReactionRemoveHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   MessageReactionRemoveHandler
		args args
	}{
		{
			name: "run do",
			eh: MessageReactionRemoveHandler(func(e *discordgo.MessageReactionRemove) {
				confirmDo <- true
			}),
			args: args{&discordgo.MessageReactionRemove{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestPresencesReplaceHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   PresencesReplaceHandler
		args args
	}{
		{
			name: "run do",
			eh: PresencesReplaceHandler(func(e *discordgo.PresencesReplace) {
				confirmDo <- true
			}),
			args: args{&discordgo.PresencesReplace{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestPresenceUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   PresenceUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: PresenceUpdateHandler(func(e *discordgo.PresenceUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.PresenceUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestResumedHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   ResumedHandler
		args args
	}{
		{
			name: "run do",
			eh: ResumedHandler(func(e *discordgo.Resumed) {
				confirmDo <- true
			}),
			args: args{&discordgo.Resumed{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestRelationshipAddHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   RelationshipAddHandler
		args args
	}{
		{
			name: "run do",
			eh: RelationshipAddHandler(func(e *discordgo.RelationshipAdd) {
				confirmDo <- true
			}),
			args: args{&discordgo.RelationshipAdd{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestRelationshipRemoveHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   RelationshipRemoveHandler
		args args
	}{
		{
			name: "run do",
			eh: RelationshipRemoveHandler(func(e *discordgo.RelationshipRemove) {
				confirmDo <- true
			}),
			args: args{&discordgo.RelationshipRemove{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestTypingStartHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   TypingStartHandler
		args args
	}{
		{
			name: "run do",
			eh: TypingStartHandler(func(e *discordgo.TypingStart) {
				confirmDo <- true
			}),
			args: args{&discordgo.TypingStart{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestUserUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   UserUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: UserUpdateHandler(func(e *discordgo.UserUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.UserUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestUserSettingsUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   UserSettingsUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: UserSettingsUpdateHandler(func(e *discordgo.UserSettingsUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.UserSettingsUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestUserGuildSettingsUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   UserGuildSettingsUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: UserGuildSettingsUpdateHandler(func(e *discordgo.UserGuildSettingsUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.UserGuildSettingsUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestVoiceServerUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   VoiceServerUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: VoiceServerUpdateHandler(func(e *discordgo.VoiceServerUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.VoiceServerUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}

func TestVoiceStateUpdateHandler_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		eh   VoiceStateUpdateHandler
		args args
	}{
		{
			name: "run do",
			eh: VoiceStateUpdateHandler(func(e *discordgo.VoiceStateUpdate) {
				confirmDo <- true
			}),
			args: args{&discordgo.VoiceStateUpdate{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.eh.Do(tt.args.v)
			select {
			case <-confirmDo:
			case <-time.After(time.Second * 1):
				t.Errorf("listener timed out.")
			}
		})
	}
}


