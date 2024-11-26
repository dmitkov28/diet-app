package data

import "log"

type Measurement struct {
	ID        int
	WeekStart string
	WeekEnd   string
	Weight    float32
	Calories  float32
}

type MeasurementRepository struct {
	db *DB
}

func NewMeasurementsRepository(db *DB) *MeasurementRepository {
	return &MeasurementRepository{db: db}
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
