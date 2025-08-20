package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func getUserPermInfo(userID bbs.UUserID, c *gin.Context) (userPermInfo *schema.UserPermInfo, err error) {
	userPermInfo, err = schema.GetUserPermInfo(userID)
	if err != nil {
		return nil, err
	}
	if userPermInfo != nil {
		return userPermInfo, nil
	}

	updateNanoTS := types.NowNanoTS()
	_, _, err = tryGetUserInfo(userID, userID, updateNanoTS, c)
	if err != nil {
		return nil, err
	}

	return schema.GetUserPermInfo(userID)
}
