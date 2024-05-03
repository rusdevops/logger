package logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]any

// Logger is our contract for the logger
type Logger interface {
	SimpleLogger
	FormattedLogger
	StructedLogger
}

type SimpleLogger interface {
	// Debug uses fmt.Sprint to construct and log a message.
	Debug(args ...any)

	// Info uses fmt.Sprint to construct and log a message.
	Info(args ...any)

	// Warn uses fmt.Sprint to construct and log a message.
	Warn(args ...any)

	// Error uses fmt.Sprint to construct and log a message.
	Error(args ...any)

	// Panic uses fmt.Sprint to construct and log a message, then panics.
	Panic(args ...any)

	// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
	Fatal(args ...any)
}

type FormattedLogger interface {
	// Debugf uses fmt.Sprintf to log a templated message.
	Debugf(template string, args ...any)

	// Infof uses fmt.Sprintf to log a templated message.
	Infof(template string, args ...any)

	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(template string, args ...any)

	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(template string, args ...any)

	// Panicf uses fmt.Sprintf to log a templated message, then panics.
	Panicf(template string, args ...any)

	// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
	Fatalf(template string, args ...any)
}

type StructedLogger interface {
	WithFields(keyValues Fields) Logger

	WithField(key string, val any) Logger

	WithError(err error) Logger
}
