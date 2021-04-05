package handlers

import (
	"net/http"
	"strconv"

	"github.com/michaelt0520/nfc-card/daos"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/gin-gonic/gin"
)

// SocialInformationHandler : struct
type SocialInformationHandler struct {
	userDao *daos.UserDAO
}

// NewSocialInformationHandler ...
func NewSocialInformationHandler(userDAO *daos.UserDAO) *SocialInformationHandler {
	return &SocialInformationHandler{
		userDao: userDAO,
	}
}

// Create ...
func (u *SocialInformationHandler) Create(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, "id is invalid")
		return
	}

	res, err := u.userDao.Find(id)
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

// Destroy ...
func (u *SocialInformationHandler) Destroy(c *gin.Context) {
}