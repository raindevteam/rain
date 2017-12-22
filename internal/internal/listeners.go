package internal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/internal/hail"
	"github.com/raindevteam/rain/internal/handler"
	"github.com/raindevteam/rain/internal/rbot"
)

// Listeners contains all the internal listners for the bot. It also
// has a reference to the bot.
type Listeners struct{ bot *rbot.Bot }

// AddMessageCreate adds the messageCreate internal listener.
func (l *Listeners) AddMessageCreate() {
	_ = handler.AddInternalListener(l.MessageCreate)
}

// MessageCreate logs messages sent in all channels the bot has access to.
func (l *Listeners) MessageCreate(e *discordgo.MessageCreate) {
	hail.Infof(hail.Finternal, "MessageCreate: %s: %s\n",
		e.Author.Username, e.Message.Content)
}
