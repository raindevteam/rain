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

package setup

import (
	"errors"
	"strconv"
	"strings"

	"github.com/raindevteam/rain/rbot"
	"github.com/raindevteam/rain/rlog"
	"gopkg.in/sorcix/irc.v1"
)

// The Listeners struct holds all predefined listeners available
type Listeners struct{ bot *rbot.Bot }

///////////////////////////                IRC Listeners                 ///////////////////////////

// Privmsg will parse messages for commands and invoke them via the handler appropriately. It also
// has an inbound rate limiter to prevent command flooding.
func (l *Listeners) Privmsg(msg *irc.Message) {
	if l.bot.Parser.IsCommand(msg.Trailing) {
		if <-l.bot.InboundLimiter() {
			command, Params := l.bot.Parser.ParseCommand(msg.Trailing)
			l.bot.Handler.Invoke(msg, rbot.CommandName(command), Params)
		}
	}
}

// RplTopic will set the topic for a created channel
func (l *Listeners) RplTopic(msg *irc.Message) {
	l.bot.Mu.Lock()

	name := msg.Params[1]
	topic := msg.Trailing

	channel, ok := l.bot.Channels[strings.ToLower(name)]

	if !ok {
		rlog.Warn("Bot", "Got RPL_TOPIC for a channel that hasn't been initialized yet")
	} else {
		channel.Topic = topic
	}

	l.bot.Mu.Unlock()
}

// RplNamreply will insert all users into the corresponding channel
func (l *Listeners) RplNamreply(msg *irc.Message) {
	l.bot.Mu.Lock()
	defer l.bot.Mu.Unlock()

	channel, ok := l.bot.Channels[strings.ToLower(msg.Params[2])]
	users := strings.Split(strings.Trim(msg.Trailing, " "), " ")

	if !ok {
		rlog.Warn("Bot", "Got RPL_NAMREPLY for a channel that hasn't been initialized yet")
		return
	}

	for _, user := range users {
		var name, rank string

		if strings.ContainsAny(string(user[0]), "+ & ~ & @ & &") {
			name, rank = user[1:], string(user[0])
		} else {
			name, rank = user, ""
		}

		channel.Users[name] = rank
	}
}

// Join will create a new channel if the bot is the one joining. Otherwise it updates a channel's
// user list
func (l *Listeners) Join(msg *irc.Message) {
	l.bot.Mu.Lock()
	defer l.bot.Mu.Unlock()

	who := msg.Prefix.Name
	where := msg.Trailing

	if who == l.bot.GetNick() {
		rlog.Info("Bot", "Joined: "+where)

		caller, ok := l.bot.ToJoinChs[strings.ToLower(where)]
		if ok {
			l.bot.Say(caller, "Joined "+where)
		}

		channel := rbot.NewChannel(where)
		l.bot.Channels[strings.ToLower(where)] = channel
		return
	}

	channel := l.bot.Channels[strings.ToLower(where)]
	channel.Users[who] = ""
}

// Kick will remove a user from a channel list, or delete a channel if the bot was kicked.
func (l *Listeners) Kick(msg *irc.Message) {
	l.bot.Mu.Lock()

	who := msg.Params[1]
	where := msg.Params[0]

	l.bot.RemoveUser(who, where)

	l.bot.Mu.Unlock()
}

// Kill will remove a user from any present channels, or call the quit chan on the bot.
func (l *Listeners) Kill(msg *irc.Message) {
	l.bot.Mu.Lock()
	defer l.bot.Mu.Unlock()

	who := msg.Params[1]

	if who == l.bot.GetNick() {
		l.bot.Quit() <- errors.New("Bot was killed")
		return
	}

	for _, channel := range l.bot.Channels {
		delete(channel.Users, who)
	}
}

// Part will remove a user from the channel they have parted. If the bot parts, the channel is
// removed instead
func (l *Listeners) Part(msg *irc.Message) {
	l.bot.Mu.Lock()

	who := msg.Prefix.Name
	where := msg.Params[0]

	l.bot.RemoveUser(who, where)

	l.bot.Mu.Unlock()
}

// Quit will remove a user from any channels the bot is connected to. If the bot quits, the bot's
// quit chan is fulfilled.
func (l *Listeners) Quit(msg *irc.Message) {
	l.bot.Mu.Lock()
	defer l.bot.Mu.Unlock()

	who := msg.Prefix.Name

	if who == l.bot.GetNick() {
		l.bot.Quit() <- errors.New("Bot has quit from the irc server")
	}

	for _, channel := range l.bot.Channels {
		delete(channel.Users, who)
	}
}

// Nick will change a user's nick when it is changed, including the bot's.
func (l *Listeners) Nick(msg *irc.Message) {
	l.bot.Mu.Lock()
	defer l.bot.Mu.Unlock()

	who := msg.Prefix.Name

	if who == l.bot.GetNick() {
		l.bot.SetNick(msg.Params[0])
		return
	}

	for _, channel := range l.bot.Channels {
		rank, ok := channel.Users[who]

		if ok {
			delete(channel.Users, who)
			channel.Users[msg.Params[0]] = rank
		}
	}
}

// ErrNicknameInUse will check when a nick we choose is already in use, and will attempt to
// append a number until we have valid name.
func (l *Listeners) ErrNicknameInUse(msg *irc.Message) {
	l.bot.Mu.Lock()
	defer l.bot.Mu.Unlock()

	nick := msg.Params[0]
	lastChar := nick[len(nick)-1]

	num, err := strconv.Atoi(string(lastChar))

	if err != nil {
		l.bot.SetNick(nick + "1")
	} else {
		l.bot.SetNick(nick + strconv.Itoa(num+1))
	}
}
