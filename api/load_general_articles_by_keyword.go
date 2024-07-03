package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	ptttypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
)

func LoadGeneralArticlesByKeyword(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadGeneralArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	// is board-valid-user
	_, statusCode, err = isBoardValidUser(boardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	// get article-summaries
	keywordList := types.StringsSplitAsRune(theParams.Keyword, " ")

	var createNanoTS types.NanoTS
	var articleID bbs.ArticleID
	var createTSTime4 ptttypes.Time4
	if theParams.StartIdx != "" {
		createTSTime4, articleID, err = bbs.DeserializeArticleIdxStr(theParams.StartIdx)
		if err != nil {
			return nil, 400, err
		}
		createNanoTS = types.Time4ToNanoTS(createTSTime4)
	}

	articleSummaries_db, err := schema.GetArticleSummariesByRegex(boardID, keywordList, createNanoTS, articleID, theParams.Descending, theParams.Max+1)
	if err != nil {
		return nil, 500, err
	}

	// nextIdx
	var nextIdx string
	if len(articleSummaries_db) > theParams.Max {
		nextIdx = articleSummaries_db[theParams.Max].Idx
		articleSummaries_db = articleSummaries_db[:theParams.Max]
	}

	userReadArticleMap, err := getUserReadArticleMap(userID, boardID, articleSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	return NewLoadGeneralArticlesResultByKeyword(articleSummaries_db, userReadArticleMap, nextIdx, userID), 200, nil
}

func getUserReadArticleMap(userID bbs.UUserID, boardID bbs.BBoardID, theList []*schema.ArticleSummary) (userReadArticleMap map[bbs.ArticleID]types.NanoTS, err error) {
	queryArticleIDs := make([]bbs.ArticleID, len(theList))
	for idx, each := range theList {
		queryArticleIDs[idx] = each.ArticleID
	}

	dbResults, err := schema.FindUserReadArticles(userID, boardID, queryArticleIDs)
	if err != nil {
		return nil, err
	}

	userReadArticleMap = make(map[bbs.ArticleID]types.NanoTS)
	for _, each := range dbResults {
		userReadArticleMap[each.ArticleID] = each.UpdateNanoTS
	}

	return userReadArticleMap, nil
}

func NewLoadGeneralArticlesResultByKeyword(a_db []*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]types.NanoTS, nextIdx string, userID bbs.UUserID) (result *LoadGeneralArticlesResult) {
	theList := make([]*apitypes.ArticleSummary, len(a_db))
	for i, each_db := range a_db {
		theList[i] = apitypes.NewArticleSummary(each_db, "")
		readNanoTS, ok := userReadArticleMap[each_db.ArticleID]
		if !ok {
			continue
		}
		if readNanoTS > each_db.MTime {
			theList[i].Read = true
		} else if readNanoTS > each_db.CreateTime {
			theList[i].Read = true
		}
	}

	return &LoadGeneralArticlesResult{
		List:    theList,
		NextIdx: nextIdx,

		TokenUser: userID,
	}
}
