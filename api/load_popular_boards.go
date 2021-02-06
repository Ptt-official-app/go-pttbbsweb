package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_POPULAR_BOARDS_R = "/boards/popular"

type LoadPopularBoardsResult struct {
	List []*apitypes.BoardSummary `json:"list"`
}

func LoadPopularBoardsWrapper(c *gin.Context) {
	LoginRequiredQuery(LoadPopularBoards, nil, c)
}

func LoadPopularBoards(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	//get data
	var result_b *pttbbsapi.LoadHotBoardsResult

	url := pttbbsapi.LOAD_HOT_BOARDS_R
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadPopularBoardsResult(boardSummaries_db)

	//check isRead
	err = checkBoardInfo(userID, userBoardInfoMap, r.List)
	if err != nil {
		return nil, 500, err
	}

	return r, 200, nil

}

func NewLoadPopularBoardsResult(boardSummaries_db []*schema.BoardSummary) *LoadPopularBoardsResult {

	theList := make([]*apitypes.BoardSummary, len(boardSummaries_db))
	for i, each_db := range boardSummaries_db {
		theList[i] = apitypes.NewBoardSummary(each_db, "")
	}

	return &LoadPopularBoardsResult{
		List: theList,
	}
}
