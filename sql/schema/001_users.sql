--  goose postgres postgres://postgres:7fq4%2366%23OoP5@localhost:5432/rssagg

-- +goose Up

CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
    );

-- +goose Down

DROP TABLE users;