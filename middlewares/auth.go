package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/jwt"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
)

// Auth validate token and authorizes user
func Auth(role models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

		repoUser := repositories.NewUserRepository()
		user, err := repoUser.Find(map[string]interface{}{"jwt": token})
		if err != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		_, errExtract := jwt.ExtractToken(token)
		if errExtract != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		if role != 0 && user.Role != role {
			c.JSON(http.StatusUnauthorized, "not have permission")
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

// MemberAuth authorize user role
func MemberAuth() gin.HandlerFunc {
	return Auth(models.UserMember)
}

// AdminAuth authorize user role
func AdminAuth() gin.HandlerFunc {
	return Auth(models.UserAdmin)
}

// AllAuth authorize except role
func AllAuth() gin.HandlerFunc {
	return Auth(0)
}
