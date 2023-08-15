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
	BotToken            string        `yaml:"botToken"`
	Admins              []AdminConfig `yaml:"admins"`
	ParolesOnlyChannel  string        `yaml:"parolesOnlyChannel"`
	PlateChannel        string        `yaml:"plateChannel"`
	POSTGRES_HOST       string        `yaml:"postgres_host"`
	POSTGRES_USER       string        `yaml:"postgres_user"`
	POSTGRES_PASS       string        `yaml:"postgres_password"`
	BNET_CLIENT_ID      string        `yaml:"bnet_client_id"`
	BNET_CLIENT_SECRET  string        `yaml:"bnet_client_secret"`
	PUBG_API_KEY        string        `yaml:"pubg_api_key"`
	PUBG_CURRENT_SEASON string        `yaml:"pubg_current_season"`
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
