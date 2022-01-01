package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/sirupsen/logrus"
)

func RetryLoadFullClassBoards(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			logrus.Infof("RetryLoadFullClassBoards: to LoadFullClassBoards")
			_ = LoadFullClassBoards()
			logrus.Infof("RetryLoadFullClassBoards: to sleep 1 hr")
			time.Sleep(1 * time.Hour)
		}
	}
}

func LoadFullClassBoards() (err error) {
	nextBid := ptttype.Bid(1)

	count := 0
	for {
		boardSummaries, newNextBid, err := loadFullClassBoards(nextBid)
		if err != nil {
			logrus.Errorf("cron.LoadFullClassBoards: unable to loadFullClassBoards: nextBid: %v e: %v", nextBid, err)
			return err
		}
		count += len(boardSummaries)

		if newNextBid == 0 {
			logrus.Infof("cron.LoadFullClassBoards: load classes: %v", count)
			return nil
		}

		nextBid = newNextBid
	}
}

func loadFullClassBoards(startBid ptttype.Bid) (boardSummaries []*schema.BoardSummary, nextBid ptttype.Bid, err error) {
	// backend load-general-baords
	theParams_b := &pttbbsapi.LoadFullClassBoardsParams{
		StartBid: startBid,
		NBoards:  N_BOARDS,
		IsSystem: true,
	}
	var result_b *pttbbsapi.LoadFullClassBoardsResult

	statusCode, err := utils.BackendGet(nil, pttbbsapi.LOAD_FULL_CLASS_BOARDS_R, theParams_b, nil, &result_b)
	if err != nil {
		return nil, 0, err
	}
	if statusCode != 200 {
		return nil, 0, fmt.Errorf("invalid statusCode: statusCode: %v", statusCode)
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries, err = api.DeserializeBoardsAndUpdateDB(result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, 0, err
	}

	return boardSummaries, result_b.NextBid, nil
}
