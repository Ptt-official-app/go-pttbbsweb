package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-pttbbsweb/api"
	"github.com/sirupsen/logrus"
)

func RetryLoadPopularBoards(ctx context.Context) error {
	time.Sleep(5 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadPopularBoards: to LoadPopularBoards")
			_ = LoadPopularBoards()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadPopularBoards: to sleep 10 min")
				time.Sleep(10 * time.Minute)
			}
		}
	}
}

func LoadPopularBoards() (err error) {
	boardSummaries, err := api.TryLoadPttWebPopularBoards(nil)
	if err != nil {
		logrus.Errorf("cron.LoadPopularBoards: unable to LoadPopularBoards: e: %v", err)
		return err
	}
	logrus.Infof("cron.LoadPopularBoards: boardSummaries: %v", len(boardSummaries))

	return nil
}
