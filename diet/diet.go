package diet

import "github.com/dmitkov28/dietapp/data"

func CalculateBMR(weight float64, height, age int, sex string) float64 {
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

func CaclulateExpectedDietDuration(currentWeight, targetWeight, targetWeightLossRate float64) float64 {
	weightToLose := currentWeight - targetWeight
	weightToLosePerWeek := currentWeight * targetWeightLossRate
	return weightToLose / weightToLosePerWeek
}

func CalculateAverageWeight(items []data.WeightCalories) float64 {
	if len(items) == 0 {
		return 0
	}
	sum := float64(0)
	for _, item := range items {
		sum += item.Weight
	}
	return sum / float64(len(items))

}
