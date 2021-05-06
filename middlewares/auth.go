package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/jwt"
	"github.com/michaelt0520/nfc-card/models"
)

// BasicAuth check basic auth
func BasicAuth(c *gin.Context) (*jwt.PayLoad, error) {
	// get token from header
	token := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

	// extract token
	payload, errExtract := jwt.ExtractToken(token)
	if errExtract != nil {
		return nil, errExtract
	}

	c.Set("user_id", payload.UserID)

	return payload, nil
}

// StandardAuth authorize user role
func StandardAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := BasicAuth(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		if payload.UserRole != models.UserStandard && payload.UserRole != models.UserCompanyMember {
			c.JSON(http.StatusUnauthorized, errors.DontHavePermission.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

// CompanyAuth authorize company
func CompanyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := BasicAuth(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		if !(payload.UserRole == models.UserCompanyManager && payload.UserType == models.Business) {
			c.JSON(http.StatusUnauthorized, errors.DontHavePermission.Error())
			c.Abort()
			return
		}

		c.Set("company_id", payload.CompanyID)
		c.Next()
	}
}

// AdminAuth authorize user role
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := BasicAuth(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.InvalidToken.Error())
			c.Abort()
			return
		}

		if payload.UserRole != models.UserAdmin {
			c.JSON(http.StatusUnauthorized, errors.DontHavePermission.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

// AllAuth authorize except role
func AllAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := BasicAuth(c); err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
