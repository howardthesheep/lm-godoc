package main

import (
	"os"
)

var lmlogger = Logger{}
func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		lmlogger.Debugf("%s", "You need more args")
		lmlogger.Warnf("%s", "You need more args")
		lmlogger.Errorf("%s", "You need more args")
	}
}