package data

import (
	"database/sql"
	"fmt"
	"log"
)

type Settings struct {
	ID                      int
	Current_weight          float64
	Target_weight           float64
	Target_weight_loss_rate float64
	Age                     int
	Height                  int
	Sex                     string
	Activity_level          float64
	User_id                 int
}

type SettingsRepository struct {
	db *DB
}

func NewSettingsRepository(db *DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (repo *SettingsRepository) CreateSettings(s Settings) (Settings, error) {
	result, err := repo.db.db.Exec("INSERT OR REPLACE INTO settings(user_id, current_weight, target_weight, target_weight_loss_rate, age, height, sex, activity_level) VALUES(?,?,?,?,?,?,?,?)", s.User_id, s.Current_weight, s.Target_weight, s.Target_weight_loss_rate, s.Age, s.Height, s.Sex, s.Activity_level)
	if err != nil {
		return Settings{}, fmt.Errorf("couldn't write to db: %v", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return Settings{}, fmt.Errorf("couldn't retrieve last inserted ID: %v", err)
	}

	var newSettings Settings
	err = repo.db.db.QueryRow(
		"SELECT id, user_id, current_weight, target_weight, target_weight_loss_rate, age, height, sex, activity_level FROM settings WHERE id = ?",
		lastID,
	).Scan(&newSettings.ID, &newSettings.User_id, &newSettings.Current_weight, &newSettings.Target_weight, &newSettings.Target_weight_loss_rate, &newSettings.Age, &newSettings.Height, &newSettings.Sex, &newSettings.Activity_level)

	if err != nil {
		return Settings{}, fmt.Errorf("couldn't retrieve new settings: %v", err)
	}
	return newSettings, nil

}

func (repo *SettingsRepository) ListSettings() []Settings {
	rows, err := repo.db.db.Query("SELECT id, current_weight, target_weight, target_weight_loss_rate, age, height, sex, activity_level FROM settings")
	if err != nil {
		log.Fatalf("Failed to query the table: %v", err)
	}
	defer rows.Close()
	var data []Settings
	for rows.Next() {
		var m Settings
		err := rows.Scan(&m.ID, &m.Current_weight, &m.Target_weight, &m.Target_weight_loss_rate, &m.Age, &m.Height, &m.Sex, &m.Activity_level)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		data = append(data, m)
	}
	return data
}

func (repo *SettingsRepository) GetSettingsByUserID(userId int) (Settings, error) {
	res := repo.db.db.QueryRow("SELECT id, user_id, current_weight, target_weight, target_weight_loss_rate, age, height, sex, activity_level FROM settings WHERE user_id = ?", userId)
	var settings Settings
	err := res.Scan(&settings.ID, &settings.User_id, &settings.Current_weight, &settings.Target_weight, &settings.Target_weight_loss_rate, &settings.Age, &settings.Height, &settings.Sex, &settings.Activity_level)

	if err != nil {
		if err == sql.ErrNoRows {
			return Settings{}, nil
		}
		return Settings{}, err
	}

	return settings, nil
}