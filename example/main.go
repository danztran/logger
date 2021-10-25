package main

import (
	"time"

	logger "github.com/carousell/gologger"
)

var (
	// there are many ways to create a logger
	// log, err := logger.Named("main")
	// log, err := logger.New()
	// log := logger.MustNamed("main")
	log = logger.MustNew()

	// or just import and user default logger
	// import "github.com/carousell/gologger/log"
)

func main() {
	// simple log with info level
	log.Info("execution start")         // INFO "execution start"
	defer log.Infod()("execution time") // INFO "execution time: xx.xx"
	logDuration()
	logDurationWithAutoLevel()
	logDurationWithPrefixMessage("123")
}

func logDuration() {
	// log function time
	defer log.Debugd()("do something A took") // DEBUG "do something A took: 312ms"
	time.Sleep(300 * time.Millisecond)
}

func logDurationWithAutoLevel() {
	defer log.Autod(1 * time.Second)("do something B") // WARN "do something B: 2.015s"
	defer log.Autod(3 * time.Second)("do something B") // DEBUG "do something B: 2.012s"
	time.Sleep(2 * time.Second)
}

func logDurationWithPrefixMessage(userID string) {
	log := log.Withf("[user:%s]", userID)
	defer log.Warnd(50 * time.Millisecond)("do something C")  // WARN "[user:123] do something C: 200.123ms"
	defer log.Warnd(300 * time.Millisecond)("do something C") // <not logging>
	time.Sleep(200 * time.Millisecond)
}
