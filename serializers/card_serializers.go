package serializers

type CardCreateRequest struct {
	Code      string `json:"code" binding:"required"`
	Activated bool   `json:"activated"`
}

type CardUpdateRequest struct {
	Code      string `json:"code"`
	UserID    uint   `json:"user_id"`
	Activated bool   `json:"activated"`
}

type CardResponse struct {
	Code      string       `json:"code"`
	Activated bool         `json:"activated"`
	User      UserResponse `json:"user"`
}
