package models

import "github.com/google/uuid"

type User struct {
	ID               uuid.UUID
	UnhashedPassword string `validate:"required,min=8"`
	HashedPassword   string
	Name             string `validate:"omitempty,min=4,max=50"`
	Surname          string `validate:"omitempty,min=4,max=50"`
	Description      string `validate:"omitempty,min=4,max=1000"`
	Phone            string `validate:"required,min=4,max=15"`
	Country          string `validate:"omitempty,min=2,max=2"`
	AvatarID         uuid.UUID
}

type CreateUser struct {
	// TODO: email
	Password string `validate:"required,min=8"`
	Name     string `validate:"omitempty,min=4,max=50"`
}
