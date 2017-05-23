// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

// Package hail stands for Heavily Aspired Irresponsible Logger.
package hail

import (
	"io"
	"log"
)

type hail struct {
	logmodes int
}

const (
	Lemerg = 1 << iota
	Lalert
	Lcrit
	Lerr
	Lwarning
	Lnotice
	Linfo
	Ldebug
	LDefaultFlags = Lemerg | Lalert | Lcrit | Lerr |
		Lwarning | Lnotice | Linfo | Ldebug
)

var l = hail{LDefaultFlags}

// Defaults will set default settings for hail.
func Defaults() {
	l.logmodes = Linfo | Lnotice | Lerr | Lcrit | Lalert | Lemerg
	SetLogFlags(0)
}

// SetOutput will set the internal logger's output to the given writer.
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// SetLogFlags sets the internal logger's flags. Make sure the flags you pass
// are from the standard go logger.
func SetLogFlags(flags int) {
	log.SetFlags(flags)
}

// SetFlags sets hail's logging flags.
func SetFlags(flags int) {
	l.logmodes |= flags
}

// Facility takes a facility constant and returns a corresponding string.
func Facility(facility int) string {
	var fstr string

	switch facility {
	case Fbot:
		fstr = "Bot"
	case Fdrophand:
		fstr = "DH"
	case Feventhand:
		fstr = "EH"
	case Frain:
		fstr = "Rain"
	case Fhail:
		fstr = "Hail"
	}

	return fstr
}

// Emerg is like Emergf but adds a new line. Does not accept format specifiers.
func Emerg(facility int, msg string) {
	Emergf(facility, msg+"\n")
}

// Emergf takes a facility specifier, format string and arguments to parse into
// said string.
func Emergf(facility int, msgf string, v ...interface{}) {
	if Lemerg&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, facility, Semerg, v)
	}
}

// Alert is like Alertf but adds a new line. Does not accept format specifiers.
func Alert(facility int, msg string) {
	Alertf(facility, msg+"\n")
}

// Alertf takes a facility specifier, format string and arguments to parse into
// said string.
func Alertf(facility int, msgf string, v ...interface{}) {
	if Lalert&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, v)
	}
}

// Err is like Errf but adds a new line. Does not accept format specifiers.
func Err(facility int, msg string) {
	Errf(facility, msg+"\n")
}

// Errf takes a facility specifier, format string and arguments to parse into
// said string.
func Errf(facility int, msgf string, v ...interface{}) {
	if Lerr&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, v)
	}
}

// Crit is like Critf but adds a new line. Does not accept format specifiers.
func Crit(facility int, msg string) {
	Critf(facility, msg+"\n")
}

// Critf takes a facility specifier, format string and arguments to parse into
// said string.
func Critf(facility int, msgf string, v ...interface{}) {
	if Lcrit&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, v)
	}
}

// Warning is like Warningf but adds a new line. Does not accept format specifiers.
func Warning(facility int, msg string) {
	Warningf(facility, msg+"\n")
}

// Warningf takes a facility specifier, format string and arguments to parse into
// said string.
func Warningf(facility int, msgf string, v ...interface{}) {
	if Lwarning&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, v)
	}
}

// Notice is like Noticef but adds a new line. Does not accept format specifiers.
func Notice(facility int, msg string) {
	Noticef(facility, msg+"\n")
}

// Noticef takes a facility specifier, format string and arguments to parse into
// said string.
func Noticef(facility int, msgf string, v ...interface{}) {
	if Lnotice&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, v)
	}
}

// Info is like Infof but adds a new line. Does not accept format specifiers.
func Info(facility int, msg string) {
	Infof(facility, msg+"\n")
}

// Infof takes a facility specifier, format string and arguments to parse into
// said string.
func Infof(facility int, msgf string, v ...interface{}) {
	if Linfo&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, v)
	}
}

// Debug is like Debugf but adds a new line. Does not accept format specifiers.
func Debug(facility int, msg string) {
	Debugf(facility, msg+"\n")
}

// Debugf takes a facility specifier, format string and arguments to parse into
// said string.
func Debugf(facility int, msgf string, v ...interface{}) {
	if Ldebug&l.logmodes != 0 {
		log.Printf("[%s] (%d): "+msgf, facility, Sdebug, v)
	}
}
