package storage

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DbName   string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	MigPath  string `env:"MIGRATE_PATH"`
}

func NewConfig() *Config {
	cfg := &Config{}
	// Здесь cleanenv считывает переменные окружения напрямую в cfg
	if err := cleanenv.ReadEnv(cfg); err != nil {
		panic("Failed to read DB config: " + err.Error())
	}
	return cfg
}
