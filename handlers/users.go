package handlers

import (
	"github.com/michaelt0520/nfc-card/repositories"
)

// UserHandler : struct
type UserHandler struct {
	userRepo *repositories.UserRepository
}

// NewUserHandler ...
func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// Find ...
// func (h *UserHandler) Find(c *gin.Context) {
// 	var userVals serializers.UserRequest
// 	if err := c.BindUri(&userVals); err != nil {
// 		respondError(c, http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}

// 	var data map[string]interface{}
// 	if err := serializers.ConvertSerializer(userVals, &data); err != nil {
// 		respondError(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	res, err := h.userRepo.Find(data)
// 	if err != nil {
// 		respondError(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	if res == nil {
// 		respondError(c, http.StatusNoContent, errors.RecordNotFound.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, serializers.Resp{Result: res, Error: nil})
// }
