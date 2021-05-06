package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// CompanyUserHandler : struct
type CompanyUserHandler struct {
	compSrv *services.CompanyService
	userSrv *services.UserService
}

// NewCompanyUserHandler ...
func NewCompanyUserHandler(compSrv *services.CompanyService, userSrv *services.UserService) *CompanyUserHandler {
	return &CompanyUserHandler{
		compSrv: compSrv,
		userSrv: userSrv,
	}
}

// Index : list all users
func (h *CompanyUserHandler) Index(c *gin.Context) {
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	paramQuery := c.Query("q")
	var filterUser = map[string]interface{}{
		"company_id": currentCompany.ID,
		"name":       paramQuery,
	}

	var users []models.User
	if err := h.userSrv.FindMany(&users, filterUser, c); err != nil {
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
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	var filterUser = map[string]interface{}{
		"company_id": currentCompany.ID,
		"username":   c.Param("username"),
	}

	var user models.User
	if errUser := h.userSrv.FindOne(&user, filterUser); errUser != nil {
		respondError(c, http.StatusUnauthorized, errUser.Error())
		return
	}

	if err := h.compSrv.RemoveUserFromCompany(currentCompany, &user); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "removed", Error: nil})
}

// ShowPersonalUsers
func (h *CompanyUserHandler) ShowPersonalUsers(c *gin.Context) {
	paramQuery := c.Query("q")
	var filterUser = map[string]interface{}{
		"role": models.Personal,
		"name": paramQuery,
	}

	var users []models.User
	if err := h.userSrv.FindMany(&users, filterUser, c); err != nil {
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
