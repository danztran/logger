package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ReadConfig(name string) (*zap.Config, error) {
	var (
		logLevel    = getEnv("LOG_LEVEL", "debug")      // or info, warn, error, panic, fatal
		logColor    = getEnv("LOG_COLOR", "")           // or true to enable
		logEncoding = getEnv("LOG_ENCODING", "console") // json
	)

	if name != "" {
		envLevel := fmt.Sprintf("LOG_LEVEL_%s", strings.ToUpper(name))
		logLevel = getEnv(envLevel, logLevel)
	}

	// encoding
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.ConsoleSeparator = "\t"
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	if logColor == "true" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	if name == "" {
		encoderCfg.NameKey = ""
	}

	// config
	config := zap.NewProductionConfig()
	config.Encoding = strings.ToLower(logEncoding)
	config.EncoderConfig = encoderCfg
	if err := config.Level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, fmt.Errorf("error parse log level: %s / %w", logLevel, err)
	}

	return &config, nil
}

func NewZap(name string) (*zap.SugaredLogger, error) {
	config, err := ReadConfig(name)
	if err != nil {
		return nil, fmt.Errorf("error parse zap config / %w", err)
	}

	log, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("error build logger / %w", err)
	}

	log = log.
		Named(name).
		WithOptions(zap.AddCallerSkip(1))

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
