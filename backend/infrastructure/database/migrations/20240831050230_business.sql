-- +goose Up
-- +goose StatementBegin
CREATE TABLE "business" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "cnjp" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "latitude" FLOAT NOT NULL,
    "longitude" FLOAT NOT NULL,
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
