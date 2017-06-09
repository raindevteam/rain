package handler

import (
	"github.com/raindevteam/rain/hail"
)

//go:generate go run ../tools/main.go

// Handler contains all information necessary for handling listeners.
type Handler struct {
	Status   bool
	registry *Registry
}

// H is the global singleton instance of the handler. It should be instantiated
// by calling CreateHandler().
var H *Handler

// CreateHandler will instantiate H. If it was done so in another way or H was
// already created, CreateHandler will return an error.
func CreateHandler() error {
	if H != nil {
		return hail.Err(hail.Feventhand, "a handler has already been created")
	}
	H = &Handler{
		Status:   true,
		registry: &Registry{},
	}
	H.registry.Initialize()
	return nil
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
func AddInternalListener(v interface{}) Listener {
	return H.registry.CreateListener(v, true)
}
