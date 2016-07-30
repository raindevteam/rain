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

package tmpl

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// ModuleConstruct is used to pass data to a template.
type ModuleConstruct struct {
	Name string
}

// GetModTemplate returns a predefined string representing the template type passed to it.
func GetModTemplate(m string) string {
	switch m {
	case "go":
		return GOMTemplate
	case "js":
		return JSMTemplate
	case "py":
		return PYMTemplate
	}
	return ""
}

// CreateModFile takes a template, file name and extension to create a file with the forementioned
// template.
func CreateModFile(tmpl string, name string, ext string) error {
	t := template.New(name)
	t, err := t.Parse(tmpl)

	if err != nil {
		fmt.Println(err)
	}

	construct := ModuleConstruct{strings.Title(name)}
	f, err := os.Create(strings.ToLower(name) + "." + ext)
	if err != nil {
		return err
	}
	t.Execute(f, construct)
	f.Close()
	return nil
}

// CreateModTemplate will write a template of appropriate type to a file.
func CreateModTemplate(mod string, name string) {
	tmpl := GetModTemplate(mod)

	var err error

	switch tmpl {
	case GOMTemplate:
		err = CreateModFile(tmpl, name, "go")
	case JSMTemplate:
		err = CreateModFile(tmpl, name, "js")
	case PYMTemplate:
		err = CreateModFile(tmpl, name, "py")
	}

	if err != nil {
		fmt.Println(err)
	}
}
