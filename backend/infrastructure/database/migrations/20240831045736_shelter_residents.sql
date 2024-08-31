-- +goose Up
-- +goose StatementBegin
CREATE TABLE "shelter_residents" (
    "id" SERIAL PRIMARY KEY,
    "shelter_id" INT NOT NULL,
    "resident_name" VARCHAR(255) NOT NULL
);

ALTER TABLE "shelter_residents" ADD CONSTRAINT "fk_shelter_residents_shelter_id" FOREIGN KEY ("shelter_id") REFERENCES "shelter"("id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "shelter_residents";
-- +goose StatementEnd
