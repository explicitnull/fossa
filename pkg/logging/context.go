package logging

import (
	"context"
)

type loggerKey struct{}

func PackContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func UnpackContext(ctx context.Context) *Logger {
	val := ctx.Value(loggerKey{})
	if val == nil {
		panic("logger not found in context")
	}

	logger, ok := val.(*Logger)
	if !ok {
		panic("context logging type is not *Logger")
	}

	return logger
}
