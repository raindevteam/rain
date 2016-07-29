package parser

import (
	"reflect"
	"testing"
)

var parser = NewParser(".")

func TestParseModuleValue(t *testing.T) {
	var (
		route   string
		options []string
	)

	route, options = parser.ParseModuleValue(".:pip noload")
	if route != "." || !reflect.DeepEqual(options, []string{"pip", "noload"}) {
		t.Fatal("\".:pip noload\" incorrectly parsed by parser")
	}

	route, options = parser.ParseModuleValue(".")
	if route != "." || !reflect.DeepEqual(options, []string{}) {
		t.Fatal("\".\" incorrectly parsed by parser")
	}

	route, options = parser.ParseModuleValue("modules/js:npm")
	if route != "modules/js" || !reflect.DeepEqual(options, []string{"npm"}) {
		t.Fatal("\".:pip noload\" incorrectly parsed by parser")
	}
}

func TestIsCommand(t *testing.T) {
	// Loop these
	check := parser.IsCommand(".yes")
	check = !parser.IsCommand("no, not a command") && check
	check = parser.IsCommand(".yes I am a command") && check
	check = !parser.IsCommand(";nope, I am not a command either") && check
	if !check {
		t.Fatal("Parser unable to correctly parse commands")
	}
}

func TestParseCommand(t *testing.T) {
	var (
		name string
		args []string
	)

	name, args = parser.ParseCommand(".say")
	if name != "say" || len(args) != 0 {
		t.Fatal("\".say\" incorrectly parsed by parser")
	}

	name, args = parser.ParseCommand(".echo I am broot")
	if name != "echo" || !reflect.DeepEqual(args, []string{"I", "am", "broot"}) {
		t.Fatal("\".echo I am broot\" incorrectly parsed by parser")
	}
}

func TestParsePrefix(t *testing.T) {
	var who, host string

	who, host = parser.ParsePrefix("NimBot")
	if who != "NimBot" || host != "" {
		t.Fatal("\"NimBot\" incorrectly parsed by parser")
	}

	who, host = parser.ParsePrefix("NimBot!this.host")
	if who != "NimBot" || host != "this.host" {
		t.Fatal("\"NimBot!this.host\" incorrectly parsed by parser")
	}
}
