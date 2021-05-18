package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename16            = "temp11"
	testContentAll16          []byte
	testContent16             []byte
	testSignature16           []byte
	testComment16             []byte
	testFirstCommentsDBCS16   []byte
	testTheRestCommentsDBCS16 []byte
	testContent16Big5         [][]*types.Rune
	testContent16Utf8         [][]*types.Rune

	testFirstComments16 []*schema.Comment
)

func initTest16() {
	testContentAll16, testContent16, testSignature16, testComment16, testFirstCommentsDBCS16, testTheRestCommentsDBCS16 = loadTest(testFilename16)

	testContent16Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: mapinkie (\xb6\xb2) \xac\xdd\xaaO: GreysAnatomy"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: mapinkie (\xb6\xb2) \xac\xdd\xaaO: GreysAnatomy\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] S16E16 (\xb9p\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] S16E16 (\xb9p\xa1^\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Mar  6 16:49:03 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Mar  6 16:49:03 2020\r"),
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

				Big5:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xacO\xa7\xda\xad^\xbby\xaet\xb2z\xb8\xd1\xbf\xf9\xbb~\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xacO\xa7\xda\xad^\xbby\xaet\xb2z\xb8\xd1\xbf\xf9\xbb~\xa1C\r"),
			},
		},
		{ //5
			{

				Big5:   []byte("\xb3\xcc\xab\xe1\xc1\xd9\xacO\xa8\xba\xa5y\xa1G\xa4j\xb6\xfd\xa7\xda\xab\xeb\xa7A\xa1I\xa1I\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3\xcc\xab\xe1\xc1\xd9\xacO\xa8\xba\xa5y\xa1G\xa4j\xb6\xfd\xa7\xda\xab\xeb\xa7A\xa1I\xa1I\xa1I\r"),
			},
		},
		{ //6
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //7
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("P.S."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("P.S.\r"),
			},
		},
		{ //9
			{

				Big5:   []byte("\xaa\xa6\xa4\xe5\xaa\xba\xae\xc9\xad\xd4\xb5o\xb2{\xaa\xa9\xa4\xcd\xad\xcc\xa6n\xb9\xb3\xa4@\xb6}\xa9l\xb3\xa3\xab\xdc\xb0Q\xb9\xbdLexie\xa9MApril\xa1A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaa\xa6\xa4\xe5\xaa\xba\xae\xc9\xad\xd4\xb5o\xb2{\xaa\xa9\xa4\xcd\xad\xcc\xa6n\xb9\xb3\xa4@\xb6}\xa9l\xb3\xa3\xab\xdc\xb0Q\xb9\xbdLexie\xa9MApril\xa1A\r"),
			},
		},
		{ //10
			{

				Big5:   []byte("\xa6\xfd\xa8\xec\xab\xe1\xa8\xd3\xadn\xa8\xe2\xa6\xec\xa8\xa4\xa6\xe2\xa8\xab\xaa\xba\xae\xc9\xab\xe1\xa4j\xaea\xabo\xab\xdc\xa4\xa3\xb1\xcb\xb1o\xa6o\xad\xcc\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6\xfd\xa8\xec\xab\xe1\xa8\xd3\xadn\xa8\xe2\xa6\xec\xa8\xa4\xa6\xe2\xa8\xab\xaa\xba\xae\xc9\xab\xe1\xa4j\xaea\xabo\xab\xdc\xa4\xa3\xb1\xcb\xb1o\xa6o\xad\xcc\xa1C\r"),
			},
		},
		{ //11
			{

				Big5:   []byte("\xbd\xd0\xb0\xdd\xa7A\xad\xcc\xacO\xb1q\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb6}\xa9l\xb3\xdf\xc5w\xa6o\xad\xcc\xaa\xba\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbd\xd0\xb0\xdd\xa7A\xad\xcc\xacO\xb1q\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb6}\xa9l\xb3\xdf\xc5w\xa6o\xad\xcc\xaa\xba\xa1H\r"),
			},
		},
	}

	testContent16Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: mapinkie (雯) 看板: GreysAnatomy",
				Big5:   []byte("\xa7@\xaa\xcc: mapinkie (\xb6\xb2) \xac\xdd\xaaO: GreysAnatomy"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: mapinkie (\xb6\xb2) \xac\xdd\xaaO: GreysAnatomy\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: [心得] S16E16 (雷）",
				Big5:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] S16E16 (\xb9p\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xa4\xdf\xb1o] S16E16 (\xb9p\xa1^\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Fri Mar  6 16:49:03 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Mar  6 16:49:03 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Mar  6 16:49:03 2020\r"),
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
				Utf8:   "我真的希望是我英語差理解錯誤。",
				Big5:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xacO\xa7\xda\xad^\xbby\xaet\xb2z\xb8\xd1\xbf\xf9\xbb~\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xacO\xa7\xda\xad^\xbby\xaet\xb2z\xb8\xd1\xbf\xf9\xbb~\xa1C\r"),
			},
		},
		{ //5
			{
				Utf8:   "最後還是那句：大媽我恨你！！！",
				Big5:   []byte("\xb3\xcc\xab\xe1\xc1\xd9\xacO\xa8\xba\xa5y\xa1G\xa4j\xb6\xfd\xa7\xda\xab\xeb\xa7A\xa1I\xa1I\xa1I"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3\xcc\xab\xe1\xc1\xd9\xacO\xa8\xba\xa5y\xa1G\xa4j\xb6\xfd\xa7\xda\xab\xeb\xa7A\xa1I\xa1I\xa1I\r"),
			},
		},
		{ //6
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //7
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //8
			{

				Utf8:   "P.S.",
				Big5:   []byte("P.S."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("P.S.\r"),
			},
		},
		{ //9
			{
				Utf8:   "爬文的時候發現版友們好像一開始都很討厭Lexie和April，",
				Big5:   []byte("\xaa\xa6\xa4\xe5\xaa\xba\xae\xc9\xad\xd4\xb5o\xb2{\xaa\xa9\xa4\xcd\xad\xcc\xa6n\xb9\xb3\xa4@\xb6}\xa9l\xb3\xa3\xab\xdc\xb0Q\xb9\xbdLexie\xa9MApril\xa1A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaa\xa6\xa4\xe5\xaa\xba\xae\xc9\xad\xd4\xb5o\xb2{\xaa\xa9\xa4\xcd\xad\xcc\xa6n\xb9\xb3\xa4@\xb6}\xa9l\xb3\xa3\xab\xdc\xb0Q\xb9\xbdLexie\xa9MApril\xa1A\r"),
			},
		},
		{ //10
			{
				Utf8:   "但到後來要兩位角色走的時後大家卻很不捨得她們。",
				Big5:   []byte("\xa6\xfd\xa8\xec\xab\xe1\xa8\xd3\xadn\xa8\xe2\xa6\xec\xa8\xa4\xa6\xe2\xa8\xab\xaa\xba\xae\xc9\xab\xe1\xa4j\xaea\xabo\xab\xdc\xa4\xa3\xb1\xcb\xb1o\xa6o\xad\xcc\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6\xfd\xa8\xec\xab\xe1\xa8\xd3\xadn\xa8\xe2\xa6\xec\xa8\xa4\xa6\xe2\xa8\xab\xaa\xba\xae\xc9\xab\xe1\xa4j\xaea\xabo\xab\xdc\xa4\xa3\xb1\xcb\xb1o\xa6o\xad\xcc\xa1C\r"),
			},
		},
		{ //11
			{
				Utf8:   "請問你們是從什麼時候開始喜歡她們的？",
				Big5:   []byte("\xbd\xd0\xb0\xdd\xa7A\xad\xcc\xacO\xb1q\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb6}\xa9l\xb3\xdf\xc5w\xa6o\xad\xcc\xaa\xba\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbd\xd0\xb0\xdd\xa7A\xad\xcc\xacO\xb1q\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xb6}\xa9l\xb3\xdf\xc5w\xa6o\xad\xcc\xaa\xba\xa1H\r"),
			},
		},
	}

}
