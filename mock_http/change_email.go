package mock_http

import "github.com/Ptt-official-app/go-pttbbs/api"

func ChangeEmail(params *api.ChangeEmailParams) (ret *api.ChangeEmailResult) {
	return &api.ChangeEmailResult{
		UserID: "SYSOP",
		Email:  "test@ptt.test",
	}
}
