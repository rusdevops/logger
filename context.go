package logger

import (
	"context"
)

const (
	contextKey = "logger"
)

func NewContext(ctx context.Context, log *Logger) context.Context {
	return context.WithValue(ctx, contextKey, log)
}

func FromContext(ctx context.Context) Logger {
	value := ctx.Value(contextKey)
	if value == nil {
		return L()
	}
	log, ok := value.(Logger)
	if !ok {
		return L()
	}
	return log
}
