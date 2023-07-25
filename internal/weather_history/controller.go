package weatherhistory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	response "github.com/ssonumkar/weather-report-api/internal/http_response"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

type WeatherHistoryController struct {
	weatherHistoryService IWeatherHistoryService
	logger log.CustomLogger
}

func NewWeatherHistoryController(weatherHistoryService IWeatherHistoryService, logger log.CustomLogger) *WeatherHistoryController {
	return &WeatherHistoryController{weatherHistoryService, logger}
}
func (c *WeatherHistoryController) AddSearchedWeather(w http.ResponseWriter, r *http.Request){
	c.logger.UpdateEndpoint(log.Weather_Hist_Post)
	c.logger.Info("---------------------------------------------------")

	var weather WeatherHistory 
	if r.Body == nil{
		c.logger.Error("No input data provided")
		response.RespondWithError(w, http.StatusBadRequest, "Body is empty")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil{
		c.logger.Error(fmt.Sprint("Error decoding body into weather: ", err.Error()))
		response.RespondWithError(w, http.StatusBadRequest, "Incorrect data")
		return
	}
	c.logger.Info(fmt.Sprint("Request params received: weather: ", weather))

	err = c.weatherHistoryService.InsertWeatherSearch(weather, c.logger)
	if err != nil{
		response.RespondWithError(w, http.StatusInternalServerError, "Incorrect data")
		return
	}
	c.logger.Info("Weather search POST Successful")

	response.RespondWithJSON(w, http.StatusAccepted, map[string]string{
		"message": "Record inserted successfully",
	})
}
func (c *WeatherHistoryController) DeleteWeatherHistory(w http.ResponseWriter, r *http.Request) {
	c.logger.UpdateEndpoint(log.Weather_Hist_Delete)
	c.logger.Info("---------------------------------------------------")

	vars := mux.Vars(r)
	historyID := vars["history_id"]
	c.logger.Info(fmt.Sprint("Request Param received: HistoryID:", historyID))
	if historyID == ""{
		c.logger.Error("History ID cannot be empty")
		response.RespondWithError(w, http.StatusBadRequest, "History ID cannot be empty")
		return
	}
	err := c.weatherHistoryService.DeleteWeatherHistory(historyID, c.logger)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error deleting history: %s", err.Error()))
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.logger.Info("Weather search Delete Successful")
	response.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Weather history deleted",
	})
}

// DeleteWeatherHistory handles the request to delete a weather history entry
func (c *WeatherHistoryController) BulkDeleteWeatherHistory(w http.ResponseWriter, r *http.Request) {
	c.logger.UpdateEndpoint(log.Weather_Hist_Bulk_Delete)
	c.logger.Info("---------------------------------------------------")
	
	var historyIDs BulkDeleteWeatherHistory
	err :=  json.NewDecoder(r.Body).Decode(&historyIDs)
	if err != nil {
		c.logger.Error(fmt.Sprint("Failed to decode bulk weather history:", err))
		response.RespondWithError(w, http.StatusBadRequest, "Incorrect data")
		return
	}
	
	c.logger.Info(fmt.Sprint("Request params: HistoryIDs: ",historyIDs))
	if len(historyIDs.Ids) == 0 {
		c.logger.Error("History IDs cannot be empty")
		response.RespondWithError(w, http.StatusBadRequest, "Incorrect data")
		return
	}
	err = c.weatherHistoryService.BulkDeleteWeatherHistory(historyIDs, c.logger)
	if err != nil {
		c.logger.Error(err.Error())
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.logger.Info("Weather search Bulk Delete Successful")
	response.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Weather searched bulk deleted",
	})
}

func (c *WeatherHistoryController) GetWeatherSearchHistory(w http.ResponseWriter, r *http.Request){
	c.logger.UpdateEndpoint(log.Weather_Hist_Get)
	c.logger.Info("-------------------------------------------------")
	vars := mux.Vars(r)
	userID := vars["user_id"]
	c.logger.Debug(fmt.Sprintf("User ID: %s",userID))
	if userID == ""{
		c.logger.Error("History ID cannot be empty")
		response.RespondWithError(w, http.StatusBadRequest, "User ID cannot be empty")
		return
	}
	history, err := c.weatherHistoryService.GetWeatherSearchHistory(userID, c.logger)
	if err != nil{
		c.logger.Error(err.Error())
		response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.logger.Info("Weather search Get successfull")
	response.RespondWithJSON(w, http.StatusAccepted, history)
}
