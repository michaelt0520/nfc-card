package handlers

import (
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/gin-gonic/gin"
)

// respondError ...
func respondError(c *gin.Context, status int, msg string) {
	c.JSON(status, serializers.Resp{Error: map[string]interface{}{
		"status":  status,
		"message": msg,
	}})
}
