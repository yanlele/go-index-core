package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Addr        string `yaml:"addr"`
	DSN         string `yaml:"dsn"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
}

var config *Config

func LoadConfig(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(result, &config)
}

func GetConfig() *Config {
	return config
}
