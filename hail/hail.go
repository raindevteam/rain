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
	"fmt"
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

var hlog = hail{LDefaultFlags}

// Defaults will set default settings for hail.
func Defaults() {
	hlog.logmodes = LDefaultFlags
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
	hlog.logmodes |= flags
}

// Facility takes a facility constant and returns a corresponding string.
func Facility(facility int) string {
	var fstr string

	switch facility {
	case Fbot:
		fstr = "Bot"
	case Fcore:
		fstr = "Core"
	case Fdrophand:
		fstr = "DH"
	case Feventhand:
		fstr = "EH"
	case Frain:
		fstr = "Rain"
	case Fhail:
		fstr = "HAIL"
	}

	return fstr
}

func prefix(f int, s int) string {
	return fmt.Sprintf("[%s] (%d): ", Facility(f), s)
}
