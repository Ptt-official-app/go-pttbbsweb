package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_DETAIL_R = "/boards/:bid"

type GetBoardDetailParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" url:"fields,omitempty"`
}

func NewGetBoardDetailParams() *GetBoardDetailParams {
	return &GetBoardDetailParams{}
}

type GetBoardDetailPath struct {
	BBoardID bbs.BBoardID `uri:"bid"`
}

type GetBoardDetailResult struct {
	BBoardID bbs.BBoardID `json:"bid"`

	Brdname string `json:"brdname"`

	Title     string          `json:"title,omitempty"`
	BrdAttr   ptttype.BrdAttr `json:"flag,omitempty"`
	BoardType string          `json:"boardType,omitempty"`
	Category  string          `json:"cat,omitempty"`
	NUser     int             `json:"onlineCount,omitempty"`
	BMs       []string        `json:"moderators,omitempty"`

	VoteLimitLogins      int          `json:"voteLimitLogins,omitempty"`
	UpdateTimeTS         types.Time8  `json:"updateTime,omitempty"`
	PostLimitLogins      int          `json:"postLimitLogins,omitempty"`
	NVote                int          `json:"vote,omitempty"`
	VoteClosingTimeTS    types.Time8  `json:"vtime,omitempty"`
	Level                ptttype.PERM `json:"level,omitempty"`
	LastSetTimeTS        types.Time8  `json:"lastSetTime,omitempty"`
	PostExpire           int          `json:"postExpire,omitempty"`
	EndGambleTS          types.Time8  `json:"endGamble,omitempty"`
	PostType             string       `json:"postType,omitempty"`
	FastRecommendPauseTS types.Time8  `json:"fastRecommendPause,omitempty"`
	VoteLimitBadpost     int          `json:"voteLimitBadpost,omitempty"`
	PostLimitBadpost     int          `json:"postLimitBadpost,omitempty"`

	Read  bool `json:"read,omitempty"`
	Total int  `json:"total,omitempty"`
}

type GetBoardDetailFailResult struct {
	BBoard bbs.BBoardID `json:"boardID"`
	BMs    []string     `json:"moderators"`
	Reason string       `json:"reason"`
}

func GetBoardDetail(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	thePath, ok := path.(*GetBoardDetailPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	result = &GetBoardDetailResult{
		BBoardID: thePath.BBoardID,
	}

	return result, 200, nil
}
