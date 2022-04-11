package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var appConfig *Config

type ImageInfo struct {
	Text string `yaml:"text"`
	Url  string `yaml:"url"`
}

type UserConfig struct {
	Username  string      `yaml:"name"`
	UserID    string      `yaml:"userID"`
	ChannelID string      `yaml:"channelID"`
	Images    []ImageInfo `yaml:"images"`
}

type Config struct {
	BotToken        string       `yaml:"botToken"`
	ParolaChannelID string       `yaml:"parolaChannelID"`
	Users           []UserConfig `yaml:"users"`
	Banlist         []string     `yaml:"banlist"`
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
