-- +goose Up
-- +goose StatementBegin
CREATE TYPE "volunteer_type" AS ENUM ('GENERAL', 'THERAPIST', 'DOCTOR', 'NURSE', 'SOCIAL_WORKER', 'VETERINARIAN', 'OTHER');

CREATE TABLE "volunteer" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "password_hash" VARCHAR(255) NOT NULL,
    "type" volunteer_type NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "volunteer";
DROP TYPE IF EXISTS volunteer_type;
-- +goose StatementEnd


