package data

import "database/sql"

type User struct {
	ID       int
	Email    string
	Password string
}

type UsersRepository struct {
	db *DB
}

func NewUsersRepository(db *DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (repo *UsersRepository) GetUserByEmail(email string) (User, error) {
	row := repo.db.db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, sql.ErrNoRows
		}
		return User{}, err
	}
	return user, nil
}
