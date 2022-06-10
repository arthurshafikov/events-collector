package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App
}

type App struct {
	Port string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		App: App{
			Port: os.Getenv("APP_PORT"),
		},
	}
}
