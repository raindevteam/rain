package clibot

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Bowery/prompt"
	"github.com/RyanPrintup/nimbus"
	"github.com/wolfchase/rainbot/bot"
)

type CLIMode string

const (
	IrcMode CLIMode = "irc"
	MsgMode CLIMode = "msg"
)

type CLIClient struct {
	Mode      CLIMode
	CLIPrefix string
	listeners map[string][]nimbus.Listener
	config    *rainbot.Config
	quit      chan error
}

func (c CLIClient) GetNick() string {
	return c.config.Nick
}

func (c CLIClient) GetChannels() []string {
	return c.config.Channel
}

func (c CLIClient) Send(raw ...string) {
	switch c.Mode {
	case IrcMode:
		fmt.Println("RainBot >> " + strings.Join(raw[:len(raw)-2], " "))
	case MsgMode:
		lasti := len(raw) - 1

		var rawjoined string

		if raw[lasti][0] == ':' {
			rawjoined = strings.Join(raw[:lasti], " ") + raw[lasti]
		} else {
			rawjoined = strings.Join(raw, " ")
		}

		msg, err := nimbus.ParseMessage(rawjoined)

		if err != nil {
			fmt.Println(c.GetNick() + "> " + err.Error())
		}

		if msg.Command == nimbus.PRIVMSG {
			fmt.Println(c.GetNick() + "> " + msg.Trailing)
		} else {
			fmt.Println(c.GetNick() + "> " + "(" + msg.Command + ") ")
		}
	}
}

func (c CLIClient) Say(channel string, text string) {
	c.Send(nimbus.PRIVMSG, channel+" :", text)
}

func NewCLIBot(rainConfig *rainbot.Config) *rainbot.Bot {
	cli := &CLIClient{
		MsgMode,
		rainConfig.CmdPrefix + ":",
		make(map[string][]nimbus.Listener),
		rainConfig,
		make(chan error),
	}

	bot := &rainbot.Bot{
		/* Client      */ cli,
		/* Version     */ "Alpha 0.4.0 (Monterey Jack)",
		/* Modules     */ make(map[string]*rainbot.Module),
		/* Channels    */ make(map[string]*rainbot.Channel),
		/* Parser      */ rainbot.NewParser(rainConfig.CmdPrefix),
		/* Handler     */ rainbot.NewHandler(),
		/* Mutex       */ sync.Mutex{},
	}

	return bot
}

func (c CLIClient) Connect(callback func(error)) error {
	callback(nil)
	return nil
}

func (c CLIClient) Quit() chan error {
	return c.quit
}

func (c CLIClient) Listen() {
	for {
		line, err := prompt.Basic("> ", true)

		raw := ":RainBot " + nimbus.PRIVMSG + " #cli :" + line
		msg, err := nimbus.ParseMessage(raw)

		if err != nil {
			fmt.Println(err)
		}

		c.Emit(nimbus.PRIVMSG, msg)
	}
}