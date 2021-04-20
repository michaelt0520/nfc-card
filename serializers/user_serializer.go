package serializers

import "github.com/michaelt0520/nfc-card/models"

type UserResponse struct {
	Name         string          `json:"name"`
	Avatar       string          `json:"avatar"`
	Username     string          `json:"username"`
	Email        string          `json:"email"`
	Type         string          `json:"type"`
  Role         string          `json:"role"`
	Company      CompanyResponse `json:"company"`
	Informations []InfoResponse  `json:"informations"`
}

type UserUpdateRequest struct {
	Name      string          `json:"name"`
	Avatar    string          `json:"avatar"`
	Username  string          `json:"username"`
	Email     string          `json:"email"`
	Type      models.CardType `json:"type"`
	CompanyID uint            `json:"company_id"`
}

type AdminUserCreateRequest struct {
	Avatar   string          `json:"avatar" binding:"required"`
	Email    string          `json:"email" binding:"required"`
	Password string          `json:"password" binding:"required"`
	Username string          `json:"username" binding:"required"`
	Name     string          `json:"name" binding:"required"`
	Type     models.CardType `json:"type" binding:"required"`
	Role     models.UserRole `json:"role" binding:"required"`
}

type CompanyUserCreateRequest struct {
	Avatar   string          `json:"avatar" binding:"required"`
	Email    string          `json:"email" binding:"required"`
	Password string          `json:"password" binding:"required"`
	Username string          `json:"username" binding:"required"`
	Name     string          `json:"name" binding:"required"`
	Type     models.CardType `json:"type" binding:"required"`
}
