package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testFilename19            = "temp14"
	testContentAll19          []byte
	testContent19             []byte
	testSignature19           []byte
	testComment19             []byte
	testFirstCommentsDBCS19   []byte
	testTheRestCommentsDBCS19 []byte
	testContent19Big5         [][]*types.Rune
	testContent19Utf8         [][]*types.Rune

	testFirstComments19       []*schema.Comment
	testFullFirstComments19   []*schema.Comment
	testTheRestComments19     []*schema.Comment
	testFullTheRestComments19 []*schema.Comment
)

func initTest19() {
	testContentAll19, testContent19, testSignature19, testComment19, testFirstCommentsDBCS19, testTheRestCommentsDBCS19 = loadTest(testFilename19)

	testContent19Big5 = [][]*types.Rune{
		{ // 0
			{

				Big5:   []byte("\xa7@\xaa\xcc: PttACT (PTT\xac\xa1\xb0\xca\xb3\xa1\xb1M\xa5\xce\xb1b\xb8\xb9) \xac\xdd\xaaO: SYSOP"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: PttACT (PTT\xac\xa1\xb0\xca\xb3\xa1\xb1M\xa5\xce\xb1b\xb8\xb9) \xac\xdd\xaaO: SYSOP\r"),
			},
		},
		{ // 1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9 \xa8M\xbf\xef\xa7\xeb\xb2\xbc\xa9\xe2\xbc\xfa\xab~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9 \xa8M\xbf\xef\xa7\xeb\xb2\xbc\xa9\xe2\xbc\xfa\xab~\r"),
			},
		},
		{ // 2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Jan 26 17:18:22 2018"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Jan 26 17:18:22 2018\r"),
			},
		},
		{ // 3
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 4
			{
				Big5:   []byte("\xa5\xb4\xb3\xc2\xb1N\xbaN\xb4X\xb0\xe9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9\xa8M\xbf\xef\xb2\xd7\xa9\xf3\xb6}\xa9l\xc5o\xa1I\xa7Y\xb0_\xb6i\xa4J\xa8M\xbf\xef\xb8\xd5\xc5\xa5\xb6\xa5\xacq\xa1I\r"),
			},
		},
		{ // 5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 6
			{
				Big5:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb0\xd1\xc1\xc9\xaa\xcc\xa4Q\xb1j\xaa\xba\xb2\xc4\xa4G\xad\xba\xa6\xb1\xa4l\xa4w\xa9\xf3 \x1b[36;1mOriginalSong\xaaO\x1b[m \xc4\xc0\xa5X\r"),
			},
		},
		{ // 8
			{

				Big5:   []byte("\xb6R\xaa\xd1\xb2\xbc\xa6s\xaa\xd1\xbe\xde\xbdL"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xeb\xb2\xbc\xa4\xe9\xb4\xc1\xb9w\xadp\xac\xb0\xa1G1/29 ~ 2/9\r"),
			},
		},
		{ // 9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 10
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 11
			{
				Big5:   []byte("\xa8M\xbf\xef\xb1\xc0\xa4\xe5\xb2\xbc\xbf\xef\xb0\xa3\xa4F\x1b[33;1m\xc0H\xbe\xf7\xbf\xef\xa4\xad\xa6\xec\xb6m\xa5\xc1\xb5o\xa9\xf1 3000 \xa7\xe5\xb9\xf4\x1b[m"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa8M\xbf\xef\xb1\xc0\xa4\xe5\xb2\xbc\xbf\xef\xb0\xa3\xa4F\x1b[33;1m\xc0H\xbe\xf7\xbf\xef\xa4\xad\xa6\xec\xb6m\xa5\xc1\xb5o\xa9\xf1 3000 \xa7\xe5\xb9\xf4\x1b[m\r"),
			},
		},
		{ // 12
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 13
			{
				Big5: []byte("\xa8M\xbf\xef\xb1\xc0\xa4\xe5\xb2\xbc\xbf\xef\xb0\xa3\xa4F\x1b[33;1m\xc0H\xbe\xf7\xbf\xef\xa4\xad\xa6\xec\xb6m\xa5\xc1\xb5o\xa9\xf1 3000 \xa7\xe5\xb9\xf4\x1b[m"), Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5t\xa5~\xa6A\xa5[\xbdX\xa1A\xc0H\xbe\xf7\xa9\xe2\xa5X\xa4@\xa6\xec\xb0\xd1\xbbP\xa8M\xbf\xef\xb2\xbc\xbf\xef\xaa\xba\xb6m\xa5\xc1\r"),
			},
		},
		{ // 14
			{
				Big5:   []byte("\x1b[33;1m\xc3\xd8\xb0e\xac\xf6\xa9\xc0\xab~\xa1u\xaf\xc2\xa4\xe2\xa4u\xa6L\xa8\xea\xa1DPTT \xafS\xa7O\xaa\xa9 LOGO \xa7\xf4\xa4f\xb3U\xa1v\xa4@\xad\xd3\xa1I\xa1I\x1b[m"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[33;1m\xc3\xd8\xb0e\xac\xf6\xa9\xc0\xab~\xa1u\xaf\xc2\xa4\xe2\xa4u\xa6L\xa8\xea\xa1DPTT \xafS\xa7O\xaa\xa9 LOGO \xa7\xf4\xa4f\xb3U\xa1v\xa4@\xad\xd3\xa1I\xa1I\x1b[m\r"),
			},
		},
		{ // 15
			{
				Big5:   []byte("\xb6W\xaf\xc5\xad\xad\xb6q\xb4\xda\xa1A\xabO\xc3\xd2\xb5\xb4\xaa\xa9\xb0\xb1\xb2\xa3\xa1A\xb0\xa3\xa4F\xa4\xf1\xc1\xc9\xabe\xa4\xad\xa6W\xa1A\xb4N\xa5u\xa6\xb3\xa7A\xa6\xb3\xb3\xe1\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6W\xaf\xc5\xad\xad\xb6q\xb4\xda\xa1A\xabO\xc3\xd2\xb5\xb4\xaa\xa9\xb0\xb1\xb2\xa3\xa1A\xb0\xa3\xa4F\xa4\xf1\xc1\xc9\xabe\xa4\xad\xa6W\xa1A\xb4N\xa5u\xa6\xb3\xa7A\xa6\xb3\xb3\xe1\xa1I\r"),
			},
		},
		{ // 16
			{
				Big5:   []byte("\xa6p\xb9\xcf\xa1Ghttps://imgur.com/a/yeRfy"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6p\xb9\xcf\xa1Ghttps://imgur.com/a/yeRfy\r"),
			},
		},
		{ // 17
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 18
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 16
			{
				Big5:   []byte("\xa4\xdf\xb0\xca\xa4F\xb6\xdc\xa1H\xb8\xd4\xb1\xa1\xa7\xd6\xa8\xec OriginalSong\xaaO \xa5\xfd\xc5\xa5\xac\xb0\xa7\xd6\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4\xdf\xb0\xca\xa4F\xb6\xdc\xa1H\xb8\xd4\xb1\xa1\xa7\xd6\xa8\xec OriginalSong\xaaO \xa5\xfd\xc5\xa5\xac\xb0\xa7\xd6\xa1I\r"),
			},
		},
		{ // 17
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 18
			{
				Big5:   []byte("PTT\xac\xa1\xb0\xca\xb3\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("PTT\xac\xa1\xb0\xca\xb3\xa1\r"),
			},
		},
		{ // 19
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 20
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ // 21
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 22
			{
				Big5:   []byte("\xa5\xab\xa5\xc1\xbcs\xb3\xf5     \xb3\xf8\xa7i\xaf\xb8\xaa\xf8 PTT\xabr\xa7\xda"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xab\xa5\xc1\xbcs\xb3\xf5     \xb3\xf8\xa7i\xaf\xb8\xaa\xf8 PTT\xabr\xa7\xda\r"),
			},
		},
		{ // 23
			{
				Big5:   []byte("   PttAct        \xa5\xbb\xaf\xb8 \xa3U\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("   PttAct        \xa5\xbb\xaf\xb8 \xa3U\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1\r"),
			},
		},
		{ // 24
			{
				Big5:   []byte("     OriginalSong  \xad\xec\xb3\xd0 \xa1\xb72017 PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("     OriginalSong  \xad\xec\xb3\xd0 \xa1\xb72017 PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9\r"),
			},
		},
	}

	testContent19Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:   "作者: PttACT (PTT活動部專用帳號) 看板: SYSOP",
				Big5:   []byte("\xa7@\xaa\xcc: PttACT (PTT\xac\xa1\xb0\xca\xb3\xa1\xb1M\xa5\xce\xb1b\xb8\xb9) \xac\xdd\xaaO: SYSOP"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: PttACT (PTT\xac\xa1\xb0\xca\xb3\xa1\xb1M\xa5\xce\xb1b\xb8\xb9) \xac\xdd\xaaO: SYSOP\r"),
			},
		},
		{ // 1
			{
				Utf8:   "標題: [情報] PTT自創曲大賽 決選投票抽獎品",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9 \xa8M\xbf\xef\xa7\xeb\xb2\xbc\xa9\xe2\xbc\xfa\xab~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb1\xa1\xb3\xf8] PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9 \xa8M\xbf\xef\xa7\xeb\xb2\xbc\xa9\xe2\xbc\xfa\xab~\r"),
			},
		},
		{ // 2
			{
				Utf8:   "時間: Fri Jan 26 17:18:22 2018",
				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Jan 26 17:18:22 2018"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Jan 26 17:18:22 2018\r"),
			},
		},
		{ // 3
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 4
			{
				Utf8:   "PTT自創曲大賽決選終於開始囉！即起進入決選試聽階段！",
				Big5:   []byte("PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9\xa8M\xbf\xef\xb2\xd7\xa9\xf3\xb6}\xa9l\xc5o\xa1I\xa7Y\xb0_\xb6i\xa4J\xa8M\xbf\xef\xb8\xd5\xc5\xa5\xb6\xa5\xacq\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9\xa8M\xbf\xef\xb2\xd7\xa9\xf3\xb6}\xa9l\xc5o\xa1I\xa7Y\xb0_\xb6i\xa4J\xa8M\xbf\xef\xb8\xd5\xc5\xa5\xb6\xa5\xacq\xa1I\r"),
			},
		},
		{ // 5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 6
			{
				Utf8:   "參賽者十強的第二首曲子已於 ",
				Big5:   []byte("\xb0\xd1\xc1\xc9\xaa\xcc\xa4Q\xb1j\xaa\xba\xb2\xc4\xa4G\xad\xba\xa6\xb1\xa4l\xa4w\xa9\xf3 "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb0\xd1\xc1\xc9\xaa\xcc\xa4Q\xb1j\xaa\xba\xb2\xc4\xa4G\xad\xba\xa6\xb1\xa4l\xa4w\xa9\xf3 "),
			},
			{
				Utf8:   "OriginalSong板",
				Big5:   []byte("OriginalSong\xaaO"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[36;1mOriginalSong\xaaO"),
			},
			{
				Utf8:   " 釋出",
				Big5:   []byte(" \xc4\xc0\xa5X"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m \xc4\xc0\xa5X\r"),
			},
		},
		{ // 7
			{
				Utf8:   "投票日期預計為：1/29 ~ 2/9",
				Big5:   []byte("\xa7\xeb\xb2\xbc\xa4\xe9\xb4\xc1\xb9w\xadp\xac\xb0\xa1G1/29 ~ 2/9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xeb\xb2\xbc\xa4\xe9\xb4\xc1\xb9w\xadp\xac\xb0\xa1G1/29 ~ 2/9\r"),
			},
		},
		{ // 8
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 10
			{
				Utf8:   "決選推文票選除了",
				Big5:   []byte("\xa8M\xbf\xef\xb1\xc0\xa4\xe5\xb2\xbc\xbf\xef\xb0\xa3\xa4F"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa8M\xbf\xef\xb1\xc0\xa4\xe5\xb2\xbc\xbf\xef\xb0\xa3\xa4F"),
			},
			{
				Utf8:   "隨機選五位鄉民發放 3000 批幣",
				Big5:   []byte("\xc0H\xbe\xf7\xbf\xef\xa4\xad\xa6\xec\xb6m\xa5\xc1\xb5o\xa9\xf1 3000 \xa7\xe5\xb9\xf4"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33;1m\xc0H\xbe\xf7\xbf\xef\xa4\xad\xa6\xec\xb6m\xa5\xc1\xb5o\xa9\xf1 3000 \xa7\xe5\xb9\xf4"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 12
			{
				Utf8:   "另外再加碼，隨機抽出一位參與決選票選的鄉民",
				Big5:   []byte("\xa5t\xa5~\xa6A\xa5[\xbdX\xa1A\xc0H\xbe\xf7\xa9\xe2\xa5X\xa4@\xa6\xec\xb0\xd1\xbbP\xa8M\xbf\xef\xb2\xbc\xbf\xef\xaa\xba\xb6m\xa5\xc1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5t\xa5~\xa6A\xa5[\xbdX\xa1A\xc0H\xbe\xf7\xa9\xe2\xa5X\xa4@\xa6\xec\xb0\xd1\xbbP\xa8M\xbf\xef\xb2\xbc\xbf\xef\xaa\xba\xb6m\xa5\xc1\r"),
			},
		},
		{ // 13
			{
				Utf8:   "贈送紀念品「純手工印刷．PTT 特別版 LOGO 束口袋」一個！！",
				Big5:   []byte("\xc3\xd8\xb0e\xac\xf6\xa9\xc0\xab~\xa1u\xaf\xc2\xa4\xe2\xa4u\xa6L\xa8\xea\xa1DPTT \xafS\xa7O\xaa\xa9 LOGO \xa7\xf4\xa4f\xb3U\xa1v\xa4@\xad\xd3\xa1I\xa1I"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Highlight: true, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Highlight: true, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[33;1m\xc3\xd8\xb0e\xac\xf6\xa9\xc0\xab~\xa1u\xaf\xc2\xa4\xe2\xa4u\xa6L\xa8\xea\xa1DPTT \xafS\xa7O\xaa\xa9 LOGO \xa7\xf4\xa4f\xb3U\xa1v\xa4@\xad\xd3\xa1I\xa1I"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 14
			{
				Utf8:   "超級限量款，保證絕版停產，除了比賽前五名，就只有你有喔！",
				Big5:   []byte("\xb6W\xaf\xc5\xad\xad\xb6q\xb4\xda\xa1A\xabO\xc3\xd2\xb5\xb4\xaa\xa9\xb0\xb1\xb2\xa3\xa1A\xb0\xa3\xa4F\xa4\xf1\xc1\xc9\xabe\xa4\xad\xa6W\xa1A\xb4N\xa5u\xa6\xb3\xa7A\xa6\xb3\xb3\xe1\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6W\xaf\xc5\xad\xad\xb6q\xb4\xda\xa1A\xabO\xc3\xd2\xb5\xb4\xaa\xa9\xb0\xb1\xb2\xa3\xa1A\xb0\xa3\xa4F\xa4\xf1\xc1\xc9\xabe\xa4\xad\xa6W\xa1A\xb4N\xa5u\xa6\xb3\xa7A\xa6\xb3\xb3\xe1\xa1I\r"),
			},
		},
		{ // 16
			{
				Utf8:   "如圖：https://imgur.com/a/yeRfy",
				Big5:   []byte("\xa6p\xb9\xcf\xa1Ghttps://imgur.com/a/yeRfy"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6p\xb9\xcf\xa1Ghttps://imgur.com/a/yeRfy\r"),
			},
		},
		{ // 17
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 18
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 16
			{
				Utf8:   "心動了嗎？詳情快到 OriginalSong板 先聽為快！",
				Big5:   []byte("\xa4\xdf\xb0\xca\xa4F\xb6\xdc\xa1H\xb8\xd4\xb1\xa1\xa7\xd6\xa8\xec OriginalSong\xaaO \xa5\xfd\xc5\xa5\xac\xb0\xa7\xd6\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4\xdf\xb0\xca\xa4F\xb6\xdc\xa1H\xb8\xd4\xb1\xa1\xa7\xd6\xa8\xec OriginalSong\xaaO \xa5\xfd\xc5\xa5\xac\xb0\xa7\xd6\xa1I\r"),
			},
		},
		{ // 17
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 18
			{
				Utf8:   "PTT活動部",
				Big5:   []byte("PTT\xac\xa1\xb0\xca\xb3\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("PTT\xac\xa1\xb0\xca\xb3\xa1\r"),
			},
		},
		{ // 19
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 20
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ // 21
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 22
			{
				Utf8:   "市民廣場     報告站長 PTT咬我",
				Big5:   []byte("\xa5\xab\xa5\xc1\xbcs\xb3\xf5     \xb3\xf8\xa7i\xaf\xb8\xaa\xf8 PTT\xabr\xa7\xda"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xab\xa5\xc1\xbcs\xb3\xf5     \xb3\xf8\xa7i\xaf\xb8\xaa\xf8 PTT\xabr\xa7\xda\r"),
			},
		},
		{ // 23
			{
				Utf8:   "   PttAct        本站 Σ批踢踢活動部",
				Big5:   []byte("   PttAct        \xa5\xbb\xaf\xb8 \xa3U\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("   PttAct        \xa5\xbb\xaf\xb8 \xa3U\xa7\xe5\xbd\xf0\xbd\xf0\xac\xa1\xb0\xca\xb3\xa1\r"),
			},
		},
		{ // 24
			{
				Utf8:   "     OriginalSong  原創 ◎2017 PTT自創曲大賽",
				Big5:   []byte("     OriginalSong  \xad\xec\xb3\xd0 \xa1\xb72017 PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("     OriginalSong  \xad\xec\xb3\xd0 \xa1\xb72017 PTT\xa6\xdb\xb3\xd0\xa6\xb1\xa4j\xc1\xc9\r"),
			},
		},
	}

	testFirstComments19 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "OriginalSong",
						Big5:   []byte("OriginalSong"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("OriginalSong"),
					},
				},
			},
			MD5:     "7QtBKKVmDogg0WyEsOMG0g",
			TheDate: "01/26 17:19",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mPttACT\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO OriginalSong\x1b[m                                  01/26 17:19\r"),
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "PttEarnMoney",
						Big5:   []byte("PttEarnMoney"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("PttEarnMoney"),
					},
				},
			},
			TheDate: "01/26 17:50",
			MD5:     "UxxLgfh1OKFy1snEV-4fTA",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mPttACT\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO PttEarnMoney\x1b[m                                  01/26 17:50\r"),
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Gossiping",
						Big5:   []byte("Gossiping"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Gossiping"),
					},
				},
			},
			MD5:     "YvzzA1oPoZ2WX1PZDVxNXg",
			TheDate: "01/26 22:25",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32msimons\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Gossiping\x1b[m                                     01/26 22:25\r"),
		},
		{ // 3
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "WomenTalk",
						Big5:   []byte("WomenTalk"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("WomenTalk"),
					},
				},
			},
			MD5:     "ZfEvXcD02JwSuTXREa6l5g",
			TheDate: "01/26 22:34",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32msimons\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO WomenTalk\x1b[m                                     01/26 22:34\r"),
		},
		{ // 4
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("mono5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "talk",
						Big5:   []byte("talk"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("talk"),
					},
				},
			},
			MD5:     "Ikc3AvccHX6YWe252vt5-w",
			TheDate: "01/26 22:39",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mmono5566\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO talk\x1b[m                                        01/26 22:39\r"),
		},
		{ // 5
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("a6234709"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "marvel",
						Big5:   []byte("marvel"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("marvel"),
					},
				},
			},
			MD5:     "2yN-w0xN9lsYzx1lbtZM-Q",
			TheDate: "01/27 00:43",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32ma6234709\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO marvel\x1b[m                                      01/27 00:43\r"),
		},
		{ // 6
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "KMU",
						Big5:   []byte("KMU"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("KMU"),
					},
				},
			},
			MD5:     "7zFE3bNz01gCEmX-hg4TxA",
			TheDate: "01/28 12:10",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mgghost1002\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO KMU\x1b[m                                       01/28 12:10\r"),
		},
		{ // 7
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Kaohsiung",
						Big5:   []byte("Kaohsiung"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Kaohsiung"),
					},
				},
			},
			MD5:     "AegTTCs475pm0Zx5M4ErTA",
			TheDate: "01/28 12:11",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mgghost1002\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Kaohsiung\x1b[m                                 01/28 12:11\r"),
		},
		{ // 8
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("fennyccc"),
			MD5:     "a8ZbQMrVxHxbUMd4IvtPcA",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "讚",
						Big5:   []byte("\xc6g                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc6g                                    "),
					},
				},
			},
			IP:      "115.82.149.161",
			TheDate: "01/28 12:34",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mfennyccc\x1b[m\x1b[33m: \xc6g                                    \x1b[m 115.82.149.161 01/28 12:34\r"),
		},
		{ // 9
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "某隱形看板",
						Big5:   []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO"),
					},
				},
			},
			MD5:     "unnUE1r9jqOU-NdpMHfWwA",
			TheDate: "01/29 02:39",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO\x1b[m                                         01/29 02:39\r"),
		},
	}

	testFullFirstComments19 = []*schema.Comment{
		{ // 0
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "OriginalSong",
						Big5:   []byte("OriginalSong"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("OriginalSong"),
					},
				},
			},
			MD5:        "7QtBKKVmDogg0WyEsOMG0g",
			TheDate:    "01/26 17:19",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mPttACT\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO OriginalSong\x1b[m                                  01/26 17:19\r"),
			CommentID:  "FQ1RkrLEKAA:7QtBKKVmDogg0WyEsOMG0g",
			CreateTime: 1516958340000000000,
			SortTime:   1516958340000000000,
		},
		{ // 1
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "PttEarnMoney",
						Big5:   []byte("PttEarnMoney"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("PttEarnMoney"),
					},
				},
			},
			TheDate:    "01/26 17:50",
			MD5:        "UxxLgfh1OKFy1snEV-4fTA",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mPttACT\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO PttEarnMoney\x1b[m                                  01/26 17:50\r"),
			CommentID:  "FQ1TQ8Nn0AA:UxxLgfh1OKFy1snEV-4fTA",
			CreateTime: 1516960200000000000,
			SortTime:   1516960200000000000,
		},
		{ // 2
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Gossiping",
						Big5:   []byte("Gossiping"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Gossiping"),
					},
				},
			},
			MD5:        "YvzzA1oPoZ2WX1PZDVxNXg",
			TheDate:    "01/26 22:25",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32msimons\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Gossiping\x1b[m                                     01/26 22:25\r"),
			CommentID:  "FQ1iRXgLWAA:YvzzA1oPoZ2WX1PZDVxNXg",
			CreateTime: 1516976700000000000,
			SortTime:   1516976700000000000,
		},
		{ // 3
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "WomenTalk",
						Big5:   []byte("WomenTalk"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("WomenTalk"),
					},
				},
			},
			MD5:        "ZfEvXcD02JwSuTXREa6l5g",
			TheDate:    "01/26 22:34",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32msimons\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO WomenTalk\x1b[m                                     01/26 22:34\r"),
			CommentID:  "FQ1iwzKNcAA:ZfEvXcD02JwSuTXREa6l5g",
			CreateTime: 1516977240000000000,
			SortTime:   1516977240000000000,
		},
		{ // 4
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("mono5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "talk",
						Big5:   []byte("talk"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("talk"),
					},
				},
			},
			MD5:        "Ikc3AvccHX6YWe252vt5-w",
			TheDate:    "01/26 22:39",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mmono5566\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO talk\x1b[m                                        01/26 22:39\r"),
			CommentID:  "FQ1jCQvyKAA:Ikc3AvccHX6YWe252vt5-w",
			CreateTime: 1516977540000000000,
			SortTime:   1516977540000000000,
		},
		{ // 5
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("a6234709"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "marvel",
						Big5:   []byte("marvel"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("marvel"),
					},
				},
			},
			MD5:        "2yN-w0xN9lsYzx1lbtZM-Q",
			TheDate:    "01/27 00:43",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32ma6234709\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO marvel\x1b[m                                      01/27 00:43\r"),
			CommentID:  "FQ1pzU6AyAA:2yN-w0xN9lsYzx1lbtZM-Q",
			CreateTime: 1516984980000000000,
			SortTime:   1516984980000000000,
		},
		{ // 6
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "KMU",
						Big5:   []byte("KMU"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("KMU"),
					},
				},
			},
			MD5:        "7zFE3bNz01gCEmX-hg4TxA",
			TheDate:    "01/28 12:10",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mgghost1002\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO KMU\x1b[m                                       01/28 12:10\r"),
			CommentID:  "FQ3d3ydE8AA:7zFE3bNz01gCEmX-hg4TxA",
			CreateTime: 1517112600000000000,
			SortTime:   1517112600000000000,
		},
		{ // 7
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Kaohsiung",
						Big5:   []byte("Kaohsiung"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Kaohsiung"),
					},
				},
			},
			MD5:        "AegTTCs475pm0Zx5M4ErTA",
			TheDate:    "01/28 12:11",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mgghost1002\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Kaohsiung\x1b[m                                 01/28 12:11\r"),
			CommentID:  "FQ3d7R-MSAA:AegTTCs475pm0Zx5M4ErTA",
			CreateTime: 1517112660000000000,
			SortTime:   1517112660000000000,
		},
		{ // 8
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("fennyccc"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "讚",
						Big5:   []byte("\xc6g                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc6g                                    "),
					},
				},
			},
			MD5:        "a8ZbQMrVxHxbUMd4IvtPcA",
			IP:         "115.82.149.161",
			TheDate:    "01/28 12:34",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mfennyccc\x1b[m\x1b[33m: \xc6g                                    \x1b[m 115.82.149.161 01/28 12:34\r"),
			CommentID:  "FQ3fLm31MAA:a8ZbQMrVxHxbUMd4IvtPcA",
			CreateTime: 1517114040000000000,
			SortTime:   1517114040000000000,
		},
		{ // 9
			BBoardID:  "test",
			ArticleID: "test19",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "某隱形看板",
						Big5:   []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO"),
					},
				},
			},
			MD5:        "unnUE1r9jqOU-NdpMHfWwA",
			TheDate:    "01/29 02:39",
			DBCS:       []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO\x1b[m                                         01/29 02:39\r"),
			CommentID:  "FQ4NSvFyqAA:unnUE1r9jqOU-NdpMHfWwA",
			CreateTime: 1517164740000000000,
			SortTime:   1517164740000000000,
		},
	}

	testTheRestComments19 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "賺 於是你借他錢 後來他真的賺了 你就可以",
						Big5:   []byte("\xc1\xc8 \xa9\xf3\xacO\xa7A\xad\xc9\xa5L\xbf\xfa \xab\xe1\xa8\xd3\xa5L\xafu\xaa\xba\xc1\xc8\xa4F \xa7A\xb4N\xa5i\xa5H "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc8 \xa9\xf3\xacO\xa7A\xad\xc9\xa5L\xbf\xfa \xab\xe1\xa8\xd3\xa5L\xafu\xaa\xba\xc1\xc8\xa4F \xa7A\xb4N\xa5i\xa5H "),
					},
				},
			},
			MD5:     "hxSTT_5bOL5jCmWV5vXs2g",
			IP:      "122.119.131.149",
			TheDate: "03/20 09:50",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xc1\xc8 \xa9\xf3\xacO\xa7A\xad\xc9\xa5L\xbf\xfa \xab\xe1\xa8\xd3\xa5L\xafu\xaa\xba\xc1\xc8\xa4F \xa7A\xb4N\xa5i\xa5H \x1b[m122.119.131.149 03/20 09:50\r"),
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Besorgen"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "叔叔我看過打麻將，好像也有募資行為。",
						Big5:   []byte("\xa8\xfb\xa8\xfb\xa7\xda\xac\xdd\xb9L\xa5\xb4\xb3\xc2\xb1N\xa1A\xa6n\xb9\xb3\xa4]\xa6\xb3\xb6\xd2\xb8\xea\xa6\xe6\xac\xb0\xa1C  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa8\xfb\xa8\xfb\xa7\xda\xac\xdd\xb9L\xa5\xb4\xb3\xc2\xb1N\xa1A\xa6n\xb9\xb3\xa4]\xa6\xb3\xb6\xd2\xb8\xea\xa6\xe6\xac\xb0\xa1C  "),
					},
				},
			},
			MD5:     "hx_riWhj_HuLg0UIlBSeuQ",
			IP:      "125.230.35.42",
			TheDate: "03/20 09:51",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mBesorgen\x1b[m\x1b[33m: \xa8\xfb\xa8\xfb\xa7\xda\xac\xdd\xb9L\xa5\xb4\xb3\xc2\xb1N\xa1A\xa6n\xb9\xb3\xa4]\xa6\xb3\xb6\xd2\xb8\xea\xa6\xe6\xac\xb0\xa1C  \x1b[m  125.230.35.42 03/20 09:51\r"),
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "得到本金 以及利息",
						Big5:   []byte("\xb1o\xa8\xec\xa5\xbb\xaa\xf7 \xa5H\xa4\xce\xa7Q\xae\xa7                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1o\xa8\xec\xa5\xbb\xaa\xf7 \xa5H\xa4\xce\xa7Q\xae\xa7                       "),
					},
				},
			},
			MD5:     "9_hUpOAIUE1HQcxhmr5YcQ",
			IP:      "122.119.131.149",
			TheDate: "03/20 09:51",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xb1o\xa8\xec\xa5\xbb\xaa\xf7 \xa5H\xa4\xce\xa7Q\xae\xa7                       \x1b[m122.119.131.149 03/20 09:51\r"),
		},
		{ // 3
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "拿回",
						Big5:   []byte("\xae\xb3\xa6^                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xae\xb3\xa6^                                    "),
					},
				},
			},
			MD5:     "Bkl1I9Dexr0692-6nBt7PQ",
			IP:      "122.119.131.149",
			TheDate: "03/20 09:52",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xae\xb3\xa6^                                    \x1b[m122.119.131.149 03/20 09:52\r"),
		},
		{ // 4
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("cc03233"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "重點 看誰做莊",
						Big5:   []byte("\xad\xab\xc2I \xac\xdd\xbd\xd6\xb0\xb5\xb2\xf8                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xad\xab\xc2I \xac\xdd\xbd\xd6\xb0\xb5\xb2\xf8                          "),
					},
				},
			},
			MD5:     "my7ntqlgU9SpujIvhhpV9g",
			IP:      "1.161.225.69",
			TheDate: "03/20 09:53",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mcc03233\x1b[m\x1b[33m: \xad\xab\xc2I \xac\xdd\xbd\xd6\xb0\xb5\xb2\xf8                          \x1b[m   1.161.225.69 03/20 09:53\r"),
		},
		{ // 5
			TheType: ptttype.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("RuleAllWorld"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "你先找出一副有成長性的麻將給我看",
						Big5:   []byte("\xa7A\xa5\xfd\xa7\xe4\xa5X\xa4@\xb0\xc6\xa6\xb3\xa6\xa8\xaa\xf8\xa9\xca\xaa\xba\xb3\xc2\xb1N\xb5\xb9\xa7\xda\xac\xdd  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7A\xa5\xfd\xa7\xe4\xa5X\xa4@\xb0\xc6\xa6\xb3\xa6\xa8\xaa\xf8\xa9\xca\xaa\xba\xb3\xc2\xb1N\xb5\xb9\xa7\xda\xac\xdd  "),
					},
				},
			},
			MD5:     "JHZlo-XfqbkQjxtH4jwk4g",
			IP:      "49.216.4.199",
			TheDate: "03/20 09:54",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mRuleAllWorld\x1b[m\x1b[33m: \xa7A\xa5\xfd\xa7\xe4\xa5X\xa4@\xb0\xc6\xa6\xb3\xa6\xa8\xaa\xf8\xa9\xca\xaa\xba\xb3\xc2\xb1N\xb5\xb9\xa7\xda\xac\xdd  \x1b[m   49.216.4.199 03/20 09:54\r"),
		},
		{ // 6
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("DinoZavolta"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "台灣人多的是把股市當賭場玩的",
						Big5:   []byte("\xa5x\xc6W\xa4H\xa6h\xaa\xba\xacO\xa7\xe2\xaa\xd1\xa5\xab\xb7\xed\xbd\xe4\xb3\xf5\xaa\xb1\xaa\xba       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5x\xc6W\xa4H\xa6h\xaa\xba\xacO\xa7\xe2\xaa\xd1\xa5\xab\xb7\xed\xbd\xe4\xb3\xf5\xaa\xb1\xaa\xba       "),
					},
				},
			},
			MD5:     "wSG9KCFXJmDKZqu4nH962Q",
			IP:      "112.105.78.63",
			TheDate: "03/20 10:27",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDinoZavolta\x1b[m\x1b[33m: \xa5x\xc6W\xa4H\xa6h\xaa\xba\xacO\xa7\xe2\xaa\xd1\xa5\xab\xb7\xed\xbd\xe4\xb3\xf5\xaa\xb1\xaa\xba       \x1b[m  112.105.78.63 03/20 10:27\r"),
		},
		{ // 7
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("oyui111"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "因為大家把股票當麻將玩",
						Big5:   []byte("\xa6]\xac\xb0\xa4j\xaea\xa7\xe2\xaa\xd1\xb2\xbc\xb7\xed\xb3\xc2\xb1N\xaa\xb1                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6]\xac\xb0\xa4j\xaea\xa7\xe2\xaa\xd1\xb2\xbc\xb7\xed\xb3\xc2\xb1N\xaa\xb1                 "),
					},
				},
			},
			MD5:     "Boqz1QHBlDpR3u9t2EUMag",
			IP:      "223.141.120.153",
			TheDate: "03/20 10:29",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33moyui111\x1b[m\x1b[33m: \xa6]\xac\xb0\xa4j\xaea\xa7\xe2\xaa\xd1\xb2\xbc\xb7\xed\xb3\xc2\xb1N\xaa\xb1                 \x1b[m223.141.120.153 03/20 10:29\r"),
		},
		{ // 8
			TheType: ptttype.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("jackeman"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "稅稅稅稅稅",
						Big5:   []byte("\xb5|\xb5|\xb5|\xb5|\xb5|                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb5|\xb5|\xb5|\xb5|\xb5|                            "),
					},
				},
			},
			MD5:     "YGqJ97m7yvJxtp5p7ylxPA",
			IP:      "61.223.69.223",
			TheDate: "03/20 10:33",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mjackeman\x1b[m\x1b[33m: \xb5|\xb5|\xb5|\xb5|\xb5|                            \x1b[m  61.223.69.223 03/20 10:33\r"),
		},
		{ // 9
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("tomxyz"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "症腐做莊的都合法喔 科科",
						Big5:   []byte("\xafg\xbbG\xb0\xb5\xb2\xf8\xaa\xba\xb3\xa3\xa6X\xaak\xb3\xe1 \xac\xec\xac\xec                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xafg\xbbG\xb0\xb5\xb2\xf8\xaa\xba\xb3\xa3\xa6X\xaak\xb3\xe1 \xac\xec\xac\xec                 "),
					},
				},
			},
			MD5:     "-_eTj8tTDmfX7iSxnqvURA",
			IP:      "42.72.59.33",
			TheDate: "03/20 10:57",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mtomxyz\x1b[m\x1b[33m: \xafg\xbbG\xb0\xb5\xb2\xf8\xaa\xba\xb3\xa3\xa6X\xaak\xb3\xe1 \xac\xec\xac\xec                 \x1b[m    42.72.59.33 03/20 10:57\r"),
		},
		{ // 10
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("bryanma"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "國家開的啊",
						Big5:   []byte("\xb0\xea\xaea\xb6}\xaa\xba\xb0\xda                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb0\xea\xaea\xb6}\xaa\xba\xb0\xda                             "),
					},
				},
			},
			MD5:     "-vyYfu3vDRxZT2HYvjERFw",
			IP:      "223.136.131.5",
			TheDate: "03/21 19:36",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mbryanma\x1b[m\x1b[33m: \xb0\xea\xaea\xb6}\xaa\xba\xb0\xda                             \x1b[m  223.136.131.5 03/21 19:36\r"),
		},
	}

	testFullTheRestComments19 = []*schema.Comment{
		{ // 0
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("jasome"),
			CommentID:  "FQ4NkMrXYAA:UmFoFJMkme_eoCJYAGGhOQ",
			CreateTime: 1517165040000000000,
			SortTime:   1517165040000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Lifeismoney",
						Big5:   []byte("Lifeismoney"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Lifeismoney"),
					},
				},
			},
			MD5:     "UmFoFJMkme_eoCJYAGGhOQ",
			TheDate: "01/29 02:44",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Lifeismoney\x1b[m                                   01/29 02:44\r"),
		},
		{ // 1
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("jasome"),
			CommentID:  "FQ4NurOtaAA:uGjATTO2E-igp_hil8pJew",
			CreateTime: 1517165220000000000,
			SortTime:   1517165220000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "MenTalk",
						Big5:   []byte("MenTalk"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("MenTalk"),
					},
				},
			},
			MD5:     "uGjATTO2E-igp_hil8pJew",
			TheDate: "01/29 02:47",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO MenTalk\x1b[m                                       01/29 02:47\r"),
		},
		{ // 2
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("jasome"),
			CommentID:  "FQ4ODoVZeAA:i_afAwPRZ-1POaOm_-R8iw",
			CreateTime: 1517165580000000000,
			SortTime:   1517165580000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "SENIORHIGH",
						Big5:   []byte("SENIORHIGH"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("SENIORHIGH"),
					},
				},
			},
			MD5:     "i_afAwPRZ-1POaOm_-R8iw",
			TheDate: "01/29 02:53",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO SENIORHIGH\x1b[m                                    01/29 02:53\r"),
		},
		{ // 3
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("jasome"),
			CommentID:  "FQ4OVF6-MAA:Or9MD3IqSpASZwdRwsLUqw",
			CreateTime: 1517165880000000000,
			SortTime:   1517165880000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "某隱形看板",
						Big5:   []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO"),
					},
				},
			},
			MD5:     "Or9MD3IqSpASZwdRwsLUqw",
			TheDate: "01/29 02:58",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO\x1b[m                                         01/29 02:58\r"),
		},
		{ // 4
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("frojet"),
			CommentID:  "FQ4SehPsUAA:4eCrt82X642AQp8vTDPc4Q",
			CreateTime: 1517170440000000000,
			SortTime:   1517170440000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "ONE_PIECE",
						Big5:   []byte("ONE_PIECE"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("ONE_PIECE"),
					},
				},
			},
			MD5:     "4eCrt82X642AQp8vTDPc4Q",
			TheDate: "01/29 04:14",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mfrojet\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO ONE_PIECE\x1b[m                                     01/29 04:14\r"),
		},
		{ // 5
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("a6234709"),
			CommentID:  "FQ4oQAseeAA:DQi4Af0sToGUHHHnb6k7UQ",
			CreateTime: 1517194380000000000,
			SortTime:   1517194380000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "movie",
						Big5:   []byte("movie"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("movie"),
					},
				},
			},
			MD5:     "DQi4Af0sToGUHHHnb6k7UQ",
			TheDate: "01/29 10:53",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32ma6234709\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO movie\x1b[m                                       01/29 10:53\r"),
		},
		{ // 6
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("jasome"),
			CommentID:  "FQ48qMNZCAA:WYQQxndOvwAgE4OYt2aTVg",
			CreateTime: 1517216820000000000,
			SortTime:   1517216820000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Gossiping",
						Big5:   []byte("Gossiping"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Gossiping"),
					},
				},
			},
			MD5:     "WYQQxndOvwAgE4OYt2aTVg",
			TheDate: "01/29 17:07",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Gossiping\x1b[m                                     01/29 17:07\r"),
		},
		{ // 7
			BBoardID:   "test",
			ArticleID:  "test19",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			Owner:      bbs.UUserID("gogin"),
			CommentID:  "FQ5wd0MfnAA:OwyLWiN4frHpPK6LGa-eaA",
			CreateTime: 1517273782000000000,
			SortTime:   1517273782000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "Wanted",
						Big5:   []byte("Wanted"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("Wanted"),
					},
				},
			},
			MD5:     "OwyLWiN4frHpPK6LGa-eaA",
			TheDate: "01/30 08:56",
			DBCS:    []byte("\xa1\xb0 \x1b[1;32mgogin\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Wanted\x1b[m                                         01/30 08:56\r"),
		},
	}
}
