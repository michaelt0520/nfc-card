package handlers

import (
	"net/http"

	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/gin-gonic/gin"
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

// Find ...
func (h *UserHandler) Find(c *gin.Context) {
	userName := c.Param("username")
	if userName == "" {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	res, err := h.userRepo.Find(userName)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if res == nil {
		respondError(c, http.StatusNoContent, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: res, Error: nil})
}
