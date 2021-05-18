package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename0            = "M.1607202239.A.30D"
	testContentAll0          []byte
	testContent0             []byte
	testSignature0           []byte
	testComment0             []byte
	testFirstCommentsDBCS0   []byte
	testTheRestCommentsDBCS0 []byte
	testContent0Big5         [][]*types.Rune
	testContent0Utf8         [][]*types.Rune

	testFirstComments0      []*schema.Comment
	testFullFirstComments0  []*schema.Comment
	testFullFirstComments01 []*schema.Comment
	testFullFirstComments02 []*schema.Comment
)

func initTest0() {
	testContentAll0, testContent0, testSignature0, testComment0, testFirstCommentsDBCS0, testTheRestCommentsDBCS0 = loadTest(testFilename0)

	testContent0Big5 = [][]*types.Rune{ // from python read
		{
			{
				Big5:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
			},
		},
		{
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xc3D] \xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xc3D] \xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
			},
		},
		{
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:03:57 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:03:57 2020"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xa7\xda\xa6b\xad\xfe\xb8\xcc\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xa6b\xad\xfe\xb8\xcc\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xa7\xda\xac\xb0\xa4\xb0\xbb\xf2\xb7|\xa6b\xb3o\xb8\xcc\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xac\xb0\xa4\xb0\xbb\xf2\xb7|\xa6b\xb3o\xb8\xcc\xa9O\xa1H\xa1\xe3"),
			},
		},
	}

	testContent0Utf8 = [][]*types.Rune{ // from python read
		{
			{
				Utf8:   "作者: SYSOP () 看板: WhoAmI",
				Big5:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
			},
		},
		{
			{
				Utf8:   "標題: [問題] 我是誰？～",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xc3D] \xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xc3D] \xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
			},
		},
		{
			{
				Utf8:   "時間: Sun Dec  6 05:03:57 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:03:57 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:03:57 2020"),
			},
		},
		{},
		{
			{
				Utf8:   "我是誰？～",
				Big5:   []byte("\xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xacO\xbd\xd6\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Utf8:   "我在哪裡？～",
				Big5:   []byte("\xa7\xda\xa6b\xad\xfe\xb8\xcc\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xa6b\xad\xfe\xb8\xcc\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Utf8:   "我為什麼會在這裡呢？～",
				Big5:   []byte("\xa7\xda\xac\xb0\xa4\xb0\xbb\xf2\xb7|\xa6b\xb3o\xb8\xcc\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xac\xb0\xa4\xb0\xbb\xf2\xb7|\xa6b\xb3o\xb8\xcc\xa9O\xa1H\xa1\xe3"),
			},
		},
	}

	testFirstComments0 = []*schema.Comment{
		{
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SYSOP"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推",
						Big5:   []byte("\xb1\xc0\xb1\xc0                                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1\xc0\xb1\xc0                                                     "),
					},
				},
			},
			MD5:     "t24G1aV7UjVPoUv-6_T93A",
			TheDate: "12/13 03:51",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0\xb1\xc0                                                     \x1b[m 12/13 03:51"),
		},
	}

	testFullFirstComments0 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("FlAQADI6aAA:t24G1aV7UjVPoUv-6_T93A"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1607802660000000000),
			SortTime:   types.NanoTS(1607802660000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推",
						Big5:   []byte("\xb1\xc0\xb1\xc0                                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1\xc0\xb1\xc0                                                     "),
					},
				},
			},
			MD5:     "t24G1aV7UjVPoUv-6_T93A",
			TheDate: "12/13 03:51",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0\xb1\xc0                                                     \x1b[m 12/13 03:51"),
		},
	}

	testFullFirstComments01 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test01"),
			CommentID:  types.CommentID("FlAQBy5eFAA:t24G1aV7UjVPoUv-6_T93A"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1607802690000000000),
			SortTime:   types.NanoTS(1607802690000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推",
						Big5:   []byte("\xb1\xc0\xb1\xc0                                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1\xc0\xb1\xc0                                                     "),
					},
				},
			},
			MD5:     "t24G1aV7UjVPoUv-6_T93A",
			TheDate: "12/13 03:51",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0\xb1\xc0                                                     \x1b[m 12/13 03:51"),
		},
	}

	testFullFirstComments02 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test02"),
			CommentID:  types.CommentID("FsAZ01_daAA:t24G1aV7UjVPoUv-6_T93A"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1639338660000000000),
			SortTime:   types.NanoTS(1639338660000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推",
						Big5:   []byte("\xb1\xc0\xb1\xc0                                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1\xc0\xb1\xc0                                                     "),
					},
				},
			},
			MD5:     "t24G1aV7UjVPoUv-6_T93A",
			TheDate: "12/13 03:51",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSYSOP\x1b[m\x1b[33m:\xb1\xc0\xb1\xc0                                                     \x1b[m 12/13 03:51"),
		},
	}
}
