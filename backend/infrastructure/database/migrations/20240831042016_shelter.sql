-- +goose Up
-- +goose StatementBegin
CREATE TABLE "shelter" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    "capacity" INT NOT NULL,
    "current_occupancy" INT NOT NULL,
    "capacity_pets" INT NOT NULL,
    "current_occupancy_pets" INT NOT NULL,
    "acessibility" BOOLEAN NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "shelter";
-- +goose StatementEnd
