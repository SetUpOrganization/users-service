package main

import (
	"github.com/SetUpOrganization/users-service/internal/config"
	"log/slog"

	"github.com/SetUpOrganization/users-service/internal/transport/grpc"
)

func main() {
	// Init config
	cfg := config.NewConfig()

	if err := grpc.StartGRPCServer(cfg); err != nil {
		slog.Error("failed to start gRPC server", slog.Any("error", err))
	}
}
