package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_DETAIL_R = "/Board/Boards/:board"

type GetBoardDetailParams struct {
	Fields string `json:"fields,omitempty"`
}

func NewGetBoardDetailParams() *GetBoardDetailParams {
	return &GetBoardDetailParams{}
}

type GetBoardDetailPath struct {
	BoardID string `json:"board"`
}

type GetBoardDetailResult struct {
	BoardID string `json:"boardID"`

	Title     string          `json:"title"`
	BrdAttr   ptttype.BrdAttr `json:"flag"`
	BoardType string          `json:"boardType"`
	Category  string          `json:"cat"`
	NUser     int             `json:"onlineCount"`
	BMs       []string        `json:"moderators"`

	VoteLimitLogins      int          `json:"voteLimitLogins"`
	UpdateTimeTS         types.Time8  `json:"updateTime"`
	PostLimitLogins      int          `json:"postLimitLogins"`
	NVote                int          `json:"vote"`
	VoteClosingTimeTS    types.Time8  `json:"vtime"`
	Level                ptttype.PERM `json:"level"`
	LastSetTimeTS        types.Time8  `json:"lastSetTime"`
	PostExpire           int          `json:"postExpire"`
	EndGambleTS          types.Time8  `json:"endGamble"`
	PostType             string       `json:"postType"`
	FastRecommendPauseTS types.Time8  `json:"fastRecommendPause"`
	VoteLimitBadpost     int          `json:"voteLimitBadpost"`
	PostLimitBadpost     int          `json:"postLimitBadpost"`

	Read  bool `json:"read"`
	Total int  `json:"total"`
}

type GetBoardDetailFailResult struct {
	BoardID string   `json:"boardID"`
	BMs     []string `json:"moderators"`
	Reason  string   `json:"reason"`
}

func GetBoardDetail(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	thePath, ok := path.(*GetBoardDetailPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	result = &GetBoardDetailResult{
		BoardID: thePath.BoardID,
	}

	return result, 200, nil
}
