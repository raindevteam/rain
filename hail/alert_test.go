package hail

import (
	"testing"
)

func TestAlert(t *testing.T) {
	type args struct {
		facility int
		msg      string
	}
	tests := []struct {
		name    string
		args    args
		output  string
		wantErr bool
	}{
		{
			name: "simple alert print",
			args: args{
				facility: Fhail,
				msg:      "a unit test",
			},
			output:  "[HAIL] (2): a unit test\n",
			wantErr: true,
		},
	}
	Defaults()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				if err := Alert(tt.args.facility, tt.args.msg); (err != nil) !=
					tt.wantErr {
					t.Errorf("Alert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
			if output != tt.output {
				t.Errorf("output = %s, want %s", output, tt.output)
			}
		})
	}
}

func TestAlertf(t *testing.T) {
	type args struct {
		f    int
		msgf string
		v    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		output  string
		wantErr bool
	}{
		{
			name: "simple alertf print",
			args: args{
				f:    Fhail,
				msgf: "%s #%d",
				v:    []interface{}{"unit", 1},
			},
			output:  "[HAIL] (2): unit #1",
			wantErr: true,
		},
	}
	Defaults()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				err := Alertf(tt.args.f, tt.args.msgf, tt.args.v...)
				if (err != nil) != tt.wantErr {
					t.Errorf("Alertf() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
			if output != tt.output {
				t.Errorf("Output = %s, want %s", output, tt.output)
			}
		})
	}
}
