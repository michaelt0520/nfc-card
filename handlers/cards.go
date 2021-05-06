package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// CardHandler : struct
type CardHandler struct {
	userSrv *services.UserService
	cardSrv *services.CardService
}

// NewCardHandler ...
func NewCardHandler(userSrv *services.UserService, cardSrv *services.CardService) *CardHandler {
	return &CardHandler{
		userSrv: userSrv,
		cardSrv: cardSrv,
	}
}

// Create ...
func (h *CardHandler) Create(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var cardVals struct {
		Code *string `json:"code,omitempty"`
	}
	if err := c.ShouldBindJSON(&cardVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var filterCard = map[string]interface{}{
		"code": cardVals.Code,
	}

	var card models.Card
	if err := h.cardSrv.FindOne(&card, filterCard); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if card.UserID != nil {
		respondError(c, http.StatusUnprocessableEntity, errors.UseableCard.Error())
		return
	}

	if err := h.userSrv.AddCardToUser(currentUser, &card); err != nil {
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

// Update ...
func (h *CardHandler) Update(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	cardCode := c.Param("code")
	var filterCard = map[string]interface{}{
		"code":    cardCode,
		"user_id": currentUser.ID,
	}

	var card models.Card
	if err := h.cardSrv.FindOne(&card, filterCard); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var cardVals struct {
		Activated *bool `json:"activated,omitempty"`
	}
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

// Destroy ...
func (h *CardHandler) Destroy(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	paramQuery := c.Param("code")
	var filterCard = map[string]interface{}{
		"code":    paramQuery,
		"user_id": currentUser.ID,
	}

	var card models.Card
	if err := h.cardSrv.FindOne(&card, filterCard); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userSrv.RemoveCardFromUser(currentUser, &card); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "removed", Error: nil})
}
