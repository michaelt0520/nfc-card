package serializers

type UserRequest struct {
	Username string `uri:"username" json:"username"`
}
