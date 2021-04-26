package customLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/techievee/frankieone/customLogger/debugcore"
)

type options struct {
	env          string
	isProduction bool
	fields       []zapcore.Field
}

// Option overrides behavior of Logger.
type Option interface {
	apply(*options)
}

// define a func type to use it faster in function operations
// inspired from https://github.com/uber-go/guide/blob/master/style.md#functional-options
type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// WithServiceName defines optional service name.
func WithServiceName(s string) Option {
	return optionFunc(func(o *options) {
		o.fields = append(o.fields, zap.String("service", s))
	})
}

// WithExtraFields add extra fields to logger.
func WithExtraFields(fields ...zap.Field) Option {
	return optionFunc(func(o *options) {
		o.fields = append(o.fields, fields...)
	})
}

// WithIsProduction defines whether it is prod env.
func WithIsProduction(b bool) Option {
	return optionFunc(func(o *options) {
		o.isProduction = b
	})
}

// loggerImpl implements debugcore.Logger interface.
type loggerImpl struct {
	sugar *zap.SugaredLogger
}

// NewLogger returns a new logger with customized options.
func NewLogger(env string, opts ...Option) debugcore.Logger {
	configOpt := options{env: env}
	configOpt.fields = append(configOpt.fields, zap.String("environment", env))
	for _, o := range opts {
		o.apply(&configOpt)
	}

	loggerConfig := debugConfig
	if configOpt.isProduction || configOpt.env == envProd {
		loggerConfig = prodConfig
	}

	l, _ := loggerConfig.Build()
	return &loggerImpl{
		sugar: l.With(configOpt.fields...).Sugar(),
	}
}

// Info ...
func (l *loggerImpl) Info(msg string, keysAndValues ...interface{}) {
	defer l.sugar.Sync()
	l.sugar.Infow(msg, keysAndValues...)
}

// Warn ...
func (l *loggerImpl) Warn(msg string, keysAndValues ...interface{}) {
	defer l.sugar.Sync()
	l.sugar.Warnw(msg, keysAndValues...)
}

// Debug ...
func (l *loggerImpl) Debug(msg string, keysAndValues ...interface{}) {
	defer l.sugar.Sync()
	l.sugar.Debugw(msg, keysAndValues...)
}

// Error ...
func (l *loggerImpl) Error(msg string, keysAndValues ...interface{}) {
	defer l.sugar.Sync()
	l.sugar.Errorw(msg, keysAndValues...)
}
