package handlers

import (
	"errors"
	"github.com/danielsantil/weather-api/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func returnError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
	} else {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return
}
