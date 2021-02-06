package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_SUMMARY_R = "/board/:bid/summary"

type GetBoardSummaryParams struct {
}

type GetBoardSummaryPath struct {
	BBoardID bbs.BBoardID `uri:"bid"`
}

type GetBoardSummaryResult *apitypes.BoardSummary

func GetBoardSummaryWrapper(c *gin.Context) {
	params := &GetBoardSummaryParams{}
	path := &GetBoardSummaryPath{}
	LoginRequiredPathQuery(GetBoardSummary, params, path, c)
}

func GetBoardSummary(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	thePath, ok := path.(*GetBoardSummaryPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	//backend get-board-summary
	theParams_b := &pttbbsapi.LoadBoardSummaryParams{}

	var result_b pttbbsapi.LoadBoardSummaryResult

	urlMap := map[string]string{
		"bid": string(thePath.BBoardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_BOARD_SUMMARY_R)
	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update to db
	theList_b := []*bbs.BoardSummary{(*bbs.BoardSummary)(result_b)}
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, theList_b, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	boardSummary_db := boardSummaries_db[0]
	boardSummary := apitypes.NewBoardSummary(boardSummary_db, "")
	theList := []*apitypes.BoardSummary{boardSummary}

	//check isRead
	err = checkBoardInfo(userID, userBoardInfoMap, theList)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = GetBoardSummaryResult(boardSummary)

	return result, 200, nil
}
