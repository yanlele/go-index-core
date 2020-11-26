package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Addr 			string		`yaml:"addr"`
	DSN				string		`yaml:"dsn"`
	MaxIdleConn		int			`yaml:"max_idle_conn"`
}

var config *Config

func init() {
	// 加载配置
	err := load("config/config.yaml")
	if err != nil {
		fmt.Println("Failed to load configuration")
		return
	}
}

func load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
