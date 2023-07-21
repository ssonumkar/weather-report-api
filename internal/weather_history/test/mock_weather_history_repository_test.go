package weatherhistory

import (
	"fmt"
	"time"

	"github.com/ssonumkar/weather-report-api/internal/log"
	weatherhistory "github.com/ssonumkar/weather-report-api/internal/weather_history"
)

type MockSuccessWeatherHistoryRepository struct {
}

func (m *MockSuccessWeatherHistoryRepository) InsertWeatherSearch(weather weatherhistory.WeatherHistory, logger log.CustomLogger) error  {
	return nil
}
func (m *MockSuccessWeatherHistoryRepository) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error  {
	return nil
}
func (m *MockSuccessWeatherHistoryRepository) GetWeatherSearchHistory(userID string, logger log.CustomLogger) ([]weatherhistory.WeatherHistory, error) {
	var list []weatherhistory.WeatherHistory
	list = append(list, weatherhistory.WeatherHistory{Id: 1, UserId: 2, City: "pune", TempMin: 123, TempMax: 234, FeelsLike: 122, CreatedAt: time.Now()})
	list = append(list, weatherhistory.WeatherHistory{Id: 2, UserId: 12, City: "mumbai", TempMin: 1123, TempMax: 2234, FeelsLike: 1122, CreatedAt: time.Now()})
	return list, nil
}
func (m *MockSuccessWeatherHistoryRepository) BulkDeleteWeatherHistory(historyIDs weatherhistory.BulkDeleteWeatherHistory, logger log.CustomLogger) error {
	return nil
}

type MockFailWeatherHistoryRepository struct {
}
func (m *MockFailWeatherHistoryRepository) InsertWeatherSearch(weather weatherhistory.WeatherHistory, logger log.CustomLogger) error  {
	return fmt.Errorf("")
}
func (m *MockFailWeatherHistoryRepository) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error  {
	
	return fmt.Errorf("")

}
func (m *MockFailWeatherHistoryRepository) GetWeatherSearchHistory(userID string, logger log.CustomLogger) ([]weatherhistory.WeatherHistory, error) {
	return nil, fmt.Errorf("")
}
func (m *MockFailWeatherHistoryRepository) BulkDeleteWeatherHistory(historyIDs weatherhistory.BulkDeleteWeatherHistory, logger log.CustomLogger) error  {
	return fmt.Errorf("")
}