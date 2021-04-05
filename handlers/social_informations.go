package handlers

import (
	"net/http"
	"strconv"

	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/gin-gonic/gin"
)

// SocialInformationHandler : struct
type SocialInformationHandler struct {
	repoSocialInfo *repositories.SocialInformationRepository
}

// NewSocialInformationHandler ...
func NewSocialInformationHandler(repoSocialInfo *repositories.SocialInformationRepository) *SocialInformationHandler {
	return &SocialInformationHandler{
		repoSocialInfo: repoSocialInfo,
	}
}

// Create ...
func (h *SocialInformationHandler) Create(c *gin.Context) {
	var socialValues serializers.SocialInfoCreateRequest
	if err := c.ShouldBindJSON(&socialValues); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var socialInfo models.SocialInformation
	err := serializers.ConvertSerializer(socialValues, &socialInfo)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.repoSocialInfo.Create(&socialInfo); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &socialInfo, Error: nil})
}

// Destroy ...
func (h *SocialInformationHandler) Destroy(c *gin.Context) {
	var socialValues serializers.SocialInfoRequest
	if err := c.ShouldBindUri(&socialValues); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(socialValues.SocialInformationID)
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