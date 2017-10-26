package hail

import (
	"fmt"
)

// Fatal is like Emergf but adds a new line. Does not accept format specifiers.
func Fatal(facility int, msg string) error {
	return Fatalf(facility, msg+"\n")
}

// Fatalf takes a facility specifier, format string and arguments to parse into
// said string. It returns an error with msgf as the error string.
func Fatalf(f int, msgf string, v ...interface{}) error {
	err := fmt.Errorf(msgf, v...)
	if Lfatal&hlog.logmodes != 0 {
		fmt.Print(prefix(f, Sfatal) + err.Error())
	}
	return err
}

// Crit is like Critf but adds a new line. Does not accept format specifiers.
func Crit(facility int, msg string) error {
	return Critf(facility, msg+"\n")
}

// Critf takes a facility specifier, format string and arguments to parse into
// said string.
func Critf(f int, msgf string, v ...interface{}) error {
	err := fmt.Errorf(msgf, v...)
	if Lcrit&hlog.logmodes != 0 {
		fmt.Print(prefix(f, Scrit) + err.Error())
	}
	return err
}

// Err is like Errf but adds a new line. Does not accept format specifiers.
func Err(facility int, msg string) error {
	return Errf(facility, msg+"\n")
}

// Errf takes a facility specifier, format string and arguments to parse into
// said string.
func Errf(f int, msgf string, v ...interface{}) error {
	err := fmt.Errorf(msgf, v...)
	if Lerr&hlog.logmodes != 0 {
		fmt.Print(prefix(f, Serr) + err.Error())
	}
	return err
}

// Warn is like Warningf but adds a new line. Does not accept format specifiers.
func Warn(facility int, msg string) {
	Warnf(facility, msg+"\n")
}

// Warnf takes a facility specifier, format string and arguments to parse into
// said string.
func Warnf(f int, msgf string, v ...interface{}) {
	if Lwarn&hlog.logmodes != 0 {
		fmt.Printf(prefix(f, Swarn) + fmt.Sprintf(msgf, v...))
	}
}

// Info is like Infof but adds a new line. Does not accept format specifiers.
func Info(facility int, msg string) {
	Infof(facility, msg+"\n")
}

// Infof takes a facility specifier, format string and arguments to parse into
// said string.
func Infof(f int, msgf string, v ...interface{}) {
	if Linfo&hlog.logmodes != 0 {
		fmt.Printf(prefix(f, Sinfo) + fmt.Sprintf(msgf, v...))
	}
}

// Debug is like Debugf but adds a new line. Does not accept format specifiers.
func Debug(facility int, msg string) {
	Debugf(facility, msg+"\n")
}

// Debugf takes a facility specifier, format string and arguments to parse into
// said string.
func Debugf(f int, msgf string, v ...interface{}) {
	if Ldebug&hlog.logmodes != 0 {
		fmt.Printf(prefix(f, Sdebug) + fmt.Sprintf(msgf, v...))
	}
}
