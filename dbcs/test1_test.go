package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename1            = "M.1607203395.A.00D"
	testContentAll1          []byte
	testContent1             []byte
	testSignature1           []byte
	testComment1             []byte
	testFirstCommentsDBCS1   []byte
	testTheRestCommentsDBCS1 []byte
	testContent1Big5         [][]*types.Rune
	testContent1Utf8         [][]*types.Rune

	testFirstComments1 []*schema.Comment
)

func initTest1() {
	testContentAll1, testContent1, testSignature1, testComment1, testFirstCommentsDBCS1, testTheRestCommentsDBCS1 = loadTest(testFilename1)

	testContent1Big5 = [][]*types.Rune{ // from python read
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
				Big5:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] \xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] \xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
			},
		},
		{
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:23:13 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:23:13 2020"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xb5M\xab\xe1\xa9O\xa1\xe3\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb5M\xab\xe1\xa9O\xa1\xe3\xa1H"),
			},
		},
		{},
	}

	testContent1Utf8 = [][]*types.Rune{ // from python read
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
				Utf8:   "標題: [心得] 然後呢？～",
				Big5:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] \xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] \xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
			},
		},
		{
			{
				Utf8:   "時間: Sun Dec  6 05:23:13 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:23:13 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sun Dec  6 05:23:13 2020"),
			},
		},
		{},
		{
			{
				Utf8:   "然後呢～？",
				Big5:   []byte("\xb5M\xab\xe1\xa9O\xa1\xe3\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb5M\xab\xe1\xa9O\xa1\xe3\xa1H"),
			},
		},
		{},
	}

	testFirstComments1 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EX64IZAFEAA:uzQhiFhT_R5HGcpcXa10Fg"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1260647400000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推～",
						Big5:   []byte("\xb1\xc0\xb1\xc0\xa1\xe3                                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "uzQhiFhT_R5HGcpcXa10Fg",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EX64L4hMaAA:t24G1aV7UjVPoUv-6_T93A"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1260647460000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推",
						Big5:   []byte("\xb1\xc0\xb1\xc0                                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "t24G1aV7UjVPoUv-6_T93A",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EX64PYCTwAA:HWj0bTiKkHE3DnJEz3TO0A"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1260647520000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推推",
						Big5:   []byte("\xb1\xc0\xb1\xc0                                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "HWj0bTiKkHE3DnJEz3TO0A",
		},
	}
}
