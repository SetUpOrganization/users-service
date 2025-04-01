-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    password TEXT NOT NULL,
    name VARCHAR(50),
    surname VARCHAR(50),
    description TEXT,
    phone VARCHAR(15) NOT NULL,
    country VARCHAR(2),
    avatar_id TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
