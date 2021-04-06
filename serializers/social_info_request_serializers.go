package serializers

type SocialInfoRequest struct {
	SocialInformationID string `uri:"id" json:"id"`
}

type SocialInfoCreateRequest struct {
	Provider  string `json:"provider"`
	Data      string `json:"data"`
	IsPrivate bool   `json:"is_private"`
}
