package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func IsBoardValidUser() (ret *api.IsBoardValidUserResult) {
	ret = &api.IsBoardValidUserResult{
		IsValid: true,
	}

	return ret
}
