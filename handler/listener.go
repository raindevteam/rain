// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package handler

import (
	"github.com/raindevteam/rain/droplet"
)

type Handler struct {
	Status   bool
	registry *Registry
}

type Action interface{}

///////////////////                Listener                  ///////////////////

// Listener is the interface for both internal listeners and droplet
// listeners. Use the Owner() function to discern if a listener is internal. If
// it is, it will return the string constant "__INTERNAL__".
type Listener interface {
	Run(v interface{})
	Action() Action
	Owner() string
}

// InternalListener is the struct for internal listeners. It does not have a
// droplet field like DropletListener.
type InternalListener struct {
	enabled bool
	act     Action
}

// Run will execute the listener's Action.
func (il InternalListener) Run(v interface{}) {
	runAction(v, il.act)
}

// Action will return the listener's Action.
func (il InternalListener) Action() Action {
	return il.act
}

// Owner for InternalListener returns the string constant "__INTERNAL__".
func (il InternalListener) Owner(v interface{}) string {
	return Internal
}

type DropletListener struct {
	drop *droplet.Droplet
	act  Action
}

func (l DropletListener) Run(v interface{}) {
}

func (h Handler) AddInternalListener(l Listener) {
	h.registry.addListener(l)
}

///////////////////                 Command                  ///////////////////
