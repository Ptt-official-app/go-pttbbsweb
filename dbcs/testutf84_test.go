package dbcs

import (
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

var (
	testUtf8Filename4            = "M.1644505664.A.7C5.utf8"
	testUtf8ContentAll4          []byte
	testUtf8Content4             []byte
	testUtf8Signature4           []byte
	testUtf8Comment4             []byte
	testUtf8FirstCommentsDBCS4   []byte
	testUtf8TheRestCommentsDBCS4 []byte
	testUtf8Content4Utf8         [][]*types.Rune

	testUtf8FirstComments4     []*schema.Comment
	testUtf8FullFirstComments4 []*schema.Comment
	testUtf8TheRestComments4   []*schema.Comment
)

func initTestUtf84() {
	testUtf8ContentAll4, testUtf8Content4, testUtf8Signature4, testUtf8Comment4, testUtf8FirstCommentsDBCS4, testUtf8TheRestCommentsDBCS4 = loadTest(testUtf8Filename4)

	testUtf8Content4Utf8 = [][]*types.Rune{
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
				Utf8:    "標題: [心得] 測試一下特殊字～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [心得] 測試一下特殊字～",
			},
		},
		{ // 4
			{
				Utf8:    "時間: Sat Dec 19 21:57:58 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sat Dec 19 21:57:58 2020",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "※這樣子有綠色嗎？～",
				DBCSStr: "※這樣子有綠色嗎？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
			{
				Utf8:    "※ 發信站",
				DBCSStr: "※ 發信站",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
	}
}
