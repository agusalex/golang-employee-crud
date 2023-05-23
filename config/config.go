package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var Value *Config

func init() {
	FetchConfig()
}

type Config struct {
	DB struct {
		Host     string `env:"DB_HOST" env-default:"0.0.0.0"`
		User     string `env:"DB_USER" env-default:"root"`
		Name     string `env:"DB_NAME" env-default:"mydbname"`
		Password string `env:"DB_PASS" env-default:"my-secret-pw"`
	}
	Server struct {
		Port string `env:"SERVER_PORT" env-default:"8080"`
	}
}

func FetchConfig() {
	if Value == nil {
		Value = new(Config)
	}
	_ = cleanenv.ReadEnv(Value)
}

func Get() Config {
	if Value == nil {
		panic("Config is nil")
	}
	return *Value
}
