package domain

type WeeklyStats struct {
	YearWeek      string
	AverageWeight float64
	PercentChange float64
}

func CalculateBMR(weight float64, height, age int, sex string) float64 {
	if weight == 0 || height == 0 {
		return 0
	}

	if sex == "M" {
		return (13.7516 * weight) + (5.0033 * float64(height)) - (6.755 * float64(age)) + 66.473
	} else {
		return (9.5634 * weight) + (1.8496 * float64(height)) - (4.6756 * float64(age)) + 655.0955
	}
}

func CalculateCalorieGoal(bmr, activityLevel, weight, weightLossRate float64) float64 {
	deficit := CalculateDeficit(weight, weightLossRate)
	return (bmr * activityLevel) - deficit
}

func CalculateDeficit(weight, weightLossRate float64) float64 {
	poundsPerKg := 2.2
	daysPerWeek := 7
	caloriesPerPound := 3500
	return (weight * weightLossRate * poundsPerKg * float64(caloriesPerPound)) / float64(daysPerWeek)
}

func CalculateExpectedDietDuration(currentWeight, targetWeight, targetWeightLossRate float64) float64 {
	if currentWeight == 0 || targetWeight == 0 || targetWeightLossRate == 0 {
		return 0
	}

	weightToLose := currentWeight - targetWeight
	weightToLosePerWeek := currentWeight * targetWeightLossRate
	return weightToLose / weightToLosePerWeek
}

func GetCurrentData(stats []WeeklyStats) WeeklyStats {
	if len(stats) == 0 {
		return WeeklyStats{}
	} else {
		return stats[len(stats)-1]
	}
}
