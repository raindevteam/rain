package rlog

import (
	"io"
	"log"
)

type logger struct {
	logmodes int
}

const (
	Ldebug = 1 << iota
	Linfo
	Lwarn
	Lerror
	LDefaultFlags = Linfo | Lwarn | Lerror | Ldebug // initial values for the standard logger
)

var l = logger{LDefaultFlags}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func SetLogFlags(flags int) {
	log.SetFlags(flags)
}

func SetFlags(flags int) {
	l.logmodes = flags
}

func prefix(who string, level int) string {
	var logmode string

	switch level {
	case Ldebug:
		logmode = "debug"
	case Linfo:
		logmode = "info"
	case Lwarn:
		logmode = "warn"
	case Lerror:
		logmode = "error"
	}

	return " " + who + " (" + logmode + ") "
}

func Println(v ...interface{}) {
	log.Println(v...)
}

func Debug(who string, msg string) {
	if Ldebug&l.logmodes != 0 {
		log.Println(prefix(who, Ldebug) + msg)
	}
}

func Info(who string, msg string) {
	if Linfo&l.logmodes != 0 {
		log.Println(prefix(who, Linfo) + msg)
	}
}

func Warn(who string, msg string) {
	if Lwarn&l.logmodes != 0 {
		log.Println(prefix(who, Lwarn) + msg)
	}
}

func Error(who string, msg string) {
	Errorf(who, msg)
}

func Errorf(who string, format string, v ...interface{}) {
	if Lerror&l.logmodes != 0 {
		log.Printf(prefix(who, Lerror)+format, v...)
	}
}
