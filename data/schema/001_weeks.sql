-- +goose Up
CREATE TABLE IF NOT EXISTS measurements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    week_start DATE NOT NULL,
    week_end DATE NOT NULL,
    calories REAL,
    weight REAL
);

-- +goose Down
DROP TABLE measurements;