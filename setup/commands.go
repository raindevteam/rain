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
	"github.com/raindevteam/rain/rbot"
	"gopkg.in/sorcix/irc.v1"
)

///////////////////////////             Module Manager            //////////////////////////////////

// The ModuleManager provides the M command, allowing for the management of modules loaded via the
// bot's module system.
type ModuleManager struct{ bot *rbot.Bot }

// GetCommand returns this struct's irc command wrapped in a *rbot.Command struct
func (c *ModuleManager) GetCommand() *rbot.Command {
	command := &rbot.Command{
		Help: "The RainBot Module Manager helps manage modules.",
		Fun:  c.M,
		PM:   true,
		CM:   true,
	}

	return command
}

// M routes all sub commands to their appropriate helper functions
func (c *ModuleManager) M(msg *irc.Message, Params []string) {
	if len(Params) == 0 {
		c.bot.Say(msg.Params[0], "Error: No command specified")
		return
	}

	if len(Params) == 1 {
		c.bot.Say(msg.Params[0], "Error: No module specified")
		return
	}

	switch Params[0] {
	case "reload":
		c.reload(msg, Params)
	case "kill":
		c.kill(msg, Params)
	case "load":
		c.load(msg, Params)
	case "info":
		c.info(msg, Params)
	default:
		c.bot.Say(msg.Params[0], Params[0]+" is not a valid subcommand")
	}
}

func (c *ModuleManager) reload(msg *irc.Message, Params []string) {
	if err := c.bot.ModuleReload(Params[1]); err != nil {
		c.bot.Say(msg.Params[0], "Error: "+err.Error())
		return
	}

	c.bot.Say(msg.Params[0], "Reloaded module: "+Params[1])
}

func (c *ModuleManager) kill(msg *irc.Message, Params []string) {
	// Not doable yet
}

func (c *ModuleManager) info(msg *irc.Message, Params []string) {
	info, err := c.bot.ModuleInfo(Params[1])
	if err != nil {
		c.bot.Say(msg.Params[0], "Error: "+err.Error())
		return
	}

	c.bot.Say(msg.Params[0], info)
}

func (c *ModuleManager) load(msg *irc.Message, Params []string) {
	if err := c.bot.ModuleLoad(Params[1]); err != nil {
		c.bot.Say(msg.Params[0], "Error: "+err.Error())
		return
	}

	c.bot.Say(msg.Params[0], "Loaded module: "+Params[1])
}
