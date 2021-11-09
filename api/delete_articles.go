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
// DeleteArticles provides function from api /board/:bid/deletearticles
// it will call backend api (go-pttbbs) and deleting all components like:
// comments, ranks, user_read_records about this articles
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

	// to go-pttbbs
	theParams_b := &pttbbsapi.DeleteArticlesParams{
		ArticleIDs: articleIDs,
	}
	var result_b *pttbbsapi.DeleteArticlesResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}

	url := utils.MergeURL(urlMap, pttbbsapi.DELETE_ARTICLES_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	// update to db
	// TODO backend response success deleted articles, if any failed, we should not delete failed one.
	updateNanoTS := types.NowNanoTS()
	err = schema.DeleteArticles(boardID, result_b.ArticleIDs, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}
	err = schema.DeleteCommentsByArticles(boardID, result_b.ArticleIDs, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}
	err = schema.DeleteRanks(boardID, result_b.ArticleIDs, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}
	err = schema.DeleteUserReadArticles(boardID, result_b.ArticleIDs, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	result = &DeleteArticlesResult{
		Success: true,
	}
	return result, 200, nil
}
