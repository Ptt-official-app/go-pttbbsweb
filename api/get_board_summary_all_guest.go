package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetBoardSummaryAllGuestWrapper(c *gin.Context) {
	params := &GetBoardSummaryParams{}
	path := &GetBoardSummaryPath{}

	PathQuery(GetBoardSummaryAllGuest, params, path, c)
}

func GetBoardSummaryAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	return GetBoardSummary(remoteAddr, userID, params, path, c)
}
