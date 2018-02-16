package internal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/internal/hail"
	"github.com/raindevteam/rain/internal/parser"
	"github.com/raindevteam/rain/internal/rbot"
)

// Listeners contains all the internal listners for the bot. It also
// has a reference to the bot.
type Listeners struct{ bot *rbot.Bot }

// AddCommandParser adds a messageCreate listener for commands.
func (ls *Listeners) AddCommandParser() {
	_ = ls.bot.Handler.AddInternalListener(ls.CommandParser)
}

// CommandParser parses messages for commands.
func (ls *Listeners) CommandParser(e *discordgo.MessageCreate) {
	commandData, err := parser.ParseCommand(e.Message.Content)
	if err == nil {
		hail.Infof(hail.Finternal, "CommandParser: Owner: %s, Command: %s\n",
			commandData.Owner, commandData.Command)
		ls.bot.Handler.InvokeCommand(commandData, e)
	}
}

// AddMessageCreate adds the messageCreate internal listener.
func (ls *Listeners) AddMessageCreate() {
	_ = ls.bot.Handler.AddInternalListener(ls.MessageCreate)
}

// MessageCreate logs messages sent in all channels the bot has access to.
func (ls *Listeners) MessageCreate(e *discordgo.MessageCreate) {
	hail.Infof(hail.Finternal, "MessageCreate: %s: %s\n",
		e.Author.Username, e.Message.Content)

	for _, duser := range e.Mentions {
		me, _ := ls.bot.Session.User("@me")
		if duser.ID == me.ID {
			ls.bot.Session.ChannelMessageSend(e.ChannelID,
				"If you're new to Rain, use these commmands to get going!\n"+
					"`!!help` ... Get a list of commands\n"+
					"`!!plugins` ... Get a list of plugins\n"+
					"`!!dhelp <plugin>` ... Get a list of commands for a plugin\n"+
					"`!!dinfo <plugin>` ... Show general information for a plugin\n")
		}
	}
}
