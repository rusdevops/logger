package logger

import (
	"context"
)

type contextKeyType string

const (
	contextKey contextKeyType = "logger|a809358f-7e86-4d8f-9959-054c66b6095c"
)

func NewContext(ctx context.Context, log Logger) context.Context {
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
