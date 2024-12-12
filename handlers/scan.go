package handlers

import (
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func ScanGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.ScanPage())
	}
}

func ScanBarCodeGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ean := c.Param("ean")
		return render(c, templates.FoodFacts(ean))
	}
}
