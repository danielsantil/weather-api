package handlers

import (
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (inj *Injector) HandlerGetCities(c *gin.Context) {
	var cities []database.City
	err := inj.DB.
		Order("name").
		Find(&cities).
		Error

	if err != nil {
		returnError(c, err)
		return
	}

	c.JSON(http.StatusOK, cities)
}
