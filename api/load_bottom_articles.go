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

const LOAD_BOTTOM_ARTICLES_R = "/board/:bid/articles/bottom"

type LoadBottomArticlesPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type LoadBottomArticlesResult struct {
	List []*apitypes.ArticleSummary `json:"list"`
}

func LoadBottomArticlesWrapper(c *gin.Context) {
	path := &LoadBottomArticlesPath{}
	LoginRequiredPathQuery(LoadBottomArticles, nil, path, c)
}

func LoadBottomArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	thePath, ok := path.(*LoadBottomArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	// backend load-bottom-articles
	var result_b *pttbbsapi.LoadBottomArticlesResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_BOTTOM_ARTICLES_R)
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	articleSummaries_db, userReadArticleMap, err := deserializeArticlesAndUpdateDB(userID, boardID, result_b.Articles, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	userReadArticleMap, err = checkReadArticles(userID, boardID, userReadArticleMap, articleSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadBottomArticlesResult(articleSummaries_db, userReadArticleMap, result_b)

	return r, 200, nil
}

func NewLoadBottomArticlesResult(a_db []*schema.ArticleSummaryWithRegex, userReadArticleMap map[bbs.ArticleID]bool, result_b *pttbbsapi.LoadBottomArticlesResult) *LoadBottomArticlesResult {
	theList := make([]*apitypes.ArticleSummary, len(a_db))
	for i, each_db := range a_db {
		theList[i] = apitypes.NewArticleSummaryFromWithRegex(each_db)
		articleID := each_db.ArticleID
		isRead, ok := userReadArticleMap[articleID]
		if ok && isRead {
			theList[i].Read = true
		}
	}

	return &LoadBottomArticlesResult{List: theList}
}
