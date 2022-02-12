package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUtf8Filename2            = "M.1644505576.A.831.utf8"
	testUtf8ContentAll2          []byte
	testUtf8Content2             []byte
	testUtf8Signature2           []byte
	testUtf8Comment2             []byte
	testUtf8FirstCommentsDBCS2   []byte
	testUtf8TheRestCommentsDBCS2 []byte
	testUtf8Content2Utf8         [][]*types.Rune

	testUtf8FirstComments2     []*schema.Comment
	testUtf8FullFirstComments2 []*schema.Comment
	testUtf8TheRestComments2   []*schema.Comment
)

func initTestUtf82() {
	testUtf8ContentAll2, testUtf8Content2, testUtf8Signature2, testUtf8Comment2, testUtf8FirstCommentsDBCS2, testUtf8TheRestCommentsDBCS2 = loadTest(testUtf8Filename2)

	testUtf8Content2Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:    "作者: SYSOP () 看板: WhoAmI",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: SYSOP () 看板: WhoAmI",
			},
		},
		{ // 1
			{
				Utf8:    "標題: [心得] 然後呢？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [心得] 然後呢？～",
			},
		},
		{ // 2
			{
				Utf8:    "時間: Sun Dec  6 05:23:13 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sun Dec  6 05:23:13 2020",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "然後呢～？",
				DBCSStr: "然後呢～？",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
	}

	testUtf8FirstComments2 = []*schema.Comment{
		{
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SYSOP"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "推推～",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "推推～",
					},
				},
			},
			MD5:     "PicQyDahvJqr2gRhkG05rA",
			TheDate: "12/13 03:50",
			DBCSStr: "\x1b[1;31m→ \x1b[33mSYSOP\x1b[m\x1b[33m:推推～                                                   \x1b[m 12/13 03:50",
		},
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
			MD5:     "xrp69NP3-c8UXclsSEf_Rg",
			TheDate: "12/13 03:52",
			DBCSStr: "\x1b[1;31m→ \x1b[33mSYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:52",
		},
	}
	testUtf8FullFirstComments2 = []*schema.Comment{}
	testUtf8TheRestComments2 = []*schema.Comment{}
}
