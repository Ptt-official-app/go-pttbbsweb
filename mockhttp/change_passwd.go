package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func ChangePasswd(params *api.ChangePasswdParams) (ret *api.ChangePasswdResult) {
	userID := "SYSOP"

	token, _ := createToken(userID)

	ret = &api.ChangePasswdResult{
		UserID:    bbs.UUserID(userID),
		Jwt:       token,
		TokenType: "bearer",
	}

	return ret
}
