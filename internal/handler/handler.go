package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"

	"github.com/raindevteam/rain/internal/dapi"
	"github.com/raindevteam/rain/internal/hail"
)

//go:generate go run ../../tools/generate/dispatchers.go

// Handler contains all information necessary for handling listeners.
type Handler struct {
	Status   bool
	registry *Registry
	Commands map[string]map[string]*InternalCommand

	// For the websocket handler.
	addr     string
	upgrader websocket.Upgrader
}

// NewHandler will create a new handler.
func NewHandler() *Handler {
	h := &Handler{
		Status:   true,
		registry: &Registry{},
		Commands: make(map[string]map[string]*InternalCommand),
		addr:     ":8080",
		upgrader: websocket.Upgrader{},
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
func (h *Handler) InvokeCommand(cd *CommandData,
	e *discordgo.MessageCreate) error {
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

// StartPluginServer starts a websocket server for plugins to connect to.
// Currently just echos information sent to it.
func (h *Handler) StartPluginServer() {
	http.HandleFunc("/dapi", h.dapi)
	err := http.ListenAndServe(h.addr, nil)
	if err != nil {
		hail.Crit(hail.Fhandler, err.Error())
	}
}

// dapi serves the Droplet endpoint handler.
func (h *Handler) dapi(w http.ResponseWriter, r *http.Request) {
	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		hail.Critf(hail.Fhandler, "Failed to upgrade: %s\n", err.Error())
	}
	defer c.Close()

	hail.Info(hail.Fhandler, "Now serving DAPI")

	net := dapi.NewNetHelper(c)

	err = net.Write(dapi.NewHelloMessage())
	if ne, ok := err.(*dapi.NetError); ok {
		hail.Crit(hail.Fhandler, err.Error())
		err := net.WriteCloseErrMessage(ne)
		if err != nil {
			hail.Critf(hail.Fhandler,
				"Failed to write close message: %s\n", err.Error())
		}
	}

	identified := make(chan bool, 1)
	go func() {
		defer close(identified)

		// Read message
		netmsg, err := net.Read(time.Second * 4)
		if ne, ok := err.(*dapi.NetError); ok {
			if ne.ErrType == dapi.Timeout {
				hail.Crit(hail.Fhandler, err.Error())
				err = net.WriteCloseErrMessage(&dapi.NetError{
					dapi.IdentifyTimeout,
					"failed to identify in time",
				})
			} else {
				hail.Crit(hail.Fhandler, err.Error())
				err = net.WriteCloseErrMessage(ne)
			}
			if err != nil {
				hail.Critf(hail.Fhandler,
					"failed to write close message: %s\n", err.Error())
			}
			identified <- false
			return
		}

		// Check if message is not Identify
		if netmsg.Op != dapi.OpIdentify {
			hail.Warn(hail.Fhandler, "droplet sent message before identifying.")
			err := net.WriteCloseErrMessage(
				&dapi.NetError{
					dapi.BadIdentify,
					"sent message before identifying",
				})
			if err != nil {
				hail.Critf(hail.Fhandler,
					"Failed to write close message: %s\n", err.Error())
			}
			identified <- false
			return
		}

		// If we got here then module successfully identified
		identified <- true
	}()

	select {
	case isIdentified := <-identified:
		if !isIdentified {
			// TODO: Handle
			return
		}
		hail.Info(hail.Fhandler, "A droplet successfully identified")
		hail.Info(hail.Fhandler, "Closing its connection now.")
		err := net.WriteCloseMessage()
		if err != nil {
			hail.Critf(hail.Fhandler,
				"Failed to write close message: %s\n", err.Error())
		}
	}
}
