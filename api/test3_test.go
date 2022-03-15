package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename3            = "M.1608386280.A.BC9"
	testContentAll3          []byte
	testContent3             []byte
	testSignature3           []byte
	testComment3             []byte
	testFirstCommentsDBCS3   []byte
	testTheRestCommentsDBCS3 []byte
	testContent3Big5         [][]*types.Rune
	testContent3Utf8         [][]*types.Rune
	testSignature3Utf8       [][]*types.Rune

	testFirstComments3 []*schema.Comment

	testFullFirstComments3 []*schema.Comment
)

func initTest3() {
	testContentAll3, testContent3, testSignature3, testComment3, testFirstCommentsDBCS3, testTheRestCommentsDBCS3 = loadTest(testFilename3)

	testContent3Utf8 = [][]*types.Rune{
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
	testSignature3Utf8 = [][]*types.Rune{
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
