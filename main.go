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

package main

import (
	"os"

	"github.com/raindevteam/rain/rain"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rain"
	app.Usage = "Command line tool for the management of everything Rain"
	app.Version = rain.Version()

	app.Commands = []cli.Command{
		rain.CommandInstall,
		rain.CommandModule,
		rain.CommandConfig,
		rain.CommandDepends,
		rain.CommandVersion,
	}

	app.Run(os.Args)
}
