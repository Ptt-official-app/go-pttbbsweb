package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type Board struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	Brdname   string          `bson:"brdname"`
	Title     string          `bson:"title"`
	BrdAttr   ptttype.BrdAttr `bson:"flag"`
	BoardType string          `bson:"the_type"`
	Category  string          `bson:"class"`
	//NUser     int             `bson:"nuser"`  /* use db-count to get current #users */
	BMs   []string `bson:"bms"`
	Total int      `bson:"total"` /* 需要即時知道. 因為 read 頻率高. 並且跟 last-post-time-ts 一樣 write 頻率 << read 頻率 */

	LastPostTimeTS types.Time8 `bson:"last_post_time_ts"` /* 需要即時知道來做板的已讀 */

	UpdateTimeTS types.Time8 `bson:"update_time_ts"` /* XXX 不知 c-pttbbs 對於這個的定義 */

	VoteLimitLogins  int `bson:"vote_limit_logins"`
	PostLimitLogins  int `bson:"post_limit_logins"`
	VoteLimitBadpost int `bson:"vote_limit_bad_post"`
	PostLimitBadpost int `bson:"post_limit_bad_post"`

	Parent bbs.BBoardID `bson:"parent"`

	//NVote             int         `bson:"vote"` /* use db-count to get current #vote */
	VoteClosingTimeTS types.Time8 `bson:"vtime_ts"`

	Level                ptttype.PERM `bson:"perm"`
	LastSetTimeTS        types.Time8  `bson:"last_set_time_ts"`
	PostExpire           int          `bson:"post_expire"`
	EndGambleTS          types.Time8  `bson:"end_gamble_ts"`
	PostType             string       `bson:"post_type"`
	FastRecommendPauseTS types.Time8  `bson:"fast_recommend_pause_ts"`

	IsDeleted bool `bson:"deleted"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
