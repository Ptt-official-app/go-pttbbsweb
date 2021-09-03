package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/gin-gonic/gin"
)

const LOAD_USER_ARTICLES_R = "/user/:user_id/articles"

type LoadUserArticlesParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadUserArticlesPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type LoadUserArticlesResult struct {
	List    []*apitypes.ArticleSummary `json:"list"`
	NextIdx string                     `json:"next_idx"`
}

func NewUserArticlesParams() *LoadUserArticlesParams {
	return &LoadUserArticlesParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadUserArticlesWrapper(c *gin.Context) {
	params := NewUserArticlesParams()
	path := &LoadUserArticlesPath{}
	LoginRequiredPathQuery(LoadUserArticles, params, path, c)
}

func LoadUserArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadUserArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadUserArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	articleSummaries_db, nextIdx, err := loadUserArticles(thePath.UserID, theParams.StartIdx, theParams.Descending, theParams.Max)
	if err != nil {
		return nil, 500, err
	}

	userReadArticleMap, err := checkReadUserArticles(userID, articleSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadUserArticlesResult(articleSummaries_db, userReadArticleMap, nextIdx)
	return r, 200, nil
}

// loadUserArticles
func loadUserArticles(ownerID bbs.UUserID, startIdx string, descending bool, limit int) (articleSummaries_db []*schema.ArticleSummary, nextIdx string, err error) {
	articleSummaries_db = make([]*schema.ArticleSummary, 0, limit+1)

	var nextCreateTime4 pttbbstypes.Time4
	if startIdx != "" {
		nextCreateTime4, _, err = bbs.DeserializeArticleIdxStr(startIdx)
		if err != nil {
			return nil, "", err
		}
	}
	nextCreateTime := types.Time4ToNanoTS(nextCreateTime4)

	isEndLoop := false
	remaining := limit
	for !isEndLoop && remaining > 0 {
		eachArticleSummaries_db, err := schema.GetArticleSummariesByOwnerID(ownerID, nextCreateTime, descending, limit+1)
		if err != nil {
			return nil, "", err
		}

		// check is-last query
		if len(eachArticleSummaries_db) < limit+1 {
			isEndLoop = true
			nextCreateTime = 0
			nextIdx = ""
		} else {
			// setup next
			nextArticleSummary := eachArticleSummaries_db[len(eachArticleSummaries_db)-1]
			eachArticleSummaries_db = eachArticleSummaries_db[:len(eachArticleSummaries_db)-1]

			nextCreateTime = nextArticleSummary.CreateTime
			nextIdx = nextArticleSummary.Idx
		}
		if len(eachArticleSummaries_db) == 0 {
			break
		}

		// is-valid
		validArticleSummaries_db, err := isValidArticleSummaries(eachArticleSummaries_db)
		if err != nil {
			return nil, "", err
		}

		// append
		if len(validArticleSummaries_db) > remaining {
			nextArticleSummary := validArticleSummaries_db[remaining]
			validArticleSummaries_db = validArticleSummaries_db[:remaining]

			nextCreateTime = nextArticleSummary.CreateTime
			nextIdx = nextArticleSummary.Idx
		}

		articleSummaries_db = append(articleSummaries_db, validArticleSummaries_db...)
		remaining -= len(validArticleSummaries_db)
	}

	return articleSummaries_db, nextIdx, nil
}

// isValidArticleSummaries
// XXX TODO
func isValidArticleSummaries(articleSummaries_db []*schema.ArticleSummary) ([]*schema.ArticleSummary, error) {
	return articleSummaries_db, nil
}

func checkReadUserArticles(userID bbs.UUserID, theList []*schema.ArticleSummary) (userReadArticleMap map[bbs.ArticleID]bool, err error) {
	queryArticleIDs := make([]bbs.ArticleID, 0, len(theList))
	checkArticleIDMap := make(map[bbs.ArticleID]int)
	for idx, each := range theList {
		checkArticleIDMap[each.ArticleID] = idx
		queryArticleIDs = append(queryArticleIDs, each.ArticleID)
	}

	dbResults, err := schema.FindUserReadArticlesByArticleIDs(userID, queryArticleIDs)
	if err != nil {
		return nil, err
	}

	// setup read in the list
	// no need to update db, because we don't read the article yet.
	// the Read flag is set based on the existing db.UpdateNanoTS
	userReadArticleMap = make(map[bbs.ArticleID]bool)
	for _, each := range dbResults {
		eachArticleID := each.ArticleID
		eachReadNanoTS := each.UpdateNanoTS

		listIdx, ok := checkArticleIDMap[eachArticleID]
		if !ok {
			continue
		}

		eachInTheList := theList[listIdx]
		eachPostNanoTS := eachInTheList.CreateTime
		isRead := eachReadNanoTS > eachPostNanoTS
		userReadArticleMap[eachArticleID] = isRead
	}

	return userReadArticleMap, nil
}

func NewLoadUserArticlesResult(a_db []*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]bool, nextIdx string) *LoadUserArticlesResult {
	theList := make([]*apitypes.ArticleSummary, len(a_db))
	for i, each_db := range a_db {
		theList[i] = apitypes.NewArticleSummary(each_db)
		articleID := each_db.ArticleID
		isRead, ok := userReadArticleMap[articleID]
		if ok && isRead {
			theList[i].Read = true
		}
	}

	return &LoadUserArticlesResult{
		List:    theList,
		NextIdx: nextIdx,
	}
}
