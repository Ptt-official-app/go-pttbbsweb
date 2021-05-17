package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
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

	testFirstComments14     []*schema.Comment
	testFullFirstComments14 []*schema.Comment

	testTheRestComments14 []*schema.Comment
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

	testFirstComments14 = []*schema.Comment{
		{ //0
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("guardyo"),

			MD5:     "3VevXQUX8i9UiFETWDrcTQ",
			TheDate: "03/18/2021 12:07:22",
			IP:      "111.249.44.44",
			Host:    "臺灣",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: guardyo (111.249.44.44 \xbbO\xc6W), 03/18/2021 12:07:22\r"),

			CreateTime:         1616040442000000000,
			InferredCreateTime: 1616040442000000000,
			SortTime:           1616040442000000000,
			CommentID:          "Fm1UOEgTRAA:3VevXQUX8i9UiFETWDrcTQ",
		},
		{ //1
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("yasan1029"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777",
						Big5:   []byte("77777777                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777                                         "),
					},
				},
			},
			MD5:     "-HAQCmGXnoN2PhhEhZk4zg",
			TheDate: "03/18 12:07",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33myasan1029   \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:07\r"),
		},
		{ //2
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("yue1013"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "777",
						Big5:   []byte("777                                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("777                                              "),
					},
				},
			},
			MD5:     "wZ64mftwKYpfdP5rodU4Nw",
			TheDate: "03/18 12:07",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33myue1013     \x1b[m\x1b[33m: 777                                              \x1b[m 03/18 12:07\r"),
		},
		{ //3
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Dorae5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這季打完大概又要上演人生了...",
						Big5:   []byte("\xb3o\xa9u\xa5\xb4\xa7\xb9\xa4j\xb7\xa7\xa4S\xadn\xa4W\xbat\xa4H\xa5\xcd\xa4F...                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xa9u\xa5\xb4\xa7\xb9\xa4j\xb7\xa7\xa4S\xadn\xa4W\xbat\xa4H\xa5\xcd\xa4F...                    "),
					},
				},
			},
			MD5:     "fN3AbmStA9XU50hqTnr2wQ",
			TheDate: "03/18 12:07",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDorae5566   \x1b[m\x1b[33m: \xb3o\xa9u\xa5\xb4\xa7\xb9\xa4j\xb7\xa7\xa4S\xadn\xa4W\xbat\xa4H\xa5\xcd\xa4F...                    \x1b[m 03/18 12:07\r"),
		},

		{ //4
			TheType: types.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("doooooooooog"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Mavericks",
						Big5:   []byte("Mavericks"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Mavericks"),
					},
				},
			},
			MD5:     "px43SMg0XuAbpHEULwG-Hw",
			TheDate: "03/18 12:07",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mdoooooooooog\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Mavericks\x1b[m                               03/18 12:07\r"),
		},

		{ //5
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("luck945"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777777",
						Big5:   []byte("7777777777                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777777                                       "),
					},
				},
			},
			MD5:     "UvSl9EqCLckOYXsUfwbMow",
			TheDate: "03/18 12:07",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mluck945     \x1b[m\x1b[33m: 7777777777                                       \x1b[m 03/18 12:07\r"),
		},
		{ //6
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Eddward"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "快艇電梯向下？",
						Big5:   []byte("\xa7\xd6\xb8\xa5\xb9q\xb1\xe8\xa6V\xa4U\xa1H                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xd6\xb8\xa5\xb9q\xb1\xe8\xa6V\xa4U\xa1H                                   "),
					},
				},
			},
			MD5:     "E9p3SESmurRvPhaRGzQ0rA",
			TheDate: "03/18 12:07",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mEddward     \x1b[m\x1b[33m: \xa7\xd6\xb8\xa5\xb9q\xb1\xe8\xa6V\xa4U\xa1H                                   \x1b[m 03/18 12:07\r"),
		},
		{ //7
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("kkxf336g037"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777",
						Big5:   []byte("77777777                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777                                         "),
					},
				},
			},
			MD5:     "Ra9SNKTFV7_KV7XGtEIDEQ",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mkkxf336g037 \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:08\r"),
		},
		{ //8
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("outnow5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這三分我還以為是P.league呢",
						Big5:   []byte("\xb3o\xa4T\xa4\xc0\xa7\xda\xc1\xd9\xa5H\xac\xb0\xacOP.league\xa9O                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xa4T\xa4\xc0\xa7\xda\xc1\xd9\xa5H\xac\xb0\xacOP.league\xa9O                       "),
					},
				},
			},
			MD5:     "XPC8m63E6iHQTXQoif9Tgw",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33moutnow5566  \x1b[m\x1b[33m: \xb3o\xa4T\xa4\xc0\xa7\xda\xc1\xd9\xa5H\xac\xb0\xacOP.league\xa9O                       \x1b[m 03/18 12:08\r"),
		},
		{ //9
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("doooooooooog"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777",
						Big5:   []byte("7777777                                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777                                          "),
					},
				},
			},
			MD5:     "GszoIqT5LvkbJoniDsn8nA",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mdoooooooooog\x1b[m\x1b[33m: 7777777                                          \x1b[m 03/18 12:08\r"),
		},
	}

	testFullFirstComments14 = []*schema.Comment{
		{ //0
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_EDIT,
			Owner:     bbs.UUserID("guardyo"),

			MD5:     "3VevXQUX8i9UiFETWDrcTQ",
			TheDate: "03/18/2021 12:07:22",
			IP:      "111.249.44.44",
			Host:    "臺灣",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: guardyo (111.249.44.44 \xbbO\xc6W), 03/18/2021 12:07:22\r"),

			CreateTime:         1616040442000000000,
			InferredCreateTime: 1616040442000000000,
			SortTime:           1616040442000000000,
			CommentID:          "Fm1UOEgTRAA:3VevXQUX8i9UiFETWDrcTQ",
		},
		{ //1
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("yasan1029"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777",
						Big5:   []byte("77777777                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777                                         "),
					},
				},
			},
			MD5:        "-HAQCmGXnoN2PhhEhZk4zg",
			TheDate:    "03/18 12:07",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33myasan1029   \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:07\r"),
			CommentID:  "Fm1UOEgihkA:-HAQCmGXnoN2PhhEhZk4zg",
			CreateTime: 1616040420000000000,
			SortTime:   1616040442001000000,
		},
		{ //2
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("yue1013"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "777",
						Big5:   []byte("777                                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("777                                              "),
					},
				},
			},
			MD5:        "wZ64mftwKYpfdP5rodU4Nw",
			TheDate:    "03/18 12:07",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33myue1013     \x1b[m\x1b[33m: 777                                              \x1b[m 03/18 12:07\r"),
			CommentID:  "Fm1UOEgxyIA:wZ64mftwKYpfdP5rodU4Nw",
			CreateTime: 1616040420000000000,
			SortTime:   1616040442002000000,
		},
		{ //3
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("Dorae5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這季打完大概又要上演人生了...",
						Big5:   []byte("\xb3o\xa9u\xa5\xb4\xa7\xb9\xa4j\xb7\xa7\xa4S\xadn\xa4W\xbat\xa4H\xa5\xcd\xa4F...                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xa9u\xa5\xb4\xa7\xb9\xa4j\xb7\xa7\xa4S\xadn\xa4W\xbat\xa4H\xa5\xcd\xa4F...                    "),
					},
				},
			},
			MD5:        "fN3AbmStA9XU50hqTnr2wQ",
			TheDate:    "03/18 12:07",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDorae5566   \x1b[m\x1b[33m: \xb3o\xa9u\xa5\xb4\xa7\xb9\xa4j\xb7\xa7\xa4S\xadn\xa4W\xbat\xa4H\xa5\xcd\xa4F...                    \x1b[m 03/18 12:07\r"),
			CommentID:  "Fm1UOEhBCsA:fN3AbmStA9XU50hqTnr2wQ",
			CreateTime: 1616040420000000000,
			SortTime:   1616040442003000000,
		},

		{ //4
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("doooooooooog"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Mavericks",
						Big5:   []byte("Mavericks"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Mavericks"),
					},
				},
			},
			MD5:        "px43SMg0XuAbpHEULwG-Hw",
			TheDate:    "03/18 12:07",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mdoooooooooog\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Mavericks\x1b[m                               03/18 12:07\r"),
			CommentID:  "Fm1UOEhQTQA:px43SMg0XuAbpHEULwG-Hw",
			CreateTime: 1616040420000000000,
			SortTime:   1616040442004000000,
		},

		{ //5
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("luck945"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777777",
						Big5:   []byte("7777777777                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777777                                       "),
					},
				},
			},
			MD5:        "UvSl9EqCLckOYXsUfwbMow",
			TheDate:    "03/18 12:07",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mluck945     \x1b[m\x1b[33m: 7777777777                                       \x1b[m 03/18 12:07\r"),
			CommentID:  "Fm1UOEhfj0A:UvSl9EqCLckOYXsUfwbMow",
			CreateTime: 1616040420000000000,
			SortTime:   1616040442005000000,
		},
		{ //6
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("Eddward"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "快艇電梯向下？",
						Big5:   []byte("\xa7\xd6\xb8\xa5\xb9q\xb1\xe8\xa6V\xa4U\xa1H                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xd6\xb8\xa5\xb9q\xb1\xe8\xa6V\xa4U\xa1H                                   "),
					},
				},
			},
			MD5:        "E9p3SESmurRvPhaRGzQ0rA",
			TheDate:    "03/18 12:07",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mEddward     \x1b[m\x1b[33m: \xa7\xd6\xb8\xa5\xb9q\xb1\xe8\xa6V\xa4U\xa1H                                   \x1b[m 03/18 12:07\r"),
			CommentID:  "Fm1UOEhu0YA:E9p3SESmurRvPhaRGzQ0rA",
			CreateTime: 1616040420000000000,
			SortTime:   1616040442006000000,
		},
		{ //7
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("kkxf336g037"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777",
						Big5:   []byte("77777777                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777                                         "),
					},
				},
			},
			MD5:        "Ra9SNKTFV7_KV7XGtEIDEQ",
			TheDate:    "03/18 12:08",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mkkxf336g037 \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:08\r"),
			CommentID:  "Fm1UQSENQAA:Ra9SNKTFV7_KV7XGtEIDEQ",
			CreateTime: 1616040480000000000,
			SortTime:   1616040480000000000,
		},
		{ //8
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("outnow5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這三分我還以為是P.league呢",
						Big5:   []byte("\xb3o\xa4T\xa4\xc0\xa7\xda\xc1\xd9\xa5H\xac\xb0\xacOP.league\xa9O                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xa4T\xa4\xc0\xa7\xda\xc1\xd9\xa5H\xac\xb0\xacOP.league\xa9O                       "),
					},
				},
			},
			MD5:        "XPC8m63E6iHQTXQoif9Tgw",
			TheDate:    "03/18 12:08",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33moutnow5566  \x1b[m\x1b[33m: \xb3o\xa4T\xa4\xc0\xa7\xda\xc1\xd9\xa5H\xac\xb0\xacOP.league\xa9O                       \x1b[m 03/18 12:08\r"),
			CommentID:  "Fm1UQSEcgkA:XPC8m63E6iHQTXQoif9Tgw",
			CreateTime: 1616040480000000000,
			SortTime:   1616040480001000000,
		},
		{ //9
			BBoardID:  "test",
			ArticleID: "test14",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("doooooooooog"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777",
						Big5:   []byte("7777777                                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777                                          "),
					},
				},
			},
			MD5:        "GszoIqT5LvkbJoniDsn8nA",
			TheDate:    "03/18 12:08",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mdoooooooooog\x1b[m\x1b[33m: 7777777                                          \x1b[m 03/18 12:08\r"),
			CommentID:  "Fm1UQSErxIA:GszoIqT5LvkbJoniDsn8nA",
			CreateTime: 1616040480000000000,
			SortTime:   1616040480002000000,
		},
	}

	testTheRestComments14 = []*schema.Comment{
		{ //0
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("JTFWAHQ"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777777777777777777777",
						Big5:   []byte("77777777777777777777777777                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777777777777777777777                       "),
					},
				},
			},
			MD5:     "wqpZ9lDTnYcD0moScLqLjw",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mJTFWAHQ     \x1b[m\x1b[33m: 77777777777777777777777777                       \x1b[m 03/18 12:08\r"),
		},
		{ //1
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("HardDDDD"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777777777777777",
						Big5:   []byte("77777777777777777777                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777777777777777                             "),
					},
				},
			},
			MD5:     "3lKNZOVNh8zRtqDB4fY3ag",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mHardDDDD    \x1b[m\x1b[33m: 77777777777777777777                             \x1b[m 03/18 12:08\r"),
		},
		{ //2
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("MilkTea0509"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "借轉烙賽版",
						Big5:   []byte("\xad\xc9\xc2\xe0\xafO\xc1\xc9\xaa\xa9                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xad\xc9\xc2\xe0\xafO\xc1\xc9\xaa\xa9                                       "),
					},
				},
			},
			MD5:     "Bae9hkebo2JtOUBLD8boQw",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mMilkTea0509 \x1b[m\x1b[33m: \xad\xc9\xc2\xe0\xafO\xc1\xc9\xaa\xa9                                       \x1b[m 03/18 12:08\r"),
		},
		{ //3
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("boy80421"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77跟快艇是不是有仇 印象中打快艇常爆氣",
						Big5:   []byte("77\xb8\xf2\xa7\xd6\xb8\xa5\xacO\xa4\xa3\xacO\xa6\xb3\xa4\xb3 \xa6L\xb6H\xa4\xa4\xa5\xb4\xa7\xd6\xb8\xa5\xb1`\xc3z\xae\xf0            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xb8\xf2\xa7\xd6\xb8\xa5\xacO\xa4\xa3\xacO\xa6\xb3\xa4\xb3 \xa6L\xb6H\xa4\xa4\xa5\xb4\xa7\xd6\xb8\xa5\xb1`\xc3z\xae\xf0            "),
					},
				},
			},
			MD5:     "V7SJz7XXJApwus9m5PUEkQ",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mboy80421    \x1b[m\x1b[33m: 77\xb8\xf2\xa7\xd6\xb8\xa5\xacO\xa4\xa3\xacO\xa6\xb3\xa4\xb3 \xa6L\xb6H\xa4\xa4\xa5\xb4\xa7\xd6\xb8\xa5\xb1`\xc3z\xae\xf0            \x1b[m 03/18 12:08\r"),
		},
		{ //4
			TheType: types.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("MilkTea0509"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "LaClippers",
						Big5:   []byte("LaClippers"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("LaClippers"),
					},
				},
			},
			MD5:     "7Gsdh9Cywjm66UnkM89YtQ",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mMilkTea0509\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO LaClippers\x1b[m                               03/18 12:08\r"),
		},
		{ //5
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("homocat"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77不錯",
						Big5:   []byte("77\xa4\xa3\xbf\xf9                                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xa4\xa3\xbf\xf9                                           "),
					},
				},
			},
			MD5:     "ZYr--Qi9wWKviy68ucjWHA",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mhomocat     \x1b[m\x1b[33m: 77\xa4\xa3\xbf\xf9                                           \x1b[m 03/18 12:08\r"),
		},
		{ //6
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("aa01081008tw"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "可愛會不會這季結束烙跑啊..比上季更沒戲XD",
						Big5:   []byte("\xa5i\xb7R\xb7|\xa4\xa3\xb7|\xb3o\xa9u\xb5\xb2\xa7\xf4\xafO\xb6]\xb0\xda..\xa4\xf1\xa4W\xa9u\xa7\xf3\xa8S\xc0\xb8XD         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5i\xb7R\xb7|\xa4\xa3\xb7|\xb3o\xa9u\xb5\xb2\xa7\xf4\xafO\xb6]\xb0\xda..\xa4\xf1\xa4W\xa9u\xa7\xf3\xa8S\xc0\xb8XD         "),
					},
				},
			},
			MD5:     "m6bHx4bj0mLlG3WUVNI2yw",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33maa01081008tw\x1b[m\x1b[33m: \xa5i\xb7R\xb7|\xa4\xa3\xb7|\xb3o\xa9u\xb5\xb2\xa7\xf4\xafO\xb6]\xb0\xda..\xa4\xf1\xa4W\xa9u\xa7\xf3\xa8S\xc0\xb8XD         \x1b[m 03/18 12:08\r"),
		},
		{ //7
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("chenluyang"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77這場爆幹強",
						Big5:   []byte("77\xb3o\xb3\xf5\xc3z\xb7F\xb1j                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xb3o\xb3\xf5\xc3z\xb7F\xb1j                                     "),
					},
				},
			},
			MD5:     "WOXgkIVbxhXmv5u8mkoAag",
			TheDate: "03/18 12:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mchenluyang  \x1b[m\x1b[33m: 77\xb3o\xb3\xf5\xc3z\xb7F\xb1j                                     \x1b[m 03/18 12:08\r"),
		},
		{ //8
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("jardon"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "+25  eff 45",
						Big5:   []byte("+25  eff 45                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("+25  eff 45                                      "),
					},
				},
			},
			MD5:     "-FtXbwVpV9hwndBXdqLaVA",
			TheDate: "03/18 12:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mjardon      \x1b[m\x1b[33m: +25  eff 45                                      \x1b[m 03/18 12:09\r"),
		},
		{ //9
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Toy17"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "鬼神7777777777777777777777",
						Big5:   []byte("\xb0\xad\xaf\xab7777777777777777777777                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb0\xad\xaf\xab7777777777777777777777                       "),
					},
				},
			},
			MD5:     "d3qwKrwviJaqhx5yUx8JxQ",
			TheDate: "03/18 12:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mToy17       \x1b[m\x1b[33m: \xb0\xad\xaf\xab7777777777777777777777                       \x1b[m 03/18 12:09\r"),
		},
		{ //10
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Y225"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "我777777",
						Big5:   []byte("\xa7\xda777777                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xda777777                                         "),
					},
				},
			},
			MD5:     "jm1qWGg30BCUWLvcY0XxbA",
			TheDate: "03/18 12:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mY225        \x1b[m\x1b[33m: \xa7\xda777777                                         \x1b[m 03/18 12:09\r"),
		},
		{ //11
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("darren2586"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777",
						Big5:   []byte("7777777                                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777                                          "),
					},
				},
			},
			MD5:     "TJkaWYo_F2O9ptaThfyoLw",
			TheDate: "03/18 12:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mdarren2586  \x1b[m\x1b[33m: 7777777                                          \x1b[m 03/18 12:09\r"),
		},
		{ //12
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Siika"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77好鬼",
						Big5:   []byte("77\xa6n\xb0\xad                                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xa6n\xb0\xad                                           "),
					},
				},
			},
			MD5:     "noHr5bIQhCQkeNmfS3_zPQ",
			TheDate: "03/18 12:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mSiika       \x1b[m\x1b[33m: 77\xa6n\xb0\xad                                           \x1b[m 03/18 12:09\r"),
		},
		{ //13
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("munchlax"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "慘 進攻端剩 Kawhi跟 PG在撐 這兩個又要扛防守= =",
						Big5:   []byte("\xbaG \xb6i\xa7\xf0\xba\xdd\xb3\xd1 Kawhi\xb8\xf2 PG\xa6b\xbc\xb5 \xb3o\xa8\xe2\xad\xd3\xa4S\xadn\xa6\xaa\xa8\xbe\xa6u= =   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbaG \xb6i\xa7\xf0\xba\xdd\xb3\xd1 Kawhi\xb8\xf2 PG\xa6b\xbc\xb5 \xb3o\xa8\xe2\xad\xd3\xa4S\xadn\xa6\xaa\xa8\xbe\xa6u= =   "),
					},
				},
			},
			MD5:     "tg-YCzVxKWQcxlRnU3uxPw",
			TheDate: "03/18 12:09",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mmunchlax    \x1b[m\x1b[33m: \xbaG \xb6i\xa7\xf0\xba\xdd\xb3\xd1 Kawhi\xb8\xf2 PG\xa6b\xbc\xb5 \xb3o\xa8\xe2\xad\xd3\xa4S\xadn\xa6\xaa\xa8\xbe\xa6u= =   \x1b[m 03/18 12:09\r"),
		},
		{ //14
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ryan0928"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777",
						Big5:   []byte("77777777                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777                                         "),
					},
				},
			},
			MD5:     "PEj7ECjB_hdT8ZwdnNxSMQ",
			TheDate: "03/18 12:49",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mryan0928    \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:49\r"),
		},
		{ //15
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("JessicaA1ba"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "KP今天化身鎖耶，進攻就交給萬花筒77",
						Big5:   []byte("KP\xa4\xb5\xa4\xd1\xa4\xc6\xa8\xad\xc2\xea\xadC\xa1A\xb6i\xa7\xf0\xb4N\xa5\xe6\xb5\xb9\xb8U\xaa\xe1\xb5\xa977               "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("KP\xa4\xb5\xa4\xd1\xa4\xc6\xa8\xad\xc2\xea\xadC\xa1A\xb6i\xa7\xf0\xb4N\xa5\xe6\xb5\xb9\xb8U\xaa\xe1\xb5\xa977               "),
					},
				},
			},
			MD5:     "cfRIJJ___V5vpwk338-Ggg",
			TheDate: "03/18 12:52",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mJessicaA1ba \x1b[m\x1b[33m: KP\xa4\xb5\xa4\xd1\xa4\xc6\xa8\xad\xc2\xea\xadC\xa1A\xb6i\xa7\xf0\xb4N\xa5\xe6\xb5\xb9\xb8U\xaa\xe1\xb5\xa977               \x1b[m 03/18 12:52\r"),
		},
		{ //16
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("qpeter"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "今天實在沒法子 THJ扛先發 森哥又很雷 板凳沒人可帶",
						Big5:   []byte("\xa4\xb5\xa4\xd1\xb9\xea\xa6b\xa8S\xaak\xa4l THJ\xa6\xaa\xa5\xfd\xb5o \xb4\xcb\xad\xf4\xa4S\xab\xdc\xb9p \xaaO\xb9\xb9\xa8S\xa4H\xa5i\xb1a "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xb5\xa4\xd1\xb9\xea\xa6b\xa8S\xaak\xa4l THJ\xa6\xaa\xa5\xfd\xb5o \xb4\xcb\xad\xf4\xa4S\xab\xdc\xb9p \xaaO\xb9\xb9\xa8S\xa4H\xa5i\xb1a "),
					},
				},
			},
			MD5:     "JPRAQRgnykuQrupTvv2HLQ",
			TheDate: "03/18 12:52",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mqpeter      \x1b[m\x1b[33m: \xa4\xb5\xa4\xd1\xb9\xea\xa6b\xa8S\xaak\xa4l THJ\xa6\xaa\xa5\xfd\xb5o \xb4\xcb\xad\xf4\xa4S\xab\xdc\xb9p \xaaO\xb9\xb9\xa8S\xa4H\xa5i\xb1a \x1b[m 03/18 12:52\r"),
		},
		{ //17
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("qpeter"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "只好讓077加時加班了",
						Big5:   []byte("\xa5u\xa6n\xc5\xfd077\xa5[\xae\xc9\xa5[\xafZ\xa4F                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5u\xa6n\xc5\xfd077\xa5[\xae\xc9\xa5[\xafZ\xa4F                              "),
					},
				},
			},
			MD5:     "iqUxK00D37IacQoC1n_4-Q",
			TheDate: "03/18 12:53",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mqpeter      \x1b[m\x1b[33m: \xa5u\xa6n\xc5\xfd077\xa5[\xae\xc9\xa5[\xafZ\xa4F                              \x1b[m 03/18 12:53\r"),
		},
		{ //18
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("doooooooooog"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "今天兩邊板凳很團結，11比11XD",
						Big5:   []byte("\xa4\xb5\xa4\xd1\xa8\xe2\xc3\xe4\xaaO\xb9\xb9\xab\xdc\xb9\xce\xb5\xb2\xa1A11\xa4\xf111XD                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xb5\xa4\xd1\xa8\xe2\xc3\xe4\xaaO\xb9\xb9\xab\xdc\xb9\xce\xb5\xb2\xa1A11\xa4\xf111XD                     "),
					},
				},
			},
			MD5:     "-dnzO3RC_nFemMe53_lxSg",
			TheDate: "03/18 12:54",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mdoooooooooog\x1b[m\x1b[33m: \xa4\xb5\xa4\xd1\xa8\xe2\xc3\xe4\xaaO\xb9\xb9\xab\xdc\xb9\xce\xb5\xb2\xa1A11\xa4\xf111XD                     \x1b[m 03/18 12:54\r"),
		},
		{ //19
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("dy1278dy"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "777777777",
						Big5:   []byte("777777777                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("777777777                                        "),
					},
				},
			},
			MD5:     "NC1LVO8JZtKyCogzalJPgQ",
			TheDate: "03/18 12:55",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mdy1278dy    \x1b[m\x1b[33m: 777777777                                        \x1b[m 03/18 12:55\r"),
		},
		{ //20
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("TokyoKind"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "復仇成功",
						Big5:   []byte("\xb4_\xa4\xb3\xa6\xa8\xa5\\                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb4_\xa4\xb3\xa6\xa8\xa5\\                                         "),
					},
				},
			},
			MD5:     "Z22GRmR8fWEHXJKc9vQeWw",
			TheDate: "03/18 12:56",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mTokyoKind   \x1b[m\x1b[33m: \xb4_\xa4\xb3\xa6\xa8\xa5\\                                         \x1b[m 03/18 12:56\r"),
		},
		{ //21
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("blastillocal"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77有猛到",
						Big5:   []byte("77\xa6\xb3\xb2r\xa8\xec                                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xa6\xb3\xb2r\xa8\xec                                         "),
					},
				},
			},
			MD5:     "o-eGHvDGU15tS5wo89M9Dw",
			TheDate: "03/18 12:57",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mblastillocal\x1b[m\x1b[33m: 77\xa6\xb3\xb2r\xa8\xec                                         \x1b[m 03/18 12:57\r"),
		},
		{ //22
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("YouGot5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "說好的打爆湖人咧 不要季後賽又自己先下去了",
						Big5:   []byte("\xbb\xa1\xa6n\xaa\xba\xa5\xb4\xc3z\xb4\xf2\xa4H\xab\xa8 \xa4\xa3\xadn\xa9u\xab\xe1\xc1\xc9\xa4S\xa6\xdb\xa4v\xa5\xfd\xa4U\xa5h\xa4F        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbb\xa1\xa6n\xaa\xba\xa5\xb4\xc3z\xb4\xf2\xa4H\xab\xa8 \xa4\xa3\xadn\xa9u\xab\xe1\xc1\xc9\xa4S\xa6\xdb\xa4v\xa5\xfd\xa4U\xa5h\xa4F        "),
					},
				},
			},
			MD5:     "ZUTsAvYswh4FBsQRTmkI-Q",
			TheDate: "03/18 12:57",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mYouGot5566  \x1b[m\x1b[33m: \xbb\xa1\xa6n\xaa\xba\xa5\xb4\xc3z\xb4\xf2\xa4H\xab\xa8 \xa4\xa3\xadn\xa9u\xab\xe1\xc1\xc9\xa4S\xa6\xdb\xa4v\xa5\xfd\xa4U\xa5h\xa4F        \x1b[m 03/18 12:57\r"),
		},
		{ //23
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ooxxman"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "小胖子一減肥，小俠愈來愈有競爭力，MVP來了",
						Big5:   []byte("\xa4p\xadD\xa4l\xa4@\xb4\xee\xaa\xce\xa1A\xa4p\xabL\xb7U\xa8\xd3\xb7U\xa6\xb3\xc4v\xaa\xa7\xa4O\xa1AMVP\xa8\xd3\xa4F        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4p\xadD\xa4l\xa4@\xb4\xee\xaa\xce\xa1A\xa4p\xabL\xb7U\xa8\xd3\xb7U\xa6\xb3\xc4v\xaa\xa7\xa4O\xa1AMVP\xa8\xd3\xa4F        "),
					},
				},
			},
			MD5:     "UXc9DQX55hWFXaH1sfGO3A",
			TheDate: "03/18 12:58",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mooxxman     \x1b[m\x1b[33m: \xa4p\xadD\xa4l\xa4@\xb4\xee\xaa\xce\xa1A\xa4p\xabL\xb7U\xa8\xd3\xb7U\xa6\xb3\xc4v\xaa\xa7\xa4O\xa1AMVP\xa8\xd3\xa4F        \x1b[m 03/18 12:58\r"),
		},
		{ //24
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("oo98560"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "罰球不到60％? 這也叫職業？ 阿不是很愛嘴p＋",
						Big5:   []byte("\xbb@\xb2y\xa4\xa3\xa8\xec60\xa2H? \xb3o\xa4]\xa5s\xc2\xbe\xb7~\xa1H \xaa\xfc\xa4\xa3\xacO\xab\xdc\xb7R\xbcLp\xa1\xcf       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbb@\xb2y\xa4\xa3\xa8\xec60\xa2H? \xb3o\xa4]\xa5s\xc2\xbe\xb7~\xa1H \xaa\xfc\xa4\xa3\xacO\xab\xdc\xb7R\xbcLp\xa1\xcf       "),
					},
				},
			},
			MD5:     "ejp1VCqJUeAteHxvAdkE7g",
			TheDate: "03/18 12:59",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33moo98560     \x1b[m\x1b[33m: \xbb@\xb2y\xa4\xa3\xa8\xec60\xa2H? \xb3o\xa4]\xa5s\xc2\xbe\xb7~\xa1H \xaa\xfc\xa4\xa3\xacO\xab\xdc\xb7R\xbcLp\xa1\xcf       \x1b[m 03/18 12:59\r"),
		},
		{ //25
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("iNicholas"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "歡迎可愛來我牛",
						Big5:   []byte("\xc5w\xaa\xef\xa5i\xb7R\xa8\xd3\xa7\xda\xa4\xfb                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc5w\xaa\xef\xa5i\xb7R\xa8\xd3\xa7\xda\xa4\xfb                                   "),
					},
				},
			},
			MD5:     "sy9zuwnRLuFju0TwM7qleQ",
			TheDate: "03/18 13:04",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33miNicholas   \x1b[m\x1b[33m: \xc5w\xaa\xef\xa5i\xb7R\xa8\xd3\xa7\xda\xa4\xfb                                   \x1b[m 03/18 13:04\r"),
		},
		{ //26
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("gn00152097"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "小犢加溫很久了XD先不說明星賽前的連勝 明星賽後贏",
						Big5:   []byte("\xa4p\xc3}\xa5[\xb7\xc5\xab\xdc\xa4[\xa4FXD\xa5\xfd\xa4\xa3\xbb\xa1\xa9\xfa\xacP\xc1\xc9\xabe\xaa\xba\xb3s\xb3\xd3 \xa9\xfa\xacP\xc1\xc9\xab\xe1\xc4\xb9  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4p\xc3}\xa5[\xb7\xc5\xab\xdc\xa4[\xa4FXD\xa5\xfd\xa4\xa3\xbb\xa1\xa9\xfa\xacP\xc1\xc9\xabe\xaa\xba\xb3s\xb3\xd3 \xa9\xfa\xacP\xc1\xc9\xab\xe1\xc4\xb9  "),
					},
				},
			},
			MD5:     "wEnNOVHuDeaGSfSWOmDD9A",
			TheDate: "03/18 13:06",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mgn00152097  \x1b[m\x1b[33m: \xa4p\xc3}\xa5[\xb7\xc5\xab\xdc\xa4[\xa4FXD\xa5\xfd\xa4\xa3\xbb\xa1\xa9\xfa\xacP\xc1\xc9\xabe\xaa\xba\xb3s\xb3\xd3 \xa9\xfa\xacP\xc1\xc9\xab\xe1\xc4\xb9  \x1b[m 03/18 13:06\r"),
		},
		{ //27
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("gn00152097"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "了馬刺金塊快艇 全是排名在前面的",
						Big5:   []byte("\xa4F\xb0\xa8\xa8\xeb\xaa\xf7\xb6\xf4\xa7\xd6\xb8\xa5 \xa5\xfe\xacO\xb1\xc6\xa6W\xa6b\xabe\xad\xb1\xaa\xba                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4F\xb0\xa8\xa8\xeb\xaa\xf7\xb6\xf4\xa7\xd6\xb8\xa5 \xa5\xfe\xacO\xb1\xc6\xa6W\xa6b\xabe\xad\xb1\xaa\xba                  "),
					},
				},
			},
			MD5:     "NVbEccsO7LnVSY_27F_cUA",
			TheDate: "03/18 13:06",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mgn00152097  \x1b[m\x1b[33m: \xa4F\xb0\xa8\xa8\xeb\xaa\xf7\xb6\xf4\xa7\xd6\xb8\xa5 \xa5\xfe\xacO\xb1\xc6\xa6W\xa6b\xabe\xad\xb1\xaa\xba                  \x1b[m 03/18 13:06\r"),
		},
		{ //28
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Pacers31"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "PG上場時間怎麼又開始操了  之前不是說有傷要控制?",
						Big5:   []byte("PG\xa4W\xb3\xf5\xae\xc9\xb6\xa1\xab\xe7\xbb\xf2\xa4S\xb6}\xa9l\xbe\xde\xa4F  \xa4\xa7\xabe\xa4\xa3\xacO\xbb\xa1\xa6\xb3\xb6\xcb\xadn\xb1\xb1\xa8\xee?  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("PG\xa4W\xb3\xf5\xae\xc9\xb6\xa1\xab\xe7\xbb\xf2\xa4S\xb6}\xa9l\xbe\xde\xa4F  \xa4\xa7\xabe\xa4\xa3\xacO\xbb\xa1\xa6\xb3\xb6\xcb\xadn\xb1\xb1\xa8\xee?  "),
					},
				},
			},
			MD5:     "LFJp_NKgSquSnPM3TgnkrA",
			TheDate: "03/18 13:07",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mPacers31    \x1b[m\x1b[33m: PG\xa4W\xb3\xf5\xae\xc9\xb6\xa1\xab\xe7\xbb\xf2\xa4S\xb6}\xa9l\xbe\xde\xa4F  \xa4\xa7\xabe\xa4\xa3\xacO\xbb\xa1\xa6\xb3\xb6\xcb\xadn\xb1\xb1\xa8\xee?  \x1b[m 03/18 13:07\r"),
		},
		{ //29
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("JasonIm"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "mvp",
						Big5:   []byte("mvp                                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("mvp                                              "),
					},
				},
			},
			MD5:     "UYFBYcoqJcOetb0NvYRaSA",
			TheDate: "03/18 13:16",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mJasonIm     \x1b[m\x1b[33m: mvp                                              \x1b[m 03/18 13:16\r"),
		},
		{ //30
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("DavidFoster"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "打爆快艇水啦",
						Big5:   []byte("\xa5\xb4\xc3z\xa7\xd6\xb8\xa5\xa4\xf4\xb0\xd5                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5\xb4\xc3z\xa7\xd6\xb8\xa5\xa4\xf4\xb0\xd5                                     "),
					},
				},
			},
			MD5:     "_LvHtH4CrcEFp2irMSrFMQ",
			TheDate: "03/18 13:16",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDavidFoster \x1b[m\x1b[33m: \xa5\xb4\xc3z\xa7\xd6\xb8\xa5\xa4\xf4\xb0\xd5                                     \x1b[m 03/18 13:16\r"),
		},
		{ //31
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("pippen2002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "囧爆",
						Big5:   []byte("\xca\xa8\xc3z                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xca\xa8\xc3z                                             "),
					},
				},
			},
			MD5:     "EO64jiCn_iKGWCoJPRryag",
			TheDate: "03/18 13:18",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mpippen2002  \x1b[m\x1b[33m: \xca\xa8\xc3z                                             \x1b[m 03/18 13:18\r"),
		},
		{ //32
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("remsuki"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777777",
						Big5:   []byte("77777777777                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777777                                      "),
					},
				},
			},
			MD5:     "V0xryI-8F1kjRml41TbvPw",
			TheDate: "03/18 13:18",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mremsuki     \x1b[m\x1b[33m: 77777777777                                      \x1b[m 03/18 13:18\r"),
		},
		{ //33
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("dchiu815"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "快艇那個Mann怎麼每一次都會跟77起爭議",
						Big5:   []byte("\xa7\xd6\xb8\xa5\xa8\xba\xad\xd3Mann\xab\xe7\xbb\xf2\xa8C\xa4@\xa6\xb8\xb3\xa3\xb7|\xb8\xf277\xb0_\xaa\xa7\xc4\xb3             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xd6\xb8\xa5\xa8\xba\xad\xd3Mann\xab\xe7\xbb\xf2\xa8C\xa4@\xa6\xb8\xb3\xa3\xb7|\xb8\xf277\xb0_\xaa\xa7\xc4\xb3             "),
					},
				},
			},
			MD5:     "WW6KMFt5wBBN80ew_1rdoA",
			TheDate: "03/18 13:23",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mdchiu815    \x1b[m\x1b[33m: \xa7\xd6\xb8\xa5\xa8\xba\xad\xd3Mann\xab\xe7\xbb\xf2\xa8C\xa4@\xa6\xb8\xb3\xa3\xb7|\xb8\xf277\xb0_\xaa\xa7\xc4\xb3             \x1b[m 03/18 13:23\r"),
		},
		{ //34
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("jardon"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "kp進攻還是造成快艇很大的麻煩好嗎",
						Big5:   []byte("kp\xb6i\xa7\xf0\xc1\xd9\xacO\xb3y\xa6\xa8\xa7\xd6\xb8\xa5\xab\xdc\xa4j\xaa\xba\xb3\xc2\xb7\xd0\xa6n\xb6\xdc                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("kp\xb6i\xa7\xf0\xc1\xd9\xacO\xb3y\xa6\xa8\xa7\xd6\xb8\xa5\xab\xdc\xa4j\xaa\xba\xb3\xc2\xb7\xd0\xa6n\xb6\xdc                 "),
					},
				},
			},
			MD5:     "_sqDrgglfO95jB7mR4siLw",
			TheDate: "03/18 13:32",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mjardon      \x1b[m\x1b[33m: kp\xb6i\xa7\xf0\xc1\xd9\xacO\xb3y\xa6\xa8\xa7\xd6\xb8\xa5\xab\xdc\xa4j\xaa\xba\xb3\xc2\xb7\xd0\xa6n\xb6\xdc                 \x1b[m 03/18 13:32\r"),
		},
		{ //35
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("jardon"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "進攻沒那麼開 防守有在不就是小牛最需要的",
						Big5:   []byte("\xb6i\xa7\xf0\xa8S\xa8\xba\xbb\xf2\xb6} \xa8\xbe\xa6u\xa6\xb3\xa6b\xa4\xa3\xb4N\xacO\xa4p\xa4\xfb\xb3\xcc\xbb\xdd\xadn\xaa\xba          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb6i\xa7\xf0\xa8S\xa8\xba\xbb\xf2\xb6} \xa8\xbe\xa6u\xa6\xb3\xa6b\xa4\xa3\xb4N\xacO\xa4p\xa4\xfb\xb3\xcc\xbb\xdd\xadn\xaa\xba          "),
					},
				},
			},
			MD5:     "xgNuPTLtD2sLrBuUjjvNNQ",
			TheDate: "03/18 13:33",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mjardon      \x1b[m\x1b[33m: \xb6i\xa7\xf0\xa8S\xa8\xba\xbb\xf2\xb6} \xa8\xbe\xa6u\xa6\xb3\xa6b\xa4\xa3\xb4N\xacO\xa4p\xa4\xfb\xb3\xcc\xbb\xdd\xadn\xaa\xba          \x1b[m 03/18 13:33\r"),
		},
		{ //36
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("mavnokia"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "KP這幾場爛投少和帶球切、中投變多很關鍵",
						Big5:   []byte("KP\xb3o\xb4X\xb3\xf5\xc4\xea\xa7\xeb\xa4\xd6\xa9M\xb1a\xb2y\xa4\xc1\xa1B\xa4\xa4\xa7\xeb\xc5\xdc\xa6h\xab\xdc\xc3\xf6\xc1\xe4           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("KP\xb3o\xb4X\xb3\xf5\xc4\xea\xa7\xeb\xa4\xd6\xa9M\xb1a\xb2y\xa4\xc1\xa1B\xa4\xa4\xa7\xeb\xc5\xdc\xa6h\xab\xdc\xc3\xf6\xc1\xe4           "),
					},
				},
			},
			MD5:     "MGJEtJXEQ3JVGMZ-ahBHbw",
			TheDate: "03/18 13:34",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mmavnokia    \x1b[m\x1b[33m: KP\xb3o\xb4X\xb3\xf5\xc4\xea\xa7\xeb\xa4\xd6\xa9M\xb1a\xb2y\xa4\xc1\xa1B\xa4\xa4\xa7\xeb\xc5\xdc\xa6h\xab\xdc\xc3\xf6\xc1\xe4           \x1b[m 03/18 13:34\r"),
		},
		{ //37
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("SwissMiniGun"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "如果心態是領高薪就要全包 那大概沒幾個經得起操",
						Big5:   []byte("\xa6p\xaaG\xa4\xdf\xbaA\xacO\xbb\xe2\xb0\xaa\xc1~\xb4N\xadn\xa5\xfe\xa5] \xa8\xba\xa4j\xb7\xa7\xa8S\xb4X\xad\xd3\xb8g\xb1o\xb0_\xbe\xde    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6p\xaaG\xa4\xdf\xbaA\xacO\xbb\xe2\xb0\xaa\xc1~\xb4N\xadn\xa5\xfe\xa5] \xa8\xba\xa4j\xb7\xa7\xa8S\xb4X\xad\xd3\xb8g\xb1o\xb0_\xbe\xde    "),
					},
				},
			},
			MD5:     "zaqtlDlgqsM5DMAePyVgqA",
			TheDate: "03/18 13:35",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mSwissMiniGun\x1b[m\x1b[33m: \xa6p\xaaG\xa4\xdf\xbaA\xacO\xbb\xe2\xb0\xaa\xc1~\xb4N\xadn\xa5\xfe\xa5] \xa8\xba\xa4j\xb7\xa7\xa8S\xb4X\xad\xd3\xb8g\xb1o\xb0_\xbe\xde    \x1b[m 03/18 13:35\r"),
		},
		{ //38
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SwissMiniGun"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "當你交易來一個不耐戰的 本來就不該攻防都想操他",
						Big5:   []byte("\xb7\xed\xa7A\xa5\xe6\xa9\xf6\xa8\xd3\xa4@\xad\xd3\xa4\xa3\xad@\xbe\xd4\xaa\xba \xa5\xbb\xa8\xd3\xb4N\xa4\xa3\xb8\xd3\xa7\xf0\xa8\xbe\xb3\xa3\xb7Q\xbe\xde\xa5L    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb7\xed\xa7A\xa5\xe6\xa9\xf6\xa8\xd3\xa4@\xad\xd3\xa4\xa3\xad@\xbe\xd4\xaa\xba \xa5\xbb\xa8\xd3\xb4N\xa4\xa3\xb8\xd3\xa7\xf0\xa8\xbe\xb3\xa3\xb7Q\xbe\xde\xa5L    "),
					},
				},
			},
			MD5:     "YPaZNdGx6b0WoocwYEYC5w",
			TheDate: "03/18 13:36",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSwissMiniGun\x1b[m\x1b[33m: \xb7\xed\xa7A\xa5\xe6\xa9\xf6\xa8\xd3\xa4@\xad\xd3\xa4\xa3\xad@\xbe\xd4\xaa\xba \xa5\xbb\xa8\xd3\xb4N\xa4\xa3\xb8\xd3\xa7\xf0\xa8\xbe\xb3\xa3\xb7Q\xbe\xde\xa5L    \x1b[m 03/18 13:36\r"),
		},
		{ //39
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SwissMiniGun"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "適當的讓他發力才能讓他保持上場做出貢獻",
						Big5:   []byte("\xbeA\xb7\xed\xaa\xba\xc5\xfd\xa5L\xb5o\xa4O\xa4~\xaf\xe0\xc5\xfd\xa5L\xabO\xab\xf9\xa4W\xb3\xf5\xb0\xb5\xa5X\xb0^\xc4m           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbeA\xb7\xed\xaa\xba\xc5\xfd\xa5L\xb5o\xa4O\xa4~\xaf\xe0\xc5\xfd\xa5L\xabO\xab\xf9\xa4W\xb3\xf5\xb0\xb5\xa5X\xb0^\xc4m           "),
					},
				},
			},
			MD5:     "t1pxMdHdPAkXULHcUNscjQ",
			TheDate: "03/18 13:36",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSwissMiniGun\x1b[m\x1b[33m: \xbeA\xb7\xed\xaa\xba\xc5\xfd\xa5L\xb5o\xa4O\xa4~\xaf\xe0\xc5\xfd\xa5L\xabO\xab\xf9\xa4W\xb3\xf5\xb0\xb5\xa5X\xb0^\xc4m           \x1b[m 03/18 13:36\r"),
		},
		{ //40
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("SwissMiniGun"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "經過上季沒那麼謹慎的使用 我牛這季總該學乖",
						Big5:   []byte("\xb8g\xb9L\xa4W\xa9u\xa8S\xa8\xba\xbb\xf2\xc2\xd4\xb7V\xaa\xba\xa8\xcf\xa5\xce \xa7\xda\xa4\xfb\xb3o\xa9u\xc1`\xb8\xd3\xbe\xc7\xa8\xc4        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb8g\xb9L\xa4W\xa9u\xa8S\xa8\xba\xbb\xf2\xc2\xd4\xb7V\xaa\xba\xa8\xcf\xa5\xce \xa7\xda\xa4\xfb\xb3o\xa9u\xc1`\xb8\xd3\xbe\xc7\xa8\xc4        "),
					},
				},
			},
			MD5:     "wGdHlTX0kotx84G0Fi72WA",
			TheDate: "03/18 13:37",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mSwissMiniGun\x1b[m\x1b[33m: \xb8g\xb9L\xa4W\xa9u\xa8S\xa8\xba\xbb\xf2\xc2\xd4\xb7V\xaa\xba\xa8\xcf\xa5\xce \xa7\xda\xa4\xfb\xb3o\xa9u\xc1`\xb8\xd3\xbe\xc7\xa8\xc4        \x1b[m 03/18 13:37\r"),
		},
		{ //41
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("kevin202025"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77重回MVP競爭行列",
						Big5:   []byte("77\xad\xab\xa6^MVP\xc4v\xaa\xa7\xa6\xe6\xa6C                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xad\xab\xa6^MVP\xc4v\xaa\xa7\xa6\xe6\xa6C                                "),
					},
				},
			},
			MD5:     "i05CDCp6bjuhaxKC7DP3nw",
			TheDate: "03/18 13:40",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mkevin202025 \x1b[m\x1b[33m: 77\xad\xab\xa6^MVP\xc4v\xaa\xa7\xa6\xe6\xa6C                                \x1b[m 03/18 13:40\r"),
		},
		{ //42
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("Tawara"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "可愛：這船 不太對",
						Big5:   []byte("\xa5i\xb7R\xa1G\xb3o\xb2\xee \xa4\xa3\xa4\xd3\xb9\xef                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5i\xb7R\xa1G\xb3o\xb2\xee \xa4\xa3\xa4\xd3\xb9\xef                                "),
					},
				},
			},
			MD5:     "c9SqN5m7XDGhY3Aye0PB0A",
			TheDate: "03/18 13:41",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mTawara      \x1b[m\x1b[33m: \xa5i\xb7R\xa1G\xb3o\xb2\xee \xa4\xa3\xa4\xd3\xb9\xef                                \x1b[m 03/18 13:41\r"),
		},
		{ //43
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ruiun"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "快艇最近是不是在藏啊，打算季後賽才加速",
						Big5:   []byte("\xa7\xd6\xb8\xa5\xb3\xcc\xaa\xf1\xacO\xa4\xa3\xacO\xa6b\xc2\xc3\xb0\xda\xa1A\xa5\xb4\xba\xe2\xa9u\xab\xe1\xc1\xc9\xa4~\xa5[\xb3t           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xd6\xb8\xa5\xb3\xcc\xaa\xf1\xacO\xa4\xa3\xacO\xa6b\xc2\xc3\xb0\xda\xa1A\xa5\xb4\xba\xe2\xa9u\xab\xe1\xc1\xc9\xa4~\xa5[\xb3t           "),
					},
				},
			},
			MD5:     "LBuOhmLE3NfnAO8eoVnpVw",
			TheDate: "03/18 13:43",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mruiun       \x1b[m\x1b[33m: \xa7\xd6\xb8\xa5\xb3\xcc\xaa\xf1\xacO\xa4\xa3\xacO\xa6b\xc2\xc3\xb0\xda\xa1A\xa5\xb4\xba\xe2\xa9u\xab\xe1\xc1\xc9\xa4~\xa5[\xb3t           \x1b[m 03/18 13:43\r"),
		},
		{ //44
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("cccman"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777",
						Big5:   []byte("7777777                                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777                                          "),
					},
				},
			},
			MD5:     "5Fo0L-CExiKxJ8DLmIRM-w",
			TheDate: "03/18 13:51",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mcccman      \x1b[m\x1b[33m: 7777777                                          \x1b[m 03/18 13:51\r"),
		},
		{ //45
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("providence"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77今天跟鬼一樣",
						Big5:   []byte("77\xa4\xb5\xa4\xd1\xb8\xf2\xb0\xad\xa4@\xbc\xcb                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77\xa4\xb5\xa4\xd1\xb8\xf2\xb0\xad\xa4@\xbc\xcb                                   "),
					},
				},
			},
			MD5:     "kZkjmvFrpXL-hSFIkkfyRw",
			TheDate: "03/18 13:58",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mprovidence  \x1b[m\x1b[33m: 77\xa4\xb5\xa4\xd1\xb8\xf2\xb0\xad\xa4@\xbc\xcb                                   \x1b[m 03/18 13:58\r"),
		},
		{ //46
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("age200"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "東77是鬼吧",
						Big5:   []byte("\xaaF77\xacO\xb0\xad\xa7a                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaaF77\xacO\xb0\xad\xa7a                                       "),
					},
				},
			},
			MD5:     "oD9CLBN0Sja3cDKBkiv4Ag",
			TheDate: "03/18 14:00",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mage200      \x1b[m\x1b[33m: \xaaF77\xacO\xb0\xad\xa7a                                       \x1b[m 03/18 14:00\r"),
		},
		{ //47
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("zold"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "77777777777777777",
						Big5:   []byte("77777777777777777                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("77777777777777777                                "),
					},
				},
			},
			MD5:     "nnHHFGegTCV9lL8kItKYvw",
			TheDate: "03/19 00:31",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mzold        \x1b[m\x1b[33m: 77777777777777777                                \x1b[m 03/19 00:31\r"),
		},
		{ //48
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("fantastic96"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "7777777777777",
						Big5:   []byte("7777777777777                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("7777777777777                                    "),
					},
				},
			},
			MD5:     "0uwavwBPcAcoYqzaSaBlxg",
			TheDate: "03/19 11:08",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mfantastic96 \x1b[m\x1b[33m: 7777777777777                                    \x1b[m 03/19 11:08\r"),
		},
	}
}
