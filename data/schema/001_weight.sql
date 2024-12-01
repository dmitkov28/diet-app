-- +goose Up
CREATE TABLE IF NOT EXISTS weight (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATE NOT NULL,
    weight REAL,
    user_id INT NOT NULL,
    UNIQUE(date, user_id) ON CONFLICT REPLACE,
    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE weight;