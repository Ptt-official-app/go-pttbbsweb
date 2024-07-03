package api

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/gin-gonic/gin"
)

const ADD_FAVORITE_BOARD_R = "/user/:user_id/favorites/addboard"

type AddFavoriteBoardParams struct {
	LevelIdx schema.LevelIdx   `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
	FBoardID apitypes.FBoardID `json:"bid" form:"bid" url:"bid"`
}

type AddFavoriteBoardPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type AddFavoriteResult *apitypes.BoardSummary

func AddFavoriteBoardWrapper(c *gin.Context) {
	params := &AddFavoriteBoardParams{}
	path := &AddFavoriteBoardPath{}
	LoginRequiredPathJSON(AddFavoriteBoard, params, path, c)
}

func AddFavoriteBoard(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*AddFavoriteBoardParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*AddFavoriteBoardPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	if userID != thePath.UserID {
		return nil, 403, ErrInvalidUser
	}

	boardID, err := toBoardID(theParams.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	_, statusCode, err = isBoardValidUser(boardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	bid, _, err := boardID.ToRaw()
	if err != nil {
		return nil, 500, err
	}

	userFavoritesMeta, userFavorites, err := tryGetAllUserFavorites(userID, c)
	if err != nil {
		return nil, 500, err
	}

	rootFav, err := schema.UserFavoritesToFav(&userFavoritesMeta.FolderMeta, userFavorites, 0)
	if err != nil {
		return nil, 500, err
	}

	levelIdxList := strings.Split(string(theParams.LevelIdx), ":")
	theFav, err := rootFav.LocateFav(levelIdxList)
	if err != nil {
		return nil, 500, err
	}

	theIdx, _, err := theFav.AddBoard(bid)
	if err != nil {
		return nil, 500, err
	}

	return postAddFavorite(userID, rootFav, remoteAddr, theParams.LevelIdx, theIdx, c, true)
}
