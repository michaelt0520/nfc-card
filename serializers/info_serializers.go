package serializers

type InfoCreateRequest struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	Visibled bool   `json:"visibled"`
}

type InfoUpdateRequest struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	Visibled bool   `json:"visibled"`
}

type InfoResponse struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	Visibled bool   `json:"visibled"`
}
