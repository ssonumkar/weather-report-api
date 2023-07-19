package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ssonumkar/weather-report-api/internal/config"
	"github.com/ssonumkar/weather-report-api/internal/encrypt"
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

	encrypt.InitTokenPool()

	router := routes.SetupRoutes(db, cfg.JWTSecret, *logger)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	logger.Info(fmt.Sprintf("Server is running on %s", addr))
	log.Fatal(http.ListenAndServe(addr, router))
}

func createDatabaseConnection(databaseConfig config.DatabaseConfig) (*sql.DB, error) {

	db, err := sql.Open(databaseConfig.Driver, databaseConfig.ConnectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Databse connected..!")
	return db, nil
}
