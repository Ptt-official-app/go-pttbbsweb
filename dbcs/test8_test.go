package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
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

	testFirstComments8 []*schema.Comment
)

func initTest8() {
	testContentAll8, testContent8, testSignature8, testComment8, testFirstCommentsDBCS8, testTheRestCommentsDBCS8 = loadTest(testFilename8)

	testContent8Big5 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "",
				Big5:   []byte("\xa1\xf0\xb7\xe0\xa4l\xa4k&\xa4\xd1\xc3\xc8\xa8k\xa1\xf1\xaa\xba\xb7R\xb1\xa1\xad\xe8\xb8\xa8\xb9\xf5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //1
		{ //2
			{
				Big5:   []byte("\xa9p\xc1\xd9\xa6b\xb5S\xbf\xdd\xa6\xdb\xa4v\xaa\xba\xb7R\xb1\xa1\xc2k\xb1J\xa6b\xad\xfe\xb8\xcc\xb6\xdc\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //3
			{
				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xf9\xfe\xf9\xfe    \xf9\xfe\xf9\xfe"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
		},
		{ //4
			{
				Big5:   []byte("\xb0\xa8\xa4W\xb5n\xa4J\xc1p\xa6X\xb3\xf8iPOST\xac\xa1\xb0\xca\xba\xf4\xaf\xb8                   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p\xa2p\xa2i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
		},
		{ //5
			{
				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte(" \xa2p\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("       "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("    "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("\xa2A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //6
			{
				Big5:   []byte("\xa7\xd6\xb3t\xa4\xc0\xaaR\xb1z\xaa\xba2012\xc5\xca\xb7R\xb9B\xb6\xd5                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("        \xa1\xbd\xa1\xd7\xa1X\xa1\xbd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //7
			{
				Big5:   []byte("                                                "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte(" \xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("         \xa1\xbf    \xa1\xbf"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //8
			{
				Big5:   []byte("\xa6P\xae\xc9\xa4\xc0\xa8\xc9\xb7R\xb1\xa1\xa5@\xac\xc9\xaa\xba\xbb\xc4\xb2\xa2\xadW\xbb\xb6                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte(" \xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("          \xa2A\xa2B  \xa2A\xa2B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //9
		{ //10
			{
				Big5:   []byte("\xc1p\xa6X\xabK\xa7Q\xb6K"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa7K\xb6O"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("\xbd\xd0\xa9p\xa6Y\xaak\xa6\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xba\xeb\xbdo\xb2\xa2\xc2I"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("\xa4@\xa5\xf7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //11
		{ //12
			{
				Big5:   []byte("\xb4\xa3\xabe\x83\xf2\xa5[\xac\xa1\xb0\xca"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //13
		{ //14
			{
				Big5:   []byte("\xc1\xd9\xa6\xb3\xbe\xf7\xb7|\xa9\xe2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa7K\xb6O\xaei\xc4\xfd\xb2\xbc\xa8\xe9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte("\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //15
		{ //16
			{
				Big5:   []byte("\xb3\xcc\xb2\xa2\xbbe\xb7R\xb1\xa1\xa1A\xb3\xa3\xa6b2012\xa7\xda\xaa\xba\xb7R\xb1\xa1\xa8\xd3\xa4F\xa8S\xa1I\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //17
		{ //18
			{
				Big5:   []byte("\xb3s\xb5\xb2\xa1Ghttp://event.udn.com/2012love/"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //19
		{ //20
			{
				Big5:   []byte("\xa5D\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(":\xc1p\xa6X\xabK\xa7Q\xb6K   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa8\xf3\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(":\xa5i\xa6\xb7\xa5\xa9\xa7J\xa4O  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Big5:   []byte("\xa6X\xa7@\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Big5:   []byte(":\xabB\xb4\xad\xa9R\xb2z\xba\xf4\xaf\xb8-\xb6}\xa4\xdf\xb4N\xacO\xb6}\xb9B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //21
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 06:53)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //22
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //23
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //24
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //25
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //26
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //27
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //28
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //29
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //30
	}

	testContent8Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "♀獅子女&天蠍男♂的愛情剛落幕",
				Big5:   []byte("\xa1\xf0\xb7\xe0\xa4l\xa4k&\xa4\xd1\xc3\xc8\xa8k\xa1\xf1\xaa\xba\xb7R\xb1\xa1\xad\xe8\xb8\xa8\xb9\xf5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //1
		{ //2
			{
				Utf8:   "妳還在猶豫自己的愛情歸宿在哪裡嗎？",
				Big5:   []byte("\xa9p\xc1\xd9\xa6b\xb5S\xbf\xdd\xa6\xdb\xa4v\xaa\xba\xb7R\xb1\xa1\xc2k\xb1J\xa6b\xad\xfe\xb8\xcc\xb6\xdc\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //3
			{
				Utf8:   "                                              ",
				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "▓▓    ▓▓",
				Big5:   []byte("\xf9\xfe\xf9\xfe    \xf9\xfe\xf9\xfe"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
		},
		{ //4
			{
				Utf8:   "馬上登入聯合報iPOST活動網站                   ",
				Big5:   []byte("\xb0\xa8\xa4W\xb5n\xa4J\xc1p\xa6X\xb3\xf8iPOST\xac\xa1\xb0\xca\xba\xf4\xaf\xb8                   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "▉▉▉▉▉█",
				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p\xa2p\xa2i"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
		},
		{ //5
			{
				Utf8:   "                                              ",
				Big5:   []byte("                                              "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   " ▉▉▉▉▉",
				Big5:   []byte(" \xa2p\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "       ",
				Big5:   []byte("       "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_MAGENTA, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "    ",
				Big5:   []byte("    "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "●",
				Big5:   []byte("\xa1\xb4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLUE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "∕",
				Big5:   []byte("\xa2A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //6
			{
				Utf8:   "快速分析您的2012戀愛運勢                        ",
				Big5:   []byte("\xa7\xd6\xb3t\xa4\xc0\xaaR\xb1z\xaa\xba2012\xc5\xca\xb7R\xb9B\xb6\xd5                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "▉▉▉▉",
				Big5:   []byte("\xa2p\xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "        ■＝—■",
				Big5:   []byte("        \xa1\xbd\xa1\xd7\xa1X\xa1\xbd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //7
			{
				Utf8:   "                                                ",
				Big5:   []byte("                                                "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   " ▉▉▉",
				Big5:   []byte(" \xa2p\xa2p\xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "         ▼    ▼",
				Big5:   []byte("         \xa1\xbf    \xa1\xbf"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //8
			{
				Utf8:   "同時分享愛情世界的酸甜苦辣                        ",
				Big5:   []byte("\xa6P\xae\xc9\xa4\xc0\xa8\xc9\xb7R\xb1\xa1\xa5@\xac\xc9\xaa\xba\xbb\xc4\xb2\xa2\xadW\xbb\xb6                        "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   " ▉",
				Big5:   []byte(" \xa2p"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "          ∕﹨  ∕﹨",
				Big5:   []byte("          \xa2A\xa2B  \xa2A\xa2B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //9
		{ //10
			{
				Utf8:   "聯合便利貼",
				Big5:   []byte("\xc1p\xa6X\xabK\xa7Q\xb6K"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "免費",
				Big5:   []byte("\xa7K\xb6O"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "請妳吃法式",
				Big5:   []byte("\xbd\xd0\xa9p\xa6Y\xaak\xa6\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "精緻甜點",
				Big5:   []byte("\xba\xeb\xbdo\xb2\xa2\xc2I"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "一份",
				Big5:   []byte("\xa4@\xa5\xf7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //11
		{ //12
			{
				Utf8:   "提前参加活動",
				Big5:   []byte("\xb4\xa3\xabe\x83\xf2\xa5[\xac\xa1\xb0\xca"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //13
		{ //14
			{
				Utf8:   "還有機會抽",
				Big5:   []byte("\xc1\xd9\xa6\xb3\xbe\xf7\xb7|\xa9\xe2"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "免費展覽票券",
				Big5:   []byte("\xa7K\xb6O\xaei\xc4\xfd\xb2\xbc\xa8\xe9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   "！",
				Big5:   []byte("\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //15
		{ //16
			{
				Utf8:   "最甜蜜愛情，都在2012我的愛情來了沒！？",
				Big5:   []byte("\xb3\xcc\xb2\xa2\xbbe\xb7R\xb1\xa1\xa1A\xb3\xa3\xa6b2012\xa7\xda\xaa\xba\xb7R\xb1\xa1\xa8\xd3\xa4F\xa8S\xa1I\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //17
		{ //18
			{
				Utf8:   "連結：http://event.udn.com/2012love/",
				Big5:   []byte("\xb3s\xb5\xb2\xa1Ghttp://event.udn.com/2012love/"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //19
		{ //20
			{
				Utf8:   "主辦單位",
				Big5:   []byte("\xa5D\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ":聯合便利貼   ",
				Big5:   []byte(":\xc1p\xa6X\xabK\xa7Q\xb6K   "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "協辦單位",
				Big5:   []byte("\xa8\xf3\xbf\xec\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ":可朵巧克力  ",
				Big5:   []byte(":\xa5i\xa6\xb7\xa5\xa9\xa7J\xa4O  "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
			{
				Utf8:   "合作單位",
				Big5:   []byte("\xa6X\xa7@\xb3\xe6\xa6\xec"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:   ":雨揚命理網站-開心就是開運",
				Big5:   []byte(":\xabB\xb4\xad\xa9R\xb2z\xba\xf4\xaf\xb8-\xb6}\xa4\xdf\xb4N\xacO\xb6}\xb9B"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //21
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 06:53)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 06:53)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //22
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //23
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //24
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //25
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //26
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //27
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //28
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //29
			{
				Utf8:   "※ 編輯: shiou7788       來自: 140.112.92.137       (12/22 07:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: shiou7788       \xa8\xd3\xa6\xdb: 140.112.92.137       (12/22 07:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //30
	}
}
