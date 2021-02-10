package api

import (
	"strconv"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const CREATE_BOARD_R = "/class/:cls/board"

type CreateBoardParams struct {
	Brdname      string            `json:"brdname" form:"brdname" url:"brdname"`
	BrdClass     string            `json:"class" form:"class" url:"class"`
	BrdTitle     string            `json:"title" form:"title" url:"title"`
	BMs          []bbs.UUserID     `json:"bms,omitempty" form:"bms,omitempty" url:"bms,omitempty"`
	BrdAttr      ptttype.BrdAttr   `json:"brdattr,omitempty" form:"brdattr,omitempty" url:"brdattr,omitempty"`
	Level        ptttype.PERM      `json:"level,omitempty" form:"level,omitempty" url:"level,omitempty"`
	ChessCountry ptttype.ChessCode `json:"chess_country,omitempty" form:"chess_country,omitempty" url:"chess_country,omitempty"`
	IsGroup      bool              `json:"is_group,omitempty" form:"is_group,omitempty" url:"is_group,omitempty"`
}

type CreateBoardPath struct {
	ClsBid ptttype.Bid `uri:"cls" binding:"required"`
}

type CreateBoardResult *apitypes.BoardSummary

func CreateBoardWrapper(c *gin.Context) {
	params := &CreateBoardParams{}
	path := &CreateBoardPath{}

	LoginRequiredPathJSON(CreateBoard, params, path, c)
}

func CreateBoard(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	theParams, ok := params.(*CreateBoardParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*CreateBoardPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	theClass := types.Utf8ToBig5(theParams.BrdClass)
	theTitle := types.Utf8ToBig5(theParams.BrdTitle)

	//backend
	theParams_b := &pttbbsapi.CreateBoardParams{
		Brdname:      theParams.Brdname,
		BrdClass:     theClass,
		BrdTitle:     theTitle,
		BMs:          theParams.BMs,
		BrdAttr:      theParams.BrdAttr,
		Level:        theParams.Level,
		ChessCountry: theParams.ChessCountry,
		IsGroup:      theParams.IsGroup,
	}
	var result_b pttbbsapi.CreateBoardResult

	urlMap := map[string]string{
		"cls": strconv.Itoa(int(thePath.ClsBid)),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CREATE_BOARD_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update to db
	theList_b := []*bbs.BoardSummary{(*bbs.BoardSummary)(result_b)}
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, _, err := deserializeBoardsAndUpdateDB(userID, theList_b, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	boardSummary_db := boardSummaries_db[0]
	boardSummary := apitypes.NewBoardSummary(boardSummary_db, "")

	//result
	result = CreateBoardResult(boardSummary)

	return result, 200, nil
}
