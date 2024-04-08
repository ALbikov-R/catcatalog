package routing

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port string `env:"PORT_ROUTER"`
}

func NewConfig() *Config {
	cfg := &Config{}
	// Здесь cleanenv считывает переменные окружения напрямую в cfg
	if err := cleanenv.ReadEnv(cfg); err != nil {
		panic("Failed to read DB config: " + err.Error())
	}
	return cfg
}
