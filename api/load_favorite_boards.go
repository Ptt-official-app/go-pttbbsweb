package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mock"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_FAVORITE_BOARDS_R = "/user/:user_id/favorites"

type LoadFavoriteBoardsParams struct {
	LevelIdx  string `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
	StartIdx  string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Ascending bool   `json:"asc,omitempty"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadFavoriteBoardsPath struct {
	UserID bbs.UUserID `json:"user_id"`
}

type LoadFavoriteBoardsResult struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`
}

func NewLoadFavoriteBoardsParams() *LoadFavoriteBoardsParams {
	return &LoadFavoriteBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}

func LoadFavoriteBoardsWrapper(c *gin.Context) {
	params := NewLoadFavoriteBoardsParams()
	path := &LoadFavoriteBoardsPath{}
	LoginRequiredPathQuery(LoadFavoriteBoards, params, path, c)
}

func LoadFavoriteBoards(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = mock.BoardListResult
	return result, 200, nil
}
