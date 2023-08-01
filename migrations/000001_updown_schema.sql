-- +goose Up
CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    state VARCHAR(255) NOT NULL
);
-- +goose Down
DROP TABLE users;