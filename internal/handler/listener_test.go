// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package handler

import (
	"reflect"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/internal/droplet"
)

var confirmRun = make(chan bool, 1)

func runMessageCreateEvent(e *discordgo.MessageCreate) {
	confirmRun <- true
}

func TestInternalListener_Run(t *testing.T) {
	type fields struct {
		Event   interface{}
		enabled bool
		action  Action
	}
	tests := []struct {
		name     string
		fields   fields
		wantsRun bool
	}{
		{
			name: "MessageCreate event run",
			fields: fields{
				Event:   &discordgo.MessageCreate{},
				enabled: false,
				action:  MessageCreateHandler(runMessageCreateEvent),
			},
			wantsRun: true,
		},
		// Keep in mind for the next two unit tests that a nil action and event
		// should prevent a run. Therefore it is considered a fail if the
		// listener does run.
		{
			name: "nil action run",
			fields: fields{
				Event:   &discordgo.MessageCreate{},
				enabled: false,
				action:  nil,
			},
			wantsRun: false,
		},
		{
			name: "nil event run",
			fields: fields{
				Event:   nil,
				enabled: false,
				action:  MessageCreateHandler(runMessageCreateEvent),
			},
			wantsRun: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &InternalListener{
				Event:   tt.fields.Event,
				enabled: tt.fields.enabled,
				action:  tt.fields.action,
			}
			il.Run()
			select {
			case <-confirmRun:
				if !tt.wantsRun {
					t.Errorf("listener ran, wantsRun: %t", tt.wantsRun)
				}
			case <-time.After(time.Second * 1):
				if tt.wantsRun {
					t.Errorf("listener timed out.")
				}
			}
		})
	}
}

func TestDropletListener_Run(t *testing.T) {
	type fields struct {
		Event  interface{}
		drop   *droplet.Droplet
		action Action
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "nil run",
			fields: fields{
				Event:  nil,
				drop:   nil,
				action: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := DropletListener{
				Event:  tt.fields.Event,
				drop:   tt.fields.drop,
				action: tt.fields.action,
			}
			l.Run()
		})
	}
}

func TestInternalListener_SetActionHandler(t *testing.T) {
	type fields struct {
		Event   interface{}
		enabled bool
		action  Action
	}
	type args struct {
		a Action
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set valid action handler",
			fields: fields{
				Event:   nil,
				enabled: false,
				action:  nil,
			},
			args: args{
				a: MessageCreateHandler(runMessageCreateEvent),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &InternalListener{
				Event:   tt.fields.Event,
				enabled: tt.fields.enabled,
				action:  tt.fields.action,
			}
			il.SetActionHandler(tt.args.a)
			v1 := reflect.TypeOf(il.action)
			v2 := reflect.TypeOf(tt.args.a)
			if v1 != v2 {
				t.Errorf("il.action = %s, want = %s", v1, v2)
			}
		})
	}
}

func TestDropletListener_SetActionHandler(t *testing.T) {
	type fields struct {
		Event  interface{}
		drop   *droplet.Droplet
		action Action
	}
	type args struct {
		a Action
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set valid action handler",
			fields: fields{
				Event:  nil,
				drop:   nil,
				action: nil,
			},
			args: args{
				a: MessageCreateHandler(runMessageCreateEvent),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := DropletListener{
				Event:  tt.fields.Event,
				drop:   tt.fields.drop,
				action: tt.fields.action,
			}
			l.SetActionHandler(tt.args.a)
			v1 := reflect.TypeOf(l.action)
			v2 := reflect.TypeOf(tt.args.a)
			if v1 != v2 {
				t.Errorf("il.action = %s, want = %s", v1, v2)
			}
		})
	}
}

func TestInternalListener_SetEvent(t *testing.T) {
	type fields struct {
		Event   interface{}
		enabled bool
		action  Action
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set valid event",
			fields: fields{
				Event:   nil,
				enabled: false,
				action:  nil,
			},
			args: args{
				v: &discordgo.MessageCreate{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &InternalListener{
				Event:   tt.fields.Event,
				enabled: tt.fields.enabled,
				action:  tt.fields.action,
			}
			il.SetEvent(tt.args.v)
			if !reflect.DeepEqual(il.Event, tt.args.v) {
				t.Errorf("il.action = %s, want = %s", il.Event, tt.args.v)
			}
		})
	}
}

func TestDropletListener_SetEvent(t *testing.T) {
	type fields struct {
		Event  interface{}
		drop   *droplet.Droplet
		action Action
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set valid event",
			fields: fields{
				Event:  nil,
				drop:   nil,
				action: nil,
			},
			args: args{
				v: &discordgo.MessageCreate{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DropletListener{
				Event:  tt.fields.Event,
				drop:   tt.fields.drop,
				action: tt.fields.action,
			}
			l.SetEvent(tt.args.v)
			if !reflect.DeepEqual(l.Event, tt.args.v) {
				t.Errorf("il.action = %s, want = %s", l.Event, tt.args.v)
			}
		})
	}
}

func TestInternalListener_Owner(t *testing.T) {
	type fields struct {
		Event   interface{}
		enabled bool
		action  Action
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "return valid owner",
			fields: fields{
				Event:   nil,
				enabled: false,
				action:  nil,
			},
			want: "__INTERNAL__",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &InternalListener{
				Event:   tt.fields.Event,
				enabled: tt.fields.enabled,
				action:  tt.fields.action,
			}
			if got := il.Owner(); got != tt.want {
				t.Errorf("InternalListener.Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropletListener_Owner(t *testing.T) {
	type fields struct {
		Event  interface{}
		drop   *droplet.Droplet
		action Action
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "return no owner",
			fields: fields{
				Event:  nil,
				drop:   nil,
				action: nil,
			},
			want: "__NO_DROPLET__",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DropletListener{
				Event:  tt.fields.Event,
				drop:   tt.fields.drop,
				action: tt.fields.action,
			}
			if got := l.Owner(); got != tt.want {
				t.Errorf("DropletListener.Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}
