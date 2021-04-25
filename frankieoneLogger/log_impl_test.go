package frankieoneLogger

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger("dev", WithServiceName("xero-api"), WithExtraFields(zap.Int("cpu", 30)))
	testErr := errors.New("service error")
	assert.NotPanics(t, func() {
		l.Info("there is info", "error", testErr)
		l.Warn("there is warning", "error", testErr)
		l.Debug("there is debug", "error", testErr)
		l.Error("there is error", "error", testErr)
	})

	assert.Panics(t, func() {
		l.Info("there is panic", testErr)
	})

	assert.Panics(t, func() {
		l.Warn("there is panic", testErr)
	})

	assert.Panics(t, func() {
		l.Debug("there is panic", testErr)
	})

	assert.Panics(t, func() {
		l.Error("there is panic", testErr)
	})
}
