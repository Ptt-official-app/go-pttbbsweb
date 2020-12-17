package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type User struct {
	UserID   bbs.UUserID `bson:"user_id"`
	Username string      `bson:"username"`
	Realname string      `bson:"realtime"`
	Nickname string      `bson:"nickname"`
	Avatar   []byte      `bson:"avatar"`

	Uflag        ptttype.UFlag `bson:"flag"`
	Userlevel    ptttype.PERM  `bson:"perm"`
	Numlogindays int           `bson:"login_days"` /* 考慮透過 db-count, 但是可能無法跟以前版本相容 */
	Numposts     int           `bson:"posts"`      /* 考慮透過 db-count, 因為 post 可能會被消失, 但是這個欄位不應該因為被消失的而減少? (poster 主動消失的需要減少)? */
	Firstlogin   types.Time8   `bson:"first_login_ts"`
	Lastlogin    types.Time8   `bson:"last_login_ts"` /* 考慮透過 db-max, 但是可能在拉 user-detail 時會花很多時間. */
	LastIP       string        `bson:"last_ip"`
	LastHost     string        `bson:"last_host"` //last-ip 的中文呈現, 外國則為國家.

	Money            int    `bson:"money"`
	Email            string `bson:"email"`
	EmailVerified    string `bson:"email_verified"`
	Phone            string `bson:"phone"` /* 真的要有電話資訊嗎？～ */
	PhoneVerified    string `bson:"phone_verified"`
	TwoFactorEnabled bool   `bson:"twofactor_enabled"`
	Address          string `bson:"address"`
	Justify          string `bson:"justify"`
	Over18           bool   `bson:"over18"`

	PagerUIType uint8             `bson:"pager_ui"` /* 呼叫器界面類別 (was: WATER_*) */
	Pager       ptttype.PagerMode `bson:"pager"`    /* 呼叫器狀態 */
	Invisible   uint8             `bson:"hide"`
	Exmailbox   uint32            `bson:"exmail"`

	Career        string      `bson:"career"`
	Role          uint32      `bson:"role"`
	LastSeen      types.Time8 `bson:"last_seen_ts"`
	TimeSetAngel  types.Time8 `bson:"time_set_angel_ts"`
	TimePlayAngel types.Time8 `bson:"time_play_angel_ts"`

	LastSong  types.Time8 `bson:"last_song"`
	LoginView uint32      `bson:"login_view"`

	Vlcount   int   `bson:"violation"`
	FiveWin   int   `bson:"five_win"`
	FiveLose  int   `bson:"five_lose"`
	FiveTie   int   `bson:"five_tie"`
	ChcWin    int   `bson:"chc_win"`
	ChcLose   int   `bson:"chc_lose"`
	ChcTie    int   `bson:"chc_tie"`
	Conn6Win  int   `bson:"conn6_win"`
	Conn6Lose int   `bson:"conn6_lose"`
	Conn6Tie  int   `bson:"conn6_tie"`
	GoWin     int   `bson:"go_win"`
	GoLose    int   `bson:"go_lose"`
	GoTie     int   `bson:"go_tie"`
	DarkWin   int   `bson:"dark_win"`
	DarkLose  int   `bson:"dark_lose"`
	UaVersion uint8 `bson:"ua_version"`

	Signature uint8  `bson:"signaure"` /* 慣用簽名檔 */
	BadPost   int    `bson:"bad_post"` /* 評價為壞文章數 */
	DarkTie   int    `bson:"dark_tie"` /* 暗棋戰績 和 */
	MyAngel   string `bson:"angel"`    /* 我的小天使 */

	ChessEloRating int `bson:"chess_rank"` /* 象棋等級 */

	TimeRemoveBadPost types.Time8 `bson:"time_remove_bad_post_ts"`
	TimeViolateLaw    types.Time8 `bson:"time_violate_law_ts"`

	IsDeleted    bool         `bson:"deleted"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	//NFriend int `bson:"n_friend"` /* 需要透過 db-count */
}
