package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const GET_USER_INFO_R = "/user/:user_id"

type GetUserInfoParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" url:"fields,omitempty"`
}

type GetUserInfoPath struct {
	UserID string `uri:"user_id"`
}

type GetUserInfoResult struct {
	UserID   string `json:"user_id"`
	Realname string `json:"realtime"`
	Nickname string `json:"nickname"`

	Uflag        ptttype.UFlag `json:"flag"`
	Userlevel    ptttype.PERM  `json:"perm"`
	Numlogindays int           `json:"login_days"`
	Numposts     int           `json:"posts"`
	Firstlogin   types.Time8   `json:"first_login"`
	Lastlogin    types.Time8   `json:"last_login"`
	Lasthost     string        `json:"last_ip"`
	Country      string        `json:"country"`

	Money   int    `json:"money"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Justify string `json:"justify"`
	Over18  bool   `json:"over18"`

	PagerUIType uint8             `json:"pager_ui"` /* 呼叫器界面類別 (was: WATER_*) */
	Pager       ptttype.PagerMode `json:"pager"`    /* 呼叫器狀態 */
	Invisible   uint8             `json:"hide"`
	Exmailbox   uint32            `json:"exmail"`

	Career        string      `json:"career"`
	Role          uint32      `json:"role"`
	LastSeen      types.Time8 `json:"last_seen"`
	TimeSetAngel  types.Time8 `json:"time_set_angel"`
	TimePlayAngel types.Time8 `json:"time_play_angel"`

	LastSong  types.Time8 `json:"last_song"`
	LoginView uint32      `json:"login_view"`

	Vlcount   int   `json:"violation"`
	FiveWin   int   `json:"five_win"`
	FiveLose  int   `json:"five_lose"`
	FiveTie   int   `json:"five_tie"`
	ChcWin    int   `json:"chc_win"`
	ChcLose   int   `json:"chc_lose"`
	ChcTie    int   `json:"chc_tie"`
	Conn6Win  int   `json:"conn6_win"`
	Conn6Lose int   `json:"conn6_lose"`
	Conn6Tie  int   `json:"conn6_tie"`
	GoWin     int   `json:"go_win"`
	GoLose    int   `json:"go_lose"`
	GoTie     int   `json:"go_tie"`
	DarkWin   int   `json:"dark_win"`
	DarkLose  int   `json:"dark_lose"`
	UaVersion uint8 `json:"ua_version"`

	Signature uint8  `json:"signaure"` /* 慣用簽名檔 */
	BadPost   int    `json:"bad_post"` /* 評價為壞文章數 */
	DarkTie   int    `json:"dark_tie"` /* 暗棋戰績 和 */
	MyAngel   string `json:"angel"`    /* 我的小天使 */

	ChessEloRating int `json:"chess_rank"` /* 象棋等級 */

	TimeRemoveBadPost types.Time8 `json:"time_remove_bad_post"`
	TimeViolateLaw    types.Time8 `json:"time_violate_law"`

	NFriend int `json:"n_friend"`
}

func GetUserInfo(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	result = &GetUserInfoResult{
		UserID:   "ckoool",
		Realname: "我是西k屋",
		Nickname: "我是西k屋",

		Userlevel:    ptttype.PERM_BASIC | ptttype.PERM_ANGEL,
		Numlogindays: 12345,
		Numposts:     123124,
		Firstlogin:   types.Time8(1234567890),
		Lastlogin:    types.Time8(1800000000),
		Lasthost:     "127.0.0.1",
		Money:        2114567890,

		Email:     "test@test.test",
		Address:   "台北市天龍區天龍路 87 號",
		Justify:   "[Email] test@test.test",
		Over18:    true,
		Invisible: 1,
		Exmailbox: 123,

		Career:        "某大學",
		LastSeen:      types.Time8(1800000010),
		TimeSetAngel:  types.Time8(1234568890),
		TimePlayAngel: types.Time8(123458900),

		Vlcount:   5,
		FiveWin:   1203,
		FiveLose:  312,
		FiveTie:   120,
		ChcWin:    1234,
		ChcLose:   312,
		ChcTie:    12,
		Conn6Win:  12,
		Conn6Lose: 4,
		Conn6Tie:  5,
		GoWin:     123,
		GoLose:    43,
		GoTie:     15,
		DarkWin:   1234,
		DarkLose:  12,
		DarkTie:   312,
		UaVersion: 16,

		BadPost: 123,
		MyAngel: "teamore",
	}

	return result, 200, nil
}
