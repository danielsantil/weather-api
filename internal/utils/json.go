package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

// RespondWithError formats message into a JSON response indicating an error response.
func RespondWithError(c *gin.Context, code int, message string) {
	if code >= 500 {
		log.Println("Responding with 5xx error", message)
	}

	c.JSON(code, struct {
		Error string `json:"error"`
	}{Error: message})
}
