package serializers

import "github.com/michaelt0520/nfc-card/models"

type UserResponse struct {
	Name         string          `json:"name"`
	Avatar       string          `json:"avatar"`
	Username     string          `json:"username"`
	Email        string          `json:"email"`
	Type         string          `json:"type"`
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
