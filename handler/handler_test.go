// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
//
// Contact Rodolfo at rcvallada@gmail.com for any inquires of this file.

package handler

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/droplet"
)

func tAction(e *discordgo.Ready) {
	fmt.Printf("I'm an idiot\n")
}

func TestInternalListenersRun(t *testing.T) {
	var v interface{}
	v = tAction
	il := &InternalListener{
		drop: &droplet.Droplet{},
		act:  v,
	}
	e := &discordgo.Ready{}
	il.Run(e)
}
