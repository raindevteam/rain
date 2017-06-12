package internal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/hail"
	"github.com/raindevteam/rain/handler"
	"github.com/raindevteam/rain/rbot"
)

type Listeners struct{ bot *rbot.Bot }

func (l *Listeners) AddMessageCreate() {
	_ = handler.AddInternalListener(l.messageCreate)
}

func (l *Listeners) messageCreate(e *discordgo.MessageCreate) {
	hail.Infof(hail.Fcore, "MessageCreate: %s", e.Message.Content)
}
