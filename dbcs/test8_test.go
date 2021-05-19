package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename8            = "temp3"
	testContentAll8          []byte
	testContent8             []byte
	testSignature8           []byte
	testComment8             []byte
	testFirstCommentsDBCS8   []byte
	testTheRestCommentsDBCS8 []byte

	testContent8Big5 [][]*types.Rune
	testContent8Utf8 [][]*types.Rune

	testFirstComments8     []*schema.Comment
	testFullFirstComments8 []*schema.Comment
)

func initTest8() {
	testContentAll8, testContent8, testSignature8, testComment8, testFirstCommentsDBCS8, testTheRestCommentsDBCS8 = loadTest(testFilename8)

	testContent8Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa1\xf0\xb7\xe0\xa4l\xa4k&\xa4\xd1\xc3\xc8\xa8k\xa1\xf1\xaa\xba\xb7R\xb1\xa1\xad\xe8\xb8\xa8\xb9\xf5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xf0\xb7\xe0\xa4l\xa4k&\xa4\xd1\xc3\xc8\xa8k\xa1\xf1\xaa\xba\xb7R\xb1\xa1\xad\xe8\xb8\xa8\xb9\xf5\r"),
			},
		},
		{ //1
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xa9p\xc1\xd9\xa6b\xb5S\xbf\xdd\xa6\xdb\xa4v\xaa\xba\xb7R\xb1\xa1\xc2k\xb1J\xa6b\xad\xfe\xb8\xcc\xb6\xdc\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa9p\xc1\xd9\xa6b\xb5S\xbf\xdd\xa6\xdb\xa4v\xaa\xba\xb7R\xb1\xa1\xc2k\xb1J\xa6b\xad\xfe\xb8\xcc\xb6\xdc\xa1H\r"),
			},
		},
		{ //3
			{

				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                              "),
			},
			{

				Big5:   []byte("\xf9\xfe\xf9\xfe    \xf9\xfe\xf9\xfe"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xf9\xfe\xf9\xfe    \xf9\xfe\xf9\xfe"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //4
			{

				Big5:   []byte("\xb0\xa8\xa4W\xb5n\xa4J\xc1p\xa6X\xb3\xf8iPOST\xac\xa1\xb0\xca\xba\xf4\xaf\xb8                   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb0\xa8\xa4W\xb5n\xa4J\xc1p\xa6X\xb3\xf8iPOST\xac\xa1\xb0\xca\xba\xf4\xaf\xb8                   "),
			},
			{

				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p\xa2p\xa2i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xa2p\xa2p\xa2p\xa2p\xa2p\xa2i"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //5
			{

				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                              "),
			},
			{

				Big5:   []byte(" \xa2p\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m \xa2p\xa2p\xa2p\xa2p\xa2p"),
			},
			{

				Big5:   []byte("       "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m       "),
			},
			{

				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35m\xa1\xb4"),
			},
			{

				Big5:   []byte("    "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m    "),
			},
			{

				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;34m\xa1\xb4"),
			},
			{

				Big5:   []byte("\xa2A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2A\r"),
			},
		},
		{ //6
			{

				Big5:   []byte("\xa7\xd6\xb3t\xa4\xc0\xaaR\xb1z\xaa\xba2012\xc5\xca\xb7R\xb9B\xb6\xd5                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xd6\xb3t\xa4\xc0\xaaR\xb1z\xaa\xba2012\xc5\xca\xb7R\xb9B\xb6\xd5                        "),
			},
			{

				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xa2p\xa2p\xa2p\xa2p"),
			},
			{

				Big5:   []byte("        \xa1\xbd\xa1\xd7\xa1X\xa1\xbd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m        \xa1\xbd\xa1\xd7\xa1X\xa1\xbd\r"),
			},
		},
		{ //7
			{

				Big5:   []byte("                                                "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                                "),
			},
			{

				Big5:   []byte(" \xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m \xa2p\xa2p\xa2p"),
			},
			{

				Big5:   []byte("         \xa1\xbf    \xa1\xbf"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m         \xa1\xbf    \xa1\xbf\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("\xa6P\xae\xc9\xa4\xc0\xa8\xc9\xb7R\xb1\xa1\xa5@\xac\xc9\xaa\xba\xbb\xc4\xb2\xa2\xadW\xbb\xb6                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6P\xae\xc9\xa4\xc0\xa8\xc9\xb7R\xb1\xa1\xa5@\xac\xc9\xaa\xba\xbb\xc4\xb2\xa2\xadW\xbb\xb6                        "),
			},
			{

				Big5:   []byte(" \xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m \xa2p"),
			},
			{

				Big5:   []byte("          \xa2A\xa2B  \xa2A\xa2B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m          \xa2A\xa2B  \xa2A\xa2B\r"),
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

				Big5:   []byte("\xc1p\xa6X\xabK\xa7Q\xb6K"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xc1p\xa6X\xabK\xa7Q\xb6K"),
			},
			{

				Big5:   []byte("\xa7K\xb6O"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa7K\xb6O"),
			},
			{

				Big5:   []byte("\xbd\xd0\xa9p\xa6Y\xaak\xa6\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xbd\xd0\xa9p\xa6Y\xaak\xa6\xa1"),
			},
			{

				Big5:   []byte("\xba\xeb\xbdo\xb2\xa2\xc2I"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xba\xeb\xbdo\xb2\xa2\xc2I"),
			},
			{

				Big5:   []byte("\xa4@\xa5\xf7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa4@\xa5\xf7\r"),
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

				Big5:   []byte("\xb4\xa3\xabe\x83\xf2\xa5[\xac\xa1\xb0\xca"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb4\xa3\xabe\x83\xf2\xa5[\xac\xa1\xb0\xca\r"),
			},
		},
		{ //13
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //14
			{

				Big5:   []byte("\xc1\xd9\xa6\xb3\xbe\xf7\xb7|\xa9\xe2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xc1\xd9\xa6\xb3\xbe\xf7\xb7|\xa9\xe2"),
			},
			{

				Big5:   []byte("\xa7K\xb6O\xaei\xc4\xfd\xb2\xbc\xa8\xe9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m\xa7K\xb6O\xaei\xc4\xfd\xb2\xbc\xa8\xe9"),
			},
			{

				Big5:   []byte("\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1I\r"),
			},
		},
		{ //15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //16
			{

				Big5:   []byte("\xb3\xcc\xb2\xa2\xbbe\xb7R\xb1\xa1\xa1A\xb3\xa3\xa6b2012\xa7\xda\xaa\xba\xb7R\xb1\xa1\xa8\xd3\xa4F\xa8S\xa1I\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3\xcc\xb2\xa2\xbbe\xb7R\xb1\xa1\xa1A\xb3\xa3\xa6b2012\xa7\xda\xaa\xba\xb7R\xb1\xa1\xa8\xd3\xa4F\xa8S\xa1I\xa1H\r"),
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

				Big5:   []byte("\xb3s\xb5\xb2\xa1Ghttp://event.udn.com/2012love/"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3s\xb5\xb2\xa1Ghttp://event.udn.com/2012love/\r"),
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

				Big5:   []byte("\xa5D\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xa5D\xbf\xec\xb3\xe6\xa6\xec"),
			},
			{

				Big5:   []byte(":\xc1p\xa6X\xabK\xa7Q\xb6K   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m:\xc1p\xa6X\xabK\xa7Q\xb6K   "),
			},
			{

				Big5:   []byte("\xa8\xf3\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa8\xf3\xbf\xec\xb3\xe6\xa6\xec"),
			},
			{

				Big5:   []byte(":\xa5i\xa6\xb7\xa5\xa9\xa7J\xa4O  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m:\xa5i\xa6\xb7\xa5\xa9\xa7J\xa4O  "),
			},
			{

				Big5:   []byte("\xa6X\xa7@\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa6X\xa7@\xb3\xe6\xa6\xec"),
			},
			{

				Big5:   []byte(":\xabB\xb4\xad\xa9R\xb2z\xba\xf4\xaf\xb8-\xb6}\xa4\xdf\xb4N\xacO\xb6}\xb9B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m:\xabB\xb4\xad\xa9R\xb2z\xba\xf4\xaf\xb8-\xb6}\xa4\xdf\xb4N\xacO\xb6}\xb9B\r"),
			},
		},
	}

	testContent8Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "♀獅子女&天蠍男♂的愛情剛落幕",
				Big5:   []byte("\xa1\xf0\xb7\xe0\xa4l\xa4k&\xa4\xd1\xc3\xc8\xa8k\xa1\xf1\xaa\xba\xb7R\xb1\xa1\xad\xe8\xb8\xa8\xb9\xf5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xf0\xb7\xe0\xa4l\xa4k&\xa4\xd1\xc3\xc8\xa8k\xa1\xf1\xaa\xba\xb7R\xb1\xa1\xad\xe8\xb8\xa8\xb9\xf5\r"),
			},
		},
		{ //1
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //2
			{
				Utf8:   "妳還在猶豫自己的愛情歸宿在哪裡嗎？",
				Big5:   []byte("\xa9p\xc1\xd9\xa6b\xb5S\xbf\xdd\xa6\xdb\xa4v\xaa\xba\xb7R\xb1\xa1\xc2k\xb1J\xa6b\xad\xfe\xb8\xcc\xb6\xdc\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa9p\xc1\xd9\xa6b\xb5S\xbf\xdd\xa6\xdb\xa4v\xaa\xba\xb7R\xb1\xa1\xc2k\xb1J\xa6b\xad\xfe\xb8\xcc\xb6\xdc\xa1H\r"),
			},
		},
		{ //3
			{
				Utf8:   "                                              ",
				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                              "),
			},
			{
				Utf8:   "▓▓    ▓▓",
				Big5:   []byte("\xf9\xfe\xf9\xfe    \xf9\xfe\xf9\xfe"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xf9\xfe\xf9\xfe    \xf9\xfe\xf9\xfe"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //4
			{
				Utf8:   "馬上登入聯合報iPOST活動網站                   ",
				Big5:   []byte("\xb0\xa8\xa4W\xb5n\xa4J\xc1p\xa6X\xb3\xf8iPOST\xac\xa1\xb0\xca\xba\xf4\xaf\xb8                   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb0\xa8\xa4W\xb5n\xa4J\xc1p\xa6X\xb3\xf8iPOST\xac\xa1\xb0\xca\xba\xf4\xaf\xb8                   "),
			},
			{
				Utf8:   "▉▉▉▉▉█",
				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p\xa2p\xa2i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xa2p\xa2p\xa2p\xa2p\xa2p\xa2i"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ //5
			{
				Utf8:   "                                              ",
				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                              "),
			},
			{
				Utf8:   " ▉▉▉▉▉",
				Big5:   []byte(" \xa2p\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m \xa2p\xa2p\xa2p\xa2p\xa2p"),
			},
			{
				Utf8:   "       ",
				Big5:   []byte("       "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m       "),
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;35m\xa1\xb4"),
			},
			{
				Utf8:   "    ",
				Big5:   []byte("    "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m    "),
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;34m\xa1\xb4"),
			},
			{
				Utf8:   "∕",
				Big5:   []byte("\xa2A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa2A\r"),
			},
		},
		{ //6
			{
				Utf8:   "快速分析您的2012戀愛運勢                        ",
				Big5:   []byte("\xa7\xd6\xb3t\xa4\xc0\xaaR\xb1z\xaa\xba2012\xc5\xca\xb7R\xb9B\xb6\xd5                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xd6\xb3t\xa4\xc0\xaaR\xb1z\xaa\xba2012\xc5\xca\xb7R\xb9B\xb6\xd5                        "),
			},
			{
				Utf8:   "▉▉▉▉",
				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xa2p\xa2p\xa2p\xa2p"),
			},
			{
				Utf8:   "        ■＝—■",
				Big5:   []byte("        \xa1\xbd\xa1\xd7\xa1X\xa1\xbd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m        \xa1\xbd\xa1\xd7\xa1X\xa1\xbd\r"),
			},
		},
		{ //7
			{
				Utf8:   "                                                ",
				Big5:   []byte("                                                "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                                "),
			},
			{
				Utf8:   " ▉▉▉",
				Big5:   []byte(" \xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m \xa2p\xa2p\xa2p"),
			},
			{
				Utf8:   "         ▼    ▼",
				Big5:   []byte("         \xa1\xbf    \xa1\xbf"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m         \xa1\xbf    \xa1\xbf\r"),
			},
		},
		{ //8
			{
				Utf8:   "同時分享愛情世界的酸甜苦辣                        ",
				Big5:   []byte("\xa6P\xae\xc9\xa4\xc0\xa8\xc9\xb7R\xb1\xa1\xa5@\xac\xc9\xaa\xba\xbb\xc4\xb2\xa2\xadW\xbb\xb6                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6P\xae\xc9\xa4\xc0\xa8\xc9\xb7R\xb1\xa1\xa5@\xac\xc9\xaa\xba\xbb\xc4\xb2\xa2\xadW\xbb\xb6                        "),
			},
			{
				Utf8:   " ▉",
				Big5:   []byte(" \xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m \xa2p"),
			},
			{
				Utf8:   "          ∕﹨  ∕﹨",
				Big5:   []byte("          \xa2A\xa2B  \xa2A\xa2B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m          \xa2A\xa2B  \xa2A\xa2B\r"),
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
				Utf8:   "聯合便利貼",
				Big5:   []byte("\xc1p\xa6X\xabK\xa7Q\xb6K"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xc1p\xa6X\xabK\xa7Q\xb6K"),
			},
			{
				Utf8:   "免費",
				Big5:   []byte("\xa7K\xb6O"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa7K\xb6O"),
			},
			{
				Utf8:   "請妳吃法式",
				Big5:   []byte("\xbd\xd0\xa9p\xa6Y\xaak\xa6\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xbd\xd0\xa9p\xa6Y\xaak\xa6\xa1"),
			},
			{
				Utf8:   "精緻甜點",
				Big5:   []byte("\xba\xeb\xbdo\xb2\xa2\xc2I"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xba\xeb\xbdo\xb2\xa2\xc2I"),
			},
			{
				Utf8:   "一份",
				Big5:   []byte("\xa4@\xa5\xf7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa4@\xa5\xf7\r"),
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
				Utf8:   "提前参加活動",
				Big5:   []byte("\xb4\xa3\xabe\x83\xf2\xa5[\xac\xa1\xb0\xca"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb4\xa3\xabe\x83\xf2\xa5[\xac\xa1\xb0\xca\r"),
			},
		},
		{ //13
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //14
			{
				Utf8:   "還有機會抽",
				Big5:   []byte("\xc1\xd9\xa6\xb3\xbe\xf7\xb7|\xa9\xe2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xc1\xd9\xa6\xb3\xbe\xf7\xb7|\xa9\xe2"),
			},
			{
				Utf8:   "免費展覽票券",
				Big5:   []byte("\xa7K\xb6O\xaei\xc4\xfd\xb2\xbc\xa8\xe9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;33m\xa7K\xb6O\xaei\xc4\xfd\xb2\xbc\xa8\xe9"),
			},
			{
				Utf8:   "！",
				Big5:   []byte("\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1I\r"),
			},
		},
		{ //15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //16
			{
				Utf8:   "最甜蜜愛情，都在2012我的愛情來了沒！？",
				Big5:   []byte("\xb3\xcc\xb2\xa2\xbbe\xb7R\xb1\xa1\xa1A\xb3\xa3\xa6b2012\xa7\xda\xaa\xba\xb7R\xb1\xa1\xa8\xd3\xa4F\xa8S\xa1I\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3\xcc\xb2\xa2\xbbe\xb7R\xb1\xa1\xa1A\xb3\xa3\xa6b2012\xa7\xda\xaa\xba\xb7R\xb1\xa1\xa8\xd3\xa4F\xa8S\xa1I\xa1H\r"),
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
				Utf8:   "連結：http://event.udn.com/2012love/",
				Big5:   []byte("\xb3s\xb5\xb2\xa1Ghttp://event.udn.com/2012love/"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3s\xb5\xb2\xa1Ghttp://event.udn.com/2012love/\r"),
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
				Utf8:   "主辦單位",
				Big5:   []byte("\xa5D\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m\xa5D\xbf\xec\xb3\xe6\xa6\xec"),
			},
			{
				Utf8:   ":聯合便利貼   ",
				Big5:   []byte(":\xc1p\xa6X\xabK\xa7Q\xb6K   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m:\xc1p\xa6X\xabK\xa7Q\xb6K   "),
			},
			{
				Utf8:   "協辦單位",
				Big5:   []byte("\xa8\xf3\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa8\xf3\xbf\xec\xb3\xe6\xa6\xec"),
			},
			{
				Utf8:   ":可朵巧克力  ",
				Big5:   []byte(":\xa5i\xa6\xb7\xa5\xa9\xa7J\xa4O  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m:\xa5i\xa6\xb7\xa5\xa9\xa7J\xa4O  "),
			},
			{
				Utf8:   "合作單位",
				Big5:   []byte("\xa6X\xa7@\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa6X\xa7@\xb3\xe6\xa6\xec"),
			},
			{
				Utf8:   ":雨揚命理網站-開心就是開運",
				Big5:   []byte(":\xabB\xb4\xad\xa9R\xb2z\xba\xf4\xaf\xb8-\xb6}\xa4\xdf\xb4N\xacO\xb6}\xb9B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m:\xabB\xb4\xad\xa9R\xb2z\xba\xf4\xaf\xb8-\xb6}\xa4\xdf\xb4N\xacO\xb6}\xb9B\r"),
			},
		},
	}

	testFirstComments8 = []*schema.Comment{
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "KxkHBwyiFnfhZJ6u-ayFFA",
			TheDate: "12/22 06:53",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 06:53)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
		{
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("shiou7788"),
			Content: nil,
			MD5:     "HbrbjPybbW4AtITcFsZM2A",
			TheDate: "12/22 07:00",
			IP:      "140.112.92.137",
			DBCS:    []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
		},
	}

	testFullFirstComments8 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			CommentID:  types.CommentID("E7IFC5xJ-AA:KxkHBwyiFnfhZJ6u-ayFFA"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "KxkHBwyiFnfhZJ6u-ayFFA",
			TheDate:    "12/22 06:53",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 06:53)\r"),
			CreateTime: 1419202380000000000,
			SortTime:   1419202380000000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			CommentID:  "E7IFbWY9YAA:HbrbjPybbW4AtITcFsZM2A",
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CreateTime: 1419202800000000000,
			SortTime:   1419202800000000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			CommentID:  "E7IFbWZMokA:HbrbjPybbW4AtITcFsZM2A",
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CreateTime: 1419202800000000000,
			SortTime:   1419202800001000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CommentID:  "E7IFbWZb5IA:HbrbjPybbW4AtITcFsZM2A",
			CreateTime: 1419202800000000000,
			SortTime:   1419202800002000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CommentID:  "E7IFbWZrJsA:HbrbjPybbW4AtITcFsZM2A",
			CreateTime: 1419202800000000000,
			SortTime:   1419202800003000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CommentID:  "E7IFbWZ6aQA:HbrbjPybbW4AtITcFsZM2A",
			CreateTime: 1419202800000000000,
			SortTime:   1419202800004000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CommentID:  "E7IFbWaJq0A:HbrbjPybbW4AtITcFsZM2A",
			CreateTime: 1419202800000000000,
			SortTime:   1419202800005000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CommentID:  "E7IFbWaY7YA:HbrbjPybbW4AtITcFsZM2A",
			CreateTime: 1419202800000000000,
			SortTime:   1419202800006000000,
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test8"),
			TheType:    types.COMMENT_TYPE_EDIT,
			Owner:      bbs.UUserID("shiou7788"),
			Content:    nil,
			MD5:        "HbrbjPybbW4AtITcFsZM2A",
			TheDate:    "12/22 07:00",
			IP:         "140.112.92.137",
			DBCS:       []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)\r"),
			CommentID:  "E7IFbWaoL8A:HbrbjPybbW4AtITcFsZM2A",
			CreateTime: 1419202800000000000,
			SortTime:   1419202800007000000,
		},
	}
}
