package handlers

import (
	"errors"
	"github.com/danielsantil/weather-api/internal/models"
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/danielsantil/weather-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type GetWeatherRequest struct {
	Id string `uri:"id" binding:"uuid"`
}

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
	err := inj.DB.Preload(clause.Associations).
		Where("city_id = ?", cityId).
		Order("date_utc_millis DESC").
		Last(&lastWeather).Error

	if err != nil {
		returnError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.WeatherSummary{
		Id:          lastWeather.ID,
		CityId:      lastWeather.CityId,
		Name:        lastWeather.Name,
		Country:     lastWeather.Sys.Country,
		IconId:      lastWeather.Conditions[0].Icon,
		Description: lastWeather.Conditions[0].Description,
		Temp:        lastWeather.Main.Temp,
		FeelsLike:   lastWeather.Main.FeelsLike,
	})
}

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

func returnError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
	} else {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return
}
