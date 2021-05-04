package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// AppHandler : struct
type AppHandler struct {
	cardRepo    *repositories.CardRepository
	contactRepo *repositories.ContactRepository
	userRepo    *repositories.UserRepository
}

// NewAppHandler ...
func NewAppHandler(cardRepo *repositories.CardRepository, contactRepo *repositories.ContactRepository, userRepo *repositories.UserRepository) *AppHandler {
	return &AppHandler{
		cardRepo:    cardRepo,
		contactRepo: contactRepo,
		userRepo:    userRepo,
	}
}

// ShowCard ...
func (h *AppHandler) ShowCard(c *gin.Context) {
	var resCard models.Card
	if _, err := h.cardRepo.Where(&resCard, map[string]interface{}{"code": c.Param("code")}); err != nil {
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

	c.JSON(http.StatusOK, serializers.Resp{Result: &card, Error: nil})
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
	if err := serializers.ConvertSerializer(contactVals, &contact); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// save to db
	if _, err := h.contactRepo.Create(&contact); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &contact, Error: nil})
}

// Index : list all users
func (h *AppHandler) SearchUser(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		respondError(c, http.StatusBadRequest, errors.ParameterInvalid.Error())
		return
	}

	var users []models.User
	if err := h.userRepo.UserTable().Where("name like '%?%'", query).Or("email = '%?%'", query).Or("phone_number = '%?%'", query).Find(&users).Error; err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result []serializers.UserResponse
	if err := serializers.ConvertSerializer(users, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &result, Error: nil})
}
