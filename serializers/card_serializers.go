package serializers

type CardCreateRequest struct {
	Code      string `json:"code" binding:"required"`
	Activated bool   `json:"activated"`
}

type CardUpdateRequest struct {
	UserID    *uint `json:"user_id,omitempty"`
	Activated *bool `json:"activated,omitempty"`
}

type CardResponse struct {
	ID        uint         `json:"id"`
	Code      string       `json:"code"`
	Activated bool         `json:"activated"`
	User      UserResponse `json:"user"`
}

type CardParametersRequest struct {
	Code *string `form:"code" json:"code,omitempty"`
}
