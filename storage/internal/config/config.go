package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App
	DB
}

type App struct {
	Port            string
	BufferSizeLimit int
}

type DB struct {
	Address  string
	Database string
	Username string
	Password string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	bufferSizeLimit, err := strconv.Atoi(os.Getenv("APP_BUFFER_SIZE"))
	if err != nil {
		panic(err)
	}

	config := Config{
		App: App{
			Port:            os.Getenv("APP_PORT"),
			BufferSizeLimit: bufferSizeLimit,
		},
		DB: DB{
			Address:  os.Getenv("CLICKHOUSE_ADDRESS"),
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Username: os.Getenv("CLICKHOUSE_USERNAME"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
	}

	return &config
}
