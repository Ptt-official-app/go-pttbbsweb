package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUtf8Filename1            = "M.1621089154.A.B39.utf8"
	testUtf8ContentAll1          []byte
	testUtf8Content1             []byte
	testUtf8Signature1           []byte
	testUtf8Comment1             []byte
	testUtf8FirstCommentsDBCS1   []byte
	testUtf8TheRestCommentsDBCS1 []byte
	testUtf8Content1Utf8         [][]*types.Rune

	testUtf8FirstComments1     []*schema.Comment
	testUtf8FullFirstComments1 []*schema.Comment
	testUtf8TheRestComments1   []*schema.Comment
)

func initTestUtf81() {
	testUtf8ContentAll1, testUtf8Content1, testUtf8Signature1, testUtf8Comment1, testUtf8FirstCommentsDBCS1, testUtf8TheRestCommentsDBCS1 = loadTest(testUtf8Filename1)

	testUtf8Content1Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:    "作者: PttACT (批踢踢活動部) 看板: SYSOP",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: PttACT (批踢踢活動部) 看板: SYSOP",
			},
		},
		{ // 1
			{
				Utf8:    "標題: Re: [情報] 疫情後的兒童節… (中獎名單)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: Re: [情報] 疫情後的兒童節… (中獎名單)",
			},
		},
		{ // 2
			{
				Utf8:    "時間: Thu May  6 12:50:37 2021",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Thu May  6 12:50:37 2021",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "  感謝每一位支持與關心兒童節活動的板主與鄉民朋友們",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  感謝每一位支持與關心兒童節活動的板主與鄉民朋友們",
			},
		},
		{ // 5
			{
				Utf8:    "  活動部特別準備限量的鶯歌陶瓷吸水杯墊作為抽獎禮物",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  活動部特別準備限量的鶯歌陶瓷吸水杯墊作為抽獎禮物",
			},
		},
		{ // 6
			{
				Utf8:    "  https://i.imgur.com/8629B6y.jpg",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  https://i.imgur.com/8629B6y.jpg",
			},
		},
		{ // 7
			{
				Utf8:    "  杯墊圖案的設計是來自於關愛之家小朋友們所著色的繪畫",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  杯墊圖案的設計是來自於關愛之家小朋友們所著色的繪畫",
			},
		},
		{ // 8
		},
		{ // 9
			{
				Utf8:    "  謝謝大家的分享，讓更多人看見孩子們的需要",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  謝謝大家的分享，讓更多人看見孩子們的需要",
			},
		},
		{ // 10
			{
				Utf8:    "  希望我們的關心與支持，能帶給孩子們溫暖和力量",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  希望我們的關心與支持，能帶給孩子們溫暖和力量",
			},
		},
		{ // 11

		},
		{ // 12
		},
		{ // 13
			{
				Utf8:    " 本次活動的抽獎結果如下： ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCSStr: "\x1b[1;33;41m 本次活動的抽獎結果如下： ",
			},
			{
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "\x1b[m",
			},
		},
		{ // 14

		},
		{ // 15
			{
				Utf8:    "　",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "　",
			},
			{
				Utf8:    "◆ 板主站內轉錄，抽10名：",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCSStr: "\x1b[1;36m◆ 板主站內轉錄，抽10名：",
			},
			{
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "\x1b[m",
			},
		},
		{ // 16
			{
				Utf8:    "　　 https://i.imgur.com/fo6ZVwB.gif",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "　　 https://i.imgur.com/fo6ZVwB.gif",
			},
		},
		{ // 17
			{
				Utf8:    "     ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "     ",
			},
			{
				Utf8:    "(重複轉錄以第一篇文準)",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCSStr: "\x1b[1;30m(重複轉錄以第一篇文準)",
			},
			{
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "\x1b[m",
			},
		},
		{ // 18

		},
		{ // 19
			{
				Utf8:    "　",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "　",
			},
			{
				Utf8:    "◆ FB公開分享貼文，抽10名：",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCSStr: "\x1b[1;36m◆ FB公開分享貼文，抽10名：",
			},
			{
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "\x1b[m",
			},
		},

		{ // 1
			{
				Utf8:    "　　 https://i.imgur.com/IHeeTUo.gif",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "　　 https://i.imgur.com/IHeeTUo.gif",
			},
		},

		{ // 21

		},
		{ // 22
			{
				Utf8:    "  以上獎勵將由活動部協助聯繫與發放",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  以上獎勵將由活動部協助聯繫與發放",
			},
		},
		{ // 23
			{
				Utf8:    "  預計於 5/20 前會逐一完成聯繫",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  預計於 5/20 前會逐一完成聯繫",
			},
		},
		{ // 24
			{
				Utf8:    "  如有問題可向活動部 teemocogs 聯繫",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  如有問題可向活動部 teemocogs 聯繫",
			},
		},
		{ // 25
			{
				Utf8:    "  再次感謝鄉民朋友們的參與！",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "  再次感謝鄉民朋友們的參與！",
			},
		},
		{ // 26

		},
		{ // 27

		},
		{ // 28
			{
				Utf8:    "                                  PTT 活動部",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "                                  PTT 活動部",
			},
		},
	}

	testUtf8FirstComments1 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ericf129"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "辛苦了 ><",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "辛苦了 ><",
					},
				},
			},
			MD5:     "tYigIR7gGaYMy0vC_fnamw",
			IP:      "1.200.29.12",
			TheDate: "05/16 22:57",
			DBCSStr: "\x1b[1;37m推 \x1b[33mericf129\x1b[m\x1b[33m: 辛苦了 ><                             \x1b[m    1.200.29.12 05/16 22:57",
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("zhibb"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "謝謝活動部 :)",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "謝謝活動部 :)",
					},
				},
			},
			IP:      "140.113.0.229",
			TheDate: "05/17 09:49",
			MD5:     "Ob8O34Rg8ePDDzcGTBVlHQ",
			DBCSStr: "\x1b[1;37m推 \x1b[33mzhibb\x1b[m\x1b[33m: 謝謝活動部 :)                            \x1b[m  140.113.0.229 05/17 09:49",
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Japan2001"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "謝謝活動部",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "謝謝活動部",
					},
				},
			},
			MD5:     "VXS_ljQ_WDumtxTwQ7r4CQ",
			IP:      "180.214.183.155",
			TheDate: "05/18 18:57",
			DBCSStr: "\x1b[1;37m推 \x1b[33mJapan2001\x1b[m\x1b[33m: 謝謝活動部                           \x1b[m180.214.183.155 05/18 18:57",
		},
	}

	testUtf8FullFirstComments1 = []*schema.Comment{
		{ // 0
			BBoardID:   "testUtf8",
			ArticleID:  "testUtf81",
			CommentID:  "Fn9D_yV31kA:7GEtBNWCH8N6QsG9aAGoLA",
			CreateTime: 16131301000000000,
			SortTime:   1621089154001000000,
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("ericf129"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "辛苦了 ><",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "\x1b[1;37m推 \x1b[33mericf129\x1b[m\x1b[33m: 辛苦了 ><                             \x1b[m    1.200.29.12 05/16 22:57",
					},
				},
			},
			MD5:     "7GEtBNWCH8N6QsG9aAGoLA",
			IP:      "1.10.29.12",
			TheDate: "05/06 22:57",
			DBCSStr: "x1b[1;37m推 \x1b[33mericf129\x1b[m\x1b[33m: 辛苦了 ><                             \x1b[m    1.200.29.12 05/16 22:57",
		},
		{ // 1
			BBoardID:   "testUtf8",
			ArticleID:  "testUtf81",
			CommentID:  "Fn9D_yWHGIA:9fJnW9k0qnOtNnani_mlaw",
			CreateTime: 161352140000000000,
			SortTime:   162108915400100000,
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("zhibb"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "謝謝活動部 :)",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "謝謝活動部 :)",
					},
				},
			},
			IP:      "140.113.0.229",
			TheDate: "05/07 09:49",
			MD5:     "9fJnW9k0qnOtNnani_mlaw",
			DBCSStr: "\x1b[1;37m推 \x1b[33mzhibb\x1b[m\x1b[33m: 謝謝活動部 :)                            \x1b[m  140.113.0.229 05/17 09:49",
		},
		{ // 2
			BBoardID:   "testUtf8",
			ArticleID:  "testUtf81",
			CommentID:  "Fn0SK73F2AA:SLDr8nBNFVjFaxS6C9iG-Q",
			CreateTime: 16147141000000000,
			SortTime:   16147141000000000,
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Japan101"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "謝謝活動部",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "謝謝活動部",
					},
				},
			},
			MD5:     "SLDr8nBNFVjFaxS6C9iG-Q",
			IP:      "180.214.183.155",
			TheDate: "05/08 18:57",
			DBCSStr: "\x1b[1;37m推 \x1b[33mJapan2001\x1b[m\x1b[33m: 謝謝活動部                           \x1b[m180.214.183.155 05/18 18:57",
		},
	}

	testUtf8TheRestComments1 = []*schema.Comment{}
}
