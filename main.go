package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/rs/cors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ssonumkar/weather-report-api/internal/config"
	custom_log "github.com/ssonumkar/weather-report-api/internal/log"
	"github.com/ssonumkar/weather-report-api/internal/routes"
)

func main() {

	logger := custom_log.NewCustomLogger()
	cfg, err := config.LoadConfig("config", "yaml", path.Join("internal", "config"))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load config: %s", err.Error()))
	}

	db, err := createDatabaseConnection(cfg.Database)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to connect to the database: %s", err.Error()))
	}
	defer db.Close()
	router := routes.SetupRoutes(db, cfg.JWTSecret, *logger)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	logger.Info(fmt.Sprintf("Server is running on %s", addr))
	log.Fatal(http.ListenAndServe(addr, handler))
}

func createDatabaseConnection(databaseConfig config.DatabaseConfig) (*sql.DB, error) {

	db, err := sql.Open(databaseConfig.Driver, databaseConfig.ConnectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Databse connected..!")
	return db, nil
}
