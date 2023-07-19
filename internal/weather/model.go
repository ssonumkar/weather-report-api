package weather

import "github.com/ssonumkar/weather-report-api/internal/log"

// Weather represents weather conditions
type Weather struct {
	City        string    `json:"city"`
	Report Report `json:"main"`
}
type Report struct{
	TempMin float64   `json:"temp_min"`
	TempMax float64   `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
}
type Coordinates struct{
	Lat float64
	Lon float64
}

type CoordinateResp struct{
	Name string `json:"name"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	State string `json:"state"`
	Country string `json:"country"`
}

type IWeatherService interface{
	GetCurrentWeather(city string, lat float64, lon float64, apiKey string, logger log.CustomLogger) (*Weather, error)
}