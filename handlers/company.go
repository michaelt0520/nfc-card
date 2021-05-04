package handlers

import (
	"net/http"

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

// Show : show current user
func (h *CompanyHandler) Show(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	var result serializers.CompanyResponse
	if err := serializers.ConvertSerializer(currentCompany, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}

// Update ...
func (h *CompanyHandler) Update(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	// get data update from body
	var compVals serializers.CompanyUpdateRequest
	if err := c.ShouldBindJSON(&compVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// binding to map
	var data map[string]interface{}
	if err := serializers.ConvertSerializer(compVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// update body data to company
	if _, err := h.compRepo.Update(currentCompany, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &company, Error: nil})
}
