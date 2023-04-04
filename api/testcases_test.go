package api

import (
	"io"
	"io/ioutil"
	"os"

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
		FBoardID:  "test9",
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
		Idx:          "3",
		Gid:          3,
		Bid:          9,
	}

	testBoardSummary8 = &apitypes.BoardSummary{
		FBoardID:  "test8",
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
		Idx:          "5",
		Gid:          3,
		Bid:          8,
	}

	testBoardSummary10 = &apitypes.BoardSummary{
		FBoardID:  "WhoAmI",
		Brdname:   "WhoAmI",
		Title:     "呵呵，猜猜我是誰！",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "嘰哩",
		NUser:     0,
		BMs:       []bbs.UUserID{},
		Total:     0,

		LastPostTime: 0,
		StatAttr:     ptttype.NBRD_FAV,
		Idx:          "0",
		Gid:          5,
		Bid:          10,
	}

	testFavoriteBoards0 = []*apitypes.BoardSummary{
		{StatAttr: ptttype.NBRD_LINE, Idx: "0"},
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":1", Idx: "1"},
		{StatAttr: ptttype.NBRD_LINE, Idx: "2"},
		testBoardSummary9,
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":4", Idx: "4"},
		testBoardSummary8,
	}

	testBoardSummaryWhoAmI_b = &bbs.BoardSummary{
		Gid:      5,
		Bid:      10,
		BBoardID: bbs.BBoardID("10_WhoAmI"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "WhoAmI",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xa8, 0xfe, 0xa8, 0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71,
			0xa7, 0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x49,
		},
		BoardType:  []byte{0xa1, 0xb7},
		BM:         []bbs.UUserID{},
		IdxByName:  "WhoAmI",
		IdxByClass: "vFSt-Q@WhoAmI",
	}

	testBoardSummarySYSOP_b = &bbs.BoardSummary{
		Gid:      2,
		Bid:      1,
		BBoardID: bbs.BBoardID("1_SYSOP"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "SYSOP",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		},
		BoardType:  []byte{0xa1, 0xb7},
		BM:         []bbs.UUserID{},
		IdxByName:  "SYSOP",
		IdxByClass: "vFSt-Q@SYSOP",
	}
)

func initTest() {
	if testContentAll4 != nil {
		return
	}

	initTest3()
	initTestUtf83()
	initTest4()
	initTest11()
	initTestUtf85()
	initTestUtf86()
	initTestUtf88()
}

func loadTest(filename string) (contentAll []byte, content []byte, signature []byte, recommend []byte, firstComments []byte, theRestComments []byte) {
	// content-all
	fullFilename := "testcase/" + filename
	file0, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file0.Close()

	r := io.Reader(file0)
	contentAll, _ = ioutil.ReadAll(r)

	// content
	fullFilename = "testcase/" + filename + ".content"
	file1, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file1.Close()

	r = io.Reader(file1)
	content, _ = ioutil.ReadAll(r)

	if len(content) == 0 {
		content = nil
	}

	// signature
	fullFilename = "testcase/" + filename + ".signature"
	file2, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file2.Close()

	r = io.Reader(file2)
	signature, _ = ioutil.ReadAll(r)

	if len(signature) == 0 {
		signature = nil
	}

	// recommend
	fullFilename = "testcase/" + filename + ".recommend"
	file3, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file3.Close()

	r = io.Reader(file3)
	recommend, _ = ioutil.ReadAll(r)

	if len(recommend) == 0 {
		recommend = nil
	}

	// firstComments
	fullFilename = "testcase/" + filename + ".firstComments"
	file4, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file4.Close()

	r = io.Reader(file4)
	firstComments, _ = ioutil.ReadAll(r)

	if len(firstComments) == 0 {
		firstComments = nil
	}

	// theRestComments
	fullFilename = "testcase/" + filename + ".theRestComments"
	file5, err := os.Open(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}
	defer file5.Close()

	r = io.Reader(file5)
	theRestComments, _ = ioutil.ReadAll(r)

	if len(theRestComments) == 0 {
		theRestComments = nil
	}

	return contentAll, content, signature, recommend, firstComments, theRestComments
}
