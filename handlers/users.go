package handlers

import (
	"fmt"
	"net/http"
	"os"

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

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}

// Upload ...
func (h *UserHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
	}

	filename := fmt.Sprintf("./public/avatars/%s", file.Filename)

	if err = c.SaveUploadedFile(file, filename); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := fmt.Sprintf("%s/public/avatars/%s", os.Getenv("app_host"), file.Filename)

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}
