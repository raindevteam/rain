package setup

import (
	"strings"

	"gopkg.in/sorcix/irc.v1"

	"github.com/raindevteam/rain/rbot"
)

type moduleManager struct{ bot *rbot.Bot }

func (c *moduleManager) info(msg *irc.Message, Params []string) {
	if len(Params) == 1 {
		c.bot.Say(msg.Params[0], "Error: No module specified")
		return
	}

	info, err := c.bot.ModuleInfo(Params[1])
	if err != nil {
		c.bot.Say(msg.Params[0], "Error: "+err.Error())
		return
	}

	c.bot.Say(msg.Params[0], info)
}

func (c *moduleManager) load(msg *irc.Message, Params []string) {
	if len(Params) == 1 {
		c.bot.Say(msg.Params[0], "Error: No module specified")
		return
	}

	err := c.bot.ModuleLoad(Params[1])
	if err != nil {
		c.bot.Say(msg.Params[0], "Error: "+err.Error())
		return
	}

	c.bot.Say(msg.Params[0], "Loaded module: "+Params[1])
}

func (c *moduleManager) m(msg *irc.Message, Params []string) {
	if len(Params) == 0 {
		c.bot.Say(msg.Params[0], "Error: No command specified")
		return
	}

	switch Params[0] {
	case "reload":
		if len(Params) == 1 {
			c.bot.Say(msg.Params[0], "Error: No module specified")
			return
		}

		if err := c.bot.ModuleReload(Params[1]); err != nil {
			c.bot.Say(msg.Params[0], "Error: "+err.Error())
			return
		}

		c.bot.Say(msg.Params[0], "Reloaded module: "+Params[1])

	case "kill":
		if len(Params) == 1 {
			c.bot.Say(msg.Params[0], "Error: No module specified")
			return
		}

	case "load":
		c.load(msg, Params)
	case "info":
		c.info(msg, Params)
	}
}

func Default(b *rbot.Bot) {
	mm := moduleManager{b}

	b.Handler.AddInternalCommand("m", &rbot.Command{
		Help: "The RainBot Module Manager helps manage modules.",
		Fun:  mm.m,
		PM:   true,
		CM:   true,
	})

	// Commands listener
	b.AddListener(irc.PRIVMSG, func(msg *irc.Message) {
		if b.Parser.IsCommand(msg.Trailing) {
			command, Params := b.Parser.ParseCommand(msg.Trailing)
			b.Handler.Invoke(msg, rbot.CommandName(command), Params)
		}
	})

	// Add/Update channel when topic is received
	b.AddListener(irc.RPL_TOPIC, func(msg *irc.Message) {
		b.Mu.Lock()

		name := msg.Params[1]
		topic := msg.Trailing

		b.Channels[strings.ToLower(name)].Topic = topic

		b.Mu.Unlock()
	})

	// Update users for channel
	b.AddListener(irc.RPL_NAMREPLY, func(msg *irc.Message) {
		b.Mu.Lock()

		channel := b.Channels[strings.ToLower(msg.Params[2])]
		users := strings.Split(strings.Trim(msg.Trailing, " "), " ")

		for _, user := range users {
			var name, rank string

			if strings.ContainsAny(string(user[0]), "+ & ~ & @ & &") {
				name, rank = user[1:], string(user[0])
			} else {
				name, rank = user, ""
			}

			channel.Users[name] = rank
		}

		b.Mu.Unlock()
	})

	// Update on user Join
	b.AddListener(irc.JOIN, func(msg *irc.Message) {
		b.Mu.Lock()
		defer b.Mu.Unlock()

		who := msg.Prefix.Name
		where := msg.Trailing

		if who == b.GetNick() {
			channel := rbot.NewChannel(where)
			b.Channels[strings.ToLower(where)] = channel
			return
		}

		channel := b.Channels[strings.ToLower(where)]
		channel.Users[who] = ""
	})

	// Update on user Kick
	b.AddListener(irc.KICK, func(msg *irc.Message) {
		b.Mu.Lock()

		who := msg.Params[1]
		where := msg.Params[0]

		b.RemoveUser(who, where)

		b.Mu.Unlock()
	})

	// Update on user Kill
	b.AddListener(irc.KILL, func(msg *irc.Message) {
		b.Mu.Lock()

		// Implement getInfo(msg) function?
		who := msg.Prefix.Name
		where := msg.Params[0]

		b.RemoveUser(who, where)

		b.Mu.Unlock()
	})

	// Update on user part
	b.AddListener(irc.PART, func(msg *irc.Message) {
		b.Mu.Lock()

		who := msg.Prefix.Name
		where := msg.Params[0]

		b.RemoveUser(who, where)

		b.Mu.Unlock()
	})

	// Update on user quit
	b.AddListener(irc.QUIT, func(msg *irc.Message) {
		b.Mu.Lock()

		who := msg.Prefix.Name

		if who == b.GetNick() {
			// We quit
		}

		for _, channel := range b.Channels {
			delete(channel.Users, who)
		}

		b.Mu.Unlock()
	})

	// Update on nick change
	b.AddListener(irc.NICK, func(msg *irc.Message) {
		b.Mu.Lock()

		b.Mu.Unlock()
	})
}
