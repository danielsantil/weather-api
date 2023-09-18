package handlers

import (
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

func (inj *Injector) HandlerGetWeatherHistory(c *gin.Context) {
	cityId := c.Param("cityId")
	var weatherData []database.Weather
	inj.DB.Preload(clause.Associations).
		Where("city_id = ?", cityId).
		Order("date_utc_millis DESC").
		Find(&weatherData)

	c.JSON(http.StatusOK, weatherData)
}

func (inj *Injector) HandlerGetLastWeather(c *gin.Context) {
	cityId := c.Param("cityId")
	var lastWeather database.Weather
	inj.DB.Preload(clause.Associations).
		Where("city_id = ?", cityId).
		Order("date_utc_millis DESC").
		Last(&lastWeather)

	c.JSON(http.StatusOK, lastWeather)
}
