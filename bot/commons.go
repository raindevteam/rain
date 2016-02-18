package rainbot

import "github.com/RyanPrintup/nimbus"

// CommandFun denotes a function to be used for IRC bot commands
type CommandFun func(*nimbus.Message, []string)

// TriggerFun denotes a function that determines whether a trigger should fire or not
type TriggerFun func(*nimbus.Message) bool

// Listener denotes a raw IRC listener function
type Listener func(*nimbus.Message)

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

// IrcData bundles together an Event with it's corresponding nimbus.Message. Used mostly to Send
// IRC data to modules via rpc.
type IrcData struct {
	Event Event
	Msg   *nimbus.Message
}

// CommandData bundles together a nimbus.Message, command name and it's given arguments to a module
// so that it may fire that respective command.
type CommandData struct {
	Msg  *nimbus.Message
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
