package setup

import (
	"strings"

	"github.com/RyanPrintup/nimbus"
	"github.com/raindevteam/rain/rbot"
)

type commands struct{ bot *rbot.Bot }

func (c *commands) ModuleManager(msg *nimbus.Message, args []string) {
	if len(args) == 0 {
		c.bot.Say(msg.Args[0], "Error: No command specified")
		return
	}

	switch args[0] {
	case "reload":
		if len(args) == 1 {
			c.bot.Say(msg.Args[0], "Error: No module specified")
			return
		}

		if err := c.bot.ModuleReload(args[1]); err != nil {
			c.bot.Say(msg.Args[0], "Error: Could not reload module, view output in console")
			return
		}

		c.bot.Say(msg.Args[0], "Reloaded module: "+args[1])
	case "kill":
		if len(args) == 1 {
			c.bot.Say(msg.Args[0], "Error: No module specified")
			return
		}
	}
}

func Default(b *rbot.Bot) {
	c := commands{b}

	b.Handler.AddInternalCommand("m", &rbot.Command{
		Help: "The RainBot Module Manager helps manage modules.",
		Fun:  c.ModuleManager,
		PM:   true,
		CM:   true,
	})

	// Commands listener
	b.AddListener(nimbus.PRIVMSG, func(msg *nimbus.Message) {
		if b.Parser.IsCommand(msg.Trailing) {
			command, args := b.Parser.ParseCommand(msg.Trailing)
			b.Handler.Invoke(msg, rbot.CommandName(command), args)
		}
	})

	// Add/Update channel when topic is received
	b.AddListener(nimbus.RPL_TOPIC, func(msg *nimbus.Message) {
		b.Mu.Lock()

		name := msg.Args[1]
		topic := msg.Trailing

		b.Channels[strings.ToLower(name)].Topic = topic

		b.Mu.Unlock()
	})

	// Update users for channel
	b.AddListener(nimbus.RPL_NAMEREPLY, func(msg *nimbus.Message) {
		b.Mu.Lock()

		channel := b.Channels[strings.ToLower(msg.Args[2])]
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
	b.AddListener(nimbus.JOIN, func(msg *nimbus.Message) {
		b.Mu.Lock()
		defer b.Mu.Unlock()

		who, _ := b.Parser.ParsePrefix(msg.Prefix)
		where := msg.Args[0]

		if who == b.GetNick() {
			channel := rbot.NewChannel(where)
			b.Channels[strings.ToLower(where)] = channel
			return
		}

		channel := b.Channels[strings.ToLower(where)]
		channel.Users[who] = ""
	})

	// Update on user Kick
	b.AddListener(nimbus.KICK, func(msg *nimbus.Message) {
		b.Mu.Lock()

		who := msg.Args[2]
		where := msg.Args[0]

		b.RemoveUser(who, where)

		b.Mu.Unlock()
	})

	// Update on user Kill
	b.AddListener(nimbus.KILL, func(msg *nimbus.Message) {
		b.Mu.Lock()

		// Implement getInfo(msg) function?
		who, _ := b.Parser.ParsePrefix(msg.Prefix)
		where := msg.Args[0]

		b.RemoveUser(who, where)

		b.Mu.Unlock()
	})

	// Update on user part
	b.AddListener(nimbus.PART, func(msg *nimbus.Message) {
		b.Mu.Lock()

		who, _ := b.Parser.ParsePrefix(msg.Prefix)
		where := msg.Args[0]

		b.RemoveUser(who, where)

		b.Mu.Unlock()
	})

	// Update on user quit
	b.AddListener(nimbus.QUIT, func(msg *nimbus.Message) {
		b.Mu.Lock()

		who, _ := b.Parser.ParsePrefix(msg.Prefix)

		if who == b.GetNick() {
			// We quit
		}

		for _, channel := range b.Channels {
			delete(channel.Users, who)
		}

		b.Mu.Unlock()
	})

	// Update on nick change
	b.AddListener(nimbus.NICK, func(msg *nimbus.Message) {
		b.Mu.Lock()

		b.Mu.Unlock()
	})
}
