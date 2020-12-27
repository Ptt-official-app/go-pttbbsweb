package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
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
			MD5:             "a1nbh9m8KYnC0QQevDRyqA",
			IsFirstComments: true,
			UpdateNanoTS:    types.NanoTS(1334567890000000000),
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
			MD5:             "u0m4ezxFyxG8CX56gLwr4Q",
			IsFirstComments: true,
			UpdateNanoTS:    types.NanoTS(1334567890000000000),
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
)
