// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
//
// Contact Rodolfo at rcvallada@gmail.com for any inquires of this file.

package hail

const (
	// Fcore specifies core logs
	Fcore = iota
	// Fbot specifies bot logs
	Fbot
	// Fdroplet specifies droplet handler logs
	Fdroplet
	// Fhandler specifies event handler logs
	Fhandler
	// Finternal specifies internal listeners/commands logs.
	Finternal
	// Frain specifies rain logs
	Frain
	// Fhail specifies hail logs
	Fhail
)

const (
	// Sfatal is the emergency level status
	Sfatal = iota + 1
	// Scrit is the critical level status
	Scrit
	// Serr is the error level status
	Serr
	// Swarn is the warning level status
	Swarn
	// Sinfo is the notice level status
	Sinfo
	// Sdebug is the debug level status
	Sdebug
)
