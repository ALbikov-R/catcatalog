package service

import (
	"carcat/internal/storage"
	"carcat/internal/transport/routing"
)

type config struct {
	db     *storage.Config
	router *routing.Config
}

func NewConfig() *config {
	return &config{
		db:     storage.NewConfig(),
		router: routing.NewConfig(),
	}
}
