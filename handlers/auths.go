package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/jwt"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// AuthHandler : Session Handler
type AuthHandler struct {
	repoUser *repositories.UserRepository
}

const (
	ExpiredCookie = 3600 * 24 * 30
	DefaultAvatar = "/public/images/default_avatar.png"
)

// NewAuthHandler : Constructor
func NewAuthHandler(userRepo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		repoUser: userRepo,
	}
}

// Signup : POST #signup
func (h *AuthHandler) Signup(c *gin.Context) {
	var signupVals serializers.SignUpRequest
	if err := c.ShouldBindJSON(&signupVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User
	err := serializers.ConvertSerializer(signupVals, &user)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Avatar = fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar)
	user.Type = models.Personal

	if err := h.repoUser.Create(&user); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: user, Error: nil})
}

// Signin : POST #signin
func (h *AuthHandler) Signin(c *gin.Context) {
	var loginVals serializers.SigninRequest
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(loginVals, &data); err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.repoUser.Find(data)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	} else if res == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	if err := res.CheckPassword(loginVals.Password); err != nil {
		respondError(c, http.StatusUnprocessableEntity, errors.PasswordIncorrect.Error())
		return
	}

	token, err := jwt.CreateToken(res)
	if err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response := serializers.SigninResponse{
		User:  res,
		Token: token,
	}

	c.SetCookie("Authorized", token, ExpiredCookie, "/", os.Getenv("app_host"), false, true)
	c.JSON(http.StatusOK, serializers.Resp{Result: response, Error: nil})
}

// Signout : DELETE #signout
func (h *AuthHandler) Signout(c *gin.Context) {
	c.SetCookie("Authorized", "", -1, "", os.Getenv("app_host"), false, true)
	c.JSON(http.StatusOK, serializers.Resp{Result: "logout successfully", Error: nil})
}
