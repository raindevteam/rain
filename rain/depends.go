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
	"os/exec"
)

var AllDepends = []string{
	"github.com/RyanPrintup/Nimbus",
	"gopkg.in/readline.v1",
	"gopkg.in/yaml.v2",
}

var CoreDepends = []string{
	"github.com/RyanPrintup/Nimbus",
	"gopkg.in/yaml.v2",
}

var NeverfreeDepends = []string{
	"gopkg.in/readline.v1",
}

func GetDepends(depends []string) error {
	for _, dep := range depends {
		output, err := exec.Command("go", "get", "-u", dep).CombinedOutput()
		if err != nil {
			fmt.Printf(" Internal command error\n ----\n%s\n ----\n\n", string(output[:]))
			return err
		}
	}
	return nil
}
