package serializers

import "github.com/michaelt0520/nfc-card/models"

type CardRequest struct {
	CardCode string `uri:"code" json:"code"`
}

type CardCreateRequest struct {
	Code      string `json:"code"`
	Activated bool   `json:"activated"`
}

type CardUpdateRequest struct {
	Code      string `json:"code"`
	UserID    uint   `json:"user_id"`
	Activated bool   `json:"activated"`
}

type CardResponse struct {
	Card *models.Card
	User *models.User
}
