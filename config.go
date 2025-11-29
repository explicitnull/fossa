package main

import (
	"fossa/api/httpserver"
	"fossa/pkg/jiraclient"
	"fossa/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

const defaultConfFileName = "fossa.yaml"

type Config struct {
	App    AppConfig      `yaml:"app"`
	Logger logging.Config `yaml:"logger"`
	// Postgres   postgres.Config   `yaml:"postgres"`
	HTTPServer httpserver.Config `yaml:"http"`
	// SignIn     security.Config   `yaml:"sign_in"`
	Jira jiraclient.Config `yaml:"jira"`
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

	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		return nil, errors.Wrap(err, "can't parse config")
	}

	// confBytes, err := os.ReadFile(path)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "can't read configuration")
	// }

	// err = json.Unmarshal(confBytes, &conf)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "can't unmarshal configuration")
	// }

	return &conf, nil
}
