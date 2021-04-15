package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

// Index : list all users
func (h *UserHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, serializers.Resp{Result: h.userRepo.All(), Error: nil})
}

// Update ...
func (h *UserHandler) Update(c *gin.Context) {
	// get data from body
	var userVals serializers.UserUpdateRequest
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// query user from database
	user, errGetUser := h.userRepo.Find(map[string]interface{}{"id": c.Param("username")})
	if errGetUser != nil {
		respondError(c, http.StatusNotFound, errGetUser.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(userVals, &data)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userRepo.Update(user, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(user, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	resUser.Type = user.TypeToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}
