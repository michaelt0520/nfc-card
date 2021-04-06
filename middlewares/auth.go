package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/jwt"
)

// Auth validate token and authorizes user
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorized")
		if err != nil {
			c.JSON(http.StatusUnauthorized, "token not found")
			c.Abort()
			return
		}

		payload, err := jwt.ExtractToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Set("UserID", payload.UserID)
	}
}
