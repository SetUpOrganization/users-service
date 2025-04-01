package repo

import (
	"context"
	"github.com/SetUpOrganization/users-service/internal/infrastructure/db/sqlc/storage"
	"github.com/SetUpOrganization/users-service/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user models.CreateUser) (*models.User, error)
}

type usersRepository struct {
	queries *storage.Queries
}

func NewUsersRepository(queries *storage.Queries) UsersRepository {
	return &usersRepository{queries: queries}
}

func (r *usersRepository) CreateUser(ctx context.Context, user models.CreateUser) (*models.User, error) {
	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}

	createUserParams := storage.CreateUserParams{
		Password: string(hashedPassword),
		Name: pgtype.Text{
			String: user.Name,
			Valid:  true,
		},
	}

	id, err := r.queries.CreateUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}

	createdUser := models.User{
		ID:   id,
		Name: user.Name,
	}

	return &createdUser, nil
}
