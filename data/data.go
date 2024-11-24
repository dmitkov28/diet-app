package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

type MeasurementRepository struct {
	db *DB
}

func NewMeasurementsRepository(db *DB) *MeasurementRepository {
	return &MeasurementRepository{db: db}
}

type Measurement struct {
	ID        int
	WeekStart string
	WeekEnd   string
	Weight    float32
	Calories  float32
}

func (repo *MeasurementRepository) ListMeasurements() []Measurement {
	rows, err := repo.db.db.Query("SELECT id, week_start, week_end, calories, weight FROM measurements")
	if err != nil {
		log.Fatalf("Failed to query the table: %v", err)
	}
	defer rows.Close()
	var data []Measurement
	var m Measurement
	for rows.Next() {

		err := rows.Scan(&m.ID, &m.WeekStart, &m.WeekEnd, &m.Calories, &m.Weight)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
	}
	return data

}

type SettingsRepository struct {
	db *DB
}

func NewSettingsRepository(db *DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

type Settings struct {
	ID             int
	Current_weight float64
	Target_weight  float64
	Goal_deadline  string
}

func (repo *SettingsRepository) CreateSettings(s Settings) error {
	_, err := repo.db.db.Exec("INSERT INTO settings(current_weight, target_weight, goal_deadline) VALUES(?,?,?)", s.Current_weight, s.Target_weight, s.Goal_deadline)
	if err != nil {
		return fmt.Errorf("couldn't write to db: %v", err)
	}
	return nil

}

func (repo *SettingsRepository) ListSettings() []Settings {
	rows, err := repo.db.db.Query("SELECT id, current_weight, target_weight, goal_deadline FROM settings")
	if err != nil {
		log.Fatalf("Failed to query the table: %v", err)
	}
	defer rows.Close()
	var data []Settings
	for rows.Next() {
		var m Settings
		err := rows.Scan(&m.ID, &m.Current_weight, &m.Target_weight, &m.Goal_deadline)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}	
		data = append(data, m)
	}
	return data
}
