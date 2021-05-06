package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// InformationHandler : struct
type InformationHandler struct {
	userSrv *services.UserService
	infoSrv *services.InformationService
}

// NewInformationHandler ...
func NewInformationHandler(userSrv *services.UserService, infoSrv *services.InformationService) *InformationHandler {
	return &InformationHandler{
		userSrv: userSrv,
		infoSrv: infoSrv,
	}
}

// Create ...
func (h *InformationHandler) Create(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
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

	if err := h.userSrv.AddInformationToUser(currentUser, &info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var result serializers.InfoResponse
	if err := serializers.ConvertSerializer(info, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}

// Update ...
func (h *InformationHandler) Update(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	infoID := c.Param("id")
	var info models.Information
	if err := h.userSrv.GetInformationFromUser(currentUser, infoID, &info); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
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

	if err := h.infoSrv.Repo().Update(&info, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var result serializers.InfoResponse
	if err := serializers.ConvertSerializer(info, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}

// Destroy ...
func (h *InformationHandler) Destroy(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	infoID := c.Param("id")
	var info models.Information
	if err := h.userSrv.GetInformationFromUser(currentUser, infoID, &info); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	if err := h.infoSrv.Repo().Destroy(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
