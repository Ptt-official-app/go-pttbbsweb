package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var (
	testUtf8Filename0            = "M.1644505564.A.172.utf8"
	testUtf8ContentAll0          []byte
	testUtf8Content0             []byte
	testUtf8Signature0           []byte
	testUtf8Comment0             []byte
	testUtf8FirstCommentsDBCS0   []byte
	testUtf8TheRestCommentsDBCS0 []byte
	testUtf8Content0Utf8         [][]*types.Rune

	testUtf8FirstComments0      []*schema.Comment
	testUtf8FullFirstComments0  []*schema.Comment
	testUtf8FullFirstComments01 []*schema.Comment
	testUtf8FullFirstComments02 []*schema.Comment
	testUtf8TheRestComments0    []*schema.Comment
)

func initTestUtf80() {
	testUtf8ContentAll0, testUtf8Content0, testUtf8Signature0, testUtf8Comment0, testUtf8FirstCommentsDBCS0, testUtf8TheRestCommentsDBCS0 = loadTest(testUtf8Filename0)

	testUtf8Content0Utf8 = [][]*types.Rune{ // from python read
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
				Utf8:    "標題: [問題] 我是誰？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [問題] 我是誰？～",
			},
		},
		{
			{
				Utf8:    "時間: Sun Dec  6 05:03:57 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sun Dec  6 05:03:57 2020",
			},
		},
		{},
		{
			{
				Utf8:    "我是誰？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "我是誰？～",
			},
		},
		{},
		{
			{
				Utf8:    "我在哪裡？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "我在哪裡？～",
			},
		},
		{},
		{
			{
				Utf8:    "我為什麼會在這裡呢？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "我為什麼會在這裡呢？～",
			},
		},
	}

	testUtf8FirstComments0 = []*schema.Comment{
		{
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SYSOP"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "推推",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "推推",
					},
				},
			},
			MD5:     "8eWYZUhy5fspMX9cgnhwPg",
			TheDate: "12/13 03:51",
			DBCSStr: "\x1b[1;31m→ \x1b[33mSYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:51",
		},
	}

	testUtf8FullFirstComments0 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("testUtf8"),
			ArticleID:  bbs.ArticleID("testUtf8"),
			CommentID:  types.CommentID("FlAQADI6aAA:8eWYZUhy5fspMX9cgnhwPg"),
			TheType:    ptttype.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1607802660000000000),
			SortTime:   types.NanoTS(1607802660000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "推推",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "推推",
					},
				},
			},
			MD5:     "8eWYZUhy5fspMX9cgnhwPg",
			TheDate: "12/13 03:51",
			DBCSStr: "\x1b[1;31m→ \x1b[33mSYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:51",
		},
	}

	testUtf8FullFirstComments01 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("testUtf8"),
			ArticleID:  bbs.ArticleID("testUtf801"),
			CommentID:  types.CommentID("FlAQBy5eFAA:8eWYZUhy5fspMX9cgnhwPg"),
			TheType:    ptttype.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("SYSOP"),
			CreateTime: types.NanoTS(1607802690000000000),
			SortTime:   types.NanoTS(1607802690000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "推推",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "推推",
					},
				},
			},
			MD5:     "8eWYZUhy5fspMX9cgnhwPg",
			TheDate: "12/13 03:51",
			DBCSStr: "\x1b[1;31m→ \x1b[33mSYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:51",
		},
	}
}
