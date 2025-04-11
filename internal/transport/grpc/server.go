package grpc

import (
	"context"
	"fmt"
	"github.com/SetUpOrganization/users-service/internal/config"
	"github.com/SetUpOrganization/users-service/internal/infrastructure/db/sqlc/storage"
	"github.com/SetUpOrganization/users-service/internal/models"
	"github.com/SetUpOrganization/users-service/internal/repo"
	"github.com/SetUpOrganization/users-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net"

	"github.com/SetUpOrganization/protos/gen/go/users"
	"google.golang.org/grpc"
)

type server struct {
	users.UnimplementedUsersServer
	service service.UsersService
}

func (s *server) SignUp(ctx context.Context, req *users.SignUpRequest) (*users.SignUpResponse, error) {
	slog.Info("Received request", slog.String("name", req.Name))

	user := models.CreateUser{
		Name:     req.Name,
		Password: req.Password,
	}

	createdUser, err := s.service.CreateUser(ctx, user)
	if err != nil {
		return &users.SignUpResponse{Success: false, Message: err.Message}, err.Err()
	}

	return &users.SignUpResponse{Success: true, Message: fmt.Sprintf("Successfully registered user with name %s and id %s", createdUser.Name, createdUser.ID)}, nil
}

func StartGRPCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	// Init db connection
	conn, err := InitDB(context.Background(), cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to connect to db", slog.Any("err", err))
	}

	// Init SQL queries
	queries := storage.New(conn)

	usersRepo := repo.NewUsersRepository(queries)
	usersService := service.NewUsersService(usersRepo)

	serv := server{
		service: usersService,
	}
	users.RegisterUsersServer(s, &serv)

	slog.Info(fmt.Sprintf("gRPC server listening on %s", cfg.GRPCPort))
	return s.Serve(lis)
}

func InitDB(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	pool.Config().MaxConns = 100 // Max 100 connections

	return pool, nil
}
