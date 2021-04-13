package serializers

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninResponse struct {
	User  UserSerializer
	Token string
}
