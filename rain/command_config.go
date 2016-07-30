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
	"os"

	"github.com/urfave/cli"
)

var ExampleConfig = `# Rain config
# 0.6.0-alpha.1

# Server information
host : irc.examplehost.net
port : 6667

# Channel Information
# Channels listed here will be automatically joined upon server connect
channels :
  - "#Your"
  - "#Channels"

# User information
nick     : MyBot
realname : Wolfstein Jr. II
username : wolfstein
modes    : +B

# Command Information
# The command prefix is used to denote commands from IRC users
commandprefix : .

# Modules
# These are all optional. A '.' as a module path means its located in the default path. You may 
# set a specific path if you wish. You may specify extra options for the module as a space separated
# list by putting a ':' after its route. We currently only support a few options.
modules :
  # JavaScript Modules
  js :
    umbrella : modulesOverHere/js:npm noload
    mlpshow  : .

  # Python Modules
  py :
    echo  : ../andSomeMoreOverHere/py:pip
    ltest : .

  # Go Modules
  go :
    raincore : github.com/raindevteam

# These are optional
# We don't recommend setting a default go route as you may have packages installed in different
# places such as github.com and gopkg.in. You also can't move go packages outside your GOPATH.
default-routes :
  js : modules/js
  py : modules/py

# Global options define options to be used for all modules that fall under their respective type
global-options :
  js : npm
  py : noload
  go : noload
`

func config(c *cli.Context) error {
	fmt.Println(" Making an example config.yaml...")

	f, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println(" Oops, something went wrong: " + err.Error())
		return err
	}

	_, err = f.Write([]byte(ExampleConfig))
	if err != nil {
		fmt.Println(" Couldn't write config: " + err.Error())
		return err
	}

	f.Close()
	fmt.Println(" All done!")
	return nil
}

var CommandConfig = cli.Command{
	Name:        "config",
	Aliases:     []string{"c"},
	Usage:       "Creates an example YAML config",
	Description: "Will create a new example config.yaml file in the current directory.",
	Action:      config,
}
