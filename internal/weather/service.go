package weather

import (
	"encoding/json"
	"fmt"

	"github.com/ssonumkar/weather-report-api/internal/config"
	"github.com/ssonumkar/weather-report-api/internal/log"
	"github.com/ssonumkar/weather-report-api/internal/utils"
)

// WeatherService handles weather-related operations
type WeatherService struct {
	// Add dependencies as needed
}

// NewWeatherService creates a new instance of WeatherService
func NewWeatherService() *WeatherService {
	return &WeatherService{}
}


// GetCurrentWeather gets the current weather for the specified city
func (s *WeatherService) GetCurrentWeather(city string, lat float64, lon float64, apiKey string, logger log.CustomLogger) (*Weather, error) {
	urlCfg, err := config.LoadUrls()
	if err != nil{
		return nil, err
	}
	
	weather, err := getWeather(urlCfg, lat, lon, apiKey, logger)	
	if err != nil{
		return nil, err
	}
	weather.City = city
	logger.Debug("Weather data received successfully")
	return weather, nil
}
func getWeather(urlCfg *config.UrlConfig, lat float64, lon float64, apiKey string, logger log.CustomLogger) (*Weather, error){
	updateWeatherApiConfig(urlCfg, lat, lon, apiKey)
	logger.Debug(fmt.Sprintf("Url configs are: %v", urlCfg.UrlList["weather_report"]))
	resp, err := utils.CallApi(urlCfg.UrlList["weather_report"], logger)
	if err != nil{
		return nil, err
	}

	var weather Weather

	err = json.Unmarshal(resp, &weather)
	if err != nil{
		logger.Error(fmt.Sprint("Error unmarshalling weather api response: ", err.Error()))
		return nil, err
	}
	logger.Debug(fmt.Sprintf("Api response is: %v", weather))
	return &weather, nil
}

func updateWeatherApiConfig(url *config.UrlConfig, lat float64, lon float64, apiKey string) {
	url.UrlList["weather_report"].Parameters["lat"] = lat
	url.UrlList["weather_report"].Parameters["lon"] = lon
	url.UrlList["weather_report"].Parameters["appid"] = apiKey
}