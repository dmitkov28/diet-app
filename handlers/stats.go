package handlers

import (
	"fmt"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func StatsGETHandler(repo *data.MeasurementRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		items, err := repo.GetMeasurementsByUserId(userId)
		if err != nil {
			fmt.Println(err)
		}
		
		return render(c, templates.StatsPage(89, items))
	}
}
