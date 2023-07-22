package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	auth "github.com/ssonumkar/weather-report-api/internal/auth"
	"github.com/ssonumkar/weather-report-api/internal/encrypt"
	"github.com/ssonumkar/weather-report-api/internal/log"
	"github.com/ssonumkar/weather-report-api/internal/middleware"
	"github.com/ssonumkar/weather-report-api/internal/weather"
	weatherhistory "github.com/ssonumkar/weather-report-api/internal/weather_history"
)

func SetupRoutes(db *sql.DB, JWTSecretKey string, logger log.CustomLogger) *mux.Router {
	router := mux.NewRouter()

	tokenPool := encrypt.NewAuthTokenPool()
	passwordManager := encrypt.NewPasswordManager()
	userRepository := auth.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository, tokenPool, passwordManager, JWTSecretKey)
	weatherService := weather.NewWeatherService()
	weatherHistoryRepository := weatherhistory.NewWeatherHistoryRepository(db) 
	weatherHistoryService := weatherhistory.NewWeatherHistoryService(weatherHistoryRepository)

	authController := auth.NewAuthController(authService, logger)
	weatherController := weather.NewWeatherController(weatherService, logger)
	weatherHistoryController := weatherhistory.NewWeatherHistoryController(weatherHistoryService, logger)

	_middleware := middleware.NewAuthMiddleware(JWTSecretKey)
	router.HandleFunc("/api/login", authController.Login).Methods("POST")
	router.HandleFunc("/api/logout", _middleware.Authenticate(tokenPool, logger, authController.Logout)).Methods("POST")
	router.HandleFunc("/api/register", authController.Register).Methods("POST")
	router.HandleFunc("/api/weather", _middleware.Authenticate(tokenPool, logger, weatherController.GetCurrentWeather)).Methods("GET")
	router.HandleFunc("/api/weather/history/{user_id}", _middleware.Authenticate(tokenPool, logger, weatherHistoryController.GetWeatherSearchHistory)).Methods("GET")
	router.HandleFunc("/api/weather/history", _middleware.Authenticate(tokenPool, logger, weatherHistoryController.AddSearchedWeather)).Methods("POST")
	router.HandleFunc("/api/weather/history", _middleware.Authenticate(tokenPool, logger, weatherHistoryController.BulkDeleteWeatherHistory)).Methods("PATCH")
	router.HandleFunc("/api/weather/history/{history_id}", _middleware.Authenticate(tokenPool, logger, weatherHistoryController.DeleteWeatherHistory)).Methods("DELETE")

	return router
}
