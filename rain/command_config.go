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
	"os"

	"github.com/urfave/cli"
)

// ExampleConfig is an example YAML configuration displaying all possible configuration options.
var ExampleConfig = `# Rain config
# -----------

# Server information
# ------------------
host : irc.examplehost.net
port : 6667

# Channel Information
# -------------------
# Channels listed here will be automatically joined upon server connect
channels :
  - "#Your"
  - "#Channels"

# User information
# ----------------
# nick is the only required field, the rest can be omitted and will default to the Nimbus default.
nick     : MyBot
realname : Wolfstein Jr. II
username : wolfstein
modes    : +B

# Debug level for internal IRC client
# -----------------------------------
# 1 = Print received IRC Messages
# 2 = Also print parsed values from IRC messages
nimbus-debug: 2

# Command Information
# -------------------
# The command prefix is used to denote commands from IRC users
commandprefix : .

# Modules (Optional)
# ------------------
# A '.' as a module path means its located in the default path. You can't use this for Go modules
# however, as they may be installed in different places such as github.com or gopkg.in. You may 
# set a specific path if you wish. You may also specify extra options for the module as a space 
# separated list by putting a ':' after its route. We currently only support the noload option.
modules :
  # JavaScript Modules
  js :
    umbrella : modulesOverHere/js:noload
    mlpshow  : .

  # Python Modules
  py :
    echo  : ../andSomeMoreOverHere/py
    ltest : .

  # Go Modules
  go :
    raincore : github.com/raindevteam

# Default Routes (Optional)
# -------------------------
# Changing these will overwrite the default module routes for their respective module type. Keep in
# mind that modules/<type> is already the default type for js and py. Go has no default module route
# set due to how Go modules are installed.
#
# We don't recommend setting a default go route as you may have packages installed in different
# places such as github.com and gopkg.in. You also can't move go packages outside your GOPATH.
default-routes :
  js : modules/js
  py : modules/py

# Global Options (Optional)
# -------------------------
# Global options define options to be used for all modules that fall under their respective type
global-options :
  js : npm
  py : noload
  go : noload
`

func config(c *cli.Context) error {
	fmt.Println("Making an example config.yaml...")

	f, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println("Oops, something went wrong: " + err.Error())
		return err
	}

	_, err = f.Write([]byte(ExampleConfig))
	if err != nil {
		fmt.Println("Couldn't write config: " + err.Error())
		return err
	}

	f.Close()
	fmt.Println("All done!")
	return nil
}

// CommandConfig is a command that allows the user to create an example config.yaml. The config will
// show all possible config options a user can assign.
var CommandConfig = cli.Command{
	Name:        "config",
	Aliases:     []string{"c"},
	Usage:       "Creates an example YAML config file",
	Description: "Will create a new example config.yaml file in the current directory.",
	Action:      config,
}
