package service

import (
	"context"
	"github.com/SetUpOrganization/users-service/internal/models"
	"github.com/SetUpOrganization/users-service/internal/repo"
	"github.com/go-playground/validator/v10"
)

type UsersService interface {
	CreateUser(ctx context.Context, user models.CreateUser) (*models.User, error)
}

type usersService struct {
	validate *validator.Validate
	repo     repo.UsersRepository
}

func NewUsersService(repo repo.UsersRepository) UsersService {
	return &usersService{
		validate: validator.New(),
		repo:     repo,
	}
}

func (s *usersService) CreateUser(ctx context.Context, user models.CreateUser) (*models.User, error) {
	err := s.validate.Struct(user)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateUser(ctx, user)
}
