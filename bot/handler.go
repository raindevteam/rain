package rainbot

import (
	"net/rpc"
	"strings"

	"github.com/RyanPrintup/nimbus"
)

type Handler struct {
	Commands map[CommandName]ModuleName
	Triggers map[Event][]ModuleName
	Numerics map[Numeric][]ModuleName

	InternalCommands map[CommandName]*Command
	InternalTriggers map[Event][]*Trigger
	InternalNumerics map[Numeric][]*Listener

	Modules map[ModuleName]*rpc.Client
}

func NewHandler() *Handler {
	handler := &Handler{
		Commands: make(map[CommandName]ModuleName),
		Triggers: make(map[Event][]ModuleName),
		Numerics: make(map[Numeric][]ModuleName),

		InternalCommands: make(map[CommandName]*Command),
		InternalTriggers: make(map[Event][]*Trigger),
		InternalNumerics: make(map[Numeric][]*Listener),

		Modules: make(map[ModuleName]*rpc.Client),
	}
	return handler
}

// AddModule adds a new module to the handler by adding its respective rpc client.
func (h *Handler) AddModule(mName ModuleName, module *rpc.Client) {
	h.Modules[mName] = module
}

// ModuleExists checks if a module exists within the handler.
func (h *Handler) ModuleExists(mName string) bool {
	_, ok := h.Modules[ModuleName(mName)]
	return ok
}

// SignalKill calls the Cleanup method for the given module. Returns an error if module did not
// successfully cleanup after itself.
func (h *Handler) SignalKill(mName string) error {
	result := ""
	err := h.Modules[ModuleName(mName)].Call(mName+".Cleanup", nil, &result)

	if err != nil {
		return err
	}

	return nil
}

// AddCommand adds a command to the handler by using its name as a key for the module's name. The
// module name is then used to invoke the command via rpc.
func (h *Handler) AddCommand(cmd CommandName, mName ModuleName) {
	h.Commands[cmd] = mName
}

// AddInternalCommand adds an internal command to the handler. It uses the name as a key for the
// command's respective struct.
func (h *Handler) AddInternalCommand(cname CommandName, cfunc *Command) {
	h.InternalCommands[cname] = cfunc
}

// AddTrigger adds an event to the handler using the event as a key.
func (h *Handler) AddTrigger(mName ModuleName, e Event) {
	h.Triggers[e] = append(h.Triggers[e], mName)
}

// AddInternalTrigger does the same as AddTrigger except it instead adds a trigger struct to the
// list of internal triggers.
func (h *Handler) AddInternalTrigger(event Event, trigger *Trigger) {
	h.InternalTriggers[event] = append(h.InternalTriggers[event], trigger)
}

// IsInternalCommand checks if a command name is an internal command.
func (h *Handler) IsInternalCommand(cmd CommandName) bool {
	_, ok := h.InternalCommands[cmd]
	return ok
}

// Invoke runs a command. Commands are ran by calling the respective module's "InvokeCommand" rpc
// method. The module then handles the firing.
func (h *Handler) Invoke(msg *nimbus.Message, cmd CommandName, args []string) {
	if h.IsInternalCommand(cmd) {
		h.InvokeInternal(msg, cmd, args)
	} else {
		mName, ok := h.Commands[CommandName(strings.ToLower(string(cmd)))]
		if !ok {
			return
		}
		result := ""
		h.Modules[mName].Call(string(mName)+".InvokeCommand",
			&CommandData{msg, cmd, args}, &result)
		// Handle Error
	}
}

// InvokeInternal runs an internal command.
func (h *Handler) InvokeInternal(msg *nimbus.Message, cmd CommandName, args []string) {
	hook := h.InternalCommands[cmd]
	hook.Fun(msg, args)
}

func (h *Handler) Fire(msg *nimbus.Message, e Event) {
	h.FireInternal(msg, e)

	for _, mName := range h.Triggers[e] {
		result := ""
		go h.Modules[mName].Call(string(mName)+".FireTriggers", msg, &result)
	}
}

func (h *Handler) FireInternal(msg *nimbus.Message, e Event) {
	for _, trigger := range h.InternalTriggers[e] {
		go func(msg *nimbus.Message) {
			if trigger.Check(msg) {
				trigger.Fun(msg)
			}
		}(msg)
	}
}
