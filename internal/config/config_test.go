package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T){
	wd, _ := os.Getwd()
	tmpFile, err := ioutil.TempFile(wd, "test-config-*.yaml")
	if err != nil {
		t.Fatal("Failed to create temporary file:", err)
	}
	defer os.Remove(tmpFile.Name())
	s := `server:
	host: localhost
	port: 8082
  database:
	driver: mysql
	connectionString: "sagar:password@tcp(localhost:3306)/weather_report"
  jwtSecret: YOUR_JWT_SECRET_KEY
   `
   _, err = tmpFile.WriteString(s)
   if err != nil {
	t.Fatal("Failed to write temporary file:", err)
	}
	
	
	expectedCfg := Config{
			ServerConfig{"localhost", "8082"}, 
			DatabaseConfig{"mysql", "sagar:password@tcp(localhost:3306)/weather_report"}, 
			"YOUR_JWT_SECRET_KEY",
		}
 	
	cfg, err := LoadConfig(tmpFile.Name(), "yaml", wd)
	if err != nil{
		t.Errorf(err.Error())
	}
	
	if cfg.Database.ConnectionString != expectedCfg.Database.ConnectionString {
		t.Errorf("Expected Connection String %s, got %s", expectedCfg.Database.ConnectionString, cfg.Database.ConnectionString)
	}
	if cfg.Database.Driver != expectedCfg.Database.Driver {
		t.Errorf("Expected Driver %s, got %s", expectedCfg.Database.Driver, cfg.Database.Driver)
	}
	if cfg.Server.Host != expectedCfg.Server.Host {
		t.Errorf("Expected Connection StringServer host %s, got %s", expectedCfg.Server.Host, cfg.Server.Host)
	}
	if cfg.Server.Port != expectedCfg.Server.Port {
		t.Errorf("Expected Server port %s, got %s", expectedCfg.Server.Port, cfg.Server.Port)
	}
	if cfg.JWTSecret != expectedCfg.JWTSecret {
		t.Errorf("Expected JWTSecret %s, got %s", expectedCfg.JWTSecret, cfg.JWTSecret)
	}
}