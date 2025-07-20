package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/gin-gonic/gin"
)

func LoadPopularBoardsAllGuestWrapper(c *gin.Context) {
	Query(LoadPopularBoardsAllGuest, nil, c)
}

func LoadPopularBoardsAllGuest(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	boardSummaries_db, err := schema.GetPopularBoardSummaries()
	if err != nil {
		return nil, 500, err
	}

	userBoardInfoMap := make(map[bbs.BBoardID]*apitypes.UserBoardInfo)
	for _, board := range boardSummaries_db {
		userBoardInfoMap[board.BBoardID] = &apitypes.UserBoardInfo{
			Stat: ptttype.NBRD_BOARD,
		}
	}

	result = NewLoadPopularBoardsResult(boardSummaries_db, userBoardInfoMap, "guest")

	return result, 200, nil
}
