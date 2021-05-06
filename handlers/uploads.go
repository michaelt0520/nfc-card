package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/serializers"
	"github.com/michaelt0520/nfc-card/services"
)

// UploadHandler : struct
type UploadHandler struct {
	userSrv *services.UserService
	compSrv *services.CompanyService
}

// NewUploadHandler ...
func NewUploadHandler(userSrv *services.UserService, compSrv *services.CompanyService) *UploadHandler {
	return &UploadHandler{
		userSrv: userSrv,
		compSrv: compSrv,
	}
}

// Avatar : upload avatar
func (h *UploadHandler) Avatar(c *gin.Context) {
	currentUser, err := h.userSrv.GetCurrentUser(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
	}

	userAvatarPath := fmt.Sprintf("./public/avatars/%v", currentUser.ID)
	if err := os.MkdirAll(userAvatarPath, os.ModePerm); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	filename := fmt.Sprintf("%s/%s", userAvatarPath, file.Filename)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := fmt.Sprintf("%s/public/avatars/%v/%s", os.Getenv("app_host"), currentUser.ID, file.Filename)

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}

// Logo : upload company logo
func (h *UploadHandler) Logo(c *gin.Context) {
	currentCompany, err := h.compSrv.GetCurrentCompany(c)
	if err != nil {
		respondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
	}

	CompanyLogoPath := fmt.Sprintf("./public/logos/%v", currentCompany.ID)
	if err := os.MkdirAll(CompanyLogoPath, os.ModePerm); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	filename := fmt.Sprintf("%s/%s", CompanyLogoPath, file.Filename)

	if err = c.SaveUploadedFile(file, filename); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := fmt.Sprintf("%s/public/logos/%v/%s", os.Getenv("app_host"), currentCompany.ID, file.Filename)

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}
