package logging

import (
	"github.com/explicitnull/promcommon"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	DebugLevel = "debug"
)

const (
	DefaultTimeEncoder = "epoch"
)

type Config struct {
	Level       string `yaml:"level" env:"LEVEL"`
	TimeEncoder string `yaml:"time_encoder" env:"TIME_ENCODER"`
	Development bool   `yaml:"development" env:"DEVELOPMENT"`
}

func NewLogger(cfg *Config, appName string, metrics promcommon.LoggerIncrementer) (*Logger, error) {
	if cfg == nil {
		return nil, errors.New("no configuration provided")
	}

	if appName == "" {
		return nil, errors.New("no app name provided")
	}

	zapLogger, err := newZapLogger(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot initialize zap logger")
	}

	zapLogger = zapLogger.With(zap.String("service", appName))

	zapLogger.Info("starting microservice...")

	return &Logger{
		logger:  zapLogger.Sugar(),
		metrics: metrics,
		fields:  map[string]interface{}{},
	}, nil
}

func NewNopLogger() *Logger {
	return &Logger{
		logger: zap.NewNop().Sugar(),
	}
}

func newZapLogger(cfg *Config) (*zap.Logger, error) {
	config := cfg.withDefaults()

	zapConfig, err := config.newZapConfig()
	if err != nil {
		return nil, errors.Wrap(err, "cannot initialize zap configuration")
	}

	return zapConfig.Build()
}

func (c *Config) newZapConfig() (*zap.Config, error) {
	atomicLevel, err := zap.ParseAtomicLevel(c.Level)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse log level")
	}

	var config zap.Config

	if c.Development {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.Level = atomicLevel

	return &config, nil
}

func (c *Config) withDefaults() Config {
	var config Config

	if c != nil {
		config = *c
	}

	if len(config.TimeEncoder) == 0 {
		config.TimeEncoder = DefaultTimeEncoder
	}

	return config
}
