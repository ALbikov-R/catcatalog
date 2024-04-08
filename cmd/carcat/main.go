package main

import (
	"carcat/internal/service"
	"log"

	"github.com/joho/godotenv"
)

const _configPath = "configs/config.env"

// @title Car info API
// @version 1.0
// @description Car info service.
// @host localhost:9092
// @BasePath /
func main() {
	if err := godotenv.Load(_configPath); err != nil {
		log.Fatalf("error loading config file: %s\n", err)
	}
	cfg := service.NewConfig()
	serv := service.New(cfg)
	serv.Start()
}
