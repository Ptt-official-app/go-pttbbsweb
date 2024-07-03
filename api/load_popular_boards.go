package api

import (
	"context"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/boardd"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
)

const LOAD_POPULAR_BOARDS_R = "/boards/popular"

type LoadPopularBoardsResult struct {
	List []*apitypes.BoardSummary `json:"list"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func LoadPopularBoardsWrapper(c *gin.Context) {
	LoginRequiredQuery(LoadPopularBoards, nil, c)
}

func LoadPopularBoards(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	// get data
	ctx := context.Background()
	req := &boardd.HotboardRequest{}
	resp, err := boardd.Cli.Hotboard(ctx, req)
	if err != nil {
		return nil, 500, err
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializePBBoardsAndUpdateDB(resp.Boards, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	// check isRead
	userBoardInfoMap, err = checkUserReadBoard(userID, userBoardInfoMap, boardSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	userBoardInfoMap, err = checkUserFavBoard(userID, userBoardInfoMap, boardSummaries_db, c)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadPopularBoardsResult(boardSummaries_db, userBoardInfoMap, userID)

	return r, 200, nil
}

func NewLoadPopularBoardsResult(boardSummaries_db []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, userID bbs.UUserID) *LoadPopularBoardsResult {
	theList := make([]*apitypes.BoardSummary, len(boardSummaries_db))
	for i, each_db := range boardSummaries_db {
		userBoardInfo := userBoardInfoMap[each_db.BBoardID]
		theList[i] = apitypes.NewBoardSummary(each_db, "", userBoardInfo, "")
	}

	return &LoadPopularBoardsResult{
		List:      theList,
		TokenUser: userID,
	}
}
