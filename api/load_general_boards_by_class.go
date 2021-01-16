package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_GENERAL_BOARDS_BY_CLASS_R = "/boards/byclass"

func LoadGeneralBoardsByClassWrapper(c *gin.Context) {
	params := NewLoadGeneralBoardsParams()
	LoginRequiredQuery(LoadGeneralBoardsByClass, params, c)
}

func LoadGeneralBoardsByClass(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	//backend load-general-baords
	theParams_b := &pttbbsapi.LoadGeneralBoardsParams{
		StartIdx: theParams.StartIdx,
		Keyword:  types.Utf8ToBig5(theParams.Keyword),
		NBoards:  theParams.Max,
	}
	var result_b *pttbbsapi.LoadGeneralBoardsResult

	url := pttbbsapi.LOAD_GENERAL_BOARDS_BY_CLASS_R
	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadGeneralBoardsResult(boardSummaries_db, result_b.NextIdx)

	//check isRead
	err = checkBoardInfo(userID, userBoardInfoMap, r.List)
	if err != nil {
		return nil, 500, err
	}

	return r, 200, nil
}
