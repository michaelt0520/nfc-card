package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/serializers"
)

// respondError ...
func respondError(c *gin.Context, status int, msg string) {
	c.JSON(status, serializers.Resp{Error: map[string]interface{}{
		"status":  status,
		"message": msg,
	}})
}
