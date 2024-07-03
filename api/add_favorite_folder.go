package api

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/gin-gonic/gin"
)

const ADD_FAVORITE_FOLDER_R = "/user/:user_id/favorites/addfolder"

type AddFavoriteFolderParams struct {
	LevelIdx schema.LevelIdx `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
	Title    string          `json:"title,omitempty" form:"title,omitempty" url:"title,omitempty"`
}

type AddFavoriteFolderPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

func AddFavoriteFolderWrapper(c *gin.Context) {
	params := &AddFavoriteFolderParams{}
	path := &AddFavoriteFolderPath{}
	LoginRequiredPathJSON(AddFavoriteLine, params, path, c)
}

func AddFavoriteFolder(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*AddFavoriteFolderParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*AddFavoriteFolderPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	if userID != thePath.UserID {
		return nil, 403, ErrInvalidUser
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

	title := theParams.Title
	if title == "" {
		title = DEFAULT_FAV_TITLE
	}

	theIdx, _, err := theFav.AddFolder(title)
	if err != nil {
		return nil, 500, err
	}

	return postAddFavorite(userID, rootFav, remoteAddr, theParams.LevelIdx, theIdx, c, false)
}
