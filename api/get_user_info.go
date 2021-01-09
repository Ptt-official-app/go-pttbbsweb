package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const GET_USER_INFO_R = "/user/:user_id"

type GetUserInfoParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" url:"fields,omitempty"`
}

type GetUserInfoPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type GetUserInfoResult struct {
	UserID   bbs.UUserID `json:"user_id"`
	Username string      `json:"username"`
	Realname string      `json:"realtime"`
	Nickname string      `json:"nickname"`

	Uflag        ptttype.UFlag `json:"flag"`
	Userlevel    ptttype.PERM  `json:"perm"`
	Numlogindays int           `json:"login_days"`
	Numposts     int           `json:"posts"`
	Firstlogin   types.Time8   `json:"first_login"`
	Lastlogin    types.Time8   `json:"last_login"`
	LastIP       string        `json:"last_ip"`
	LastHost     string        `json:"last_host"` //ip 的中文呈現, 外國則為國家.

	Money    int    `json:"money"`
	PttEmail string `json:"pttemail"`
	Justify  string `json:"justify"`
	Over18   bool   `json:"over18"`

	PagerUIType uint8             `json:"pager_ui"` /* 呼叫器界面類別 (was: WATER_*) */
	Pager       ptttype.PagerMode `json:"pager"`    /* 呼叫器狀態 */
	Invisible   bool              `json:"hide"`
	Exmailbox   uint32            `json:"exmail"`

	Career        string      `json:"career"`
	Role          uint32      `json:"role"`
	LastSeen      types.Time8 `json:"last_seen"`
	TimeSetAngel  types.Time8 `json:"time_set_angel"`
	TimePlayAngel types.Time8 `json:"time_play_angel"`

	LastSong  types.Time8 `json:"last_song"`
	LoginView uint32      `json:"login_view"`

	Vlcount        int `json:"violation"`
	FiveWin        int `json:"five_win"`
	FiveLose       int `json:"five_lose"`
	FiveTie        int `json:"five_tie"`
	ChcWin         int `json:"chc_win"`
	ChcLose        int `json:"chc_lose"`
	ChcTie         int `json:"chc_tie"`
	Conn6Win       int `json:"conn6_win"`
	Conn6Lose      int `json:"conn6_lose"`
	Conn6Tie       int `json:"conn6_tie"`
	GoWin          int `json:"go_win"`
	GoLose         int `json:"go_lose"`
	GoTie          int `json:"go_tie"`
	DarkWin        int `json:"dark_win"`
	DarkLose       int `json:"dark_lose"`
	DarkTie        int `json:"dark_tie"`   /* 暗棋戰績 和 */
	ChessEloRating int `json:"chess_rank"` /* 象棋等級 */

	UaVersion uint8 `json:"ua_version"`

	Signature uint8       `json:"signaure"` /* 慣用簽名檔 */
	BadPost   int         `json:"bad_post"` /* 評價為壞文章數 */
	MyAngel   bbs.UUserID `json:"angel"`    /* 我的小天使 */

	TimeRemoveBadPost types.Time8 `json:"time_remove_bad_post"`
	TimeViolateLaw    types.Time8 `json:"time_violate_law"`

	IsDeleted bool        `json:"deleted"`
	UpdateTS  types.Time8 `json:"update_ts"`

	UserLevel2 ptttype.PERM2 `bson:"perm2"`
	UpdateTS2  types.Time8   `bson:"update_ts2"`

	Avatar   []byte      `json:"avatar"`
	AvatarTS types.Time8 `json:"avatar_ts"`

	Email   string      `json:"email"`
	EmailTS types.Time8 `json:"email_ts"`

	TwoFactorEnabled   bool        `json:"twofactor_enabled"`
	TwoFactorEnabledTS types.Time8 `json:"twofactor_enabled_ts"`

	IDEmail    string      `json:"idemail"`
	IDEmailSet bool        `json:"idemail_set"`
	IDEmailTS  types.Time8 `json:"idemail_ts"`
}

func GetUserInfoWrapper(c *gin.Context) {
	params := &GetUserInfoParams{}
	path := &GetUserInfoPath{}
	LoginRequiredPathQuery(GetUserInfo, params, path, c)
}

func GetUserInfo(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	thePath, ok := path.(*GetUserInfoPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	return tryGetUserInfo(userID, thePath.UserID, c)
}

func tryGetUserInfo(userID bbs.UUserID, queryUserID bbs.UUserID, c *gin.Context) (result *GetUserInfoResult, statusCode int, err error) {
	updateNanoTS := types.NowNanoTS()

	//get backend data
	var result_b pttbbsapi.GetUserResult

	urlMap := make(map[string]string)
	urlMap["uid"] = string(queryUserID)
	url := utils.MergeURL(urlMap, pttbbsapi.GET_USER_R)

	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	userDetail, err := deserializeUserDetailAndUpdateDB(result_b, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	userNewInfo, err := schema.GetUserNewInfo(queryUserID)
	if err != nil {
		return nil, 500, err
	}

	userIDEmail, err := schema.GetUserIDEmailByUserID(queryUserID, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	result = NewUserInfoResult(userDetail, userNewInfo, userIDEmail)

	return result, 200, nil
}

func NewUserInfoResult(userDetail_db *schema.UserDetail, userNewInfo_db *schema.UserNewInfo, userIDEmail_db *schema.UserIDEmail) (result *GetUserInfoResult) {

	if userNewInfo_db == nil {
		userNewInfo_db = &schema.UserNewInfo{}
	}

	if userIDEmail_db == nil {
		userIDEmail_db = &schema.UserIDEmail{}
	}

	logrus.Infof("NewUserInfoResult: userDetail_db: %v userNewInfo_db: %v userIDEmail_db: %v", userDetail_db, userNewInfo_db, userIDEmail_db)

	result = &GetUserInfoResult{
		UserID:   userDetail_db.UserID,
		Username: userDetail_db.Username,
		Realname: userDetail_db.Realname,
		Nickname: userDetail_db.Nickname,

		Uflag:        userDetail_db.Uflag,
		Userlevel:    userDetail_db.Userlevel,
		Numlogindays: userDetail_db.Numlogindays,
		Numposts:     userDetail_db.Numposts,
		Firstlogin:   userDetail_db.Firstlogin.ToTime8(),
		Lastlogin:    userDetail_db.Lastlogin.ToTime8(),
		LastIP:       userDetail_db.LastIP,
		LastHost:     userDetail_db.LastHost,

		Money:    userDetail_db.Money,
		PttEmail: userDetail_db.PttEmail,
		Justify:  userDetail_db.Justify,
		Over18:   userDetail_db.Over18,

		PagerUIType: userDetail_db.PagerUIType,
		Pager:       userDetail_db.Pager,
		Invisible:   userDetail_db.Invisible,
		Exmailbox:   userDetail_db.Exmailbox,

		Career:        userDetail_db.Career,
		Role:          userDetail_db.Role,
		LastSeen:      userDetail_db.LastSeen.ToTime8(),
		TimeSetAngel:  userDetail_db.TimeSetAngel.ToTime8(),
		TimePlayAngel: userDetail_db.TimePlayAngel.ToTime8(),

		LastSong:  userDetail_db.LastSong.ToTime8(),
		LoginView: userDetail_db.LoginView,

		Vlcount:   userDetail_db.Vlcount,
		FiveWin:   userDetail_db.FiveWin,
		FiveLose:  userDetail_db.FiveLose,
		FiveTie:   userDetail_db.FiveTie,
		ChcWin:    userDetail_db.ChcWin,
		ChcLose:   userDetail_db.ChcLose,
		ChcTie:    userDetail_db.ChcTie,
		Conn6Win:  userDetail_db.Conn6Win,
		Conn6Lose: userDetail_db.Conn6Lose,
		Conn6Tie:  userDetail_db.Conn6Tie,
		GoWin:     userDetail_db.GoWin,
		GoLose:    userDetail_db.GoLose,
		GoTie:     userDetail_db.GoTie,
		DarkWin:   userDetail_db.DarkWin,
		DarkLose:  userDetail_db.DarkLose,
		DarkTie:   userDetail_db.DarkTie,
		UaVersion: userDetail_db.UaVersion,

		Signature: userDetail_db.Signature,
		BadPost:   userDetail_db.BadPost,
		MyAngel:   userDetail_db.MyAngel,

		ChessEloRating: userDetail_db.ChessEloRating,

		TimeRemoveBadPost: userDetail_db.TimeRemoveBadPost.ToTime8(),
		TimeViolateLaw:    userDetail_db.TimeViolateLaw.ToTime8(),

		IsDeleted: userDetail_db.IsDeleted,
		UpdateTS:  userDetail_db.UpdateNanoTS.ToTime8(),

		UserLevel2: userDetail_db.UserLevel2,
		UpdateTS2:  userDetail_db.UpdateNanoTS2.ToTime8(),

		Avatar:   userNewInfo_db.Avatar,
		AvatarTS: userNewInfo_db.AvatarNanoTS.ToTime8(),

		Email:   userNewInfo_db.Email,
		EmailTS: userNewInfo_db.EmailNanoTS.ToTime8(),

		TwoFactorEnabled:   userNewInfo_db.TwoFactorEnabled,
		TwoFactorEnabledTS: userNewInfo_db.TwoFactorEnabledNanoTS.ToTime8(),

		IDEmail:    userIDEmail_db.IDEmail,
		IDEmailTS:  userIDEmail_db.UpdateNanoTS.ToTime8(),
		IDEmailSet: userIDEmail_db.IsSet,
	}

	return result
}
