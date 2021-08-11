package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
)

func toBoardID(fboardID apitypes.FBoardID, remoteAddr string, userID bbs.UUserID, c *gin.Context) (boardID bbs.BBoardID, err error) {
	boardID, err = fboardID.ToBBoardID()
	if err == nil {
		return boardID, nil
	}

	params := &LoadAutoCompleteBoardsParams{
		Keyword:   string(fboardID),
		Ascending: true,
		Max:       1,
	}

	_, _, err = LoadAutoCompleteBoards(remoteAddr, userID, params, c)
	if err != nil {
		return "", err
	}

	return fboardID.ToBBoardID()
}

// DeserializeBoards
//
// each_b.Reason happens only with invalid permission.
func DeserializeBoardsAndUpdateDB(boardSummaries_b []*bbs.BoardSummary, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, err error) {
	boardSummaries = make([]*schema.BoardSummary, 0, len(boardSummaries_b))
	for _, each_b := range boardSummaries_b {
		if each_b.Reason != 0 {
			continue
		}
		each := schema.NewBoardSummary(each_b, updateNanoTS)

		boardSummaries = append(boardSummaries, each)
	}
	if len(boardSummaries) == 0 {
		return nil, nil
	}

	err = schema.UpdateBoardSummaries(boardSummaries, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return boardSummaries, nil
}

func deserializeBoardsAndUpdateDB(userID bbs.UUserID, boardSummaries_b []*bbs.BoardSummary, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, err error) {
	if len(boardSummaries_b) == 0 {
		return nil, nil, nil
	}

	boardSummaries, err = DeserializeBoardsAndUpdateDB(boardSummaries_b, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadBoards := make([]*schema.UserReadBoard, 0, len(boardSummaries_b))
	userBoardInfoMap = make(map[bbs.BBoardID]*apitypes.UserBoardInfo)
	for _, each_b := range boardSummaries_b {
		if each_b.Reason != 0 {
			continue
		}

		userBoardInfoMap[each_b.BBoardID] = &apitypes.UserBoardInfo{
			Read: each_b.Read,
			Stat: each_b.StatAttr,
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

func isBoardValidUser(boardID bbs.BBoardID, c *gin.Context) (isValid bool, statusCode int, err error) {
	var result_b *pttbbsapi.IsBoardValidUserResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.IS_BOARD_VALID_USER_R)
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return false, statusCode, err
	}
	if !result_b.IsValid {
		return false, 403, ErrInvalidUser
	}

	return true, 200, nil
}
