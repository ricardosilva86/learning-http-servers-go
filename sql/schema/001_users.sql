-- +goose Up
CREATE TABLE users (
    id uuid primary key default gen_random_uuid(),
    email text not null unique,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- +goose Down
DROP TABLE users;