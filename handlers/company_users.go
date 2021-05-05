package handlers

import (
	"database/sql"
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
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	query := c.Query("q")
	var users []models.User
	if err := h.userRepo.UserTable().Where("company_id = ?", currentCompany.ID, sql.Named("query", query)).Find(&users).Error; err != nil {
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

// Update : remove user from company
func (h *CompanyUserHandler) Update(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	var filterUser = map[string]interface{}{
		"company_id": currentCompany.ID,
		"username":   c.Param("username"),
	}

	var user []models.User
	if errUser := h.userRepo.Find(&user, filterUser); errUser != nil {
		respondError(c, http.StatusUnauthorized, errUser.Error())
		return
	}

	if err := repositories.GetDB().Model(&currentCompany).Association("Users").Delete(user); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "removed", Error: nil})
}

// ShowPersonalUsers
func (h *CompanyUserHandler) ShowPersonalUsers(c *gin.Context) {
	query := c.Query("q")

	var users []models.User
	if err := h.userRepo.UserTable().Where("role = ? and (name like '%@query%' or email like '%@query%' or phone_number like '%@query%')", models.Personal, sql.Named("query", query)).Find(&users).Error; err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result []serializers.UserResponse
	if err := serializers.ConvertSerializer(users, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}
