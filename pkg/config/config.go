package config

import (
	"gopkg.in/yaml.v3"
	"ioutil"
)

type Config struct {
	Bootstrap string `yaml:"bootstrap_server"`
}

func NewConfig(file ioutil.Reader) {
	return NewConfig(file.ReadFile)
}

func NewConfig(filepath string) (Config, error) {
	cfg := Config{}
	raw_config, err := ioutil.ReadFile(filepath)

	if err != nil {
		return cfg, err
	}

	yaml.Unmarshal(raw_config, &cfg)

	return cfg, nil
}
