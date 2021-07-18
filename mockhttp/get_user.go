package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func GetUser() (ret api.GetUserResult) {
	return api.GetUserResult(&bbs.Userec{
		UUserID:  "SYSOP",
		Username: "SYSOP",
		Realname: []byte("test_sysop"),
		Nickname: []byte("test_nickname"),

		Uflag:        1,
		Userlevel:    7,
		Numlogindays: 1,
		Numposts:     2,
		Firstlogin:   1234567890,
		Lastlogin:    1234567891,
		Lasthost:     "127.0.0.1",

		Money:   11,
		Email:   "test@ptt.test",
		Justify: []byte("email: test@ptt.test"),
		Over18:  true,

		PagerUIType: 1,
		Pager:       10,
		Exmailbox:   120,

		Career:        []byte("test-career"),
		Role:          12,
		LastSeen:      1234567893,
		TimeSetAngel:  1234567894,
		TimePlayAngel: 1234567895,

		LastSong:  1234567900,
		LoginView: 12,

		Vlcount:  123,
		FiveWin:  124,
		FiveLose: 125,

		Signature: 6,

		BadPost: 134,
		MyAngel: "myangel",

		ChessEloRating: 126,
	})
}
