package weatherhistory

import (
	"time"

	"github.com/ssonumkar/weather-report-api/internal/log"
)

type WeatherHistory struct {
	Id        int64   `json:"id"`
	UserId    int64   `json:"user_id"`
	City      string  `json:"city"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	CreatedAt time.Time `json:"created_at"`
}

type BulkDeleteWeatherHistory struct{
	Ids []int `json:"ids"`
}

type IWeatherHistoryService interface{
	InsertWeatherSearch(weather WeatherHistory, logger log.CustomLogger) (error)
	DeleteWeatherHistory(historyID string, logger log.CustomLogger) error
	GetWeatherSearchHistory(logger log.CustomLogger)([]WeatherHistory, error)
	BulkDeleteWeatherHistory(historyIDs BulkDeleteWeatherHistory, logger log.CustomLogger) error
}

type IWeatherHistoryRepository interface{
	InsertWeatherSearch(weather WeatherHistory, logger log.CustomLogger) error
	DeleteWeatherHistory(historyID string, logger log.CustomLogger) error
	GetWeatherSearchHistory(logger log.CustomLogger) ([]WeatherHistory, error)
	BulkDeleteWeatherHistory(historyIDs BulkDeleteWeatherHistory, logger log.CustomLogger) error
}
