// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package main

import (
	"os"

	"github.com/raindevteam/rain/rain"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "rain"
	app.Usage = "A command line tool for Rain."

	rain.NewCmdGenhand(app)

	app.Run(os.Args)
}
