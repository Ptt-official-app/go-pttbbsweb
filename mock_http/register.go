package mock_http

import "github.com/Ptt-official-app/go-openbbsmiddleware/backend"

func Register(params *backend.RegisterParams) (ret *backend.RegisterResult) {
	userID := params.UserID
	token, _ := createToken(userID)

	ret = &backend.RegisterResult{
		Jwt:       token,
		TokenType: "bearer",
	}

	return ret
}
