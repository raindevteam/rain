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

package tmpl

// GOMTemplate is the template for go modules
const GOMTemplate = `package main

import (
	"os"

	"github.com/raindevteam/gorml"
	"gopkg.in/sorcix/irc.v1"
)

type {{.Name}} struct { *module.Module }

func (m *{{.Name}}) Greet(msg *irc.Message, args []string) {
	m.Say(msg.Args[0], "Hello there!")
}

func main() {
	m := &{{.Name}}{module.NewModule("{{.Name}}", "Your module's short description")}

	m.AddCommand("greet", &module.Command{
		Help: "It greets you!",
		Fun:  m.Info,
	})

	m.Register(os.Args)
}
`
