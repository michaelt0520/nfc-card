package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// AdminUserHandler : struct
type AdminUserHandler struct {
	userSrv *services.UserService
}

// NewAdminUserHandler ...
func NewAdminUserHandler(userSrv *services.UserService) *AdminUserHandler {
	return &AdminUserHandler{
		userSrv: userSrv,
	}
}

// Index : list all users
func (h *AdminUserHandler) Index(c *gin.Context) {
	paramQuery := c.Query("q")
	var filterUser = map[string]interface{}{
		"name": paramQuery,
	}

	var filterSearch = map[string]interface{}{
		"name":         paramQuery,
		"email":        paramQuery,
		"phone_number": paramQuery,
	}

	var users []models.User
	if err := h.userSrv.FindMany(&users, filterUser, filterSearch, c); err != nil {
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

// Show : Get user by username
func (h *AdminUserHandler) Show(c *gin.Context) {
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var preloadData = map[string]interface{}{
		"Informations": nil,
		"Company":      nil,
	}

	var user models.User
	if err := h.userSrv.FindOneWithScopes(&user, filterUser, preloadData); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var result serializers.UserResponse
	if err := serializers.ConvertSerializer(user, &result); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}

// Create : create user
func (h *AdminUserHandler) Create(c *gin.Context) {
	var createVals serializers.AdminUserCreateRequest
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

	if err := h.userSrv.Repo().Create(&user); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "create success", Error: nil})
}

// Update ...
func (h *AdminUserHandler) Update(c *gin.Context) {
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var user models.User
	if err := h.userSrv.FindOne(&user, filterUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// get data from body
	var userVals serializers.AdminUserUpdateRequest
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(userVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userSrv.Repo().Update(&user, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(user, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
}

// Upload ...
func (h *AdminUserHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
	}

	filename := fmt.Sprintf("./public/avatars/%s", file.Filename)

	if err = c.SaveUploadedFile(file, filename); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := fmt.Sprintf("%s/public/avatars/%s", os.Getenv("app_host"), file.Filename)

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}

// Delete : soft delete user
func (h *AdminUserHandler) Destroy(c *gin.Context) {
	var filterUser = map[string]interface{}{
		"username": c.Param("username"),
	}

	var user models.User
	if err := h.userSrv.FindOne(&user, filterUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userSrv.Repo().Destroy(&user); err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "user deleted", Error: nil})
}
