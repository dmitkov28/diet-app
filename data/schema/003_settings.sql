-- +goose Up
CREATE TABLE IF NOT EXISTS settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_weight REAL NOT NULL,
    target_weight REAL NOT NULL,
    target_weight_loss_rate REAL NOT NULL,
    user_id INT NOT NULL UNIQUE,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE settings;