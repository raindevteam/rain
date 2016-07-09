package tmpl

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var testTemplDir = "../tests/premade-templates/"

var modtypes = map[string]string{
	"go": GOMTemplate,
	"js": JSMTemplate,
	"py": PYMTemplate,
}

func checkTmplAgainst(premade string, created string) (bool, error) {
	var err error

	file, err := ioutil.ReadFile(premade)
	if err != nil {
		return false, err
	}

	templ, err := ioutil.ReadFile(created)
	if err != nil {
		return false, err
	}

	return string(file[:]) == string(templ[:]), nil
}

func TestMain(m *testing.M) {
	ret := m.Run()
	tearDown()
	os.Exit(ret)
}

func TestGetModTemplate(t *testing.T) {
	for mod, templ := range modtypes {
		mod := GetModTemplate(mod)
		if mod != templ {
			t.Error("Couldn't match " + mod + " to template")
		}
	}
}

func TestCreateModFile(t *testing.T) {
	err := CreateModFile(GOMTemplate, "mine", "go")
	if err != nil {
		t.Error("Could not create file: " + err.Error())
	}

	checks, err := checkTmplAgainst(testTemplDir+"go_module.txt", "mine.go")
	if err != nil {
		t.Error("Error while checking templates: " + err.Error())
	}

	if checks != true {
		t.Error("mine.go didn't match predefined module")
	}
}

func TestCreateModTemplate(t *testing.T) {
	for mod, templ := range modtypes {
		CreateModTemplate(mod, templ)
		checks, err := checkTmplAgainst(testTemplDir+mod+"_module.txt", "test."+mod)
		if err != nil {
			t.Error("Error while checking templates: " + err.Error())
		}

		if checks != true {
			t.Error("test." + mod + " didn't match predefined module")
		}
	}
}

func tearDown() {
	for _, file := range []string{"mine.go, test.go, test.js, test.py"} {
		if err := os.Remove(file); err != nil {
			fmt.Println("Could not remove test files")
		}
	}
}
