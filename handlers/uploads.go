package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/serializers"
)

// UploadHandler : struct
type UploadHandler struct{}

// NewUploadHandler ...
func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// Avatar : upload avatar
func (h *UploadHandler) Avatar(c *gin.Context) {
	// get currentUser
	user, ok := c.Get("currentUser")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentUser := user.(*models.User)

	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
	}

	filename := fmt.Sprintf("./public/avatars/%s/%s", currentUser.Username, file.Filename)

	if err = c.SaveUploadedFile(file, filename); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := fmt.Sprintf("%s/public/avatars/%s/%s", os.Getenv("app_host"), currentUser.Username, file.Filename)

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}

// Logo : upload company logo
func (h *UploadHandler) Logo(c *gin.Context) {
	// get currentCompany
	company, ok := c.Get("currentCompany")
	if !ok {
		respondError(c, http.StatusUnauthorized, errors.RecordNotFound.Error())
		return
	}
	currentCompany := company.(*models.Company)

	file, err := c.FormFile("file")
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
	}

	filename := fmt.Sprintf("./public/logos/%s/%s", strconv.Itoa(int(currentCompany.ID)), file.Filename)

	if err = c.SaveUploadedFile(file, filename); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result := fmt.Sprintf("%s/public/logos/%s/%s", os.Getenv("app_host"), strconv.Itoa(int(currentCompany.ID)), file.Filename)

	c.JSON(http.StatusOK, serializers.Resp{Result: result, Error: nil})
}
