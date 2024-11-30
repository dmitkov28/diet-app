package data

import "time"

func ParseDateString(dateString string) string {
	parsed, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return "NaN"
	}
	return parsed.Format("02 Jan 06")
}

func GetPreviousWeekRange() (time.Time, time.Time) {

	now := time.Now()

	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	currentDayMidnight := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)

	currentWeekday := int(currentDayMidnight.Weekday())

	if currentWeekday == 0 {
		currentWeekday = 7
	}

	daysToPreviousMonday := currentWeekday + 6
	startOfPreviousWeek := currentDayMidnight.AddDate(0, 0, -daysToPreviousMonday)

	endOfPreviousWeek := startOfPreviousWeek.AddDate(0, 0, 6).Add(time.Hour*24 - time.Nanosecond)

	return startOfPreviousWeek, endOfPreviousWeek
}


