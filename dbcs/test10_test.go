package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename10            = "temp5"
	testContentAll10          []byte
	testContent10             []byte
	testSignature10           []byte
	testComment10             []byte
	testFirstCommentsDBCS10   []byte
	testTheRestCommentsDBCS10 []byte

	testContent10Big5 [][]*types.Rune
	testContent10Utf8 [][]*types.Rune

	testFirstComments10     []*schema.Comment
	testFullFirstComments10 []*schema.Comment
)

func initTest10() {
	testContentAll10, testContent10, testSignature10, testComment10, testFirstCommentsDBCS10, testTheRestCommentsDBCS10 = loadTest(testFilename10)

	testContent10Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: rareone (\xa5\xec\xa5\xac) \xac\xdd\xaaO: CodeJob"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: rareone (\xa5\xec\xa5\xac) \xac\xdd\xaaO: CodeJob\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] web \xaaA\xb0\xc8\xb6}\xb5o\xc5U\xb0\xdd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] web \xaaA\xb0\xc8\xb6}\xb5o\xc5U\xb0\xdd\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Wed Dec  9 16:47:09 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Wed Dec  9 16:47:09 2020\r"),
			},
		},
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //3
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //4
		{ //5
			{

				Big5:   []byte("        \xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        \xa1@"),
			},
			{

				Big5:   []byte("\xb5o\xae\xd7\xa4H"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xb5o\xae\xd7\xa4H"),
			},
			{

				Big5:   []byte("\xa1Gaaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1Gaaaaaaaa\r"),
			},
		},
		{ //6
			{

				Big5:   []byte("       "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       "),
			},
			{

				Big5:   []byte("\xc1p\xb5\xb8\xa4\xe8\xa6\xa1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xc1p\xb5\xb8\xa4\xe8\xa6\xa1"),
			},
			{

				Big5:   []byte("\xa1Gaaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1Gaaaaaaaa\r"),
			},
		},
		{ //7
			{

				Big5:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1GNTU"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1GNTU\r"),
			},
		},
		{ //8
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //8
		{ //9
			{

				Big5:   []byte("        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        "),
			},
			{

				Big5:   []byte("\xa6\xb3\xae\xc4\xae\xc9\xb6\xa1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xa6\xb3\xae\xc4\xae\xc9\xb6\xa1"),
			},
			{

				Big5:   []byte("\xa1G\xb5o\xa4\xe5\xa4\xe9\xb0_\xa6\xdc 2021/01/01"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xb5o\xa4\xe5\xa4\xe9\xb0_\xa6\xdc 2021/01/01\r"),
			},
		},
		{ //10
			{

				Big5:   []byte("        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        "),
			},
			{

				Big5:   []byte("\xb1M\xae\xd7\xbb\xa1\xa9\xfa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xb1M\xae\xd7\xbb\xa1\xa9\xfa"),
			},
			{

				Big5:   []byte("\xa1G\xa7\xda\xacO\xa4@\xa6W\xb6}\xb5o\xa4\xa4 web service \xaa\xba PM\xa1A\xa6]\xac\xb0\xacO\xb2\xc4\xa4@\xa6\xb8\xb0\xb5\xb1M\xae\xd7\xb9\xef\xa9\xf3\xb6}\xb5o\xaa\xba\xb1`\xc3\xd1\xa9\xd2\xaa\xbe\xb2\xa4\xa4\xd6\xa1A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xa7\xda\xacO\xa4@\xa6W\xb6}\xb5o\xa4\xa4 web service \xaa\xba PM\xa1A\xa6]\xac\xb0\xacO\xb2\xc4\xa4@\xa6\xb8\xb0\xb5\xb1M\xae\xd7\xb9\xef\xa9\xf3\xb6}\xb5o\xaa\xba\xb1`\xc3\xd1\xa9\xd2\xaa\xbe\xb2\xa4\xa4\xd6\xa1A\r"),
			},
		},
		{ //11
			{

				Big5:   []byte("\xa6]\xa6\xb9\xbb\xdd\xadn\xa6\xb3\xa4H\xc0\xb0\xa7U\xa7\xda\xc1\xd7\xa7K\xbd\xf2\xb9p\xa5H\xa4\xce\xa7i\xaa\xbe\xa4@\xa8\xc7\xb6}\xb5o\xac\xdb\xc3\xf6\xaa\xba\xad\xec\xabh\xa1A\xba\xde\xb2z\xb6}\xb5o\xae\xc9\xb5{\xa4W\xaa\xba\xab\xd8\xc4\xb3\xa1A\xc1`\xa6\xd3\xa8\xa5\xa4\xa7\xa7\xda\xbb\xdd\xadn\xa6\xb3\xa4H\xb8\xf2\xa7\xda\xb0Q\xbd\xd7\xb1M\xae\xd7\xa8\xd3\xb4\xee\xa4\xd6\xb3n\xc5\xe9\xb6}\xb5o\xb1`\xb5o\xa5\xcd\xbf\xf9\xbb~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6]\xa6\xb9\xbb\xdd\xadn\xa6\xb3\xa4H\xc0\xb0\xa7U\xa7\xda\xc1\xd7\xa7K\xbd\xf2\xb9p\xa5H\xa4\xce\xa7i\xaa\xbe\xa4@\xa8\xc7\xb6}\xb5o\xac\xdb\xc3\xf6\xaa\xba\xad\xec\xabh\xa1A\xba\xde\xb2z\xb6}\xb5o\xae\xc9\xb5{\xa4W\xaa\xba\xab\xd8\xc4\xb3\xa1A\xc1`\xa6\xd3\xa8\xa5\xa4\xa7\xa7\xda\xbb\xdd\xadn\xa6\xb3\xa4H\xb8\xf2\xa7\xda\xb0Q\xbd\xd7\xb1M\xae\xd7\xa8\xd3\xb4\xee\xa4\xd6\xb3n\xc5\xe9\xb6}\xb5o\xb1`\xb5o\xa5\xcd\xbf\xf9\xbb~\r"),
			},
		},
		{ //12
			{

				Big5:   []byte("I don't know what I should know."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("I don't know what I should know.\r"),
			},
		},
		{ //13
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //13
		{ //14
			{

				Big5:   []byte("        \xa1@\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        \xa1@\xa1@"),
			},
			{

				Big5:   []byte("\xb9w\xba\xe2"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xb9w\xba\xe2"),
			},
			{

				Big5:   []byte("\xa1G2000~5000/hr"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G2000~5000/hr\r"),
			},
		},
		{ //15
			{

				Big5:   []byte("      \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbd\xd0\xb1N CV \xb1H\xa8\xec\xa7\xda\xaa\xba\xabH\xbdc\xa1A\xbb\xb7\xba\xdd\xb1\xb5\xae\xd7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("      \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbd\xd0\xb1N CV \xb1H\xa8\xec\xa7\xda\xaa\xba\xabH\xbdc\xa1A\xbb\xb7\xba\xdd\xb1\xb5\xae\xd7\r"),
			},
		},
	}

	testContent10Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: rareone (伊布) 看板: CodeJob",
				Big5:   []byte("\xa7@\xaa\xcc: rareone (\xa5\xec\xa5\xac) \xac\xdd\xaaO: CodeJob"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: rareone (\xa5\xec\xa5\xac) \xac\xdd\xaaO: CodeJob\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: [發案] web 服務開發顧問",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] web \xaaA\xb0\xc8\xb6}\xb5o\xc5U\xb0\xdd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] web \xaaA\xb0\xc8\xb6}\xb5o\xc5U\xb0\xdd\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Wed Dec  9 16:47:09 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Wed Dec  9 16:47:09 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Wed Dec  9 16:47:09 2020\r"),
			},
		},
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //3
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //4
		{ //5
			{
				Utf8:   "        　",
				Big5:   []byte("        \xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        \xa1@"),
			},
			{
				Utf8:   "發案人",
				Big5:   []byte("\xb5o\xae\xd7\xa4H"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xb5o\xae\xd7\xa4H"),
			},
			{
				Utf8:   "：aaaaaaaa",
				Big5:   []byte("\xa1Gaaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1Gaaaaaaaa\r"),
			},
		},
		{ //6
			{
				Utf8:   "       ",
				Big5:   []byte("       "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       "),
			},
			{
				Utf8:   "聯絡方式",
				Big5:   []byte("\xc1p\xb5\xb8\xa4\xe8\xa6\xa1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xc1p\xb5\xb8\xa4\xe8\xa6\xa1"),
			},
			{
				Utf8:   "：aaaaaaaa",
				Big5:   []byte("\xa1Gaaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1Gaaaaaaaa\r"),
			},
		},
		{ //7
			{
				Utf8:   "       所在地區 ：NTU",
				Big5:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1GNTU"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1GNTU\r"),
			},
		},
		{ //8
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //8
		{ //9
			{
				Utf8:   "        ",
				Big5:   []byte("        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        "),
			},
			{
				Utf8:   "有效時間",
				Big5:   []byte("\xa6\xb3\xae\xc4\xae\xc9\xb6\xa1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xa6\xb3\xae\xc4\xae\xc9\xb6\xa1"),
			},
			{
				Utf8:   "：發文日起至 2021/01/01",
				Big5:   []byte("\xa1G\xb5o\xa4\xe5\xa4\xe9\xb0_\xa6\xdc 2021/01/01"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xb5o\xa4\xe5\xa4\xe9\xb0_\xa6\xdc 2021/01/01\r"),
			},
		},
		{ //10
			{
				Utf8:   "        ",
				Big5:   []byte("        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        "),
			},
			{
				Utf8:   "專案說明",
				Big5:   []byte("\xb1M\xae\xd7\xbb\xa1\xa9\xfa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xb1M\xae\xd7\xbb\xa1\xa9\xfa"),
			},
			{
				Utf8:   "：我是一名開發中 web service 的 PM，因為是第一次做專案對於開發的常識所知略少，",
				Big5:   []byte("\xa1G\xa7\xda\xacO\xa4@\xa6W\xb6}\xb5o\xa4\xa4 web service \xaa\xba PM\xa1A\xa6]\xac\xb0\xacO\xb2\xc4\xa4@\xa6\xb8\xb0\xb5\xb1M\xae\xd7\xb9\xef\xa9\xf3\xb6}\xb5o\xaa\xba\xb1`\xc3\xd1\xa9\xd2\xaa\xbe\xb2\xa4\xa4\xd6\xa1A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xa7\xda\xacO\xa4@\xa6W\xb6}\xb5o\xa4\xa4 web service \xaa\xba PM\xa1A\xa6]\xac\xb0\xacO\xb2\xc4\xa4@\xa6\xb8\xb0\xb5\xb1M\xae\xd7\xb9\xef\xa9\xf3\xb6}\xb5o\xaa\xba\xb1`\xc3\xd1\xa9\xd2\xaa\xbe\xb2\xa4\xa4\xd6\xa1A\r"),
			},
		},
		{ //11
			{
				Utf8:   "因此需要有人幫助我避免踩雷以及告知一些開發相關的原則，管理開發時程上的建議，總而言之我需要有人跟我討論專案來減少軟體開發常發生錯誤",
				Big5:   []byte("\xa6]\xa6\xb9\xbb\xdd\xadn\xa6\xb3\xa4H\xc0\xb0\xa7U\xa7\xda\xc1\xd7\xa7K\xbd\xf2\xb9p\xa5H\xa4\xce\xa7i\xaa\xbe\xa4@\xa8\xc7\xb6}\xb5o\xac\xdb\xc3\xf6\xaa\xba\xad\xec\xabh\xa1A\xba\xde\xb2z\xb6}\xb5o\xae\xc9\xb5{\xa4W\xaa\xba\xab\xd8\xc4\xb3\xa1A\xc1`\xa6\xd3\xa8\xa5\xa4\xa7\xa7\xda\xbb\xdd\xadn\xa6\xb3\xa4H\xb8\xf2\xa7\xda\xb0Q\xbd\xd7\xb1M\xae\xd7\xa8\xd3\xb4\xee\xa4\xd6\xb3n\xc5\xe9\xb6}\xb5o\xb1`\xb5o\xa5\xcd\xbf\xf9\xbb~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6]\xa6\xb9\xbb\xdd\xadn\xa6\xb3\xa4H\xc0\xb0\xa7U\xa7\xda\xc1\xd7\xa7K\xbd\xf2\xb9p\xa5H\xa4\xce\xa7i\xaa\xbe\xa4@\xa8\xc7\xb6}\xb5o\xac\xdb\xc3\xf6\xaa\xba\xad\xec\xabh\xa1A\xba\xde\xb2z\xb6}\xb5o\xae\xc9\xb5{\xa4W\xaa\xba\xab\xd8\xc4\xb3\xa1A\xc1`\xa6\xd3\xa8\xa5\xa4\xa7\xa7\xda\xbb\xdd\xadn\xa6\xb3\xa4H\xb8\xf2\xa7\xda\xb0Q\xbd\xd7\xb1M\xae\xd7\xa8\xd3\xb4\xee\xa4\xd6\xb3n\xc5\xe9\xb6}\xb5o\xb1`\xb5o\xa5\xcd\xbf\xf9\xbb~\r"),
			},
		},
		{ //12
			{
				Utf8:   "I don't know what I should know.",
				Big5:   []byte("I don't know what I should know."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("I don't know what I should know.\r"),
			},
		},
		{ //13
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		}, //13
		{ //14
			{
				Utf8:   "        　　",
				Big5:   []byte("        \xa1@\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        \xa1@\xa1@"),
			},
			{
				Utf8:   "預算",
				Big5:   []byte("\xb9w\xba\xe2"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xb9w\xba\xe2"),
			},
			{
				Utf8:   "：2000~5000/hr",
				Big5:   []byte("\xa1G2000~5000/hr"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G2000~5000/hr\r"),
			},
		},
		{ //15
			{
				Utf8:   "      接案者要求：請將 CV 寄到我的信箱，遠端接案",
				Big5:   []byte("      \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbd\xd0\xb1N CV \xb1H\xa8\xec\xa7\xda\xaa\xba\xabH\xbdc\xa1A\xbb\xb7\xba\xdd\xb1\xb5\xae\xd7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("      \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbd\xd0\xb1N CV \xb1H\xa8\xec\xa7\xda\xaa\xba\xabH\xbdc\xa1A\xbb\xb7\xba\xdd\xb1\xb5\xae\xd7\r"),
			},
		},
	}

	testFirstComments10 = []*schema.Comment{
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("rareone"),
			Content: nil,
			MD5:     "WH30ZZA2DpAv_WfoE-hcBA",
			TheDate: "12/09/2020 16:53:01",
			IP:      "140.112.217.57",
			Host:    "臺灣",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: rareone (140.112.217.57 \xbbO\xc6W), 12/09/2020 16:53:01\r"),

			CreateTime:         1607503981000000000,
			InferredCreateTime: 1607503981000000000,
			SortTime:           1607503981000000000,
			CommentID:          "Fk8AWpKIAgA:WH30ZZA2DpAv_WfoE-hcBA",
		},
		{
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("rareone"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Closed 徵到了",
						Big5:   []byte("Closed \xbcx\xa8\xec\xa4F                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Closed \xbcx\xa8\xec\xa4F                                         "),
					},
				},
			},
			MD5:     "QTzPa_CHE5xLUEUF_A0aZw",
			TheDate: "12/11 02:15",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mrareone\x1b[m\x1b[33m: Closed \xbcx\xa8\xec\xa4F                                         \x1b[m 12/11 02:15\r"),
		},
	}

	testFullFirstComments10 = []*schema.Comment{
		{
			BBoardID:  "test",
			ArticleID: "test10",
			TheType:   types.COMMENT_TYPE_EDIT,
			Owner:     bbs.UUserID("rareone"),
			Content:   nil,
			MD5:       "WH30ZZA2DpAv_WfoE-hcBA",
			TheDate:   "12/09/2020 16:53:01",
			IP:        "140.112.217.57",
			Host:      "臺灣",
			DBCS:      []byte("\xa1\xb0 \xbds\xbf\xe8: rareone (140.112.217.57 \xbbO\xc6W), 12/09/2020 16:53:01\r"),

			CreateTime:         1607503981000000000,
			InferredCreateTime: 1607503981000000000,
			SortTime:           1607503981000000000,
			CommentID:          "Fk8AWpKIAgA:WH30ZZA2DpAv_WfoE-hcBA",
		},
		{
			BBoardID:  "test",
			ArticleID: "test10",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("rareone"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Closed 徵到了",
						Big5:   []byte("Closed \xbcx\xa8\xec\xa4F                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Closed \xbcx\xa8\xec\xa4F                                         "),
					},
				},
			},
			MD5:        "QTzPa_CHE5xLUEUF_A0aZw",
			TheDate:    "12/11 02:15",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mrareone\x1b[m\x1b[33m: Closed \xbcx\xa8\xec\xa4F                                         \x1b[m 12/11 02:15\r"),
			CommentID:  "Fk9tn0_DjgA:QTzPa_CHE5xLUEUF_A0aZw",
			SortTime:   1607624123000000000,
			CreateTime: 1607624123000000000,
		},
	}
}
