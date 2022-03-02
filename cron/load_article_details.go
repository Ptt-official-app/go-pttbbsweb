package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/sirupsen/logrus"
)

func RetryLoadArticleDetails(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadGeneralArticles: to LoadGeneralArticles")
			_ = LoadArticleDetails()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadGeneralArticles: to sleep 1 min")
				time.Sleep(1 * time.Minute)
			}
		}
	}
}

func LoadArticleDetails() (err error) {
	nextBrdname := ""
	count := 0

	for {
		boardIDs, err := schema.GetBoardIDs(nextBrdname, false, N_BOARDS+1, true)
		if err != nil {
			return err
		}

		newNextBrdname := ""
		if len(boardIDs) == N_BOARDS+1 {
			newNextBoardID := boardIDs[N_BOARDS]
			newNextBrdname = newNextBoardID.Brdname
			boardIDs = boardIDs[:N_BOARDS]
		}

		for _, each := range boardIDs {
			err = loadArticleDetails(each.BBoardID)
			if err == nil {
				count++
			}
		}

		if newNextBrdname == "" {
			logrus.Infof("cron.LoadArticleDetails: load %v boards", count)
			return nil

		}

		nextBrdname = newNextBrdname
	}
}

func loadArticleDetails(boardID bbs.BBoardID) (err error) {
	nextIdx := ""
	count := 0
	for {
		articleDetailSummaries, err := schema.GetArticleDetailSummaries(boardID, nextIdx, false, N_ARTICLE_DETAILS+1, true)
		if err != nil {
			return err
		}

		newNextIdx := ""
		if len(articleDetailSummaries) == N_ARTICLE_DETAILS+1 {
			nextDetailSummary := articleDetailSummaries[N_ARTICLE_DETAILS]
			newNextIdx = nextDetailSummary.Idx

			articleDetailSummaries = articleDetailSummaries[:N_ARTICLE_DETAILS]
		}

		origCount := count
		for _, each := range articleDetailSummaries {
			if each.MTime < each.ContentMTime {
				continue
			}

			_, _, _, _, _, _, _, _, _, _, _, _, err = api.TryGetArticleContentInfo("SYSOP", each.BBoardID, each.ArticleID, nil, true, false, false)
			if err == nil {
				count++
			}
		}

		if origCount != count {
			logrus.Infof("cron.loadArticleDetails: bid: %v count: %v", boardID, count)
		}

		if newNextIdx == "" {
			logrus.Infof("cron.loadArticleDetails: bid: %v load %v articles", boardID, count)
			return nil
		}

		nextIdx = newNextIdx
	}
}
