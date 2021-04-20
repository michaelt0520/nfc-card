package serializers

type CompanyCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Logo    string `json:"logo" binding:"required"`
	Website string `json:"website" binding:"required"`
}

type CompanyUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
	Website string `json:"website"`
}

type CompanyResponse struct {
	Name    string         `json:"name"`
	Address string         `json:"address"`
	Logo    string         `json:"logo"`
	Website string         `json:"website"`
	Users   []UserResponse `json:"users"`
}
