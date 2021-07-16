package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func AttemptChangeEmail(params *api.AttemptChangeEmailParams) (ret *api.AttemptChangeEmailResult) {
	jwt, _ := api.CreateEmailToken("SYSOP", "", params.Email, api.CONTEXT_CHANGE_EMAIL)

	ret = &api.AttemptChangeEmailResult{
		UserID: "SYSOP",
		Jwt:    jwt,
	}

	return ret
}
