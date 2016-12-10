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
	"os/exec"
)

// AllDepends are all the dependencies needed to install any premade bot.
var AllDepends = []string{
	"github.com/RyanPrintup/nimbus",
	"gopkg.in/readline.v1",
	"gopkg.in/yaml.v2",
	"golang.org/x/time/rate",
}

// CoreDepends are all the dependencies needed for for all premade bot.
var CoreDepends = []string{
	"github.com/RyanPrintup/nimbus",
	"gopkg.in/yaml.v2",
	"golang.org/x/time/rate",
}

// NeverfreeDepends are all the depends needed to install the premade Neverfree bot.
var NeverfreeDepends = []string{
	"gopkg.in/readline.v1",
}

// GetDepends will "go get" the dependencies specified as an array.
func GetDepends(depends []string) error {
	for _, dep := range depends {
		fmt.Printf("get %s... ", dep)
		output, err := exec.Command("go", "get", "-u", dep).CombinedOutput()
		fmt.Println("done!")
		if err != nil {
			fmt.Println("Internal command error")
			fmt.Println(fmtCmdOutput(output))
			return err
		}
	}
	return nil
}
