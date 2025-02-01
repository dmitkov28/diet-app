package use_cases

import (
	"strconv"

	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
)

func AddWeightUseCase(service services.IMeasurementsService, userId int, weightValue string, dateValue string) error {

	weight, err := strconv.ParseFloat(weightValue, 64)

	if err != nil || weight <= 0 {
		return err
	}

	newWeight := repositories.Weight{
		User_id: userId,
		Weight:  weight,
		Date:    dateValue,
	}

	_, err = service.CreateWeight(newWeight)
	if err != nil {
		return err
	}

	return nil
}
