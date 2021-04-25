package frankieoneLogger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const envProd = "prod"

var (
	debugConfig zap.Config
	prodConfig  zap.Config
	config      zap.Config
	sugar       *zap.SugaredLogger
)

func init() {
	debugConfig = NewConfig()

	prodConfig = NewConfig()
	prodConfig.Development = false
	prodConfig.DisableStacktrace = true
	prodConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	if os.Getenv("APP_ENV") != envProd {
		config = debugConfig
	} else {
		config = prodConfig
	}

	logger, _ := config.Build()
	sugar = logger.Sugar()
}

// NewEncoderConfig returns an opinionated EncoderConfig
func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "",
		CallerKey:      "",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// NewConfig is a reasonable development logging configuration.
// Logging is enabled at DebugLevel and above.
//
// It enables development mode (which makes DPanicLevel logs panic), uses a
// console encoder, writes to standard error, and disables sampling.
// Stacktraces are automatically included on logs of WarnLevel and above.
func NewConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		DisableCaller:    true,
		Encoding:         "json",
		EncoderConfig:    NewEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// Info will accept custom argument pairs that you wish to xeroLog.
// It will produce a json-encoded string.
// See: https://godoc.org/go.uber.org/zap
func Info(msg string, keysAndValues ...interface{}) {
	defer sugar.Sync()
	sugar.Infow(msg, keysAndValues...)
}

// Warn will accept custom argument pairs that you wish to xeroLog.
// It will produce a json-encoded string.
// See: https://godoc.org/go.uber.org/zap
func Warn(msg string, keysAndValues ...interface{}) {
	defer sugar.Sync()
	sugar.Warnw(msg, keysAndValues...)
}

// Debug will accept custom argument pairs that you wish to xeroLog.
// It will produce a json-encoded string.
// See: https://godoc.org/go.uber.org/zap
func Debug(msg string, keysAndValues ...interface{}) {
	defer sugar.Sync()
	sugar.Debugw(msg, keysAndValues...)
}

// Error will accept custom argument pairs that you wish to xeroLog.
// It will produce a json-encoded string.
// See: https://godoc.org/go.uber.org/zap
func Error(msg string, keysAndValues ...interface{}) {
	defer sugar.Sync()
	sugar.Errorw(msg, keysAndValues...)
}
