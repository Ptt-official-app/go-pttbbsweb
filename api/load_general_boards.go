package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const LOAD_GENERAL_BOARDS_R = "/Board/search"

type LoadGeneralBoardsParams struct {
	Keyword  string `json:"name,omitempty"`
	StartIdx string `json:"startIdx,omitempty"`
	Max      int    `json:"max,omitempty"`
}

func NewLoadGeneralBoardsParams() *LoadGeneralBoardsParams {
	return &LoadGeneralBoardsParams{
		Max: DEFAULT_MAX_LIST,
	}
}

type LoadGeneralBoardsResult struct {
	List    []*types.BoardSummary `json:"list"`
	NextIdx string                `json:"nextIdx"`
}

func LoadGeneralBoards(remoteAddr string, userID string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}
	log.Infof("LoadGeneralBoards: theParams: %v", theParams)

	//backend load-general-baords
	theParams_b := &backend.LoadGeneralBoardsParams{
		StartIdx: theParams.StartIdx,
		Keyword:  theParams.Keyword,
		NBoards:  theParams.Max,
	}
	var result_b *backend.LoadGeneralBoardsResult

	url := backend.WithPrefix(backend.LOAD_GENERAL_BOARDS_R)
	statusCode, err = utils.HttpPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	r := &LoadGeneralBoardsResult{}
	r.Deserialize(result_b)

	//check isRead
	err = checkNewPost(userID, r.List)
	if err != nil {
		return nil, 500, err
	}

	return r, 200, nil
}

func (r *LoadGeneralBoardsResult) Deserialize(r_b *backend.LoadGeneralBoardsResult) {

	r.List = make([]*types.BoardSummary, len(r_b.Boards))
	for i := 0; i < len(r.List); i++ {
		each := &types.BoardSummary{}
		each.Deserialize(r_b.Boards[i])
		r.List[i] = each
	}

	r.NextIdx = r_b.NextIdx
}

//https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L953
func checkNewPost(userID string, theList []*types.BoardSummary) error {
	checkBoardIDMap := make(map[string]int)
	queryBoardIDs := make([]string, 0, len(theList))
	for idx, each := range theList {
		if each.Read {
			continue
		}
		if each.BrdAttr&(ptttype.BRD_GROUPBOARD|ptttype.BRD_SYMBOLIC) != 0 {
			continue
		}
		if each.StatAttr_d&(ptttype.NBRD_LINE|ptttype.NBRD_FOLDER) != 0 {
			continue
		}
		if each.Total == 0 {
			continue
		}

		//check with read-time
		checkBoardIDMap[each.BoardID] = idx
		queryBoardIDs = append(queryBoardIDs, each.BoardID)
	}

	//query
	query := make(map[string]interface{})
	query[schema.USER_READ_BOARD_USER_ID_b] = userID
	queryBoards := make(map[string]interface{})
	queryBoards["$in"] = queryBoardIDs
	query[schema.USER_READ_BOARD_BOARD_ID_b] = queryBoards

	var dbResults []*schema.UserReadBoard
	err := schema.UserReadBoard_c.Find(query, 0, &dbResults, nil)
	if err != nil {
		return err
	}

	for _, each := range dbResults {
		eachBoardID := each.BoardID
		eachReadNanoTS := each.UpdateNanoTS

		listIdx := checkBoardIDMap[eachBoardID]
		eachInTheList := theList[listIdx]
		eachLastPostNanoTS := utils.TSToNanoTS(int64(eachInTheList.LastPostTime_d))
		eachInTheList.Read = eachReadNanoTS > eachLastPostNanoTS
	}

	return nil
}
