package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename14            = "temp9"
	testContentAll14          []byte
	testContent14             []byte
	testSignature14           []byte
	testComment14             []byte
	testFirstCommentsDBCS14   []byte
	testTheRestCommentsDBCS14 []byte
	testContent14Big5         [][]*types.Rune
	testContent14Utf8         [][]*types.Rune

	testFirstComments14 []*schema.Comment
)

func initTest14() {
	testContentAll14, testContent14, testSignature14, testComment14, testFirstCommentsDBCS14, testTheRestCommentsDBCS14 = loadTest(testFilename14)

	testContent14Big5 = [][]*types.Rune{
		{ //0
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("2020-21 NBA Regular Season"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[m\x1b[1;44m2020-21 NBA Regular Season"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1m\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //3
			{

				Big5:   []byte(" LA Clippers "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCS:   []byte("\x1b[1;41m LA Clippers "),
			},
			{

				Big5:   []byte(" @ "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m @ "),
			},
			{

				Big5:   []byte(" Dallas Mavericks "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				DBCS:   []byte("\x1b[1;45m Dallas Mavericks "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //4
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
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

				Big5:   []byte("Team  1st   2nd   3rd   4th  Total "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;46mTeam  1st   2nd   3rd   4th  Total "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //7
			{

				Big5:   []byte("LAC    29    22    20    18    89"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("LAC    29    22    20    18    89\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("DAL    24    32    21    28   105"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("DAL    24    32    21    28   105\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //11
			{

				Big5:   []byte("Arena Stats             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;46mArena Stats             "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //12
			{

				Big5:   []byte("Stadium: "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Stadium: "),
			},
			{

				Big5:   []byte("American Airlines Center"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mAmerican Airlines Center"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //13
			{

				Big5:   []byte("         "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("         "),
			},
			{

				Big5:   []byte("Dallas, TX"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mDallas, TX"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //14
			{

				Big5:   []byte("Attendance: "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Attendance: "),
			},
			{

				Big5:   []byte("3975"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m3975"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //15
			{

				Big5:   []byte("Duration:   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Duration:   "),
			},
			{

				Big5:   []byte("2:15"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m2:15"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
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

				Big5:   []byte("Scoring                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;46mScoring                 "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //19
			{

				Big5:   []byte("Lead Changes: "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Lead Changes: "),
			},
			{

				Big5:   []byte("11"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m11"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //20
			{

				Big5:   []byte("Times Tied:   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Times Tied:   "),
			},
			{

				Big5:   []byte("5"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m5"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //21
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //22
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //23
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //24
			{

				Big5:   []byte(" LA Clippers  (26-16)                                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCS:   []byte("\x1b[1;41m LA Clippers  (26-16)                                                          "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //25
			{

				Big5:   []byte("PLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[1;42mPLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //26
			{

				Big5:   []byte("K. Leonard    F 36:55  9-21 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("K. Leonard    F 36:55  9-21 "),
			},
			{

				Big5:   []byte("  1-4   1-4  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  1-4   1-4  "),
			},
			{

				Big5:   []byte("-8  2  5  7  7  3  2  2  0  2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m-8  2  5  7  7  3  2  2  0  2 "),
			},
			{

				Big5:   []byte(" 20  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m 20  "),
			},
			{

				Big5:   []byte("19"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m19\r"),
			},
		},
		{ //27
			{

				Big5:   []byte("M. Morris Sr. F 36:03  5-14 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("M. Morris Sr. F 36:03  5-14 "),
			},
			{

				Big5:   []byte("  1-9   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  1-9   "),
			},
			{

				Big5:   []byte("0-0 -11  1  6  7  2  3  2  0  0  0  11  13"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0 -11  1  6  7  2  3  2  0  0  0  11  13\r"),
			},
		},
		{ //28
			{

				Big5:   []byte("I. Zubac      C 35:02 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("I. Zubac      C 35:02 "),
			},
			{

				Big5:   []byte("  7-9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m  7-9"),
			},
			{

				Big5:   []byte("   0-0   0-0 -15  4  3  7  2  2  3  2  1  1  14  23"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m   0-0   0-0 -15  4  3  7  2  2  3  2  1  1  14  23\r"),
			},
		},
		{ //29
			{

				Big5:   []byte("P. George     G 41:21 10-20 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("P. George     G 41:21 10-20 "),
			},
			{

				Big5:   []byte("  5-8 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m  5-8 "),
			},
			{

				Big5:   []byte("  3-3 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31m  3-3 "),
			},
			{

				Big5:   []byte("-10  0  7  7  5  0  1  4  1  1 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m-10  0  7  7  5  0  1  4  1  1 "),
			},
			{

				Big5:   []byte(" 28  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m 28  "),
			},
			{

				Big5:   []byte("28"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m28\r"),
			},
		},
		{ //30
			{

				Big5:   []byte("R. Jackson    G 19:16 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("R. Jackson    G 19:16 "),
			},
			{

				Big5:   []byte("  2-8   0-3   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  2-8   0-3   "),
			},
			{

				Big5:   []byte("1-2 -13  0  2  2  1  2  2  1  0  1   5   2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m1-2 -13  0  2  2  1  2  2  1  0  1   5   2\r"),
			},
		},
		{ //31
			{

				Big5:   []byte("N. Batum        29:00   1-3   1-2   0-0  -3  0  2  2  2  2  2  1  0  0   3   6"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("N. Batum        29:00   1-3   1-2   0-0  -3  0  2  2  2  2  2  1  0  0   3   6\r"),
			},
		},
		{ //32
			{

				Big5:   []byte("L. Williams     24:24 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("L. Williams     24:24 "),
			},
			{

				Big5:   []byte(" 2-10   1-4   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m 2-10   1-4   "),
			},
			{

				Big5:   []byte("0-0  -5  0  1  1  3  0  0  1  1  2   5   1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0  -5  0  1  1  3  0  0  1  1  2   5   1\r"),
			},
		},
		{ //33
			{

				Big5:   []byte("T. Mann         13:33   1-2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("T. Mann         13:33   1-2 "),
			},
			{

				Big5:   []byte("  0-1   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  0-1   "),
			},
			{

				Big5:   []byte("1-2  -6  1  3  4  0  3  0  0  0  0   3   5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m1-2  -6  1  3  4  0  3  0  0  0  0   3   5\r"),
			},
		},
		{ //34
			{

				Big5:   []byte("A. Coffey       01:32   0-0   0-0   0-0  -3  0  0  0  0  0  0  0  0  0   0   0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("A. Coffey       01:32   0-0   0-0   0-0  -3  0  0  0  0  0  0  0  0  0   0   0\r"),
			},
		},
		{ //35
			{

				Big5:   []byte("D. Oturu        01:27   0-0   0-0   0-0  -3  0  0  0  0  0  0  1  0  0   0 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("D. Oturu        01:27   0-0   0-0   0-0  -3  0  0  0  0  0  0  1  0  0   0 "),
			},
			{

				Big5:   []byte(" -1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m -1"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //36
			{

				Big5:   []byte("M. Kabengele    01:27 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("M. Kabengele    01:27 "),
			},
			{

				Big5:   []byte("  0-1   0-1   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  0-1   0-1   "),
			},
			{

				Big5:   []byte("0-0  -3  0  0  0  0  0  0  0  0  0   0 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0  -3  0  0  0  0  0  0  0  0  0   0 "),
			},
			{

				Big5:   []byte(" -1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m -1"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //37
			{

				Big5:   []byte("L. Kennard       DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mL. Kennard       DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //38
			{

				Big5:   []byte("P. Patterson     DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mP. Patterson     DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //39
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //40
			{

				Big5:   []byte("Total            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[1;42mTotal            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //41
			{

				Big5:   []byte("                 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                 "),
			},
			{

				Big5:   []byte("240  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1m240  "),
			},
			{

				Big5:   []byte("37-88 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m37-88 "),
			},
			{

				Big5:   []byte(" 9-32 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m 9-32 "),
			},
			{

				Big5:   []byte(" 6-11     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m 6-11     "),
			},
			{

				Big5:   []byte(" 8 29 37 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m 8 29 37 "),
			},
			{

				Big5:   []byte("22 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m22 "),
			},
			{

				Big5:   []byte("15 12 12  3  7  89"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m15 12 12  3  7  89\r"),
			},
		},
		{ //42
			{

				Big5:   []byte("                      "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                      "),
			},
			{

				Big5:   []byte("42.0% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m42.0% "),
			},
			{

				Big5:   []byte("28.1% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[32m28.1% "),
			},
			{

				Big5:   []byte("54.5%     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m54.5%     "),
			},
			{

				Big5:   []byte("Team Rebs: 47  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[mTeam Rebs: 47  "),
			},
			{

				Big5:   []byte("Total TO: 13"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32mTotal TO: 13"),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //43
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //44
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //45
			{

				Big5:   []byte(" Dallas Mavericks  (21-18)                                                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				DBCS:   []byte("\x1b[1;45m Dallas Mavericks  (21-18)                                                     "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //46
			{

				Big5:   []byte("PLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				DBCS:   []byte("\x1b[1;43mPLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //47
			{

				Big5:   []byte("Hardaway Jr.  F 36:34  6-12   3-7   0-0 +17  0  6  6  1  1  0  2  0  0  15  14"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Hardaway Jr.  F 36:34  6-12   3-7   0-0 +17  0  6  6  1  1  0  2  0  0  15  14\r"),
			},
		},
		{ //48
			{

				Big5:   []byte("M. Kleber     F 40:54   4-8   2-5 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("M. Kleber     F 40:54   4-8   2-5 "),
			},
			{

				Big5:   []byte("  2-2 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m  2-2 "),
			},
			{

				Big5:   []byte("+24  0  6  6  0  3  0  1  0  0  12  13"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m+24  0  6  6  0  3  0  1  0  0  12  13\r"),
			},
		},
		{ //49
			{

				Big5:   []byte("K. Porzingis  C 34:19  5-14"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[1;44mK. Porzingis  C 34:19  5-14"),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[;37;44m "),
			},
			{

				Big5:   []byte("  1-4   0-1 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  1-4   0-1 "),
			},
			{

				Big5:   []byte(" +9  3 10"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[37m +9  3 10"),
			},
			{

				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[;37;44m "),
			},
			{

				Big5:   []byte("13 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[1;31m13 "),
			},
			{

				Big5:   []byte(" 0  2  0  4  2  1  11  12 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[37m 0  2  0  4  2  1  11  12 "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //50
			{

				Big5:   []byte("J. Richardson G 38:51   5-9   2-4 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("J. Richardson G 38:51   5-9   2-4 "),
			},
			{

				Big5:   []byte("  2-2 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m  2-2 "),
			},
			{

				Big5:   []byte("+16  1  4  5  2  2  3  1  1  0  14  20"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m+16  1  4  5  2  2  3  1  1  0  14  20\r"),
			},
		},
		{ //51
			{

				Big5:   []byte("L. Doncic     G 42:48 16-28  6-11   4-5 +29  1  5  6  9  2  3  4  2  2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("L. Doncic     G 42:48 16-28  6-11   4-5 +29  1  5  6  9  2  3  4  2  2 "),
			},
			{

				Big5:   []byte(" 42  45"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m 42  45"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //52
			{

				Big5:   []byte("J. Brunson      15:44   2-5   1-2   0-0 -11  1  2  3  2  2  1  5  0  0   5   3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("J. Brunson      15:44   2-5   1-2   0-0 -11  1  2  3  2  2  1  5  0  0   5   3\r"),
			},
		},
		{ //53
			{

				Big5:   []byte("W. Iwundu       10:37   0-0   0-0   0-0  -2  0  2  2  0  2  0  1  0  0   0   1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("W. Iwundu       10:37   0-0   0-0   0-0  -2  0  2  2  0  2  0  1  0  0   0   1\r"),
			},
		},
		{ //54
			{

				Big5:   []byte("T. Burke        04:59 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("T. Burke        04:59 "),
			},
			{

				Big5:   []byte("  1-1 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m  1-1 "),
			},
			{

				Big5:   []byte("  0-0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m  0-0"),
			},
			{

				Big5:   []byte("   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m   "),
			},
			{

				Big5:   []byte("0-0  -7  0  0  0  0  0  0  0  0  0   2   2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0  -7  0  0  0  0  0  0  0  0  0   2   2\r"),
			},
		},
		{ //55
			{

				Big5:   []byte("D. Powell       03:19   0-0   0-0   0-0  -5  0  0  0  0  0  0  0  0  0   0   0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("D. Powell       03:19   0-0   0-0   0-0  -5  0  0  0  0  0  0  0  0  0   0   0\r"),
			},
		},
		{ //56
			{

				Big5:   []byte("Cauley-Stein    11:55   1-2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Cauley-Stein    11:55   1-2 "),
			},
			{

				Big5:   []byte("  0-1 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  0-1 "),
			},
			{

				Big5:   []byte("  2-2 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31m  2-2 "),
			},
			{

				Big5:   []byte("+10  0  2  2  0  2  0  0  2  0   4   7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m+10  0  2  2  0  2  0  0  2  0   4   7\r"),
			},
		},
		{ //57
			{

				Big5:   []byte("T. Bey           DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mT. Bey           DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //58
			{

				Big5:   []byte("J. Green         DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mJ. Green         DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //59
			{

				Big5:   []byte("N. Hinton        DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mN. Hinton        DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //60
			{

				Big5:   []byte("B. Marjanovic    DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mB. Marjanovic    DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //61
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //62
			{

				Big5:   []byte("Total            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				DBCS:   []byte("\x1b[1;43mTotal            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //63
			{

				Big5:   []byte("                 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                 "),
			},
			{

				Big5:   []byte("240  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1m240  "),
			},
			{

				Big5:   []byte("40-79 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m40-79 "),
			},
			{

				Big5:   []byte("15-34 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m15-34 "),
			},
			{

				Big5:   []byte("10-12      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m10-12      "),
			},
			{

				Big5:   []byte("6 37 43 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m6 37 43 "),
			},
			{

				Big5:   []byte("14 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m14 "),
			},
			{

				Big5:   []byte("16  7 18  7  3 105"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m16  7 18  7  3 105\r"),
			},
		},
		{ //64
			{

				Big5:   []byte("                      "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                      "),
			},
			{

				Big5:   []byte("50.6% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m50.6% "),
			},
			{

				Big5:   []byte("44.1% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[32m44.1% "),
			},
			{

				Big5:   []byte("83.3%     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m83.3%     "),
			},
			{

				Big5:   []byte("Team Rebs: 50  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[mTeam Rebs: 50  "),
			},
			{

				Big5:   []byte("Total TO: 18"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32mTotal TO: 18"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //65
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //66
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //67
			{

				Big5:   []byte("Team Statistics      LAC   DAL  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				DBCS:   []byte("\x1b[1;33;45mTeam Statistics      LAC   DAL  "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //68
			{

				Big5:   []byte("Fast Break Points"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mFast Break Points"),
			},
			{

				Big5:   []byte("     13     7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m     13     7\r"),
			},
		},
		{ //69
			{

				Big5:   []byte("Biggest Lead  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mBiggest Lead  "),
			},
			{

				Big5:   []byte("         7    16"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m         7    16\r"),
			},
		},
		{ //70
			{

				Big5:   []byte("Points In The Paint"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mPoints In The Paint"),
			},
			{

				Big5:   []byte("   42    30"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m   42    30\r"),
			},
		},
		{ //71
			{

				Big5:   []byte("Bench Points  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mBench Points  "),
			},
			{

				Big5:   []byte("        11    11"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m        11    11\r"),
			},
		},
		{ //72
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //73
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
	}

	testContent14Utf8 = [][]*types.Rune{
		{ //0
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //1
			{
				Utf8:   "2020-21 NBA Regular Season",
				Big5:   []byte("2020-21 NBA Regular Season"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[m\x1b[1;44m2020-21 NBA Regular Season"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //2
			{
				Utf8:   "￣￣￣￣￣￣￣￣￣￣￣￣￣",
				Big5:   []byte("\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1m\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3\xa1\xc3"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //3
			{
				Utf8:   " LA Clippers ",
				Big5:   []byte(" LA Clippers "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCS:   []byte("\x1b[1;41m LA Clippers "),
			},
			{
				Utf8:   " @ ",
				Big5:   []byte(" @ "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m @ "),
			},
			{
				Utf8:   " Dallas Mavericks ",
				Big5:   []byte(" Dallas Mavericks "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				DBCS:   []byte("\x1b[1;45m Dallas Mavericks "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //4
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
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
				Utf8:   "Team  1st   2nd   3rd   4th  Total ",
				Big5:   []byte("Team  1st   2nd   3rd   4th  Total "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;46mTeam  1st   2nd   3rd   4th  Total "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //7
			{
				Utf8:   "LAC    29    22    20    18    89",
				Big5:   []byte("LAC    29    22    20    18    89"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("LAC    29    22    20    18    89\r"),
			},
		},
		{ //8
			{
				Utf8:   "DAL    24    32    21    28   105",
				Big5:   []byte("DAL    24    32    21    28   105"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("DAL    24    32    21    28   105\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //11
			{
				Utf8:   "Arena Stats             ",
				Big5:   []byte("Arena Stats             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;46mArena Stats             "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //12
			{
				Utf8:   "Stadium: ",
				Big5:   []byte("Stadium: "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Stadium: "),
			},
			{
				Utf8:   "American Airlines Center",
				Big5:   []byte("American Airlines Center"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mAmerican Airlines Center"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //13
			{
				Utf8:   "         ",
				Big5:   []byte("         "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("         "),
			},
			{
				Utf8:   "Dallas, TX",
				Big5:   []byte("Dallas, TX"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mDallas, TX"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //14
			{
				Utf8:   "Attendance: ",
				Big5:   []byte("Attendance: "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Attendance: "),
			},
			{
				Utf8:   "3975",
				Big5:   []byte("3975"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m3975"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //15
			{
				Utf8:   "Duration:   ",
				Big5:   []byte("Duration:   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Duration:   "),
			},
			{
				Utf8:   "2:15",
				Big5:   []byte("2:15"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m2:15"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
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
				Utf8:   "Scoring                 ",
				Big5:   []byte("Scoring                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;46mScoring                 "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //19
			{
				Utf8:   "Lead Changes: ",
				Big5:   []byte("Lead Changes: "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Lead Changes: "),
			},
			{
				Utf8:   "11",
				Big5:   []byte("11"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m11"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //20
			{
				Utf8:   "Times Tied:   ",
				Big5:   []byte("Times Tied:   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Times Tied:   "),
			},
			{
				Utf8:   "5",
				Big5:   []byte("5"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m5"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //21
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //22
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //23
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //24
			{
				Utf8:   " LA Clippers  (26-16)                                                          ",
				Big5:   []byte(" LA Clippers  (26-16)                                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCS:   []byte("\x1b[1;41m LA Clippers  (26-16)                                                          "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //25
			{
				Utf8:   "PLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF ",
				Big5:   []byte("PLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[1;42mPLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //26
			{
				Utf8:   "K. Leonard    F 36:55  9-21 ",
				Big5:   []byte("K. Leonard    F 36:55  9-21 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("K. Leonard    F 36:55  9-21 "),
			},
			{
				Utf8:   "  1-4   1-4  ",
				Big5:   []byte("  1-4   1-4  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  1-4   1-4  "),
			},
			{
				Utf8:   "-8  2  5  7  7  3  2  2  0  2 ",
				Big5:   []byte("-8  2  5  7  7  3  2  2  0  2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m-8  2  5  7  7  3  2  2  0  2 "),
			},
			{
				Utf8:   " 20  ",
				Big5:   []byte(" 20  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m 20  "),
			},
			{
				Utf8:   "19",
				Big5:   []byte("19"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m19\r"),
			},
		},
		{ //27
			{
				Utf8:   "M. Morris Sr. F 36:03  5-14 ",
				Big5:   []byte("M. Morris Sr. F 36:03  5-14 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("M. Morris Sr. F 36:03  5-14 "),
			},
			{
				Utf8:   "  1-9   ",
				Big5:   []byte("  1-9   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  1-9   "),
			},
			{
				Utf8:   "0-0 -11  1  6  7  2  3  2  0  0  0  11  13",
				Big5:   []byte("0-0 -11  1  6  7  2  3  2  0  0  0  11  13"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0 -11  1  6  7  2  3  2  0  0  0  11  13\r"),
			},
		},
		{ //28
			{
				Utf8:   "I. Zubac      C 35:02 ",
				Big5:   []byte("I. Zubac      C 35:02 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("I. Zubac      C 35:02 "),
			},
			{
				Utf8:   "  7-9",
				Big5:   []byte("  7-9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m  7-9"),
			},
			{
				Utf8:   "   0-0   0-0 -15  4  3  7  2  2  3  2  1  1  14  23",
				Big5:   []byte("   0-0   0-0 -15  4  3  7  2  2  3  2  1  1  14  23"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m   0-0   0-0 -15  4  3  7  2  2  3  2  1  1  14  23\r"),
			},
		},
		{ //29
			{
				Utf8:   "P. George     G 41:21 10-20 ",
				Big5:   []byte("P. George     G 41:21 10-20 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("P. George     G 41:21 10-20 "),
			},
			{
				Utf8:   "  5-8 ",
				Big5:   []byte("  5-8 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m  5-8 "),
			},
			{
				Utf8:   "  3-3 ",
				Big5:   []byte("  3-3 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31m  3-3 "),
			},
			{
				Utf8:   "-10  0  7  7  5  0  1  4  1  1 ",
				Big5:   []byte("-10  0  7  7  5  0  1  4  1  1 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m-10  0  7  7  5  0  1  4  1  1 "),
			},
			{
				Utf8:   " 28  ",
				Big5:   []byte(" 28  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m 28  "),
			},
			{
				Utf8:   "28",
				Big5:   []byte("28"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m28\r"),
			},
		},
		{ //30
			{
				Utf8:   "R. Jackson    G 19:16 ",
				Big5:   []byte("R. Jackson    G 19:16 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("R. Jackson    G 19:16 "),
			},
			{
				Utf8:   "  2-8   0-3   ",
				Big5:   []byte("  2-8   0-3   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  2-8   0-3   "),
			},
			{
				Utf8:   "1-2 -13  0  2  2  1  2  2  1  0  1   5   2",
				Big5:   []byte("1-2 -13  0  2  2  1  2  2  1  0  1   5   2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m1-2 -13  0  2  2  1  2  2  1  0  1   5   2\r"),
			},
		},
		{ //31
			{
				Utf8:   "N. Batum        29:00   1-3   1-2   0-0  -3  0  2  2  2  2  2  1  0  0   3   6",
				Big5:   []byte("N. Batum        29:00   1-3   1-2   0-0  -3  0  2  2  2  2  2  1  0  0   3   6"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("N. Batum        29:00   1-3   1-2   0-0  -3  0  2  2  2  2  2  1  0  0   3   6\r"),
			},
		},
		{ //32
			{
				Utf8:   "L. Williams     24:24 ",
				Big5:   []byte("L. Williams     24:24 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("L. Williams     24:24 "),
			},
			{
				Utf8:   " 2-10   1-4   ",
				Big5:   []byte(" 2-10   1-4   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m 2-10   1-4   "),
			},
			{
				Utf8:   "0-0  -5  0  1  1  3  0  0  1  1  2   5   1",
				Big5:   []byte("0-0  -5  0  1  1  3  0  0  1  1  2   5   1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0  -5  0  1  1  3  0  0  1  1  2   5   1\r"),
			},
		},
		{ //33
			{
				Utf8:   "T. Mann         13:33   1-2 ",
				Big5:   []byte("T. Mann         13:33   1-2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("T. Mann         13:33   1-2 "),
			},
			{
				Utf8:   "  0-1   ",
				Big5:   []byte("  0-1   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  0-1   "),
			},
			{
				Utf8:   "1-2  -6  1  3  4  0  3  0  0  0  0   3   5",
				Big5:   []byte("1-2  -6  1  3  4  0  3  0  0  0  0   3   5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m1-2  -6  1  3  4  0  3  0  0  0  0   3   5\r"),
			},
		},
		{ //34
			{
				Utf8:   "A. Coffey       01:32   0-0   0-0   0-0  -3  0  0  0  0  0  0  0  0  0   0   0",
				Big5:   []byte("A. Coffey       01:32   0-0   0-0   0-0  -3  0  0  0  0  0  0  0  0  0   0   0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("A. Coffey       01:32   0-0   0-0   0-0  -3  0  0  0  0  0  0  0  0  0   0   0\r"),
			},
		},
		{ //35
			{
				Utf8:   "D. Oturu        01:27   0-0   0-0   0-0  -3  0  0  0  0  0  0  1  0  0   0 ",
				Big5:   []byte("D. Oturu        01:27   0-0   0-0   0-0  -3  0  0  0  0  0  0  1  0  0   0 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("D. Oturu        01:27   0-0   0-0   0-0  -3  0  0  0  0  0  0  1  0  0   0 "),
			},
			{
				Utf8:   " -1",
				Big5:   []byte(" -1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m -1"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //36
			{
				Utf8:   "M. Kabengele    01:27 ",
				Big5:   []byte("M. Kabengele    01:27 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("M. Kabengele    01:27 "),
			},
			{
				Utf8:   "  0-1   0-1   ",
				Big5:   []byte("  0-1   0-1   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  0-1   0-1   "),
			},
			{
				Utf8:   "0-0  -3  0  0  0  0  0  0  0  0  0   0 ",
				Big5:   []byte("0-0  -3  0  0  0  0  0  0  0  0  0   0 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0  -3  0  0  0  0  0  0  0  0  0   0 "),
			},
			{
				Utf8:   " -1",
				Big5:   []byte(" -1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m -1"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //37
			{
				Utf8:   "L. Kennard       DNP    Coach's Decision",
				Big5:   []byte("L. Kennard       DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mL. Kennard       DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //38
			{
				Utf8:   "P. Patterson     DNP    Coach's Decision",
				Big5:   []byte("P. Patterson     DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mP. Patterson     DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //39
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //40
			{
				Utf8:   "Total            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     ",
				Big5:   []byte("Total            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_GREEN, Highlight: true},
				DBCS:   []byte("\x1b[1;42mTotal            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //41
			{
				Utf8:   "                 ",
				Big5:   []byte("                 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                 "),
			},
			{
				Utf8:   "240  ",
				Big5:   []byte("240  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1m240  "),
			},
			{
				Utf8:   "37-88 ",
				Big5:   []byte("37-88 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m37-88 "),
			},
			{
				Utf8:   " 9-32 ",
				Big5:   []byte(" 9-32 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m 9-32 "),
			},
			{
				Utf8:   " 6-11     ",
				Big5:   []byte(" 6-11     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m 6-11     "),
			},
			{
				Utf8:   " 8 29 37 ",
				Big5:   []byte(" 8 29 37 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m 8 29 37 "),
			},
			{
				Utf8:   "22 ",
				Big5:   []byte("22 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m22 "),
			},
			{
				Utf8:   "15 12 12  3  7  89",
				Big5:   []byte("15 12 12  3  7  89"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m15 12 12  3  7  89\r"),
			},
		},
		{ //42
			{
				Utf8:   "                      ",
				Big5:   []byte("                      "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                      "),
			},
			{
				Utf8:   "42.0% ",
				Big5:   []byte("42.0% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m42.0% "),
			},
			{
				Utf8:   "28.1% ",
				Big5:   []byte("28.1% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[32m28.1% "),
			},
			{
				Utf8:   "54.5%     ",
				Big5:   []byte("54.5%     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m54.5%     "),
			},
			{
				Utf8:   "Team Rebs: 47  ",
				Big5:   []byte("Team Rebs: 47  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[mTeam Rebs: 47  "),
			},
			{
				Utf8:   "Total TO: 13",
				Big5:   []byte("Total TO: 13"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32mTotal TO: 13"),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //43
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //44
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //45
			{
				Utf8:   " Dallas Mavericks  (21-18)                                                     ",
				Big5:   []byte(" Dallas Mavericks  (21-18)                                                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				DBCS:   []byte("\x1b[1;45m Dallas Mavericks  (21-18)                                                     "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //46
			{
				Utf8:   "PLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF ",
				Big5:   []byte("PLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				DBCS:   []byte("\x1b[1;43mPLAYER        P  MIN  FGM-A 3PM-A FTM-A +/- OR DR TR AS PF ST TO BS BA PTS EFF "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //47
			{
				Utf8:   "Hardaway Jr.  F 36:34  6-12   3-7   0-0 +17  0  6  6  1  1  0  2  0  0  15  14",
				Big5:   []byte("Hardaway Jr.  F 36:34  6-12   3-7   0-0 +17  0  6  6  1  1  0  2  0  0  15  14"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Hardaway Jr.  F 36:34  6-12   3-7   0-0 +17  0  6  6  1  1  0  2  0  0  15  14\r"),
			},
		},
		{ //48
			{
				Utf8:   "M. Kleber     F 40:54   4-8   2-5 ",
				Big5:   []byte("M. Kleber     F 40:54   4-8   2-5 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("M. Kleber     F 40:54   4-8   2-5 "),
			},
			{
				Utf8:   "  2-2 ",
				Big5:   []byte("  2-2 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m  2-2 "),
			},
			{
				Utf8:   "+24  0  6  6  0  3  0  1  0  0  12  13",
				Big5:   []byte("+24  0  6  6  0  3  0  1  0  0  12  13"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m+24  0  6  6  0  3  0  1  0  0  12  13\r"),
			},
		},
		{ //49
			{
				Utf8:   "K. Porzingis  C 34:19  5-14",
				Big5:   []byte("K. Porzingis  C 34:19  5-14"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[1;44mK. Porzingis  C 34:19  5-14"),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[;37;44m "),
			},
			{
				Utf8:   "  1-4   0-1 ",
				Big5:   []byte("  1-4   0-1 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  1-4   0-1 "),
			},
			{
				Utf8:   " +9  3 10",
				Big5:   []byte(" +9  3 10"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[37m +9  3 10"),
			},
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[;37;44m "),
			},
			{
				Utf8:   "13 ",
				Big5:   []byte("13 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[1;31m13 "),
			},
			{
				Utf8:   " 0  2  0  4  2  1  11  12 ",
				Big5:   []byte(" 0  2  0  4  2  1  11  12 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLUE, Highlight: true},
				DBCS:   []byte("\x1b[37m 0  2  0  4  2  1  11  12 "),
			},
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //50
			{
				Utf8:   "J. Richardson G 38:51   5-9   2-4 ",
				Big5:   []byte("J. Richardson G 38:51   5-9   2-4 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("J. Richardson G 38:51   5-9   2-4 "),
			},
			{
				Utf8:   "  2-2 ",
				Big5:   []byte("  2-2 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m  2-2 "),
			},
			{
				Utf8:   "+16  1  4  5  2  2  3  1  1  0  14  20",
				Big5:   []byte("+16  1  4  5  2  2  3  1  1  0  14  20"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m+16  1  4  5  2  2  3  1  1  0  14  20\r"),
			},
		},
		{ //51
			{
				Utf8:   "L. Doncic     G 42:48 16-28  6-11   4-5 +29  1  5  6  9  2  3  4  2  2 ",
				Big5:   []byte("L. Doncic     G 42:48 16-28  6-11   4-5 +29  1  5  6  9  2  3  4  2  2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("L. Doncic     G 42:48 16-28  6-11   4-5 +29  1  5  6  9  2  3  4  2  2 "),
			},
			{
				Utf8:   " 42  45",
				Big5:   []byte(" 42  45"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m 42  45"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //52
			{
				Utf8:   "J. Brunson      15:44   2-5   1-2   0-0 -11  1  2  3  2  2  1  5  0  0   5   3",
				Big5:   []byte("J. Brunson      15:44   2-5   1-2   0-0 -11  1  2  3  2  2  1  5  0  0   5   3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("J. Brunson      15:44   2-5   1-2   0-0 -11  1  2  3  2  2  1  5  0  0   5   3\r"),
			},
		},
		{ //53
			{
				Utf8:   "W. Iwundu       10:37   0-0   0-0   0-0  -2  0  2  2  0  2  0  1  0  0   0   1",
				Big5:   []byte("W. Iwundu       10:37   0-0   0-0   0-0  -2  0  2  2  0  2  0  1  0  0   0   1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("W. Iwundu       10:37   0-0   0-0   0-0  -2  0  2  2  0  2  0  1  0  0   0   1\r"),
			},
		},
		{ //54
			{
				Utf8:   "T. Burke        04:59 ",
				Big5:   []byte("T. Burke        04:59 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("T. Burke        04:59 "),
			},
			{
				Utf8:   "  1-1 ",
				Big5:   []byte("  1-1 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m  1-1 "),
			},
			{
				Utf8:   "  0-0",
				Big5:   []byte("  0-0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m  0-0"),
			},
			{
				Utf8:   "   ",
				Big5:   []byte("   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m   "),
			},
			{
				Utf8:   "0-0  -7  0  0  0  0  0  0  0  0  0   2   2",
				Big5:   []byte("0-0  -7  0  0  0  0  0  0  0  0  0   2   2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m0-0  -7  0  0  0  0  0  0  0  0  0   2   2\r"),
			},
		},
		{ //55
			{
				Utf8:   "D. Powell       03:19   0-0   0-0   0-0  -5  0  0  0  0  0  0  0  0  0   0   0",
				Big5:   []byte("D. Powell       03:19   0-0   0-0   0-0  -5  0  0  0  0  0  0  0  0  0   0   0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("D. Powell       03:19   0-0   0-0   0-0  -5  0  0  0  0  0  0  0  0  0   0   0\r"),
			},
		},
		{ //56
			{
				Utf8:   "Cauley-Stein    11:55   1-2 ",
				Big5:   []byte("Cauley-Stein    11:55   1-2 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("Cauley-Stein    11:55   1-2 "),
			},
			{
				Utf8:   "  0-1 ",
				Big5:   []byte("  0-1 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m  0-1 "),
			},
			{
				Utf8:   "  2-2 ",
				Big5:   []byte("  2-2 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31m  2-2 "),
			},
			{
				Utf8:   "+10  0  2  2  0  2  0  0  2  0   4   7",
				Big5:   []byte("+10  0  2  2  0  2  0  0  2  0   4   7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m+10  0  2  2  0  2  0  0  2  0   4   7\r"),
			},
		},
		{ //57
			{
				Utf8:   "T. Bey           DNP    Coach's Decision",
				Big5:   []byte("T. Bey           DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mT. Bey           DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //58
			{
				Utf8:   "J. Green         DNP    Coach's Decision",
				Big5:   []byte("J. Green         DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mJ. Green         DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //59
			{
				Utf8:   "N. Hinton        DNP    Coach's Decision",
				Big5:   []byte("N. Hinton        DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mN. Hinton        DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //60
			{
				Utf8:   "B. Marjanovic    DNP    Coach's Decision",
				Big5:   []byte("B. Marjanovic    DNP    Coach's Decision"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32mB. Marjanovic    DNP    Coach's Decision"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //61
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //62
			{
				Utf8:   "Total            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     ",
				Big5:   []byte("Total            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_YELLOW, Highlight: true},
				DBCS:   []byte("\x1b[1;43mTotal            MIN  FGM-A 3PM-A FTM-A     OR DR TR AS PF ST TO BS BA PTS     "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //63
			{
				Utf8:   "                 ",
				Big5:   []byte("                 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                 "),
			},
			{
				Utf8:   "240  ",
				Big5:   []byte("240  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1m240  "),
			},
			{
				Utf8:   "40-79 ",
				Big5:   []byte("40-79 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m40-79 "),
			},
			{
				Utf8:   "15-34 ",
				Big5:   []byte("15-34 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32m15-34 "),
			},
			{
				Utf8:   "10-12      ",
				Big5:   []byte("10-12      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m10-12      "),
			},
			{
				Utf8:   "6 37 43 ",
				Big5:   []byte("6 37 43 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m6 37 43 "),
			},
			{
				Utf8:   "14 ",
				Big5:   []byte("14 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36m14 "),
			},
			{
				Utf8:   "16  7 18  7  3 105",
				Big5:   []byte("16  7 18  7  3 105"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m16  7 18  7  3 105\r"),
			},
		},
		{ //64
			{
				Utf8:   "                      ",
				Big5:   []byte("                      "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                      "),
			},
			{
				Utf8:   "50.6% ",
				Big5:   []byte("50.6% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m50.6% "),
			},
			{
				Utf8:   "44.1% ",
				Big5:   []byte("44.1% "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[32m44.1% "),
			},
			{
				Utf8:   "83.3%     ",
				Big5:   []byte("83.3%     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33m83.3%     "),
			},
			{
				Utf8:   "Team Rebs: 50  ",
				Big5:   []byte("Team Rebs: 50  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[mTeam Rebs: 50  "),
			},
			{
				Utf8:   "Total TO: 18",
				Big5:   []byte("Total TO: 18"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;32mTotal TO: 18"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //65
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //66
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //67
			{
				Utf8:   "Team Statistics      LAC   DAL  ",
				Big5:   []byte("Team Statistics      LAC   DAL  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_MAGENTA, Highlight: true},
				DBCS:   []byte("\x1b[1;33;45mTeam Statistics      LAC   DAL  "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //68
			{
				Utf8:   "Fast Break Points",
				Big5:   []byte("Fast Break Points"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mFast Break Points"),
			},
			{
				Utf8:   "     13     7",
				Big5:   []byte("     13     7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m     13     7\r"),
			},
		},
		{ //69
			{
				Utf8:   "Biggest Lead  ",
				Big5:   []byte("Biggest Lead  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mBiggest Lead  "),
			},
			{
				Utf8:   "         7    16",
				Big5:   []byte("         7    16"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m         7    16\r"),
			},
		},
		{ //70
			{
				Utf8:   "Points In The Paint",
				Big5:   []byte("Points In The Paint"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mPoints In The Paint"),
			},
			{
				Utf8:   "   42    30",
				Big5:   []byte("   42    30"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m   42    30\r"),
			},
		},
		{ //71
			{
				Utf8:   "Bench Points  ",
				Big5:   []byte("Bench Points  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33mBench Points  "),
			},
			{
				Utf8:   "        11    11",
				Big5:   []byte("        11    11"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m        11    11\r"),
			},
		},
		{ //72
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //73
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
	}
}
