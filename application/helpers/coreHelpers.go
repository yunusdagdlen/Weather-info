package helpers

import (
	"encoding/json"
	"io/ioutil"
	"main/application/models"
	"net/http"
	"strings"
)

func loadApiConf(filename string) (models.ApiConfig, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return models.ApiConfig{}, err
	}
	var c models.ApiConfig
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return models.ApiConfig{}, err
	}
	return c, nil
}

func WeatherApiQuery(user_ip string, city string) (models.WeatherData, error) {
	resp, err := http.Get("http://api.weatherstack.com/current?access_key=abdefe66d5ea2a236eb6d91a44f8fdc2&query=" + city)
	defer resp.Body.Close()
	var weatherData models.WeatherData
	if err != nil {
		//log basıp boi+ş data ve error dönmek doğru yaklaşım
		panic(err.Error())
	}
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		panic(err.Error())
	}
	if weatherData.Current.WeatherDescriptions != nil {
		_ = InsertEmployee(user_ip, city, strings.Join(weatherData.Current.WeatherDescriptions, ", "))
	}
	return weatherData, err
}
