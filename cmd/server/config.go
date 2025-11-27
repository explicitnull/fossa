package main

import (
	"encoding/json"
	"fossa/api/httpserver"
	"fossa/pkg/logging"
	"os"

	"github.com/pkg/errors"
)

const defaultConfFileName = "fossa.conf"

type Config struct {
	App    AppConfig      `yaml:"app"`
	Logger logging.Config `yaml:"logger"`
	// Postgres   postgres.Config   `yaml:"postgres"`
	HTTPServer httpserver.Config `yaml:"http"`
	// SignIn     security.Config   `yaml:"sign_in"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig(path string) (*Config, error) {
	var conf Config

	// for faster config reload I want to read JSON via HTTP from consul in future

	// if err := cleanenv.ReadConfig(path, &config); err != nil {
	// 	return nil, errors.Wrap(err, "cannot parse config")
	// }

	confBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't read configuration")
	}

	err = json.Unmarshal(confBytes, &conf)
	if err != nil {
		return nil, errors.Wrap(err, "can't unmarshal configuration")
	}

	return &conf, nil
}
