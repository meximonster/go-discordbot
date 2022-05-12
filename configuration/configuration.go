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

type CntConfig struct {
	Name      string      `yaml:"name"`
	UserID    string      `yaml:"userID"`
	ChannelID string      `yaml:"channelID"`
	IsHuman   bool        `yaml:"isHuman"`
	IsPet     bool        `yaml:"isPet"`
	Variety   bool        `yaml:"variety"`
	Images    []ImageInfo `yaml:"images"`
}

type Config struct {
	BotToken        string      `yaml:"botToken"`
	ParolaChannelID string      `yaml:"parolaChannelID"`
	Content         []CntConfig `yaml:"content"`
	Banlist         []string    `yaml:"banlist"`
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

func Write(c *Config) error {
	newEnv, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(".env", newEnv, 0)
}
