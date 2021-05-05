package serializers

import "github.com/michaelt0520/nfc-card/models"

type UserResponse struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Avatar       string          `json:"avatar"`
	Username     string          `json:"username"`
	Email        string          `json:"email"`
	Address      string          `json:"address"`
	PhoneNumber  string          `json:"phone_number"`
	Type         string          `json:"type"`
	Role         string          `json:"role"`
	Company      CompanyResponse `json:"company"`
	Informations []InfoResponse  `json:"informations"`
}

type UserUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Avatar      *string `json:"avatar,omitempty"`
	Address     *string `json:"address,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

type UserUpdatePasswordRequest struct {
	OldPassowrd          *string `json:"old_password" binding:"required"`
	NewPassword          *string `json:"new_password" binding:"required"`
	PasswordConfirmation *string `json:"password_confirmation" binding:"required"`
}

type AdminUserCreateRequest struct {
	Email       string          `json:"email" binding:"required"`
	Password    string          `json:"password" binding:"required"`
	Username    string          `json:"username" binding:"required"`
	Name        string          `json:"name" binding:"required"`
	Address     string          `json:"address"`
	PhoneNumber string          `json:"phone_number"`
	Type        models.CardType `json:"type" binding:"required"`
	Role        models.UserRole `json:"role" binding:"required"`
}
