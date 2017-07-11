package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	User UserConfig `yaml: "user"`
}

type UserConfig struct {
	Account string `yaml:"account"`
	Password string `yaml:"password"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	confData, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(confData, &config)
	if err != nil {
		return config, err
	}

	return config, err
}
