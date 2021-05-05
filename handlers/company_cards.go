package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// CompanyCardHandler : struct
type CompanyCardHandler struct {
	cardRepo *repositories.CardRepository
}

// NewCompanyCardHandler ...
func NewCompanyCardHandler(cardRepo *repositories.CardRepository) *CompanyCardHandler {
	return &CompanyCardHandler{
		cardRepo: cardRepo,
	}
}

// Index : list all cards
func (h *CompanyCardHandler) Index(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	// get parameters
	var paramVals serializers.CardParametersRequest
	if err := c.ShouldBind(&paramVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(paramVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	data["CompanyID"] = currentCompany.ID

	var cards []models.Card
	if _, err := h.cardRepo.Where(&cards, data, repositories.Paginate(c)); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result []serializers.CardResponse
	if err := serializers.ConvertSerializer(cards, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &cards, Error: nil})
}

// Update ...
func (h *CompanyCardHandler) Update(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	// query card from database
	var card models.Card
	if _, err := h.cardRepo.Where(&card, map[string]interface{}{"code": c.Param("code"), "company_id": currentCompany.ID}); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	var cardVals serializers.CardUpdateRequest
	if err := c.ShouldBindJSON(&cardVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(cardVals, &data)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.cardRepo.Update(&card, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &card, Error: nil})
}
