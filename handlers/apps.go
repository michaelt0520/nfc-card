package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// AppHandler : struct
type AppHandler struct {
	cardSrv    *services.CardService
	contactSrv *services.ContactService
	userSrv    *services.UserService
}

// NewAppHandler ...
func NewAppHandler(cardSrv *services.CardService, contactSrv *services.ContactService, userSrv *services.UserService) *AppHandler {
	return &AppHandler{
		cardSrv:    cardSrv,
		contactSrv: contactSrv,
		userSrv:    userSrv,
	}
}

// ShowCard ...
func (h *AppHandler) ShowCard(c *gin.Context) {
	var filterCard = map[string]interface{}{
		"code": c.Param("code"),
	}

	var preloadData = map[string]interface{}{
		"User":         nil,
		"User.Company": nil,
	}

	var resCard models.Card
	if err := h.cardSrv.FindOneWithScopes(&resCard, filterCard, preloadData); err != nil {
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
	if err := h.contactSrv.Repo().Create(&contact); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: &contact, Error: nil})
}

// Index : list all users
func (h *AppHandler) SearchUser(c *gin.Context) {
	query := c.Query("q")

	var filterUser = map[string]interface{}{
		"name": query,
	}

	var users []models.User
	if err := h.userSrv.FindMany(&users, filterUser); err != nil {
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
