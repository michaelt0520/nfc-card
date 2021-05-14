package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// AdminCardHandler : struct
type AdminCardHandler struct {
	cardSrv *services.CardService
}

// NewAdminCardHandler ...
func NewAdminCardHandler(cardSrv *services.CardService) *AdminCardHandler {
	return &AdminCardHandler{
		cardSrv: cardSrv,
	}
}

// Index : list all cards
func (h *AdminCardHandler) Index(c *gin.Context) {
	paramQuery := c.Query("q")
	var filterCard = map[string]interface{}{
		"code": paramQuery,
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

// Show ...
func (h *AdminCardHandler) Show(c *gin.Context) {
	paramQuery := c.Param("code")
	var filterCard = map[string]interface{}{
		"code": paramQuery,
	}

	var preloadData = map[string]interface{}{
		"User": nil,
	}

	var card models.Card
	if err := h.cardSrv.FindOneWithScopes(&card, filterCard, preloadData); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result serializers.CardResponse
	if err := serializers.ConvertSerializer(card, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}

// Create ...
func (h *AdminCardHandler) Create(c *gin.Context) {
	var cardVals serializers.CardCreateRequest
	if err := c.ShouldBindJSON(&cardVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var card models.Card
	err := serializers.ConvertSerializer(cardVals, &card)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.cardSrv.Repo().Create(&card); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &card, Error: nil})
}

// Update ...
func (h *AdminCardHandler) Update(c *gin.Context) {
	paramQuery := c.Param("code")
	var filterCard = map[string]interface{}{
		"code": paramQuery,
	}

	var card models.Card
	if err := h.cardSrv.FindOne(&card, filterCard); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
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

// Destroy ...
func (h *AdminCardHandler) Destroy(c *gin.Context) {
	paramQuery := c.Param("code")
	var filterCard = map[string]interface{}{
		"code": paramQuery,
	}

	if err := h.cardSrv.RemoveCard(filterCard); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
