package log

import (
	"log"
	"os"
)

type CustomLogger struct {
	logger *log.Logger
	endpoint Endpoint
} 
func NewCustomLogger() *CustomLogger{
	return &CustomLogger{
		logger:   log.New(os.Stdout, "", log.LstdFlags),
	}
}
func (c *CustomLogger) UpdateEndpoint(endpoint Endpoint){
	c.endpoint = endpoint
}
func (c *CustomLogger) logEntry(logLevel string, message string ){
	c.logger.Printf("%s %s %s %s\n", logLevel, c.endpoint.Name, c.endpoint.Method, message )
}
func (c *CustomLogger) Debug(message string){
	c.logEntry("DEBUG", message)
}

func (c *CustomLogger) Info(message string){
	c.logEntry("INFO", message )
}

func (c *CustomLogger) Error(message string){
	c.logEntry("DEBUG", message )
}

func (c *CustomLogger) Fatal(message string){
	c.logEntry("ERROR", message)
	os.Exit(1)
}