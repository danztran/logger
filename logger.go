package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	prefix string
	logger *zap.SugaredLogger
}

// MustNew is like New but panics the logger can not be built
func MustNew(args ...string) *Logger {
	log, err := New(args...)
	if err != nil {
		panic(err)
	}

	return log
}

// New build a simple wrapped zap logger
func New(args ...string) (*Logger, error) {
	module := ""
	if len(args) > 0 {
		module = args[0]
	}

	logger, err := newZapLogger(module)
	if err != nil {
		return nil, err
	}

	log := &Logger{
		prefix: "",
		logger: logger,
	}

	return log, nil
}

func (c *Logger) newWith(submsg string) *Logger {
	log := &Logger{
		prefix: submsg,
		logger: c.logger.Named(""),
	}

	return log
}

// Withf makes new instance and uses fmt.Sprintf to make prefix for next log messages
func (c *Logger) Withf(template string, args ...interface{}) *Logger {
	message := fmt.Sprintf(template, args...)
	return c.newWith(c.prefix + message + " ")
}

// New makes new instance with blank prefix message.
func (c *Logger) New() *Logger {
	return c.newWith("")
}

// Debugf uses fmt.Sprintf to log a templated message with prefix message.
func (c *Logger) Debugf(template string, args ...interface{}) {
	c.logger.Debugf(c.prefix+template, args...)
}

// Infof uses fmt.Sprintf to log a templated message with prefix message.
func (c *Logger) Infof(template string, args ...interface{}) {
	c.logger.Infof(c.prefix+template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message with prefix message.
func (c *Logger) Warnf(template string, args ...interface{}) {
	c.logger.Warnf(c.prefix+template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message with prefix message.
func (c *Logger) Errorf(template string, args ...interface{}) {
	c.logger.Errorf(c.prefix+template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message with prefix message.
func (c *Logger) Panicf(template string, args ...interface{}) {
	c.logger.Panicf(c.prefix+template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message with prefix message.
func (c *Logger) Fatalf(template string, args ...interface{}) {
	c.logger.Fatalf(c.prefix+template, args...)
}

// Infod is like Infof but return a function to log message with duration.
func (c *Logger) Infod(template string, args ...interface{}) func() {
	ts := time.Now()
	return func() {
		msg := fmt.Sprintf(c.prefix+template, args...)
		c.logger.Infof("%s: %s", msg, time.Since(ts))
	}
}

// Debugd is like Debugf but return a function to log message with duration.
func (c *Logger) Debugd(template string, args ...interface{}) func() {
	ts := time.Now()
	return func() {
		msg := fmt.Sprintf(c.prefix+template, args...)
		c.logger.Debugf("%s: %s", msg, time.Since(ts))
	}
}

// Skip return new instances that increases the number of callers skipped by caller annotation
func (c *Logger) Skip(skip int) *Logger {
	logger := c.logger.Desugar().WithOptions(zap.AddCallerSkip(skip)).Sugar()
	return &Logger{
		prefix: c.prefix,
		logger: logger,
	}
}
