package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// CardHandler : struct
type CardHandler struct {
	cardRepo *repositories.CardRepository
	userRepo *repositories.UserRepository
}

// NewCardHandler ...
func NewCardHandler(cardRepo *repositories.CardRepository, userRepo *repositories.UserRepository) *CardHandler {
	return &CardHandler{
		cardRepo: cardRepo,
		userRepo: userRepo,
	}
}

// Find ...
func (h *CardHandler) Find(c *gin.Context) {
	var cardVals serializers.CardRequest
	if err := c.BindUri(&cardVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resCard, err := h.cardRepo.Find(cardVals.CardCode)
	if err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}
	if resCard == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	var userParams = make(map[string]interface{})
	userParams["id"] = resCard.UserID

	resUser, err := h.userRepo.Find(userParams)
	if err != nil {
		respondError(c, http.StatusNotFound, err.Error())
		return
	}
	if resUser == nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	response := serializers.CardResponse{
		Card: resCard,
		User: resUser,
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: response, Error: nil})
}

// Create ...
func (h *CardHandler) Create(c *gin.Context) {
	var cardVals serializers.CardCreateRequest
	if err := c.ShouldBindJSON(&cardVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var card models.Card
	err := serializers.ConvertSerializer(cardVals, &card)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
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
	var cardVals serializers.CardUpdateRequest
	if err := c.BindUri(&cardVals); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var data map[string]interface{}
	err := serializers.ConvertSerializer(cardVals, &data)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	resCard, err := h.cardRepo.Update(data)
	if resCard != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if resCard == nil {
		respondError(c, http.StatusNoContent, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &resCard, Error: nil})
}
