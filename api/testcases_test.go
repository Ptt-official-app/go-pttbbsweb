package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUserInfoResult0 = &GetUserInfoResult{
		UserID:   "SYSOP",
		Username: "SYSOP",
		Realname: "test_sysop",
		Nickname: "test_nickname",

		Uflag:        1,
		Userlevel:    7,
		Numlogindays: 1,
		Numposts:     2,
		Firstlogin:   1234567890,
		Lastlogin:    1234567891,
		LastIP:       "127.0.0.1",
		LastHost:     "127.0.0.1",

		Money:    11,
		PttEmail: "test@ptt.test",
		Justify:  "email: test@ptt.test",
		Over18:   true,

		PagerUIType: 1,
		Pager:       10,
		Invisible:   false,
		Exmailbox:   120,

		Career:        "test-career",
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
	}

	testTitle0 = "新的目錄"

	testBoardSummary9 = &apitypes.BoardSummary{
		BBoardID:  "9_test9",
		Brdname:   "test9",
		Title:     "測試9",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     100,
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,

		LastPostTime: 1234567890,
		StatAttr:     ptttype.NBRD_BOARD,
	}

	testBoardSummary8 = &apitypes.BoardSummary{
		BBoardID:  "8_test8",
		Brdname:   "test8",
		Title:     "測試8",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     101,
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,

		LastPostTime: 1300000000,
		StatAttr:     ptttype.NBRD_BOARD,
	}

	testFavoriteBoards0 = []*apitypes.BoardSummary{
		{StatAttr: ptttype.NBRD_LINE},
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":1"},
		{StatAttr: ptttype.NBRD_LINE},
		testBoardSummary9,
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":4"},
		testBoardSummary8,
	}
)
