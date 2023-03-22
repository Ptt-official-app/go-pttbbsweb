package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testUtf8Filename3            = "M.1608386280.A.BC9"
	testUtf8ContentAll3          []byte
	testUtf8Content3             []byte
	testUtf8Signature3           []byte
	testUtf8Comment3             []byte
	testUtf8FirstCommentsDBCS3   []byte
	testUtf8TheRestCommentsDBCS3 []byte
	testUtf8Content3Big5         [][]*types.Rune
	testUtf8Content3Utf8         [][]*types.Rune
	testUtf8Signature3Utf8       [][]*types.Rune

	testUtf8FirstComments3 []*schema.Comment

	testUtf8FullFirstComments3 []*schema.Comment
)

func initTestUtf83() {
	testUtf8ContentAll3, testUtf8Content3, testUtf8Signature3, testUtf8Comment3, testUtf8FirstCommentsDBCS3, testUtf8TheRestCommentsDBCS3 = loadTest(testUtf8Filename3)

	testUtf8Content3Utf8 = [][]*types.Rune{
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
				Utf8:    "標題: [心得] 測試一下特殊字～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [心得] 測試一下特殊字～",
			},
		},
		{
			{
				Utf8:    "時間: Sat Dec 19 21:57:58 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sat Dec 19 21:57:58 2020",
			},
		},
		{},
		{
			{
				Utf8:    "※這樣子有綠色嗎？～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "※這樣子有綠色嗎？～",
			},
		},
		{
			{
				Utf8:    "※ 發信站",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "※ 發信站",
			},
		},
	}
	testUtf8Signature3Utf8 = [][]*types.Rune{
		{},
		{
			{
				Utf8:    "--",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "--",
			},
		},
		{
			{
				Utf8:    "※ 發信站: 批踢踢 docker(pttdocker.test), 來自: 172.22.0.1",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "※ 發信站: 批踢踢 docker(pttdocker.test), 來自: 172.22.0.1",
			},
		},
	}
}
