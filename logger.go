package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Logger wrapper the base zap SugaredLogger for some extra features
type Logger struct {
	prefix string
	core   *zap.SugaredLogger
}

// MustNew is like New but panics if the logger can not be built
func MustNew() *Logger {
	return MustNamed("")
}

// MustNamed is like Named but panics if the logger can not be built
func MustNamed(name string) *Logger {
	log, err := Named(name)
	if err != nil {
		panic(err)
	}

	return log
}

// New register a wrapped zap logger with no name
func New() (*Logger, error) {
	return Named("")
}

// Named build a simple wrapped zap logger
func Named(name string) (*Logger, error) {
	logger, err := NewZap(name)
	if err != nil {
		return nil, err
	}

	log := &Logger{
		prefix: "",
		core:   logger,
	}

	return log, nil
}

// Wrap wrap zap logger with pkg logger
func Wrap(logger *zap.SugaredLogger) *Logger {
	log := &Logger{
		prefix: "",
		core:   logger,
	}

	return log
}

// Core returns core logger instance, so you can log directly. (for performance concern)
func (c *Logger) Core() *zap.SugaredLogger {
	return c.core
}

// Skip return new instances that increases the number of callers skipped by caller annotation
func (c *Logger) Skip(skip int) *Logger {
	logger := c.core.Desugar().WithOptions(zap.AddCallerSkip(skip)).Sugar()
	return &Logger{
		prefix: c.prefix,
		core:   logger.Named(""),
	}
}

// With makes new instance and uses fmt.Sprint to make prefix for later logs
func (c *Logger) With(args ...interface{}) *Logger {
	message := fmt.Sprint(args...)
	return &Logger{
		prefix: c.prefix + message + " ",
		core:   c.core.Named(""),
	}
}

// Withf makes new instance and uses fmt.Sprintf to make prefix for later logs
func (c *Logger) Withf(template string, args ...interface{}) *Logger {
	message := fmt.Sprintf(template, args...)
	return &Logger{
		prefix: c.prefix + message + " ",
		core:   c.core.Named(""),
	}
}

// Withw makes new instance included additional context keys and values for later logs (prefer for json logs).
func (c *Logger) Withw(keysAndValues ...interface{}) *Logger {
	return &Logger{
		prefix: c.prefix,
		core:   c.core.With(keysAndValues),
	}
}

// Debug uses fmt.Sprint to construct and log a message with prefix.
func (c *Logger) Debug(args ...interface{}) {
	msg := c.prefix + fmt.Sprint(args...)
	c.core.Debug(msg)
}

// Debugf uses fmt.Sprintf to log a templated message with prefix.
func (c *Logger) Debugf(template string, args ...interface{}) {
	c.core.Debugf(c.prefix+template, args...)
}

// Debugw logs a message with some additional context with keys and values (prefer for json logs).
func (c *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	c.core.Debugw(c.prefix+msg, keysAndValues...)
}

// Info uses fmt.Sprint to construct and log a message with prefix.
func (c *Logger) Info(args ...interface{}) {
	msg := c.prefix + fmt.Sprint(args...)
	c.core.Info(msg)
}

// Infof uses fmt.Sprintf to log a templated message with prefix.
func (c *Logger) Infof(template string, args ...interface{}) {
	c.core.Infof(c.prefix+template, args...)
}

// Infow logs a message with some additional context with keys and values (prefer for json logs).
func (c *Logger) Infow(msg string, keysAndValues ...interface{}) {
	c.core.Infow(c.prefix+msg, keysAndValues...)
}

// Warn uses fmt.Sprint to construct and log a message with prefix.
func (c *Logger) Warn(args ...interface{}) {
	msg := c.prefix + fmt.Sprint(args...)
	c.core.Warn(msg)
}

// Warnf uses fmt.Sprintf to log a templated message with prefix.
func (c *Logger) Warnf(template string, args ...interface{}) {
	c.core.Warnf(c.prefix+template, args...)
}

// Warnw logs a message with some additional context with keys and values (prefer for json logs).
func (c *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	c.core.Warnw(c.prefix+msg, keysAndValues...)
}

// Error uses fmt.Sprint to construct and log a message with prefix.
func (c *Logger) Error(args ...interface{}) {
	msg := c.prefix + fmt.Sprint(args...)
	c.core.Error(msg)
}

// Errorf uses fmt.Sprintf to log a templated message with prefix.
func (c *Logger) Errorf(template string, args ...interface{}) {
	c.core.Errorf(c.prefix+template, args...)
}

// Errorw logs a message with some additional context with keys and values (prefer for json logs).
func (c *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	c.core.Errorw(c.prefix+msg, keysAndValues...)
}

// Panic uses fmt.Sprint to construct and log a message with prefix.
func (c *Logger) Panic(args ...interface{}) {
	msg := c.prefix + fmt.Sprint(args...)
	c.core.Panic(msg)
}

// Panicf uses fmt.Sprintf to log a templated message with prefix.
func (c *Logger) Panicf(template string, args ...interface{}) {
	c.core.Panicf(c.prefix+template, args...)
}

// Panicw logs a message with some additional context with keys and values (prefer for json logs).
func (c *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	c.core.Panicw(c.prefix+msg, keysAndValues...)
}

// Fatal uses fmt.Sprint to construct and log a message with prefix.
func (c *Logger) Fatal(args ...interface{}) {
	msg := c.prefix + fmt.Sprint(args...)
	c.core.Fatal(msg)
}

// Fatalf uses fmt.Sprintf to log a templated message with prefix.
func (c *Logger) Fatalf(template string, args ...interface{}) {
	c.core.Fatalf(c.prefix+template, args...)
}

// Fatalw logs a message with some additional context with keys and values (prefer for json logs).
func (c *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	c.core.Fatalw(c.prefix+msg, keysAndValues...)
}

// Infod is like Infof but return a function to log message with duration.
func (c *Logger) Infod() func(template string, args ...interface{}) {
	ts := time.Now()
	return func(template string, args ...interface{}) {
		msg := fmt.Sprintf(c.prefix+template, args...)
		c.core.Infof("%s: %s", msg, time.Since(ts))
	}
}

// Warnd is like Warnf but return a function to log message with duration only if the duration is longer than expected
func (c *Logger) Warnd(dur time.Duration) func(template string, args ...interface{}) {
	ts := time.Now()
	return func(template string, args ...interface{}) {
		el := time.Since(ts)
		if el > dur {
			msg := fmt.Sprintf(c.prefix+template, args...)
			c.core.Warnf("%s: %s", msg, el)
		}
	}
}

// Debugd is like Debugf but return a function to log message with duration.
func (c *Logger) Debugd() func(template string, args ...interface{}) {
	ts := time.Now()
	return func(template string, args ...interface{}) {
		msg := fmt.Sprintf(c.prefix+template, args...)
		c.core.Debugf("%s: %s", msg, time.Since(ts))
	}
}

// Autod return a function to log message with elapsed duration, log level could be debug or warn depended on duration time
func (c *Logger) Autod(dur time.Duration) func(template string, args ...interface{}) {
	ts := time.Now()
	return func(template string, args ...interface{}) {
		el := time.Since(ts)
		msg := fmt.Sprintf(c.prefix+template, args...)
		if el > dur {
			c.core.Warnf("%s: %s", msg, el)
		} else {
			c.core.Debugf("%s: %s", msg, el)
		}
	}
}
