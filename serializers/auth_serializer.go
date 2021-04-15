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

type UserSigninResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"jwt"`
}

type SigninResponse struct {
	User UserSigninResponse `json:"user"`
}
