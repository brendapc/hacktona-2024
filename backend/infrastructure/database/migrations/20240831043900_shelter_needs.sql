-- +goose Up
-- +goose StatementBegin
CREATE TABLE "shelter_needs" (
    "id" SERIAL PRIMARY KEY,
    "shelter_id" INT NOT NULL,
    "bedding_item" BOOLEAN NOT NULL DEFAULT FALSE,
    "food_non_perishable" BOOLEAN NOT NULL DEFAULT FALSE,
    "food_perishable" BOOLEAN NOT NULL DEFAULT FALSE,
    "hygiene_products" BOOLEAN NOT NULL DEFAULT FALSE,
    "clothing_male" BOOLEAN NOT NULL DEFAULT FALSE,
    "clothing_female" BOOLEAN NOT NULL DEFAULT FALSE,
    "clothing_children" BOOLEAN NOT NULL DEFAULT FALSE,
    "medical_supplies" BOOLEAN NOT NULL DEFAULT FALSE,
    "pet_food_dogs" BOOLEAN NOT NULL DEFAULT FALSE,
    "pet_food_cats" BOOLEAN NOT NULL DEFAULT FALSE,
    "cleaning_supplies" BOOLEAN NOT NULL DEFAULT FALSE
);

ALTER TABLE "shelter_needs" ADD CONSTRAINT "fk_shelter_needs_shelter_id" FOREIGN KEY ("shelter_id") REFERENCES "shelter"("id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "shelter_needs";
-- +goose StatementEnd
