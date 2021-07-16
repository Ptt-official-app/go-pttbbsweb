package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func CheckExistsUser(params *api.CheckExistsUserParams) (ret *api.CheckExistsUserResult) {
	if params.Username == "SYSOP2" {
		ret = &api.CheckExistsUserResult{
			UserID:   "SYSOP2",
			IsExists: false,
		}
	} else if params.Username == "SYSOP" {
		ret = &api.CheckExistsUserResult{
			UserID:   "SYSOP",
			IsExists: true,
		}
	} else {
		ret = nil
	}

	return ret
}
