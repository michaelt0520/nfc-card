package serializers

type CompanyRequest struct {
	InformationID string `uri:"id" json:"id"`
}

type CompanyCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
	Website string `json:"website"`
}

type CompanyUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
	Website string `json:"website"`
}
