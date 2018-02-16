package parser

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		want      *parser.CommandData
		wantError bool
	}{
		{
			name: "Test internal Command",
			args: args{
				input: "!!src",
			},
			want: &parser.CommandData{
				Owner:   "__INTERNAL__",
				Command: "src",
			},
			wantError: false,
		},
		{
			name: "Test droplet (non-specific) Command",
			args: args{
				input: "!roll",
			},
			want: &parser.CommandData{
				Owner:   "",
				Command: "roll",
			},
			wantError: false,
		},
		{
			name: "Test droplet (specific) Command",
			args: args{
				input: "!game:roll",
			},
			want: &parser.CommandData{
				Owner:   "game",
				Command: "roll",
			},
			wantError: false,
		},
		{
			name: "Test invalid command",
			args: args{
				input: "game:roll",
			},
			want:      nil,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := ParseCommand(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
				if err != nil && !tt.wantError {
					t.Errorf("ParseCommand() wantError = %v, got %v", tt.wantError, err)
				}
			}
		})
	}
}
