package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_SUMMARY_R = "/board/:bid/summary"

type GetBoardSummaryParams struct {
}

type GetBoardSummaryPath struct {
	BBoardID bbs.BBoardID `uri:"bid"`
}

type GetBoardSummaryResult struct {
	*types.BoardSummary
}

func GetBoardSummary(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	_, ok := path.(*GetBoardSummaryPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	result = &GetBoardSummaryResult{
		BoardSummary: &types.BoardSummary{
			BBoardID:  bbs.BBoardID("10_WhoAmI"),
			Brdname:   "WhoAmI",
			Title:     "我～是～誰？～",
			BrdAttr:   0,
			BoardType: "◎",
			Category:  "嘰哩",
			NUser:     39,
			BMs:       []string{"okcool", "teemo"},
			Reason:    "",
			Read:      true,
			Total:     134,
		},
	}

	return result, 200, nil
}
