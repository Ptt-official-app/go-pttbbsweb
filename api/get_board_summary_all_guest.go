package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/gin-gonic/gin"
)

func GetBoardSummaryAllGuestWrapper(c *gin.Context) {
	params := &GetBoardSummaryParams{}
	path := &GetBoardSummaryPath{}

	PathQuery(GetBoardSummaryAllGuest, params, path, c)
}

func GetBoardSummaryAllGuest(remoteAddr string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := bbs.UUserID("guest")

	thePath, ok := path.(*GetBoardSummaryPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 400, err
	}

	_, err = CheckUserBoardPermReadable(userID, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	boardSummary_db, err := schema.GetBoardSummary(boardID)
	if err != nil {
		return nil, 500, err
	}

	userBoardInfo := &apitypes.UserBoardInfo{
		Stat: ptttype.NBRD_BOARD,
	}

	boardSummary := apitypes.NewBoardSummary(boardSummary_db, "", userBoardInfo, userID)

	// result
	result = GetBoardSummaryResult(boardSummary)

	return result, 200, nil
}
