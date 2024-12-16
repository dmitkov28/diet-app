package main

import (
	"log"
	"net/http"

	"github.com/dmitkov28/dietapp/data"
	"github.com/dmitkov28/dietapp/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {

	db, err := data.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	usersRepo := data.NewUsersRepository(db)
	sessionsRepo := data.NewSessionsRepository(db)
	settingsRepo := data.NewSettingsRepository(db)
	measurementsRepo := data.NewMeasurementsRepository(db)
	foodLogRepo := data.NewFoodLogsRepository(db)

	e := echo.New()
	e.Static("/static", "static")
	e.File("/favicon.ico", "static/img/favicon/favicon.ico")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/stats")
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	}, authMiddleware(sessionsRepo))

	e.GET("/dashboard", handlers.DashboardGETHandler(measurementsRepo, settingsRepo), authMiddleware(sessionsRepo))

	e.GET("/settings", handlers.SettingsGETHandler(settingsRepo), authMiddleware(sessionsRepo))
	e.POST("/settings", handlers.SettingsPOSTHandler(settingsRepo), authMiddleware(sessionsRepo))

	e.GET("/weight", handlers.WeightGETHandler(measurementsRepo), authMiddleware(sessionsRepo))
	e.POST("/weight", handlers.WeightPOSTHandler(measurementsRepo), authMiddleware(sessionsRepo))

	e.GET("/stats", handlers.StatsGETHandler(measurementsRepo), authMiddleware(sessionsRepo))
	e.DELETE("/stats/:id", handlers.StatsDELETEHandler(measurementsRepo), authMiddleware(sessionsRepo))

	e.GET("/calories", handlers.CaloriesGETHandler(measurementsRepo), authMiddleware(sessionsRepo))
	e.POST("/calories", handlers.CaloriesPOSTHandler(measurementsRepo), authMiddleware(sessionsRepo))

	e.GET("/scan", handlers.ScanGETHandler(), authMiddleware(sessionsRepo))
	e.GET("/scan/:ean", handlers.ScanBarCodeGETHandler(), authMiddleware(sessionsRepo))

	e.GET("/search", handlers.SearchFoodGETHandler(measurementsRepo), authMiddleware(sessionsRepo))
	e.GET("/search_food", handlers.SearchFoodGetHandlerWithParams(measurementsRepo), authMiddleware(sessionsRepo))

	e.GET("/food_log", handlers.FoodLogGETHandler(foodLogRepo, settingsRepo), authMiddleware(sessionsRepo))
	e.GET("/refresh_totals", handlers.FoodLogRefreshTotalsGETHandler(foodLogRepo, settingsRepo), authMiddleware(sessionsRepo))
	e.POST("/food_log", handlers.FoodLogPOSTHandler(foodLogRepo, settingsRepo), authMiddleware(sessionsRepo))
	e.DELETE("/food_log/:id", handlers.FoodLogDELETEHandler(foodLogRepo), authMiddleware(sessionsRepo))

	e.GET("/login", handlers.LoginGETHandler())
	e.POST("/login", handlers.LoginPOSTHandler(usersRepo, sessionsRepo))

	e.Logger.Fatal(e.Start(":1323"))
}
