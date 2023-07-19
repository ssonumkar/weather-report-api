package weather

import (
	"github.com/ssonumkar/weather-report-api/internal/log"
	"github.com/ssonumkar/weather-report-api/internal/weather"
)

type MockSuccessWeatherService struct {
}

func (s *MockSuccessWeatherService) GetCurrentWeather(city string, lat float64, lon float64, apiKey string, logger log.CustomLogger) (*weather.Weather, error) {
	return &weather.Weather{}, nil
}

type MockFailWeatherService struct {
}

func (s *MockFailWeatherService) GetCurrentWeather(city string, lat float64, lon float64, apiKey string, logger log.CustomLogger) (*weather.Weather, error) {
	return &weather.Weather{}, nil
}