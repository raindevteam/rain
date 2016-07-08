package rbot

/*

import (
	"errors"
	"net/rpc"
	"strings"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rlog"
	"github.com/wolfchase/rainbot/bot"
)

// BotAPI is the exposed api served via the bot's master consumer connection
type BotAPI struct {
	bot *rbot.Bot
}

// A Ticket is used to connect to a provider connection via rpc.
type Ticket struct {
	Port       string
	ModuleName string
}

// A CommandRequest is used to register commands via the handler.
type CommandRequest struct {
	CommandName string
	ModuleName  string
}

// A TriggerRequest is used to register triggers via the handler.
type TriggerRequest struct {
	Name  ModuleName
	Event Event
}

// Send transmits a message over irc as a PRIVMSG
func (b BotAPI) Send(raw string, result *string) error {

	b.bot.Send(nimbus.PRIVMSG, raw)
	return nil
}

// RegisterCommand registers a command from a module with the handler. The command request
// holds the command's name and the module it belongs to (used to signal the module to fire the
// command).
func (b BotAPI) RegisterCommand(cr CommandRequest, result *string) error {
	b.bot.Handler.AddCommand(cr.Name, cr.Module)
	rlog.Debug("Bot", "Added: "+string(cr.Name)+" for module: "+string(cr.Module))
	return nil
}

// RegisterTrigger will register a trigger from a module with the bot handler. If this the first
// trigger for its corresponding event, the bot will add a new listener that handles trigger firing
// for this event.
func (b BotAPI) RegisterTrigger(tr TriggerRequest, result *string) error {
	listeners := b.bot.GetListeners(string(tr.Event))
	if len(listeners) == 0 {
		b.bot.AddListener(string(tr.Event), func(msg *nimbus.Message) {
			b.bot.Handler.Fire(msg, tr.Event)
		})
	}
	b.bot.Handler.AddTrigger(tr.Name, tr.Event)
	return nil
}

// GetVersion will return the bot's current version.
func (b BotAPI) GetVersion(mName string, result *string) error {
	*result = b.bot.Version
	return nil
}

// GetConnectedUsers will return a user map (where every user has an IRC rank as a value).
func (b BotAPI) GetConnectedUsers(channel string, result *map[string]string) error {
	*result = b.bot.Channels[strings.ToLower(channel)].Users
	return nil
}

// GetTopic returns the channel's topic.
func (b BotAPI) GetTopic(channel string, result *string) error {
	if _, ok := b.bot.Channels[strings.ToLower(channel)]; !ok {
		*result = ""
		return errors.New("Channel doesn't exist")
	}

	*result = b.bot.Channels[strings.ToLower(channel)].Topic
	return nil
}

// Register registers a module with the bot. With the given port number in the Ticket, the bot
// creates a new rpc provider client connection to the module. The module is kept in the handler
// for event dispatching and module management.
func (b BotAPI) Register(t Ticket, result *string) error {
	module := rpc.NewClientWithCodec(RpcCodecClientWithPort(t.Port))
	if module == nil {
		rlog.Warn("Bot", "Could not register:"+string(t.Name))
		return errors.New("Failed to regsiter module")
	}
	b.bot.Handler.AddModule(ModuleName(strings.ToLower(string(t.Name))), module)
	rlog.Debug("Bot", "Registered "+string(t.Name)+" on port "+t.Port)
	return nil
}
*/
