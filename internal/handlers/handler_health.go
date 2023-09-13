package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerHealth(c *gin.Context) {
	c.JSON(http.StatusOK, struct{}{})
}
