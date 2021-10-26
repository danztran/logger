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
- Easy to start with many simple ways to create a basic logger.
- Easy to switch between logger and sugared logger with `Wrap` and `Unwrap` method.

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
# set log level for named logger (with uppercase),
# for example: LOG_LEVEL_MAIN, LOG_LEVEL_DB
LOG_LEVEL_MAIN="debug"

# log color: false (default), true
# colorize on level field
LOG_COLOR="true"

# log encoding: console (default), json
#   console: simple content (easy to read)
#   json: structured json logs (easy to search)
LOG_ENCODING="console"

# timestamp format: rfc3339 (default), rfc3339nano, iso8601, s, ms, ns, disabled
LOG_TIMESTAMP="rfc3339"
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
	defer log.Infod()("execution time") // INFO "execution time: xx.xx"

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
```

```bash
$ LOG_COLOR=true go run example/main.go
2021-10-26T18:00:30.792Z    info    example/main.go:22      start
2021-10-26T18:00:32.794Z    warn    example/main.go:54      [user:123] do C: 2.0010509s
2021-10-26T18:00:32.794Z    debug   example/main.go:54      do B: 2.0011264s
2021-10-26T18:00:32.794Z    warn    example/main.go:54      do B: 2.0011339s
2021-10-26T18:00:32.794Z    debug   example/main.go:54      do A took: 2.0011446s
2021-10-26T18:00:32.794Z    info    example/main.go:26      execution time: 2.0011492s

$ LOG_TIMESTAMP=ts LOG_ENCODING=json go run example/main.go
{"level":"info","ts":1635257396952,"caller":"example/main.go:22","msg":"start"}
{"level":"warn","ts":1635257398953,"caller":"example/main.go:54","msg":"[user:123] do C: 2.0010486s"}
{"level":"debug","ts":1635257398953,"caller":"example/main.go:54","msg":"do B: 2.0011176s"}
{"level":"warn","ts":1635257398953,"caller":"example/main.go:54","msg":"do B: 2.0011231s"}
{"level":"debug","ts":1635257398953,"caller":"example/main.go:54","msg":"do A took: 2.0011263s"}
{"level":"info","ts":1635257398953,"caller":"example/main.go:26","msg":"execution time: 2.0011293s"}
```
