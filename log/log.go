// package log is the expose default logger instance
package log

import "github.com/danztran/logger"

var (
	log     = logger.MustNew()
	Core    = log.Unwrap
	Skip    = log.Skip
	AddHook = log.AddHook
	Sync    = log.Sync

	With  = log.With
	Withf = log.Withf
	Withw = log.Withw

	Debug  = log.Debug
	Debugw = log.Debugw
	Debugf = log.Debugf
	Debugd = log.Debugd

	Info  = log.Info
	Infof = log.Infof
	Infow = log.Infow
	Infod = log.Infod

	Warn  = log.Warn
	Warnf = log.Warnf
	Warnw = log.Warnw
	Warnd = log.Warnd

	Error  = log.Error
	Errorf = log.Errorf
	Errorw = log.Errorw

	Panic  = log.Panic
	Panicf = log.Panicf
	Panicw = log.Panicw

	Fatal  = log.Fatal
	Fatalf = log.Fatalf
	Fatalw = log.Fatalw

	Autod = log.Autod
)
