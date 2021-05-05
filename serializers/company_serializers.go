package serializers

type CompanyCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Website string `json:"website" binding:"required"`
	Hotline string `json:"hotline" binding:"required"`
}

type CompanyUpdateRequest struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Logo    string `json:"logo,omitempty"`
	Website string `json:"website,omitempty"`
	Hotline string `json:"hotline,omitempty"`
}

type CompanyResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
	Website string `json:"website"`
	Hotline string `json:"hotline"`
}
