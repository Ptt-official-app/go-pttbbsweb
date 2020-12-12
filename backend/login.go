package backend

const LOGIN_R = "/token"

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResults struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
