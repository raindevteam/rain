package Thelpers

import (
	"sync"

	"github.com/raindevteam/rain/bot"
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
