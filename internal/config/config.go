package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
