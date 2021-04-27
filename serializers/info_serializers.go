package serializers

type InfoCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Data     string `json:"data" binding:"required"`
	Visibled bool   `json:"visibled"`
}

type InfoUpdateRequest struct {
	Name     *string `json:"name,omitempty"`
	Data     *string `json:"data,omitempty"`
	Visibled *bool   `json:"visibled,omitempty"`
}

type InfoResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Visibled bool   `json:"visibled"`
}
