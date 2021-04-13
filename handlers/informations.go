package handlers

import (
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
	repoInfo *repositories.InformationRepository
}

// NewInformationHandler ...
func NewInformationHandler(repoInfo *repositories.InformationRepository) *InformationHandler {
	return &InformationHandler{
		repoInfo: repoInfo,
	}
}

// Create ...
func (h *InformationHandler) Create(c *gin.Context) {
	userID, ok := c.Get("UserID")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}

	var infoValues serializers.InfoCreateRequest
	if err := c.ShouldBindJSON(&infoValues); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var info models.Information
	err := serializers.ConvertSerializer(infoValues, &info)
	if err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	info.UserID = userID.(uint)

	if err := h.repoInfo.Create(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &info, Error: nil})
}

// Update ...
func (h *InformationHandler) Update(c *gin.Context) {
	// get info's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query info from database
	info, errGetInfo := h.repoInfo.Find(uint(id))
	if errGetInfo != nil {
		respondError(c, http.StatusNotFound, errGetInfo.Error())
		return
	}
	if info == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	var infoVals serializers.InfoUpdateRequest
	if err := c.ShouldBindJSON(&infoVals); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(infoVals, &data)
	if err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	resInfo, err := h.repoInfo.Update(info, data)
	if err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if resInfo == nil {
		respondError(c, http.StatusNoContent, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: resInfo, Error: nil})
}

// Destroy ...
func (h *InformationHandler) Destroy(c *gin.Context) {
	// get info's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	if err := h.repoInfo.Destroy(uint(id)); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
