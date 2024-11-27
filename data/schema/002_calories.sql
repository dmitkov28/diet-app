-- +goose Up
CREATE TABLE IF NOT EXISTS calories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATE NOT NULL,
    calories REAL,
    user_id INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE calories;