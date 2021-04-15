package serializers

type InfoRequest struct {
	InformationID string `uri:"id" json:"id"`
}

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
