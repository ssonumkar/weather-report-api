package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ssonumkar/weather-report-api/internal/config"
	"github.com/ssonumkar/weather-report-api/internal/log"
)

func CallApi(url config.Url, logger log.CustomLogger) ([]byte, error){
	apiUrl := buildUrl(url);
	logger.Debug(fmt.Sprintf("Url received is: %s and Method is %s", url.UrlString, url.Method))
	var resp *http.Response;
	var err error = nil
	
	switch url.Method {
	case "GET", "get", "Get":
		resp, err = http.Get(apiUrl)
	case "POST", "post", "Post":
		resp, err = http.Post(apiUrl, "application/json", bytes.NewBuffer([]byte(url.RequestBody)))
	default: 
		logger.Error(fmt.Sprint("Invalid HTTP Method: ", url.Method))
		return nil, fmt.Errorf("invalid HTTP Method: %v", url.Method)
	}
	if err != nil{
		logger.Error(fmt.Sprint("Failed to call api: ", err.Error()))
		return nil, err
	}
	
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	
	if err != nil{
		logger.Error(fmt.Sprint("Error reading response: ", err.Error()))
		return nil, err
	}
	return responseBody, nil
}
func buildUrl(url config.Url) (string) {
	queryParams := buildQueryParams(url.Parameters)

	apiUrl := fmt.Sprintf("%v?%v", url.UrlString, queryParams)
	fmt.Println("Api Url is : ", apiUrl)
	return apiUrl
}
func buildQueryParams(parameters map[string]interface{}) (string){
	params := ""
	// fmt.Println(parameters)
	for k, v := range(parameters) {
		params = fmt.Sprintf("%v&%v=%v",params, k, v)
	}
	return params[1:]
}