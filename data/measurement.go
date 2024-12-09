package data

import (
	"database/sql"

	"golang.org/x/exp/constraints"
)

const (
	ItemsPerPage = 10
)

type Weight struct {
	ID      int
	User_id int
	Weight  float64
	Date    string
}

type Calories struct {
	ID       int
	User_id  int
	Calories int
	Date     string
}

type MeasurementRepository struct {
	db *DB
}

type Number interface {
	constraints.Float | int | int32 | int64
}

func CalculatePercentageDifference[T Number](value1, value2 T) T {

	return ((value2 - value1) * 100) / value1
}

func NewMeasurementsRepository(db *DB) *MeasurementRepository {
	return &MeasurementRepository{db: db}
}

func (repo *MeasurementRepository) CreateWeight(weight Weight) (Weight, error) {
	res, err := repo.db.db.Exec("INSERT INTO weight(user_id, weight, date) VALUES(?, ?, ?)", weight.User_id, weight.Weight, weight.Date)
	if err != nil {
		return Weight{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return Weight{}, err
	}

	var newWeight Weight

	row := repo.db.db.QueryRow("SELECT id, weight, date FROM weight WHERE id = ?", lastId)
	err = row.Scan(&newWeight.ID, &newWeight.Weight, &newWeight.Date)
	if err != nil {
		return Weight{}, err
	}
	return newWeight, nil

}

func (repo *MeasurementRepository) CreateCalories(calories Calories) (Calories, error) {
	res, err := repo.db.db.Exec("INSERT INTO calories(user_id, calories, date) VALUES(?, ?, ?)", calories.User_id, calories.Calories, calories.Date)
	if err != nil {
		return Calories{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return Calories{}, err
	}

	var newCalories Calories

	row := repo.db.db.QueryRow("SELECT id, calories, date FROM calories WHERE id = ?", lastId)
	err = row.Scan(&newCalories.ID, &newCalories.Calories, &newCalories.Date)
	if err != nil {
		return Calories{}, err
	}
	return newCalories, nil

}

type WeightCalories struct {
	WeightID     int     `json:"weight_id"`
	Weight       float64 `json:"weight"`
	WeightDate   string  `json:"weight_date"`
	CaloriesID   *int    `json:"calories_id,omitempty"`
	Calories     *int    `json:"calories,omitempty"`
	CaloriesDate *string `json:"calories_date,omitempty"`
	UserID       int     `json:"user_id"`
}

func (repo *MeasurementRepository) GetMeasurementsByUserId(userId, offset int) ([]WeightCalories, error) {
	query := `
        SELECT 
            w.id, 
            w.weight, 
            w.date as weight_date,
            c.id as calories_id,
            c.calories,
            c.date as calories_date,
            w.user_id
        FROM weight w
        LEFT JOIN calories c ON w.date = c.date AND w.user_id = c.user_id
        WHERE w.user_id = ?
        ORDER BY w.date DESC
		LIMIT 10
		OFFSET ?
		`

	rows, err := repo.db.db.Query(query, userId, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []WeightCalories
	for rows.Next() {
		var wc WeightCalories
		var caloriesID sql.NullInt64
		var calories sql.NullFloat64
		var caloriesDate sql.NullString

		err := rows.Scan(
			&wc.WeightID,
			&wc.Weight,
			&wc.WeightDate,
			&caloriesID,
			&calories,
			&caloriesDate,
			&wc.UserID,
		)
		if err != nil {
			return nil, err
		}

		if caloriesID.Valid {
			val := int(caloriesID.Int64)
			wc.CaloriesID = &val
		}
		if calories.Valid {
			val := int(calories.Float64)
			wc.Calories = &val
		}
		if caloriesDate.Valid {
			wc.CaloriesDate = &caloriesDate.String
		}

		results = append(results, wc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *MeasurementRepository) DeleteWeightAndCaloriesByWeightID(weightID string) error {
	tx, err := repo.db.db.Begin()
	if err != nil {
		return err
	}
	var date string
	tx.QueryRow("SELECT date FROM weight WHERE id = ?", weightID).Scan(&date)
	_, err = tx.Exec("DELETE FROM calories WHERE date = ?", date)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM weight WHERE id = ?", weightID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type WeeklyStats struct {
	YearWeek      string
	AverageWeight float64
	PercentChange float64
}

func (repo *MeasurementRepository) GetWeeklyStats(userId, weeks int) ([]WeeklyStats, error) {
	query := `
		SELECT week, avg_weight
		FROM (
			SELECT 
				strftime('%Y-%W', date) as week,
				AVG(weight) as avg_weight
			FROM weight
			WHERE user_id = ?
			GROUP BY week
			ORDER BY week DESC
			LIMIT ?
		) subquery
		ORDER BY week ASC;`

	rows, err := repo.db.db.Query(query, userId, weeks+1)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var stats []WeeklyStats
	var prevWeight float64

	for rows.Next() {
		var yearWeek string
		var avgWeight float64

		if err := rows.Scan(&yearWeek, &avgWeight); err != nil {
			return nil, err
		}

		var percentChange float64
		if len(stats) > 0 {
			percentChange = ((avgWeight - prevWeight) / prevWeight) * 100
		}

		stats = append(stats, WeeklyStats{
			YearWeek:      yearWeek,
			AverageWeight: avgWeight,
			PercentChange: percentChange,
		})

		prevWeight = avgWeight
	}

	if len(stats) == 0 {
		return []WeeklyStats{}, nil
	}

	if len(stats) < weeks {
		return stats, nil
	}

	return stats[1:], nil
}
