package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App
}

type App struct {
	Port           string
	StorageAddress string
}

func NewConfig(envFile string) *Config {
	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}

	return &Config{
		App: App{
			Port:           os.Getenv("APP_PORT"),
			StorageAddress: os.Getenv("STORAGE_ADDRESS"),
		},
	}
}
