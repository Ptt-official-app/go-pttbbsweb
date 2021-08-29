package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const LOAD_CLASS_BOARDS_R = "/cls/:clsid"

type LoadClassBoardsParams struct {
	StartIdx string          `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	SortBy   ptttype.BSortBy `json:"sortby,omitempty" form:"sortby,omitempty" url:"sortby,omitempty"`

	Ascending bool `json:"asc,omitempty"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int  `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadClassBoardsPath struct {
	ClsID ptttype.Bid `uri:"clsid"`
}

type LoadClassBoardsResult struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`
}

func NewLoadClassBoardsParams() *LoadClassBoardsParams {
	return &LoadClassBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}

func LoadClassBoardsWrapper(c *gin.Context) {
	params := NewLoadClassBoardsParams()
	path := &LoadClassBoardsPath{}
	LoginRequiredPathQuery(LoadClassBoards, params, path, c)
}

func LoadClassBoards(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadClassBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadClassBoardsPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := bidToBoardID(thePath.ClsID)
	if err != nil {
		return nil, 500, err
	}

	// is board-valid-user
	_, statusCode, err = isBoardValidUser(boardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	boardSummaries_db, userBoardInfoMap, nextIdx, err := loadClassBoards(userID, thePath.ClsID, theParams.StartIdx, theParams.Ascending, theParams.Max, theParams.SortBy, c)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadClassBoardsResult(boardSummaries_db, userBoardInfoMap, nextIdx, theParams.SortBy)

	return r, 200, nil
}

func loadClassBoards(userID bbs.UUserID, clsID ptttype.Bid, startIdx string, asc bool, limit int, sortBy ptttype.BSortBy, c *gin.Context) (boardSummaries_db []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, nextIdx string, err error) {
	boardSummaries_db = make([]*schema.BoardSummary, 0, limit)
	userBoardInfoMap = make(map[bbs.BBoardID]*apitypes.UserBoardInfo)

	nextIdx = startIdx

	isEndLoop := false
	remaining := limit
	for remaining > 0 && !isEndLoop {
		eachBoardBids, err := schema.GetBoardBidsByClsID(clsID, nextIdx, asc, limit+1, sortBy)
		if err != nil {
			break
		}
		if len(eachBoardBids) < limit+1 {
			isEndLoop = true
			nextIdx = ""
		} else {
			nextBoardBid := eachBoardBids[len(eachBoardBids)-1]
			eachBoardBids = eachBoardBids[:len(eachBoardBids)-1]

			switch sortBy {
			case ptttype.BSORT_BY_NAME:
				nextIdx = nextBoardBid.IdxByName
			case ptttype.BSORT_BY_CLASS:
				nextIdx = nextBoardBid.IdxByClass
			}
		}
		if len(eachBoardBids) == 0 {
			break
		}

		// is-valid
		eachBids := make([]ptttype.Bid, len(eachBoardBids))
		for idx, each := range eachBoardBids {
			eachBids[idx] = each.Bid
		}

		boardSummaryMap_db, eachUserBoardInfoMap, _, err := getBoardSummaryMapFromBids(userID, eachBids, c)
		if err != nil {
			return nil, nil, "", err
		}

		validBoardSummaries_db := make([]*schema.BoardSummary, 0, len(boardSummaryMap_db))
		for _, eachBid := range eachBids {
			eachBoardSummary, ok := boardSummaryMap_db[eachBid]
			if !ok {
				continue
			}
			validBoardSummaries_db = append(validBoardSummaries_db, eachBoardSummary)
		}

		// append
		if len(validBoardSummaries_db) > remaining {
			nextBoardSummary := validBoardSummaries_db[remaining]
			validBoardSummaries_db = validBoardSummaries_db[:remaining]

			switch sortBy {
			case ptttype.BSORT_BY_NAME:
				nextIdx = nextBoardSummary.IdxByName
			case ptttype.BSORT_BY_CLASS:
				nextIdx = nextBoardSummary.IdxByClass
			}
		}

		boardSummaries_db = append(boardSummaries_db, validBoardSummaries_db...)
		remaining -= len(validBoardSummaries_db)

		for eachBoardID, eachBoardInfo := range eachUserBoardInfoMap {
			userBoardInfoMap[eachBoardID] = eachBoardInfo
		}
	}

	// check isRead
	userBoardInfoMap, err = checkUserReadBoard(userID, userBoardInfoMap, boardSummaries_db)
	if err != nil {
		return nil, nil, "", err
	}

	return boardSummaries_db, userBoardInfoMap, nextIdx, nil
}

func NewLoadClassBoardsResult(boardSummaries_db []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, nextIdx string, sortBy ptttype.BSortBy) (ret *LoadClassBoardsResult) {
	theList := make([]*apitypes.BoardSummary, len(boardSummaries_db))

	isByClass := sortBy == ptttype.BSORT_BY_CLASS
	for i, each_db := range boardSummaries_db {
		idx := each_db.IdxByName
		if isByClass {
			idx = each_db.IdxByClass
		}

		userBoardInfo, ok := userBoardInfoMap[each_db.BBoardID]
		if !ok {
			continue
		}

		each := apitypes.NewBoardSummary(each_db, idx, userBoardInfo)
		theList[i] = each
	}

	return &LoadClassBoardsResult{
		List:    theList,
		NextIdx: nextIdx,
	}
}
