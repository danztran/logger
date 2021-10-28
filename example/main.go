package main

import (
	"time"

	logger "github.com/carousell/gologger"
)

var (
	// there are many ways to create a logger
	// log, err = logger.Named("main")
	// log, err = logger.New()
	// log = logger.MustNamed("main")
	log = logger.MustNew()

	// or just import and use default logger
	// import "github.com/carousell/gologger/log"
)

func main() {
	// simple log with info level
	log.Info("start")                   // INFO "start"
	defer log.Infod()("execution time") // INFO "execution time: 2.016s"

	doSomething()
}

func doSomething() {
	defer log.Debugd()("do A took") // DEBUG "do A took: 2.0011446s"

	// Autod log with warn level if duration is longer than expected,
	// otherwise log with debug level.

	// log with warn level because duration (2s) is longer than expected (1s)
	defer log.Autod(1 * time.Second)("do B") // WARN "do B: 2.015s"

	// log with debug level because duration (2s) is shorter than expected (3s)
	defer log.Autod(3 * time.Second)("do B") // DEBUG "do B: 2.012s"

	// make new logger with prefix message
	userID := 123
	log := log.Withf("[user:%d]", userID)

	// Warnd log warn if duration took longer than expected,
	// otherwise nothing will be logged.

	// log as warn because duration (2s) is longer than expected (1s)
	defer log.Warnd(1 * time.Second)("do C") // WARN "[user:123] do C: 2.011s"

	// not logging because duration (2s) is shorter than expected (3s)
	defer log.Warnd(3 * time.Second)("do C")

	time.Sleep(2 * time.Second)
}
