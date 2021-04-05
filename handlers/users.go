package handlers

import (
	"net/http"
	"strconv"

	"github.com/michaelt0520/nfc-card/daos"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/gin-gonic/gin"
)

// UserHandler : struct
type UserHandler struct {
	userDao *daos.UserDAO
}

// NewUserHandler ...
func NewUserHandler(userDAO *daos.UserDAO) *UserHandler {
	return &UserHandler{
		userDao: userDAO,
	}
}

// Find ...
func (h *UserHandler) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("username"))
	if err == nil || id == "" {
		respondError(c, http.StatusBadRequest, "id is invalid")
		return
	}

	res, err := h.userDao.Find(id)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if res == nil {
		respondError(c, http.StatusNoContent, "user not found")
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: res, Error: nil})
}
