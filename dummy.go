package logger

type dummyLogger struct {
}

// NewDummyLogger create new dummy logger
func NewDummyLogger() Logger {
	return &dummyLogger{}
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *dummyLogger) Debug(_ ...interface{}) {
}

// Info uses fmt.Sprint to construct and log a message.
func (l *dummyLogger) Info(_ ...interface{}) {
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *dummyLogger) Warn(_ ...interface{}) {
}

// Error uses fmt.Sprint to construct and log a message.
func (l *dummyLogger) Error(_ ...interface{}) {
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *dummyLogger) Panic(_ ...interface{}) {
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *dummyLogger) Fatal(_ ...interface{}) {
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *dummyLogger) Debugf(_ string, _ ...interface{}) {
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *dummyLogger) Infof(_ string, _ ...interface{}) {
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *dummyLogger) Warnf(_ string, _ ...interface{}) {
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *dummyLogger) Errorf(_ string, _ ...interface{}) {
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *dummyLogger) Panicf(_ string, _ ...interface{}) {
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *dummyLogger) Fatalf(_ string, _ ...interface{}) {
}

// WithFields adds fields to the logging context
func (l *dummyLogger) WithFields(_ Fields) Logger {
	return l
}

// WithField add field to the logging context
func (l *dummyLogger) WithField(_ string, _ any) Logger {
	return l
}

// WithError adds error to the logging context
func (l *dummyLogger) WithError(_ error) Logger {
	return l
}
