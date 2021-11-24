package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapConfig read log config from environment
// default value will be used if env is invalid
func NewZapConfig(name string) *zap.Config {
	var (
		logLevel     = getEnv("LOG_LEVEL", "debug")       // debug, info, warn, error, panic, fatal
		logColor     = getEnv("LOG_COLOR", "")            // true to enable
		logEncoding  = getEnv("LOG_ENCODING", "console")  // console, json
		logTimestamp = getEnv("LOG_TIMESTAMP", "rfc3339") // rfc3339, rfc3339nano, iso8601, s, ms, ns, disabled
		logSeparator = getEnv("LOG_SEPARATOR", "\t")
	)

	if name != "" {
		envLevel := fmt.Sprintf("LOG_LEVEL_%s", strings.ToUpper(name))
		logLevel = getEnv(envLevel, logLevel)
	}

	// encoding
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.ConsoleSeparator = logSeparator
	encoderCfg.EncodeTime = parseTimeEncoding(logTimestamp)
	encoderCfg.TimeKey = parseTimeKey(logTimestamp)
	encoderCfg.EncodeLevel = parseLogEncoder(logColor)
	encoderCfg.NameKey = parseNameKey(name)

	// config
	config := zap.NewProductionConfig()
	config.Encoding = parseLogEncoding(logEncoding)
	config.Level = zap.NewAtomicLevelAt(parseLogLevel(logLevel))
	config.EncoderConfig = encoderCfg

	return &config
}

// NewZap returns core logger as sugared logger with
// default basic config
func NewZap(name string) (*zap.SugaredLogger, error) {
	config := NewZapConfig(name)
	log, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("error build logger / %w", err)
	}

	log = log.Named(name)
	return log.Sugar(), nil
}

func parseTimeEncoding(key string) zapcore.TimeEncoder {
	switch key {
	case "s":
		return UnixTimeEncoder
	case "ms":
		return UnixMilliTimeEncoder
	case "ns":
		return zapcore.EpochNanosTimeEncoder
	case "rfc3339nano":
		return zapcore.RFC3339NanoTimeEncoder
	case "rfc3339":
		return zapcore.RFC3339TimeEncoder
	case "iso8601":
		return zapcore.ISO8601TimeEncoder
	case "disabled":
		return func(time.Time, zapcore.PrimitiveArrayEncoder) {}
	default:
		return zapcore.RFC3339TimeEncoder
	}
}

func parseTimeKey(key string) string {
	if key == "disabled" {
		return ""
	}
	return "ts"
}

func parseLogEncoder(key string) zapcore.LevelEncoder {
	if key == "true" {
		return zapcore.LowercaseColorLevelEncoder
	}
	return zapcore.LowercaseLevelEncoder
}

func parseNameKey(key string) string {
	if key != "" {
		return "name"
	}
	return ""
}

func parseLogLevel(key string) zapcore.Level {
	switch strings.ToLower(key) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func parseLogEncoding(key string) string {
	if key == "json" {
		return key
	}
	return "console"
}

func UnixTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.Unix())
}

func UnixMilliTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixMilli())
}

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
