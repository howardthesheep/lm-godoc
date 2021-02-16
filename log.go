package main

import (
	"fmt"
)

// Logging Color Constants
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// Logging Level Constants
const (
	LevelEmerg uint = iota
	LevelAlert
	LevelCrit
	LevelError
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug
)

// Our Logging struct which implements allmanf
type Logger struct{}


/******* Allmanf Method Implementations *******/

func (Logger) Emergf(format string, args ...interface{}) {
	lmlog(LevelEmerg,format, args...)
}

func (Logger) Critf(format string, args ...interface{}) {
	lmlog(LevelCrit,format, args...)
}

func (Logger) Errorf(format string, args ...interface{}) {
	lmlog(LevelError,format, args...)
}

func (Logger) Alertf(format string, args ...interface{}) {
	lmlog(LevelAlert,format, args...)
}

func (Logger) Warnf(format string, args ...interface{}) {
	lmlog(LevelWarning,format, args...)
}

func (Logger) Noticef(format string, args ...interface{}) {
	lmlog(LevelNotice,format, args...)
}

func (Logger) Infof(format string, args ...interface{}) {
	lmlog(LevelInfo, format, args...)
}

func (Logger) Debugf(format string, args ...interface{}) {
	lmlog(LevelDebug, format, args...)
}


/******* Helper Methods *******/

// lmlog prints a log to the console using fmt.Printf
// It formats the output with a predefined color based on
// level provided
func lmlog(level uint, format string, args ...interface{}) {

	prefixColor := ""

	switch level {
	case LevelDebug:
		prefixColor = ""
		break
	case LevelInfo:
		prefixColor = White
		break
	case LevelNotice:
		prefixColor = White
		break
	case LevelWarning:
		prefixColor = Yellow
		break
	case LevelAlert:
		prefixColor = Yellow
		break
	case LevelError:
		prefixColor = Red
		break
	case LevelCrit:
		prefixColor = Red
		break
	case LevelEmerg:
		prefixColor = Red
		break
	default:
		fmt.Printf("%s", "Invalid Log Level Provided")
	}


	fmt.Printf(prefixColor + format + Reset + "\n", args...)
}


