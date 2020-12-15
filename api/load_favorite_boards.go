package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/gin-gonic/gin"
)

const LOAD_FAVORITE_BOARDS_R = "/user/:user_id/favorites"

type LoadFavoriteBoardsParams struct{}

type LoadFavoriteBoardsPath struct {
	UserID string `json:"user_id"`
}

type LoadFavoriteBoardsResult struct {
	List    []*types.BoardSummary `json:"list"`
	NextIdx string                `json:"next_idx"`
}

func LoadFavoriteBoards(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.BoardListResult
	return result, 200, nil
}
