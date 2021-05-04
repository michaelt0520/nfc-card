package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/jwt"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
)

// BasicAuth check basic auth
func BasicAuth(c *gin.Context) *models.User {
	// get token from header
	token := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

	// extract token
	payload, errExtract := jwt.ExtractToken(token)
	if errExtract != nil {
		return nil
	}

	// get user from token
	userRepo := repositories.NewUserRepository()
	var currentUser models.User
	_, err := userRepo.Find(&currentUser, map[string]interface{}{"id": payload.UserID})

	// valid current user with token
	if currentUser.JWT != token || err != nil {
		return nil
	}

	return &currentUser
}

// StandardAuth authorize user role
func StandardAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := BasicAuth(c)
		if currentUser == nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		if currentUser.Role != models.UserStandard && currentUser.Role != models.UserCompanyMember {
			c.JSON(http.StatusUnauthorized, errors.DontHavePermission.Error())
			c.Abort()
			return
		}

		c.Set("currentUser", currentUser)
		c.Next()
	}
}

// CompanyAuth authorize company
func CompanyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := BasicAuth(c)
		if currentUser == nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		if !(currentUser.Role == models.UserCompanyManager && currentUser.Type == models.Business) {
			c.JSON(http.StatusUnauthorized, errors.DontHavePermission.Error())
			c.Abort()
			return
		}

		c.Set("currentUser", currentUser)
		c.Set("currentCompany", &currentUser.Company)
		c.Next()
	}
}

// AdminAuth authorize user role
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := BasicAuth(c)
		if currentUser == nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		if currentUser.Role != models.UserAdmin {
			c.JSON(http.StatusUnauthorized, errors.DontHavePermission.Error())
			c.Abort()
			return
		}

		c.Set("currentUser", currentUser)
		c.Next()
	}
}

// AllAuth authorize except role
func AllAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := BasicAuth(c)
		if currentUser == nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		c.Set("currentUser", currentUser)
		c.Next()
	}
}
