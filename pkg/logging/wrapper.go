package logging

import (
	"github.com/explicitnull/promcommon"
	"go.uber.org/zap"
)

const (
	contextKey = "context"

	contextScenarioKey = "scenario"
	contextEventKey    = "event"
)

type Logger struct {
	logger  *zap.SugaredLogger
	metrics promcommon.LoggerIncrementer

	fields map[string]interface{}
}

// Deprecated: Use WithContext.
func (log *Logger) With(fields ...interface{}) *Logger {
	return &Logger{
		logger:  log.logger.With(fields...),
		metrics: log.metrics,
		fields:  log.copyFields(),
	}
}

func (log *Logger) WithContext(key string, obj any) *Logger {
	newLogger := &Logger{
		logger:  log.logger,
		metrics: log.metrics,
		fields:  log.copyFields(),
	}

	newLogger.fields[key] = obj

	return newLogger
}

func (log *Logger) copyFields() map[string]interface{} {
	fields := map[string]interface{}{}

	for k, v := range log.fields {
		fields[k] = v
	}

	return fields
}

func (log *Logger) withContext() *zap.SugaredLogger {
	return log.logger.With(zap.Any(contextKey, log.fields))
}

func (log *Logger) Debug(msg string, keysAndValues ...interface{}) {
	log.withContext().Debugw(msg, keysAndValues...)
}

func (log *Logger) Info(msg string, keysAndValues ...interface{}) {
	log.withContext().Infow(msg, keysAndValues...)
}

func (log *Logger) Warn(msg string, keysAndValues ...interface{}) {
	if log.metrics != nil {
		log.metrics.IncLogWarns()
	}

	log.withContext().Warnw(msg, keysAndValues...)
}

func (log *Logger) Error(msg string, keysAndValues ...interface{}) {
	if log.metrics != nil {
		log.metrics.IncLogErrors()
	}

	log.withContext().Errorw(msg, keysAndValues...)
}

func (log *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	log.withContext().Fatalw(msg, keysAndValues...)
}
