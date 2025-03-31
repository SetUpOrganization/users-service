package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/SetUpOrganization/protos/gen/go/users"
	"google.golang.org/grpc"
)

type server struct {
	users.UnimplementedUsersServer
}

func (s *server) SignUp(ctx context.Context, req *users.SignUpRequest) (*users.SignUpResponse, error) {
	slog.Info("Received request", slog.String("name", req.Name))
	return &users.SignUpResponse{Success: true, Message: "Successfully registered"}, nil
}

func StartGRPCServer(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	users.RegisterUsersServer(s, &server{})

	slog.Info(fmt.Sprintf("gRPC server listening on %s", port))
	return s.Serve(lis)
}
