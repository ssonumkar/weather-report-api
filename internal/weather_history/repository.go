package weatherhistory

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ssonumkar/weather-report-api/internal/log"
)

type WeatherHistoryRepository struct {
	db *sql.DB
}

func NewWeatherHistoryRepository(db *sql.DB) *WeatherHistoryRepository {
	return &WeatherHistoryRepository{db}
}

func (r *WeatherHistoryRepository) InsertWeatherSearch(weather WeatherHistory, logger log.CustomLogger) error{
	if weather.City == ""{
		return fmt.Errorf("city cannot be nil")
	}
	stmt, err := r.db.Prepare("CALL AddWeatherHistory(?,?,?,?,?,?)")
	
	if err != nil{
		logger.Error(fmt.Sprint("Error executiong prepare : ", err.Error()))
		return fmt.Errorf("failed to prepare: %s", err.Error())
	}
	logger.Debug("Prepare statement successful")

	golangDateTime := time.Now().Format("2006-01-02 15:04:05")
	_ , err = stmt.Exec(weather.UserId, weather.City, weather.TempMin, weather.TempMax, weather.FeelsLike, golangDateTime)
	defer stmt.Close()
	if err != nil{
		logger.Error(fmt.Sprint("Error executiong Exec: ", err.Error()))
		return err
	}
	logger.Debug("Weather search data inserted successfully in database")
	return nil
}
func (r *WeatherHistoryRepository) DeleteWeatherHistory(historyID string, logger log.CustomLogger) error {
	stmt, err := r.db.Prepare("CALL DeleteWeatherHistory(?)");
	
	if err != nil {
		logger.Error(fmt.Sprint("Error executiong prepare: ", err.Error()))
		return err
	}
	logger.Debug("Prepare statement successful")
	defer stmt.Close()
	_, err = stmt.Exec(historyID)
	if err != nil {
		logger.Error(fmt.Sprint("Error executiong DeleteWeatherHistory procedure: ", err.Error()))
		return err
	}
	logger.Debug("Weather search data deleted successfully from database")
	return nil
}

func (r *WeatherHistoryRepository) GetWeatherSearchHistory(userID string, logger log.CustomLogger) ([]WeatherHistory, error){
	query := fmt.Sprintf("SELECT * FROM WeatherHistory where user_id = %s", userID)
	rows, err := r.db.Query(query)
	logger.Debug(query)
    if err != nil {
		logger.Error(fmt.Sprint("Error executiong select query: ", err.Error()))
        return nil, err
    }
    defer rows.Close()

    var weatherSerchHistory []WeatherHistory

    for rows.Next() {
        var search WeatherHistory
        if err := rows.Scan(&search.Id, &search.UserId, &search.City, &search.TempMin, &search.TempMin, &search.FeelsLike, &search.CreatedAt); err != nil {
			logger.Error(fmt.Sprint("Error scanning rows: ", err.Error()))
			return weatherSerchHistory, err
        }
        weatherSerchHistory = append(weatherSerchHistory, search)
    }
    if err = rows.Err(); err != nil {
		logger.Error(fmt.Sprint("Error reading rows: ", err.Error()))
        return weatherSerchHistory, fmt.Errorf("error reading rows: %s", err.Error())
    }
	logger.Debug("Weather search data fetched successfully from database")
    return weatherSerchHistory, nil
}
func (r *WeatherHistoryRepository)	BulkDeleteWeatherHistory(historyIDs BulkDeleteWeatherHistory, logger log.CustomLogger) error{
	
	query := "Delete from WeatherHistory where id in"
	inClause := ""
	for _, v := range historyIDs.Ids {
		inClause = fmt.Sprintf("%s, %d ", inClause, v)
	}
	query = fmt.Sprintf("%s(%s)", query, inClause[2:])
	logger.Debug(fmt.Sprintf("Prepared query: %s",query))

	stmt, err := r.db.Prepare(query);
	if err != nil {
		logger.Error(fmt.Sprint("Error preparing query: ", err.Error()))
		return err
	}
	logger.Debug("Prepare statement successful")
	_, err = stmt.Exec()
	if err != nil {
		logger.Error(fmt.Sprint("Error executing prepared query: ", err.Error()))
		return err
	}
	logger.Debug("Weather search bulk data deleted successfully from database")
	return nil
}
