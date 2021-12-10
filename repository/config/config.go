package config

import (
	"github.com/bykovme/goconfig"
)

// Config - structure of config file
type Config struct {
	ServerPort string   `json:"server_port"`
	Nodes      []Node   `json:"nodes"`
	DBConfig   DBConfig `json:"db"`
}

// Node -
type Node struct {
	UUID       string `json:"uuid"`
	Blockchain string `json:"blockchain"`
}

// DBConfig ...
type DBConfig struct {
	// User ...
	User string `json:"user"`
	// Password ...
	Password string `json:"password"`
	// Dbname ...
	Dbname string `json:"dbname"`
	// Sslmode ...
	Sslmode string `json:"sslmode"`
}

const cConfigPath = "etc/black-rocket/controller.conf"
const cConfigPathReserve = "etc/ada-rocket/controller.conf"

// var loadedConfig Config

func LoadConfig() (loadedConfig Config, err error) {
	usrHomePath, err := goconfig.GetUserHomePath()
	if err != nil {
		return loadedConfig, err
	}

	err = goconfig.LoadConfig(usrHomePath+cConfigPath, &loadedConfig)
	if err == nil {
		return loadedConfig, nil
	}

	err = goconfig.LoadConfig(usrHomePath+cConfigPathReserve, &loadedConfig)

	return loadedConfig, err
}
