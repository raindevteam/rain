// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU Affero General Public License as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License along with this program.
// If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

import (
	"net/rpc"
	"strings"
	"sync"

	"gopkg.in/sorcix/irc.v1"

	"github.com/raindevteam/rain/rlog"
)

// The Handler is used to manage RPC clients and their respective commands and listeners. It will
// invoke appropriate RPC calls when needed. It will also in the future handle pings, inbound and
// outbound, to/from modules. It also manages internal commands and listeners. A callback function
// should be registered with the handler so that the bot will no when RPC clients are removed.
type Handler struct {
	// Listeners for modules
	Commands map[CommandName]ModuleName
	Triggers map[Event][]ModuleName
	Numerics map[Numeric][]ModuleName

	// Listeners for Bot
	InternalCommands map[CommandName]*Command
	InternalTriggers map[Event][]*Trigger
	InternalNumerics map[Numeric][]*Listener

	Modules        map[ModuleName]*rpc.Client
	removeCallback func(ModuleName)
	mu             sync.RWMutex
}

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	handler := &Handler{
		Commands: make(map[CommandName]ModuleName),
		Triggers: make(map[Event][]ModuleName),
		Numerics: make(map[Numeric][]ModuleName),

		InternalCommands: make(map[CommandName]*Command),
		InternalTriggers: make(map[Event][]*Trigger),
		InternalNumerics: make(map[Numeric][]*Listener),

		Modules: make(map[ModuleName]*rpc.Client),
		mu:      sync.RWMutex{},
	}
	return handler
}

// AddModule adds a new module to the handler by adding its respective rpc client.
func (h *Handler) AddModule(name ModuleName, module *rpc.Client) {
	h.Modules[name] = module
}

// AddRemoveCallback will register a callback function, which is called when a module has been
// removed. The module's name is passed as a parameter.
func (h *Handler) AddRemoveCallback(callback func(ModuleName)) {
	h.removeCallback = callback
}

// ModuleExists checks if a module exists within the handler.
func (h *Handler) ModuleExists(name string) bool {
	_, ok := h.Modules[ModuleName(name)]
	return ok
}

// SignalCleanup calls the Cleanup method for the given module. Returns an error if module did not
// successfully cleanup after itself.
func (h *Handler) SignalCleanup(name ModuleName) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	result := ""
	err := h.Modules[name].Call(string(name)+".Cleanup", nil, &result)

	if err != nil {
		return err
	}

	return nil
}

// AddCommand adds a command to the handler by using its name as a key for the module's name. The
// module name is then used to invoke the command via rpc.
func (h *Handler) AddCommand(cmd CommandName, name ModuleName) {
	h.Commands[cmd] = ModuleName(strings.ToLower(string(name)))
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
func (h *Handler) Invoke(msg *irc.Message, cmd CommandName, args []string) {
	h.mu.RLock()

	if h.IsInternalCommand(cmd) {
		h.InvokeInternal(msg, cmd, args)
		h.mu.RUnlock()
	} else {
		mName, ok := h.Commands[CommandName(strings.ToLower(string(cmd)))]
		if !ok {
			rlog.Debug("Handler", string(cmd)+" [Command Listener] does not exist")
			return
		}

		_, err := h.doRPC(mName, "InvokeCommand", &CommandData{msg, cmd, args})
		h.mu.RUnlock()
		if err != nil {
			h.removeUnstableClient(mName, err)
		}
	}
}

// InvokeInternal runs an internal command.
func (h *Handler) InvokeInternal(msg *irc.Message, cmd CommandName, args []string) {
	listener := h.InternalCommands[cmd]
	listener.Fun(msg, args)
}

// Fire calls all modules that have registered an event
func (h *Handler) Fire(msg *irc.Message, e Event) {
	h.FireInternal(msg, e)

	for _, mName := range h.Triggers[e] {
		go func(h *Handler, name ModuleName, msg *irc.Message) {
			h.mu.RLock()

			result := ""
			h.Modules[name].Call(string(name)+".FireTriggers", msg, &result)

			h.mu.RUnlock()
		}(h, mName, msg)
	}
}

// FireInternal will fire all internal listeners of the bot.
func (h *Handler) FireInternal(msg *irc.Message, e Event) {
	for _, trigger := range h.InternalTriggers[e] {
		go func(msg *irc.Message, trigger *Trigger) {
			if trigger.Check(msg) {
				trigger.Fun(msg)
			}
		}(msg, trigger)
	}
}

// RemoveModule will first remove all commands and listeners of a module from the handler. It will
// then remove its RPC client.
func (h *Handler) RemoveModule(name ModuleName) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Let's make sure we haven't already removed the module, as removal can be initiated by
	// multiple sources.
	if !h.ModuleExists(string(name)) {
		// Our job here is/was done
		return
	}

	// Remove all commands
	for cmd, modname := range h.Commands {
		if modname == name {
			delete(h.Commands, cmd)
		}
	}

	// Remove all triggers
	for event, modnames := range h.Triggers {
		for i, modname := range modnames {
			if modname == name {
				rlog.Debug("here", "here")
				h.Triggers[event] = append(h.Triggers[event][:i], h.Triggers[event][i+1:]...)
			}
		}
	}

	// And finally remove from handler
	delete(h.Modules, name)
}

func (h *Handler) removeUnstableClient(name ModuleName, err error) {
	// RPC call failed, module deemed unstable and is therefore removed

	// Signal a clean up to module
	rlog.Warn("Handler", "Removing "+string(name)+" [Module Client] due to RPC error: "+
		err.Error())

	err = h.SignalCleanup(name)
	if err != nil {
		rlog.Warn("Handler", string(name)+
			" [Module Client] did not successfully clean up, it will still be removed.")
	}

	h.RemoveModule(name)
	h.signalRemove(name)
}

func (h *Handler) signalRemove(name ModuleName) {
	h.removeCallback(name)
}

func (h *Handler) doRPC(name ModuleName, procedure string, data interface{}) (string, error) {
	result := ""
	err := h.Modules[name].Call(string(name)+"."+procedure, data, &result)
	return result, err
}
