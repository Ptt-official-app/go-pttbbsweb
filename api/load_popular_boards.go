package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_POPULAR_BOARDS_R = "/boards/popular"

type LoadPopularBoardsParams struct {
	StartIdx  string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Ascending bool   `json:"asc,omitempty"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadPopularBoardsResult struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`
}

func NewLoadPopularBoardsParams() *LoadPopularBoardsParams {
	return &LoadPopularBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}
func LoadPopularBoardsWrapper(c *gin.Context) {
	params := NewLoadPopularBoardsParams()
	LoginRequiredQuery(LoadPopularBoards, params, c)
}

func LoadPopularBoards(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.BoardListResult
	return result, 200, nil
}
