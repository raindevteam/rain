// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package cbot

import (
	"strings"

	"gopkg.in/readline.v1"
	"gopkg.in/sorcix/irc.v1"

	"github.com/RyanPrintup/nimbus"
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

func (c CliClient) SetNick(nick string) {
	c.config.User.Nick = nick
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

	bot := rbot.NewBot("CLI", conf)
	bot.Client = cli // Oh what? I can do this?

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
