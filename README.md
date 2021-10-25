# logger

![actions](https://github.com/carousell/gologger/actions/workflows/main.yml/badge.svg)
[![travis](https://app.travis-ci.com/carousell/gologger.svg?branch=main)](https://app.travis-ci.com/carousell/gologger)

Package logger is a wrapper of structured Zap Logger with extended features.

## Features

- Config logger via env.
- Name key is optional, add custom prefix message.
- Add log duration functions (Infod, Debugd, Warnd, Autod).
- Add default logger instance in package `log`.
- Share common logger across repositories.
- Simple but fully log format information.
- More easier to use with very basic init functions.

## Quick start

### Install

```bash
go get -d github.com/carousell/gologger
```

### Options
```bash
# log levels: debug (default), info, warn, error, fatal, panic
# set default log level
LOG_LEVEL="debug"
# set log level for named logger
LOG_LEVEL_$NAME="debug"

# log color: false (default), true
# colorize on level field 
LOG_COLOR="true"

# log encoding: console (default), json
#   console: simple content (easy to read)
#   json: structured json logs (easy to search)
LOG_ENCODING="console"
```

### Example

```go
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

	// or just import and use default logger
	// import "github.com/carousell/gologger/log"
)

func main() {
	// simple log with info level
	log.Info("start") // INFO "start"
	
	// log duration
	defer log.Infod()("execution time") // INFO "execution time: xx.xx"
	logDuration()
	logDurationWithAutoLevel()
	logDurationWithPrefixMessage("123")
}

func logDuration() {
	// log duration
	defer log.Debugd()("do A took") // DEBUG "do A took: 312ms"
	time.Sleep(300 * time.Millisecond)
}

func logDurationWithAutoLevel() {
        // Autod log with warn level if duration is longer than expected, otherwise log with debug level
	// log with warn level because duration (2s) is longer than expected (1s)
	defer log.Autod(1 * time.Second)("do B") // WARN "do B: 2.015s"
	
	// log with debug level because duration (2s) is shorter than expected (3s)
	defer log.Autod(3 * time.Second)("do B") // DEBUG "do B: 2.012s"
	
	time.Sleep(2 * time.Second)
}

func logDurationWithPrefixMessage(userID string) {
        // make new logger with prefix message
	log := log.Withf("[user:%s]", userID)
	
	// Warnd log warn if duration tooks longer than epected, otherwise nothing will be logged
	// log as warn because duration (200ms) is longer than expected (50ms)
	defer log.Warnd(50 * time.Millisecond)("do C")  // WARN "[user:123] do C: 200.123ms"
	
	// not logging because duration (200ms) is shorter than expected (300ms)
	defer log.Warnd(300 * time.Millisecond)("do C")
	
	time.Sleep(200 * time.Millisecond)
}
```
