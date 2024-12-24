package services

import "github.com/dmitkov28/dietapp/internal/data"

type IMeasurementsService interface {
	CreateWeight(weight data.Weight) (data.Weight, error)
	CreateCalories(calories data.Calories) (data.Calories, error)
	GetMeasurementsByUserId(userId, offset int) ([]data.WeightCalories, error)
	DeleteWeightAndCaloriesByWeightID(weightID string) error
	GetWeeklyStats(userId, weeks int) ([]data.WeeklyStats, error)
}

type MeasurementsService struct {
	measurementsRepo *data.MeasurementRepository
}

func NewMeasurementsService(repo *data.MeasurementRepository) IMeasurementsService {
	return &MeasurementsService{measurementsRepo: repo}
}

func (s *MeasurementsService) CreateWeight(weight data.Weight) (data.Weight, error) {
	return s.measurementsRepo.CreateWeight(weight)
}

func (s *MeasurementsService) CreateCalories(calories data.Calories) (data.Calories, error) {
	return s.measurementsRepo.CreateCalories(calories)
}

func (s *MeasurementsService) GetMeasurementsByUserId(userId, offset int) ([]data.WeightCalories, error) {
	return s.measurementsRepo.GetMeasurementsByUserId(userId, offset)
}

func (s *MeasurementsService) GetWeeklyStats(userId int, weeks int) ([]data.WeeklyStats, error) {
	return s.measurementsRepo.GetWeeklyStats(userId, weeks)
}

func (s *MeasurementsService) DeleteWeightAndCaloriesByWeightID(weightID string) error {
	return s.measurementsRepo.DeleteWeightAndCaloriesByWeightID(weightID)
}
