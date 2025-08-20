package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/sirupsen/logrus"
)

func RetryLoadArticleDetails(ctx context.Context) error {
	time.Sleep(20 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadArticleDetails: to LoadArticleDetails")
			_ = LoadArticleDetails()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadArticleDetails: to sleep 10 min")
				time.Sleep(10 * time.Minute)
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

		for _, each := range articleDetailSummaries {
			if each.MTime <= each.ContentMTime && each.MTime < each.ContentUpdateNanoTS {
				continue
			}

			_, _, _, _, _, _, _, _, _, _, _, _, err = api.TryGetArticleContentInfo("SYSOP", each.BBoardID, each.ArticleID, nil, true, false, false)

			if err == nil {
				count++
			}
		}

		if newNextIdx == "" {
			return nil
		}

		nextIdx = newNextIdx
	}
}
