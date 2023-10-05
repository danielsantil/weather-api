package handlers

import (
	"github.com/danielsantil/weather-api/internal/models"
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"net/http"
)

type GetWeatherRequest struct {
	Id string `uri:"id" binding:"uuid"`
}

type GetWeatherHistoryRequest struct {
	CityId string `uri:"cityId" binding:"required"`
	Limit  int    `uri:"limit" binding:"required"`
}

// HandlerGetWeather returns weather data for a specific weather id.
//
// If an error occurs while retrieving from the database, a 4xx status code is returned
func (inj *Injector) HandlerGetWeather(c *gin.Context) {
	var req GetWeatherRequest
	if err := c.ShouldBindUri(&req); err != nil {
		returnError(c, err)
		return
	}

	var weather database.Weather
	id, _ := uuid.Parse(req.Id)
	err := inj.DB.Preload(clause.Associations).First(&weather, id).Error

	if err != nil {
		returnError(c, err)
		return
	}

	c.JSON(http.StatusOK, weather)
}

// HandlerGetWeatherHistory returns weather history data for a city id, ordered by date.
//
// If an error occurs while retrieving from the database, a 4xx status code is returned
func (inj *Injector) HandlerGetWeatherHistory(c *gin.Context) {
	var req GetWeatherHistoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		returnError(c, err)
		return
	}

	var weatherData []database.Weather
	inj.DB.Preload(clause.Associations).
		Limit(req.Limit).
		Where("city_id = ?", req.CityId).
		Order("date_utc_millis DESC").
		Find(&weatherData)

	c.JSON(http.StatusOK, weatherData)
}

// HandlerGetLastWeather returns last weather data for a city id.
//
// If an error occurs while retrieving from the database, a 4xx status code is returned
func (inj *Injector) HandlerGetLastWeather(c *gin.Context) {
	cityId := c.Param("cityId")
	var lastWeather database.Weather
	err := inj.DB.Preload(clause.Associations).
		Where("city_id = ?", cityId).
		Order("date_utc_millis DESC").
		Last(&lastWeather).Error

	if err != nil {
		returnError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.MapToWeatherSummary(lastWeather))
}
