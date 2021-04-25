package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// AppHandler : struct
type AppHandler struct {
	cardRepo    *repositories.CardRepository
	contactRepo *repositories.ContactRepository
}

// NewAppHandler ...
func NewAppHandler(cardRepo *repositories.CardRepository, contactRepo *repositories.ContactRepository) *AppHandler {
	return &AppHandler{
		cardRepo:    cardRepo,
		contactRepo: contactRepo,
	}
}

// ShowCard ...
func (h *AppHandler) ShowCard(c *gin.Context) {
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

// CreateContact
func (h *AppHandler) CreateContact(c *gin.Context) {
	// get create data from body
	var contactVals serializers.ContactCreateRequest
	if err := c.ShouldBindJSON(&contactVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// binding data to model
	var contact models.Contact
	err := serializers.ConvertSerializer(contactVals, &contact)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// save to db
	if err := h.contactRepo.Create(&contact); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &contact, Error: nil})
}
