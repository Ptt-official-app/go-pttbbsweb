package cron

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/api"
	"github.com/Ptt-official-app/go-pttbbsweb/mand"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/sirupsen/logrus"
)

func RetryLoadManArticles(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("RetryLoadManArticles: Recovered r: %v stack: %v", r, string(debug.Stack()))
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadManArticles: to LoadManArticles")
			_ = LoadManArticles()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadManArticles: to sleep 1 min")
				time.Sleep(1 * time.Minute)
			}
		}
	}
}

func LoadManArticles() (err error) {
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
			err = loadManArticles(each.BBoardID)
			if err == nil {
				count++
			}
		}

		if newNextBrdname == "" {
			logrus.Infof("cron.LoadManArticles: load %v boards", count)
			return nil

		}

		nextBrdname = newNextBrdname
	}
}

func loadManArticles(boardID bbs.BBoardID) (err error) {
	count, err := loadManArticlesCore(boardID, "")
	if err != nil {
		logrus.Errorf("cron.loadManArticles: unable to loadManArticles: e: %v", err)
		return err
	}

	logrus.Infof("cron.loadManArticles: bid: %v count: %v", boardID, count)

	return nil
}

func loadManArticlesCore(boardID bbs.BBoardID, levelIdx types.ManArticleID) (count int, err error) {
	// backend load-general-articles
	brdname := boardID.ToBrdname()
	ctx := context.Background()
	req := &mand.ListRequest{
		BoardName: brdname,
		Path:      string(levelIdx),
	}
	resp, err := mand.Cli.List(ctx, req)
	if err != nil {
		return 0, err
	}

	entries := resp.Entries

	// update to db
	updateNanoTS := types.NowNanoTS()
	articleSummaries, err := api.DeserializePBManArticlesAndUpdateDB(boardID, levelIdx, entries, updateNanoTS)
	if err != nil {
		return 0, err
	}

	count += len(articleSummaries)

	// recursively loop through dirs.
	for _, each := range articleSummaries {
		if !each.IsDir {
			continue
		}
		eachCount, err := loadManArticlesCore(boardID, each.ArticleID)
		if err != nil {
			logrus.Errorf("cron.loadManArticlesCore: unable to loadManArticlesCore: boardID: %v path: %v each: %v e: %v", boardID, levelIdx, each.ArticleID, err)
		}

		count += eachCount
	}

	return count, nil
}
