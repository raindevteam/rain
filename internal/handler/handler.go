package handler

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"github.com/raindevteam/rain/internal/hail"
)

//go:generate go run ../../tools/generate/dispatchers.go

// Handler contains all information necessary for handling listeners.
type Handler struct {
	Status   bool
	registry *Registry
	Commands map[string]map[string]*InternalCommand
}

// NewHandler will create a new handler.
func NewHandler() *Handler {
	h := &Handler{
		Status:   true,
		registry: &Registry{},
		Commands: make(map[string]map[string]*InternalCommand),
	}
	h.registry.Initialize()
	return h
}

// CommandData holds information pertaining to a command. The command will
// always be specified but the owner is not guaranteed to be filled.
type CommandData struct {
	Owner   string
	Command string
}

// An Action is an interface wrapper for listener commands. It is primarily used
// to execute the function by passing it a discord event struct. The function
// will handle the type conversion.
type Action interface {
	Do(v interface{})
}

// AddInternalListener takes a function and wraps it into a listener. This
// listener is then added to registry and returned to the caller. If the value
// passed to this function is not valid, listener will be nil.
func (h *Handler) AddInternalListener(v interface{}) Listener {
	return h.registry.CreateListener(v, true)
}

// AddInternalCommand takes a function and wraps it into a command.
func (h *Handler) AddInternalCommand(name string,
	f func(string, *discordgo.MessageCreate) error) {
	ic := &InternalCommand{}
	ic.CommandFunc = f
	ic.SetName(name)
	h.registry.AddCommand(ic)
}

// InvokeCommand takes a CommandData struct and uses it to run a specified
// command.
func (h *Handler) InvokeCommand(cd *CommandData, e *discordgo.MessageCreate) error {
	// Droplets aren't implement yet.
	if cd.Owner != Internal {
		return errors.New("droplet commands are not implemented yet")
	}

	// TODO: Create registry get methods for commands and listeners.
	command, ok := h.registry.Commands[cd.Owner][cd.Command]
	if !ok {
		hail.Warn(hail.Fhandler, "not a command")
		return nil
	}
	command.SetArguments("")
	command.SetEvent(e)
	command.Invoke()

	return nil
}
