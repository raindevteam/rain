package tmpl

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type JsModuleConstruct struct {
	Name string
}

type GoModuleConstruct struct {
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

/*
func GetModTemplate(m string) string {
	switch m {
	case "go":
		return GOMTemplate
	case "js":
		return JSMTemplate
	}
}

func CreateModTemplate(lib string, name string) {
	construct := TestConstruct{lib, strings.Title(name)}

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
}
*/
