package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// CompanyHandler : struct
type CompanyHandler struct {
	compSrv *services.CompanyService
}

// NewCompanyHandler ...
func NewCompanyHandler(compSrv *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		compSrv: compSrv,
	}
}

// Show : show current user
func (h *CompanyHandler) Show(c *gin.Context) {
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var result serializers.CompanyResponse
	if err := serializers.ConvertSerializer(currentCompany, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}

// Update ...
func (h *CompanyHandler) Update(c *gin.Context) {
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
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
	if err := serializers.ConvertSerializer(compVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// update body data to company
	if err := h.compSrv.Repo().Update(currentCompany, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &currentCompany, Error: nil})
}
