package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmitkov28/dietapp/internal/diet"
	"github.com/dmitkov28/dietapp/internal/handlers"
	"github.com/dmitkov28/dietapp/internal/httputils"
	customMiddleware "github.com/dmitkov28/dietapp/internal/middleware"
	"github.com/dmitkov28/dietapp/internal/repositories"
	"github.com/dmitkov28/dietapp/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {

	db, err := repositories.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// repositories
	usersRepo := repositories.NewUsersRepository(db)
	sessionsRepo := repositories.NewSessionsRepository(db)
	settingsRepo := repositories.NewSettingsRepository(db)
	measurementsRepo := repositories.NewMeasurementsRepository(db)
	foodLogRepo := repositories.NewFoodLogsRepository(db)

	httpClient := http.Client{}
	apiClient := httputils.NewAPIClient(&httpClient)
	nutritionixAPIClient, err := diet.NewNutritionixAPIClient(apiClient)

	if err != nil {
		fmt.Println(err)
	}

	// services
	authService := services.NewAuthService(usersRepo, sessionsRepo)
	measurementsService := services.NewMeasurementsService(measurementsRepo)
	foodLogService := services.NewFoodLogService(foodLogRepo)
	settingsService := services.NewSettingsService(settingsRepo)
	chartService := services.NewChartService()

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

	e.GET("/dashboard", handlers.DashboardGETHandler(measurementsService, settingsService, chartService), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/settings", handlers.SettingsGETHandler(settingsService), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/settings", handlers.SettingsPOSTHandler(settingsService), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/weight", handlers.WeightGETHandler(), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/weight", handlers.WeightPOSTHandler(measurementsService), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/stats", handlers.StatsGETHandler(measurementsService), customMiddleware.AuthMiddleware(sessionsRepo))
	e.DELETE("/stats/:id", handlers.StatsDELETEHandler(measurementsService), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/calories", handlers.CaloriesGETHandler(), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/calories", handlers.CaloriesPOSTHandler(measurementsService), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/scan", handlers.ScanGETHandler(), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/scan/:ean", handlers.ScanBarCodeGETHandler(), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/search", handlers.SearchFoodGETHandler(foodLogService), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/search_food", handlers.SearchFoodGetHandlerWithParams(nutritionixAPIClient), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/search_food/modal", handlers.SearchFoodModalGETHandler(nutritionixAPIClient), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/food_log", handlers.FoodLogGETHandler(foodLogService), customMiddleware.AuthMiddleware(sessionsRepo))
	e.GET("/refresh_totals", handlers.FoodLogRefreshTotalsGETHandler(foodLogService), customMiddleware.AuthMiddleware(sessionsRepo))
	e.POST("/food_log", handlers.FoodLogPOSTHandler(foodLogService), customMiddleware.AuthMiddleware(sessionsRepo))
	e.DELETE("/food_log/:id", handlers.FoodLogDELETEHandler(foodLogService), customMiddleware.AuthMiddleware(sessionsRepo))

	e.GET("/login", handlers.LoginGETHandler(authService))
	e.POST("/login", handlers.LoginPOSTHandler(authService))

	e.Logger.Fatal(e.Start(":1323"))
}
