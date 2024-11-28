package data

import (
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
	User_id                 int
}

type SettingsRepository struct {
	db *DB
}

func NewSettingsRepository(db *DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (repo *SettingsRepository) CreateSettings(s Settings) (Settings, error) {
	result, err := repo.db.db.Exec("INSERT OR REPLACE INTO settings(user_id, current_weight, target_weight, target_weight_loss_rate, age, height, sex) VALUES(?,?,?,?,?,?,?)", s.User_id, s.Current_weight, s.Target_weight, s.Target_weight_loss_rate, s.Age, s.Height, s.Sex)
	if err != nil {
		return Settings{}, fmt.Errorf("couldn't write to db: %v", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return Settings{}, fmt.Errorf("couldn't retrieve last inserted ID: %v", err)
	}

	var newSettings Settings
	err = repo.db.db.QueryRow(
		"SELECT id, user_id, current_weight, target_weight, target_weight_loss_rate, age, height, sex FROM settings WHERE id = ?",
		lastID,
	).Scan(&newSettings.ID, &newSettings.User_id, &newSettings.Current_weight, &newSettings.Target_weight, &newSettings.Target_weight_loss_rate, &newSettings.Age, &newSettings.Height, &newSettings.Sex)

	if err != nil {
		return Settings{}, fmt.Errorf("couldn't retrieve new settings: %v", err)
	}
	return newSettings, nil

}

func (repo *SettingsRepository) ListSettings() []Settings {
	rows, err := repo.db.db.Query("SELECT id, current_weight, target_weight, target_weight_loss_rate, age, height, sex FROM settings")
	if err != nil {
		log.Fatalf("Failed to query the table: %v", err)
	}
	defer rows.Close()
	var data []Settings
	for rows.Next() {
		var m Settings
		err := rows.Scan(&m.ID, &m.Current_weight, &m.Target_weight, &m.Target_weight_loss_rate, &m.Age, &m.Height, &m.Sex)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		data = append(data, m)
	}
	return data
}

func (repo *SettingsRepository) GetSettingsByUserID(userId string) (Settings, error) {
	res := repo.db.db.QueryRow("SELECT id, user_id, current_weight, target_weight, target_weight_loss_rate, age, height, sex FROM settings WHERE user_id = ?", userId)
	var settings Settings
	err := res.Scan(&settings.ID, &settings.User_id, &settings.Current_weight, &settings.Target_weight, &settings.Target_weight_loss_rate, &settings.Age, &settings.Height, &settings.Sex)

	if err != nil {
		return Settings{}, err
	}

	return settings, nil
}
