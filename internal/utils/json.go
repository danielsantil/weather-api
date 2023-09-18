package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func RespondWithError(c *gin.Context, code int, message string) {
	if code >= 500 {
		log.Println("Responding with 5xx error", message)
	}

	c.JSON(code, struct {
		Error string `json:"error"`
	}{Error: message})
}
