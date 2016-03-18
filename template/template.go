package tmpl

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type ModuleConstruct struct {
	Name string
}

type TestConstruct struct {
	Lib       string
	SuiteName string
}

func CreateTestTemplate(lib string, name string) {
	construct := TestConstruct{lib, strings.Title(name)}

	var err error
	t := template.New("test")
	t, err = t.Parse(TestTemplate)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(name + "_test.go")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(f, construct)
	f.Close()
}

func GetModTemplate(m string) string {
	switch m {
	case "go":
		return GOMTemplate
	case "js":
		return JSMTemplate
	}
	return ""
}

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

func CreateModTemplate(mod string, name string) {
	tmpl := GetModTemplate(mod)

	var err error
	switch tmpl {
	case GOMTemplate:
		err = CreateModFile(tmpl, name, "go")
	case JSMTemplate:
		err = CreateModFile(tmpl, name, "js")
	}

	if err != nil {
		fmt.Println(err)
	}
}
