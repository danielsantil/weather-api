package config

import (
	"fmt"
	"github.com/danielsantil/weather-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func AddRouter(port int, injector handlers.Injector) *gin.Engine {
	router := gin.Default()
	router.GET("health", handlers.HandlerHealth)

	router.GET("weather/:id", injector.HandlerGetWeather)
	router.GET("weather-last/:cityId", injector.HandlerGetLastWeather)
	router.GET("weather-history/:cityId/:limit", injector.HandlerGetWeatherHistory)

	router.GET("cities", injector.HandlerGetCities)

	serverErr := router.Run(fmt.Sprintf("localhost:%d", port))
	if serverErr != nil {
		log.Fatal(serverErr)
	}

	return router
}
