package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type BoardDetail struct {
	FBoardID     FBoardID        `json:"bid"`
	Brdname      string          `json:"brdname"`
	Title        string          `json:"title"`
	BrdAttr      ptttype.BrdAttr `json:"flag"`
	BoardType    string          `json:"type"`
	Category     string          `json:"class"`
	NUser        int             `json:"nuser"`
	BMs          []bbs.UUserID   `json:"moderators"`
	Reason       string          `json:"reason"`
	Read         bool            `json:"read"`
	Total        int             `json:"total"`
	LastPostTime types.Time8     `json:"last_post_time"`

	UpdateTime types.Time8 `json:"update_time"`

	VoteLimitLogins  int `json:"vote_limit_logins"`
	PostLimitLogins  int `json:"post_limit_logins"`
	VoteLimitBadpost int `json:"vote_limit_bad_post"`
	PostLimitBadpost int `json:"post_limit_bad_post"`

	NVote           int         `json:"vote"`
	VoteClosingTime types.Time8 `json:"vtime"`

	Level              ptttype.PERM `json:"level"`
	LastSetTime        types.Time8  `json:"last_set_time"` /* perm-reload */
	PostExpire         ptttype.Bid  `json:"link_pttbid"`   /* 看板連結的 bid */
	PostType           []string     `json:"post_type"`
	PostTemplate       []bool       `json:"post_tmpl"`
	EndGambleNanoTS    types.Time8  `json:"end_gamble"`
	FastRecommendPause types.Time8  `json:"fast_recommend_pause"`

	ChessCountry ptttype.ChessCode `json:"chesscountry"`

	LevelIdx schema.LevelIdx `json:"level_idx,omitempty"`

	Gid ptttype.Bid `json:"gid"`
	Bid ptttype.Bid `json:"pttbid"`

	Idx string `json:"idx"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

var BOARD_DETAIL_FIELD_MAP = map[string]string{
	"type":                 schema.BOARD_BOARD_TYPE_b,
	"moderators":           schema.BOARD_BMS_b,
	"last_post_time":       schema.BOARD_LAST_POST_TIME_b,
	"update_time":          schema.BOARD_UPDATE_NANO_TS_b,
	"vtime":                schema.BOARD_VOTE_CLOSING_TIME_b,
	"last_set_time":        schema.BOARD_LAST_SET_TIME_b,
	"link_pttbid":          schema.BOARD_POST_EXPIRE_b,
	"end_gamble":           schema.BOARD_END_GAMBLE_NANO_TS_b,
	"fast_recommend_pause": schema.BOARD_FAST_RECOMMEND_PAUSE_b,
	"gid":                  schema.BOARD_GID_b,
}

func NewBoardDetail(b_db *schema.BoardDetail, idx string, userID bbs.UUserID) *BoardDetail {
	if b_db == nil {
		return nil
	}

	return &BoardDetail{
		FBoardID:     ToFBoardID(b_db.BBoardID),
		Brdname:      b_db.Brdname,
		Title:        b_db.Title,
		BrdAttr:      b_db.BrdAttr,
		BoardType:    b_db.BoardType,
		Category:     b_db.Category,
		BMs:          b_db.BMs,
		Total:        b_db.Total,
		LastPostTime: b_db.LastPostTime.ToTime8(),
		NUser:        b_db.NUser,

		UpdateTime: b_db.UpdateTime.ToTime8(),

		VoteLimitLogins:  b_db.VoteLimitLogins,
		PostLimitLogins:  b_db.PostLimitLogins,
		VoteLimitBadpost: b_db.VoteLimitBadpost,
		PostLimitBadpost: b_db.PostLimitBadpost,

		NVote:           b_db.NVote,
		VoteClosingTime: b_db.VoteClosingTime.ToTime8(),

		Level:              b_db.Level,
		LastSetTime:        b_db.LastSetTime.ToTime8(),
		PostExpire:         b_db.PostExpire,
		PostType:           b_db.PostType,
		PostTemplate:       b_db.PostTemplate,
		EndGambleNanoTS:    b_db.EndGambleNanoTS.ToTime8(),
		FastRecommendPause: b_db.FastRecommendPause.ToTime8(),

		ChessCountry: b_db.ChessCountry,

		Gid: b_db.Gid,
		Bid: b_db.Bid,

		Idx: idx,

		TokenUser: userID,
	}
}
