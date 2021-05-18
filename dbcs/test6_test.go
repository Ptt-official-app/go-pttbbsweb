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

	testFirstComments6     []*schema.Comment
	testFullFirstComments6 []*schema.Comment
)

func initTest6() {

	testContentAll6, testContent6, testSignature6, testComment6, testFirstCommentsDBCS6, testTheRestCommentsDBCS6 = loadTest(testFilename6)

	testContent6Big5 = [][]*types.Rune{
		{ //0
			{
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: RedSox"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: RedSox\r"),
			},
		},
		{ //1
			{
				Big5:   []byte("\xbc\xd0\xc3D: Fw: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: Fw: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:46:44 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:46:44 2020\r"),
			},
		},
		{ //3
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //4
			{

				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb MLB \xac\xdd\xaaO #1UH3bcSy ]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb MLB \xac\xdd\xaaO #1UH3bcSy ]\r"),
			},
		},
		{ //5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //6
			{

				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: MLB"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: MLB\r"),
			},
		},
		{ //7
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:44:02 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:44:02 2020\r"),
			},
		},
		{ //9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //10
			{

				Big5:   []byte("https://twitter.com/BNightengale/status/1227634955699789824"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("https://twitter.com/BNightengale/status/1227634955699789824\r"),
			},
		},
		{ //11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //12
			{

				Big5:   []byte("Chris Sale has flu that turned into pneumonia and expected in #Redsox camp"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Chris Sale has flu that turned into pneumonia and expected in #Redsox camp\r"),
			},
		},
		{ //13
			{

				Big5:   []byte("Friday."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Friday.\r"),
			},
		},
		{ //14
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //15
			{

				Big5:   []byte("Chris Sale \xad\xec\xa5\xbb\xb1o\xa8\xec\xaa\xba\xacy\xb7P\xb2{\xa6b\xa4w\xb8g\xc5\xdc\xa6\xa8\xaa\xcd\xaa\xa2\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Chris Sale \xad\xec\xa5\xbb\xb1o\xa8\xec\xaa\xba\xacy\xb7P\xb2{\xa6b\xa4w\xb8g\xc5\xdc\xa6\xa8\xaa\xcd\xaa\xa2\xa1C\r"),
			},
		},
		{ //16
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //17
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //18
			{

				Big5:   []byte("..\xa9\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("..\xa9\xc8\r"),
			},
		},
		{ //19
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //20
			{

				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ //21
			{

				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{

				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m\xa2\xa8"),
			},
			{

				Big5:   []byte("      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[42m      "),
			},
			{

				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[40m\xa2\xa9"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //22
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{

				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42m\xb4"),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte(" "),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m "),
			},
			{

				Big5:   []byte("\xa2i"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2i"),
			},
			{

				Big5:   []byte("\xa1\xdd"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[30;47m\xa1\xdd"),
			},
			{

				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[32m\xa2\xaa"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[42m"),
			},
			{

				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\xa2\x1b[40m\xa9"),
			},
			{

				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("  "),
			},
			{

				Big5:   []byte("\xa2\xa8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xa8"),
			},
			{

				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m                    "),
			},
			{

				Big5:   []byte("\xa2\xa9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xa9\r"),
			},
		},
		{ //23
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{

				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42mj"),
			},
			{

				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m  "),
			},
			{

				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37;40m\xa1C"),
			},
			{

				Big5:   []byte(" \\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				DBCS:   []byte("\x1b[;30;47m \\"),
			},
			{

				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37;40m\xa1C"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[32;42m"),
			},
			{

				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\xa1\x1b[;32;40mi"),
			},
			{

				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("  "),
			},
			{

				Big5:   []byte("     \xc1\xd9\xacO\xb0Q\xb9\xbd\xa4U\xabB\xa4\xd1     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				DBCS:   []byte("\x1b[30;47m     \xc1\xd9\xacO\xb0Q\xb9\xbd\xa4U\xabB\xa4\xd1     "),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //24
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{

				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42m\xb4"),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte(" "),
			},
			{

				Big5:   []byte(" //\xa1t\\\\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[30;47m //\xa1t\\\\"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[32;42m"),
			},
			{

				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\xa1\x1b[40m\xb4"),
			},
			{

				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("  "),
			},
			{

				Big5:   []byte("\xa2\xaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xaa"),
			},
			{

				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m                    "),
			},
			{

				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xab\r"),
			},
		},
		{ //25
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{

				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42mj"),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte(" "),
			},
			{

				Big5:   []byte(" \xa2\xa2\xa2\xa4\xa2\xa3"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				DBCS:   []byte("\x1b[1;31;47m \xa2\xa2\xa2\xa4\xa2\xa3"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[;32;42m"),
			},
			{

				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\xa1\x1b[40mi"),
			},
			{

				Big5:   []byte("    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("    "),
			},
			{

				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xab\r"),
			},
		},
		{ //26
			{

				Big5:   []byte("  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  "),
			},
			{

				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[35m\xa2\xa8"),
			},
			{

				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[34;42m\xa2\xaa"),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
				DBCS:   []byte("\x1b[35;45m "),
			},
			{

				Big5:   []byte("\xa2\xab"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[34;42m\xa2\xab"),
			},
			{

				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[35;40m\xa2\xa9"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //27
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //28
			{

				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ //29
			{

				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc), \xa8\xd3\xa6\xdb: 110.50.185.154 (\xbbO\xc6W)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc), \xa8\xd3\xa6\xdb: 110.50.185.154 (\xbbO\xc6W)\r"),
			},
		},
		{ //30
			{

				Big5:   []byte("\xa1\xb0 \xa4\xe5\xb3\xb9\xba\xf4\xa7}: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xa4\xe5\xb3\xb9\xba\xf4\xa7}: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html\r"),
			},
		},
		{ //31
			{

				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{

				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mZaneTrout"),
			},
			{

				Big5:   []byte(": \xa8\xd3\xa6\xdb\xaaZ\xba~\xaa\xba\xaa\xcd\xaa\xa2\xa1H                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xa8\xd3\xa6\xdb\xaaZ\xba~\xaa\xba\xaa\xcd\xaa\xa2\xa1H                                    "),
			},
			{

				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:44\r"),
			},
		},
		{ //32
			{

				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{

				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mZaneTrout"),
			},
			{

				Big5:   []byte(": \xa6\xad\xa4\xe9\xb1d\xb4_                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xa6\xad\xa4\xe9\xb1d\xb4_                                            "),
			},
			{

				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:44\r"),
			},
		},
		{ //33
			{

				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{

				Big5:   []byte("ganhua"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mganhua"),
			},
			{

				Big5:   []byte(": \xabO\xad\xab...                                                "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xabO\xad\xab...                                                "),
			},
			{

				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:44\r"),
			},
		},
		{ //34
			{

				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{

				Big5:   []byte("x6073123"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mx6073123"),
			},
			{

				Big5:   []byte(": \xabz\xbea \xb3o\xae\xc9\xbe\xf7\xc2I...                                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xabz\xbea \xb3o\xae\xc9\xbe\xf7\xc2I...                                     "),
			},
			{

				Big5:   []byte(" 02/13 01:45"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:45\r"),
			},
		},
	}

	testContent6Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: currykukuo (陳菊濕汗) 看板: RedSox",
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: RedSox"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: RedSox\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: Fw: [情報] Chris Sale 得肺炎",
				Big5:   []byte("\xbc\xd0\xc3D: Fw: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: Fw: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Thu Feb 13 01:46:44 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:46:44 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:46:44 2020\r"),
			},
		},
		{ //3
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //4
			{
				Utf8:   "※ [本文轉錄自 MLB 看板 #1UH3bcSy ]",
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb MLB \xac\xdd\xaaO #1UH3bcSy ]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb MLB \xac\xdd\xaaO #1UH3bcSy ]\r"),
			},
		},
		{ //5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //6
			{
				Utf8:   "作者: currykukuo (陳菊濕汗) 看板: MLB",
				Big5:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: MLB"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: currykukuo (\xb3\xaf\xb5\xe2\xc0\xe3\xa6\xbd) \xac\xdd\xaaO: MLB\r"),
			},
		},
		{ //7
			{
				Utf8:   "標題: [情報] Chris Sale 得肺炎",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] Chris Sale \xb1o\xaa\xcd\xaa\xa2\r"),
			},
		},
		{ //8
			{
				Utf8:   "時間: Thu Feb 13 01:44:02 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:44:02 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Feb 13 01:44:02 2020\r"),
			},
		},
		{ //9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //10
			{
				Utf8:   "https://twitter.com/BNightengale/status/1227634955699789824",
				Big5:   []byte("https://twitter.com/BNightengale/status/1227634955699789824"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("https://twitter.com/BNightengale/status/1227634955699789824\r"),
			},
		},
		{ //11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //12
			{
				Utf8:   "Chris Sale has flu that turned into pneumonia and expected in #Redsox camp",
				Big5:   []byte("Chris Sale has flu that turned into pneumonia and expected in #Redsox camp"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Chris Sale has flu that turned into pneumonia and expected in #Redsox camp\r"),
			},
		},
		{ //13
			{
				Utf8:   "Friday.",
				Big5:   []byte("Friday."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Friday.\r"),
			},
		},
		{ //14
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //15
			{
				Utf8:   "Chris Sale 原本得到的流感現在已經變成肺炎。",
				Big5:   []byte("Chris Sale \xad\xec\xa5\xbb\xb1o\xa8\xec\xaa\xba\xacy\xb7P\xb2{\xa6b\xa4w\xb8g\xc5\xdc\xa6\xa8\xaa\xcd\xaa\xa2\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Chris Sale \xad\xec\xa5\xbb\xb1o\xa8\xec\xaa\xba\xacy\xb7P\xb2{\xa6b\xa4w\xb8g\xc5\xdc\xa6\xa8\xaa\xcd\xaa\xa2\xa1C\r"),
			},
		},
		{ //16
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //17
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //18
			{
				Utf8:   "..怕",
				Big5:   []byte("..\xa9\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("..\xa9\xc8\r"),
			},
		},
		{ //19
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //20
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ //21
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "◢",
				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m\xa2\xa8"),
			},
			{
				Utf8:   "      ",
				Big5:   []byte("      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[42m      "),
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[40m\xa2\xa9"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //22
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42m\xb4"),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte(" "),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m "),
			},
			{
				Utf8:   "█",
				Big5:   []byte("\xa2i"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2i"),
			},
			{
				Utf8:   "≡",
				Big5:   []byte("\xa1\xdd"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[30;47m\xa1\xdd"),
			},
			{
				Utf8:   "◥",
				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[32m\xa2\xaa"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[42m"),
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\xa2\x1b[40m\xa9"),
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("  "),
			},
			{
				Utf8:   "◢",
				Big5:   []byte("\xa2\xa8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xa8"),
			},
			{
				Utf8:   "                    ",
				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m                    "),
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xa9\r"),
			},
		},
		{ //23
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{
				Utf8:   "】",
				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42mj"),
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m  "),
			},
			{
				Utf8:   "。",
				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37;40m\xa1C"),
			},
			{
				Utf8:   " \\",
				Big5:   []byte(" \\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				DBCS:   []byte("\x1b[;30;47m \\"),
			},
			{
				Utf8:   "。",
				Big5:   []byte("\xa1C"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37;40m\xa1C"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[32;42m"),
			},
			{
				Utf8:   "【",
				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\xa1\x1b[;32;40mi"),
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("  "),
			},
			{
				Utf8:   "     還是討厭下雨天     ",
				Big5:   []byte("     \xc1\xd9\xacO\xb0Q\xb9\xbd\xa4U\xabB\xa4\xd1     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				DBCS:   []byte("\x1b[30;47m     \xc1\xd9\xacO\xb0Q\xb9\xbd\xa4U\xabB\xa4\xd1     "),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //24
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42m\xb4"),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte(" "),
			},
			{
				Utf8:   " //﹀\\\\",
				Big5:   []byte(" //\xa1t\\\\"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[30;47m //\xa1t\\\\"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[32;42m"),
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\xa1\x1b[40m\xb4"),
			},
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("  "),
			},
			{
				Utf8:   "◥",
				Big5:   []byte("\xa2\xaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xaa"),
			},
			{
				Utf8:   "                    ",
				Big5:   []byte("                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_WHITE},
				DBCS:   []byte("\x1b[47m                    "),
			},
			{
				Utf8:   "◤",
				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xab\r"),
			},
		},
		{ //25
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[32m"),
			},
			{
				Utf8:   "】",
				Big5:   []byte("\xa1j"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\xa1\x1b[42mj"),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte(" "),
			},
			{
				Utf8:   " ╰═╯",
				Big5:   []byte(" \xa2\xa2\xa2\xa4\xa2\xa3"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_WHITE, Highlight: true},
				DBCS:   []byte("\x1b[1;31;47m \xa2\xa2\xa2\xa4\xa2\xa3"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[;32;42m"),
			},
			{
				Utf8:   "【",
				Big5:   []byte("\xa1i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\xa1\x1b[40mi"),
			},
			{
				Utf8:   "    ",
				Big5:   []byte("    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("    "),
			},
			{
				Utf8:   "◤",
				Big5:   []byte("\xa2\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2\xab\r"),
			},
		},
		{ //26
			{
				Utf8:   "  ",
				Big5:   []byte("  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  "),
			},
			{
				Utf8:   "◢",
				Big5:   []byte("\xa2\xa8"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[35m\xa2\xa8"),
			},
			{
				Utf8:   "◥",
				Big5:   []byte("\xa2\xaa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[34;42m\xa2\xaa"),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_MAGENTA},
				DBCS:   []byte("\x1b[35;45m "),
			},
			{
				Utf8:   "◤",
				Big5:   []byte("\xa2\xab"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_GREEN},
				DBCS:   []byte("\x1b[34;42m\xa2\xab"),
			},
			{
				Utf8:   "◣",
				Big5:   []byte("\xa2\xa9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[35;40m\xa2\xa9"),
			},
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //27
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //28
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ //29
			{
				Utf8:   "※ 發信站: 批踢踢實業坊(ptt.cc), 來自: 110.50.185.154 (臺灣)",
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc), \xa8\xd3\xa6\xdb: 110.50.185.154 (\xbbO\xc6W)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc), \xa8\xd3\xa6\xdb: 110.50.185.154 (\xbbO\xc6W)\r"),
			},
		},
		{ //30
			{
				Utf8:   "※ 文章網址: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html",
				Big5:   []byte("\xa1\xb0 \xa4\xe5\xb3\xb9\xba\xf4\xa7}: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xa4\xe5\xb3\xb9\xba\xf4\xa7}: https://www.ptt.cc/bbs/MLB/M.1581529446.A.73C.html\r"),
			},
		},
		{ //31
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "ZaneTrout",
				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mZaneTrout"),
			},
			{
				Utf8:   ": 來自武漢的肺炎？                                    ",
				Big5:   []byte(": \xa8\xd3\xa6\xdb\xaaZ\xba~\xaa\xba\xaa\xcd\xaa\xa2\xa1H                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xa8\xd3\xa6\xdb\xaaZ\xba~\xaa\xba\xaa\xcd\xaa\xa2\xa1H                                    "),
			},
			{
				Utf8:   " 02/13 01:44",
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:44\r"),
			},
		},
		{ //32
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "ZaneTrout",
				Big5:   []byte("ZaneTrout"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mZaneTrout"),
			},
			{
				Utf8:   ": 早日康復                                            ",
				Big5:   []byte(": \xa6\xad\xa4\xe9\xb1d\xb4_                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xa6\xad\xa4\xe9\xb1d\xb4_                                            "),
			},
			{
				Utf8:   " 02/13 01:44",
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:44\r"),
			},
		},
		{ //33
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "ganhua",
				Big5:   []byte("ganhua"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mganhua"),
			},
			{
				Utf8:   ": 保重...                                                ",
				Big5:   []byte(": \xabO\xad\xab...                                                "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xabO\xad\xab...                                                "),
			},
			{
				Utf8:   " 02/13 01:44",
				Big5:   []byte(" 02/13 01:44"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:44\r"),
			},
		},
		{ //34
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "x6073123",
				Big5:   []byte("x6073123"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mx6073123"),
			},
			{
				Utf8:   ": 哇靠 這時機點...                                     ",
				Big5:   []byte(": \xabz\xbea \xb3o\xae\xc9\xbe\xf7\xc2I...                                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m: \xabz\xbea \xb3o\xae\xc9\xbe\xf7\xc2I...                                     "),
			},
			{
				Utf8:   " 02/13 01:45",
				Big5:   []byte(" 02/13 01:45"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m 02/13 01:45\r"),
			},
		},
	}

	testFirstComments6 = []*schema.Comment{
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("try107799"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   ".....,.",
						Big5:   []byte(".....,.                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte(".....,.                                             "),
					},
				},
			},
			MD5:     "6T-ZR97m1lqVecQZAmlMuA",
			TheDate: "02/13 03:01",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mtry107799\x1b[m\x1b[33m: .....,.                                             \x1b[m 02/13 03:01\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("hahabis"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "..............",
						Big5:   []byte("..............                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("..............                                        "),
					},
				},
			},
			MD5:     "ivEDjzbs9GQE0tA2qP9PVQ",
			TheDate: "02/13 03:04",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mhahabis\x1b[m\x1b[33m: ..............                                        \x1b[m 02/13 03:04\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Unleashed"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "哪招...",
						Big5:   []byte("\xad\xfe\xa9\xdb...                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xad\xfe\xa9\xdb...                                             "),
					},
				},
			},
			MD5:     "p5dkEHAzRTSQL4gzP6GvNg",
			TheDate: "02/13 03:14",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mUnleashed\x1b[m\x1b[33m: \xad\xfe\xa9\xdb...                                             \x1b[m 02/13 03:14\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("piconeko"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不是拉肚子就是肺炎,明年清出這些人Sale,Elvo,JD買Betts",
						Big5:   []byte("\xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts "),
					},
				},
			},
			MD5:     "PKyW17rlcp07gdY_wBC4sw",
			TheDate: "02/13 03:29",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mpiconeko\x1b[m\x1b[33m: \xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts \x1b[m 02/13 03:29\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("DavFlow"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "....",
						Big5:   []byte("....                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("....                                                  "),
					},
				},
			},
			MD5:     "py4oWIA-ee0qAUqWJ4WO5A",
			TheDate: "02/13 03:41",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDavFlow\x1b[m\x1b[33m: ....                                                  \x1b[m 02/13 03:41\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ekpum135"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "希望沒事.....",
						Big5:   []byte("\xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        "),
					},
				},
			},
			MD5:     "ag_pH6PvnPs4D5Guezh9kQ",
			TheDate: "02/13 04:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mekpum135\x1b[m\x1b[33m: \xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        \x1b[m 02/13 04:09\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("seekforever"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "他還在正常訓練，應該還好",
						Big5:   []byte("\xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          "),
					},
				},
			},
			MD5:     "WB3pkDhBHnLC-eJYMdIeEQ",
			TheDate: "02/13 04:52",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mseekforever\x1b[m\x1b[33m: \xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          \x1b[m 02/13 04:52\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("tortoise2006"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "早日康復啊",
						Big5:   []byte("\xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       "),
					},
				},
			},
			MD5:     "lxWkMcBKgzfkQU58S75mDw",
			TheDate: "02/13 07:43",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mtortoise2006\x1b[m\x1b[33m: \xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       \x1b[m 02/13 07:43\r"),
		},
		{
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("triff"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這個時間點得肺炎...",
						Big5:   []byte("\xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     "),
					},
				},
			},
			MD5:     "VsL5xJ9ukmfB5g3kITeCNQ",
			TheDate: "02/13 08:02",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mtriff\x1b[m\x1b[33m: \xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     \x1b[m 02/13 08:02\r"),
		},
		{
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("LBJKO"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "........=.=lll 傻眼 早日康復阿",
						Big5:   []byte("........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          "),
					},
				},
			},
			MD5:     "9JBlv1b11fR1u5J1QDEHRQ",
			TheDate: "02/13 08:44",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mLBJKO\x1b[m\x1b[33m: ........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          \x1b[m 02/13 08:44\r"),
		},
	}

	testFullFirstComments6 = []*schema.Comment{
		{ //0
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfK82SZ7OAA:6T-ZR97m1lqVecQZAmlMuA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("try107799"),
			CreateTime: types.NanoTS(1581534060000000000),
			SortTime:   types.NanoTS(1581534060000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   ".....,.",
						Big5:   []byte(".....,.                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte(".....,.                                             "),
					},
				},
			},
			MD5:     "6T-ZR97m1lqVecQZAmlMuA",
			TheDate: "02/13 03:01",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mtry107799\x1b[m\x1b[33m: .....,.                                             \x1b[m 02/13 03:01\r"),
		},
		{ //1
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfK9Aw9RQAA:ivEDjzbs9GQE0tA2qP9PVQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("hahabis"),
			CreateTime: types.NanoTS(1581534240000000000),
			SortTime:   types.NanoTS(1581534240000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "..............",
						Big5:   []byte("..............                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("..............                                        "),
					},
				},
			},
			MD5:     "ivEDjzbs9GQE0tA2qP9PVQ",
			TheDate: "02/13 03:04",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mhahabis\x1b[m\x1b[33m: ..............                                        \x1b[m 02/13 03:04\r"),
		},
		{ //2
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfK9jsIasAA:p5dkEHAzRTSQL4gzP6GvNg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Unleashed"),
			CreateTime: types.NanoTS(1581534840000000000),
			SortTime:   types.NanoTS(1581534840000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "哪招...",
						Big5:   []byte("\xad\xfe\xa9\xdb...                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xad\xfe\xa9\xdb...                                             "),
					},
				},
			},
			MD5:     "p5dkEHAzRTSQL4gzP6GvNg",
			TheDate: "02/13 03:14",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mUnleashed\x1b[m\x1b[33m: \xad\xfe\xa9\xdb...                                             \x1b[m 02/13 03:14\r"),
		},
		{ //3
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfK-YE5I2AA:PKyW17rlcp07gdY_wBC4sw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("piconeko"),
			CreateTime: types.NanoTS(1581535740000000000),
			SortTime:   types.NanoTS(1581535740000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不是拉肚子就是肺炎,明年清出這些人Sale,Elvo,JD買Betts",
						Big5:   []byte("\xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts "),
					},
				},
			},
			MD5:     "PKyW17rlcp07gdY_wBC4sw",
			TheDate: "02/13 03:29",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mpiconeko\x1b[m\x1b[33m: \xa4\xa3\xacO\xa9\xd4\xa8{\xa4l\xb4N\xacO\xaa\xcd\xaa\xa2,\xa9\xfa\xa6~\xb2M\xa5X\xb3o\xa8\xc7\xa4HSale,Elvo,JD\xb6RBetts \x1b[m 02/13 03:29\r"),
		},
		{ //4
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfK_B_Gg-AA:py4oWIA-ee0qAUqWJ4WO5A"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("DavFlow"),
			CreateTime: types.NanoTS(1581536460000000000),
			SortTime:   types.NanoTS(1581536460000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "....",
						Big5:   []byte("....                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("....                                                  "),
					},
				},
			},
			MD5:     "py4oWIA-ee0qAUqWJ4WO5A",
			TheDate: "02/13 03:41",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDavFlow\x1b[m\x1b[33m: ....                                                  \x1b[m 02/13 03:41\r"),
		},
		{ //5
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfLAjxlumAA:ag_pH6PvnPs4D5Guezh9kQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("ekpum135"),
			CreateTime: types.NanoTS(1581538140000000000),
			SortTime:   types.NanoTS(1581538140000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "希望沒事.....",
						Big5:   []byte("\xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        "),
					},
				},
			},
			MD5:     "ag_pH6PvnPs4D5Guezh9kQ",
			TheDate: "02/13 04:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mekpum135\x1b[m\x1b[33m: \xa7\xc6\xb1\xe6\xa8S\xa8\xc6.....                                        \x1b[m 02/13 04:09\r"),
		},
		{ //6
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfLC581qYAA:WB3pkDhBHnLC-eJYMdIeEQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("seekforever"),
			CreateTime: types.NanoTS(1581540720000000000),
			SortTime:   types.NanoTS(1581540720000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "他還在正常訓練，應該還好",
						Big5:   []byte("\xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          "),
					},
				},
			},
			MD5:     "WB3pkDhBHnLC-eJYMdIeEQ",
			TheDate: "02/13 04:52",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mseekforever\x1b[m\x1b[33m: \xa5L\xc1\xd9\xa6b\xa5\xbf\xb1`\xb0V\xbdm\xa1A\xc0\xb3\xb8\xd3\xc1\xd9\xa6n                          \x1b[m 02/13 04:52\r"),
		},
		{ //7
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfLMPKUSKAA:lxWkMcBKgzfkQU58S75mDw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("tortoise2006"),
			CreateTime: types.NanoTS(1581550980000000000),
			SortTime:   types.NanoTS(1581550980000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "早日康復啊",
						Big5:   []byte("\xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       "),
					},
				},
			},
			MD5:     "lxWkMcBKgzfkQU58S75mDw",
			TheDate: "02/13 07:43",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mtortoise2006\x1b[m\x1b[33m: \xa6\xad\xa4\xe9\xb1d\xb4_\xb0\xda                                       \x1b[m 02/13 07:43\r"),
		},
		{ //8
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfLNRhJdsAA:VsL5xJ9ukmfB5g3kITeCNQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("triff"),
			CreateTime: types.NanoTS(1581552120000000000),
			SortTime:   types.NanoTS(1581552120000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這個時間點得肺炎...",
						Big5:   []byte("\xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     "),
					},
				},
			},
			MD5:     "VsL5xJ9ukmfB5g3kITeCNQ",
			TheDate: "02/13 08:02",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mtriff\x1b[m\x1b[33m: \xb3o\xad\xd3\xae\xc9\xb6\xa1\xc2I\xb1o\xaa\xcd\xaa\xa2...                                     \x1b[m 02/13 08:02\r"),
		},
		{ //9
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test6"),
			CommentID:  types.CommentID("FfLPkM4SIAA:9JBlv1b11fR1u5J1QDEHRQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("LBJKO"),
			CreateTime: types.NanoTS(1581554640000000000),
			SortTime:   types.NanoTS(1581554640000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "........=.=lll 傻眼 早日康復阿",
						Big5:   []byte("........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          "),
					},
				},
			},
			MD5:     "9JBlv1b11fR1u5J1QDEHRQ",
			TheDate: "02/13 08:44",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mLBJKO\x1b[m\x1b[33m: ........=.=lll \xb6\xcc\xb2\xb4 \xa6\xad\xa4\xe9\xb1d\xb4_\xaa\xfc                          \x1b[m 02/13 08:44\r"),
		},
	}
}
