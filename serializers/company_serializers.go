package serializers

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

type CompanyResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
	Website string `json:"website"`
}
