package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmitkov28/dietapp/internal/data"
	"github.com/dmitkov28/dietapp/internal/diet"
	"github.com/dmitkov28/dietapp/internal/handlers"
	"github.com/dmitkov28/dietapp/internal/httputils"
	customMiddleware "github.com/dmitkov28/dietapp/internal/middleware"
	"github.com/dmitkov28/dietapp/internal/services"
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

	// repositories
	usersRepo := data.NewUsersRepository(db)
	sessionsRepo := data.NewSessionsRepository(db)
	settingsRepo := data.NewSettingsRepository(db)
	measurementsRepo := data.NewMeasurementsRepository(db)
	foodLogRepo := data.NewFoodLogsRepository(db)

	httpClient := http.Client{}
	apiClient := httputils.NewAPIClient(&httpClient)
	nutritionixAPIClient, err := diet.NewNutritionixAPIClient(apiClient)

	if err != nil {
		fmt.Println(err)
	}

	// services
	authService := services.NewAuthService(usersRepo, sessionsRepo)

	

	e := echo.New()
	e.Static("/static", "static")
	e.File("/favicon.ico", "static/img/favicon/favicon.ico")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/stats")
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	}, customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/dashboard", handlers.DashboardGETHandler(measurementsRepo, settingsRepo), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/settings", handlers.SettingsGETHandler(settingsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/settings", handlers.SettingsPOSTHandler(settingsRepo), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/weight", handlers.WeightGETHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/weight", handlers.WeightPOSTHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/stats", handlers.StatsGETHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.DELETE("/stats/:id", handlers.StatsDELETEHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/calories", handlers.CaloriesGETHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/calories", handlers.CaloriesPOSTHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/scan", handlers.ScanGETHandler(), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/scan/:ean", handlers.ScanBarCodeGETHandler(), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/search", handlers.SearchFoodGETHandler(measurementsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/search_food", handlers.SearchFoodGetHandlerWithParams(nutritionixAPIClient), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/search_food/modal", handlers.SearchFoodModalGETHandler(nutritionixAPIClient), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/food_log", handlers.FoodLogGETHandler(foodLogRepo, settingsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/refresh_totals", handlers.FoodLogRefreshTotalsGETHandler(foodLogRepo, settingsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/food_log", handlers.FoodLogPOSTHandler(foodLogRepo, settingsRepo), customMiddleware.AuthMiddleware(sessionsRepo))
	e.DELETE("/food_log/:id", handlers.FoodLogDELETEHandler(foodLogRepo), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/login", handlers.LoginGETHandler(authService))
	e.POST("/login", handlers.LoginPOSTHandler(authService))

	e.POST("/test", handlers.TestPOSTHandler())

	e.Logger.Fatal(e.Start(":1323"))
}
