package mockhttp

import "github.com/Ptt-official-app/go-pttbbs/api"

func CheckExistsUser(params *api.CheckExistsUserParams) (ret *api.CheckExistsUserResult) {
	switch params.Username {
	case "SYSOP2":
		ret = &api.CheckExistsUserResult{
			UserID:   "SYSOP2",
			IsExists: false,
		}
	case "SYSOP":
		ret = &api.CheckExistsUserResult{
			UserID:   "SYSOP",
			IsExists: true,
		}
	default:
		ret = nil
	}

	return ret
}
