package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename9            = "temp4"
	testContentAll9          []byte
	testContent9             []byte
	testSignature9           []byte
	testComment9             []byte
	testFirstCommentsDBCS9   []byte
	testTheRestCommentsDBCS9 []byte

	testContent9Big5 [][]*types.Rune
	testContent9Utf8 [][]*types.Rune

	testFirstComments9     []*schema.Comment
	testFullFirstComments9 []*schema.Comment
)

func initTest9() {
	testContentAll9, testContent9, testSignature9, testComment9, testFirstCommentsDBCS9, testTheRestCommentsDBCS9 = loadTest(testFilename9)

	testContent9Big5 = [][]*types.Rune{

		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: xxFrency (\xaa\xfc\xad\xf5) \xac\xdd\xaaO: CodeJob"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: xxFrency (\xaa\xfc\xad\xf5) \xac\xdd\xaaO: CodeJob\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] \xb1\xd0\xbe\xc7 \xb6\xb3\xba\xdd\xa5D\xbe\xf7\xc1\xca\xaa\xab\xba\xf4\xad\xb6\xba\xfb\xb9B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] \xb1\xd0\xbe\xc7 \xb6\xb3\xba\xdd\xa5D\xbe\xf7\xc1\xca\xaa\xab\xba\xf4\xad\xb6\xba\xfb\xb9B\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Dec 19 14:46:08 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sat Dec 19 14:46:08 2020\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
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

				Big5:   []byte("\xa1G\xbcB\xa5\xfd\xa5\xcd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xbcB\xa5\xfd\xa5\xcd\r"),
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

				Big5:   []byte("\xc1p\xb5\xb8\xa4\xe8\xa6\xa11"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xc1p\xb5\xb8\xa4\xe8\xa6\xa11"),
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

				Big5:   []byte("       \xc1p\xb5\xb8\xa4\xe8\xa6\xa12\xa1Gaaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       \xc1p\xb5\xb8\xa4\xe8\xa6\xa12\xa1Gaaaaaaaa\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1G\xa5x\xa5_\xa5\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1G\xa5x\xa5_\xa5\xab\r"),
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

				Big5:   []byte("\xa1G12/31"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G12/31\r"),
			},
		},
		{ //11
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

				Big5:   []byte("\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\r"),
			},
		},
		{ //12
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //13
			{

				Big5:   []byte("               \xaa\xac\xaap\xb4y\xadz\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xaa\xac\xaap\xb4y\xadz\xa1G\r"),
			},
		},
		{ //14
			{

				Big5:   []byte("               \xa5\xd8\xabe\xa4w\xb8g\xa6\xb3\xba\xf4\xaf\xb8\xb3\xa1\xc4\xdd\xa6b\xaa\xfc\xa8\xbd\xb6\xb3\xa4W(LNMP\xa1A\xa8t\xb2\xceCentOS 7.7)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xa5\xd8\xabe\xa4w\xb8g\xa6\xb3\xba\xf4\xaf\xb8\xb3\xa1\xc4\xdd\xa6b\xaa\xfc\xa8\xbd\xb6\xb3\xa4W(LNMP\xa1A\xa8t\xb2\xceCentOS 7.7)\r"),
			},
		},
		{ //15
			{

				Big5:   []byte("               \xb1\xb5\xa4U\xa8\xd3\xa5i\xaf\xe0\xb7|\xad\xb1\xc1{\xa4j\xb6q\xacy\xb6q\xb3y\xb3X\xa1A\xb7Q\xb9\xef\xa6\xf8\xaaA\xbe\xb9\xb0\xb5\xc0u\xa4\xc6\xa1B\xb0t\xb8m\xa1B\xad\xb0\xa7C\xa5i\xaf\xe0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xb1\xb5\xa4U\xa8\xd3\xa5i\xaf\xe0\xb7|\xad\xb1\xc1{\xa4j\xb6q\xacy\xb6q\xb3y\xb3X\xa1A\xb7Q\xb9\xef\xa6\xf8\xaaA\xbe\xb9\xb0\xb5\xc0u\xa4\xc6\xa1B\xb0t\xb8m\xa1B\xad\xb0\xa7C\xa5i\xaf\xe0\r"),
			},
		},
		{ //16
			{

				Big5:   []byte("               \xad\xb1\xc1{\xaa\xba\xad\xb7\xc0I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xad\xb1\xc1{\xaa\xba\xad\xb7\xc0I\r"),
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

				Big5:   []byte("\xa1G\xae\xc9\xc1~800\xb0_\xa1A\xa6p\xaaG\xc4\xb1\xb1o\xa4\xd3\xa7C\xa1A\xc5w\xaa\xef\xb1H\xabH\xa7i\xb6D\xa7\xda\xa7A\xaa\xba\xb8g"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xae\xc9\xc1~800\xb0_\xa1A\xa6p\xaaG\xc4\xb1\xb1o\xa4\xd3\xa7C\xa1A\xc5w\xaa\xef\xb1H\xabH\xa7i\xb6D\xa7\xda\xa7A\xaa\xba\xb8g\r"),
			},
		},
		{ //19
			{

				Big5:   []byte("                               \xc5\xe7\xa1A\xa7\xda\xb8\xf2\xa4\xbd\xa5q\xb0\xd3\xb6q\xb3q\xb9L\xab\xe1\xb4N\xa5i\xa5H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                               \xc5\xe7\xa1A\xa7\xda\xb8\xf2\xa4\xbd\xa5q\xb0\xd3\xb6q\xb3q\xb9L\xab\xe1\xb4N\xa5i\xa5H\r"),
			},
		},
		{ //20
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //21
			{

				Big5:   []byte("                \xa5\xb2\xb6\xb7\xb6\xf1\xbcg\xb9w\xba\xe2\xbdd\xb3\xf2\xa1A\xadY\xb5L\xb6\xf1\xbcg\xb9H\xa4\xcf\xaaO\xb3W"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa5\xb2\xb6\xb7\xb6\xf1\xbcg\xb9w\xba\xe2\xbdd\xb3\xf2\xa1A\xadY\xb5L\xb6\xf1\xbcg\xb9H\xa4\xcf\xaaO\xb3W\r"),
			},
		},
		{ //22
			{

				Big5:   []byte("            \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbe\xd6\xa6\xb3\xaa\xfc\xa8\xbd\xb6\xb3\xa9\xceAWS\xb9\xea\xbb\xda\xbe\xde\xa7@\xb8g\xc5\xe7\xa1A\xa8\xc3\xb4\xbf\xb8g\xb3B\xb2z\xb9L\xc0\xfe\xb6\xa1\xa4j\xb6q\xacy\xaa\xba"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("            \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbe\xd6\xa6\xb3\xaa\xfc\xa8\xbd\xb6\xb3\xa9\xceAWS\xb9\xea\xbb\xda\xbe\xde\xa7@\xb8g\xc5\xe7\xa1A\xa8\xc3\xb4\xbf\xb8g\xb3B\xb2z\xb9L\xc0\xfe\xb6\xa1\xa4j\xb6q\xacy\xaa\xba\r"),
			},
		},
		{ //23
			{

				Big5:   []byte("            \xaa\xac\xaap"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("            \xaa\xac\xaap\r"),
			},
		},
		{ //24
			{

				Big5:   []byte("        \xa1@\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        \xa1@\xa1@\r"),
			},
		},
		{ //25
			{

				Big5:   []byte("            \xaa\xfe\xb5\xf9\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("            \xaa\xfe\xb5\xf9\xa1G\r"),
			},
		},
		{ //26
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //27
			{

				Big5:   []byte("                \xb1\xd0\xbe\xc7\xad\xab\xc2I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xb1\xd0\xbe\xc7\xad\xab\xc2I\r"),
			},
		},
		{ //28
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //29
			{

				Big5:   []byte("                \xa4W\xbd\xd2\xae\xc9\xb6\xa1\xa1G\xa4@\xb6g1\xa1\xe32\xa6\xb8\xa1A\xa5i\xb0Q\xbd\xd7\xa1A\xa4@\xa6\xb8\xac\xf91~2\xa4p\xae\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa4W\xbd\xd2\xae\xc9\xb6\xa1\xa1G\xa4@\xb6g1\xa1\xe32\xa6\xb8\xa1A\xa5i\xb0Q\xbd\xd7\xa1A\xa4@\xa6\xb8\xac\xf91~2\xa4p\xae\xc9\r"),
			},
		},
		{ //30
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //31
			{

				Big5:   []byte("                1.\xba\xf4\xaf\xb8\xb9B\xa6\xe6\xac[\xbac\xab\xd8\xc4\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                1.\xba\xf4\xaf\xb8\xb9B\xa6\xe6\xac[\xbac\xab\xd8\xc4\xb3\r"),
			},
		},
		{ //32
			{

				Big5:   []byte("                (\xa8\xd2\xa6p\xa1GMySql\xaa\xbd\xb1\xb5\xb9B\xa6\xe6\xa6b\xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa4W\xa9\xce\xa8\xcf\xa5\xceRDS\xa1A\xa4\xb0\xbb\xf2\xb1\xa1\xaap\xb8\xd3\xa8\xcf\xa5\xceRDS)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                (\xa8\xd2\xa6p\xa1GMySql\xaa\xbd\xb1\xb5\xb9B\xa6\xe6\xa6b\xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa4W\xa9\xce\xa8\xcf\xa5\xceRDS\xa1A\xa4\xb0\xbb\xf2\xb1\xa1\xaap\xb8\xd3\xa8\xcf\xa5\xceRDS)\r"),
			},
		},
		{ //33
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //34
			{

				Big5:   []byte("                2. Load Balancer\xadt\xb8\xfc\xa5\xad\xbf\xc5\xb3]\xa9w"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                2. Load Balancer\xadt\xb8\xfc\xa5\xad\xbf\xc5\xb3]\xa9w\r"),
			},
		},
		{ //35
			{

				Big5:   []byte("                \xaa\xbd\xb1\xb5\xb6i\xaa\xfc\xa8\xbd\xb6\xb3\xab\xe1\xa5x\xa1A\xa4\xe2\xa7\xe2\xa4\xe2\xb1\xd0\xa7\xda\xb3]\xa9wLoad Balancer\xa1A\xa6p\xaaG\xa7A\xa5u\xa6\xb3AWS"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xaa\xbd\xb1\xb5\xb6i\xaa\xfc\xa8\xbd\xb6\xb3\xab\xe1\xa5x\xa1A\xa4\xe2\xa7\xe2\xa4\xe2\xb1\xd0\xa7\xda\xb3]\xa9wLoad Balancer\xa1A\xa6p\xaaG\xa7A\xa5u\xa6\xb3AWS\r"),
			},
		},
		{ //36
			{

				Big5:   []byte("                Load Balancer\xb8g\xc5\xe7\xa4]\xa5i\xa5H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                Load Balancer\xb8g\xc5\xe7\xa4]\xa5i\xa5H\r"),
			},
		},
		{ //37
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //38
			{

				Big5:   []byte("                \xa8\xe4\xa5L\xb0\xdd\xc3D\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa8\xe4\xa5L\xb0\xdd\xc3D\xa1G\r"),
			},
		},
		{ //39
			{

				Big5:   []byte("\xa1@\xa1@                    -\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb8\xd3\xa8\xcf\xa5\xceLoad Balancer \xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j\xb4N\xa6n"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@                    -\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb8\xd3\xa8\xcf\xa5\xceLoad Balancer \xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j\xb4N\xa6n\r"),
			},
		},
		{ //40
			{

				Big5:   []byte("\xa1@\xa1@                    -\xa7\xda\xb8\xd3\xab\xe7\xbb\xf2\xaa\xbe\xb9D\xbe\xf7\xbe\xb9\xaa\xba\xa9\xd3\xb8\xfc\xaf\xe0\xa4O\xa1A\xacO\xa7_\xbb\xdd\xadn\xb0\xb5\xc0\xa3\xa4O\xb4\xfa\xb8\xd5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@                    -\xa7\xda\xb8\xd3\xab\xe7\xbb\xf2\xaa\xbe\xb9D\xbe\xf7\xbe\xb9\xaa\xba\xa9\xd3\xb8\xfc\xaf\xe0\xa4O\xa1A\xacO\xa7_\xbb\xdd\xadn\xb0\xb5\xc0\xa3\xa4O\xb4\xfa\xb8\xd5\r"),
			},
		},
		{ //41
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //42
			{

				Big5:   []byte("                3.\xa6\xdb\xb0\xca\xa9w\xb4\xc1\xb3\xc6\xa5\xf7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                3.\xa6\xdb\xb0\xca\xa9w\xb4\xc1\xb3\xc6\xa5\xf7\r"),
			},
		},
		{ //43
			{

				Big5:   []byte("                \xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa5i\xa5H\xa7\xd6\xb7\xd3\xb3\xc6\xa5\xf7\xa1A\xa6\xfd\xa6p\xaaG\xacORDS+ECS\xaa\xba\xaa\xac\xaap\xb8\xd3\xab\xe7\xbb\xf2\xb3\xc6\xa5\xf7\xa4\xf1\xb8\xfb\xa6X\xbeA"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa5i\xa5H\xa7\xd6\xb7\xd3\xb3\xc6\xa5\xf7\xa1A\xa6\xfd\xa6p\xaaG\xacORDS+ECS\xaa\xba\xaa\xac\xaap\xb8\xd3\xab\xe7\xbb\xf2\xb3\xc6\xa5\xf7\xa4\xf1\xb8\xfb\xa6X\xbeA\r"),
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

				Big5:   []byte("                4. \xb1\xd0\xa7\xda\xa8t\xb2\xce\xba\xca\xb1\xb1\xaa\xba\xb9\xee\xac\xdd\xad\xab\xc2I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                4. \xb1\xd0\xa7\xda\xa8t\xb2\xce\xba\xca\xb1\xb1\xaa\xba\xb9\xee\xac\xdd\xad\xab\xc2I\r"),
			},
		},
		{ //46
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //47
			{

				Big5:   []byte("                5. \xa6w\xa5\xfe\xa9\xca\xa4\xce\xa8\xbe\xa4\xf5\xc0\xf0\xb3]\xb8m\xb1\xd0\xbe\xc7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                5. \xa6w\xa5\xfe\xa9\xca\xa4\xce\xa8\xbe\xa4\xf5\xc0\xf0\xb3]\xb8m\xb1\xd0\xbe\xc7\r"),
			},
		},
		{ //48
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //49
			{

				Big5:   []byte("                6. (\xabD\xa5\xb2\xadn)NGINX\xba\xf4\xad\xb6\xa6\xf8\xaaA\xbe\xb9\xb1\xd0\xbe\xc7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                6. (\xabD\xa5\xb2\xadn)NGINX\xba\xf4\xad\xb6\xa6\xf8\xaaA\xbe\xb9\xb1\xd0\xbe\xc7\r"),
			},
		},
	}

	testContent9Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: xxFrency (阿哲) 看板: CodeJob",
				Big5:   []byte("\xa7@\xaa\xcc: xxFrency (\xaa\xfc\xad\xf5) \xac\xdd\xaaO: CodeJob"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: xxFrency (\xaa\xfc\xad\xf5) \xac\xdd\xaaO: CodeJob\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: [發案] 教學 雲端主機購物網頁維運",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] \xb1\xd0\xbe\xc7 \xb6\xb3\xba\xdd\xa5D\xbe\xf7\xc1\xca\xaa\xab\xba\xf4\xad\xb6\xba\xfb\xb9B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb5o\xae\xd7] \xb1\xd0\xbe\xc7 \xb6\xb3\xba\xdd\xa5D\xbe\xf7\xc1\xca\xaa\xab\xba\xf4\xad\xb6\xba\xfb\xb9B\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Sat Dec 19 14:46:08 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Dec 19 14:46:08 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sat Dec 19 14:46:08 2020\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
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
				Utf8:   "：劉先生",
				Big5:   []byte("\xa1G\xbcB\xa5\xfd\xa5\xcd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xbcB\xa5\xfd\xa5\xcd\r"),
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
				Utf8:   "聯絡方式1",
				Big5:   []byte("\xc1p\xb5\xb8\xa4\xe8\xa6\xa11"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35;40m\xc1p\xb5\xb8\xa4\xe8\xa6\xa11"),
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
				Utf8:   "       聯絡方式2：aaaaaaaa",
				Big5:   []byte("       \xc1p\xb5\xb8\xa4\xe8\xa6\xa12\xa1Gaaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       \xc1p\xb5\xb8\xa4\xe8\xa6\xa12\xa1Gaaaaaaaa\r"),
			},
		},
		{ //8
			{
				Utf8:   "       所在地區 ：台北市",
				Big5:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1G\xa5x\xa5_\xa5\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("       \xa9\xd2\xa6b\xa6a\xb0\xcf \xa1G\xa5x\xa5_\xa5\xab\r"),
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
				Utf8:   "：12/31",
				Big5:   []byte("\xa1G12/31"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G12/31\r"),
			},
		},
		{ //11
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
				Utf8:   "：",
				Big5:   []byte("\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\r"),
			},
		},
		{ //12
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //13
			{
				Utf8:   "               狀況描述：",
				Big5:   []byte("               \xaa\xac\xaap\xb4y\xadz\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xaa\xac\xaap\xb4y\xadz\xa1G\r"),
			},
		},
		{ //14
			{
				Utf8:   "               目前已經有網站部屬在阿里雲上(LNMP，系統CentOS 7.7)",
				Big5:   []byte("               \xa5\xd8\xabe\xa4w\xb8g\xa6\xb3\xba\xf4\xaf\xb8\xb3\xa1\xc4\xdd\xa6b\xaa\xfc\xa8\xbd\xb6\xb3\xa4W(LNMP\xa1A\xa8t\xb2\xceCentOS 7.7)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xa5\xd8\xabe\xa4w\xb8g\xa6\xb3\xba\xf4\xaf\xb8\xb3\xa1\xc4\xdd\xa6b\xaa\xfc\xa8\xbd\xb6\xb3\xa4W(LNMP\xa1A\xa8t\xb2\xceCentOS 7.7)\r"),
			},
		},
		{ //15
			{
				Utf8:   "               接下來可能會面臨大量流量造訪，想對伺服器做優化、配置、降低可能",
				Big5:   []byte("               \xb1\xb5\xa4U\xa8\xd3\xa5i\xaf\xe0\xb7|\xad\xb1\xc1{\xa4j\xb6q\xacy\xb6q\xb3y\xb3X\xa1A\xb7Q\xb9\xef\xa6\xf8\xaaA\xbe\xb9\xb0\xb5\xc0u\xa4\xc6\xa1B\xb0t\xb8m\xa1B\xad\xb0\xa7C\xa5i\xaf\xe0"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xb1\xb5\xa4U\xa8\xd3\xa5i\xaf\xe0\xb7|\xad\xb1\xc1{\xa4j\xb6q\xacy\xb6q\xb3y\xb3X\xa1A\xb7Q\xb9\xef\xa6\xf8\xaaA\xbe\xb9\xb0\xb5\xc0u\xa4\xc6\xa1B\xb0t\xb8m\xa1B\xad\xb0\xa7C\xa5i\xaf\xe0\r"),
			},
		},
		{ //16
			{
				Utf8:   "               面臨的風險",
				Big5:   []byte("               \xad\xb1\xc1{\xaa\xba\xad\xb7\xc0I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("               \xad\xb1\xc1{\xaa\xba\xad\xb7\xc0I\r"),
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
				Utf8:   "：時薪800起，如果覺得太低，歡迎寄信告訴我你的經",
				Big5:   []byte("\xa1G\xae\xc9\xc1~800\xb0_\xa1A\xa6p\xaaG\xc4\xb1\xb1o\xa4\xd3\xa7C\xa1A\xc5w\xaa\xef\xb1H\xabH\xa7i\xb6D\xa7\xda\xa7A\xaa\xba\xb8g"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1G\xae\xc9\xc1~800\xb0_\xa1A\xa6p\xaaG\xc4\xb1\xb1o\xa4\xd3\xa7C\xa1A\xc5w\xaa\xef\xb1H\xabH\xa7i\xb6D\xa7\xda\xa7A\xaa\xba\xb8g\r"),
			},
		},
		{ //19
			{

				Utf8:   "                               驗，我跟公司商量通過後就可以",
				Big5:   []byte("                               \xc5\xe7\xa1A\xa7\xda\xb8\xf2\xa4\xbd\xa5q\xb0\xd3\xb6q\xb3q\xb9L\xab\xe1\xb4N\xa5i\xa5H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                               \xc5\xe7\xa1A\xa7\xda\xb8\xf2\xa4\xbd\xa5q\xb0\xd3\xb6q\xb3q\xb9L\xab\xe1\xb4N\xa5i\xa5H\r"),
			},
		},
		{ //20
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //21
			{
				Utf8:   "                必須填寫預算範圍，若無填寫違反板規",
				Big5:   []byte("                \xa5\xb2\xb6\xb7\xb6\xf1\xbcg\xb9w\xba\xe2\xbdd\xb3\xf2\xa1A\xadY\xb5L\xb6\xf1\xbcg\xb9H\xa4\xcf\xaaO\xb3W"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa5\xb2\xb6\xb7\xb6\xf1\xbcg\xb9w\xba\xe2\xbdd\xb3\xf2\xa1A\xadY\xb5L\xb6\xf1\xbcg\xb9H\xa4\xcf\xaaO\xb3W\r"),
			},
		},
		{ //22
			{
				Utf8:   "            接案者要求：擁有阿里雲或AWS實際操作經驗，並曾經處理過瞬間大量流的",
				Big5:   []byte("            \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbe\xd6\xa6\xb3\xaa\xfc\xa8\xbd\xb6\xb3\xa9\xceAWS\xb9\xea\xbb\xda\xbe\xde\xa7@\xb8g\xc5\xe7\xa1A\xa8\xc3\xb4\xbf\xb8g\xb3B\xb2z\xb9L\xc0\xfe\xb6\xa1\xa4j\xb6q\xacy\xaa\xba"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("            \xb1\xb5\xae\xd7\xaa\xcc\xadn\xa8D\xa1G\xbe\xd6\xa6\xb3\xaa\xfc\xa8\xbd\xb6\xb3\xa9\xceAWS\xb9\xea\xbb\xda\xbe\xde\xa7@\xb8g\xc5\xe7\xa1A\xa8\xc3\xb4\xbf\xb8g\xb3B\xb2z\xb9L\xc0\xfe\xb6\xa1\xa4j\xb6q\xacy\xaa\xba\r"),
			},
		},
		{ //23
			{
				Utf8:   "            狀況",
				Big5:   []byte("            \xaa\xac\xaap"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("            \xaa\xac\xaap\r"),
			},
		},
		{ //24
			{
				Utf8:   "        　　",
				Big5:   []byte("        \xa1@\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("        \xa1@\xa1@\r"),
			},
		},
		{ //25
			{
				Utf8:   "            附註：",
				Big5:   []byte("            \xaa\xfe\xb5\xf9\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("            \xaa\xfe\xb5\xf9\xa1G\r"),
			},
		},
		{ //26
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //27
			{
				Utf8:   "                教學重點",
				Big5:   []byte("                \xb1\xd0\xbe\xc7\xad\xab\xc2I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xb1\xd0\xbe\xc7\xad\xab\xc2I\r"),
			},
		},
		{ //28
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //29
			{
				Utf8:   "                上課時間：一週1～2次，可討論，一次約1~2小時",
				Big5:   []byte("                \xa4W\xbd\xd2\xae\xc9\xb6\xa1\xa1G\xa4@\xb6g1\xa1\xe32\xa6\xb8\xa1A\xa5i\xb0Q\xbd\xd7\xa1A\xa4@\xa6\xb8\xac\xf91~2\xa4p\xae\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa4W\xbd\xd2\xae\xc9\xb6\xa1\xa1G\xa4@\xb6g1\xa1\xe32\xa6\xb8\xa1A\xa5i\xb0Q\xbd\xd7\xa1A\xa4@\xa6\xb8\xac\xf91~2\xa4p\xae\xc9\r"),
			},
		},
		{ //30
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //31
			{
				Utf8:   "                1.網站運行架構建議",
				Big5:   []byte("                1.\xba\xf4\xaf\xb8\xb9B\xa6\xe6\xac[\xbac\xab\xd8\xc4\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                1.\xba\xf4\xaf\xb8\xb9B\xa6\xe6\xac[\xbac\xab\xd8\xc4\xb3\r"),
			},
		},
		{ //32
			{
				Utf8:   "                (例如：MySql直接運行在同一台機器上或使用RDS，什麼情況該使用RDS)",
				Big5:   []byte("                (\xa8\xd2\xa6p\xa1GMySql\xaa\xbd\xb1\xb5\xb9B\xa6\xe6\xa6b\xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa4W\xa9\xce\xa8\xcf\xa5\xceRDS\xa1A\xa4\xb0\xbb\xf2\xb1\xa1\xaap\xb8\xd3\xa8\xcf\xa5\xceRDS)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                (\xa8\xd2\xa6p\xa1GMySql\xaa\xbd\xb1\xb5\xb9B\xa6\xe6\xa6b\xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa4W\xa9\xce\xa8\xcf\xa5\xceRDS\xa1A\xa4\xb0\xbb\xf2\xb1\xa1\xaap\xb8\xd3\xa8\xcf\xa5\xceRDS)\r"),
			},
		},
		{ //33
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //34
			{
				Utf8:   "                2. Load Balancer負載平衡設定",
				Big5:   []byte("                2. Load Balancer\xadt\xb8\xfc\xa5\xad\xbf\xc5\xb3]\xa9w"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                2. Load Balancer\xadt\xb8\xfc\xa5\xad\xbf\xc5\xb3]\xa9w\r"),
			},
		},
		{ //35
			{
				Utf8:   "                直接進阿里雲後台，手把手教我設定Load Balancer，如果你只有AWS",
				Big5:   []byte("                \xaa\xbd\xb1\xb5\xb6i\xaa\xfc\xa8\xbd\xb6\xb3\xab\xe1\xa5x\xa1A\xa4\xe2\xa7\xe2\xa4\xe2\xb1\xd0\xa7\xda\xb3]\xa9wLoad Balancer\xa1A\xa6p\xaaG\xa7A\xa5u\xa6\xb3AWS"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xaa\xbd\xb1\xb5\xb6i\xaa\xfc\xa8\xbd\xb6\xb3\xab\xe1\xa5x\xa1A\xa4\xe2\xa7\xe2\xa4\xe2\xb1\xd0\xa7\xda\xb3]\xa9wLoad Balancer\xa1A\xa6p\xaaG\xa7A\xa5u\xa6\xb3AWS\r"),
			},
		},
		{ //36
			{
				Utf8:   "                Load Balancer經驗也可以",
				Big5:   []byte("                Load Balancer\xb8g\xc5\xe7\xa4]\xa5i\xa5H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                Load Balancer\xb8g\xc5\xe7\xa4]\xa5i\xa5H\r"),
			},
		},
		{ //37
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //38
			{
				Utf8:   "                其他問題：",
				Big5:   []byte("                \xa8\xe4\xa5L\xb0\xdd\xc3D\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa8\xe4\xa5L\xb0\xdd\xc3D\xa1G\r"),
			},
		},
		{ //39
			{
				Utf8:   "　　                    -什麼時候該使用Load Balancer 什麼時候把機器加大就好",
				Big5:   []byte("\xa1@\xa1@                    -\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb8\xd3\xa8\xcf\xa5\xceLoad Balancer \xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j\xb4N\xa6n"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@                    -\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb8\xd3\xa8\xcf\xa5\xceLoad Balancer \xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j\xb4N\xa6n\r"),
			},
		},
		{ //40
			{
				Utf8:   "　　                    -我該怎麼知道機器的承載能力，是否需要做壓力測試",
				Big5:   []byte("\xa1@\xa1@                    -\xa7\xda\xb8\xd3\xab\xe7\xbb\xf2\xaa\xbe\xb9D\xbe\xf7\xbe\xb9\xaa\xba\xa9\xd3\xb8\xfc\xaf\xe0\xa4O\xa1A\xacO\xa7_\xbb\xdd\xadn\xb0\xb5\xc0\xa3\xa4O\xb4\xfa\xb8\xd5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@                    -\xa7\xda\xb8\xd3\xab\xe7\xbb\xf2\xaa\xbe\xb9D\xbe\xf7\xbe\xb9\xaa\xba\xa9\xd3\xb8\xfc\xaf\xe0\xa4O\xa1A\xacO\xa7_\xbb\xdd\xadn\xb0\xb5\xc0\xa3\xa4O\xb4\xfa\xb8\xd5\r"),
			},
		},
		{ //41
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //42
			{
				Utf8:   "                3.自動定期備份",
				Big5:   []byte("                3.\xa6\xdb\xb0\xca\xa9w\xb4\xc1\xb3\xc6\xa5\xf7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                3.\xa6\xdb\xb0\xca\xa9w\xb4\xc1\xb3\xc6\xa5\xf7\r"),
			},
		},
		{ //43
			{
				Utf8:   "                同一台機器可以快照備份，但如果是RDS+ECS的狀況該怎麼備份比較合適",
				Big5:   []byte("                \xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa5i\xa5H\xa7\xd6\xb7\xd3\xb3\xc6\xa5\xf7\xa1A\xa6\xfd\xa6p\xaaG\xacORDS+ECS\xaa\xba\xaa\xac\xaap\xb8\xd3\xab\xe7\xbb\xf2\xb3\xc6\xa5\xf7\xa4\xf1\xb8\xfb\xa6X\xbeA"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                \xa6P\xa4@\xa5x\xbe\xf7\xbe\xb9\xa5i\xa5H\xa7\xd6\xb7\xd3\xb3\xc6\xa5\xf7\xa1A\xa6\xfd\xa6p\xaaG\xacORDS+ECS\xaa\xba\xaa\xac\xaap\xb8\xd3\xab\xe7\xbb\xf2\xb3\xc6\xa5\xf7\xa4\xf1\xb8\xfb\xa6X\xbeA\r"),
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
				Utf8:   "                4. 教我系統監控的察看重點",
				Big5:   []byte("                4. \xb1\xd0\xa7\xda\xa8t\xb2\xce\xba\xca\xb1\xb1\xaa\xba\xb9\xee\xac\xdd\xad\xab\xc2I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                4. \xb1\xd0\xa7\xda\xa8t\xb2\xce\xba\xca\xb1\xb1\xaa\xba\xb9\xee\xac\xdd\xad\xab\xc2I\r"),
			},
		},
		{ //46
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //47
			{
				Utf8:   "                5. 安全性及防火牆設置教學",
				Big5:   []byte("                5. \xa6w\xa5\xfe\xa9\xca\xa4\xce\xa8\xbe\xa4\xf5\xc0\xf0\xb3]\xb8m\xb1\xd0\xbe\xc7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                5. \xa6w\xa5\xfe\xa9\xca\xa4\xce\xa8\xbe\xa4\xf5\xc0\xf0\xb3]\xb8m\xb1\xd0\xbe\xc7\r"),
			},
		},
		{ //48
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //49
			{
				Utf8:   "                6. (非必要)NGINX網頁伺服器教學",
				Big5:   []byte("                6. (\xabD\xa5\xb2\xadn)NGINX\xba\xf4\xad\xb6\xa6\xf8\xaaA\xbe\xb9\xb1\xd0\xbe\xc7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                6. (\xabD\xa5\xb2\xadn)NGINX\xba\xf4\xad\xb6\xa6\xf8\xaaA\xbe\xb9\xb1\xd0\xbe\xc7\r"),
			},
		},
	}

	testFirstComments9 = []*schema.Comment{
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("xxFrency"),
			Content: nil,
			MD5:     "2dn14lVR7wDw3BC6XOEizw",
			TheDate: "12/19/2020 14:47:01",
			IP:      "150.117.235.183",
			Host:    "臺灣",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: xxFrency (150.117.235.183 \xbbO\xc6W), 12/19/2020 14:47:01\r"),

			CreateTime:         1608360421000000000,
			InferredCreateTime: 1608360421000000000,
			SortTime:           1608360421000000000,
			CommentID:          "FlILSAyAsgA:2dn14lVR7wDw3BC6XOEizw",
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("xxFrency"),
			Content: nil,
			MD5:     "L1kywzRrQ3nO1w4NU_C3ig",
			TheDate: "12/19/2020 14:47:46",
			IP:      "150.117.235.183",
			Host:    "臺灣",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: xxFrency (150.117.235.183 \xbbO\xc6W), 12/19/2020 14:47:46\r"),

			CreateTime:         1608360466000000000,
			InferredCreateTime: 1608360466000000000,
			SortTime:           1608360466000000000,
			CommentID:          "FlILUoa2NAA:L1kywzRrQ3nO1w4NU_C3ig",
		},
		{
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("asdfghjklasd"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這種的不是要找 Tech Lead 就是要找顧問",
						Big5:   []byte("\xb3o\xba\xd8\xaa\xba\xa4\xa3\xacO\xadn\xa7\xe4 Tech Lead \xb4N\xacO\xadn\xa7\xe4\xc5U\xb0\xdd            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xba\xd8\xaa\xba\xa4\xa3\xacO\xadn\xa7\xe4 Tech Lead \xb4N\xacO\xadn\xa7\xe4\xc5U\xb0\xdd            "),
					},
				},
			},
			MD5:     "LqbCx26fVtY50dudiEeOjw",
			TheDate: "12/19 17:21",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33masdfghjklasd\x1b[m\x1b[33m: \xb3o\xba\xd8\xaa\xba\xa4\xa3\xacO\xadn\xa7\xe4 Tech Lead \xb4N\xacO\xadn\xa7\xe4\xc5U\xb0\xdd            \x1b[m 12/19 17:21\r"),
		},
		{
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("dinos"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不是把機器加大,是把機器加多",
						Big5:   []byte("\xa4\xa3\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j,\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa6h                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xa3\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j,\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa6h                             "),
					},
				},
			},
			MD5:     "WWPwvPGXeBZ5R8IbyEqW-g",
			TheDate: "12/19 18:38",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mdinos\x1b[m\x1b[33m: \xa4\xa3\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j,\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa6h                             \x1b[m 12/19 18:38\r"),
		},
		{
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("bestwishes"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這些跟程式架構也有關係 不是設定調一調就好了",
						Big5:   []byte("\xb3o\xa8\xc7\xb8\xf2\xb5{\xa6\xa1\xac[\xbac\xa4]\xa6\xb3\xc3\xf6\xabY \xa4\xa3\xacO\xb3]\xa9w\xbd\xd5\xa4@\xbd\xd5\xb4N\xa6n\xa4F        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xa8\xc7\xb8\xf2\xb5{\xa6\xa1\xac[\xbac\xa4]\xa6\xb3\xc3\xf6\xabY \xa4\xa3\xacO\xb3]\xa9w\xbd\xd5\xa4@\xbd\xd5\xb4N\xa6n\xa4F        "),
					},
				},
			},
			MD5:     "I1Rz90BfBPhbAcLrIK1Bzg",
			TheDate: "12/19 18:57",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mbestwishes\x1b[m\x1b[33m: \xb3o\xa8\xc7\xb8\xf2\xb5{\xa6\xa1\xac[\xbac\xa4]\xa6\xb3\xc3\xf6\xabY \xa4\xa3\xacO\xb3]\xa9w\xbd\xd5\xa4@\xbd\xd5\xb4N\xa6n\xa4F        \x1b[m 12/19 18:57\r"),
		},
	}

	testFullFirstComments9 = []*schema.Comment{
		{
			BBoardID:  "test",
			ArticleID: "test9",
			TheType:   types.COMMENT_TYPE_EDIT,
			Owner:     bbs.UUserID("xxFrency"),
			Content:   nil,
			MD5:       "2dn14lVR7wDw3BC6XOEizw",
			TheDate:   "12/19/2020 14:47:01",
			IP:        "150.117.235.183",
			Host:      "臺灣",
			DBCS:      []byte("\xa1\xb0 \xbds\xbf\xe8: xxFrency (150.117.235.183 \xbbO\xc6W), 12/19/2020 14:47:01\r"),

			CreateTime:         1608360421000000000,
			InferredCreateTime: 1608360421000000000,
			SortTime:           1608360421000000000,
			CommentID:          "FlILSAyAsgA:2dn14lVR7wDw3BC6XOEizw",
		},
		{
			BBoardID:  "test",
			ArticleID: "test9",
			TheType:   types.COMMENT_TYPE_EDIT,
			Owner:     bbs.UUserID("xxFrency"),
			Content:   nil,
			MD5:       "L1kywzRrQ3nO1w4NU_C3ig",
			TheDate:   "12/19/2020 14:47:46",
			IP:        "150.117.235.183",
			Host:      "臺灣",
			DBCS:      []byte("\xa1\xb0 \xbds\xbf\xe8: xxFrency (150.117.235.183 \xbbO\xc6W), 12/19/2020 14:47:46\r"),

			CreateTime:         1608360466000000000,
			InferredCreateTime: 1608360466000000000,
			SortTime:           1608360466000000000,
			CommentID:          "FlILUoa2NAA:L1kywzRrQ3nO1w4NU_C3ig",
		},
		{
			BBoardID:  "test",
			ArticleID: "test9",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("asdfghjklasd"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這種的不是要找 Tech Lead 就是要找顧問",
						Big5:   []byte("\xb3o\xba\xd8\xaa\xba\xa4\xa3\xacO\xadn\xa7\xe4 Tech Lead \xb4N\xacO\xadn\xa7\xe4\xc5U\xb0\xdd            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xba\xd8\xaa\xba\xa4\xa3\xacO\xadn\xa7\xe4 Tech Lead \xb4N\xacO\xadn\xa7\xe4\xc5U\xb0\xdd            "),
					},
				},
			},
			MD5:        "LqbCx26fVtY50dudiEeOjw",
			TheDate:    "12/19 17:21",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33masdfghjklasd\x1b[m\x1b[33m: \xb3o\xba\xd8\xaa\xba\xa4\xa3\xacO\xadn\xa7\xe4 Tech Lead \xb4N\xacO\xadn\xa7\xe4\xc5U\xb0\xdd            \x1b[m 12/19 17:21\r"),
			CommentID:  "FlITryvQ2AA:LqbCx26fVtY50dudiEeOjw",
			CreateTime: 1608369660000000000,
			SortTime:   1608369660000000000,
		},
		{
			BBoardID:  "test",
			ArticleID: "test9",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("dinos"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不是把機器加大,是把機器加多",
						Big5:   []byte("\xa4\xa3\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j,\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa6h                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xa3\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j,\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa6h                             "),
					},
				},
			},
			MD5:        "WWPwvPGXeBZ5R8IbyEqW-g",
			TheDate:    "12/19 18:38",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mdinos\x1b[m\x1b[33m: \xa4\xa3\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa4j,\xacO\xa7\xe2\xbe\xf7\xbe\xb9\xa5[\xa6h                             \x1b[m 12/19 18:38\r"),
			CommentID:  "FlIX4tlGUAA:WWPwvPGXeBZ5R8IbyEqW-g",
			CreateTime: 1608374280000000000,
			SortTime:   1608374280000000000,
		},
		{
			BBoardID:  "test",
			ArticleID: "test9",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("bestwishes"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這些跟程式架構也有關係 不是設定調一調就好了",
						Big5:   []byte("\xb3o\xa8\xc7\xb8\xf2\xb5{\xa6\xa1\xac[\xbac\xa4]\xa6\xb3\xc3\xf6\xabY \xa4\xa3\xacO\xb3]\xa9w\xbd\xd5\xa4@\xbd\xd5\xb4N\xa6n\xa4F        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xa8\xc7\xb8\xf2\xb5{\xa6\xa1\xac[\xbac\xa4]\xa6\xb3\xc3\xf6\xabY \xa4\xa3\xacO\xb3]\xa9w\xbd\xd5\xa4@\xbd\xd5\xb4N\xa6n\xa4F        "),
					},
				},
			},
			MD5:        "I1Rz90BfBPhbAcLrIK1Bzg",
			TheDate:    "12/19 18:57",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mbestwishes\x1b[m\x1b[33m: \xb3o\xa8\xc7\xb8\xf2\xb5{\xa6\xa1\xac[\xbac\xa4]\xa6\xb3\xc3\xf6\xabY \xa4\xa3\xacO\xb3]\xa9w\xbd\xd5\xa4@\xbd\xd5\xb4N\xa6n\xa4F        \x1b[m 12/19 18:57\r"),
			CommentID:  "FlIY8o_lJgA:I1Rz90BfBPhbAcLrIK1Bzg",
			CreateTime: 1608375447000000000,
			SortTime:   1608375447000000000,
		},
	}
}
