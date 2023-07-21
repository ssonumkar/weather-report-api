package weatherhistory

import "github.com/ssonumkar/weather-report-api/internal/log"

// WeatherHistoryService handles weather history-related operations
type WeatherHistoryService struct {
	repo IWeatherHistoryRepository
}

// NewWeatherHistoryService creates a new instance of WeatherHistoryService
func NewWeatherHistoryService(repo IWeatherHistoryRepository) *WeatherHistoryService {
	return &WeatherHistoryService{repo}
}

func (s *WeatherHistoryService) InsertWeatherSearch(weather WeatherHistory, logger log.CustomLogger) (error){
	return s.repo.InsertWeatherSearch(weather, logger)
}
// DeleteWeatherHistory deletes a weather history entry
func (s *WeatherHistoryService) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error {
	// Call the repository to delete the weather history
	return s.repo.DeleteWeatherHistory(historyID, logger)
}

func (s *WeatherHistoryService) GetWeatherSearchHistory(userID string, logger log.CustomLogger) ([]WeatherHistory, error){
	return s.repo.GetWeatherSearchHistory(userID, logger)
}

func (s *WeatherHistoryService)	BulkDeleteWeatherHistory(historyIDs BulkDeleteWeatherHistory, logger log.CustomLogger) error{
	return s.repo.BulkDeleteWeatherHistory(historyIDs, logger)
}
