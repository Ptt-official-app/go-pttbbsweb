package backend

const REGISTER_R = "/register"

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Over18 bool `json:"over18,omitempty"`

	Email    string `json:"email,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Realname string `json:"realname,omitempty"`
	Career   string `json:"career,omitempty"`
	Address  string `json:"address,omitempty"`
}

type RegisterResults struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
