-- +goose Up
CREATE TABLE IF NOT EXISTS weight (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATE NOT NULL,
    weight REAL,
    user_id INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE calories;