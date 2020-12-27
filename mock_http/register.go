package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func Register(params *api.RegisterParams) (ret *api.RegisterResult) {
	userID := params.Username
	token, _ := createToken(userID)

	ret = &api.RegisterResult{
		UserID:    bbs.UUserID(userID),
		Jwt:       token,
		TokenType: "bearer",
	}

	return ret
}
