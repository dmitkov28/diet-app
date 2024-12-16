package data

import (
	"strconv"
	"strings"
	"time"
)

func ParseDateString(dateString string) string {
	parsed, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		return "NaN"
	}
	return parsed.Format("02 Jan 06")
}

func ParseWeekYearString(weekYearString string) (int, int) {
	splitStr := strings.Split(weekYearString, "-")
	if len(splitStr) != 2 {
		return 0, 0
	}

	year, err := strconv.ParseInt(splitStr[0], 10, 64)
	if err != nil {
		return 0, 0
	}
	week, err := strconv.ParseInt(splitStr[1], 10, 64)

	if err != nil {
		return 0, 0
	}

	return int(year), int(week)

}

func HasCurrentWeek(v WeeklyStats) bool {
	currentYear, currentWeek := time.Now().ISOWeek()
	year, week := ParseWeekYearString(v.YearWeek)
	if currentYear != year || currentWeek != week {
		return false
	}
	return true
}
