package rbot

import (
	"errors"
	"net/rpc"
	"strings"

	"gopkg.in/sorcix/irc.v1"

	"github.com/raindevteam/rain/rlog"
)

// BotAPI is the exposed api served via the bot's master consumer connection
type BotAPI struct {
	bot *Bot
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

type JoinRequest struct {
	Channel  string
	Password string
}

// Send transmits a message over irc as a PRIVMSG
func (b BotAPI) Send(raw string, result *string) error {
	b.bot.Send(irc.PRIVMSG, raw)
	return nil
}

// RegisterCommand registers a command from a module with the handler. The command request
// holds the command's name and the module it belongs to (used to signal the module to fire the
// command).
func (b BotAPI) RegisterCommand(cr CommandRequest, result *string) error {
	b.bot.Handler.AddCommand(CommandName(cr.CommandName), ModuleName(cr.ModuleName))
	rlog.Debug("Bot", "Added: "+cr.CommandName+" for module: "+cr.ModuleName)
	return nil
}

// RegisterTrigger will register a trigger from a module with the bot handler. If this the first
// trigger for its corresponding event, the bot will add a new listener that handles trigger firing
// for this event.
func (b BotAPI) RegisterTrigger(tr TriggerRequest, result *string) error {
	listeners := b.bot.GetListeners(string(tr.Event))
	if len(listeners) == 0 {
		b.bot.AddListener(string(tr.Event), func(msg *irc.Message) {
			b.bot.Handler.Fire(msg, tr.Event)
		})
	}
	b.bot.Handler.AddTrigger(tr.Name, tr.Event)
	return nil
}

func (b BotAPI) JoinChannel(jr JoinRequest, result *string) error {
	if jr.Password != "" {
		b.bot.Send(irc.JOIN, jr.Channel, jr.Password)
	} else {
		b.bot.Send(irc.JOIN, jr.Channel)
	}

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
	rlog.Debug("Bot", "Starting registration for "+t.ModuleName+" [Module Client]")
	client, err := RpcCodecClientWithPort(t.Port)
	rlog.Debug("Bot", "Client created")
	if err != nil {
		rlog.Error("Bot", "Could not establish an RPC client: "+err.Error())
	}

	module := rpc.NewClientWithCodec(client)
	if module == nil {
		rlog.Warn("Bot", "Could not register:"+t.ModuleName)
		return errors.New("Failed to regsiter module")
	}

	b.bot.Handler.AddModule(ModuleName(strings.ToLower(t.ModuleName)), module)
	b.bot.Modules[strings.ToLower(t.ModuleName)].Registered = true
	rlog.Debug("Bot", "Registered "+t.ModuleName+" on port "+t.Port)
	return nil
}
