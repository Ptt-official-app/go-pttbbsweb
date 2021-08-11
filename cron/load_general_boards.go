package cron

import (
	"fmt"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RetryLoadGeneralBoards() {
	for {
		logrus.Infof("RetryLoadGeneralBoards: to LoadGeneralBoards")
		_ = LoadGeneralBoards()
		logrus.Infof("RetryLoadGeneralBoards: to sleep 1 min")
		time.Sleep(1 * time.Minute)
	}
}

func LoadGeneralBoards() (err error) {
	nextIdx := ""
	count := 0
	for {
		boardSummaries, newNextIdx, err := loadGeneralBoards(nextIdx)
		if err != nil {
			logrus.Errorf("cron.LoadGeneralBoards: unable to loadGeneralBoards: nextIdx: %v e: %v", nextIdx, err)
			return err
		}

		for idx, each := range boardSummaries {
			logrus.Infof("cron.LoadGeneralBoards: (%v/%v) to LoadGeneralArticles: %v", idx, len(boardSummaries), each.BBoardID)
			err = LoadGeneralArticles(each)
			if err == nil {
				count++
			}
		}

		if newNextIdx == "" {
			logrus.Infof("cron.LoadGeneralBoards: load %v boards", count)
			return nil
		}

		nextIdx = newNextIdx
	}
}

func loadGeneralBoards(startIdx string) (boardSummaries []*schema.BoardSummary, nextIdx string, err error) {
	// backend load-general-baords
	theParams_b := &pttbbsapi.LoadGeneralBoardsParams{
		StartIdx: startIdx,
		NBoards:  N_BOARDS,
		Asc:      true,
		IsSystem: true,
	}
	var result_b *pttbbsapi.LoadGeneralBoardsResult

	c := &gin.Context{}
	statusCode, err := utils.BackendGet(c, pttbbsapi.LOAD_GENERAL_BOARDS_R, theParams_b, nil, &result_b)
	if err != nil {
		return nil, "", err
	}
	if statusCode != 200 {
		return nil, "", fmt.Errorf("invalid statusCode: statusCode: %v", statusCode)
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries, err = api.DeserializeBoardsAndUpdateDB(result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, "", err
	}

	return boardSummaries, result_b.NextIdx, nil
}
