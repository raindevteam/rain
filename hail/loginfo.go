// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
//
// Contact Rodolfo at rcvallada@gmail.com for any inquires of this file.

package hail

const (
	// Fbot specifies bot logs
	Fbot = iota
	// Fdrophand specifies droplet handler logs
	Fdrophand
	// Feventhand specifies event handler logs
	Feventhand
	// Frain specifies rain logs
	Frain
	// Fhail specifies hail logs
	Fhail
)

const (
	// Semerg is the emergency level status
	Semerg = iota
	// Salert is the alert level status
	Salert
	// Scrit is the critical level status
	Scrit
	// Serr is the error level status
	Serr
	// Swarning is the warning level status
	Swarning
	// Snotice is the notice level status
	Snotice
	// Sinfo is the notice level status
	Sinfo
	// Sdebug is the debug level status
	Sdebug
)
