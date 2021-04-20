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
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(models.User)

	var infoValues serializers.InfoCreateRequest
	if err := c.ShouldBindJSON(&infoValues); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var info models.Information
	err := serializers.ConvertSerializer(infoValues, &info)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	info.UserID = currentUser.ID
	info.Visibled = true

	if err := h.repoInfo.Create(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &info, Error: nil})
}

// Update ...
func (h *InformationHandler) Update(c *gin.Context) {
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(models.User)

	// get info's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query info from database
	info, errGetInfo := h.repoInfo.Find(map[string]interface{}{"id": id, "user_id": currentUser.ID})
	if errGetInfo != nil {
		respondError(c, http.StatusNotFound, errGetInfo.Error())
		return
	}

	var infoVals serializers.InfoUpdateRequest
	if err := c.ShouldBindJSON(&infoVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(infoVals, &data)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repoInfo.Update(info, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: info, Error: nil})
}

// Destroy ...
func (h *InformationHandler) Destroy(c *gin.Context) {
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(models.User)

	// get info's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query info from database
	info, errGetInfo := h.repoInfo.Find(map[string]interface{}{"id": id, "user_id": currentUser.ID})
	if errGetInfo != nil {
		respondError(c, http.StatusNotFound, errGetInfo.Error())
		return
	}

	if err := h.repoInfo.Destroy(info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
