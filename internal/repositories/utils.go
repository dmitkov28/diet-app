package repositories

import (
	"time"
	"github.com/dmitkov28/dietapp/internal/utils"
)

func HasCurrentWeek(v WeeklyStats) bool {
	currentYear, currentWeek := time.Now().ISOWeek()
	year, week := utils.ParseWeekYearString(v.YearWeek)
	if currentYear != year || currentWeek != week {
		return false
	}
	return true
}
