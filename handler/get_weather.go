package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"weather_api/api"
)

func GetWeatherHandler(c *gin.Context) {
    city := c.Param("city")

  
    weather, err := api.GetWeather(city)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
        return
    }

    c.JSON(http.StatusOK, weather)
}


func getWeather(city string){

}
