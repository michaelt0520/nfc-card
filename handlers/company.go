package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
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
	if err := h.compRepo.Update(company, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &company, Error: nil})
}
