package serializers

import "github.com/michaelt0520/nfc-card/models"

type UserSerializer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateSerializer struct {
	Name      string          `json:"name"`
	Avatar    string          `json:"avatar"`
	Username  string          `json:"username"`
	Email     string          `json:"email"`
	Type      models.CardType `json:"type"`
	CompanyID uint            `json:"company_id"`
}
