package serializers

type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type SigninRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SigninResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
