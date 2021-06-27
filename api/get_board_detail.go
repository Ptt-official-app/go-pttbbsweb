package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_DETAIL_R = "/board/:bid"

type GetBoardDetailParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" url:"fields,omitempty"`
}

type GetBoardDetailPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type GetBoardDetailResult struct {
	*apitypes.BoardSummary

	UpdateTimeTS types.Time8 `json:"update_time"`

	VoteLimitLogins  int `json:"vote_limit_logins"`
	PostLimitLogins  int `json:"post_limit_logins"`
	VoteLimitBadpost int `json:"vote_limit_bad_post"`
	PostLimitBadpost int `json:"post_limit_bad_post"`

	NVote             int         `json:"vote"`
	VoteClosingTimeTS types.Time8 `json:"vtime"`

	Level                ptttype.PERM `json:"perm"`
	LastSetTimeTS        types.Time8  `json:"last_set_time"`
	PostExpire           int          `json:"post_expire"`
	EndGambleTS          types.Time8  `json:"end_gamble"`
	PostType             string       `json:"post_type"`
	FastRecommendPauseTS types.Time8  `json:"fast_recommend_pause"`
}

type GetBoardDetailFailResult struct {
	BBoard apitypes.FBoardID `json:"bid"`
	BMs    []bbs.UUserID     `json:"moderators"`
	Reason string            `json:"reason"`
}

func GetBoardDetailWrapper(c *gin.Context) {
	params := &GetBoardDetailParams{}
	path := &GetBoardDetailPath{}
	LoginRequiredPathQuery(GetBoardDetail, params, path, c)
}

func GetBoardDetail(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	_, ok := path.(*GetBoardDetailPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	result = &GetBoardDetailResult{
		BoardSummary: &apitypes.BoardSummary{
			FBoardID:  apitypes.FBoardID("WhoAmI"),
			Brdname:   "WhoAmI",
			Title:     "我～是～誰？～",
			BrdAttr:   0,
			BoardType: "◎",
			Category:  "嘰哩",
			NUser:     39,
			BMs:       []bbs.UUserID{"okcool", "teemo"},

			Read:  true,
			Total: 134,
		},

		UpdateTimeTS: 1234567890,

		VoteLimitLogins:  10,
		PostLimitLogins:  10,
		VoteLimitBadpost: 0,
		PostLimitBadpost: 0,

		NVote:                120,
		VoteClosingTimeTS:    1712300000,
		EndGambleTS:          1712300000,
		FastRecommendPauseTS: 10,
	}

	return result, 200, nil
}
