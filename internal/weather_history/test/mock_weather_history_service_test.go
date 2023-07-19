package weatherhistory

// import (
// 	"github.com/ssonumkar/weather-report-api/internal/weatherhistory"
// 	service "github.com/ssonumkar/weather-report-api/internal/service/test"
// 	"github.com/ssonumkar/weather-report-api/internal/utils"
// )

// type MockSuccessWeatherHistoryService struct {
// 	repo service.MockSuccessWeatherHistoryRepository
// }

// func (s *MockSuccessWeatherHistoryService) InsertWeatherSearch(weather weatherhistory.WeatherHistory, logger log.CustomLogger) (error){
// 	return s.repo.InsertWeatherSearch(weather, logger)
// }

// func (s *MockSuccessWeatherHistoryService) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error {
// 	return s.repo.DeleteWeatherHistory(historyID, logger)
// }

// func (s *MockSuccessWeatherHistoryService) GetWeatherSearchHistory(logger log.CustomLogger) ([]weatherhistory.WeatherHistory, error){
// 	return s.repo.GetWeatherSearchHistory(logger)
// }

// func (s *MockSuccessWeatherHistoryService)	BulkDeleteWeatherHistory(historyIDs weatherhistory.BulkDeleteWeatherHistory, logger log.CustomLogger) error{
// 	return s.BulkDeleteWeatherHistory(historyIDs, logger)
// }

// type MockFailWeatherHistoryService struct {
// 	repo service.MockFailWeatherHistoryRepository
// }

// func (s *MockFailWeatherHistoryService) InsertWeatherSearch(weather weatherhistory.WeatherHistory, logger log.CustomLogger) (error){
// 	return s.repo.InsertWeatherSearch(weather, logger)
// }

// func (s *MockFailWeatherHistoryService) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error {
// 	return s.repo.DeleteWeatherHistory(historyID, logger)
// }

// func (s *MockFailWeatherHistoryService) GetWeatherSearchHistory(logger log.CustomLogger) ([]weatherhistory.WeatherHistory, error){
// 	return s.repo.GetWeatherSearchHistory(logger)
// }

// func (s *MockFailWeatherHistoryService)	BulkDeleteWeatherHistory(historyIDs weatherhistory.BulkDeleteWeatherHistory, logger log.CustomLogger) error{
// 	return s.BulkDeleteWeatherHistory(historyIDs, logger)
// }