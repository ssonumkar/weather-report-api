package weather

import (
	"fmt"
	"net/http"
	"strconv"

	response "github.com/ssonumkar/weather-report-api/internal/http_response"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

// WeatherController handles weather-related requests
type WeatherController struct {
	weatherService IWeatherService
	logger log.CustomLogger
}

// NewWeatherController creates a new instance of WeatherController
func NewWeatherController(weatherService IWeatherService, logger log.CustomLogger) *WeatherController {
	return &WeatherController{weatherService, logger}
}

// GetCurrentWeather handles the request to get the current weather
func (c *WeatherController) GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	c.logger.UpdateEndpoint(log.Weather_Get)
	c.logger.Info("-----------------------------------------------------")
	lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
	if err != nil{
		c.logger.Debug(fmt.Sprint("cannot this convert string to float: ", r.FormValue("lat")))
		response.RespondWithError(w, http.StatusBadRequest, "Incorrect value for lattitude.")
		return
	}
	lon, err := strconv.ParseFloat(r.FormValue("lon"), 64)
	if err != nil{
		c.logger.Debug(fmt.Sprint("cannot this convert string to float: ", r.FormValue("lat")))
		response.RespondWithError(w, http.StatusBadRequest, "Incorrect value for lattitude.")
		return
	}
	city := r.FormValue("city")
	apiKey := r.FormValue("apiKey")
	c.logger.Info(fmt.Sprintf("Request params are: city: %s, lat: %f, lon: %f, apiKey: %s", city, lat, lon, apiKey))
	weather, err := c.weatherService.GetCurrentWeather(city, lat, lon, apiKey, c.logger)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.logger.Info("Weather Data Get successfull")
	response.RespondWithJSON(w, http.StatusOK, weather)
}

