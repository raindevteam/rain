package hail

import "fmt"

// Crit is like Critf but adds a new line. Does not accept format specifiers.
func Crit(facility int, msg string) error {
	return Critf(facility, msg+"\n")
}

// Critf takes a facility specifier, format string and arguments to parse into
// said string.
func Critf(f int, msgf string, v ...interface{}) error {
	err := fmt.Errorf(msgf, v...)
	if Lcrit&l.logmodes != 0 {
		fmt.Print(prefix(f, Scrit) + err.Error())
	}
	return err
}
