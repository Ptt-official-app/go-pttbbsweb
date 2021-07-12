package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const GET_USER_ID_R = "/userid"

type GetUserIDResult struct {
	UserID bbs.UUserID `json:"user_id"`
}

func GetUserIDWrapper(c *gin.Context) {
	LoginRequiredQuery(GetUserID, nil, c)
}

func GetUserID(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	return &GetUserIDResult{
		UserID: userID,
	}, 200, nil
}
