package debugcore

// Logger is interface for logging of debug package.
type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

// NoOpsLogger is default logger with no operation
type NoOpsLogger struct{}

// Info ...
func (l *NoOpsLogger) Info(_ string, _ ...interface{}) {
	// do nothing
}

// Warn ...
func (l *NoOpsLogger) Warn(_ string, _ ...interface{}) {
	// do nothing
}

// Debug ...
func (l *NoOpsLogger) Debug(_ string, _ ...interface{}) {
	// do nothing
}

// Error ...
func (l *NoOpsLogger) Error(_ string, _ ...interface{}) {
	// do nothing
}
