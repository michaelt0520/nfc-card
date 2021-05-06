package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// CompanyCardHandler : struct
type CompanyCardHandler struct {
	compSrv *services.CompanyService
	cardSrv *services.CardService
}

// NewCompanyCardHandler ...
func NewCompanyCardHandler(compSrv *services.CompanyService, cardSrv *services.CardService) *CompanyCardHandler {
	return &CompanyCardHandler{
		compSrv: compSrv,
		cardSrv: cardSrv,
	}
}

// Index : list all cards
func (h *CompanyCardHandler) Index(c *gin.Context) {
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	paramQuery := c.Query("q")
	var filterCard = map[string]interface{}{
		"company_id": currentCompany.ID,
		"code":       paramQuery,
	}

	var preloadData = map[string]interface{}{
		"User": nil,
	}

	var cards []models.Card
	if err := h.cardSrv.FindManyWithScopes(&cards, filterCard, preloadData, c); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result []serializers.CardResponse
	if err := serializers.ConvertSerializer(cards, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}

// Update ...
func (h *CompanyCardHandler) Update(c *gin.Context) {
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var filterCard = map[string]interface{}{
		"code":       c.Param("code"),
		"company_id": currentCompany.ID,
	}

	// query card from database
	var card models.Card
	if err := h.cardSrv.FindOne(&card, filterCard); err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	var cardVals serializers.CardUpdateRequest
	if err := c.ShouldBindJSON(&cardVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(cardVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.cardSrv.Repo().Update(&card, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var result serializers.CardResponse
	if err := serializers.ConvertSerializer(card, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}
