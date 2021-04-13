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
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Visibled bool   `json:"visibled"`
}
