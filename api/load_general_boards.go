package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const LOAD_GENERAL_BOARDS_R = "/boards"

type LoadGeneralBoardsParams struct {
	Title     string `json:"title,omitempty" form:"title,omitempty" url:"title,omitempty"`
	Keyword   string `json:"keyword,omitempty" form:"keyword,omitempty" url:"keyword,omitempty"`
	StartIdx  string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Ascending bool   `json:"asc,omitempty"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadGeneralBoardsResult struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`
}

func NewLoadGeneralBoardsParams() *LoadGeneralBoardsParams {
	return &LoadGeneralBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}

func LoadGeneralBoardsWrapper(c *gin.Context) {
	params := NewLoadGeneralBoardsParams()
	LoginRequiredQuery(LoadGeneralBoards, params, c)
}

func LoadGeneralBoards(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	return loadGeneralBoardsCore(remoteAddr, userID, params, c, pttbbsapi.LOAD_GENERAL_BOARDS_R)
}

func loadGeneralBoardsCore(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context, url string) (result interface{}, statusCode int, err error) {

	theParams, ok := params.(*LoadGeneralBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	//backend load-general-baords
	theParams_b := &pttbbsapi.LoadGeneralBoardsParams{
		StartIdx: theParams.StartIdx,
		Title:    types.Utf8ToBig5(theParams.Title),
		Keyword:  types.Utf8ToBig5(theParams.Keyword),
		NBoards:  theParams.Max,
		Asc:      theParams.Ascending,
	}
	var result_b *pttbbsapi.LoadGeneralBoardsResult

	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	return postLoadBoards(userID, result_b, url)
}

func postLoadBoards(userID bbs.UUserID, result_b *pttbbsapi.LoadGeneralBoardsResult, url string) (result *LoadGeneralBoardsResult, statusCode int, err error) {

	//update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	//check isRead
	userBoardInfoMap, err = checkUserReadBoard(userID, userBoardInfoMap, boardSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadGeneralBoardsResult(boardSummaries_db, userBoardInfoMap, result_b.NextIdx, url)

	return r, 200, nil

}

//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L953
func checkUserReadBoard(userID bbs.UUserID, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, theList []*schema.BoardSummary) (newUserBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, err error) {
	checkBBoardIDMap := make(map[bbs.BBoardID]int)
	queryBBoardIDs := make([]bbs.BBoardID, 0, len(theList))
	for idx, each := range theList {
		if each == nil {
			continue
		}

		eachBoardInfo, ok := userBoardInfoMap[each.BBoardID]
		if (eachBoardInfo.Stat&ptttype.NBRD_LINE != 0) || (eachBoardInfo.Stat&ptttype.NBRD_FOLDER != 0) {
			continue
		}

		if ok && eachBoardInfo.Read {
			continue
		}

		if each.BrdAttr&(ptttype.BRD_GROUPBOARD|ptttype.BRD_SYMBOLIC) != 0 {
			continue
		}

		if each.Total == 0 {
			continue
		}

		//check with read-time
		checkBBoardIDMap[each.BBoardID] = idx
		queryBBoardIDs = append(queryBBoardIDs, each.BBoardID)
	}

	dbResults, err := schema.FindUserReadBoards(userID, queryBBoardIDs)
	if err != nil {
		return nil, err
	}

	//setup read in the list
	//no need to update db, because we don't read the newest yet.
	//the Read flag is set based on the existing db.UpdateNanoTS
	for _, each := range dbResults {
		eachBoardID := each.BBoardID
		eachReadNanoTS := each.UpdateNanoTS

		eachBoardInfo, ok := userBoardInfoMap[eachBoardID]
		if !ok {
			continue
		}

		listIdx, ok := checkBBoardIDMap[eachBoardID]
		if !ok {
			continue
		}

		eachInTheList := theList[listIdx]
		eachLastPostNanoTS := eachInTheList.LastPostTime

		isRead := eachReadNanoTS > eachLastPostNanoTS
		eachBoardInfo.Read = isRead
	}

	return userBoardInfoMap, nil
}

func NewLoadGeneralBoardsResult(boardSummaries_db []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, nextIdx string, url string) *LoadGeneralBoardsResult {

	theList := make([]*apitypes.BoardSummary, len(boardSummaries_db))

	isByClass := url == pttbbsapi.LOAD_GENERAL_BOARDS_BY_CLASS_R
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

	return &LoadGeneralBoardsResult{
		List:    theList,
		NextIdx: nextIdx,
	}
}
