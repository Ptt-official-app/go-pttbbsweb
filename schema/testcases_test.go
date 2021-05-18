package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/fav"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
)

var (
	testComments0 = []*Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("foolfighter"),
			CreateTime: types.NanoTS(1261396680001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "土柯糞幾成有台北市戶籍啊？",
						Big5:   []byte("\xa4g\xac_\xc1T\xb4X\xa6\xa8\xa6\xb3\xa5x\xa5_\xa5\xab\xa4\xe1\xc4y\xb0\xda\xa1H                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "Es26f7U0EXdr7Gp4a9N8pQ",
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("yehpi"),
			CreateTime: types.NanoTS(1261396680002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "黃有超大支票（物理）",
						Big5:   []byte("\xb6\xc0\xa6\xb3\xb6W\xa4j\xa4\xe4\xb2\xbc\xa1]\xaa\xab\xb2z\xa1^                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "gmrKWXE7BjV-1U89GcPqHg",
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("lockeyman"),
			CreateTime: types.NanoTS(1261396680003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噗",
						Big5:   []byte("\xbcP                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "cpqbGyLoF_jIyITF4bv-rQ",
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R"),
			RefIDs:     []types.CommentID{"EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"},
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("cheinshin"),
			CreateTime: types.NanoTS(1261396680003100000),
			Content: [][]*types.Rune{
				{},
				{
					{
						Utf8:   "test123123",
						Big5:   []byte("test123123"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
				{},
				{
					{
						Utf8:   "test124124",
						Big5:   []byte("test124124"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
				{},
				{
					{
						Utf8:   "test125125",
						Big5:   []byte("test125125"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "YdlPEuzP2CXzn8nA-n92Ow",
			IP:           "49.216.65.39",
			Host:         "臺灣",
			EditNanoTS:   types.NanoTS(1608551574000000000),
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFuT-Ew6AA:ALE6XIa5ARhXunryJTB3xg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("deathdancer"),
			CreateTime: types.NanoTS(1261410660000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "阿北才沒有輸",
						Big5:   []byte("\xaa\xfc\xa5_\xa4~\xa8S\xa6\xb3\xbf\xe9                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "ALE6XIa5ARhXunryJTB3xg",
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
	}

	testComments1 = []*Comment{
		{ //
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhYQdheQA:a1nbh9m8KYnC0QQevDRyqA"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("carrey8"),
			CreateTime: types.NanoTS(1261396440004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "搞個錘子 連其邁都贏不了?",
						Big5:   []byte("\xb7d\xad\xd3\xc1\xe8\xa4l \xb3s\xa8\xe4\xc1\xda\xb3\xa3\xc4\xb9\xa4\xa3\xa4F?                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "a1nbh9m8KYnC0QQevDRyqA",
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{ //
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv9ryAA:u0m4ezxFyxG8CX56gLwr4Q"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("hyde711034"),
			CreateTime: types.NanoTS(1261396500000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "台南人可憐還開心指數最高",
						Big5:   []byte("\xa5x\xabn\xa4H\xa5i\xbc\xa6\xc1\xd9\xb6}\xa4\xdf\xab\xfc\xbc\xc6\xb3\xcc\xb0\xaa                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:          "u0m4ezxFyxG8CX56gLwr4Q",
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
	}

	testReply0 = &Comment{ //41 (43)
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ:R"),
		RefIDs:     []types.CommentID{"EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ"},
		TheType:    types.COMMENT_TYPE_REPLY,
		Owner:      bbs.UUserID("test123"),
		CreateTime: types.NanoTS(1234567890000000000),
		Content: [][]*types.Rune{
			{},
			{
				{
					Utf8:   "     123123 123123 123123 123123     ",
					Big5:   []byte("     123123 123123 123123 123123     "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
				{
					Utf8:   "     456456 456456 456456 456456     ",
					Big5:   []byte("     456456 456456 456456 456456     "),
					Color0: types.DefaultColor,
					Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_CYAN},
				},
			},
			{},
			{},
			{
				{
					Utf8:   "     123123 123123 123123 123123     ",
					Big5:   []byte("     123123 123123 123123 123123     "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{
				{
					Utf8:   "※ 編輯: cheinshin (49.216.65.39 臺灣), 12/21/2020 19:52:54        ",
					Big5:   []byte(""),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{},
			{},
			{},
		},
		IP:         "localhost",
		Host:       "臺灣",
		MD5:        "m4zgWiP6HQDKfLxk-6bAhw",
		EditNanoTS: types.NanoTS(1234567890000000001),
	}

	testExpectedReply0 = &Comment{ //41 (43)
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ:R"),
		RefIDs:     []types.CommentID{"EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ"},
		TheType:    types.COMMENT_TYPE_REPLY,
		Owner:      bbs.UUserID("test123"),
		CreateTime: types.NanoTS(1234567890000000000),
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "     123123 123123 123123 123123     ",
					Big5:   []byte("     123123 123123 123123 123123     "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
				{
					Utf8:   "     456456 456456 456456 456456     ",
					Big5:   []byte("     456456 456456 456456 456456     "),
					Color0: types.DefaultColor,
					Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_CYAN},
				},
			},
			{},
			{},
			{
				{
					Utf8:   "     123123 123123 123123 123123     ",
					Big5:   []byte("     123123 123123 123123 123123     "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
		},
		IP:         "localhost",
		Host:       "臺灣",
		MD5:        "m4zgWiP6HQDKfLxk-6bAhw",
		EditNanoTS: types.NanoTS(1234567890000000001),
	}

	testReply1 = &Comment{ //41 (43)
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ:R"),
		RefIDs:     []types.CommentID{"EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ"},
		TheType:    types.COMMENT_TYPE_REPLY,
		Owner:      bbs.UUserID("test123"),
		CreateTime: types.NanoTS(1234567890000000000),
		Content: [][]*types.Rune{
			{},
			{},
			{},
			{},
			{},
			{
				{
					Utf8:   "※ 編輯: cheinshin (49.216.65.39 臺灣), 12/21/2020 19:52:54        ",
					Big5:   []byte(""),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{},
			{},
			{},
		},
		IP:         "localhost",
		Host:       "臺灣",
		MD5:        "m4zgWiP6HQDKfLxk-6bAhw",
		EditNanoTS: types.NanoTS(1234567890000000001),
	}

	testExpectedReply1 = &Comment{ //41 (43)
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ:R"),
		RefIDs:     []types.CommentID{"EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ"},
		TheType:    types.COMMENT_TYPE_REPLY,
		Owner:      bbs.UUserID("test123"),
		CreateTime: types.NanoTS(1234567890000000000),
		Content:    nil,
		IP:         "localhost",
		Host:       "臺灣",
		MD5:        "m4zgWiP6HQDKfLxk-6bAhw",
		EditNanoTS: types.NanoTS(1234567890000000001),
	}

	testCommentSummaries0 = []*CommentSummary{
		{
			CommentID:    types.CommentID("EYFhbv9ryAA:u0m4ezxFyxG8CX56gLwr4Q"),
			CreateTime:   types.NanoTS(1261396500000000000),
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			CommentID:    types.CommentID("EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ"),
			CreateTime:   types.NanoTS(1261396680001000000),
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			CommentID:    types.CommentID("EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg"),
			CreateTime:   types.NanoTS(1261396680002000000),
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			CommentID:    types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"),
			CreateTime:   types.NanoTS(1261396680003000000),
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
		{
			CommentID:    types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R"),
			CreateTime:   types.NanoTS(1261396680003100000),
			UpdateNanoTS: types.NanoTS(1334567890000000000),
		},
	}

	testUser0 = &bbs.Userec{
		UUserID:  "test_userid",
		Username: "test_username",
		Realname: []byte("test_realname"),
		Nickname: []byte("test_nickname"),

		Uflag:        1,
		Userlevel:    2,
		Numlogindays: 11,
		Numposts:     12,
		Firstlogin:   1234567890,
		Lastlogin:    1234567891,
		Lasthost:     "localhost",

		Money:   13,
		Email:   "test@ptt.test",
		Justify: []byte("Email: test@ptt.test 12/23/2008 12:13:20"),
		Over18:  true,

		PagerUIType: 1,
		Pager:       12,
		Invisible:   false,
		Exmailbox:   10,

		Career:        []byte("test_career"),
		Role:          1,
		LastSeen:      1234567892,
		TimeSetAngel:  1234567893,
		TimePlayAngel: 1234567894,

		LastSong:  1234567895,
		LoginView: 1,

		Vlcount:  12,
		FiveWin:  13,
		FiveLose: 14,
		FiveTie:  15,

		UaVersion: 3,

		Signature: 3,
		BadPost:   12,
		MyAngel:   "test_myangel",

		ChessEloRating: 1,

		TimeRemoveBadPost: 1234567895,
		TimeViolateLaw:    1234567897,
	}

	testUserDetail0 = &UserDetail{
		UserID:   "test_userid",
		Username: "test_username",
		Realname: "test_realname",
		Nickname: "test_nickname",

		Uflag:        1,
		Userlevel:    2,
		Numlogindays: 11,
		Numposts:     12,
		Firstlogin:   1234567890000000000,
		Lastlogin:    1234567891000000000,
		LastIP:       "localhost",
		LastHost:     "localhost",

		Money:    13,
		PttEmail: "test@ptt.test",
		Justify:  "Email: test@ptt.test 12/23/2008 12:13:20",
		Over18:   true,

		PagerUIType: 1,
		Pager:       12,
		Invisible:   false,
		Exmailbox:   10,

		Career:        "test_career",
		Role:          1,
		LastSeen:      1234567892000000000,
		TimeSetAngel:  1234567893000000000,
		TimePlayAngel: 1234567894000000000,

		LastSong:  1234567895000000000,
		LoginView: 1,

		Vlcount:  12,
		FiveWin:  13,
		FiveLose: 14,
		FiveTie:  15,

		UaVersion: 3,

		Signature: 3,
		BadPost:   12,
		MyAngel:   "test_myangel",

		ChessEloRating: 1,

		TimeRemoveBadPost: 1234567895000000000,
		TimeViolateLaw:    1234567897000000000,

		UpdateNanoTS: 123456790000000000,
	}

	testUserInfoSummary0 = &UserInfoSummary{
		UserID:       "test_userid",
		UpdateNanoTS: 1234567890000000003,
	}

	testUserNewInfo0 = &UserNewInfo{}
)

var (
	testTitle0  = "新的目錄"
	testSubFav0 = &fav.Fav{
		FavNum:   1,
		Depth:    1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*fav.FavType{
			{
				FavIdx:  6,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 1, LastVisit: 0, Attr: 0},
			},
		},
	}

	testFav0 = &fav.Fav{
		NBoards:  3,
		NLines:   2,
		NFolders: 1,
		LineID:   2,
		FolderID: 1,
		FavNum:   7,
		Favh: []*fav.FavType{
			{
				FavIdx:  0,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 1, LastVisit: 0, Attr: 0},
			},
			{
				FavIdx:  1,
				TheType: pttbbsfav.FAVT_LINE,
				Attr:    1,
				Fp:      &fav.FavLine{Lid: 1},
			},
			{
				FavIdx:  2,
				TheType: pttbbsfav.FAVT_FOLDER,
				Attr:    1,
				Fp:      &fav.FavFolder{Fid: 1, Title: testTitle0, ThisFolder: testSubFav0},
			},
			{
				FavIdx:  3,
				TheType: pttbbsfav.FAVT_LINE,
				Attr:    1,
				Fp:      &fav.FavLine{Lid: 2},
			},
			{
				FavIdx:  4,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 9, LastVisit: 0, Attr: 0},
			},
			{
				FavIdx:  5,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 8, LastVisit: 0, Attr: 0},
			},
		},
	}

	testFavMeta0 = &UserFavoritesMeta{
		UserID:       "SYSOP",
		UpdateNanoTS: 1234567890000000000,
		MTime:        1234567890000000000,
		FolderMeta: FolderMeta{
			FavNum:   7,
			NBoards:  3,
			NLines:   2,
			NFolders: 1,
		},
	}

	testUserFavorites0 = []*UserFavorites{
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          0,
			LevelIdx:        "",
			Idx:             0,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   1,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          1,
			LevelIdx:        "",
			Idx:             1,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_LINE,
			Attr:    1,
			TheID:   1,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          2,
			LevelIdx:        "",
			Idx:             2,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_FOLDER,
			Attr:    1,
			TheID:   1,

			FolderTitle: testTitle0,
			FolderMeta: &FolderMeta{
				FavNum:  1,
				NBoards: 1,
			},
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          3,
			LevelIdx:        "",
			Idx:             3,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_LINE,
			Attr:    1,
			TheID:   2,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          4,
			LevelIdx:        "",
			Idx:             4,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   9,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          5,
			LevelIdx:        "",
			Idx:             5,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   8,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          6,
			LevelIdx:        ":2",
			Idx:             0,
			UpdateNanoTS:    1234567890000000000,
			MTime:           1234567890000000000,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   1,
		},
	}

	testSubFav1 = &fav.Fav{
		FavNum:   1,
		Depth:    1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*fav.FavType{
			{
				FavIdx:  6,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 1, LastVisit: 0, Attr: 0},
			},
		},
	}

	testSubFav2 = &fav.Fav{
		FavNum:   1,
		Depth:    1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*fav.FavType{
			{
				FavIdx:  7,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 1, LastVisit: 0, Attr: 0},
			},
		},
	}

	testFav1 = &fav.Fav{
		FavNum:   8,
		NBoards:  2,
		NLines:   2,
		NFolders: 2,
		LineID:   2,
		FolderID: 2,
		Favh: []*fav.FavType{
			{
				FavIdx:  0,
				TheType: pttbbsfav.FAVT_LINE,
				Attr:    1,
				Fp:      &fav.FavLine{Lid: 1},
			},
			{
				FavIdx:  1,
				TheType: pttbbsfav.FAVT_FOLDER,
				Attr:    1,
				Fp:      &fav.FavFolder{Fid: 1, Title: testTitle0, ThisFolder: testSubFav1},
			},
			{
				FavIdx:  2,
				TheType: pttbbsfav.FAVT_LINE,
				Attr:    1,
				Fp:      &fav.FavLine{Lid: 2},
			},
			{
				FavIdx:  3,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 9, LastVisit: 0, Attr: 0},
			},
			{
				FavIdx:  4,
				TheType: pttbbsfav.FAVT_FOLDER,
				Attr:    1,
				Fp:      &fav.FavFolder{Fid: 2, Title: testTitle0, ThisFolder: testSubFav2},
			},
			{
				FavIdx:  5,
				TheType: pttbbsfav.FAVT_BOARD,
				Attr:    1,
				Fp:      &fav.FavBoard{Bid: 8, LastVisit: 0, Attr: 0},
			},
		},
	}

	testFavMeta1 = &UserFavoritesMeta{
		UserID:       "SYSOP",
		UpdateNanoTS: 1234567890000000001,
		MTime:        1234567890000000001,
		FolderMeta: FolderMeta{
			FavNum:   8,
			NBoards:  2,
			NLines:   2,
			NFolders: 2,
		},
	}

	testUserFavorites1 = []*UserFavorites{
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          0,
			LevelIdx:        "",
			Idx:             0,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType: pttbbsfav.FAVT_LINE,
			Attr:    1,
			TheID:   1,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          1,
			LevelIdx:        "",
			Idx:             1,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType:     pttbbsfav.FAVT_FOLDER,
			Attr:        1,
			TheID:       1,
			FolderTitle: testTitle0,
			FolderMeta: &FolderMeta{
				FavNum:  1,
				NBoards: 1,
			},
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          2,
			LevelIdx:        "",
			Idx:             2,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType: pttbbsfav.FAVT_LINE,
			Attr:    1,
			TheID:   2,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          3,
			LevelIdx:        "",
			Idx:             3,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   9,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          4,
			LevelIdx:        "",
			Idx:             4,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType:     pttbbsfav.FAVT_FOLDER,
			Attr:        1,
			TheID:       2,
			FolderTitle: testTitle0,
			FolderMeta: &FolderMeta{
				FavNum:  1,
				NBoards: 1,
			},
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          5,
			LevelIdx:        "",
			Idx:             5,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   8,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          6,
			LevelIdx:        ":1",
			Idx:             0,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   1,
		},
		{
			UserID:          "SYSOP",
			DoubleBufferIdx: 0,
			FavIdx:          7,
			LevelIdx:        ":4",
			Idx:             0,
			UpdateNanoTS:    1234567890000000001,
			MTime:           1234567890000000001,

			TheType: pttbbsfav.FAVT_BOARD,
			Attr:    1,
			TheID:   1,
		},
	}
)
