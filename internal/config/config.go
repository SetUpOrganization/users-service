package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	DatabaseURL string
	GRPCPort    string
}

func NewConfig() *Config {
	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		slog.Error("DATABASE_URL environment variable not set")
		os.Exit(1)
	}

	grpcPort := os.Getenv("GRPC_SERVER_PORT")
	if grpcPort == "" {
		slog.Error("GRPC_SERVER_PORT environment variable not set")
		os.Exit(1)
	}

	return &Config{
		DatabaseURL: databaseURL,
		GRPCPort:    grpcPort,
	}
}
