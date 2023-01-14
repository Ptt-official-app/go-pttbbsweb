package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testFilename20            = "M.1621089154.A.B39"
	testContentAll20          []byte
	testContent20             []byte
	testSignature20           []byte
	testComment20             []byte
	testFirstCommentsDBCS20   []byte
	testTheRestCommentsDBCS20 []byte
	testContent20Big5         [][]*types.Rune
	testContent20Utf8         [][]*types.Rune

	testFirstComments20       []*schema.Comment
	testFullFirstComments20   []*schema.Comment
	testTheRestComments20     []*schema.Comment
	testFullTheRestComments20 []*schema.Comment
)

func initTest20() {
	testContentAll20, testContent20, testSignature20, testComment20, testFirstCommentsDBCS20, testTheRestCommentsDBCS20 = loadTest(testFilename20)

	testContent20Big5 = [][]*types.Rune{
		{ // 0
			{
				Big5:   []byte("\xa7@\xaa\xcc: PttACT (\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1) \xac\xdd\xaaO: SYSOP"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: PttACT (\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1) \xac\xdd\xaaO: SYSOP"),
			},
		},
		{ // 1
			{
				Big5:   []byte("\xbc\xd0\xc3D: Re: [\xb1\xa1\xb3\xf8] \xac\xcc\xb1\xa1\xab\xe1\xaa\xba\xa8\xe0\xb5\xa3\xb8`\xa1K (\xa4\xa4\xbc\xfa\xa6W\xb3\xe6)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: Re: [\xb1\xa1\xb3\xf8] \xac\xcc\xb1\xa1\xab\xe1\xaa\xba\xa8\xe0\xb5\xa3\xb8`\xa1K (\xa4\xa4\xbc\xfa\xa6W\xb3\xe6)"),
			},
		},
		{ // 2
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu May  6 12:50:37 2021"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu May  6 12:50:37 2021"),
			},
		},
		{ // 3
		},
		{ // 4
			{
				Big5:   []byte("  \xb7P\xc1\xc2\xa8C\xa4@\xa6\xec\xa4\xe4\xab\xf9\xbbP\xc3\xf6\xa4\xdf\xa8\xe0\xb5\xa3\xb8`\xac\xa1\xb0\xca\xaa\xba\xaaO\xa5D\xbbP\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xb7P\xc1\xc2\xa8C\xa4@\xa6\xec\xa4\xe4\xab\xf9\xbbP\xc3\xf6\xa4\xdf\xa8\xe0\xb5\xa3\xb8`\xac\xa1\xb0\xca\xaa\xba\xaaO\xa5D\xbbP\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc"),
			},
		},
		{ // 5
			{
				Big5:   []byte("  \xac\xa1\xb0\xca\xb3\xa1\xafS\xa7O\xb7\xc7\xb3\xc6\xad\xad\xb6q\xaa\xba\xc5a\xbaq\xb3\xb3\xb2\xa1\xa7l\xa4\xf4\xaaM\xb9\xd4\xa7@\xac\xb0\xa9\xe2\xbc\xfa\xc2\xa7\xaa\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xac\xa1\xb0\xca\xb3\xa1\xafS\xa7O\xb7\xc7\xb3\xc6\xad\xad\xb6q\xaa\xba\xc5a\xbaq\xb3\xb3\xb2\xa1\xa7l\xa4\xf4\xaaM\xb9\xd4\xa7@\xac\xb0\xa9\xe2\xbc\xfa\xc2\xa7\xaa\xab"),
			},
		},
		{ // 6
			{
				Big5:   []byte("  https://i.imgur.com/8629B6y.jpg"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  https://i.imgur.com/8629B6y.jpg"),
			},
		},
		{ // 7
			{
				Big5:   []byte("  \xaaM\xb9\xd4\xb9\xcf\xae\xd7\xaa\xba\xb3]\xadp\xacO\xa8\xd3\xa6\xdb\xa9\xf3\xc3\xf6\xb7R\xa4\xa7\xaea\xa4p\xaaB\xa4\xcd\xad\xcc\xa9\xd2\xb5\xdb\xa6\xe2\xaa\xba\xc3\xb8\xb5e"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xaaM\xb9\xd4\xb9\xcf\xae\xd7\xaa\xba\xb3]\xadp\xacO\xa8\xd3\xa6\xdb\xa9\xf3\xc3\xf6\xb7R\xa4\xa7\xaea\xa4p\xaaB\xa4\xcd\xad\xcc\xa9\xd2\xb5\xdb\xa6\xe2\xaa\xba\xc3\xb8\xb5e"),
			},
		},
		{ // 8
		},
		{ // 9
			{
				Big5:   []byte("  \xc1\xc2\xc1\xc2\xa4j\xaea\xaa\xba\xa4\xc0\xa8\xc9\xa1A\xc5\xfd\xa7\xf3\xa6h\xa4H\xac\xdd\xa8\xa3\xab\xc4\xa4l\xad\xcc\xaa\xba\xbb\xdd\xadn"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xc1\xc2\xc1\xc2\xa4j\xaea\xaa\xba\xa4\xc0\xa8\xc9\xa1A\xc5\xfd\xa7\xf3\xa6h\xa4H\xac\xdd\xa8\xa3\xab\xc4\xa4l\xad\xcc\xaa\xba\xbb\xdd\xadn"),
			},
		},
		{ // 10
			{
				Big5:   []byte("  \xa7\xc6\xb1\xe6\xa7\xda\xad\xcc\xaa\xba\xc3\xf6\xa4\xdf\xbbP\xa4\xe4\xab\xf9\xa1A\xaf\xe0\xb1a\xb5\xb9\xab\xc4\xa4l\xad\xcc\xb7\xc5\xb7x\xa9M\xa4O\xb6q"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa7\xc6\xb1\xe6\xa7\xda\xad\xcc\xaa\xba\xc3\xf6\xa4\xdf\xbbP\xa4\xe4\xab\xf9\xa1A\xaf\xe0\xb1a\xb5\xb9\xab\xc4\xa4l\xad\xcc\xb7\xc5\xb7x\xa9M\xa4O\xb6q"),
			},
		},
		{ // 11
		},
		{ // 12
		},
		{ // 13
			{
				Big5:   []byte(" \xa5\xbb\xa6\xb8\xac\xa1\xb0\xca\xaa\xba\xa9\xe2\xbc\xfa\xb5\xb2\xaaG\xa6p\xa4U\xa1G "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCS:   []byte("\x1b[1;33;41m \xa5\xbb\xa6\xb8\xac\xa1\xb0\xca\xaa\xba\xa9\xe2\xbc\xfa\xb5\xb2\xaaG\xa6p\xa4U\xa1G "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},
		{ // 14

		},
		{ // 15
			{
				Big5:   []byte("\xb6W\xaf\xc5\xad\xad\xb6q\xb4\xda\xa1A\xabO\xc3\xd2\xb5\xb4\xaa\xa9\xb0\xb1\xb2\xa3\xa1A\xb0\xa3\xa4F\xa4\xf1\xc1\xc9\xabe\xa4\xad\xa6W\xa1A\xb4N\xa5u\xa6\xb3\xa7A\xa6\xb3\xb3\xe1\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@"),
			},
			{
				Big5:   []byte("\xa1\xbb \xaaO\xa5D\xaf\xb8\xa4\xba\xc2\xe0\xbf\xfd\xa1A\xa9\xe210\xa6W\xa1G"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;36m\xa1\xbb \xaaO\xa5D\xaf\xb8\xa4\xba\xc2\xe0\xbf\xfd\xa1A\xa9\xe210\xa6W\xa1G"),
			},
			{
				Big5:   []byte("\xb6W\xaf\xc5\xad\xad\xb6q\xb4\xda\xa1A\xabO\xc3\xd2\xb5\xb4\xaa\xa9\xb0\xb1\xb2\xa3\xa1A\xb0\xa3\xa4F\xa4\xf1\xc1\xc9\xabe\xa4\xad\xa6W\xa1A\xb4N\xa5u\xa6\xb3\xa7A\xa6\xb3\xb3\xe1\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},
		{ // 16
			{
				Big5:   []byte("\xa1@\xa1@ https://i.imgur.com/fo6ZVwB.gif"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@ https://i.imgur.com/fo6ZVwB.gif"),
			},
		},
		{ // 17
			{
				Big5:   []byte("     "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("     "),
			},
			{
				Big5:   []byte("(\xad\xab\xbd\xc6\xc2\xe0\xbf\xfd\xa5H\xb2\xc4\xa4@\xbdg\xa4\xe5\xb7\xc7)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;30m(\xad\xab\xbd\xc6\xc2\xe0\xbf\xfd\xa5H\xb2\xc4\xa4@\xbdg\xa4\xe5\xb7\xc7)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},
		{ // 18

		},
		{ // 19
			{
				Big5:   []byte("\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@"),
			},
			{
				Big5:   []byte("\xa1\xbb FB\xa4\xbd\xb6}\xa4\xc0\xa8\xc9\xb6K\xa4\xe5\xa1A\xa9\xe210\xa6W\xa1G"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Highlight: true},
				DBCS:   []byte("\x1b[1;36m\xa1\xbb FB\xa4\xbd\xb6}\xa4\xc0\xa8\xc9\xb6K\xa4\xe5\xa1A\xa9\xe210\xa6W\xa1G"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},

		{ // 20
			{
				Big5:   []byte("\xa1@\xa1@ https://i.imgur.com/IHeeTUo.gif"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@ https://i.imgur.com/IHeeTUo.gif"),
			},
		},

		{ // 21

		},
		{ // 22
			{
				Big5:   []byte("  \xa5H\xa4W\xbc\xfa\xc0y\xb1N\xa5\xd1\xac\xa1\xb0\xca\xb3\xa1\xa8\xf3\xa7U\xc1p\xc3\xb4\xbbP\xb5o\xa9\xf1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa5H\xa4W\xbc\xfa\xc0y\xb1N\xa5\xd1\xac\xa1\xb0\xca\xb3\xa1\xa8\xf3\xa7U\xc1p\xc3\xb4\xbbP\xb5o\xa9\xf1"),
			},
		},
		{ // 23
			{
				Big5:   []byte("  \xb9w\xadp\xa9\xf3 5/20 \xabe\xb7|\xb3v\xa4@\xa7\xb9\xa6\xa8\xc1p\xc3\xb4"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xb9w\xadp\xa9\xf3 5/20 \xabe\xb7|\xb3v\xa4@\xa7\xb9\xa6\xa8\xc1p\xc3\xb4"),
			},
		},
		{ // 24
			{
				Big5:   []byte("  \xa6p\xa6\xb3\xb0\xdd\xc3D\xa5i\xa6V\xac\xa1\xb0\xca\xb3\xa1 teemocogs \xc1p\xc3\xb4"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa6p\xa6\xb3\xb0\xdd\xc3D\xa5i\xa6V\xac\xa1\xb0\xca\xb3\xa1 teemocogs \xc1p\xc3\xb4"),
			},
		},
		{ // 25
			{
				Big5:   []byte("  \xa6A\xa6\xb8\xb7P\xc1\xc2\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc\xaa\xba\xb0\xd1\xbbP\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa6A\xa6\xb8\xb7P\xc1\xc2\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc\xaa\xba\xb0\xd1\xbbP\xa1I"),
			},
		},
		{ // 26

		},
		{ // 27

		},
		{ // 28
			{
				Big5:   []byte("                                  PTT \xac\xa1\xb0\xca\xb3\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                  PTT \xac\xa1\xb0\xca\xb3\xa1"),
			},
		},
	}

	testContent20Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:   "作者: PttACT (批踢踢活動部) 看板: SYSOP",
				Big5:   []byte("\xa7@\xaa\xcc: PttACT (\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1) \xac\xdd\xaaO: SYSOP"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: PttACT (\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1) \xac\xdd\xaaO: SYSOP"),
			},
		},
		{ // 1
			{
				Utf8:   "標題: Re: [情報] 疫情後的兒童節… (中獎名單)",
				Big5:   []byte("\xbc\xd0\xc3D: Re: [\xb1\xa1\xb3\xf8] \xac\xcc\xb1\xa1\xab\xe1\xaa\xba\xa8\xe0\xb5\xa3\xb8`\xa1K (\xa4\xa4\xbc\xfa\xa6W\xb3\xe6)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: Re: [\xb1\xa1\xb3\xf8] \xac\xcc\xb1\xa1\xab\xe1\xaa\xba\xa8\xe0\xb5\xa3\xb8`\xa1K (\xa4\xa4\xbc\xfa\xa6W\xb3\xe6)"),
			},
		},
		{ // 2
			{
				Utf8:   "時間: Thu May  6 12:50:37 2021",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu May  6 12:50:37 2021"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu May  6 12:50:37 2021"),
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:   "  感謝每一位支持與關心兒童節活動的板主與鄉民朋友們",
				Big5:   []byte("  \xb7P\xc1\xc2\xa8C\xa4@\xa6\xec\xa4\xe4\xab\xf9\xbbP\xc3\xf6\xa4\xdf\xa8\xe0\xb5\xa3\xb8`\xac\xa1\xb0\xca\xaa\xba\xaaO\xa5D\xbbP\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xb7P\xc1\xc2\xa8C\xa4@\xa6\xec\xa4\xe4\xab\xf9\xbbP\xc3\xf6\xa4\xdf\xa8\xe0\xb5\xa3\xb8`\xac\xa1\xb0\xca\xaa\xba\xaaO\xa5D\xbbP\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc"),
			},
		},
		{ // 5
			{
				Utf8:   "  活動部特別準備限量的鶯歌陶瓷吸水杯墊作為抽獎禮物",
				Big5:   []byte("  \xac\xa1\xb0\xca\xb3\xa1\xafS\xa7O\xb7\xc7\xb3\xc6\xad\xad\xb6q\xaa\xba\xc5a\xbaq\xb3\xb3\xb2\xa1\xa7l\xa4\xf4\xaaM\xb9\xd4\xa7@\xac\xb0\xa9\xe2\xbc\xfa\xc2\xa7\xaa\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xac\xa1\xb0\xca\xb3\xa1\xafS\xa7O\xb7\xc7\xb3\xc6\xad\xad\xb6q\xaa\xba\xc5a\xbaq\xb3\xb3\xb2\xa1\xa7l\xa4\xf4\xaaM\xb9\xd4\xa7@\xac\xb0\xa9\xe2\xbc\xfa\xc2\xa7\xaa\xab"),
			},
		},
		{ // 6
			{
				Utf8:   "  https://i.imgur.com/8629B6y.jpg",
				Big5:   []byte("  https://i.imgur.com/8629B6y.jpg"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  https://i.imgur.com/8629B6y.jpg"),
			},
		},
		{ // 7
			{
				Utf8:   "  杯墊圖案的設計是來自於關愛之家小朋友們所著色的繪畫",
				Big5:   []byte("  \xaaM\xb9\xd4\xb9\xcf\xae\xd7\xaa\xba\xb3]\xadp\xacO\xa8\xd3\xa6\xdb\xa9\xf3\xc3\xf6\xb7R\xa4\xa7\xaea\xa4p\xaaB\xa4\xcd\xad\xcc\xa9\xd2\xb5\xdb\xa6\xe2\xaa\xba\xc3\xb8\xb5e"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xaaM\xb9\xd4\xb9\xcf\xae\xd7\xaa\xba\xb3]\xadp\xacO\xa8\xd3\xa6\xdb\xa9\xf3\xc3\xf6\xb7R\xa4\xa7\xaea\xa4p\xaaB\xa4\xcd\xad\xcc\xa9\xd2\xb5\xdb\xa6\xe2\xaa\xba\xc3\xb8\xb5e"),
			},
		},
		{ // 8
		},
		{ // 9
			{
				Utf8:   "  謝謝大家的分享，讓更多人看見孩子們的需要",
				Big5:   []byte("  \xc1\xc2\xc1\xc2\xa4j\xaea\xaa\xba\xa4\xc0\xa8\xc9\xa1A\xc5\xfd\xa7\xf3\xa6h\xa4H\xac\xdd\xa8\xa3\xab\xc4\xa4l\xad\xcc\xaa\xba\xbb\xdd\xadn"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xc1\xc2\xc1\xc2\xa4j\xaea\xaa\xba\xa4\xc0\xa8\xc9\xa1A\xc5\xfd\xa7\xf3\xa6h\xa4H\xac\xdd\xa8\xa3\xab\xc4\xa4l\xad\xcc\xaa\xba\xbb\xdd\xadn"),
			},
		},
		{ // 10
			{
				Utf8:   "  希望我們的關心與支持，能帶給孩子們溫暖和力量",
				Big5:   []byte("  \xa7\xc6\xb1\xe6\xa7\xda\xad\xcc\xaa\xba\xc3\xf6\xa4\xdf\xbbP\xa4\xe4\xab\xf9\xa1A\xaf\xe0\xb1a\xb5\xb9\xab\xc4\xa4l\xad\xcc\xb7\xc5\xb7x\xa9M\xa4O\xb6q"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa7\xc6\xb1\xe6\xa7\xda\xad\xcc\xaa\xba\xc3\xf6\xa4\xdf\xbbP\xa4\xe4\xab\xf9\xa1A\xaf\xe0\xb1a\xb5\xb9\xab\xc4\xa4l\xad\xcc\xb7\xc5\xb7x\xa9M\xa4O\xb6q"),
			},
		},
		{ // 11

		},
		{ // 12
		},
		{ // 13
			{
				Utf8:   " 本次活動的抽獎結果如下： ",
				Big5:   []byte(" \xa5\xbb\xa6\xb8\xac\xa1\xb0\xca\xaa\xba\xa9\xe2\xbc\xfa\xb5\xb2\xaaG\xa6p\xa4U\xa1G "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_RED, Highlight: true},
				DBCS:   []byte("\x1b[1;33;41m \xa5\xbb\xa6\xb8\xac\xa1\xb0\xca\xaa\xba\xa9\xe2\xbc\xfa\xb5\xb2\xaaG\xa6p\xa4U\xa1G "),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},
		{ // 14

		},
		{ // 15
			{
				Utf8:   "　",
				Big5:   []byte("\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@"),
			},
			{
				Utf8:   "◆ 板主站內轉錄，抽10名：",
				Big5:   []byte("\xa1\xbb \xaaO\xa5D\xaf\xb8\xa4\xba\xc2\xe0\xbf\xfd\xa1A\xa9\xe210\xa6W\xa1G"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;36m\xa1\xbb \xaaO\xa5D\xaf\xb8\xa4\xba\xc2\xe0\xbf\xfd\xa1A\xa9\xe210\xa6W\xa1G"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},
		{ // 16
			{
				Utf8:   "　　 https://i.imgur.com/fo6ZVwB.gif",
				Big5:   []byte("\xa1@\xa1@ https://i.imgur.com/fo6ZVwB.gif"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@ https://i.imgur.com/fo6ZVwB.gif"),
			},
		},
		{ // 17
			{
				Utf8:   "     ",
				Big5:   []byte("     "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("     "),
			},
			{
				Utf8:   "(重複轉錄以第一篇文準)",
				Big5:   []byte("(\xad\xab\xbd\xc6\xc2\xe0\xbf\xfd\xa5H\xb2\xc4\xa4@\xbdg\xa4\xe5\xb7\xc7)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_BLACK, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;30m(\xad\xab\xbd\xc6\xc2\xe0\xbf\xfd\xa5H\xb2\xc4\xa4@\xbdg\xa4\xe5\xb7\xc7)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},
		{ // 18

		},
		{ // 19
			{
				Utf8:   "　",
				Big5:   []byte("\xa1@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@"),
			},
			{
				Utf8:   "◆ FB公開分享貼文，抽10名：",
				Big5:   []byte("\xa1\xbb FB\xa4\xbd\xb6}\xa4\xc0\xa8\xc9\xb6K\xa4\xe5\xa1A\xa9\xe210\xa6W\xa1G"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;36m\xa1\xbb FB\xa4\xbd\xb6}\xa4\xc0\xa8\xc9\xb6K\xa4\xe5\xa1A\xa9\xe210\xa6W\xa1G"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m"),
			},
		},

		{ // 20
			{
				Utf8:   "　　 https://i.imgur.com/IHeeTUo.gif",
				Big5:   []byte("\xa1@\xa1@ https://i.imgur.com/IHeeTUo.gif"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1@\xa1@ https://i.imgur.com/IHeeTUo.gif"),
			},
		},

		{ // 21

		},
		{ // 22
			{
				Utf8:   "  以上獎勵將由活動部協助聯繫與發放",
				Big5:   []byte("  \xa5H\xa4W\xbc\xfa\xc0y\xb1N\xa5\xd1\xac\xa1\xb0\xca\xb3\xa1\xa8\xf3\xa7U\xc1p\xc3\xb4\xbbP\xb5o\xa9\xf1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa5H\xa4W\xbc\xfa\xc0y\xb1N\xa5\xd1\xac\xa1\xb0\xca\xb3\xa1\xa8\xf3\xa7U\xc1p\xc3\xb4\xbbP\xb5o\xa9\xf1"),
			},
		},
		{ // 23
			{
				Utf8:   "  預計於 5/20 前會逐一完成聯繫",
				Big5:   []byte("  \xb9w\xadp\xa9\xf3 5/20 \xabe\xb7|\xb3v\xa4@\xa7\xb9\xa6\xa8\xc1p\xc3\xb4"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xb9w\xadp\xa9\xf3 5/20 \xabe\xb7|\xb3v\xa4@\xa7\xb9\xa6\xa8\xc1p\xc3\xb4"),
			},
		},
		{ // 24
			{
				Utf8:   "  如有問題可向活動部 teemocogs 聯繫",
				Big5:   []byte("  \xa6p\xa6\xb3\xb0\xdd\xc3D\xa5i\xa6V\xac\xa1\xb0\xca\xb3\xa1 teemocogs \xc1p\xc3\xb4"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa6p\xa6\xb3\xb0\xdd\xc3D\xa5i\xa6V\xac\xa1\xb0\xca\xb3\xa1 teemocogs \xc1p\xc3\xb4"),
			},
		},
		{ // 25
			{
				Utf8:   "  再次感謝鄉民朋友們的參與！",
				Big5:   []byte("  \xa6A\xa6\xb8\xb7P\xc1\xc2\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc\xaa\xba\xb0\xd1\xbbP\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa6A\xa6\xb8\xb7P\xc1\xc2\xb6m\xa5\xc1\xaaB\xa4\xcd\xad\xcc\xaa\xba\xb0\xd1\xbbP\xa1I"),
			},
		},
		{ // 26

		},
		{ // 27

		},
		{ // 28
			{
				Utf8:   "                                  PTT 活動部",
				Big5:   []byte("                                  PTT \xac\xa1\xb0\xca\xb3\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("                                  PTT \xac\xa1\xb0\xca\xb3\xa1"),
			},
		},
	}

	testFirstComments20 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ericf129"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "辛苦了 ><",
						Big5:   []byte("\xa8\xaf\xadW\xa4F ><                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa8\xaf\xadW\xa4F ><                             "),
					},
				},
			},
			MD5:     "7GEtBNWCH8N6QsG9aAGoLA",
			IP:      "1.200.29.12",
			TheDate: "05/06 22:57",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mericf129\x1b[m\x1b[33m: \xa8\xaf\xadW\xa4F ><                             \x1b[m    1.200.29.12 05/06 22:57"),
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("zhibb"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "謝謝活動部 :)",
						Big5:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1 :)                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1 :)                            "),
					},
				},
			},
			IP:      "140.113.0.229",
			TheDate: "05/07 09:49",
			MD5:     "9fJnW9k0qnOtNnani_mlaw",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mzhibb\x1b[m\x1b[33m: \xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1 :)                            \x1b[m  140.113.0.229 05/07 09:49"),
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Japan2001"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "謝謝活動部",
						Big5:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1                           "),
					},
				},
			},
			MD5:     "SLDr8nBNFVjFaxS6C9iG-Q",
			IP:      "180.214.183.155",
			TheDate: "05/08 18:57",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mJapan2001\x1b[m\x1b[33m: \xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1                           \x1b[m180.214.183.155 05/08 18:57"),
		},
	}

	testFullFirstComments20 = []*schema.Comment{
		{ // 0
			BBoardID:   "test",
			ArticleID:  "test20",
			CommentID:  "Fn9D_yV31kA:7GEtBNWCH8N6QsG9aAGoLA",
			CreateTime: 1620313020000000000,
			SortTime:   1621089154001000000,
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("ericf129"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "辛苦了 ><",
						Big5:   []byte("\xa8\xaf\xadW\xa4F ><                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa8\xaf\xadW\xa4F ><                             "),
					},
				},
			},
			MD5:     "7GEtBNWCH8N6QsG9aAGoLA",
			IP:      "1.200.29.12",
			TheDate: "05/06 22:57",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mericf129\x1b[m\x1b[33m: \xa8\xaf\xadW\xa4F ><                             \x1b[m    1.200.29.12 05/06 22:57"),
		},
		{ // 1
			BBoardID:   "test",
			ArticleID:  "test20",
			CommentID:  "Fn9D_yWHGIA:9fJnW9k0qnOtNnani_mlaw",
			CreateTime: 1620352140000000000,
			SortTime:   1621089154002000000,
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("zhibb"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "謝謝活動部 :)",
						Big5:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1 :)                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1 :)                            "),
					},
				},
			},
			IP:      "140.113.0.229",
			TheDate: "05/07 09:49",
			MD5:     "9fJnW9k0qnOtNnani_mlaw",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mzhibb\x1b[m\x1b[33m: \xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1 :)                            \x1b[m  140.113.0.229 05/07 09:49"),
		},
		{ // 2
			BBoardID:   "test",
			ArticleID:  "test20",
			CommentID:  "Fn0SK73F2AA:SLDr8nBNFVjFaxS6C9iG-Q",
			CreateTime: 1620471420000000000,
			SortTime:   1620471420000000000,
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Japan2001"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "謝謝活動部",
						Big5:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1                           "),
					},
				},
			},
			MD5:     "SLDr8nBNFVjFaxS6C9iG-Q",
			IP:      "180.214.183.155",
			TheDate: "05/08 18:57",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mJapan2001\x1b[m\x1b[33m: \xc1\xc2\xc1\xc2\xac\xa1\xb0\xca\xb3\xa1                           \x1b[m180.214.183.155 05/08 18:57"),
		},
	}

	testTheRestComments20 = []*schema.Comment{}
}
