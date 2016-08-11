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
	"gopkg.in/sorcix/irc.v1"

	"github.com/raindevteam/rain/rbot"
)

// Default is the default setup for listeners on the bot. It includes the module manager command,
// a privmsg listener to parse commands, and common irc housekeeping listeners (Join, kick, etc).
func Default(b *rbot.Bot) {
	mm := ModuleManager{b}
	l := Listeners{b}

	b.Handler.AddInternalCommand("m", mm.GetCommand())

	b.AddListener(irc.RPL_TOPIC, l.RplTopic)
	b.AddListener(irc.RPL_NAMREPLY, l.RplNamreply)

	b.AddListener(irc.PRIVMSG, l.Privmsg)
	b.AddListener(irc.JOIN, l.Join)
	b.AddListener(irc.KICK, l.Kick)
	b.AddListener(irc.KILL, l.Kill)
	b.AddListener(irc.PART, l.Part)
	b.AddListener(irc.QUIT, l.Quit)
	b.AddListener(irc.NICK, l.Nick)

	b.AddListener(irc.ERR_NICKNAMEINUSE, l.ErrNicknameInUse)
}
