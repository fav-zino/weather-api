package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"weather_api/handler"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()

	router.GET("/weather/:city", handler.GetWeatherHandler)

	if err := router.Run(":8080"); err != nil {
		log.Panicf("Error starting server: %s", err)
	}
}
