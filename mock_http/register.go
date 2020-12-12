package mock_http

import "github.com/Ptt-official-app/go-openbbsmiddleware/backend"

func Register(params *backend.RegisterParams) (ret *backend.RegisterResults) {
	userID := params.Username
	token, _ := createToken(userID)

	ret = &backend.RegisterResults{
		AccessToken: token,
		TokenType:   "bearer",
	}

	return ret
}
