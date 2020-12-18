package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename6            = "temp1"
	testContentAll6          []byte
	testContent6             []byte
	testSignature6           []byte
	testComment6             []byte
	testFirstCommentsDBCS6   []byte
	testTheRestCommentsDBCS6 []byte
	testContent6Big5         [][]*types.Rune
	testContent6Utf8         [][]*types.Rune

	testFirstComments6 []*schema.Comment
)

func initTest6() {

	testContentAll6, testContent6, testSignature6, testComment6, testFirstCommentsDBCS6, testTheRestCommentsDBCS6 = loadTest(testFilename6)

	testContent6Big5 = [][]*types.Rune{
		{ //0
			{
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: RedSox"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //1
			{
				Big5:   []byte("\xbc\xd0\xc3D: Fw: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //2
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:46:44 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //3
		{ //4
			{
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb MLB \xac\xdd\xaaO #1UH3bcSy ]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //5
		{ //6
			{
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: MLB"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //7
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //8
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:44:02 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //9
		{ //10
			{
				Big5:   []byte("https://twitter.com/BNightengale/status/1227634955699789824"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //11
		{ //12
			{
				Big5:   []byte("Chris Sale has flu that turned into pneumonia and expected in #Redsox camp"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //13
			{
				Big5:   []byte("Friday."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //14
		{ //15
			{
				Big5:   []byte("Chris Sale \xad\xec\xa5\xbb\xb1o\xa8\xec\xaa\xba\xacy\xb7P\xb2{\xa6b\xa4w\xb8g\xc5\xdc\xa6\xa8\xaa\xcd\xaa\xa2\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //16
		{}, //17
		{ //18
			{
				Big5:   []byte("..\xa9\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //19
		{ //20
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //21
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Big5:   []byte("      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ //22
			{
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa2i"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa1\xdd"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Big5:   []byte("\xa2\xa8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa2\xa9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //23
			{
				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(" \\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
			{
				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("     \xc1\xd9\xacO\xb0Q\xb9\xbd\xa4U\xabB\xa4\xd1     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
		},
		{ //24
			{
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" //\xa1t\\\\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Big5:   []byte("\xa2\xaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //25
			{
				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" \xa2\xa2\xa2\xa4\xa2\xa3"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
			{
				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //26
			{
				Big5:   []byte("  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
			},
			{
				Big5:   []byte("\xa2\xab"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{}, //
		{ //28
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //29
			{
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc), \xa8\xd3\xa6\xdb: 110.50.185.154 (\xbbO\xc6W)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //30
			{
				Big5:   []byte("\xa1\xb0 \xa4\xe5\xb3\xb9\xba\xf4\xa7}: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //31
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(": \xa8\xd3\xa6\xdb\xaaZ\xba~\xaa\xba\xaa\xcd\xaa\xa2\xa1H                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{ //32
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(": \xa6\xad\xa4\xe9\xb1d\xb4_                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{ //33
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("ganhua"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(": \xabO\xad\xab...                                                "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{ //34
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("x6073123"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(": \xabz\xbea \xb3o\xae\xc9\xbe\xf7\xc2I...                                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Big5:   []byte(" 02/13 01:45"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{}, //35
	}

	testContent6Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: currykukuo (陳菊濕汗) 看板: RedSox",
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: RedSox"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //1
			{
				Utf8:   "標題: Fw: [情報] Chris Sale 得肺炎",
				Big5:   []byte("\xbc\xd0\xc3D: Fw: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //2
			{
				Utf8:   "時間: Thu Feb 13 01:46:44 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:46:44 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //3
		{ //4
			{
				Utf8:   "※ [本文轉錄自 MLB 看板 #1UH3bcSy ]",
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb MLB \xac\xdd\xaaO #1UH3bcSy ]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //5
		{ //6
			{
				Utf8:   "作者: currykukuo (陳菊濕汗) 看板: MLB",
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: MLB"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //7
			{
				Utf8:   "標題: [情報] Chris Sale 得肺炎",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //8
			{
				Utf8:   "時間: Thu Feb 13 01:44:02 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:44:02 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //9
		{ //10
			{
				Utf8:   "https://twitter.com/BNightengale/status/1227634955699789824",
				Big5:   []byte("https://twitter.com/BNightengale/status/1227634955699789824"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //11
		{ //12
			{
				Utf8:   "Chris Sale has flu that turned into pneumonia and expected in #Redsox camp",
				Big5:   []byte("Chris Sale has flu that turned into pneumonia and expected in #Redsox camp"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //13
			{
				Utf8:   "Friday.",
				Big5:   []byte("Friday."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //14
		{ //15
			{
				Utf8:   "Chris Sale 原本得到的流感現在已經變成肺炎。",
				Big5:   []byte("Chris Sale \xad\xec\xa5\xbb\xb1o\xa8\xec\xaa\xba\xacy\xb7P\xb2{\xa6b\xa4w\xb8g\xc5\xdc\xa6\xa8\xaa\xcd\xaa\xa2\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //16
		{}, //17
		{ //18
			{
				Utf8:   "..怕",
				Big5:   []byte("..\xa9\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //19
		{ //20
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //21
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "◢",
				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:   "      ",
				Big5:   []byte("      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{ //22
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "█",
				Big5:   []byte("\xa2i"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "≡",
				Big5:   []byte("\xa1\xdd"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "◥",
				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:   "◢",
				Big5:   []byte("\xa2\xa8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "                    ",
				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //23
			{
				Utf8:   "】",
				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "。",
				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   " \\",
				Big5:   []byte(" \\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
			{
				Utf8:   "。",
				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "【",
				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "     還是討厭下雨天     ",
				Big5:   []byte("     \xc1\xd9\xacO\xb0Q\xb9\xbd\xa4U\xabB\xa4\xd1     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
		},
		{ //24
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " //﹀\\\\",
				Big5:   []byte(" //\xa1t\\\\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:   "◥",
				Big5:   []byte("\xa2\xaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "                    ",
				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
			},
			{
				Utf8:   "◤",
				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //25
			{
				Utf8:   "】",
				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " ╰═╯",
				Big5:   []byte(" \xa2\xa2\xa2\xa4\xa2\xa3"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
			},
			{
				Utf8:   "【",
				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "    ",
				Big5:   []byte("    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "◤",
				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //26
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "◢",
				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
			},
			{
				Utf8:   "◥",
				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
			},
			{
				Utf8:   "◤",
				Big5:   []byte("\xa2\xab"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
			},
		},
		{}, //
		{ //28
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //29
			{
				Utf8:   "※ 發信站: 批踢踢實業坊(ptt.cc), 來自: 110.50.185.154 (臺灣)",
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc), \xa8\xd3\xa6\xdb: 110.50.185.154 (\xbbO\xc6W)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //30
			{
				Utf8:   "※ 文章網址: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html",
				Big5:   []byte("\xa1\xb0 \xa4\xe5\xb3\xb9\xba\xf4\xa7}: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //31
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "ZaneTrout",
				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ": 來自武漢的肺炎？                                    ",
				Big5:   []byte(": \xa8\xd3\xa6\xdb\xaaZ\xba~\xaa\xba\xaa\xcd\xaa\xa2\xa1H                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Utf8:   " 02/13 01:44",
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{ //32
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "ZaneTrout",
				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ": 早日康復                                            ",
				Big5:   []byte(": \xa6\xad\xa4\xe9\xb1d\xb4_                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Utf8:   " 02/13 01:44",
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{ //33
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "ganhua",
				Big5:   []byte("ganhua"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ": 保重...                                                ",
				Big5:   []byte(": \xabO\xad\xab...                                                "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Utf8:   " 02/13 01:44",
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{ //34
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "x6073123",
				Big5:   []byte("x6073123"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ": 哇靠 這時機點...                                     ",
				Big5:   []byte(": \xabz\xbea \xb3o\xae\xc9\xbe\xf7\xc2I...                                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
			{
				Utf8:   " 02/13 01:45",
				Big5:   []byte(" 02/13 01:45"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
			},
		},
		{}, //35
	}

	testFirstComments6 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZG9cDt_OAA:6T-ZR97m1lqVecQZAmlMuA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("try107799"),
			CreateTime: types.NanoTS(1266001260000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   ".....,.",
						Big5:   []byte(".....,.                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "6T-ZR97m1lqVecQZAmlMuA",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZG9miRVQAA:ivEDjzbs9GQE0tA2qP9PVQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("hahabis"),
			CreateTime: types.NanoTS(1266001440000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "..............",
						Big5:   []byte("..............                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "ivEDjzbs9GQE0tA2qP9PVQ",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZG-JdcesAA:p5dkEHAzRTSQL4gzP6GvNg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Unleashed"),
			CreateTime: types.NanoTS(1266002040000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "哪招...",
						Big5:   []byte("\xad\xfe\xa9\xdb...                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "p5dkEHAzRTSQL4gzP6GvNg",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZG-92NM2AA:PKyW17rlcp07gdY_wBC4sw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("piconeko"),
			CreateTime: types.NanoTS(1266002940000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不是拉肚子就是肺炎,明年清出這些人Sale,Elvo,JD買Betts",
						Big5:   []byte("\xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "PKyW17rlcp07gdY_wBC4sw",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZG_nwak-AA:py4oWIA-ee0qAUqWJ4WO5A"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("DavFlow"),
			CreateTime: types.NanoTS(1266003660000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "....",
						Big5:   []byte("....                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "py4oWIA-ee0qAUqWJ4WO5A",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZHBJi5ymAA:ag_pH6PvnPs4D5Guezh9kQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("ekpum135"),
			CreateTime: types.NanoTS(1266005340000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "希望沒事.....",
						Big5:   []byte("\xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "ag_pH6PvnPs4D5Guezh9kQ",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZHDfuJuYAA:WB3pkDhBHnLC-eJYMdIeEQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("seekforever"),
			CreateTime: types.NanoTS(1266007920000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "他還在正常訓練，應該還好",
						Big5:   []byte("\xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "WB3pkDhBHnLC-eJYMdIeEQ",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZHM07oWKAA:lxWkMcBKgzfkQU58S75mDw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("tortoise2006"),
			CreateTime: types.NanoTS(1266018180000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "早日康復啊",
						Big5:   []byte("\xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "lxWkMcBKgzfkQU58S75mDw",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZHN3SdhsAA:VsL5xJ9ukmfB5g3kITeCNQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("triff"),
			CreateTime: types.NanoTS(1266019320000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這個時間點得肺炎...",
						Big5:   []byte("\xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "VsL5xJ9ukmfB5g3kITeCNQ",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EZHQJ-MWIAA:9JBlv1b11fR1u5J1QDEHRQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("LBJKO"),
			CreateTime: types.NanoTS(1266021840000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "........=.=lll 傻眼 早日康復阿",
						Big5:   []byte("........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "9JBlv1b11fR1u5J1QDEHRQ",
		},
	}
}
