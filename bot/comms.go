package rainbot

import "github.com/RyanPrintup/nimbus"

type CommandFun func(*nimbus.Message, []string)
type TriggerFun func(*nimbus.Message) bool
type Listener func(*nimbus.Message)

type Event string

type CommandName string
type ModuleName string
type Numeric string

type Command struct {
	Help string
	Fun  CommandFun
	PM   bool
	CM   bool
}

type Trigger struct {
	Check TriggerFun
	Fun   Listener
}

type IrcData struct {
	Event Event
	Msg   *nimbus.Message
}

type CommandData struct {
	Msg  *nimbus.Message
	Name CommandName
	Args []string
}
