-- +goose Up
-- +goose StatementBegin
CREATE TABLE "shelter_needs" (
    "id" SERIAL PRIMARY KEY,
    "shelter_id" INT NOT NULL,
    "bedding_item" BOOLEAN NOT NULL,
    "food_non_perishable" BOOLEAN NOT NULL,
    "food_perishable" BOOLEAN NOT NULL,
    "hygiene_products" BOOLEAN NOT NULL,
    "clothing_male" BOOLEAN NOT NULL,
    "clothing_female" BOOLEAN NOT NULL,
    "clothing_children" BOOLEAN NOT NULL,
    "medical_supplies" BOOLEAN NOT NULL,
    "pet_food_dogs" BOOLEAN NOT NULL,
    "pet_food_cats" BOOLEAN NOT NULL,
    "cleaning_supplies" BOOLEAN NOT NULL
);

ALTER TABLE "shelter_needs" ADD CONSTRAINT "fk_shelter_needs_shelter_id" FOREIGN KEY ("shelter_id") REFERENCES "shelter"("id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "shelter_needs";
-- +goose StatementEnd
