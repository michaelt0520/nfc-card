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

// AdminInformationHandler : struct
type AdminInformationHandler struct {
	infoRepo *repositories.InformationRepository
	userRepo *repositories.UserRepository
}

// NewAdminInformationHandler ...
func NewAdminInformationHandler(infoRepo *repositories.InformationRepository, userRepo *repositories.UserRepository) *AdminInformationHandler {
	return &AdminInformationHandler{
		infoRepo: infoRepo,
		userRepo: userRepo,
	}
}

// Index : list all info of user
func (h *AdminInformationHandler) Index(c *gin.Context) {
	var user models.User
	if _, err := h.userRepo.Find(&user, map[string]interface{}{"username": c.Param("username")}); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
	}

	var infos []serializers.InfoResponse
	if err := serializers.ConvertSerializer(user.Informations, &infos); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: infos, Error: nil})
}

// Create ...
func (h *AdminInformationHandler) Create(c *gin.Context) {
	var user models.User
	if _, err := h.userRepo.Find(&user, map[string]interface{}{"username": c.Param("username")}); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
	}

	var infoValues serializers.InfoCreateRequest
	if err := c.ShouldBindJSON(&infoValues); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var info models.Information
	if err := serializers.ConvertSerializer(infoValues, &info); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	info.UserID = user.ID
	info.Visibled = true

	if _, err := h.infoRepo.Create(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: info, Error: nil})
}

// Update ...
func (h *AdminInformationHandler) Update(c *gin.Context) {
	var user models.User
	if _, err := h.userRepo.Find(&user, map[string]interface{}{"username": c.Param("username")}); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
	}

	// get info's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query info from database
  var info models.Information
	if _, errGetInfo := h.infoRepo.Find(&info, map[string]interface{}{"id": id, "user_id": user.ID}); errGetInfo != nil {
		respondError(c, http.StatusNotFound, errGetInfo.Error())
		return
	}

	var infoVals serializers.InfoUpdateRequest
	if err := c.ShouldBindJSON(&infoVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(infoVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.infoRepo.Update(&info, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &info, Error: nil})
}

// Destroy ...
func (h *AdminInformationHandler) Destroy(c *gin.Context) {
	var user models.User
	if _, err := h.userRepo.Find(&user, map[string]interface{}{"username": c.Param("username")}); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
	}

	// get info's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query info from database
  var info models.Information
	if _, err := h.infoRepo.Find(&info, map[string]interface{}{"id": id, "user_id": user.ID}); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	if _, err := h.infoRepo.Destroy(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "info deleted", Error: nil})
}
