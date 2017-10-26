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



func Test_dispatchConnect(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.Connect
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch Connect listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Connect{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch Connect listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Connect{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch Connect listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Connect{},
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
				dispatchConnect(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ConnectHandler(func(e *discordgo.Connect) {
				l1chan <- true
			}))
			H.registry.ConnectListeners[l.Owner()] = append(H.registry.ConnectListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchConnect(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ConnectHandler(func(e *discordgo.Connect) {
				l1chan <- true
			}))
			H.registry.ConnectListeners[l1.Owner()] = append(H.registry.ConnectListeners[l1.Owner()], l1)
			l2.SetActionHandler(ConnectHandler(func(e *discordgo.Connect) {
				l1chan <- true
			}))
			H.registry.ConnectListeners[l2.Owner()] = append(H.registry.ConnectListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchConnect(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchDisconnect(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.Disconnect
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch Disconnect listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Disconnect{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch Disconnect listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Disconnect{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch Disconnect listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Disconnect{},
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
				dispatchDisconnect(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(DisconnectHandler(func(e *discordgo.Disconnect) {
				l1chan <- true
			}))
			H.registry.DisconnectListeners[l.Owner()] = append(H.registry.DisconnectListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchDisconnect(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(DisconnectHandler(func(e *discordgo.Disconnect) {
				l1chan <- true
			}))
			H.registry.DisconnectListeners[l1.Owner()] = append(H.registry.DisconnectListeners[l1.Owner()], l1)
			l2.SetActionHandler(DisconnectHandler(func(e *discordgo.Disconnect) {
				l1chan <- true
			}))
			H.registry.DisconnectListeners[l2.Owner()] = append(H.registry.DisconnectListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchDisconnect(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchRateLimit(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.RateLimit
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch RateLimit listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RateLimit{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch RateLimit listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RateLimit{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch RateLimit listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RateLimit{},
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
				dispatchRateLimit(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(RateLimitHandler(func(e *discordgo.RateLimit) {
				l1chan <- true
			}))
			H.registry.RateLimitListeners[l.Owner()] = append(H.registry.RateLimitListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchRateLimit(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(RateLimitHandler(func(e *discordgo.RateLimit) {
				l1chan <- true
			}))
			H.registry.RateLimitListeners[l1.Owner()] = append(H.registry.RateLimitListeners[l1.Owner()], l1)
			l2.SetActionHandler(RateLimitHandler(func(e *discordgo.RateLimit) {
				l1chan <- true
			}))
			H.registry.RateLimitListeners[l2.Owner()] = append(H.registry.RateLimitListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchRateLimit(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchEvent(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.Event
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch Event listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Event{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch Event listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Event{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch Event listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Event{},
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
				dispatchEvent(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(EventHandler(func(e *discordgo.Event) {
				l1chan <- true
			}))
			H.registry.EventListeners[l.Owner()] = append(H.registry.EventListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchEvent(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(EventHandler(func(e *discordgo.Event) {
				l1chan <- true
			}))
			H.registry.EventListeners[l1.Owner()] = append(H.registry.EventListeners[l1.Owner()], l1)
			l2.SetActionHandler(EventHandler(func(e *discordgo.Event) {
				l1chan <- true
			}))
			H.registry.EventListeners[l2.Owner()] = append(H.registry.EventListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchEvent(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchReady(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.Ready
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch Ready listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Ready{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch Ready listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Ready{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch Ready listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Ready{},
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
				dispatchReady(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ReadyHandler(func(e *discordgo.Ready) {
				l1chan <- true
			}))
			H.registry.ReadyListeners[l.Owner()] = append(H.registry.ReadyListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchReady(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ReadyHandler(func(e *discordgo.Ready) {
				l1chan <- true
			}))
			H.registry.ReadyListeners[l1.Owner()] = append(H.registry.ReadyListeners[l1.Owner()], l1)
			l2.SetActionHandler(ReadyHandler(func(e *discordgo.Ready) {
				l1chan <- true
			}))
			H.registry.ReadyListeners[l2.Owner()] = append(H.registry.ReadyListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchReady(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchChannelCreate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.ChannelCreate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch ChannelCreate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelCreate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch ChannelCreate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelCreate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch ChannelCreate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelCreate{},
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
				dispatchChannelCreate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ChannelCreateHandler(func(e *discordgo.ChannelCreate) {
				l1chan <- true
			}))
			H.registry.ChannelCreateListeners[l.Owner()] = append(H.registry.ChannelCreateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelCreate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ChannelCreateHandler(func(e *discordgo.ChannelCreate) {
				l1chan <- true
			}))
			H.registry.ChannelCreateListeners[l1.Owner()] = append(H.registry.ChannelCreateListeners[l1.Owner()], l1)
			l2.SetActionHandler(ChannelCreateHandler(func(e *discordgo.ChannelCreate) {
				l1chan <- true
			}))
			H.registry.ChannelCreateListeners[l2.Owner()] = append(H.registry.ChannelCreateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelCreate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchChannelUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.ChannelUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch ChannelUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch ChannelUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch ChannelUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelUpdate{},
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
				dispatchChannelUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ChannelUpdateHandler(func(e *discordgo.ChannelUpdate) {
				l1chan <- true
			}))
			H.registry.ChannelUpdateListeners[l.Owner()] = append(H.registry.ChannelUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ChannelUpdateHandler(func(e *discordgo.ChannelUpdate) {
				l1chan <- true
			}))
			H.registry.ChannelUpdateListeners[l1.Owner()] = append(H.registry.ChannelUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(ChannelUpdateHandler(func(e *discordgo.ChannelUpdate) {
				l1chan <- true
			}))
			H.registry.ChannelUpdateListeners[l2.Owner()] = append(H.registry.ChannelUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchChannelDelete(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.ChannelDelete
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch ChannelDelete listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelDelete{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch ChannelDelete listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelDelete{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch ChannelDelete listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelDelete{},
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
				dispatchChannelDelete(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ChannelDeleteHandler(func(e *discordgo.ChannelDelete) {
				l1chan <- true
			}))
			H.registry.ChannelDeleteListeners[l.Owner()] = append(H.registry.ChannelDeleteListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelDelete(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ChannelDeleteHandler(func(e *discordgo.ChannelDelete) {
				l1chan <- true
			}))
			H.registry.ChannelDeleteListeners[l1.Owner()] = append(H.registry.ChannelDeleteListeners[l1.Owner()], l1)
			l2.SetActionHandler(ChannelDeleteHandler(func(e *discordgo.ChannelDelete) {
				l1chan <- true
			}))
			H.registry.ChannelDeleteListeners[l2.Owner()] = append(H.registry.ChannelDeleteListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelDelete(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchChannelPinsUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.ChannelPinsUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch ChannelPinsUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelPinsUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch ChannelPinsUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelPinsUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch ChannelPinsUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.ChannelPinsUpdate{},
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
				dispatchChannelPinsUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ChannelPinsUpdateHandler(func(e *discordgo.ChannelPinsUpdate) {
				l1chan <- true
			}))
			H.registry.ChannelPinsUpdateListeners[l.Owner()] = append(H.registry.ChannelPinsUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelPinsUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ChannelPinsUpdateHandler(func(e *discordgo.ChannelPinsUpdate) {
				l1chan <- true
			}))
			H.registry.ChannelPinsUpdateListeners[l1.Owner()] = append(H.registry.ChannelPinsUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(ChannelPinsUpdateHandler(func(e *discordgo.ChannelPinsUpdate) {
				l1chan <- true
			}))
			H.registry.ChannelPinsUpdateListeners[l2.Owner()] = append(H.registry.ChannelPinsUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchChannelPinsUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildCreate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildCreate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildCreate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildCreate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildCreate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildCreate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildCreate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildCreate{},
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
				dispatchGuildCreate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildCreateHandler(func(e *discordgo.GuildCreate) {
				l1chan <- true
			}))
			H.registry.GuildCreateListeners[l.Owner()] = append(H.registry.GuildCreateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildCreate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildCreateHandler(func(e *discordgo.GuildCreate) {
				l1chan <- true
			}))
			H.registry.GuildCreateListeners[l1.Owner()] = append(H.registry.GuildCreateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildCreateHandler(func(e *discordgo.GuildCreate) {
				l1chan <- true
			}))
			H.registry.GuildCreateListeners[l2.Owner()] = append(H.registry.GuildCreateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildCreate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildUpdate{},
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
				dispatchGuildUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildUpdateHandler(func(e *discordgo.GuildUpdate) {
				l1chan <- true
			}))
			H.registry.GuildUpdateListeners[l.Owner()] = append(H.registry.GuildUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildUpdateHandler(func(e *discordgo.GuildUpdate) {
				l1chan <- true
			}))
			H.registry.GuildUpdateListeners[l1.Owner()] = append(H.registry.GuildUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildUpdateHandler(func(e *discordgo.GuildUpdate) {
				l1chan <- true
			}))
			H.registry.GuildUpdateListeners[l2.Owner()] = append(H.registry.GuildUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildDelete(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildDelete
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildDelete listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildDelete{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildDelete listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildDelete{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildDelete listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildDelete{},
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
				dispatchGuildDelete(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildDeleteHandler(func(e *discordgo.GuildDelete) {
				l1chan <- true
			}))
			H.registry.GuildDeleteListeners[l.Owner()] = append(H.registry.GuildDeleteListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildDelete(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildDeleteHandler(func(e *discordgo.GuildDelete) {
				l1chan <- true
			}))
			H.registry.GuildDeleteListeners[l1.Owner()] = append(H.registry.GuildDeleteListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildDeleteHandler(func(e *discordgo.GuildDelete) {
				l1chan <- true
			}))
			H.registry.GuildDeleteListeners[l2.Owner()] = append(H.registry.GuildDeleteListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildDelete(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildBanAdd(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildBanAdd
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildBanAdd listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildBanAdd{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildBanAdd listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildBanAdd{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildBanAdd listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildBanAdd{},
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
				dispatchGuildBanAdd(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildBanAddHandler(func(e *discordgo.GuildBanAdd) {
				l1chan <- true
			}))
			H.registry.GuildBanAddListeners[l.Owner()] = append(H.registry.GuildBanAddListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildBanAdd(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildBanAddHandler(func(e *discordgo.GuildBanAdd) {
				l1chan <- true
			}))
			H.registry.GuildBanAddListeners[l1.Owner()] = append(H.registry.GuildBanAddListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildBanAddHandler(func(e *discordgo.GuildBanAdd) {
				l1chan <- true
			}))
			H.registry.GuildBanAddListeners[l2.Owner()] = append(H.registry.GuildBanAddListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildBanAdd(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildBanRemove(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildBanRemove
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildBanRemove listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildBanRemove{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildBanRemove listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildBanRemove{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildBanRemove listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildBanRemove{},
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
				dispatchGuildBanRemove(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildBanRemoveHandler(func(e *discordgo.GuildBanRemove) {
				l1chan <- true
			}))
			H.registry.GuildBanRemoveListeners[l.Owner()] = append(H.registry.GuildBanRemoveListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildBanRemove(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildBanRemoveHandler(func(e *discordgo.GuildBanRemove) {
				l1chan <- true
			}))
			H.registry.GuildBanRemoveListeners[l1.Owner()] = append(H.registry.GuildBanRemoveListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildBanRemoveHandler(func(e *discordgo.GuildBanRemove) {
				l1chan <- true
			}))
			H.registry.GuildBanRemoveListeners[l2.Owner()] = append(H.registry.GuildBanRemoveListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildBanRemove(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildMemberAdd(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildMemberAdd
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildMemberAdd listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberAdd{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildMemberAdd listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberAdd{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildMemberAdd listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberAdd{},
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
				dispatchGuildMemberAdd(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildMemberAddHandler(func(e *discordgo.GuildMemberAdd) {
				l1chan <- true
			}))
			H.registry.GuildMemberAddListeners[l.Owner()] = append(H.registry.GuildMemberAddListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMemberAdd(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildMemberAddHandler(func(e *discordgo.GuildMemberAdd) {
				l1chan <- true
			}))
			H.registry.GuildMemberAddListeners[l1.Owner()] = append(H.registry.GuildMemberAddListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildMemberAddHandler(func(e *discordgo.GuildMemberAdd) {
				l1chan <- true
			}))
			H.registry.GuildMemberAddListeners[l2.Owner()] = append(H.registry.GuildMemberAddListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMemberAdd(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildMemberUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildMemberUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildMemberUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildMemberUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildMemberUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberUpdate{},
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
				dispatchGuildMemberUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildMemberUpdateHandler(func(e *discordgo.GuildMemberUpdate) {
				l1chan <- true
			}))
			H.registry.GuildMemberUpdateListeners[l.Owner()] = append(H.registry.GuildMemberUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMemberUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildMemberUpdateHandler(func(e *discordgo.GuildMemberUpdate) {
				l1chan <- true
			}))
			H.registry.GuildMemberUpdateListeners[l1.Owner()] = append(H.registry.GuildMemberUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildMemberUpdateHandler(func(e *discordgo.GuildMemberUpdate) {
				l1chan <- true
			}))
			H.registry.GuildMemberUpdateListeners[l2.Owner()] = append(H.registry.GuildMemberUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMemberUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildMemberRemove(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildMemberRemove
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildMemberRemove listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberRemove{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildMemberRemove listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberRemove{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildMemberRemove listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMemberRemove{},
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
				dispatchGuildMemberRemove(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildMemberRemoveHandler(func(e *discordgo.GuildMemberRemove) {
				l1chan <- true
			}))
			H.registry.GuildMemberRemoveListeners[l.Owner()] = append(H.registry.GuildMemberRemoveListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMemberRemove(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildMemberRemoveHandler(func(e *discordgo.GuildMemberRemove) {
				l1chan <- true
			}))
			H.registry.GuildMemberRemoveListeners[l1.Owner()] = append(H.registry.GuildMemberRemoveListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildMemberRemoveHandler(func(e *discordgo.GuildMemberRemove) {
				l1chan <- true
			}))
			H.registry.GuildMemberRemoveListeners[l2.Owner()] = append(H.registry.GuildMemberRemoveListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMemberRemove(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildRoleCreate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildRoleCreate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildRoleCreate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleCreate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildRoleCreate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleCreate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildRoleCreate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleCreate{},
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
				dispatchGuildRoleCreate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildRoleCreateHandler(func(e *discordgo.GuildRoleCreate) {
				l1chan <- true
			}))
			H.registry.GuildRoleCreateListeners[l.Owner()] = append(H.registry.GuildRoleCreateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildRoleCreate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildRoleCreateHandler(func(e *discordgo.GuildRoleCreate) {
				l1chan <- true
			}))
			H.registry.GuildRoleCreateListeners[l1.Owner()] = append(H.registry.GuildRoleCreateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildRoleCreateHandler(func(e *discordgo.GuildRoleCreate) {
				l1chan <- true
			}))
			H.registry.GuildRoleCreateListeners[l2.Owner()] = append(H.registry.GuildRoleCreateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildRoleCreate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildRoleUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildRoleUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildRoleUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildRoleUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildRoleUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleUpdate{},
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
				dispatchGuildRoleUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildRoleUpdateHandler(func(e *discordgo.GuildRoleUpdate) {
				l1chan <- true
			}))
			H.registry.GuildRoleUpdateListeners[l.Owner()] = append(H.registry.GuildRoleUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildRoleUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildRoleUpdateHandler(func(e *discordgo.GuildRoleUpdate) {
				l1chan <- true
			}))
			H.registry.GuildRoleUpdateListeners[l1.Owner()] = append(H.registry.GuildRoleUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildRoleUpdateHandler(func(e *discordgo.GuildRoleUpdate) {
				l1chan <- true
			}))
			H.registry.GuildRoleUpdateListeners[l2.Owner()] = append(H.registry.GuildRoleUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildRoleUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildRoleDelete(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildRoleDelete
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildRoleDelete listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleDelete{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildRoleDelete listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleDelete{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildRoleDelete listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildRoleDelete{},
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
				dispatchGuildRoleDelete(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildRoleDeleteHandler(func(e *discordgo.GuildRoleDelete) {
				l1chan <- true
			}))
			H.registry.GuildRoleDeleteListeners[l.Owner()] = append(H.registry.GuildRoleDeleteListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildRoleDelete(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildRoleDeleteHandler(func(e *discordgo.GuildRoleDelete) {
				l1chan <- true
			}))
			H.registry.GuildRoleDeleteListeners[l1.Owner()] = append(H.registry.GuildRoleDeleteListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildRoleDeleteHandler(func(e *discordgo.GuildRoleDelete) {
				l1chan <- true
			}))
			H.registry.GuildRoleDeleteListeners[l2.Owner()] = append(H.registry.GuildRoleDeleteListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildRoleDelete(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildEmojisUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildEmojisUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildEmojisUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildEmojisUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildEmojisUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildEmojisUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildEmojisUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildEmojisUpdate{},
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
				dispatchGuildEmojisUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildEmojisUpdateHandler(func(e *discordgo.GuildEmojisUpdate) {
				l1chan <- true
			}))
			H.registry.GuildEmojisUpdateListeners[l.Owner()] = append(H.registry.GuildEmojisUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildEmojisUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildEmojisUpdateHandler(func(e *discordgo.GuildEmojisUpdate) {
				l1chan <- true
			}))
			H.registry.GuildEmojisUpdateListeners[l1.Owner()] = append(H.registry.GuildEmojisUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildEmojisUpdateHandler(func(e *discordgo.GuildEmojisUpdate) {
				l1chan <- true
			}))
			H.registry.GuildEmojisUpdateListeners[l2.Owner()] = append(H.registry.GuildEmojisUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildEmojisUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildMembersChunk(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildMembersChunk
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildMembersChunk listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMembersChunk{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildMembersChunk listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMembersChunk{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildMembersChunk listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildMembersChunk{},
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
				dispatchGuildMembersChunk(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildMembersChunkHandler(func(e *discordgo.GuildMembersChunk) {
				l1chan <- true
			}))
			H.registry.GuildMembersChunkListeners[l.Owner()] = append(H.registry.GuildMembersChunkListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMembersChunk(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildMembersChunkHandler(func(e *discordgo.GuildMembersChunk) {
				l1chan <- true
			}))
			H.registry.GuildMembersChunkListeners[l1.Owner()] = append(H.registry.GuildMembersChunkListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildMembersChunkHandler(func(e *discordgo.GuildMembersChunk) {
				l1chan <- true
			}))
			H.registry.GuildMembersChunkListeners[l2.Owner()] = append(H.registry.GuildMembersChunkListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildMembersChunk(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchGuildIntegrationsUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.GuildIntegrationsUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch GuildIntegrationsUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildIntegrationsUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch GuildIntegrationsUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildIntegrationsUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch GuildIntegrationsUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.GuildIntegrationsUpdate{},
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
				dispatchGuildIntegrationsUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(GuildIntegrationsUpdateHandler(func(e *discordgo.GuildIntegrationsUpdate) {
				l1chan <- true
			}))
			H.registry.GuildIntegrationsUpdateListeners[l.Owner()] = append(H.registry.GuildIntegrationsUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildIntegrationsUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(GuildIntegrationsUpdateHandler(func(e *discordgo.GuildIntegrationsUpdate) {
				l1chan <- true
			}))
			H.registry.GuildIntegrationsUpdateListeners[l1.Owner()] = append(H.registry.GuildIntegrationsUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(GuildIntegrationsUpdateHandler(func(e *discordgo.GuildIntegrationsUpdate) {
				l1chan <- true
			}))
			H.registry.GuildIntegrationsUpdateListeners[l2.Owner()] = append(H.registry.GuildIntegrationsUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchGuildIntegrationsUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchMessageAck(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.MessageAck
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch MessageAck listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageAck{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch MessageAck listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageAck{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch MessageAck listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageAck{},
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
				dispatchMessageAck(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(MessageAckHandler(func(e *discordgo.MessageAck) {
				l1chan <- true
			}))
			H.registry.MessageAckListeners[l.Owner()] = append(H.registry.MessageAckListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageAck(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(MessageAckHandler(func(e *discordgo.MessageAck) {
				l1chan <- true
			}))
			H.registry.MessageAckListeners[l1.Owner()] = append(H.registry.MessageAckListeners[l1.Owner()], l1)
			l2.SetActionHandler(MessageAckHandler(func(e *discordgo.MessageAck) {
				l1chan <- true
			}))
			H.registry.MessageAckListeners[l2.Owner()] = append(H.registry.MessageAckListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageAck(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchMessageCreate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.MessageCreate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch MessageCreate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageCreate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch MessageCreate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageCreate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch MessageCreate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageCreate{},
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
				dispatchMessageCreate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(MessageCreateHandler(func(e *discordgo.MessageCreate) {
				l1chan <- true
			}))
			H.registry.MessageCreateListeners[l.Owner()] = append(H.registry.MessageCreateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageCreate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(MessageCreateHandler(func(e *discordgo.MessageCreate) {
				l1chan <- true
			}))
			H.registry.MessageCreateListeners[l1.Owner()] = append(H.registry.MessageCreateListeners[l1.Owner()], l1)
			l2.SetActionHandler(MessageCreateHandler(func(e *discordgo.MessageCreate) {
				l1chan <- true
			}))
			H.registry.MessageCreateListeners[l2.Owner()] = append(H.registry.MessageCreateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageCreate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchMessageUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.MessageUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch MessageUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch MessageUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch MessageUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageUpdate{},
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
				dispatchMessageUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(MessageUpdateHandler(func(e *discordgo.MessageUpdate) {
				l1chan <- true
			}))
			H.registry.MessageUpdateListeners[l.Owner()] = append(H.registry.MessageUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(MessageUpdateHandler(func(e *discordgo.MessageUpdate) {
				l1chan <- true
			}))
			H.registry.MessageUpdateListeners[l1.Owner()] = append(H.registry.MessageUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(MessageUpdateHandler(func(e *discordgo.MessageUpdate) {
				l1chan <- true
			}))
			H.registry.MessageUpdateListeners[l2.Owner()] = append(H.registry.MessageUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchMessageDelete(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.MessageDelete
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch MessageDelete listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageDelete{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch MessageDelete listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageDelete{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch MessageDelete listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageDelete{},
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
				dispatchMessageDelete(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(MessageDeleteHandler(func(e *discordgo.MessageDelete) {
				l1chan <- true
			}))
			H.registry.MessageDeleteListeners[l.Owner()] = append(H.registry.MessageDeleteListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageDelete(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(MessageDeleteHandler(func(e *discordgo.MessageDelete) {
				l1chan <- true
			}))
			H.registry.MessageDeleteListeners[l1.Owner()] = append(H.registry.MessageDeleteListeners[l1.Owner()], l1)
			l2.SetActionHandler(MessageDeleteHandler(func(e *discordgo.MessageDelete) {
				l1chan <- true
			}))
			H.registry.MessageDeleteListeners[l2.Owner()] = append(H.registry.MessageDeleteListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageDelete(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchMessageReactionAdd(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.MessageReactionAdd
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch MessageReactionAdd listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageReactionAdd{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch MessageReactionAdd listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageReactionAdd{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch MessageReactionAdd listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageReactionAdd{},
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
				dispatchMessageReactionAdd(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(MessageReactionAddHandler(func(e *discordgo.MessageReactionAdd) {
				l1chan <- true
			}))
			H.registry.MessageReactionAddListeners[l.Owner()] = append(H.registry.MessageReactionAddListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageReactionAdd(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(MessageReactionAddHandler(func(e *discordgo.MessageReactionAdd) {
				l1chan <- true
			}))
			H.registry.MessageReactionAddListeners[l1.Owner()] = append(H.registry.MessageReactionAddListeners[l1.Owner()], l1)
			l2.SetActionHandler(MessageReactionAddHandler(func(e *discordgo.MessageReactionAdd) {
				l1chan <- true
			}))
			H.registry.MessageReactionAddListeners[l2.Owner()] = append(H.registry.MessageReactionAddListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageReactionAdd(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchMessageReactionRemove(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.MessageReactionRemove
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch MessageReactionRemove listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageReactionRemove{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch MessageReactionRemove listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageReactionRemove{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch MessageReactionRemove listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.MessageReactionRemove{},
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
				dispatchMessageReactionRemove(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(MessageReactionRemoveHandler(func(e *discordgo.MessageReactionRemove) {
				l1chan <- true
			}))
			H.registry.MessageReactionRemoveListeners[l.Owner()] = append(H.registry.MessageReactionRemoveListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageReactionRemove(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(MessageReactionRemoveHandler(func(e *discordgo.MessageReactionRemove) {
				l1chan <- true
			}))
			H.registry.MessageReactionRemoveListeners[l1.Owner()] = append(H.registry.MessageReactionRemoveListeners[l1.Owner()], l1)
			l2.SetActionHandler(MessageReactionRemoveHandler(func(e *discordgo.MessageReactionRemove) {
				l1chan <- true
			}))
			H.registry.MessageReactionRemoveListeners[l2.Owner()] = append(H.registry.MessageReactionRemoveListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchMessageReactionRemove(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchPresencesReplace(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.PresencesReplace
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch PresencesReplace listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.PresencesReplace{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch PresencesReplace listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.PresencesReplace{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch PresencesReplace listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.PresencesReplace{},
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
				dispatchPresencesReplace(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(PresencesReplaceHandler(func(e *discordgo.PresencesReplace) {
				l1chan <- true
			}))
			H.registry.PresencesReplaceListeners[l.Owner()] = append(H.registry.PresencesReplaceListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchPresencesReplace(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(PresencesReplaceHandler(func(e *discordgo.PresencesReplace) {
				l1chan <- true
			}))
			H.registry.PresencesReplaceListeners[l1.Owner()] = append(H.registry.PresencesReplaceListeners[l1.Owner()], l1)
			l2.SetActionHandler(PresencesReplaceHandler(func(e *discordgo.PresencesReplace) {
				l1chan <- true
			}))
			H.registry.PresencesReplaceListeners[l2.Owner()] = append(H.registry.PresencesReplaceListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchPresencesReplace(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchPresenceUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.PresenceUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch PresenceUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.PresenceUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch PresenceUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.PresenceUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch PresenceUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.PresenceUpdate{},
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
				dispatchPresenceUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(PresenceUpdateHandler(func(e *discordgo.PresenceUpdate) {
				l1chan <- true
			}))
			H.registry.PresenceUpdateListeners[l.Owner()] = append(H.registry.PresenceUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchPresenceUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(PresenceUpdateHandler(func(e *discordgo.PresenceUpdate) {
				l1chan <- true
			}))
			H.registry.PresenceUpdateListeners[l1.Owner()] = append(H.registry.PresenceUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(PresenceUpdateHandler(func(e *discordgo.PresenceUpdate) {
				l1chan <- true
			}))
			H.registry.PresenceUpdateListeners[l2.Owner()] = append(H.registry.PresenceUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchPresenceUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchResumed(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.Resumed
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch Resumed listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Resumed{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch Resumed listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Resumed{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch Resumed listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.Resumed{},
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
				dispatchResumed(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(ResumedHandler(func(e *discordgo.Resumed) {
				l1chan <- true
			}))
			H.registry.ResumedListeners[l.Owner()] = append(H.registry.ResumedListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchResumed(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(ResumedHandler(func(e *discordgo.Resumed) {
				l1chan <- true
			}))
			H.registry.ResumedListeners[l1.Owner()] = append(H.registry.ResumedListeners[l1.Owner()], l1)
			l2.SetActionHandler(ResumedHandler(func(e *discordgo.Resumed) {
				l1chan <- true
			}))
			H.registry.ResumedListeners[l2.Owner()] = append(H.registry.ResumedListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchResumed(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchRelationshipAdd(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.RelationshipAdd
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch RelationshipAdd listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RelationshipAdd{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch RelationshipAdd listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RelationshipAdd{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch RelationshipAdd listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RelationshipAdd{},
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
				dispatchRelationshipAdd(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(RelationshipAddHandler(func(e *discordgo.RelationshipAdd) {
				l1chan <- true
			}))
			H.registry.RelationshipAddListeners[l.Owner()] = append(H.registry.RelationshipAddListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchRelationshipAdd(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(RelationshipAddHandler(func(e *discordgo.RelationshipAdd) {
				l1chan <- true
			}))
			H.registry.RelationshipAddListeners[l1.Owner()] = append(H.registry.RelationshipAddListeners[l1.Owner()], l1)
			l2.SetActionHandler(RelationshipAddHandler(func(e *discordgo.RelationshipAdd) {
				l1chan <- true
			}))
			H.registry.RelationshipAddListeners[l2.Owner()] = append(H.registry.RelationshipAddListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchRelationshipAdd(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchRelationshipRemove(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.RelationshipRemove
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch RelationshipRemove listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RelationshipRemove{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch RelationshipRemove listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RelationshipRemove{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch RelationshipRemove listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.RelationshipRemove{},
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
				dispatchRelationshipRemove(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(RelationshipRemoveHandler(func(e *discordgo.RelationshipRemove) {
				l1chan <- true
			}))
			H.registry.RelationshipRemoveListeners[l.Owner()] = append(H.registry.RelationshipRemoveListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchRelationshipRemove(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(RelationshipRemoveHandler(func(e *discordgo.RelationshipRemove) {
				l1chan <- true
			}))
			H.registry.RelationshipRemoveListeners[l1.Owner()] = append(H.registry.RelationshipRemoveListeners[l1.Owner()], l1)
			l2.SetActionHandler(RelationshipRemoveHandler(func(e *discordgo.RelationshipRemove) {
				l1chan <- true
			}))
			H.registry.RelationshipRemoveListeners[l2.Owner()] = append(H.registry.RelationshipRemoveListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchRelationshipRemove(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchTypingStart(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.TypingStart
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch TypingStart listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.TypingStart{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch TypingStart listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.TypingStart{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch TypingStart listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.TypingStart{},
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
				dispatchTypingStart(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(TypingStartHandler(func(e *discordgo.TypingStart) {
				l1chan <- true
			}))
			H.registry.TypingStartListeners[l.Owner()] = append(H.registry.TypingStartListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchTypingStart(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(TypingStartHandler(func(e *discordgo.TypingStart) {
				l1chan <- true
			}))
			H.registry.TypingStartListeners[l1.Owner()] = append(H.registry.TypingStartListeners[l1.Owner()], l1)
			l2.SetActionHandler(TypingStartHandler(func(e *discordgo.TypingStart) {
				l1chan <- true
			}))
			H.registry.TypingStartListeners[l2.Owner()] = append(H.registry.TypingStartListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchTypingStart(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchUserUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.UserUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch UserUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch UserUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch UserUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserUpdate{},
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
				dispatchUserUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(UserUpdateHandler(func(e *discordgo.UserUpdate) {
				l1chan <- true
			}))
			H.registry.UserUpdateListeners[l.Owner()] = append(H.registry.UserUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchUserUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(UserUpdateHandler(func(e *discordgo.UserUpdate) {
				l1chan <- true
			}))
			H.registry.UserUpdateListeners[l1.Owner()] = append(H.registry.UserUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(UserUpdateHandler(func(e *discordgo.UserUpdate) {
				l1chan <- true
			}))
			H.registry.UserUpdateListeners[l2.Owner()] = append(H.registry.UserUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchUserUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchUserSettingsUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.UserSettingsUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch UserSettingsUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserSettingsUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch UserSettingsUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserSettingsUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch UserSettingsUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserSettingsUpdate{},
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
				dispatchUserSettingsUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(UserSettingsUpdateHandler(func(e *discordgo.UserSettingsUpdate) {
				l1chan <- true
			}))
			H.registry.UserSettingsUpdateListeners[l.Owner()] = append(H.registry.UserSettingsUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchUserSettingsUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(UserSettingsUpdateHandler(func(e *discordgo.UserSettingsUpdate) {
				l1chan <- true
			}))
			H.registry.UserSettingsUpdateListeners[l1.Owner()] = append(H.registry.UserSettingsUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(UserSettingsUpdateHandler(func(e *discordgo.UserSettingsUpdate) {
				l1chan <- true
			}))
			H.registry.UserSettingsUpdateListeners[l2.Owner()] = append(H.registry.UserSettingsUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchUserSettingsUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchUserGuildSettingsUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.UserGuildSettingsUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch UserGuildSettingsUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserGuildSettingsUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch UserGuildSettingsUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserGuildSettingsUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch UserGuildSettingsUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.UserGuildSettingsUpdate{},
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
				dispatchUserGuildSettingsUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(UserGuildSettingsUpdateHandler(func(e *discordgo.UserGuildSettingsUpdate) {
				l1chan <- true
			}))
			H.registry.UserGuildSettingsUpdateListeners[l.Owner()] = append(H.registry.UserGuildSettingsUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchUserGuildSettingsUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(UserGuildSettingsUpdateHandler(func(e *discordgo.UserGuildSettingsUpdate) {
				l1chan <- true
			}))
			H.registry.UserGuildSettingsUpdateListeners[l1.Owner()] = append(H.registry.UserGuildSettingsUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(UserGuildSettingsUpdateHandler(func(e *discordgo.UserGuildSettingsUpdate) {
				l1chan <- true
			}))
			H.registry.UserGuildSettingsUpdateListeners[l2.Owner()] = append(H.registry.UserGuildSettingsUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchUserGuildSettingsUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchVoiceServerUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.VoiceServerUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch VoiceServerUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.VoiceServerUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch VoiceServerUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.VoiceServerUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch VoiceServerUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.VoiceServerUpdate{},
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
				dispatchVoiceServerUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(VoiceServerUpdateHandler(func(e *discordgo.VoiceServerUpdate) {
				l1chan <- true
			}))
			H.registry.VoiceServerUpdateListeners[l.Owner()] = append(H.registry.VoiceServerUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchVoiceServerUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(VoiceServerUpdateHandler(func(e *discordgo.VoiceServerUpdate) {
				l1chan <- true
			}))
			H.registry.VoiceServerUpdateListeners[l1.Owner()] = append(H.registry.VoiceServerUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(VoiceServerUpdateHandler(func(e *discordgo.VoiceServerUpdate) {
				l1chan <- true
			}))
			H.registry.VoiceServerUpdateListeners[l2.Owner()] = append(H.registry.VoiceServerUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchVoiceServerUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}

func Test_dispatchVoiceStateUpdate(t *testing.T) {
	type args struct {
		s *discordgo.Session
		e *discordgo.VoiceStateUpdate
	}
	tests := []struct {
		name                string
		args                args
		listenersToRegister int
	}{
		{
			name: "dispatch VoiceStateUpdate listeners (empty)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.VoiceStateUpdate{},
			},
			listenersToRegister: 0,
		},
		{
			name: "dispatch VoiceStateUpdate listeners (One)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.VoiceStateUpdate{},
			},
			listenersToRegister: 1,
		},
		{
			name: "dispatch VoiceStateUpdate listeners (Two)",
			args: args{
				s: &discordgo.Session{},
				e: &discordgo.VoiceStateUpdate{},
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
				dispatchVoiceStateUpdate(tt.args.s, tt.args.e)
			})
		case 1:
			// Listener callback chan.
			l1chan := make(chan bool, 1)
			// Create and add listner to registry.
			l := &InternalListener{}
			l.SetActionHandler(VoiceStateUpdateHandler(func(e *discordgo.VoiceStateUpdate) {
				l1chan <- true
			}))
			H.registry.VoiceStateUpdateListeners[l.Owner()] = append(H.registry.VoiceStateUpdateListeners[l.Owner()], l)
			// Run test.
			t.Run(tt.name, func(t *testing.T) {
				dispatchVoiceStateUpdate(tt.args.s, tt.args.e)
			})
			select {
			case <-l1chan:
			case <-time.After(time.Second * 1):
				t.Errorf("Case %d: listener timed out.", 1)
			}
		case 2:
			// Listeners callback chan.
			l1chan := make(chan bool, 2)
			// Create Listeners.
			l1 := &InternalListener{}
			l2 := &InternalListener{}
			// Add listeners to registry.
			l1.SetActionHandler(VoiceStateUpdateHandler(func(e *discordgo.VoiceStateUpdate) {
				l1chan <- true
			}))
			H.registry.VoiceStateUpdateListeners[l1.Owner()] = append(H.registry.VoiceStateUpdateListeners[l1.Owner()], l1)
			l2.SetActionHandler(VoiceStateUpdateHandler(func(e *discordgo.VoiceStateUpdate) {
				l1chan <- true
			}))
			H.registry.VoiceStateUpdateListeners[l2.Owner()] = append(H.registry.VoiceStateUpdateListeners[l2.Owner()], l2)
			t.Run(tt.name, func(t *testing.T) {
				dispatchVoiceStateUpdate(tt.args.s, tt.args.e)
			})
			// Two listeners to callback on.
			for i := 0; i < 2; i++ {
				select {
				case <-l1chan:
				case <-time.After(time.Second * 1):
					t.Errorf("Case %d: listener (%d) timed out.", 2, i+1)
				}
			}
		}
	}
	H = nil
}


