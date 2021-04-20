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

	var cards []serializers.CardResponse
	if err := serializers.ConvertSerializer(currentCompany.Cards, &cards); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: cards, Error: nil})
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

	// get card code from params
	cardCode := c.Param("code")

	// query card from database
	card, errGetCard := h.cardRepo.Find(map[string]interface{}{"code": cardCode, "company_id": currentCompany.ID})
	if errGetCard != nil {
		respondError(c, http.StatusNotFound, errGetCard.Error())
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

	if err := h.cardRepo.Update(card, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &card, Error: nil})
}
