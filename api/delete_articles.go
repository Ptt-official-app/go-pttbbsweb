package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const DELETE_ARTICLES_R = "/board/:bid/deletearticles"

type DeleteArticlesParams struct {
	ArticleIDs []apitypes.FArticleID `json:"aids" form:"aids" url:"aids" binding:"required"`
}

type DeleteArticlesPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
}

type DeleteArticlesResult struct {
	Success bool `json:"success"`
}

func DeleteArticlesWrapper(c *gin.Context) {
	params := &DeleteArticlesParams{}
	path := &DeleteArticlesPath{}
	LoginRequiredPathJSON(DeleteArticles, params, path, c)
}

func DeleteArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*DeleteArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*DeleteArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	var articleIDs []bbs.ArticleID
	for _, articleID := range theParams.ArticleIDs {
		articleIDs = append(articleIDs, articleID.ToArticleID())
	}

	result = &DeleteArticlesResult{
		Success: true,
	}
	return result, 200, nil
}
