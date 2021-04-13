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

// CompanyHandler : struct
type CompanyHandler struct {
	compRepo *repositories.CompanyRepository
}

// NewCompanyHandler ...
func NewCompanyHandler(compRepo *repositories.CompanyRepository) *CompanyHandler {
	return &CompanyHandler{
		compRepo: compRepo,
	}
}

// Index : list all companies
func (h *CompanyHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, serializers.Resp{Result: h.compRepo.All(), Error: nil})
}

// Show ...
func (h *CompanyHandler) Show(c *gin.Context) {
	// get company's id from params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query company by id
	resComp, err := h.compRepo.Find(uint(id))
	if err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}
	if resComp == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: resComp, Error: nil})
}

// Create ...
func (h *CompanyHandler) Create(c *gin.Context) {
	// get create data from body
	var compVals serializers.CompanyCreateRequest
	if err := c.ShouldBindJSON(&compVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// binding data to model
	var comp models.Company
	err := serializers.ConvertSerializer(compVals, &comp)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// save to db
	if err := h.compRepo.Create(&comp); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &comp, Error: nil})
}

// Update ...
func (h *CompanyHandler) Update(c *gin.Context) {
	// get company's id from params
	id, errGetID := strconv.Atoi(c.Param("id"))
	if errGetID != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// query company from database
	company, errGetComp := h.compRepo.Find(uint(id))
	if errGetComp != nil {
		respondError(c, http.StatusNotFound, errGetComp.Error())
		return
	}
	if company == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	// get data update from body
	var compVals serializers.CompanyUpdateRequest
	if err := c.ShouldBindJSON(&compVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// binding to map
	var data map[string]interface{}
	err := serializers.ConvertSerializer(compVals, &data)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// update body data to company
	resComp, err := h.compRepo.Update(company, data)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if resComp == nil {
		respondError(c, http.StatusNoContent, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &resComp, Error: nil})
}

// Destroy ...
func (h *CompanyHandler) Destroy(c *gin.Context) {
	// get company's id from params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	// update delete company to db
	if err := h.compRepo.Destroy(uint(id)); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
