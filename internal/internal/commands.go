package internal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/internal/hail"
	"github.com/raindevteam/rain/internal/rbot"
)

type Commands struct{ bot *rbot.Bot }

func (cs *Commands) AddSource() {
	cs.bot.Handler.AddInternalCommand("source", cs.Source)
}

func (cs *Commands) Source(args string, e *discordgo.MessageCreate) error {
	_, err := cs.bot.Session.ChannelMessageSend(e.Message.ChannelID, "github.com/raindevteam/rain")
	if err != nil {
		hail.Critf(hail.Finternal, "Error sending message to discord channel: %s", err)
		return err
	}
	return nil
}

func (cs *Commands) AddHelp() {
	cs.bot.Handler.AddInternalCommand("help", cs.Help)
}

func (cs *Commands) Help(args string, e *discordgo.MessageCreate) error {
	_, err := cs.bot.Session.ChannelMessageSend(e.Message.ChannelID, "The `help` command is still undergoing work")
	if err != nil {
		hail.Critf(hail.Finternal, "Error sending message to discord channel: %s", err)
		return err
	}
	return nil
}
