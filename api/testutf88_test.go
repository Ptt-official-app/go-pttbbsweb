package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

var (
	testUtf8Filename8            = "M.1624846335.A.4BB.utf8"
	testUtf8ContentAll8          []byte
	testUtf8Content8             []byte
	testUtf8Signature8           []byte
	testUtf8Comment8             []byte
	testUtf8FirstCommentsDBCS8   []byte
	testUtf8TheRestCommentsDBCS8 []byte
	testUtf8Content8Utf8         [][]*types.Rune

	testUtf8FirstComments8     []*schema.Comment
	testUtf8FullFirstComments8 []*schema.Comment
	testUtf8TheRestComments8   []*schema.Comment
)

func initTestUtf88() {
	testUtf8ContentAll8, testUtf8Content8, testUtf8Signature8, testUtf8Comment8, testUtf8FirstCommentsDBCS8, testUtf8TheRestCommentsDBCS8 = loadTest(testUtf8Filename8)

	testUtf8Content8Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:    "作者: peter50505 (大虱哥) 看板: CMS_99_S3G",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: peter50505 (大虱哥) 看板: CMS_99_S3G",
			},
		},
		{ // 1
			{
				Utf8:    "標題: [轉錄] 特殊色彩效果、控制碼",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [轉錄] 特殊色彩效果、控制碼",
			},
		},
		{ // 2
			{
				Utf8:    "時間: Sun Oct  3 14:42:11 2010",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sun Oct  3 14:42:11 2010",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "※ [本文轉錄自 CMU_CM45 看板 #1CPU2Nrt ]",
				DBCSStr: "※ [本文轉錄自 CMU_CM45 看板 #1CPU2Nrt ]",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 6
			{
				Utf8:    "作者: doggyhu (笨狗) 看板: CMU_CM45",
				DBCSStr: "作者: doggyhu (笨狗) 看板: CMU_CM45",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 7
			{
				Utf8:    "標題: [轉錄][轉錄] 特殊色彩效果、控制碼",
				DBCSStr: "標題: [轉錄][轉錄] 特殊色彩效果、控制碼",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 8
			{
				Utf8:    "時間: Sat Aug 14 08:17:26 2010",
				DBCSStr: "時間: Sat Aug 14 08:17:26 2010",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 9
		},
		{ // 10
			{
				Utf8:    "※ [本文轉錄自 KMSH_99xin 看板 #1CMeh78J ]",
				DBCSStr: "※ [本文轉錄自 KMSH_99xin 看板 #1CMeh78J ]",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 11

		},
		{ // 12
			{
				Utf8:    "作者: brianzzy (王子維)",
				DBCSStr: "作者: brianzzy (王子維)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 13
			{
				Utf8:    "標題: [轉錄][轉錄] 特殊色彩效果、控制碼",
				DBCSStr: "標題: [轉錄][轉錄] 特殊色彩效果、控制碼",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 14
			{
				Utf8:    "時間: Thu Aug  5 17:56:22 2010",
				DBCSStr: "時間: Thu Aug  5 17:56:22 2010",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 15

		},
		{ // 16
			{
				Utf8:    "這份整理的還不錯",
				DBCSStr: "這份整理的還不錯",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 17

		},
		{ // 18
			{
				Utf8:    "慢看吧~",
				DBCSStr: "慢看吧~",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 19

		},
		{ // 20

		},
		{ // 21
			{
				Utf8:    "※ [本文轉錄自 brianzzy 信箱]",
				DBCSStr: "※ [本文轉錄自 brianzzy 信箱]",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 22

		},
		{ // 23
			{
				Utf8:    "大家一起來ASCII Art入門",
				DBCSStr: "\x1b[1;30m大家一起來ASCII Art入門",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 24
			{
				Utf8:    "一步一步玩，從零開始學吧<(￣︶￣)>",
				DBCSStr: "\x1b[1;30m一步一步玩，從零開始學吧<(￣︶￣)>",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 25
			{
				Utf8:    "以下資料整理自water4567、台大資管鳴蟬小站及我自己",
				DBCSStr: "\x1b[1;30m以下資料整理自water4567、台大資管鳴蟬小站及我自己",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 26

		},
		{ // 27
			{
				Utf8:    " ◆ ASCII核心顏色碼：",
				DBCSStr: "\x1b[1;44m ◆ ASCII核心顏色碼：",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 28
			{
				Utf8:    "一、先從著色開始，在要上色的文字",
				DBCSStr: "一、先從著色開始，在要上色的文字",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "前後",
				DBCSStr: "\x1b[1;33m前後",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "打上",
				DBCSStr: "\x1b[m打上",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[m",
				DBCSStr: "\x1b[1;33m*[m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 29

		},
		{ // 30
			{
				Utf8:    "方法Ａ；",
				DBCSStr: "方法Ａ；",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "Crtl+u ",
				DBCSStr: "\x1b[1;33mCrtl+u ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "(*)",
				DBCSStr: "\x1b[m(*)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    " [",
				DBCSStr: "\x1b[1;33m [",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "(在字母P右邊)",
				DBCSStr: "\x1b[m(在字母P右邊)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "m",
				DBCSStr: "\x1b[1;33mm",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 31
			{
				Utf8:    "方法Ｂ(建議使用)；",
				DBCSStr: "方法Ｂ(建議使用)；",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "Crtl+c",
				DBCSStr: "\x1b[1;33mCrtl+c",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "(*[m三個符號一次滿足)",
				DBCSStr: "\x1b[m(*[m三個符號一次滿足)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 32

		},
		{ // 33
			{
				Utf8:    "二、接著將遊標移至前面那一個*[m，並在m前面打上數字作為指令(下列圖表會解釋)",
				DBCSStr: "二、接著將遊標移至前面那一個*[m，並在m前面打上數字作為指令(下列圖表會解釋)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 34

		},
		{ // 35
			{
				Utf8:    "三、預覽與否的切換按Crtl+v，右下角的顯示",
				DBCSStr: "三、預覽與否的切換按Crtl+v，右下角的顯示",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "║插入│aipr║",
				DBCSStr: "\x1b[30;47m║插入│aipr║",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:    "為一般模式",
				DBCSStr: "\x1b[m為一般模式",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 36
			{
				Utf8:    "    預覽時會變成",
				DBCSStr: "    預覽時會變成",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "║插入│Aipr║",
				DBCSStr: "\x1b[30;47m║插入│Aipr║",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 37

		},
		{ // 38
			{
				Utf8:    " ◆ 特殊性質指令 ",
				DBCSStr: "\x1b[1;44m ◆ 特殊性質指令 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 39

		},
		{ // 40
			{
				Utf8:    "    ┌─────────────────────┐",
				DBCSStr: "    ┌─────────────────────┐",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 41
			{
				Utf8:    "    │",
				DBCSStr: "    │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "指令    功能     原始碼         呈現效果",
				DBCSStr: "\x1b[1;33m指令    功能     原始碼         呈現效果",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "  │",
				DBCSStr: "\x1b[m  │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 42
			{
				Utf8:    "    ├─────────────────────┤",
				DBCSStr: "    ├─────────────────────┤",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 43
			{
				Utf8:    "    │  0    無效果    *[0m文字★*[m   文字★   │ ",
				DBCSStr: "    │  0    無效果    *[0m文字★*[m   文字★   │ ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "// 預設值：無效果",
				DBCSStr: "\x1b[1;30m// 預設值：無效果",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 44
			{
				Utf8:    "    │  1    高亮度    *[1m文字★*[m   ",
				DBCSStr: "    │  1    高亮度    *[1m文字★*[m   ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[1m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 45
			{
				Utf8:    "    │  4     底線     *[4m文字★*[m   文字★   │ ",
				DBCSStr: "    │  4     底線     *[4m文字★*[m   文字★   │ ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "// 有些軟體不支援",
				DBCSStr: "\x1b[1;30m// 有些軟體不支援",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 46
			{
				Utf8:    "    │  5     閃爍     *[5m文字★*[m   ",
				DBCSStr: "    │  5     閃爍     *[5m文字★*[m   ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[5m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 47
			{
				Utf8:    "    │  7     反白     *[7m文字★*[m   ",
				DBCSStr: "    │  7     反白     *[7m文字★*[m   ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[30;47m文字★",

				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 48
			{
				Utf8:    "    │  8     黑字     *[8m文字★*[m   文字★   │",
				DBCSStr: "    │  8     黑字     *[8m文字★*[m   文字★   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 49
			{
				Utf8:    "    └─────────────────────┘",
				DBCSStr: "    └─────────────────────┘",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 50

		},
		{ // 51
			{
				Utf8:    " ◆ 前景顏色指令 ",
				DBCSStr: "\x1b[1;44m ◆ 前景顏色指令 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 52

		},
		{ // 53
			{
				Utf8:    "    ┌─────────────────────┐",
				DBCSStr: "    ┌─────────────────────┐",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 54
			{
				Utf8:    "    │",
				DBCSStr: "    │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "指令  前景顏色   原始碼         呈現效果",
				DBCSStr: "\x1b[1;33m指令  前景顏色   原始碼         呈現效果",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "  │",
				DBCSStr: "\x1b[m  │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 55
			{
				Utf8:    "    ├─────────────────────┤",
				DBCSStr: "    ├─────────────────────┤",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 56
			{
				Utf8:    "    │ 30     黑色     *[30m文字★*[m  ",
				DBCSStr: "    │ 30     黑色     *[30m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[30m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 57
			{
				Utf8:    "    │ 31    紅棕色    *[31m文字★*[m  ",
				DBCSStr: "    │ 31    紅棕色    *[31m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[31m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 58
			{
				Utf8:    "    │ 32    墨綠色    *[32m文字★*[m  ",
				DBCSStr: "    │ 32    墨綠色    *[32m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[32m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 59
			{
				Utf8:    "    │ 33    土黃色    *[33m文字★*[m  ",
				DBCSStr: "    │ 33    土黃色    *[33m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[33m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 60
			{
				Utf8:    "    │ 34     靛色     *[34m文字★*[m  ",
				DBCSStr: "    │ 34     靛色     *[34m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[34m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 61
			{
				Utf8:    "    │ 35    深紫色    *[35m文字★*[m  ",
				DBCSStr: "    │ 35    深紫色    *[35m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[35m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 62
			{
				Utf8:    "    │ 36    暗青色    *[36m文字★*[m  ",
				DBCSStr: "    │ 36    暗青色    *[36m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[36m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 63
			{
				Utf8:    "    │ 37     白色     *[37m文字★*[m  文字★   │ ",
				DBCSStr: "    │ 37     白色     *[37m文字★*[m  文字★   │ ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "// 預設值；白字",
				DBCSStr: "\x1b[1;30m// 預設值；白字",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 64
			{
				Utf8:    "    └─────────────────────┘",
				DBCSStr: "    └─────────────────────┘",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 65

		},
		{ // 66
			{
				Utf8:    " ◆ 背景顏色指令 ",
				DBCSStr: "\x1b[1;44m ◆ 背景顏色指令 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 67

		},
		{ // 68
			{
				Utf8:    "    ┌─────────────────────┐",
				DBCSStr: "    ┌─────────────────────┐",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 69
			{
				Utf8:    "    │",
				DBCSStr: "    │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "指令  背景顏色   原始碼         呈現效果",
				DBCSStr: "\x1b[1;33m指令  背景顏色   原始碼         呈現效果",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "  │",
				DBCSStr: "\x1b[m  │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 70
			{
				Utf8:    "    ├─────────────────────┤",
				DBCSStr: "    ├─────────────────────┤",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 71
			{
				Utf8:    "    │ 40     黑色     *[40m文字★*[m  文字★   │ ",
				DBCSStr: "    │ 40     黑色     *[40m文字★*[m  文字★   │ ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "// 預設值；黑底",
				DBCSStr: "\x1b[1;30m// 預設值；黑底",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 72
			{
				Utf8:    "    │ 41    紅棕色    *[41m文字★*[m  ",
				DBCSStr: "    │ 41    紅棕色    *[41m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[41m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 73
			{
				Utf8:    "    │ 42    翠綠色    *[42m文字★*[m  ",
				DBCSStr: "    │ 42    翠綠色    *[42m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[42m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 74
			{
				Utf8:    "    │ 43    土黃色    *[43m文字★*[m  ",
				DBCSStr: "    │ 43    土黃色    *[43m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[43m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 75
			{
				Utf8:    "    │ 44     藍色     *[44m文字★*[m  ",
				DBCSStr: "    │ 44     藍色     *[44m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[44m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 76
			{
				Utf8:    "    │ 45     紫色     *[45m文字★*[m  ",
				DBCSStr: "    │ 45     紫色     *[45m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[45m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 77
			{
				Utf8:    "    │ 46     青色     *[46m文字★*[m  ",
				DBCSStr: "    │ 46     青色     *[46m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[46m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 78
			{
				Utf8:    "    │ 47     白色     *[47m文字★*[m  ",
				DBCSStr: "    │ 47     白色     *[47m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[47m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:    "   │",
				DBCSStr: "\x1b[m   │",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 79
			{
				Utf8:    "    │ 48    亮紅色    *[48m文字★*[m  ",
				DBCSStr: "    │ 48    亮紅色    *[48m文字★*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "文字★",
				DBCSStr: "\x1b[48m文字★",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "   │ ",
				DBCSStr: "\x1b[m   │ ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "某些軟體不支援",
				DBCSStr: "\x1b[1;30m某些軟體不支援",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 80
			{
				Utf8:    "    └─────────────────────┘",
				DBCSStr: "    └─────────────────────┘",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 81

		},
		{ // 82
			{
				Utf8:    "    上述顏色碼，可以多種並存（但前景、背景至多各選一種）",
				DBCSStr: "    上述顏色碼，可以多種並存（但前景、背景至多各選一種）",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 83
			{
				Utf8:    "    如果要對同一文字下多重指令，指令與指令之間要以",
				DBCSStr: "    如果要對同一文字下多重指令，指令與指令之間要以",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    ";",
				DBCSStr: "\x1b[1;33m;",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "分開",
				DBCSStr: "\x1b[m分開",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 84
			{
				Utf8:    "    例如； *[1;5;35;44m政大社會*[m→",
				DBCSStr: "    例如； *[1;5;35;44m政大社會*[m→",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "政大社會",
				DBCSStr: "\x1b[1;5;35;44m政大社會",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 85

		},
		{ // 86
			{
				Utf8:    "然後是圖塊的拼拼湊湊",
				DBCSStr: "\x1b[1;30m然後是圖塊的拼拼湊湊",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 87

		},
		{ // 88
			{
				Utf8:    " ◆ 全形特殊符號",
				DBCSStr: "\x1b[1;44m ◆ 全形特殊符號",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 89

		},
		{ // 90
			{
				Utf8:    "一、Ctrl+p 出現符號表後，按數字鍵可以切換，對照英文鍵輸入即可",
				DBCSStr: "一、Ctrl+p 出現符號表後，按數字鍵可以切換，對照英文鍵輸入即可",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 91

		},
		{ // 92
			{
				Utf8:    "二、PCman的使用者可以按(工具)→(標點符號輸入)",
				DBCSStr: "二、PCman的使用者可以按(工具)→(標點符號輸入)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 93

		},
		{ // 94
			{
				Utf8:    " ◆ 半形特殊符號",
				DBCSStr: "\x1b[1;44m ◆ 半形特殊符號",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
			},
			{
				Utf8:    "(鍵盤上都有)",
				DBCSStr: "\x1b[m(鍵盤上都有)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 95
			{
				Utf8:    "                                        _",
				DBCSStr: "                                        _",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 96
			{
				Utf8:    "          !  exclamation mark               bar (其實它是上一行的底線)",
				DBCSStr: "          !  exclamation mark               bar (其實它是上一行的底線)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 97
			{
				Utf8:    "          @  at sign                    []  brackets",
				DBCSStr: "          @  at sign                    []  brackets",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 98
			{
				Utf8:    "          #  number sign                {}  braces",
				DBCSStr: "          #  number sign                {}  braces",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 99
			{
				Utf8:    "          $  dollar sign                 :  colon",
				DBCSStr: "          $  dollar sign                 :  colon",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 100
			{
				Utf8:    "          %  percent                     ;  semicolon",
				DBCSStr: "          %  percent                     ;  semicolon",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 101
			{
				Utf8:    "          ^  caret, hat                  /  slash",
				DBCSStr: "          ^  caret, hat                  /  slash",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 102
			{
				Utf8:    "          &  ampersand, and              \\  backslash",
				DBCSStr: "          &  ampersand, and              \\  backslash",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 103
			{
				Utf8:    "          *  asterisk, star              ,  comma",
				DBCSStr: "          *  asterisk, star              ,  comma",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 104
			{
				Utf8:    "         ()  parentheses                 .  period, dot",
				DBCSStr: "         ()  parentheses                 .  period, dot",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 105
			{
				Utf8:    "          +  plus                        ?  question mark",
				DBCSStr: "          +  plus                        ?  question mark",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 106
			{
				Utf8:    "          -  minus, hyphen, dash         '  apostrophe",
				DBCSStr: "          -  minus, hyphen, dash         '  apostrophe",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 107
			{
				Utf8:    "          _  horizontal bar, underscore ''  signle quote",
				DBCSStr: "          _  horizontal bar, underscore ''  signle quote",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 108
			{
				Utf8:    "          =  equals sign                \"\"  double quote",
				DBCSStr: "          =  equals sign                \"\"  double quote",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 109
			{
				Utf8:    "          `  grave accent                >  greater than",
				DBCSStr: "          `  grave accent                >  greater than",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 110
			{
				Utf8:    "          ~  tilde                       <  less than",
				DBCSStr: "          ~  tilde                       <  less than",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 111
			{
				Utf8:    "          |  vertical bar",
				DBCSStr: "          |  vertical bar",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 112

		},
		{ // 113
			{
				Utf8:    "再來是自以為很高級的",
				DBCSStr: "\x1b[1;30m再來是自以為很高級的",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "一",
				DBCSStr: "\x1b[37m\x1b[130m一",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "字",
				DBCSStr: "\x1b[37m\x1b[130m字",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "雙",
				DBCSStr: "\x1b[37m\x1b[130m雙",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "色",
				DBCSStr: "\x1b[37m\x1b[130m色",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 114

		},
		{ // 115
			{
				Utf8:    "一、若想要做出這樣的效果，請先測試一下，",
				DBCSStr: "一、若想要做出這樣的效果，請先測試一下，",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 116
			{
				Utf8:    "    對於某個中文字，你能否將游標移到它的後半段？",
				DBCSStr: "    對於某個中文字，你能否將游標移到它的後半段？",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 117

		},
		{ // 118
			{
				Utf8:    "    §例如： ",
				DBCSStr: "    §例如： ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "字 ",
				DBCSStr: "\x1b[1;33m字 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "// 游標位在文字的右半邊",
				DBCSStr: "\x1b[30m// 游標位在文字的右半邊",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 119
			{
				Utf8:    "             ",
				DBCSStr: "             ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "￣",
				DBCSStr: "\x1b[5;30m\x1b[111;137m￣",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Blink: true, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 120
			{
				Utf8:    "    如果不行的話，請先調整一下你觀看BBS所使用之軟體的相關設定。",
				DBCSStr: "    如果不行的話，請先調整一下你觀看BBS所使用之軟體的相關設定。",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 121

		},
		{ // 122
			{
				Utf8:    "    KKman為例：1. 按 Alt + F1 啟動〔KKman設定〕",
				DBCSStr: "    KKman為例：1. 按 Alt + F1 啟動〔KKman設定〕",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 123
			{
				Utf8:    "               2. 選擇〔全形〕",
				DBCSStr: "               2. 選擇〔全形〕",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 124
			{
				Utf8:    "               3. 關閉〔全形字偵測〕",
				DBCSStr: "               3. 關閉〔全形字偵測〕",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 125
			{
				Utf8:    "    PCman：選項 → 修改目前連線的設定值(c) → 自動偵測中文雙位元組",
				DBCSStr: "    PCman：選項 → 修改目前連線的設定值(c) → 自動偵測中文雙位元組",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 126
			{
				Utf8:    "                                    ",
				DBCSStr: "                                    ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    " (有4個框框，都不要打勾)",
				DBCSStr: "\x1b[1;30m (有4個框框，都不要打勾)",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 127

		},
		{ // 128
			{
				Utf8:    "二、",
				DBCSStr: "二、",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "調整完畢之後，上色方法很簡單：在文字的左右兩邊各下一道顏色碼。",
				DBCSStr: "\x1b[m調整完畢之後，上色方法很簡單：在文字的左右兩邊各下一道顏色碼。",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 129

		},
		{ // 130
			{
				Utf8:    "    也就是說： ",
				DBCSStr: "    也就是說： ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "字  字  字  ",
				DBCSStr: "\x1b[1;33m字  字  字  ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "// 在這三個位置分別下：左半顏色、右半顏色、收尾",
				DBCSStr: "\x1b[30m// 在這三個位置分別下：左半顏色、右半顏色、收尾",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 131
			{
				Utf8:    "               ",
				DBCSStr: "               ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "￣",
				DBCSStr: "\x1b[1;5m\x1b[110;130;140m￣",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "  ",
				DBCSStr: "\x1b[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "￣",
				DBCSStr: "\x1b[5;30m\x1b[111;137m￣",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
			},
			{
				Utf8:    "    ",
				DBCSStr: "    ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
			},
			{
				Utf8:    "￣",
				DBCSStr: "\x1b[110;130;140m￣",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 132
			{
				Utf8:    "    原始碼： *[1;32m",
				DBCSStr: "    原始碼： *[1;32m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "?",
				DBCSStr: "\x1b[1;33m?",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "*[31m",
				DBCSStr: "\x1b[m*[31m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "?",
				DBCSStr: "\x1b[33m?",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Utf8:    "*[m  ",
				DBCSStr: "\x1b[m*[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    " // 打完顏色碼之後出現亂請不要緊張",
				DBCSStr: "\x1b[1;30m // 打完顏色碼之後出現亂請不要緊張",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 133
			{
				Utf8:    "    呈現效果：",
				DBCSStr: "    呈現效果：",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "字",
				DBCSStr: "\x1b[1;32m\x1b[131m字",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "                 ",
				DBCSStr: "\x1b[m                 ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "// 按 Ctrl + v 檢查有沒有成功",
				DBCSStr: "\x1b[1;30m// 按 Ctrl + v 檢查有沒有成功",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 134

		},
		{ // 135
			{
				Utf8:    "    同理，除了中文字以外，對任何全形符號，都可套用單字雙色著色方法。",
				DBCSStr: "    同理，除了中文字以外，對任何全形符號，都可套用單字雙色著色方法。",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 136

		},
		{ // 137
			{
				Utf8:    "   例如： ",
				DBCSStr: "   例如： ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "☆",
				DBCSStr: "\x1b[30m\x1b[111;133m☆",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "  只是把 ☆ 的前半部用 前景色黑色的色碼",
				DBCSStr: "\x1b[m  只是把 ☆ 的前半部用 前景色黑色的色碼",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 138

		},
		{ // 139
			{
				Utf8:    "          ",
				DBCSStr: "          ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "●",
				DBCSStr: "\x1b[31m\x1b[141m●",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_RED},
			},
			{
				Utf8:    "    只是把 ● 的後半部用 後景色暗紅色色碼",
				DBCSStr: "\x1b[m    只是把 ● 的後半部用 後景色暗紅色色碼",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 140

		},
		{ // 141

		},
		{ // 142
			{
				Utf8:    "          ",
				DBCSStr: "          ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[1;32;45m\x1b[133;142m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 143

		},
		{ // 144
			{
				Utf8:    "          ▄  一樣，先分割為 前後",
				DBCSStr: "          ▄  一樣，先分割為 前後",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 145

		},
		{ // 146

		},
		{ // 147
			{
				Utf8:    "          前面前景亮綠  後面前景亮黃",
				DBCSStr: "          前面前景亮綠  後面前景亮黃",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 148

		},
		{ // 149
			{
				Utf8:    "          前面背景暗紫  後面背景暗綠",
				DBCSStr: "          前面背景暗紫  後面背景暗綠",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 150

		},
		{ // 151
			{
				Utf8:    "--",
				DBCSStr: "--",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 152

		},
		{ // 153
			{
				Utf8:    "                ",
				DBCSStr: "                ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "◣",
				DBCSStr: "\x1b[30m\x1b[110;107;100m◣",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ // 154
			{
				Utf8:    "             ",
				DBCSStr: "             ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[31m\x1b[141m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_RED},
			},
			{
				Utf8:    "≡",
				DBCSStr: "\x1b[30;40m\x1b[110;107;100m≡",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    ". .",
				DBCSStr: "\x1b[30;47m. .",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:    "◤",
				DBCSStr: "\x1b[37m\x1b[111;130;140m◤",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "〝",
				DBCSStr: "\x1b[m〝",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 155
			{
				Utf8:    "         ",
				DBCSStr: "         ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "╴╴╴╴",
				DBCSStr: "\x1b[1m╴╴╴╴",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "╭",
				DBCSStr: "\x1b[;30;40m\x1b[147m╭",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
			{
				Utf8:    "◣",
				DBCSStr: "\x1b[m\x1b[130m◣",
				Color0:  types.DefaultColor,
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "╴╴╴╴╴╴╴╴╴",
				DBCSStr: "\x1b[1;37m╴╴╴╴╴╴╴╴╴",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    " ",
				DBCSStr: "\x1b[m ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "i",
				DBCSStr: "\x1b[1mi",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "ν",
				DBCSStr: "\x1b[;30;47mν",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
			{
				Utf8:    "y",
				DBCSStr: "\x1b[1;40my",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "v",
				DBCSStr: "\x1b[37mv",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "i",
				DBCSStr: "\x1b[mi",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    " ╴﹎",
				DBCSStr: "\x1b[1m ╴﹎",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 156
			{
				Utf8:    "                   ╯",
				DBCSStr: "                   ╯",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 157

		},
		{ // 158

		},
		{ // 159

		},
		{ // 160

		},
		{ // 161
			{
				Utf8:    "-----------------------------------------------------------------------------",
				DBCSStr: "-----------------------------------------------------------------------------",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 162

		},
		{ // 163
			{
				Utf8:    "這是補充篇，有走馬燈、閃爍跳動、和控制碼",
				DBCSStr: "\x1b[1m這是補充篇，有走馬燈、閃爍跳動、和控制碼",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "。",
				DBCSStr: "\x1b[m。",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 164
			{
				Utf8:    "              ",
				DBCSStr: "              ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▌",
				DBCSStr: "\x1b[34;45;5m▌",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Blink: true},
			},

			{
				Utf8:    "▌",
				DBCSStr: "\x1b[35;44m▌",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLUE, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLUE, Blink: true},
			},
			{
				Utf8:    "▌",
				DBCSStr: "\x1b[34;45m▌",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Blink: true},
			},
			{
				Utf8:    "  ",
				DBCSStr: "\x1b[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[45;34m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA},
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[5;34;45m\x1b[135;144m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLUE, Blink: true},
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[5;35;44m\x1b[134;145m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLUE, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA, Blink: true},
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[m\x1b[34;45m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_MAGENTA},
			},
			{
				Utf8:    "    hi,\x1b*s,真希望趕快到\x1b*b就送你蛋糕",
				DBCSStr: "\x1b[m    hi,\x1b*s,真希望趕快到\x1b*b就送你蛋糕",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 165

		},
		{ // 166

		},
		{ // 167

		},
		{ // 168
			{
				Utf8:    "        ********************************************",
				DBCSStr: "        ********************************************",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 169
			{
				Utf8:    "        *                                          *",
				DBCSStr: "        *                                          *",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 170
			{
				Utf8:    "        *        ",
				DBCSStr: "        *        ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "[教學] 控制碼使用 (補充篇)",
				DBCSStr: "\x1b[1;37m[教學] 控制碼使用 (補充篇)",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "        *",
				DBCSStr: "\x1b[m        *",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 171
			{
				Utf8:    "        *     ",
				DBCSStr: "        *     ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "http://www.wretch.cc/blog/ffivej",
				DBCSStr: "\x1b[30mhttp://www.wretch.cc/blog/ffivej",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    "     *",
				DBCSStr: "\x1b[37m     *",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 172
			{
				Utf8:    "        *                                  ",
				DBCSStr: "        *                                  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "by ",
				DBCSStr: "\x1b[1;30mby ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "f5j",
				DBCSStr: "\x1b[4mf5j", // XXX ignore "underscore"
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "  *",
				DBCSStr: "\x1b[m  *",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 173
			{
				Utf8:    "        *                                          *",
				DBCSStr: "        *                                          *",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 174
			{
				Utf8:    "        ********************************************",
				DBCSStr: "        ********************************************",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 175

		},
		{ // 176

		},
		{ // 177
			{
				Utf8:    " 色碼 ",
				DBCSStr: "\x1b[1;37;42m 色碼 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 178

		},
		{ // 179
			{
				Utf8:    "補充:",
				DBCSStr: "補充:",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 180

		},
		{ // 181
			{
				Utf8:    "*[1;30m",
				DBCSStr: "\x1b[1;30m*[1;30m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "    *[37m   ",
				DBCSStr: "\x1b[m    *[37m   ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[1;37m",
				DBCSStr: "\x1b[1;37m*[1;37m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 182

		},
		{ // 183
			{
				Utf8:    "KKman特有的 ",
				DBCSStr: "KKman特有的 ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[48m",
				DBCSStr: "\x1b[48m*[48m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "       ",
				DBCSStr: "\x1b[40m       ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[38m",
				DBCSStr: "\x1b[38m*[38m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    " 好像看不太出來...",
				DBCSStr: "\x1b[m 好像看不太出來...",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 184

		},
		{ // 185

		},
		{ // 186
			{
				Utf8:    " 一字雙色 ",
				DBCSStr: "\x1b[1;37;42m 一字雙色 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 187

		},
		{ // 188
			{
				Utf8:    "把全形的選項取消,這樣中文字(或全形符號),",
				DBCSStr: "把全形的選項取消,這樣中文字(或全形符號),",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 189
			{
				Utf8:    "往左退時他會退到字的中間,再上控制碼即可!",
				DBCSStr: "往左退時他會退到字的中間,再上控制碼即可!",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 190

		},
		{ // 191
			{
				Utf8:    "範例: ",
				DBCSStr: "範例: ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[31;46m\x1b[111;133;145m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_CYAN},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
			},
			{
				Utf8:    "    ",
				DBCSStr: "\x1b[m    ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "◥",
				DBCSStr: "\x1b[36;41m\x1b[111;132;143m◥",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_RED},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
			},
			{
				Utf8:    "  ",
				DBCSStr: "\x1b[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "＠",
				DBCSStr: "\x1b[35;42m\x1b[111;133;141m＠",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_GREEN},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
			},
			{
				Utf8:    "  ",
				DBCSStr: "\x1b[m  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[30;43m\x1b[133;140m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    " ",
				DBCSStr: "\x1b[37m ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "▄",
				DBCSStr: "\x1b[30;43m\x1b[133m▄",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_YELLOW},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 192

		},
		{ // 193
			{
				Utf8:    "註解: 把全形字當半形字用!",
				DBCSStr: "註解: 把全形字當半形字用!",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 194

		},
		{ // 195

		},
		{ // 196
			{
				Utf8:    " 走馬燈 ",
				DBCSStr: "\x1b[1;37;42m 走馬燈 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 197

		},
		{ // 198
			{
				Utf8:    "範例: ",
				DBCSStr: "範例: ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[5;31;47m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[37;41m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[31;47m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[37;41m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 199

		},
		{ // 200
			{
				Utf8:    "原碼: *[5;31;47m█*[37;41m█*[31;47m█*[37;41m█*[m",
				DBCSStr: "原碼: *[5;31;47m█*[37;41m█*[31;47m█*[37;41m█*[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 201

		},
		{ // 202
			{
				Utf8:    "註解:前景 + 背景 + 閃爍 的應用!",
				DBCSStr: "註解:前景 + 背景 + 閃爍 的應用!",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 203

		},
		{ // 204

		},
		{ // 205
			{
				Utf8:    " 閃爍跳動 ",
				DBCSStr: "\x1b[1;37;42m 閃爍跳動 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 206

		},
		{ // 207
			{
				Utf8:    "範例:",
				DBCSStr: "範例:",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 208
			{
				Utf8:    "    █      █      ",
				DBCSStr: "\x1b[5;33m    █      █      ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 209
			{
				Utf8:    "    █",
				DBCSStr: "\x1b[5;33m    █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[30;43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "    █",
				DBCSStr: "\x1b[33;40m    █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[30;43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 210
			{
				Utf8:    "    █",
				DBCSStr: "\x1b[5;33m    █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "  ",
				DBCSStr: "\x1b[43m  ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "███",
				DBCSStr: "\x1b[40m███",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[30;43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 211
			{
				Utf8:    "    █",
				DBCSStr: "\x1b[5;33m    █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "███  █",
				DBCSStr: "\x1b[30;43m███  █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 212
			{
				Utf8:    "    █",
				DBCSStr: "\x1b[5;33m    █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[30;43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "    █",
				DBCSStr: "\x1b[33;40m    █",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[30;43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 213
			{
				Utf8:    "      ",
				DBCSStr: "\x1b[5;33m      ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[30;43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "      ",
				DBCSStr: "\x1b[40m      ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Blink: true},
			},
			{
				Utf8:    "█",
				DBCSStr: "\x1b[43m█",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_YELLOW, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 214
			{
				Utf8:    "原碼:",
				DBCSStr: "原碼:",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 215
			{
				Utf8:    "*[5;33m    █      █      *[m",
				DBCSStr: "*[5;33m    █      █      *[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 216
			{
				Utf8:    "*[5;33m    █*[30;43m█*[33;40m    █*[30;43m█*[m",
				DBCSStr: "*[5;33m    █*[30;43m█*[33;40m    █*[30;43m█*[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 217
			{
				Utf8:    "*[5;33m    █*[43m  *[40m███*[30;43m█*[m",
				DBCSStr: "*[5;33m    █*[43m  *[40m███*[30;43m█*[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 218
			{
				Utf8:    "*[5;33m    █*[30;43m███  █*[m",
				DBCSStr: "*[5;33m    █*[30;43m███  █*[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 219
			{
				Utf8:    "*[5;33m    █*[30;43m█*[33;40m    █*[30;43m█*[m",
				DBCSStr: "*[5;33m    █*[30;43m█*[33;40m    █*[30;43m█*[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 220
			{
				Utf8:    "*[5;33m      *[30;43m█*[40m      *[43m█*[m",
				DBCSStr: "*[5;33m      *[30;43m█*[40m      *[43m█*[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 221

		},
		{ // 222
			{
				Utf8:    "註解:走馬燈的進階應用!",
				DBCSStr: "註解:走馬燈的進階應用!",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 223

		},
		{ // 224

		},
		{ // 225
			{
				Utf8:    " 網址底線消除 ",
				DBCSStr: "\x1b[1;37;42m 網址底線消除 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 226

		},
		{ // 227
			{
				Utf8:    "範例: ",
				DBCSStr: "範例: ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "http://",
				DBCSStr: "\x1b[1;32mhttp://",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "0rz.net/b41dH",
				DBCSStr: "\x1b[;5;30m0rz.net/b41dH",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true, Blink: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 228

		},
		{ // 229
			{
				Utf8:    "註解:前景 + 閃爍 應用",
				DBCSStr: "註解:前景 + 閃爍 應用",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 230
			{
				Utf8:    "PS:網址底線消除,好像沒看過有人使用...",
				DBCSStr: "PS:網址底線消除,好像沒看過有人使用...",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 231

		},
		{ // 232

		},
		{ // 233
			{
				Utf8:    " 控制碼調整與簡化 ",
				DBCSStr: "\x1b[1;37;42m 控制碼調整與簡化 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 234

		},
		{ // 235
			{
				Utf8:    "補充:",
				DBCSStr: "補充:",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 236
			{
				Utf8:    "  (1)   ",
				DBCSStr: "  (1)   ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[0;31;42m",
				DBCSStr: "\x1b[1;37m*[0;31;42m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},

			{
				Utf8:    "  等於  ",
				DBCSStr: "\x1b[m  等於  ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},

			{
				Utf8:    "*[;31;42m",
				DBCSStr: "\x1b[1;37m*[;31;42m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},

			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 237
			{
				Utf8:    "        若該控制碼位於該行開端  等於 ",
				DBCSStr: "        若該控制碼位於該行開端  等於 ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[31;42m",
				DBCSStr: "\x1b[1;37m*[31;42m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 238

		},
		{ // 239
			{
				Utf8:    "  (2)   欲還原成",
				DBCSStr: "  (2)   欲還原成",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "*[37m",
				DBCSStr: "\x1b[1;37m*[37m",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    " , 只需 *[m  或  *[;4Xm   (X=1~7)",
				DBCSStr: "\x1b[m , 只需 *[m  或  *[;4Xm   (X=1~7)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 240
			{
				Utf8:    "        且可於 4X 之後再加控制碼, 如: *[;4X;4;5m",
				DBCSStr: "        且可於 4X 之後再加控制碼, 如: *[;4X;4;5m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 241
			{
				Utf8:    "        因為 37 是預設色碼,不須特別補上!",
				DBCSStr: "        因為 37 是預設色碼,不須特別補上!",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 242
			{
				Utf8:    "        底色 40 亦同!",
				DBCSStr: "        底色 40 亦同!",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 243

		},
		{ // 244

		},
		{ // 245
			{
				Utf8:    " 其他控制碼 ",
				DBCSStr: "\x1b[1;37;42m 其他控制碼 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 246

		},
		{ // 247
			{
				Utf8:    " **s 顯示使用者的 id       \x1b*s",
				DBCSStr: " **s 顯示使用者的 id       \x1b*s",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 248
			{
				Utf8:    " **l 顯示使用者的上站次數  \x1b*l",
				DBCSStr: " **l 顯示使用者的上站次數  \x1b*l",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 249
			{
				Utf8:    " **b 顯示使用者的生日      \x1b*b",
				DBCSStr: " **b 顯示使用者的生日      \x1b*b",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 250
			{
				Utf8:    " **u 顯示現在的上站人數    \x1b*u",
				DBCSStr: " **u 顯示現在的上站人數    \x1b*u",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 251
			{
				Utf8:    " **p 顯示使用者的文章數    \x1b*p",
				DBCSStr: " **p 顯示使用者的文章數    \x1b*p",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 252
			{
				Utf8:    " **t 顯示現在的時間        \x1b*t",
				DBCSStr: " **t 顯示現在的時間        \x1b*t",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 253
			{
				Utf8:    " **n 顯示使用者的暱稱      \x1b*n",
				DBCSStr: " **n 顯示使用者的暱稱      \x1b*n",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 254

		},
		{ // 255
			{
				Utf8:    "第一個 * 是控制碼,",
				DBCSStr: "第一個 * 是控制碼,",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 256
			{
				Utf8:    "第二個 * 是shift+8 (或數字區9的上面)",
				DBCSStr: "第二個 * 是shift+8 (或數字區9的上面)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 257

		},
		{ // 258
			{
				Utf8:    "但有些BBS站,不支援此功能...",
				DBCSStr: "但有些BBS站,不支援此功能...",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 259

		},
		{ // 260

		},
		{ // 261
			{
				Utf8:    "--",
				DBCSStr: "--",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 262
			{
				Utf8:    "※ 發信站: 批踢踢實業坊(ptt.cc)",
				DBCSStr: "※ 發信站: 批踢踢實業坊(ptt.cc)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 263
			{
				Utf8:    "◆ From: 220.142.13.162",
				DBCSStr: "◆ From: 220.142.13.162",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 264
			{
				Utf8:    "※ 編輯: brianzzy        來自: 220.142.13.162       (08/05 17:57)",
				DBCSStr: "※ 編輯: brianzzy        來自: 220.142.13.162       (08/05 17:57)",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 265
			{
				Utf8:    "推 ",
				DBCSStr: "\x1b[1;37m推 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "NudeLotus   ",
				DBCSStr: "\x1b[33mNudeLotus   ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    ":頭推學長足感心=D                                  ",
				DBCSStr: "\x1b[m\x1b[33m:頭推學長足感心=D                                  ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    " 08/05 21:22",
				DBCSStr: "\x1b[m 08/05 21:22",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 266
			{
				Utf8:    "推 ",
				DBCSStr: "\x1b[1;37m推 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "doggyhu     ",
				DBCSStr: "\x1b[33mdoggyhu     ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    ":好文推                                            ",
				DBCSStr: "\x1b[m\x1b[33m:好文推                                            ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    " 08/10 12:22",
				DBCSStr: "\x1b[m 08/10 12:22",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 267

		},
		{ // 268
			{
				Utf8:    "--",
				DBCSStr: "--",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 269

		},
		{ // 270

		},
		{ // 271

		},
		{ // 272
			{
				Utf8:    " ",
				DBCSStr: " ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "Don't just think that this is breaking or that is others.....Just think this",
				DBCSStr: "\x1b[31mDon't just think that this is breaking or that is others.....Just think this",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ // 273

		},
		{ // 274
			{
				Utf8:    " is dancing.",
				DBCSStr: " is dancing.",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ // 275

		},
		{ // 276
			{
				Utf8:    "--",
				DBCSStr: "--",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ // 277
			{
				Utf8:    "※ 發信站: 批踢踢實業坊(ptt.cc)",
				DBCSStr: "※ 發信站: 批踢踢實業坊(ptt.cc)",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ // 278
			{
				Utf8:    "◆ From: 218.175.153.122",
				DBCSStr: "◆ From: 218.175.153.122",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ // 279
			{
				Utf8:    "推 ",
				DBCSStr: "\x1b[1;37m推 ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "peter50505",
				DBCSStr: "\x1b[33mpeter50505",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    ":好文借轉~                                           ",
				DBCSStr: "\x1b[m\x1b[33m:好文借轉~                                           ",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:    " 10/03 14:41",
				DBCSStr: "\x1b[m 10/03 14:41",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
	}

	testUtf8FirstComments8 = []*schema.Comment{
		{
			TheType: ptttype.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("peter50505"),
			MD5:     "Dh2nPIAxSkeOGqwVyrHntw",
			TheDate: "10/03 14:42",
			IP:      "163.27.69.176",
			DBCSStr: "※ 編輯: peter50505      來自: 163.27.69.176        (10/03 14:42)",
		},
		{
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("geoffrey2012"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "好文m起來",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "好文m起來",
					},
				},
			},
			MD5:     "qPcnTLxxvXTFNSynG1eBYg",
			TheDate: "10/03 16:06",
			DBCSStr: "\x1b[1;37m推 \x1b[33mgeoffrey2012\x1b[m\x1b[33m:好文m起來                                         \x1b[m 10/03 16:06",
		},
		{
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("jason610155"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "好多耶!!!哪看的完~~會用實在點喔!!",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "好多耶!!!哪看的完~~會用實在點喔!!",
					},
				},
			},
			MD5:     "-Q0MF0YE3nkRBzVVPFkx_Q",
			TheDate: "10/03 19:00",
			DBCSStr: "\x1b[1;31m→ \x1b[33mjason610155\x1b[m\x1b[33m:好多耶!!!哪看的完~~會用實在點喔!!                  \x1b[m 10/03 19:00",
		},
	}
	testUtf8FullFirstComments8 = []*schema.Comment{}
	testUtf8TheRestComments8 = []*schema.Comment{}
}
