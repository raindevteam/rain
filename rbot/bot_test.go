// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package bot

import (
	"os"
	"testing"

	"github.com/raindevteam/rain/hail"
)

func SetupTesting(t *testing.T) {
	os.Setenv("RAIN_TESTING", "TEST")
	hail.Defaults()
	hail.SetFlags(hail.Ldebug)
}

func TestNewBot(t *testing.T) {
	SetupTesting(t)

	conf, err := NewConfigFromString(configString)
	if err != nil {
		t.Fatalf("Error while creating new config: %s\n", err)
	}

	b, err := NewBot(conf, "token")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	ds := b.Session.(*DST)
	if ds.id != 0 {
		t.Fatalf("DST set incorrect ID, got %d, expecting 0\n", ds.id)
	}

	if b.config.Name != "RainBot" {
		t.Errorf("NewBot failed to set the correct name! Got %s instead",
			b.config.Name)
	}
}
