package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testFilename4            = "M.1608388506.A.85D"
	testContentAll4          []byte
	testContent4             []byte
	testSignature4           []byte
	testComment4             []byte
	testFirstCommentsDBCS4   []byte
	testTheRestCommentsDBCS4 []byte
	testContent4Big5         [][]*types.Rune
	testContent4Utf8         [][]*types.Rune

	testFirstComments4     []*schema.Comment
	testFullFirstComments4 []*schema.Comment
)

func initTest4() {
	testContentAll4, testContent4, testSignature4, testComment4, testFirstCommentsDBCS4, testTheRestCommentsDBCS4 = loadTest(testFilename4)

	testContent4Utf8 = [][]*types.Rune{
		{
			{
				Utf8:    "作者: SYSOP () 看板: WhoAmI",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: SYSOP () 看板: WhoAmI",
			},
		},
		{
			{
				Utf8:    "標題: [閒聊] 所以特殊字真的是有綠色的～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [閒聊] 所以特殊字真的是有綠色的～",
			},
		},
		{
			{
				Utf8:    "時間: Sat Dec 19 22:35:04 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sat Dec 19 22:35:04 2020",
			},
		},
		{},
		{
			{
				Utf8:    "然後 \\n 不會在 big5 結尾. 可以放心直接用 \\n 斷行.",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "然後 \\n 不會在 big5 結尾. 可以放心直接用 \\n 斷行.",
			},
		},
		{
			{
				Utf8:    "我是許功蓋",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "我是許功蓋",
			},
		},
	}

	testFirstComments4 = []*schema.Comment{
		{
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SYSOP"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                       "),
						DBCS:   []byte("\xb1\xc0                                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:     "3dK46zmOe5zmna12AC1gnQ",
			TheDate: "12/19 22:35",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0                                                       \x1b[m 12/19 22:35"),
		},
		{
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("chhsiao123"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                  "),
						DBCS:   []byte("\xb1\xc0                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:     "FQaNH8WkdAbEGD7yp2Zkvg",
			TheDate: "12/19 22:36",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mchhsiao123\x1b[m\x1b[33m:\xb1\xc0                                                  \x1b[m 12/19 22:36"),
		},
		{
			TheType: ptttype.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("chhsiao123"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噓～",
						Big5:   []byte("\xbcN\xa1\xe3                                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbcN\xa1\xe3                                                "),
					},
				},
			},
			MD5:     "cLGi8fC4fapuiBkTXHU2OA",
			TheDate: "12/19 22:37",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mchhsiao123\x1b[m\x1b[33m:\xbcN\xa1\xe3                                                \x1b[m 12/19 22:37"),
		},
	}

	testFullFirstComments4 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("10_WhoAmI"),
			ArticleID:  bbs.ArticleID("1VrooM21"),
			CommentID:  types.CommentID("FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ"),
			TheType:    ptttype.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1608388500000000000),
			SortTime:   types.NanoTS(1608388500000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1\xc0                                                       "),
					},
				},
			},
			MD5:     "3dK46zmOe5zmna12AC1gnQ",
			TheDate: "12/19 22:35",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0                                                       \x1b[m 12/19 22:35"),
		},
		{
			BBoardID:   bbs.BBoardID("10_WhoAmI"),
			ArticleID:  bbs.ArticleID("1VrooM21"),
			CommentID:  types.CommentID("FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg"),
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("chhsiao123"),
			CreateTime: types.NanoTS(1608388560000000000),
			SortTime:   types.NanoTS(1608388560000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1\xc0                                                  "),
					},
				},
			},
			MD5:     "FQaNH8WkdAbEGD7yp2Zkvg",
			TheDate: "12/19 22:36",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mchhsiao123\x1b[m\x1b[33m:\xb1\xc0                                                  \x1b[m 12/19 22:36"),
		},
		{
			BBoardID:   bbs.BBoardID("10_WhoAmI"),
			ArticleID:  bbs.ArticleID("1VrooM21"),
			CommentID:  types.CommentID("FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA"),
			TheType:    ptttype.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("chhsiao123"),
			CreateTime: types.NanoTS(1608388624000000000),
			SortTime:   types.NanoTS(1608388624000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噓～",
						Big5:   []byte("\xbcN\xa1\xe3                                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbcN\xa1\xe3                                                "),
					},
				},
			},
			MD5:     "cLGi8fC4fapuiBkTXHU2OA",
			TheDate: "12/19 22:37",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mchhsiao123\x1b[m\x1b[33m:\xbcN\xa1\xe3                                                \x1b[m 12/19 22:37"),
		},
	}
}
