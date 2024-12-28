-- +goose Up
CREATE TABLE IF NOT EXISTS food_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    food_name TEXT NOT NULL,
    serving_size REAL NOT NULL,
    number_of_servings REAL NOT NULL,
    calories REAL NOT NULL,
    protein REAL NOT NULL,
    carbs REAL NOT NULL,
    fats REAL NOT NULL,
    created_at DATE NOT NULL DEFAULT (DATE('now')),

    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE
);


-- +goose Down
DROP TABLE food_logs;