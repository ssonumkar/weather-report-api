package weatherhistory

import (
	"testing"
	"time"

	"github.com/ssonumkar/weather-report-api/internal/log"
	weatherhistory "github.com/ssonumkar/weather-report-api/internal/weather_history"
)

var weatherSuccessHistoryService weatherhistory.WeatherHistoryService
var weatherFailHistoryService weatherhistory.WeatherHistoryService
var logger log.CustomLogger

func init(){
	weatherSuccessHistoryService = *weatherhistory.NewWeatherHistoryService(&MockSuccessWeatherHistoryRepository{})
	weatherFailHistoryService = *weatherhistory.NewWeatherHistoryService(&MockFailWeatherHistoryRepository{})
	logger = *log.NewCustomLogger()

}

func TestShouldSucceedForInsertWeatherSearch(t *testing.T){
	//Given
	weather := weatherhistory.WeatherHistory{Id: 1, UserId: 2, City: "pune", TempMin: 12, TempMax: 23, FeelsLike: 12, CreatedAt: time.Now()}
	//When
	err := weatherSuccessHistoryService.InsertWeatherSearch(weather, logger)
	//Then
	if err != nil{
		t.Errorf("expected no err but got %v", err.Error())
	}
}
func TestShouldSucceedForDeleteWeatherHistory(t *testing.T){
	//Given
	histId := "12"
	//When
	err := weatherSuccessHistoryService.DeleteWeatherHistory(histId, logger)
	//Then
	if err != nil{
		t.Errorf("expected no err but got %v", err.Error())
	}
}
func TestShouldSucceedForGetWeatherSearchHistory(t *testing.T){
	//Given
	var expected []weatherhistory.WeatherHistory
	expected = append(expected, weatherhistory.WeatherHistory{Id: 1, UserId: 2, City: "pune", TempMin: 123, TempMax: 234, FeelsLike: 122, CreatedAt: time.Now()})
	expected = append(expected, weatherhistory.WeatherHistory{Id: 2, UserId: 12, City: "mumbai", TempMin: 1123, TempMax: 2234, FeelsLike: 1122, CreatedAt: time.Now()})

	//When
	actual, err := weatherSuccessHistoryService.GetWeatherSearchHistory(logger)
	if err != nil{
		t.Fatal("Could not get history data")
	}
	//Then

	for i := 0; i<len(expected);i++ {
		if expected[i] != actual[i]{
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}
func TestShouldSucceedForBulkDeleteWeatherHistory(t *testing.T){
	//Given
	historyId := "2"
	//When
	err := weatherSuccessHistoryService.DeleteWeatherHistory(historyId, logger)
	//Then
	if err != nil{
		t.Errorf("expected no err but got %v", err.Error())
	}
}

func TestShouldFailForrInsertWeatherSearch(t *testing.T){
	//Given
	var weather weatherhistory.WeatherHistory
	//When
	err := weatherFailHistoryService.InsertWeatherSearch(weather, logger)
	//Then
	if err == nil{
		t.Errorf("expected error but did not get any")
	}
}
func TestShouldFailForrDeleteWeatherHistory(t *testing.T){
	//Given
	histId := "12"
	//When
	err := weatherSuccessHistoryService.DeleteWeatherHistory(histId, logger)
	//Then
	if err != nil{
		t.Errorf("expected no err but got %v", err.Error())
	}
}
func TestShouldFailForrGetWeatherSearchHistory(t *testing.T){
	//Given
	var expected []weatherhistory.WeatherHistory
	expected = append(expected, weatherhistory.WeatherHistory{Id: 1, UserId: 2, City: "pune1", TempMin: 123, TempMax: 234, FeelsLike: 122, CreatedAt: time.Now()})
	expected = append(expected, weatherhistory.WeatherHistory{Id: 2, UserId: 12, City: "mumbai1", TempMin: 1123, TempMax: 2234, FeelsLike: 1122, CreatedAt: time.Now()})

	//When
	actual, err := weatherSuccessHistoryService.GetWeatherSearchHistory(logger)
	if err != nil{
		t.Fatal("Could not get history data")
	}
	//Then

	for i := 0; i<len(expected);i++ {
		if expected[i] == actual[i]{
			t.Errorf("%v should not match with %v", expected, actual)
		}
	}
}
func TestShouldFailForrBulkDeleteWeatherHistory(t *testing.T){
	//Given
	historyId := "2"
	//When
	err := weatherSuccessHistoryService.DeleteWeatherHistory(historyId, logger)
	//Then
	if err == nil{
		t.Errorf("expected err but got nil")
	}
}