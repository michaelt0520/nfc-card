package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// AdminInformationHandler : struct
type AdminInformationHandler struct {
	userSrv *services.UserService
	infoSrv *services.InformationService
}

// NewAdminInformationHandler ...
func NewAdminInformationHandler(userSrv *services.UserService, infoSrv *services.InformationService) *AdminInformationHandler {
	return &AdminInformationHandler{
		userSrv: userSrv,
		infoSrv: infoSrv,
	}
}

// Index : list all info of user
func (h *AdminInformationHandler) Index(c *gin.Context) {
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var preloadData = map[string]interface{}{
		"Informations": nil,
	}

	var user models.User
	if err := h.userSrv.FindOneWithScopes(&user, filterUser, preloadData); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
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
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var user models.User
	if err := h.userSrv.FindOne(&user, filterUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
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

	if err := h.userSrv.AddInformationToUser(&user, &info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &info, Error: nil})
}

// Update ...
func (h *AdminInformationHandler) Update(c *gin.Context) {
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var user models.User
	if err := h.userSrv.FindOne(&user, filterUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	infoID := c.Param("id")
	var info models.Information
	if err := h.userSrv.GetInformationFromUser(&user, infoID, &info); err != nil {
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

	c.JSON(http.StatusOK, serializers.Resp{Result: &info, Error: nil})
}

// Destroy ...
func (h *AdminInformationHandler) Destroy(c *gin.Context) {
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var user models.User
	if err := h.userSrv.FindOne(&user, filterUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	infoID := c.Param("id")
	var info models.Information
	if err := h.userSrv.GetInformationFromUser(&user, infoID, &info); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	if err := h.infoSrv.Repo().Destroy(&info); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
