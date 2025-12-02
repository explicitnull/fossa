package main

import (
	"fmt"
	"fossa/api/httpserver"
	"fossa/pkg/jiraclient"
	"fossa/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultConfFileName = "fossa.yml"

type Config struct {
	App    AppConfig      `yaml:"app"`
	Logger logging.Config `yaml:"logger"`
	// Postgres   postgres.Config   `yaml:"postgres"`
	HTTPServer httpserver.Config `yaml:"http"`
	// SignIn     security.Config   `yaml:"sign_in"`
	Jira      jiraclient.Config //`env:"jira"`
	Templates Templates         `yaml:"templates"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
}

type Templates struct {
	Path string `yaml:"path"`
}

func NewConfig(path string) (*Config, error) {
	var conf Config
	var jiraConfig jiraclient.Config

	// for faster config reload I want to read JSON via HTTP from consul in future

	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		return nil, fmt.Errorf("cant load config %v", err)
	}

	if err := cleanenv.ReadConfig(".env", &jiraConfig); err != nil {
		return nil, fmt.Errorf("cant load env data to config %v", err)
	}
	conf.Jira = jiraConfig
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
