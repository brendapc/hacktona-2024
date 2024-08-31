-- +goose Up
-- +goose StatementBegin
CREATE TABLE "business" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "cnpj" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    "needs" TEXT NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "business";
-- +goose StatementEnd
