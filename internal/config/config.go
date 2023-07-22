package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	JWTSecret string
}

type ServerConfig struct {
	Host string
	Port string
}


type DatabaseConfig struct {
	Driver           string
	ConnectionString string
}

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

	if cfg.JWTSecret == "" {
		cfg.JWTSecret = os.Getenv("JWT_SECRET")
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT secret key is not set")
	}

	return &cfg, nil
}
