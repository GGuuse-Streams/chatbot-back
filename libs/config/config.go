package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	App    *appConfig    `yaml:"app"`
	Twitch *twitchConfig `yaml:"twitch"`
	Bot    *botConfig    `yaml:"bot"`
}

type appConfig struct {
	AppEnv           string `yaml:"appEnv"`
	ConnectionString string `yaml:"connectionString"`
}

type twitchConfig struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	CallbackUrl  string `yaml:"callbackUrl"`
}

type botConfig struct {
	Username    string `yaml:"username"`
	AccessToken string `yaml:"accessToken"`
	//RefreshToken string `yaml:"refreshToken"`
}

func New() *Config {
	var cfg Config
	err := cleanenv.ReadConfig("../../config/config.yml", &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
