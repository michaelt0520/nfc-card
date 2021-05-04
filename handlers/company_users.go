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

	// get parameters
	var paramVals serializers.UserParametersRequest
	if err := c.ShouldBind(&paramVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(paramVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	data["CompanyID"] = currentCompany.ID

	var users []models.User
	if _, err := h.userRepo.Where(&users, data, repositories.Paginate(c)); err != nil {
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

	if _, err := h.userRepo.Create(&user); err != nil {
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
	var user models.User
	if _, err := h.userRepo.Where(&user, map[string]interface{}{"username": username, "company_id": currentCompany.ID}); err != nil {
		respondError(c, http.StatusNotFound, errors.RecordNotFound.Error())
		return
	}

	if _, err := h.userRepo.Destroy(&user); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "user deleted", Error: nil})
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
