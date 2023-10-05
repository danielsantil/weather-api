package handlers

import (
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerGetForecast returns forecast data for a city id
//
// If an error occurs while retrieving from the database, a 4xx status code is returned
func (inj *Injector) HandlerGetForecast(c *gin.Context) {
	cityId := c.Param("cityId")
	var forecast []database.CityForecast
	err := inj.DB.
		Preload("Forecasts.Main").Preload("Forecasts.Weather").
		Find(&forecast, cityId).Error

	if err != nil {
		returnError(c, err)
		return
	}

	c.JSON(http.StatusOK, forecast)
}
