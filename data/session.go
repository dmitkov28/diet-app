package data

import (
	"database/sql"
	"time"
)

type Session struct {
	User_id    int
	Expires_At time.Time
	Token      string
}

type SessionsRepository struct {
	db *DB
}

func NewSessionsRepository(db *DB) *SessionsRepository {
	return &SessionsRepository{db: db}
}

func (repo *SessionsRepository) CreateSession(session Session) (Session, error) {
	_, err := repo.db.db.Exec("INSERT OR REPLACE INTO sessions(user_id, token, expires_at) VALUES(?, ?, ?)", session.User_id, session.Token, session.Expires_At)
	if err != nil {
		return Session{}, nil
	}
	return session, nil
}

func (repo *SessionsRepository) GetSessionByToken(token string) (Session, error) {
	row := repo.db.db.QueryRow("SELECT user_id, expires_at, token FROM sessions WHERE token = ?", token)
	var session Session
	err := row.Scan(&session.User_id, &session.Expires_At, &session.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return Session{}, sql.ErrNoRows
		}
		return Session{}, err
	}
	return session, nil
}
