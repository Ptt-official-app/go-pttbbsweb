package mock_http

import "github.com/Ptt-official-app/go-pttbbs/api"

func AttemptSetIDEmail(params *api.AttemptSetIDEmailParams) (ret *api.AttemptSetIDEmailResult) {
	jwt, _ := api.CreateEmailToken("SYSOP", "", params.Email, api.CONTEXT_SET_ID_EMAIL)

	ret = &api.AttemptSetIDEmailResult{
		UserID: "SYSOP",
		Jwt:    jwt,
	}

	return ret
}
