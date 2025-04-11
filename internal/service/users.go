package service

import (
	"context"
	"github.com/SetUpOrganization/users-service/internal/models"
	"github.com/SetUpOrganization/users-service/internal/repo"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UsersService interface {
	CreateUser(ctx context.Context, user models.CreateUser) (*models.User, *models.HTTPError)
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

func (s *usersService) CreateUser(ctx context.Context, user models.CreateUser) (*models.User, *models.HTTPError) {
	err := s.validate.Struct(user)
	if err != nil {
		return nil, &models.HTTPError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, &models.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return createdUser, nil
}
