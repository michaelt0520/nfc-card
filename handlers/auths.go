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
	userRepo *repositories.UserRepository
}

const (
	ExpiredCookie = 3600 * 24 * 30
	DefaultAvatar = "/public/images/default_avatar.png"
)

// NewAuthHandler : Constructor
func NewAuthHandler(userRepo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
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
	user.Role = models.UserMember

	if err := h.userRepo.Create(&user); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var resUser serializers.UserSerializer
	if err := serializers.ConvertSerializer(user, &resUser); err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}

// Signin : POST #signin
func (h *AuthHandler) Signin(c *gin.Context) {
	var loginVals serializers.SigninRequest
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var loginData map[string]interface{}
	if err := serializers.ConvertSerializer(loginVals, &loginData); err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	resU, err := h.userRepo.Find(loginData)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	} else if resU == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	if err := resU.CheckPassword(loginVals.Password); err != nil {
		respondError(c, http.StatusUnprocessableEntity, errors.PasswordIncorrect.Error())
		return
	}

	token, err := jwt.CreateToken(resU)
	if err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var userParams = make(map[string]interface{})
	userParams["jwt"] = token

	resU, errUpdate := h.userRepo.Update(resU, userParams)
	if errUpdate != nil {
		respondError(c, http.StatusInternalServerError, errUpdate.Error())
		return
	}

	var resUser serializers.UserSerializer
	if err := serializers.ConvertSerializer(resU, &resUser); err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := serializers.SigninResponse{
		User:  resUser,
		Token: token,
	}

	c.SetCookie("Authorized", token, ExpiredCookie, "/", os.Getenv("app_host"), false, true)
	c.JSON(http.StatusOK, serializers.Resp{Result: response, Error: nil})
}

// Signout : DELETE #signout
func (h *AuthHandler) Signout(c *gin.Context) {
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(*models.User)

	var userJWTParams = make(map[string]interface{})
	userJWTParams["jwt"] = nil

	if _, errUpdate := h.userRepo.Update(currentUser, userJWTParams); errUpdate != nil {
		respondError(c, http.StatusInternalServerError, errUpdate.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "logout successfully", Error: nil})
}
