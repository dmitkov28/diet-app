-- +goose Up
CREATE TABLE IF NOT EXISTS settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    current_weight REAL NOT NULL,
    target_weight REAL NOT NULL,
    target_weight_loss_rate REAL NOT NULL,
    height INT NOT NULL,
    age INT NOT NULL,
    sex TEXT NOT NULL,
    activity_level REAL NOT NULL,
    user_id INT NOT NULL UNIQUE,
    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE

);

-- +goose Down
DROP TABLE settings;