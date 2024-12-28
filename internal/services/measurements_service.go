package services

import "github.com/dmitkov28/dietapp/internal/repositories"

type IMeasurementsService interface {
	CreateWeight(weight repositories.Weight) (repositories.Weight, error)
	CreateCalories(calories repositories.Calories) (repositories.Calories, error)
	GetMeasurementsByUserId(userId, offset int) ([]repositories.WeightCalories, error)
	DeleteWeightAndCaloriesByWeightID(weightID string) error
	GetWeeklyStats(userId, weeks int) ([]repositories.WeeklyStats, error)
}

type MeasurementsService struct {
	measurementsRepo repositories.IMeasurementRepository
}

func NewMeasurementsService(repo repositories.IMeasurementRepository) IMeasurementsService {
	return &MeasurementsService{measurementsRepo: repo}
}

func (s *MeasurementsService) CreateWeight(weight repositories.Weight) (repositories.Weight, error) {
	return s.measurementsRepo.CreateWeight(weight)
}

func (s *MeasurementsService) CreateCalories(calories repositories.Calories) (repositories.Calories, error) {
	return s.measurementsRepo.CreateCalories(calories)
}

func (s *MeasurementsService) GetMeasurementsByUserId(userId, offset int) ([]repositories.WeightCalories, error) {
	return s.measurementsRepo.GetMeasurementsByUserId(userId, offset)
}

func (s *MeasurementsService) GetWeeklyStats(userId int, weeks int) ([]repositories.WeeklyStats, error) {
	return s.measurementsRepo.GetWeeklyStats(userId, weeks)
}

func (s *MeasurementsService) DeleteWeightAndCaloriesByWeightID(weightID string) error {
	return s.measurementsRepo.DeleteWeightAndCaloriesByWeightID(weightID)
}
