// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package Thelpers

import (
	"sync"

	"github.com/raindevteam/rain/rbot"
)

func NewMockBot(rcon *rbot.Config) *rbot.Bot {
	bot := &rbot.Bot{
		/* Client      */ NewMockClient(rcon, []string{
			"JOIN #snowybottest",
			"JOIN #rainbot",
		}),
		/* Version     */ "Alpha 0.4.0 (Monterey Jack)",
		/* Modules     */ make(map[string]*rbot.Module),
		/* Channels    */ make(map[string]*rbot.Channel),
		/* Parser      */ rbot.NewParser(rcon.CmdPrefix),
		/* Handler     */ rbot.NewHandler(),
		/* Mutex       */ sync.Mutex{},
	}
	return bot
}
