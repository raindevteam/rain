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

	"github.com/bwmarrin/discordgo"
)

func PrebuiltRegistry() *Registry {
	r := &Registry{}
	r.Initialize()
	return r
}

func ilistenerMessageCreate(e *discordgo.MessageCreate) {
	return
}

func compareIL(a, b *InternalListener) bool {
	if &a == &b {
		return true
	}
	if a.Event != b.Event {
		return false
	}
	if a.enabled != b.enabled {
		return false
	}
	if reflect.TypeOf(a.action) != reflect.TypeOf(b.action) {
		return false
	}
	return true
}

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name    string
		want    *Handler
		wantErr bool
	}{
		{
			name: "First call to CreateHandler()",
			want: &Handler{
				Status:   true,
				registry: PrebuiltRegistry(),
			},
			wantErr: false,
		},
		{
			name: "Second call to CreateHandler()",
			want: &Handler{
				Status:   true,
				registry: PrebuiltRegistry(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateHandler(); (err != nil) != tt.wantErr {
				t.Errorf("CreateHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := H; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHandler() H = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddInternalListener(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want *InternalListener
	}{
		{
			name: "valid listener",
			args: args{
				v: ilistenerMessageCreate,
			},
			want: &InternalListener{
				Event:   nil,
				enabled: false,
				action:  MessageCreateHandler(ilistenerMessageCreate),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			H = &Handler{
				Status:   true,
				registry: &Registry{},
			}
			H.registry.Initialize()
			got := AddInternalListener(tt.args.v).(*InternalListener)
			if !compareIL(got, tt.want) {
				t.Errorf("AddInternalListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
