package api

import (
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_DETAIL_R = "/board/:bid"

type GetBoardDetailParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" uri:"fields,omitempty"`
}

type GetBoardDetailPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type GetBoardDetailResult *apitypes.BoardDetail

func GetBoardDetailWrapper(c *gin.Context) {
	params := &GetBoardDetailParams{}
	path := &GetBoardDetailPath{}
	LoginRequiredPathQuery(GetBoardDetail, params, path, c)
}

func GetBoardDetail(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*GetBoardDetailParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*GetBoardDetailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 400, err
	}

	// is board-valid-user
	_, statusCode, err = isBoardValidUser(boardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	// fields
	var fieldMap map[string]bool
	if len(theParams.Fields) > 0 {
		fields := strings.Split(theParams.Fields, ",")
		fieldMap = make(map[string]bool)
		for _, each := range fields {
			each_db, ok := apitypes.BOARD_DETAIL_FIELD_MAP[each]
			if !ok {
				each_db = each
			}
			fieldMap[each_db] = true
		}
		fieldMap[schema.BOARD_BBOARD_ID_b] = true
	}

	boardDetail_db, err := schema.GetBoardDetail(boardID, fieldMap)
	if err != nil {
		return nil, 500, err
	}

	boardDetail := apitypes.NewBoardDetail(boardDetail_db, "")

	result = GetBoardDetailResult(boardDetail)

	return result, 200, nil
}
