--  goose postgres postgres://postgres:7fq4%2366%23OoP5@localhost:5432/rssagg

-- +goose Up

ALTER TABLE users
    ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
        encode(sha256(random()::text::bytea), 'hex')
        );

-- +goose Down

ALTER TABLE users
    DROP COLUMN api_key;
