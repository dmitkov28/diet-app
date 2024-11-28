package data

import (
	"time"

	"golang.org/x/exp/constraints"
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

func ParseDateString(dateString string) string {
	parsed, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return "NaN"
	}
	return parsed.Format("02 Jan 2006")
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

func (repo *MeasurementRepository) GetWeightByUserId(userId int) ([]Weight, error) {
	rows, err := repo.db.db.Query("SELECT weight, date FROM weight WHERE user_id = ?", userId)
	if err != nil {
		return []Weight{}, err
	}
	defer rows.Close()
	var result []Weight
	for rows.Next() {
		var row Weight
		err := rows.Scan(&row.Weight, &row.Date)
		if err != nil {
			return []Weight{}, err
		}
		result = append(result, row)
	}
	return result, nil

}

func (repo *MeasurementRepository) GetWeightsBetweenDates(userId int, startDate, endDate string) ([]Weight, error) {
	rows, err := repo.db.db.Query("SELECT weight, date FROM weight WHERE user_id = ? AND date BETWEEN ? AND ?", userId, startDate, endDate)
	if err != nil {
		return []Weight{}, err
	}
	defer rows.Close()
	var result []Weight
	for rows.Next() {
		var row Weight
		err := rows.Scan(&row.Weight, &row.Date)
		if err != nil {
			return []Weight{}, err
		}
		result = append(result, row)
	}
	return result, nil

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

func (repo *MeasurementRepository) GetCaloriesByUserId(userId int) ([]Calories, error) {
	rows, err := repo.db.db.Query("SELECT calories, date FROM calories WHERE user_id = ?", userId)
	if err != nil {
		return []Calories{}, err
	}
	defer rows.Close()
	var result []Calories
	for rows.Next() {
		var row Calories
		err := rows.Scan(&row.Calories, &row.Date)
		if err != nil {
			return []Calories{}, err
		}
		result = append(result, row)
	}
	return result, nil

}

func (repo *MeasurementRepository) GetCaloriesBetweenDates(userId int, startDate, endDate string) ([]Calories, error) {
	rows, err := repo.db.db.Query("SELECT calories, date FROM calories WHERE user_id = ? AND date BETWEEN ? AND ?", userId, startDate, endDate)
	if err != nil {
		return []Calories{}, err
	}
	defer rows.Close()
	var result []Calories
	for rows.Next() {
		var row Calories
		err := rows.Scan(&row.Calories, &row.Date)
		if err != nil {
			return []Calories{}, err
		}
		result = append(result, row)
	}
	return result, nil

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
