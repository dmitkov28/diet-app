package handlers

import (
	"time"

	"github.com/dmitkov28/dietapp/templates"
	"github.com/labstack/echo/v4"
)

func DashboardGETHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		today := time.Now().Format("Jan 2, 2006")
		return render(c, templates.HomePage(today))
	}
}
