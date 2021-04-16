package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// CardHandler : struct
type CardHandler struct {
	cardRepo *repositories.CardRepository
}

// NewCardHandler ...
func NewCardHandler(cardRepo *repositories.CardRepository) *CardHandler {
	return &CardHandler{
		cardRepo: cardRepo,
	}
}

// Index : list all cards
func (h *CardHandler) Index(c *gin.Context) {
	var cards []serializers.CardResponse
	if err := serializers.ConvertSerializer(h.cardRepo.All(), &cards); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: cards, Error: nil})
}

// Show ...
func (h *CardHandler) Show(c *gin.Context) {
	resCard, err := h.cardRepo.Find(c.Param("code"))
	if err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}

	var card serializers.CardResponse
	if err := serializers.ConvertSerializer(resCard, &card); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	card.User.Type = resCard.User.TypeToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: card, Error: nil})
}

// Create ...
func (h *CardHandler) Create(c *gin.Context) {
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

	if err := h.cardRepo.Create(&card); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &card, Error: nil})
}

// Update ...
func (h *CardHandler) Update(c *gin.Context) {
	// query card from database
	card, errGetCard := h.cardRepo.Find(c.Param("code"))
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

// Destroy ...
func (h *CardHandler) Destroy(c *gin.Context) {
	if err := h.cardRepo.Destroy(c.Param("code")); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
