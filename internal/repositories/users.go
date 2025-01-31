package repositories

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Email    string
	Password string
}

type IUsersRepository interface {
	GetUserByEmail(email string) (User, error)
	CreateUser(email, password string) (User, error)
}

type UsersRepository struct {
	db *SqlDB
}

func NewUsersRepository(db *SqlDB) IUsersRepository {
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

func (repo *UsersRepository) CreateUser(email string, password string) (User, error) {
	hashedPassword, err := hashPassword(password)

	if err != nil {
		return User{}, fmt.Errorf("couldn't hash password")
	}

	res, err := repo.db.db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, hashedPassword)
	if err != nil {
		return User{}, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return User{}, err
	}
	if rowsAffected == 0 {
		return User{}, fmt.Errorf("no rows affected, user not created")
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}

	return User{
		ID:    int(lastInsertID),
		Email: email,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
