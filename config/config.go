package config

import (
	"github.com/bykovme/goconfig"
)

// Config - structure of config file
type Config struct {
	ServerPort string `json:"server_port"`
	Nodes      []Node `json:"nodes"`
}

// Node -
type Node struct {
	Ticker string `json:"ticker"`
	UUID   string `json:"uuid"`
}

const cConfigPath = "/etc/ada-rocket/controller.conf"

// var loadedConfig Config

func LoadConfig() (loadedConfig Config, err error) {
	usrHomePath, err := goconfig.GetUserHomePath()
	if err == nil {
		err = goconfig.LoadConfig(usrHomePath+cConfigPath, &loadedConfig)
		if err != nil {
			return loadedConfig, err
		}
	} else {
		return loadedConfig, err
	}
	return loadedConfig, err
}
