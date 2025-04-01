-- name: CreateUser :one
INSERT INTO users (
    password,
    name, surname,
    description,
    phone,
    country,
    avatar_id
) VALUES (
    @password,
    @name, @surname,
    @description,
    @phone,
    @country,
    @avatar_id
)
RETURNING id;
