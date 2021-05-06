package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// UserHandler : struct
type UserHandler struct {
	userSrv *services.UserService
}

// NewUserHandler ...
func NewUserHandler(userSrv *services.UserService) *UserHandler {
	return &UserHandler{
		userSrv: userSrv,
	}
}

// Show : show current user
func (h *UserHandler) Show(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(currentUser, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	resUser.Type = currentUser.TypeToString()
	resUser.Role = currentUser.RoleToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}

// Update ...
func (h *UserHandler) Update(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	// get data from body
	var userVals serializers.UserUpdateRequest
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(userVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userSrv.Repo().Update(currentUser, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(currentUser, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	resUser.Type = currentUser.TypeToString()
	resUser.Role = currentUser.RoleToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}

// UpdatePassword : update current user's password
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	// get data from body
	var userVals serializers.UserUpdatePasswordRequest
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := currentUser.CheckPassword(*userVals.OldPassowrd); err != nil {
		respondError(c, http.StatusForbidden, errors.PasswordIncorrect.Error())
		return
	}

	if *userVals.NewPassword != *userVals.PasswordConfirmation {
		respondError(c, http.StatusUnprocessableEntity, errors.ConfirmationPasswordIncorrect.Error())
		return
	}

	if errUpdate := h.userSrv.Repo().Update(currentUser, map[string]interface{}{"jwt": nil, "password": *userVals.NewPassword}); errUpdate != nil {
		respondError(c, http.StatusUnprocessableEntity, errUpdate.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "Password Updated", Error: nil})
}
