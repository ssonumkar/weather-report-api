package config

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)
type UrlConfig struct{
	UrlList   map[string]Url
}

type Url struct {
	UrlString  string
	Parameters map[string]interface{}
	Method     string
	RequestBody string
}

func LoadUrls() (*UrlConfig, error){

	viper.SetConfigName("url")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join("internal", "config"))
	
	err := viper.ReadInConfig()
	if err != nil{
		fmt.Println("Error loading url config, ", err)
		return nil, err
	}

	var urlConfig UrlConfig
	
	err = viper.Unmarshal(&urlConfig)
	if err != nil{
		fmt.Println("Error loading url config, ", err)
		return nil, err
	}

	return &urlConfig, nil

}
