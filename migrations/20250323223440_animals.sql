-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS animal(
    id UUID  DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL,
    biome TEXT NOT NULL,
    noise TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS animal;
-- +goose StatementEnd
