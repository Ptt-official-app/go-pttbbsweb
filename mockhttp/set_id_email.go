package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func SetIDEmail(params *api.SetIDEmailParams) *api.SetIDEmailResult {
	return &api.SetIDEmailResult{
		UserID: "SYSOP",
		Email:  "test@ptt.test",
	}
}
