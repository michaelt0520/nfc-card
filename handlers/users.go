package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
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
	// query user from database
	var userParams = make(map[string]interface{})
	userParams["id"] = c.Param("username")

	user, errGetUser := h.userRepo.Find(userParams)
	if errGetUser != nil {
		respondError(c, http.StatusNotFound, errGetUser.Error())
		return
	}
	if user == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	var userVals serializers.UserUpdateSerializer
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(userVals, &data)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.userRepo.Update(user, data)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if res == nil {
		respondError(c, http.StatusNoContent, errors.RecordNotFound.Error())
		return
	}

	var resUser serializers.UserSerializer
	if err := serializers.ConvertSerializer(res, &resUser); err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}
