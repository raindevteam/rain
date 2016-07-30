// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rain

import (
	"fmt"

	"github.com/urfave/cli"
)

var installs = []string{
	"default",
	"neverfree",
	"simple",
}

func install(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println(c.App.ArgsUsage)
	}

	if !isIn(installs, c.Args().First()) {
		fmt.Printf(" Not a valid install type. Valid install types are:\n%s", listPrint(installs))
	} else {
		if c.Args().Get(1) == "" {
			fmt.Println(" No name specified")
			fmt.Println(" Usage: " + c.Command.Name + " install bot-name")
			return nil
		}

		fmt.Println(" Building... ")
		if err := doInstall(c.Args().Get(0), c.Args().Get(1)); err != nil {
			fmt.Println(" Install was not successful, please report the above problem(s) as an issue")
			return err
		}
		fmt.Println(" Done!")
	}

	return nil
}

var CommandInstall = cli.Command{
	Name:    "install",
	Aliases: []string{"i"},
	Usage:   "Creates a new bot with a given install and name",
	Description: "Creates a new bot by buidling it using the \"go build\" command.\n\n" +
		"   Possible build types are:\n" + listPrint(installs),
	ArgsUsage: "install name",
	Action:    install,
}
