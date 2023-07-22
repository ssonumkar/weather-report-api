package weatherhistory

import "github.com/ssonumkar/weather-report-api/internal/log"

type WeatherHistoryService struct {
	repo IWeatherHistoryRepository
}

func NewWeatherHistoryService(repo IWeatherHistoryRepository) *WeatherHistoryService {
	return &WeatherHistoryService{repo}
}

func (s *WeatherHistoryService) InsertWeatherSearch(weather WeatherHistory, logger log.CustomLogger) (error){
	return s.repo.InsertWeatherSearch(weather, logger)
}
func (s *WeatherHistoryService) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error {
	return s.repo.DeleteWeatherHistory(historyID, logger)
}

func (s *WeatherHistoryService) GetWeatherSearchHistory(userID string, logger log.CustomLogger) ([]WeatherHistory, error){
	return s.repo.GetWeatherSearchHistory(userID, logger)
}

func (s *WeatherHistoryService)	BulkDeleteWeatherHistory(historyIDs BulkDeleteWeatherHistory, logger log.CustomLogger) error{
	return s.repo.BulkDeleteWeatherHistory(historyIDs, logger)
}
