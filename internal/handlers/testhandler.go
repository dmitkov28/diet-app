package handlers

import (
	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func TestPOSTHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return render(c, templates.FoodLogSuccess())
	}
}
