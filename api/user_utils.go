package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
)

func deserializeUserDetailAndUpdateDB(user_b pttbbsapi.GetUserResult, updateNanoTS types.NanoTS) (userDetail *schema.UserDetail, err error) {

	userDetail = schema.NewUserDetail(user_b, updateNanoTS)

	err = schema.UpdateUserDetail(userDetail)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}
