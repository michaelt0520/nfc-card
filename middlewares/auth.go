package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/jwt"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
)

// Auth validate token and authorizes currentUser
func Auth(role models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header
		token := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

		// extract token
		payload, errExtract := jwt.ExtractToken(token)
		if errExtract != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		// get user from token
		repoUser := repositories.NewUserRepository()
		currentUser, err := repoUser.Find(map[string]interface{}{"id": payload.UserID})

		// valid current user with token
		if currentUser.JWT != token || err != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		// check role user
		if role != 0 && currentUser.Role != role {
			c.JSON(http.StatusUnauthorized, "not have permission")
			c.Abort()
			return
		}

		c.Set("currentUser", currentUser)
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
