package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
)

const LOAD_GENERAL_ARTICLES_R = "/board/:bid/articles"

type LoadGeneralArticlesParams struct {
	Keyword    string `json:"title,omitempty" form:"title,omitempty" url:"title,omitempty"`
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
}

type LoadGeneralArticlesPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type LoadGeneralArticlesResult struct {
	List    []*apitypes.ArticleSummary `json:"list"`
	NextIdx string                     `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewLoadGeneralArticlesParams() *LoadGeneralArticlesParams {
	return &LoadGeneralArticlesParams{
		Max:        DEFAULT_MAX_LIST,
		Descending: DEFAULT_DESCENDING,
	}
}

func LoadGeneralArticlesWrapper(c *gin.Context) {
	params := NewLoadGeneralArticlesParams()
	path := &LoadGeneralArticlesPath{}
	LoginRequiredPathQuery(LoadGeneralArticles, params, path, c)
}

func LoadGeneralArticles(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadGeneralArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	if theParams.Keyword != "" {
		return LoadGeneralArticlesByKeyword(remoteAddr, userID, params, path, c)
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	// backend load-general-articles
	articleSummaries_db, err := schema.GetArticleSummaries(boardID, theParams.StartIdx, theParams.Descending, theParams.Max+1)
	if err != nil {
		return nil, 500, err
	}

	nextIdx := ""
	if len(articleSummaries_db) == theParams.Max+1 {
		nextArticleSummary := articleSummaries_db[theParams.Max]
		nextIdx = nextArticleSummary.Idx

		articleSummaries_db = articleSummaries_db[:theParams.Max]
	}

	userReadArticleMap := make(map[bbs.ArticleID]bool)
	userReadArticleMap, err = checkReadArticles(userID, boardID, userReadArticleMap, articleSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadGeneralArticlesResult(articleSummaries_db, userReadArticleMap, nextIdx, userID)

	// update user_read_board if is-newest
	if theParams.Descending && theParams.StartIdx == "" || !theParams.Descending && nextIdx == "" {

		updateNanoTS := types.NowNanoTS()
		err = updateUserReadBoard(userID, boardID, updateNanoTS)
		if err != nil {
			return nil, 500, err
		}
	}

	return r, 200, nil
}

func checkReadArticles(userID bbs.UUserID, boardID bbs.BBoardID, userReadArticleMap map[bbs.ArticleID]bool, theList []*schema.ArticleSummary) (newUserReadArticleMap map[bbs.ArticleID]bool, err error) {
	queryArticleIDs := make([]bbs.ArticleID, 0, len(theList))
	checkArticleIDMap := make(map[bbs.ArticleID]int)
	for idx, each := range theList {
		isRead, ok := userReadArticleMap[each.ArticleID]
		if ok && isRead {
			continue
		}

		// check with read-time
		checkArticleIDMap[each.ArticleID] = idx
		queryArticleIDs = append(queryArticleIDs, each.ArticleID)
	}

	dbResults, err := schema.FindUserReadArticles(userID, boardID, queryArticleIDs)
	if err != nil {
		return nil, err
	}

	// setup read in the list
	// no need to update db, because we don't read the article yet.
	// the Read flag is set based on the existing db.UpdateNanoTS
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

func updateUserReadBoard(userID bbs.UUserID, boardID bbs.BBoardID, updateNanoTS types.NanoTS) (err error) {
	userReadBoard := &schema.UserReadBoard{UserID: userID, BBoardID: boardID, UpdateNanoTS: updateNanoTS}

	err = schema.UpdateUserReadBoard(userReadBoard)
	if err != nil {
		return err
	}

	return nil
}

func NewLoadGeneralArticlesResult(a_db []*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]bool, nextIdx string, userID bbs.UUserID) *LoadGeneralArticlesResult {
	theList := make([]*apitypes.ArticleSummary, len(a_db))
	for i, each_db := range a_db {
		theList[i] = apitypes.NewArticleSummary(each_db, "")
		articleID := each_db.ArticleID
		isRead, ok := userReadArticleMap[articleID]
		if ok && isRead {
			theList[i].Read = true
		}
	}

	return &LoadGeneralArticlesResult{
		List:      theList,
		NextIdx:   nextIdx,
		TokenUser: userID,
	}
}
