package api

import (
	"strconv"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

func deserializeArticlesAndUpdateDB(userID bbs.UUserID, articleSummaries_b []*bbs.ArticleSummary, updateNanoTS types.NanoTS) (articleSummaries []*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]bool, err error) {
	if len(articleSummaries_b) == 0 {
		return nil, nil, nil
	}

	articleSummaries = make([]*schema.ArticleSummary, len(articleSummaries_b))
	for idx, each_b := range articleSummaries_b {
		articleSummaries[idx] = schema.NewArticleSummary(each_b, updateNanoTS)
	}

	err = schema.UpdateArticleSummaries(articleSummaries, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadArticles := make([]*schema.UserReadArticle, 0, len(articleSummaries_b))
	userReadArticleMap = make(map[bbs.ArticleID]bool)
	for _, each_b := range articleSummaries_b {
		if each_b.Read {
			each_db := &schema.UserReadArticle{
				UserID:       userID,
				ArticleID:    each_b.ArticleID,
				UpdateNanoTS: updateNanoTS,
			}
			userReadArticles = append(userReadArticles, each_db)

			userReadArticleMap[each_db.ArticleID] = true
		}
	}

	err = schema.UpdateUserReadArticles(userReadArticles, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	return articleSummaries, userReadArticleMap, err
}

func loadArticlesStartIdx(startIdxStr string, desc bool, theMax int) (newStartIdxStr string) {
	if desc == true {
		return startIdxStr
	}

	//ascending
	if startIdxStr == "" {
		startIdxStr = "1"
	}

	startIdx, err := ptttype.ToSortIdx(startIdxStr)
	if err != nil || startIdx < 1 { //guarantee startIdx >= 1
		startIdx = 1
	}

	startIdx += ptttype.SortIdx(theMax - 1) //startIdx is included.

	return strconv.Itoa(int(startIdx))
}

func reverseArticleSummaryList(s []*apitypes.ArticleSummary) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
