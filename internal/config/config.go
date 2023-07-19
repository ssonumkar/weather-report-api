package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	JWTSecret string
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Host string
	Port string
}


// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
	Driver           string
	ConnectionString string
	// Add more database-related configuration fields as needed
}

// LoadConfig loads the application configuration
func LoadConfig(fileName string, fileType string, parentDir string) (*Config, error) {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(parentDir)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	// Check if JWT secret key is set in environment variable
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = os.Getenv("JWT_SECRET")
	}

	// Validate JWT secret key
	if cfg.JWTSecret == "" {
		log.Fatal("JWT secret key is not set")
	}

	return &cfg, nil
}
