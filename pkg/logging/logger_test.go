package logging

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func newLogger() *Logger {
	return &Logger{
		logger: zap.L().Sugar(),
		fields: map[string]interface{}{},
	}
}

func getContext(logs *observer.ObservedLogs, idx int) map[string]interface{} {
	return logs.All()[idx].ContextMap()["context"].(map[string]interface{})
}

func TestSimple(t *testing.T) {
	core, logs := observer.New(zapcore.InfoLevel)
	zap.ReplaceGlobals(zap.New(core))

	logger := newLogger()

	testMap := map[string]interface{}{"hello": "world"}

	logger = logger.WithContext("test", testMap)
	logger.Info("test")

	assert.Equal(t, testMap, getContext(logs, 0)["test"])
}

func TestMultipleContext(t *testing.T) {
	core, logs := observer.New(zapcore.InfoLevel)
	zap.ReplaceGlobals(zap.New(core))

	logger := newLogger()

	testMap := map[string]interface{}{"hello": "world"}
	testMap1 := map[string]interface{}{"hello1": "world1"}

	logger = logger.
		WithContext("test", testMap).
		WithContext("test1", testMap1)

	logger.Warn("test")

	assert.Equal(t, testMap, getContext(logs, 0)["test"])
	assert.Equal(t, testMap1, getContext(logs, 0)["test1"])
}

func TestMultipleInfo(t *testing.T) {
	core, logs := observer.New(zapcore.InfoLevel)
	zap.ReplaceGlobals(zap.New(core))

	logger := newLogger()

	testMap := map[string]interface{}{"hello": "world"}
	testMapUpdated := map[string]interface{}{"1": "1"}

	testMap1 := map[string]interface{}{"hello1": "world1"}
	testMap2 := map[string]interface{}{"hello2": "world2"}

	logger = logger.
		WithContext("test", testMap).
		WithContext("test1", testMap1)

	logger.Info("test")

	logger = logger.
		WithContext("test", testMapUpdated).
		WithContext("test2", testMap2)

	logger.Error("test")

	fmt.Println(logs.All())
	assert.Equal(t, testMap, getContext(logs, 0)["test"])
	assert.Equal(t, testMap1, getContext(logs, 0)["test1"])

	assert.Equal(t, testMapUpdated, getContext(logs, 1)["test"])
	assert.Equal(t, testMap1, getContext(logs, 1)["test1"])
	assert.Equal(t, testMap2, getContext(logs, 1)["test2"])
}
