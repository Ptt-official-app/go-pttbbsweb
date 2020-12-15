package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

const LOAD_POPULAR_BOARDS_R = "/boards/popular"

type LoadPopularBoardsParams struct{}

type LoadPopularBoardsResult struct {
	List    []*types.BoardSummary `json:"list"`
	NextIdx string                `json:"next_idx"`
}

func LoadPopularBoards(remoteAddr string, userID string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.BoardListResult
	return result, 200, nil
}
