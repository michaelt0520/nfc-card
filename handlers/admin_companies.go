package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// AdminCompanyHandler : struct
type AdminCompanyHandler struct {
	compSrv *services.CompanyService
}

// NewAdminCompanyHandler ...
func NewAdminCompanyHandler(compSrv *services.CompanyService) *AdminCompanyHandler {
	return &AdminCompanyHandler{
		compSrv: compSrv,
	}
}

// Index : list all companies
func (h *AdminCompanyHandler) Index(c *gin.Context) {
	paramQuery := c.Query("q")
	var filterCompany = map[string]interface{}{
		"name": paramQuery,
	}

	var companies []models.Company
	if err := h.compSrv.FindMany(&companies, filterCompany, c); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result []serializers.CompanyResponse
	if err := serializers.ConvertSerializer(companies, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}

// Show ...
func (h *AdminCompanyHandler) Show(c *gin.Context) {
	companyID := c.Param("id")
	var filterCompany = map[string]interface{}{
		"id": companyID,
	}

	var company models.Company
	if err := h.compSrv.FindOne(&company, filterCompany); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	var result serializers.CompanyResponse
	if err := serializers.ConvertSerializer(company, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}

// Create ...
func (h *AdminCompanyHandler) Create(c *gin.Context) {
	// get create data from body
	var compVals serializers.CompanyCreateRequest
	if err := c.ShouldBindJSON(&compVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// binding data to model
	var comp models.Company
	err := serializers.ConvertSerializer(compVals, &comp)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// save to db
	if err := h.compSrv.Repo().Create(&comp); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &comp, Error: nil})
}

// Update ...
func (h *AdminCompanyHandler) Update(c *gin.Context) {
	companyID := c.Param("id")
	var filterCompany = map[string]interface{}{
		"id": companyID,
	}

	var company models.Company
	if err := h.compSrv.FindOne(&company, filterCompany); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	// get data update from body
	var compVals serializers.CompanyUpdateRequest
	if err := c.ShouldBindJSON(&compVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// binding to map
	var data map[string]interface{}
	err := serializers.ConvertSerializer(compVals, &data)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// update body data to company
	if err := h.compSrv.Repo().Update(&company, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &company, Error: nil})
}

// Destroy ...
func (h *AdminCompanyHandler) Destroy(c *gin.Context) {
	companyID := c.Param("id")
	var filterCompany = map[string]interface{}{
		"id": companyID,
	}

	var company models.Company
	if err := h.compSrv.FindOne(&company, filterCompany); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	// update delete company to db
	if err := h.compSrv.Repo().Destroy(&company); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
