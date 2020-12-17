package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_FAVORITE_BOARDS_R = "/user/:user_id/favorites"

type LoadFavoriteBoardsParams struct {
	LevelIdx string `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
	StartIdx string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max      int    `json:"max,omitempty" form:"max,omitempty" url:"max,omitempty"`
}

type LoadFavoriteBoardsPath struct {
	UserID bbs.UUserID `json:"user_id"`
}

type LoadFavoriteBoardsResult struct {
	List    []*types.BoardSummary `json:"list"`
	NextIdx string                `json:"next_idx"`
}

func LoadFavoriteBoards(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.BoardListResult
	return result, 200, nil
}
