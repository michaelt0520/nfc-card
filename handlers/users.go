package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// UserHandler : struct
type UserHandler struct {
	userRepo *repositories.UserRepository
}

// NewUserHandler ...
func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// Show : show current user
func (h *UserHandler) Show(c *gin.Context) {
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(*models.User)

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
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(*models.User)

	// get data from body
	var userVals serializers.UserUpdateRequest
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(userVals, &data)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userRepo.Update(currentUser, data); err != nil {
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
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(*models.User)

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

	if errUpdate := h.userRepo.Update(currentUser, map[string]interface{}{"jwt": nil, "password": *userVals.NewPassword}); errUpdate != nil {
		respondError(c, http.StatusUnprocessableEntity, errUpdate.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "Password Updated", Error: nil})
}
