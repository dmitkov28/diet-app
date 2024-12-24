-- +goose Up
ALTER TABLE food_logs
ADD COLUMN serving_unit TEXT;


-- +goose Down
ALTER TABLE food_logs
DROP COLUMN serving_unit;