package handlers

import (
	"fmt"

	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/dmitkov28/dietapp/internal/use_cases"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func StatsGETHandler(measurementsService services.IMeasurementsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("user_id").(int)
		pageParam := c.QueryParam("page")
		orderByParam := c.QueryParam("orderBy")
		orderParam := c.QueryParam("order")

		data, err := use_cases.GetUserStatsUseCase(measurementsService, userId, pageParam, orderByParam, orderParam)

		if err != nil {
			fmt.Println(err)
		}

		isHTMX := c.Request().Header.Get("HX-Request") != ""

		return render(c, templates.StatsPage(data.Items, data.Page, data.NoMoreResults, isHTMX, data.SortOptions))
	}
}

func StatsDELETEHandler(measurementsService services.IMeasurementsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := measurementsService.DeleteWeightAndCaloriesByWeightID(id)
		if err != nil {
			fmt.Println(err)
		}
		return c.NoContent(202)

	}
}
