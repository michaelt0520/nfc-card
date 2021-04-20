package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"github.com/michaelt0520/nfc-card/serializers"
)

// AdminUserHandler : struct
type AdminUserHandler struct {
	userRepo *repositories.UserRepository
}

// NewAdminUserHandler ...
func NewAdminUserHandler(userRepo *repositories.UserRepository) *AdminUserHandler {
	return &AdminUserHandler{
		userRepo: userRepo,
	}
}

// Index : list all users
func (h *AdminUserHandler) Index(c *gin.Context) {
	var users []serializers.UserResponse

	if err := serializers.ConvertSerializer(h.userRepo.All(), &users); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: users, Error: nil})
}

// Show : Get user by username
func (h *AdminUserHandler) Show(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userRepo.Find(map[string]interface{}{"username": username})
	if err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(user, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	resUser.Type = user.TypeToString()

	c.JSON(http.StatusOK, serializers.Resp{Result: resUser, Error: nil})
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

	if err := h.userRepo.Create(&user); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: "create success", Error: nil})
}

// Update ...
func (h *AdminUserHandler) Update(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userRepo.Find(map[string]interface{}{"username": username})
	if err != nil {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
	}

	// get data from body
	var userVals serializers.UserUpdateRequest
	if err := c.ShouldBindJSON(&userVals); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	if err := serializers.ConvertSerializer(userVals, &data); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userRepo.Update(user, data); err != nil {
		respondError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var resUser serializers.UserResponse
	if err := serializers.ConvertSerializer(user, &resUser); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	resUser.Type = user.TypeToString()
	resUser.Role = user.RoleToString()

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
	username := c.Param("username")
	user, err := h.userRepo.Find(map[string]interface{}{"username": username})
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
