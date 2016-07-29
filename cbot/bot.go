package cbot

import (
	"strings"
	"sync"

	"gopkg.in/sorcix/irc.v1"

	"github.com/RyanPrintup/nimbus"
	"github.com/chzyer/readline"
	"github.com/raindevteam/rain/parser"
	"github.com/raindevteam/rain/rbot"
	"github.com/raindevteam/rain/rlog"
)

var rl *readline.Instance

type CLIMode string

const (
	IrcMode CLIMode = "irc"
	MsgMode CLIMode = "msg"
)

type CliClient struct {
	Mode      CLIMode
	CLIPrefix string
	listeners map[string][]nimbus.Listener
	config    *rbot.Config
	quit      chan error
}

func (c CliClient) GetNick() string {
	return c.config.User.Nick
}

func (c CliClient) GetChannels() []string {
	return c.config.Server.Channels
}

func (c CliClient) Send(raw ...string) {
	rl.Clean()
	switch c.Mode {
	case IrcMode:
		rlog.Println("RainBot >> " + strings.Join(raw[:len(raw)-2], " "))
	case MsgMode:
		rawjoined := strings.Join(raw, " ")
		msg := irc.ParseMessage(rawjoined)

		if msg.Command == irc.PRIVMSG {
			rlog.Println(" " + c.GetNick() + " » " + msg.Trailing)
		} else {
			rlog.Println(" " + c.GetNick() + " » " + "(" + msg.Command + ") ")
		}
		rl.Refresh()
	}
}

func (c CliClient) Say(channel string, text string) {
	c.Send(irc.PRIVMSG, channel, ":"+text)
}

func NewCLIBot(conf *rbot.Config) *rbot.Bot {
	rlog.SetFlags(rlog.Linfo | rlog.Lwarn | rlog.Lerror | rlog.Ldebug)
	rlog.SetLogFlags(0)

	cli := &CliClient{
		MsgMode,
		conf.Command.Prefix + ":",
		make(map[string][]nimbus.Listener),
		conf,
		make(chan error),
	}

	bot := &rbot.Bot{
		/* Client      */ cli,
		/* Version     */ "CLI",
		/* Modules     */ make(map[string]*rbot.Module),
		/* Channels    */ make(map[string]*rbot.Channel),
		/* Parser      */ parser.NewParser(conf.Command.Prefix),
		/* Handler     */ rbot.NewHandler(),
		/* Config      */ conf,
		/* Mutex       */ sync.Mutex{},
	}

	return bot
}

func (c CliClient) Connect() error {
	return nil
}

func (c CliClient) Quit() chan error {
	return c.quit
}

func (c CliClient) Listen() {
	var err error

	rl, err = readline.New(" » ")

	if err != nil {
		panic(err)
	}

	defer rl.Close()

	rlog.SetOutput(rl.Stdout())

	for {
		line, err := rl.Readline()

		if len(line) == 0 {
			continue
		}

		if err != nil {
			break
		}

		raw := ":RainBot " + irc.PRIVMSG + " #cli :" + line
		msg := irc.ParseMessage(raw)

		if err != nil {
			rlog.Error("cli", err.Error())
		}

		c.Emit(irc.PRIVMSG, msg)
	}
}
