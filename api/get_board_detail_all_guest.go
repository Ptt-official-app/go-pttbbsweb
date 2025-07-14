package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetBoardDetailAllGuestWrapper(c *gin.Context) {
	params := &GetBoardDetailParams{}
	path := &GetBoardDetailPath{}
	PathQuery(GetBoardDetailAllGuest, params, path, c)
}

func GetBoardDetailAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return GetBoardDetail(remoteAddr, userID, params, path, c)
}
