package handlers

import (
	"log"

	"github.com/dmitkov28/dietapp/internal/diet"
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func ScanGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		isHTMX := c.Request().Header.Get("HX-Request") != ""
		return render(c, templates.ScanPage(isHTMX))
	}
}

func ScanBarCodeGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ean := c.Param("ean")
		if ean == "" {
			return render(c, templates.FoodFacts(diet.NutritionData{}))
		}

		data, err := diet.FetchNutritionData(ean)

		if err != nil {
			log.Println(err)
		}
		return render(c, templates.FoodFacts(data))
	}
}
