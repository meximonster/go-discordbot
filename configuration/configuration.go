package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var appConfig *Config

type Config struct {
	BotToken           string `yaml:"botToken"`
	GeneralBetAdmin    string `yaml:"GeneralBetAdmin"`
	GeneralBetChannel  string `yaml:"generalBetChannel"`
	PoloBetAdmin       string `yaml:"poloBetAdmin"`
	PoloBetChannel     string `yaml:"poloBetChannel"`
	ParolesOnlyChannel string `yaml:"parolesOnlyChannel"`
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
