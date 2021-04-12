package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// InformationHandler : struct
type InformationHandler struct {
	repoSocialInfo *repositories.InformationRepository
}

// NewInformationHandler ...
func NewInformationHandler(repoSocialInfo *repositories.InformationRepository) *InformationHandler {
	return &InformationHandler{
		repoSocialInfo: repoSocialInfo,
	}
}

// Create ...
func (h *InformationHandler) Create(c *gin.Context) {
	userID, ok := c.Get("UserID")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}

	var socialValues serializers.InfoCreateRequest
	if err := c.ShouldBindJSON(&socialValues); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var info models.Information
	err := serializers.ConvertSerializer(socialValues, &info)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	uid, err := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 32)
	if err != nil {
		respondError(c, http.StatusUnauthorized, errors.ParameterInvalid.Error())
		return
	}

	info.UserID = uint(uid)

	if err := h.repoSocialInfo.Create(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &info, Error: nil})
}

// Destroy ...
func (h *InformationHandler) Destroy(c *gin.Context) {
	var socialValues serializers.InfoRequest
	if err := c.ShouldBindUri(&socialValues); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(socialValues.InformationID)
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	if err := h.repoSocialInfo.Destroy(uint(id)); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
