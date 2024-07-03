package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var Board_c *db.Collection

type Board struct {
	Version   int             `bson:"version"`
	BBoardID  bbs.BBoardID    `bson:"bid"`
	Brdname   string          `bson:"brdname"`
	Title     string          `bson:"title"`
	BrdAttr   ptttype.BrdAttr `bson:"flag"`
	BoardType string          `bson:"the_type"`
	Category  string          `bson:"class"`
	NUser     int             `bson:"nuser"` /* use aggregate to periodically get the data */
	BMs       []bbs.UUserID   `bson:"bms"`
	Total     int             `bson:"total"` /* total articles, 需要即時知道. 因為 read 頻率高. 並且跟 last-post-time-ts 一樣 write 頻率 << read 頻率 */

	LastPostTime types.NanoTS `bson:"last_post_time_nano_ts"` /* 需要即時知道來做板的已讀 */

	UpdateTime types.NanoTS `bson:"update_time_nano_ts"` /* show 進板畫面, 目前只有 INT_MAX - 1 或 0 */

	VoteLimitLogins  int `bson:"vote_limit_logins"`
	PostLimitLogins  int `bson:"post_limit_logins"`
	VoteLimitBadpost int `bson:"vote_limit_bad_post"`
	PostLimitBadpost int `bson:"post_limit_bad_post"`

	Parent bbs.BBoardID `bson:"parent"`

	NVote           int          `bson:"vote"` /* use db-count to get current #vote */
	VoteClosingTime types.NanoTS `bson:"vtime_nano_ts"`

	Level              ptttype.PERM `bson:"perm"`
	LastSetTime        types.NanoTS `bson:"last_set_time_nano_ts"` /* perm-reload */
	PostExpire         ptttype.Bid  `bson:"post_expire"`           /* 看板連結的 bid */
	PostType           []string     `bson:"post_type"`
	PostTemplate       []bool       `bson:"post_tmpl"`
	EndGambleNanoTS    types.NanoTS `bson:"end_gamble_nano_ts"`
	FastRecommendPause types.NanoTS `bson:"fast_recommend_pause_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	Gid        ptttype.Bid `bson:"pttgid"`
	Bid        ptttype.Bid `bson:"pttbid"`
	IdxByName  string      `bson:"pttidxname"`
	IdxByClass string      `bson:"pttidxclass"`

	ChessCountry ptttype.ChessCode `bson:"chesscountry"`
}

var EMPTY_BOARD = &Board{}

var (
	BOARD_BBOARD_ID_b      = getBSONName(EMPTY_BOARD, "BBoardID")
	BOARD_BRDNAME_b        = getBSONName(EMPTY_BOARD, "Brdname")
	BOARD_TITLE_b          = getBSONName(EMPTY_BOARD, "Title")
	BOARD_BRD_ATTR_b       = getBSONName(EMPTY_BOARD, "BrdAttr")
	BOARD_BOARD_TYPE_b     = getBSONName(EMPTY_BOARD, "BoardType")
	BOARD_CATEGORY_b       = getBSONName(EMPTY_BOARD, "Category")
	BOARD_NUSER_b          = getBSONName(EMPTY_BOARD, "NUser")
	BOARD_BMS_b            = getBSONName(EMPTY_BOARD, "BMs")
	BOARD_TOTAL_b          = getBSONName(EMPTY_BOARD, "Total")
	BOARD_LAST_POST_TIME_b = getBSONName(EMPTY_BOARD, "LastPostTime")
	BOARD_UDPATE_TIME_b    = getBSONName(EMPTY_BOARD, "UpdateTime")

	BOARD_VOTE_LIMIT_LOGINS_b  = getBSONName(EMPTY_BOARD, "VoteLimitLogins")
	BOARD_POST_LIMIT_LOGINS_b  = getBSONName(EMPTY_BOARD, "PostLimitLogins")
	BOARD_VOTE_LIMIT_BADPOST_b = getBSONName(EMPTY_BOARD, "VoteLimitBadpost")
	BOARD_POST_LIMIT_BADPOST_b = getBSONName(EMPTY_BOARD, "PostLimitBadpost")

	BOARD_PARENT_b = getBSONName(EMPTY_BOARD, "Parent")

	BOARD_NVOTE_b             = getBSONName(EMPTY_BOARD, "NVote")
	BOARD_VOTE_CLOSING_TIME_b = getBSONName(EMPTY_BOARD, "VoteClosingTime")

	BOARD_LEVEL_b                = getBSONName(EMPTY_BOARD, "Level")
	BOARD_LAST_SET_TIME_b        = getBSONName(EMPTY_BOARD, "LastSetTime")
	BOARD_POST_EXPIRE_b          = getBSONName(EMPTY_BOARD, "PostExpire")
	BOARD_END_GAMBLE_NANO_TS_b   = getBSONName(EMPTY_BOARD, "EndGambleNanoTS")
	BOARD_POST_TYPE_b            = getBSONName(EMPTY_BOARD, "PostType")
	BOARD_POST_TEMPLATE_b        = getBSONName(EMPTY_BOARD, "PostTemplate")
	BOARD_FAST_RECOMMEND_PAUSE_b = getBSONName(EMPTY_BOARD, "FastRecommendPause")
	BOARD_IS_DELETED_b           = getBSONName(EMPTY_BOARD, "IsDeleted")
	BOARD_UPDATE_NANO_TS_b       = getBSONName(EMPTY_BOARD, "UpdateNanoTS")

	BOARD_GID_b          = getBSONName(EMPTY_BOARD, "Gid")
	BOARD_BID_b          = getBSONName(EMPTY_BOARD, "Bid")
	BOARD_IDX_BY_NAME_b  = getBSONName(EMPTY_BOARD, "IdxByName")
	BOARD_IDX_BY_CLASS_b = getBSONName(EMPTY_BOARD, "IdxByClass")

	BOARD_CHESS_COUNTRY_b = getBSONName(EMPTY_BOARD, "ChessCountry")
)

func assertBoardFields() error {
	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_SUMMARY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_DETAIL); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_ID); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_BID); err != nil {
		return err
	}

	return nil
}

type BoardQuery struct {
	BBoardID bbs.BBoardID `bson:"bid"`

	IsDeleted interface{} `bson:"deleted,omitempty"`
}

var EMPTY_BOARD_QUERY = &BoardQuery{}
