package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type userBoardInfo struct {
	Read  bool
	Stat  ptttype.BoardStatAttr
	NUser int
}

func deserializeBoardsAndUpdateDB(userID bbs.UUserID, boardSummaries_b []*bbs.BoardSummary, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*userBoardInfo, err error) {

	if len(boardSummaries_b) == 0 {
		return nil, nil, nil
	}

	boardSummaries = make([]*schema.BoardSummary, len(boardSummaries_b))
	for idx, each_b := range boardSummaries_b {
		boardSummaries[idx] = schema.NewBoardSummary(each_b, updateNanoTS)
	}

	err = schema.UpdateBoardSummaries(boardSummaries, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadBoards := make([]*schema.UserReadBoard, 0, len(boardSummaries_b))
	userBoardInfoMap = make(map[bbs.BBoardID]*userBoardInfo)
	for _, each_b := range boardSummaries_b {
		userBoardInfoMap[each_b.BBoardID] = &userBoardInfo{
			Read:  each_b.Read,
			Stat:  each_b.StatAttr,
			NUser: int(each_b.NUser),
		}

		if each_b.Read {
			each_db := &schema.UserReadBoard{
				UserID:       userID,
				BBoardID:     each_b.BBoardID,
				UpdateNanoTS: updateNanoTS,
			}
			userReadBoards = append(userReadBoards, each_db)
		}
	}

	err = schema.UpdateUserReadBoards(userReadBoards, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	return boardSummaries, userBoardInfoMap, err
}
