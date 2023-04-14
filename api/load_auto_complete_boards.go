package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_AUTO_COMPLETE_BOARDS_R = "/boards/autocomplete"

type LoadAutoCompleteBoardsParams struct {
	Keyword   string `json:"brdname,omitempty" form:"brdname,omitempty" url:"brdname,omitempty"`
	StartIdx  string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Ascending bool   `json:"asc,omitempty"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

func NewLoadAutoCompleteBoardsParams() *LoadAutoCompleteBoardsParams {
	return &LoadAutoCompleteBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}

func LoadAutoCompleteBoardsWrapper(c *gin.Context) {
	params := NewLoadAutoCompleteBoardsParams()
	LoginRequiredQuery(LoadAutoCompleteBoards, params, c)
}

func LoadAutoCompleteBoards(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadAutoCompleteBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	// backend load-general-baords
	theParams_b := &pttbbsapi.LoadAutoCompleteBoardsParams{
		StartIdx: theParams.StartIdx,
		Keyword:  theParams.Keyword,
		NBoards:  theParams.Max,
		Asc:      theParams.Ascending,
	}
	var result_b *pttbbsapi.LoadGeneralBoardsResult

	url := pttbbsapi.LOAD_AUTO_COMPLETE_BOARDS_R
	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	return postLoadBoards(userID, result_b, url, c)
}
