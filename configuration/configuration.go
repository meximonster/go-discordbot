package configuration

import (
	"os"

	"gopkg.in/yaml.v2"
)

var appConfig *Config

type AdminConfig struct {
	Name        string `yaml:"name"`
	Id          string `yaml:"id"`
	Channel     string `yaml:"channel"`
	Table       string `yaml:"table"`
	ExtraGraphs bool   `yaml:"extra_graphs"`
}

type Config struct {
	BotToken           string        `yaml:"botToken"`
	Admins             []AdminConfig `yaml:"admins"`
	ParolesOnlyChannel string        `yaml:"parolesOnlyChannel"`
	POSTGRES_HOST      string        `yaml:"postgres_host"`
	POSTGRES_USER      string        `yaml:"postgres_user"`
	POSTGRES_PASS      string        `yaml:"postgres_password"`
	FORWARD_ENDPOINT   string        `yaml:"forward_endpoint"`
}

func Load() error {
	var c Config
	ymlFile, err := os.ReadFile(".env")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(ymlFile, &c)
	if err != nil {
		return err
	}
	appConfig = &c
	return nil
}

func Read() *Config {
	return appConfig
}
