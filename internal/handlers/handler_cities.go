package handlers

import (
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerGetCities returns list of cities ordered by name
//
// If an error occurs while retrieving from the database, a 4xx status code is returned
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
