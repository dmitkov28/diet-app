-- +goose Up
CREATE TABLE IF NOT EXISTS calories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATE NOT NULL,
    calories REAL,
    user_id INT NOT NULL,
    UNIQUE(date, user_id) ON CONFLICT REPLACE,
    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE calories;