-- +goose Up
CREATE TABLE IF NOT EXISTS settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_weight REAL NOT NULL,
    target_weight REAL NOT NULL,
    goal_deadline DATE NOT NULL
);

-- +goose Down
DROP TABLE settings;