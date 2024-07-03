package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var (
	testUtf8Filename3            = "M.1644505590.A.80C.utf8"
	testUtf8ContentAll3          []byte
	testUtf8Content3             []byte
	testUtf8Signature3           []byte
	testUtf8Comment3             []byte
	testUtf8FirstCommentsDBCS3   []byte
	testUtf8TheRestCommentsDBCS3 []byte
	testUtf8Content3Utf8         [][]*types.Rune

	testUtf8FirstComments3     []*schema.Comment
	testUtf8FullFirstComments3 []*schema.Comment
	testUtf8TheRestComments3   []*schema.Comment
)

func initTestUtf83() {
	testUtf8ContentAll3, testUtf8Content3, testUtf8Signature3, testUtf8Comment3, testUtf8FirstCommentsDBCS3, testUtf8TheRestCommentsDBCS3 = loadTest(testUtf8Filename3)

	testUtf8Content3Utf8 = [][]*types.Rune{
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
				Utf8:    "標題: [討論] 怎樣子才可以賺 p 幣呢～？",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [討論] 怎樣子才可以賺 p 幣呢～？",
			},
		},
		{ // 3
			{
				Utf8:    "時間: Mon Dec 14 17:12:52 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Mon Dec 14 17:12:52 2020",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "揪～境～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "揪～境～",
			},
		},
		{ // 5
		},
		{ // 6
		},
		{ // 7
			{
				Utf8:    "要怎樣子才可以賺 p 幣呢？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "要怎樣子才可以賺 p 幣呢？～",
			},
		},
		{ // 6
		},
		{ // 7
			{
				Utf8:    "是不是要等待 1 min 呢～？",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "是不是要等待 1 min 呢～？",
			},
		},
		{ // 6
		},
		{ // 6
			{
				Utf8:    "p 幣難賺啊～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "p 幣難賺啊～",
			},
		},
		{ // 6
		},
		{ // 6
			{
				Utf8:    "是不是呢？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "是不是呢？～",
			},
		},
		{ // 6
		},
		{ // 6
			{
				Utf8:    "我只是想要測試一下 p 幣是怎樣子被存在 .DIR 裡而已 XD.",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "我只是想要測試一下 p 幣是怎樣子被存在 .DIR 裡而已 XD.",
			},
		},
		{ // 6
		},
		{ // 6
			{
				Utf8:    "然後呢？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "然後呢？～",
			},
		},
	}

	testUtf8FirstComments3 = nil
	testUtf8FullFirstComments3 = nil
	testUtf8TheRestComments3 = nil
}
