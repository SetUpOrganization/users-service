package main

import (
	"log/slog"

	"github.com/SetUpOrganization/users-service/internal/transport/grpc"
)

func main() {
	if err := grpc.StartGRPCServer("50051"); err != nil {
		slog.Error("failed to start gRPC server", slog.Any("error", err))
	}
}
