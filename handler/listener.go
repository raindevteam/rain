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

//go:generate go run ../tools/main.go

type Handler struct {
	Status   bool
	registry *Registry
}

type Action interface {
	Do(v interface{})
}

///////////////////                Listener                  ///////////////////

// Listener is the interface for both internal listeners and droplet
// listeners. Use the Owner() function to discern if a listener is internal. If
// it is, it will return the string constant "__INTERNAL__".
type Listener interface {
	Run()
	SetActionHandler(a Action)
	SetEvent(v interface{})
	Owner() string
}

// InternalListener is the struct for internal listeners. It does not have a
// droplet field like DropletListener.
type InternalListener struct {
	Event   interface{}
	enabled bool
	action  Action
}

// Run will execute the listener's Action.
func (il InternalListener) Run() {
	il.action.Do(il.Event)
}

// Action will return the listener's Action.
func (il InternalListener) SetActionHandler(a Action) {
	il.action = a
}

func (il InternalListener) SetEvent(v interface{}) {
	il.Event = v
}

// Owner for InternalListener returns the string constant "__INTERNAL__".
func (il InternalListener) Owner() string {
	return Internal
}

type DropletListener struct {
	Event  interface{}
	drop   *droplet.Droplet
	action Action
}

func (l DropletListener) Run() {
}

func (l DropletListener) SetActionHandler(a Action) {
	l.action = a
}

func (l DropletListener) SetEvent(v interface{}) {
	l.Event = v
}

func (l DropletListener) Owner() string {
	return "module"
}

func (h Handler) AddInternalListener(l Listener) {
}

///////////////////                 Command                  ///////////////////
