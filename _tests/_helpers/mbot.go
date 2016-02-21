package RBThelpers

import (
	"sync"

	"github.com/wolfchase/rainbot/bot"
)

func NewMockBot(rcon *rainbot.Config) *rainbot.Bot {
	bot := &rainbot.Bot{
		/* Client      */ NewMockClient(rcon, []string{
			"JOIN #snowybottest",
			"JOIN #rainbot",
		}),
		/* Version     */ "Alpha 0.4.0 (Monterey Jack)",
		/* Modules     */ make(map[string]*rainbot.Module),
		/* Channels    */ make(map[string]*rainbot.Channel),
		/* Parser      */ rainbot.NewParser(rcon.CmdPrefix),
		/* Handler     */ rainbot.NewHandler(),
		/* Mutex       */ sync.Mutex{},
	}
	return bot
}
