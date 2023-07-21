package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"weather_api/model"
)


func GetWeather(city string) ( []model.Weather, error) {
    apiKey := os.Getenv("WEATHER_API_KEY")
    url := fmt.Sprintf("https://api.weatherbit.io/v2.0/current?key=%s&city=%s", apiKey, city)

  
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()


    var data map[string]interface{}
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        return nil, err
    }

    weatherData := data["data"].([]interface{})


	var filteredweatherList []model.Weather


	for i := 0; i < len(weatherData); i++ {

		weather := model.Weather{
			City:       weatherData[i].(map[string]interface{})["city_name"].(string),
			CountryCode:       weatherData[i].(map[string]interface{})["country_code"].(string),
			Temperature: weatherData[i].(map[string]interface{})["temp"].(float64),
			Description: weatherData[i].(map[string]interface{})["weather"].(map[string]interface{})["description"].(string),
			Humidity:     weatherData[i].(map[string]interface{})["rh"].(float64),
			WindSpeed:     weatherData[i].(map[string]interface{})["wind_spd"].(float64),
		}
		filteredweatherList = append(filteredweatherList,weather)
		
	}


    return filteredweatherList, nil
}