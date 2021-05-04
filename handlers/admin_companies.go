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

// AdminCompanyHandler : struct
type AdminCompanyHandler struct {
	compRepo *repositories.CompanyRepository
}

// NewAdminCompanyHandler ...
func NewAdminCompanyHandler(compRepo *repositories.CompanyRepository) *AdminCompanyHandler {
	return &AdminCompanyHandler{
		compRepo: compRepo,
	}
}

// Index : list all companies
func (h *AdminCompanyHandler) Index(c *gin.Context) {
	var companies []serializers.CompanyResponse

	if _, err := h.compRepo.All(&companies); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &companies, Error: nil})
}

// Show ...
func (h *AdminCompanyHandler) Show(c *gin.Context) {
	// get company's id from params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query company by id
	var company serializers.CompanyResponse
	if _, err := h.compRepo.Find(&company, map[string]interface{}{"id": id}); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &company, Error: nil})
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
	if _, err := h.compRepo.Create(&comp); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &comp, Error: nil})
}

// Update ...
func (h *AdminCompanyHandler) Update(c *gin.Context) {
	// get company's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query company from database
	var company models.Company
	if _, err := h.compRepo.Find(&company, map[string]interface{}{"id": id}); err != nil {
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
	if _, err := h.compRepo.Update(&company, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &company, Error: nil})
}

// Destroy ...
func (h *AdminCompanyHandler) Destroy(c *gin.Context) {
	// get company's id from params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query company from database
	var company models.Company
	if _, err := h.compRepo.Find(&company, map[string]interface{}{"id": id}); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	// update delete company to db
	if _, err := h.compRepo.Destroy(&company); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
