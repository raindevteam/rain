package hail

import (
	"fmt"
)

// Emerg is like Emergf but adds a new line. Does not accept format specifiers.
func Emerg(facility int, msg string) error {
	return Emergf(facility, msg+"\n")
}

// Emergf takes a facility specifier, format string and arguments to parse into
// said string. It returns an error with msgf as the error string.
func Emergf(f int, msgf string, v ...interface{}) error {
	err := fmt.Errorf(msgf, v...)
	if Lemerg&l.logmodes != 0 {
		fmt.Print(prefix(f, Semerg) + err.Error())
	}
	return err
}
