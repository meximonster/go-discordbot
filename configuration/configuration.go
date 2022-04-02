package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var appConfig *Config

type Config struct {
	BotToken  string `yaml:"botToken"`
	ChannelID string `yaml:"channelID"`
	UserID    string `yaml:"userID"`
}

func Load() error {
	var c Config
	ymlFile, err := ioutil.ReadFile(".env")
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
