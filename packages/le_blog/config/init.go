package config

import "gopkg.in/ini.v1"

var Config *ini.File

func InitConfig() {
	configPath := "./.env"
	cfg, err := ini.Load(configPath)
	if err != nil {
		panic(err)
	}
	Config = cfg
}
