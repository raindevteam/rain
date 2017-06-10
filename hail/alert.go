package hail

import "fmt"

// Alert is like Alertf but adds a new line. Does not accept format specifiers.
func Alert(facility int, msg string) error {
	return Alertf(facility, msg+"\n")
}

// Alertf takes a facility specifier, format string and arguments to parse into
// said string.
func Alertf(f int, msgf string, v ...interface{}) error {
	err := fmt.Errorf(msgf, v...)
	if Lalert&l.logmodes != 0 {
		fmt.Print(prefix(f, Salert) + err.Error())
	}
	return err
}
