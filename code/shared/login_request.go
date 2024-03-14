package shared

type LoginRequest struct {
	Username string
	Password string
}

func GetNewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}
