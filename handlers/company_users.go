package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// CompanyUserHandler : struct
type CompanyUserHandler struct {
	userRepo *repositories.UserRepository
}

// NewCompanyUserHandler ...
func NewCompanyUserHandler(userRepo *repositories.UserRepository) *CompanyUserHandler {
	return &CompanyUserHandler{
		userRepo: userRepo,
	}
}

// Index : list all users
func (h *CompanyUserHandler) Index(c *gin.Context) {
	var users []serializers.UserResponse

	if err := serializers.ConvertSerializer(h.userRepo.All(), &users); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: users, Error: nil})
}

// Create : create user
func (h *CompanyUserHandler) Create(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	var createVals serializers.CompanyUserCreateRequest
	if err := c.ShouldBindJSON(&createVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	err := serializers.ConvertSerializer(createVals, &user)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	user.CompanyID = &currentCompany.ID
	user.Type = models.Business
	user.Role = models.UserCompanyMember

	if err := h.userRepo.Create(&user); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "create success", Error: nil})
}

// Delete : soft delete user
func (h *CompanyUserHandler) Destroy(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	username := c.Param("username")
	user, err := h.userRepo.Find(map[string]interface{}{"username": username, "company_id": currentCompany.ID})
	if err != nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	if err := h.userRepo.Destroy(user); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "user deleted", Error: nil})
}
