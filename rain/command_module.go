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

package rain

import (
	"fmt"

	"github.com/raindevteam/rain/tmpl"
	"github.com/urfave/cli"
)

var rmls = []string{
	"go",
	"py",
	"js",
}

func module(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println(c.App.ArgsUsage)
	}

	if !isIn(rmls, c.Args().First()) {
		fmt.Printf(" Not a valid install type. Valid install types are:\n%s", listPrint(rmls))
	} else {
		if c.Args().Get(1) == "" {
			fmt.Println(" No name specified")
			fmt.Println(" Usage: " + c.Command.Name + " rml-prefix module-name")
			return nil
		}

		tmpl.CreateModTemplate(c.Args().Get(0), c.Args().Get(1))
		fmt.Printf(" Created module template: %s.%s", c.Args().Get(1), c.Args().Get(0))
	}

	return nil
}

var CommandModule = cli.Command{
	Name:        "module",
	Aliases:     []string{"m"},
	Usage:       "Creates boilerplate for a module",
	Description: "Creates boilerplate for a module tailored to the specified RML.",
	ArgsUsage:   "rml-prefix module-name",
	Action:      module,
}
