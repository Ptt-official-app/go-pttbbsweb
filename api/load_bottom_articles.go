package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_BOTTOM_ARTICLES_R = "/board/:bid/articles/bottom"

type LoadBottomArticlesPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type LoadBottomArticlesResult struct {
	List []*apitypes.ArticleSummary `json:"list"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
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

	articleSummaries_db, err := schema.GetBottomArticleSummaries(boardID)
	if err != nil {
		return nil, 500, err
	}

	userReadArticleMap := make(map[bbs.ArticleID]bool)
	userReadArticleMap, err = checkReadArticles(userID, boardID, userReadArticleMap, articleSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadBottomArticlesResult(articleSummaries_db, userReadArticleMap, userID)

	return r, 200, nil
}

func NewLoadBottomArticlesResult(a_db []*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]bool, userID bbs.UUserID) *LoadBottomArticlesResult {
	theList := make([]*apitypes.ArticleSummary, len(a_db))
	for i, each_db := range a_db {
		theList[i] = apitypes.NewArticleSummary(each_db, "")
		articleID := each_db.ArticleID
		isRead, ok := userReadArticleMap[articleID]
		if ok && isRead {
			theList[i].Read = true
		}
	}

	return &LoadBottomArticlesResult{List: theList, TokenUser: userID}
}
