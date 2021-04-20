package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// AdminCardHandler : struct
type AdminCardHandler struct {
	cardRepo *repositories.CardRepository
}

// NewAdminCardHandler ...
func NewAdminCardHandler(cardRepo *repositories.CardRepository) *AdminCardHandler {
	return &AdminCardHandler{
		cardRepo: cardRepo,
	}
}

// Index : list all cards
func (h *AdminCardHandler) Index(c *gin.Context) {
	var cards []serializers.CardResponse
	if err := serializers.ConvertSerializer(h.cardRepo.All(), &cards); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: cards, Error: nil})
}

// Show ...
func (h *AdminCardHandler) Show(c *gin.Context) {
	resCard, err := h.cardRepo.Find(map[string]interface{}{"code": c.Param("code")})
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
	card.User.Role = resCard.User.RoleToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: card, Error: nil})
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

	card.Activated = true

	if err := h.cardRepo.Create(&card); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &card, Error: nil})
}

// Update ...
func (h *AdminCardHandler) Update(c *gin.Context) {
	// query card from database
	card, errGetCard := h.cardRepo.Find(map[string]interface{}{"code": c.Param("code")})
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
func (h *AdminCardHandler) Destroy(c *gin.Context) {
	if err := h.cardRepo.Destroy(c.Param("code")); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "deleted", Error: nil})
}
