// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package cbot

import (
	"github.com/RyanPrintup/nimbus"
	"gopkg.in/sorcix/irc.v1"
)

type Handle struct {
	listener nimbus.Listener
	done     chan bool
}

func (h *Handle) Run(msg *irc.Message) {
	h.listener(msg)
	h.done <- true
}

func (c *CliClient) AddListener(event string, l nimbus.Listener) {
	c.listeners[event] = append(c.listeners[event], l)
}

func (c *CliClient) GetListeners(event string) []nimbus.Listener {
	return c.listeners[event]
}

func (c *CliClient) Emit(event string, msg *irc.Message) {
	for _, listener := range c.listeners[event] {
		h := Handle{listener, make(chan bool)}
		go h.Run(msg)
		<-h.done
	}
}
