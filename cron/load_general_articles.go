package cron

import (
	"fmt"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoadGeneralArticles(board *schema.BoardSummary) (err error) {
	nextIdx := ""
	count := 0

	for {
		articleSummaries, newNextIdx, err := loadGeneralArticles(board.BBoardID, nextIdx)
		if err != nil {
			logrus.Errorf("cron.LoadGeneralArticles: unable to loadGeneralArticles: nextIdx: %v e: %v", nextIdx, err)
			return err
		}

		for _, each := range articleSummaries {
			err = TryGetArticleContentInfo(each.BBoardID, each.ArticleID)
			if err == nil {
				count++
			}
		}

		if newNextIdx == "" {
			logrus.Infof("cron.LoadGeneralArticles: bid: %v load %v articles", board.BBoardID, count)
			return nil
		}

		nextIdx = newNextIdx
	}
}

func loadGeneralArticles(boardID bbs.BBoardID, startIdx string) (articleSummaries []*schema.ArticleSummaryWithRegex, nextIdx string, err error) {
	// backend load-general-articles
	theParams_b := &pttbbsapi.LoadGeneralArticlesParams{
		StartIdx:  startIdx,
		NArticles: N_ARTICLES,
		Desc:      true,
		IsSystem:  true,
	}
	var result_b *pttbbsapi.LoadGeneralArticlesResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_GENERAL_ARTICLES_R)

	c := &gin.Context{}
	statusCode, err := utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, "", err
	}
	if statusCode != 200 {
		return nil, "", fmt.Errorf("invalid statusCode: statusCode: %v", statusCode)
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	articleSummaries, err = api.DeserializeArticlesAndUpdateDB(result_b.Articles, updateNanoTS)
	if err != nil {
		return nil, "", err
	}

	return articleSummaries, result_b.NextIdx, nil
}
