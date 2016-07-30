// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

import "gopkg.in/sorcix/irc.v1"

// CommandFun denotes a function to be used for IRC bot commands
type CommandFun func(*irc.Message, []string)

// TriggerFun denotes a function that determines whether a trigger should fire or not
type TriggerFun func(*irc.Message) bool

// Listener denotes a raw IRC listener function
type Listener func(*irc.Message)

// Event denotes an IRC event
type Event string

/********      These types are used mostly with the handler to keep track of modules       ********/

// CommandName is used for command names
type CommandName string

// ModuleName is used for module names
type ModuleName string

// Numeric is used to denote numeric IRC events
type Numeric string

/**************************************************************************************************/

// IrcData bundles together an Event with it's corresponding irc.Message. Used mostly to Send
// IRC data to modules via rpc.
type IrcData struct {
	Event Event
	Msg   *irc.Message
}

// CommandData bundles together a irc.Message, command name and it's given arguments to a module
// so that it may fire that respective command.
type CommandData struct {
	Msg  *irc.Message
	Name CommandName
	Args []string
}

/********            The following structs are used for internal listeners only            ********/

// A Command struct holds the function to fire for a command. Also holds help, fire on private
// message and fire on channel message properties.
type Command struct {
	Help string
	Fun  CommandFun
	PM   bool
	CM   bool
}

// A Trigger struct holds the validation check function and the function to fire if validation
// passes.
type Trigger struct {
	Check TriggerFun
	Fun   Listener
}

/**************************************************************************************************/
