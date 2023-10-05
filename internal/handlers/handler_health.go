package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerHealth returns ok. Useful to do a fast check on the API status
func HandlerHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
