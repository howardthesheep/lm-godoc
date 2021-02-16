package main

import (
	"fmt"

	"./lmgo/log"
)

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

type Logger struct{}

func (Logger) Emergf(format string, args ...interface{}) {
	lmlog(log.LevelEmerg,format, args...)
}

func (Logger) Critf(format string, args ...interface{}) {
	lmlog(log.LevelCrit,format, args...)
}

func (Logger) Errorf(format string, args ...interface{}) {
	lmlog(log.LevelError,format, args...)
}

func (Logger) Alertf(format string, args ...interface{}) {
	lmlog(log.LevelAlert,format, args...)
}

func (Logger) Warnf(format string, args ...interface{}) {
	lmlog(log.LevelWarning,format, args...)
}

func (Logger) Noticef(format string, args ...interface{}) {
	lmlog(log.LevelNotice,format, args...)
}

func (Logger) Infof(format string, args ...interface{}) {
	lmlog(log.LevelInfo, format, args...)
}

func (Logger) Debugf(format string, args ...interface{}) {
	lmlog(log.LevelDebug, format, args...)
}

func lmlog(level uint, format string, args ...interface{}) {

	prefixColor := ""

	switch level {
	case log.LevelDebug:
		prefixColor = ""
		break
	case log.LevelInfo:
		prefixColor = White
		break
	case log.LevelNotice:
		prefixColor = White
		break
	case log.LevelWarning:
		prefixColor = Yellow
		break
	case log.LevelAlert:
		prefixColor = Yellow
		break
	case log.LevelError:
		prefixColor = Red
		break
	case log.LevelCrit:
		prefixColor = Red
		break
	case log.LevelEmerg:
		prefixColor = Red
		break
	default:
		fmt.Printf("%s", "Invalid Log Level Provided")
	}


	fmt.Printf(prefixColor + format + Reset + "\n", args...)
}


