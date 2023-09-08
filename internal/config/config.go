package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database *dbConfig     `yaml:"database"`
	Server   *serverConfig `yaml:"server"`
}

type dbConfig struct {
	ConnectionString string `yaml:"connectionString"`
}

type serverConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func New() *Config {
	var cfg Config
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
