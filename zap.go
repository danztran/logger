package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(module string) (*zap.SugaredLogger, error) {
	var (
		logLevel    = getEnv("LOG_LEVEL", "debug")      // or info, warn, error, panic, fatal
		logColor    = getEnv("LOG_COLOR", "")           // or true to enable
		logEncoding = getEnv("LOG_ENCODING", "console") // json
		config      = zap.NewProductionConfig()
	)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.ConsoleSeparator = "\t"
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	if logColor == "true" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.Encoding = strings.ToLower(logEncoding)
	config.EncoderConfig = encoderCfg

	if module == "" {
		encoderCfg.NameKey = ""
	} else {
		logLevel = getEnv(fmt.Sprintf("LOG_LEVEL_%s", strings.ToUpper(module)), logLevel)
	}

	if err := config.Level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, fmt.Errorf("error parse log level: %s / %w", logLevel, err)
	}
	log, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("error build logger / %w", err)
	}

	log = log.WithOptions(zap.AddCallerSkip(1))
	if module != "" {
		log = log.Named(module)
	}

	return log.Sugar(), nil
}

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
