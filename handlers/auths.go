package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/jwt"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// AuthHandler : Session Handler
type AuthHandler struct {
	userSrv *services.UserService
}

const DefaultAvatar = "/public/avatars/default_avatar.png"

// NewAuthHandler : Constructor
func NewAuthHandler(userSrv *services.UserService) *AuthHandler {
	return &AuthHandler{
		userSrv: userSrv,
	}
}

// Signup : POST #signup
func (h *AuthHandler) Signup(c *gin.Context) {
	var signupVals serializers.SignUpRequest
	if err := c.ShouldBindJSON(&signupVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	err := serializers.ConvertSerializer(signupVals, &user)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	user.Avatar = fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar)
	user.Type = models.Personal
	user.Role = models.UserStandard

	if err := h.userSrv.Repo().Create(&user); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "signup success", Error: nil})
}

// Signin : POST #signin
func (h *AuthHandler) Signin(c *gin.Context) {
	var loginVals serializers.SigninRequest
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var filterUser = map[string]interface{}{
		"username": loginVals.Username,
	}

	var resU models.User
	if err := h.userSrv.FindOne(&resU, filterUser); err != nil {
		fmt.Println(err.Error())
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	if err := resU.CheckPassword(loginVals.Password); err != nil {
		respondError(c, http.StatusForbidden, err.Error())
		return
	}

	token, err := jwt.CreateToken(&resU)
	if err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if errUpdate := h.userSrv.Repo().Update(&resU, map[string]interface{}{"jwt": token}); errUpdate != nil {
		respondError(c, http.StatusUnprocessableEntity, errUpdate.Error())
		return
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(resU, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	resUser.Role = resU.RoleToString()
	resUser.Type = resU.TypeToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: serializers.SigninResponse{User: resUser, Token: token}, Error: nil})
}

// Signout : DELETE #signout
func (h *AuthHandler) Signout(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var dataUser = map[string]interface{}{
		"jwt": nil,
	}

	if errUpdate := h.userSrv.Repo().Update(currentUser, dataUser); errUpdate != nil {
		respondError(c, http.StatusUnprocessableEntity, errUpdate.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "logout successfully", Error: nil})
}
