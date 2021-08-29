package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func IsBoardsValidUser(params *api.IsBoardsValidUserParams) (ret *api.IsBoardsValidUserResult) {
	ret = &api.IsBoardsValidUserResult{
		IsValid: map[bbs.BBoardID]bool{
			"1_test1": true,
			"2_test2": true,
		},
	}
	return ret
}
