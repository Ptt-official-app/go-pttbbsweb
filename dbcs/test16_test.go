package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
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

	testFirstComments16     []*schema.Comment
	testFullFirstComments16 []*schema.Comment
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

	testFirstComments16 = []*schema.Comment{
		{ //0
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("mapinkie"),
			CreateTime: 1583511858000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "我最初看到是3.8分，然後現在好像升到了4.5。",
						Big5:   []byte("\xa7\xda\xb3\xcc\xaa\xec\xac\xdd\xa8\xec\xacO3.8\xa4\xc0\xa1A\xb5M\xab\xe1\xb2{\xa6b\xa6n\xb9\xb3\xa4\xc9\xa8\xec\xa4F4.5\xa1C"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xda\xb3\xcc\xaa\xec\xac\xdd\xa8\xec\xacO3.8\xa4\xc0\xa1A\xb5M\xab\xe1\xb2{\xa6b\xa6n\xb9\xb3\xa4\xc9\xa8\xec\xa4F4.5\xa1C\r"),
					},
				},
				{
					{
						Utf8:   "看來大家也跟我一樣看了很氣，我猜空難那集也沒那麼誇張。",
						Big5:   []byte("\xac\xdd\xa8\xd3\xa4j\xaea\xa4]\xb8\xf2\xa7\xda\xa4@\xbc\xcb\xac\xdd\xa4F\xab\xdc\xae\xf0\xa1A\xa7\xda\xb2q\xaa\xc5\xc3\xf8\xa8\xba\xb6\xb0\xa4]\xa8S\xa8\xba\xbb\xf2\xb8\xd8\xb1i\xa1C"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xac\xdd\xa8\xd3\xa4j\xaea\xa4]\xb8\xf2\xa7\xda\xa4@\xbc\xcb\xac\xdd\xa4F\xab\xdc\xae\xf0\xa1A\xa7\xda\xb2q\xaa\xc5\xc3\xf8\xa8\xba\xb6\xb0\xa4]\xa8S\xa8\xba\xbb\xf2\xb8\xd8\xb1i\xa1C\r"),
					},
				},
				{
					{
						Utf8:   "BTW我認真覺得絕對是大媽仍在恨演Izzie的演員，順便用這招來噁心她的。",
						Big5:   []byte("BTW\xa7\xda\xbb{\xafu\xc4\xb1\xb1o\xb5\xb4\xb9\xef\xacO\xa4j\xb6\xfd\xa4\xb4\xa6b\xab\xeb\xbatIzzie\xaa\xba\xbat\xad\xfb\xa1A\xb6\xb6\xabK\xa5\xce\xb3o\xa9\xdb\xa8\xd3\xe4\xfa\xa4\xdf\xa6o\xaa\xba\xa1C"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("BTW\xa7\xda\xbb{\xafu\xc4\xb1\xb1o\xb5\xb4\xb9\xef\xacO\xa4j\xb6\xfd\xa4\xb4\xa6b\xab\xeb\xbatIzzie\xaa\xba\xbat\xad\xfb\xa1A\xb6\xb6\xabK\xa5\xce\xb3o\xa9\xdb\xa8\xd3\xe4\xfa\xa4\xdf\xa6o\xaa\xba\xa1C"),
					},
				},
			},
			MD5:        "MlwwZY4bzjaIUZAotWmOPw",
			TheDate:    "03/07/2020 00:24:18",
			IP:         "58.152.169.221",
			Host:       "香港",
			EditNanoTS: 1583511858000000000,
			DBCS:       []byte("\r\n\xa7\xda\xb3\xcc\xaa\xec\xac\xdd\xa8\xec\xacO3.8\xa4\xc0\xa1A\xb5M\xab\xe1\xb2{\xa6b\xa6n\xb9\xb3\xa4\xc9\xa8\xec\xa4F4.5\xa1C\r\n\xac\xdd\xa8\xd3\xa4j\xaea\xa4]\xb8\xf2\xa7\xda\xa4@\xbc\xcb\xac\xdd\xa4F\xab\xdc\xae\xf0\xa1A\xa7\xda\xb2q\xaa\xc5\xc3\xf8\xa8\xba\xb6\xb0\xa4]\xa8S\xa8\xba\xbb\xf2\xb8\xd8\xb1i\xa1C\r\nBTW\xa7\xda\xbb{\xafu\xc4\xb1\xb1o\xb5\xb4\xb9\xef\xacO\xa4j\xb6\xfd\xa4\xb4\xa6b\xab\xeb\xbatIzzie\xaa\xba\xbat\xad\xfb\xa1A\xb6\xb6\xabK\xa5\xce\xb3o\xa9\xdb\xa8\xd3\xe4\xfa\xa4\xdf\xa6o\xaa\xba\xa1C\r"),
		},
		{ //1
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("meowmeow516"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這個離場我不能接受QQ",
						Big5:   []byte("\xb3o\xad\xd3\xc2\xf7\xb3\xf5\xa7\xda\xa4\xa3\xaf\xe0\xb1\xb5\xa8\xfcQQ                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xad\xd3\xc2\xf7\xb3\xf5\xa7\xda\xa4\xa3\xaf\xe0\xb1\xb5\xa8\xfcQQ                              "),
					},
				},
			},
			MD5:     "Ircl0IPEos24qAKyPbg6ZQ",
			TheDate: "03/07 00:20",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mmeowmeow516\x1b[m\x1b[33m: \xb3o\xad\xd3\xc2\xf7\xb3\xf5\xa7\xda\xa4\xa3\xaf\xe0\xb1\xb5\xa8\xfcQQ                              \x1b[m 03/07 00:20\r"),
		},
		{ //2
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "其實可以就讓他像剛消失的那幾集活在演員台詞中就好了 Ale",
						Big5:   []byte("\xa8\xe4\xb9\xea\xa5i\xa5H\xb4N\xc5\xfd\xa5L\xb9\xb3\xad\xe8\xae\xf8\xa5\xa2\xaa\xba\xa8\xba\xb4X\xb6\xb0\xac\xa1\xa6b\xbat\xad\xfb\xa5x\xb5\xfc\xa4\xa4\xb4N\xa6n\xa4F Ale "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa8\xe4\xb9\xea\xa5i\xa5H\xb4N\xc5\xfd\xa5L\xb9\xb3\xad\xe8\xae\xf8\xa5\xa2\xaa\xba\xa8\xba\xb4X\xb6\xb0\xac\xa1\xa6b\xbat\xad\xfb\xa5x\xb5\xfc\xa4\xa4\xb4N\xa6n\xa4F Ale "),
					},
				},
			},
			MD5:     "-hljJY3kGIN-cQtsNwf5Cw",
			TheDate: "03/07 00:22",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33msounan\x1b[m\x1b[33m: \xa8\xe4\xb9\xea\xa5i\xa5H\xb4N\xc5\xfd\xa5L\xb9\xb3\xad\xe8\xae\xf8\xa5\xa2\xaa\xba\xa8\xba\xb4X\xb6\xb0\xac\xa1\xa6b\xbat\xad\xfb\xa5x\xb5\xfc\xa4\xa4\xb4N\xa6n\xa4F Ale \x1b[m 03/07 00:22\r"),
		},
		{ //3
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "x去找他媽媽了 他去了哪裡之類的 為什麼沒事跟Izzie扯上關",
						Big5:   []byte("x\xa5h\xa7\xe4\xa5L\xb6\xfd\xb6\xfd\xa4F \xa5L\xa5h\xa4F\xad\xfe\xb8\xcc\xa4\xa7\xc3\xfe\xaa\xba \xac\xb0\xa4\xb0\xbb\xf2\xa8S\xa8\xc6\xb8\xf2Izzie\xa7\xe8\xa4W\xc3\xf6 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("x\xa5h\xa7\xe4\xa5L\xb6\xfd\xb6\xfd\xa4F \xa5L\xa5h\xa4F\xad\xfe\xb8\xcc\xa4\xa7\xc3\xfe\xaa\xba \xac\xb0\xa4\xb0\xbb\xf2\xa8S\xa8\xc6\xb8\xf2Izzie\xa7\xe8\xa4W\xc3\xf6 "),
					},
				},
			},
			MD5:     "SRgpd8HvsFuCyCHPbBmaZw",
			TheDate: "03/07 00:22",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: x\xa5h\xa7\xe4\xa5L\xb6\xfd\xb6\xfd\xa4F \xa5L\xa5h\xa4F\xad\xfe\xb8\xcc\xa4\xa7\xc3\xfe\xaa\xba \xac\xb0\xa4\xb0\xbb\xf2\xa8S\xa8\xc6\xb8\xf2Izzie\xa7\xe8\xa4W\xc3\xf6 \x1b[m 03/07 00:22\r"),
		},
		{ //4
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "係= =",
						Big5:   []byte("\xabY= =                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xabY= =                                                  "),
					},
				},
			},
			MD5:     "U3Ik65lHMolfna0kbPCVBQ",
			TheDate: "03/07 00:22",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xabY= =                                                  \x1b[m 03/07 00:22\r"),
		},
		{ //5
			TheType: types.COMMENT_TYPE_EDIT,
			Owner:   bbs.UUserID("mapinkie"),

			MD5:                "iNGAzCTcEvC5J-jBrqW9bQ",
			TheDate:            "03/07/2020 00:24:18",
			IP:                 "58.152.169.221",
			Host:               "香港",
			DBCS:               []byte("\xa1\xb0 \xbds\xbf\xe8: mapinkie (58.152.169.221 \xad\xbb\xb4\xe4), 03/07/2020 00:24:18\r"),
			CreateTime:         1583511858000000000,
			InferredCreateTime: 1583511858000000000,
			SortTime:           1583511858000000000,
			CommentID:          "FfnDpSG1dAA:iNGAzCTcEvC5J-jBrqW9bQ",
		},
		{ //6
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("artandcounse"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噁心她幹嘛遷怒我們的Alex！怒",
						Big5:   []byte("\xe4\xfa\xa4\xdf\xa6o\xb7F\xb9\xc0\xbeE\xab\xe3\xa7\xda\xad\xcc\xaa\xbaAlex\xa1I\xab\xe3                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xe4\xfa\xa4\xdf\xa6o\xb7F\xb9\xc0\xbeE\xab\xe3\xa7\xda\xad\xcc\xaa\xbaAlex\xa1I\xab\xe3                     "),
					},
				},
			},
			MD5:     "otKsr88jcjEFpdW-t4DrqQ",
			TheDate: "03/07 00:24",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33martandcounse\x1b[m\x1b[33m: \xe4\xfa\xa4\xdf\xa6o\xb7F\xb9\xc0\xbeE\xab\xe3\xa7\xda\xad\xcc\xaa\xbaAlex\xa1I\xab\xe3                     \x1b[m 03/07 00:24\r"),
		},
		{ //7
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這幾集被寫爛的還有Andrew 要讓Meredith有新對象就不能讓",
						Big5:   []byte("\xb3o\xb4X\xb6\xb0\xb3Q\xbcg\xc4\xea\xaa\xba\xc1\xd9\xa6\xb3Andrew \xadn\xc5\xfdMeredith\xa6\xb3\xb7s\xb9\xef\xb6H\xb4N\xa4\xa3\xaf\xe0\xc5\xfd  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xb4X\xb6\xb0\xb3Q\xbcg\xc4\xea\xaa\xba\xc1\xd9\xa6\xb3Andrew \xadn\xc5\xfdMeredith\xa6\xb3\xb7s\xb9\xef\xb6H\xb4N\xa4\xa3\xaf\xe0\xc5\xfd  "),
					},
				},
			},
			MD5:     "hBi-IucpCpvhokjuT5tI4g",
			TheDate: "03/07 00:27",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xb3o\xb4X\xb6\xb0\xb3Q\xbcg\xc4\xea\xaa\xba\xc1\xd9\xa6\xb3Andrew \xadn\xc5\xfdMeredith\xa6\xb3\xb7s\xb9\xef\xb6H\xb4N\xa4\xa3\xaf\xe0\xc5\xfd  \x1b[m 03/07 00:27\r"),
		},
		{ //8
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "她跟Andrew好聚好散嗎 非要把他寫成這樣 雖然整體來說我反",
						Big5:   []byte("\xa6o\xb8\xf2Andrew\xa6n\xbbE\xa6n\xb4\xb2\xb6\xdc \xabD\xadn\xa7\xe2\xa5L\xbcg\xa6\xa8\xb3o\xbc\xcb \xc1\xf6\xb5M\xbe\xe3\xc5\xe9\xa8\xd3\xbb\xa1\xa7\xda\xa4\xcf "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6o\xb8\xf2Andrew\xa6n\xbbE\xa6n\xb4\xb2\xb6\xdc \xabD\xadn\xa7\xe2\xa5L\xbcg\xa6\xa8\xb3o\xbc\xcb \xc1\xf6\xb5M\xbe\xe3\xc5\xe9\xa8\xd3\xbb\xa1\xa7\xda\xa4\xcf "),
					},
				},
			},
			MD5:     "Why7Wi-d6DBBM0qpl7qw0Q",
			TheDate: "03/07 00:27",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xa6o\xb8\xf2Andrew\xa6n\xbbE\xa6n\xb4\xb2\xb6\xdc \xabD\xadn\xa7\xe2\xa5L\xbcg\xa6\xa8\xb3o\xbc\xcb \xc1\xf6\xb5M\xbe\xe3\xc5\xe9\xa8\xd3\xbb\xa1\xa7\xda\xa4\xcf \x1b[m 03/07 00:27\r"),
		},
		{ //9
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "而覺得Andrew跟Maggie配對的化學反應是最好的",
						Big5:   []byte("\xa6\xd3\xc4\xb1\xb1oAndrew\xb8\xf2Maggie\xb0t\xb9\xef\xaa\xba\xa4\xc6\xbe\xc7\xa4\xcf\xc0\xb3\xacO\xb3\xcc\xa6n\xaa\xba             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6\xd3\xc4\xb1\xb1oAndrew\xb8\xf2Maggie\xb0t\xb9\xef\xaa\xba\xa4\xc6\xbe\xc7\xa4\xcf\xc0\xb3\xacO\xb3\xcc\xa6n\xaa\xba             "),
					},
				},
			},
			MD5:     "KFYDlvlUcGYwnxMH-3YMxA",
			TheDate: "03/07 00:27",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xa6\xd3\xc4\xb1\xb1oAndrew\xb8\xf2Maggie\xb0t\xb9\xef\xaa\xba\xa4\xc6\xbe\xc7\xa4\xcf\xc0\xb3\xacO\xb3\xcc\xa6n\xaa\xba             \x1b[m 03/07 00:27\r"),
		},
		{ //10
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("mapinkie"),
			CreateTime: 1583512190000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "我真的希望不要再讓Mere搞對象了，就讓她普通的當醫生和做單親媽媽吧！",
						Big5:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xa4\xa3\xadn\xa6A\xc5\xfdMere\xb7d\xb9\xef\xb6H\xa4F\xa1A\xb4N\xc5\xfd\xa6o\xb4\xb6\xb3q\xaa\xba\xb7\xed\xc2\xe5\xa5\xcd\xa9M\xb0\xb5\xb3\xe6\xbf\xcb\xb6\xfd\xb6\xfd\xa7a\xa1I"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xa4\xa3\xadn\xa6A\xc5\xfdMere\xb7d\xb9\xef\xb6H\xa4F\xa1A\xb4N\xc5\xfd\xa6o\xb4\xb6\xb3q\xaa\xba\xb7\xed\xc2\xe5\xa5\xcd\xa9M\xb0\xb5\xb3\xe6\xbf\xcb\xb6\xfd\xb6\xfd\xa7a\xa1I"),
					},
				},
			},
			MD5:        "7tkvXPK6iHO5IzFBfJCOwA",
			TheDate:    "03/07/2020 00:29:50",
			IP:         "58.152.169.221",
			EditNanoTS: 1583512190000000000,
			Host:       "香港",
			DBCS:       []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xa4\xa3\xadn\xa6A\xc5\xfdMere\xb7d\xb9\xef\xb6H\xa4F\xa1A\xb4N\xc5\xfd\xa6o\xb4\xb6\xb3q\xaa\xba\xb7\xed\xc2\xe5\xa5\xcd\xa9M\xb0\xb5\xb3\xe6\xbf\xcb\xb6\xfd\xb6\xfd\xa7a\xa1I\r"),
		},
	}

	testFullFirstComments16 = []*schema.Comment{
		{ //0
			BBoardID:   "test",
			ArticleID:  "test16",
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("mapinkie"),
			CreateTime: 1583511858000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "我最初看到是3.8分，然後現在好像升到了4.5。",
						Big5:   []byte("\xa7\xda\xb3\xcc\xaa\xec\xac\xdd\xa8\xec\xacO3.8\xa4\xc0\xa1A\xb5M\xab\xe1\xb2{\xa6b\xa6n\xb9\xb3\xa4\xc9\xa8\xec\xa4F4.5\xa1C"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xda\xb3\xcc\xaa\xec\xac\xdd\xa8\xec\xacO3.8\xa4\xc0\xa1A\xb5M\xab\xe1\xb2{\xa6b\xa6n\xb9\xb3\xa4\xc9\xa8\xec\xa4F4.5\xa1C\r"),
					},
				},
				{
					{
						Utf8:   "看來大家也跟我一樣看了很氣，我猜空難那集也沒那麼誇張。",
						Big5:   []byte("\xac\xdd\xa8\xd3\xa4j\xaea\xa4]\xb8\xf2\xa7\xda\xa4@\xbc\xcb\xac\xdd\xa4F\xab\xdc\xae\xf0\xa1A\xa7\xda\xb2q\xaa\xc5\xc3\xf8\xa8\xba\xb6\xb0\xa4]\xa8S\xa8\xba\xbb\xf2\xb8\xd8\xb1i\xa1C"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xac\xdd\xa8\xd3\xa4j\xaea\xa4]\xb8\xf2\xa7\xda\xa4@\xbc\xcb\xac\xdd\xa4F\xab\xdc\xae\xf0\xa1A\xa7\xda\xb2q\xaa\xc5\xc3\xf8\xa8\xba\xb6\xb0\xa4]\xa8S\xa8\xba\xbb\xf2\xb8\xd8\xb1i\xa1C\r"),
					},
				},
				{
					{
						Utf8:   "BTW我認真覺得絕對是大媽仍在恨演Izzie的演員，順便用這招來噁心她的。",
						Big5:   []byte("BTW\xa7\xda\xbb{\xafu\xc4\xb1\xb1o\xb5\xb4\xb9\xef\xacO\xa4j\xb6\xfd\xa4\xb4\xa6b\xab\xeb\xbatIzzie\xaa\xba\xbat\xad\xfb\xa1A\xb6\xb6\xabK\xa5\xce\xb3o\xa9\xdb\xa8\xd3\xe4\xfa\xa4\xdf\xa6o\xaa\xba\xa1C"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("BTW\xa7\xda\xbb{\xafu\xc4\xb1\xb1o\xb5\xb4\xb9\xef\xacO\xa4j\xb6\xfd\xa4\xb4\xa6b\xab\xeb\xbatIzzie\xaa\xba\xbat\xad\xfb\xa1A\xb6\xb6\xabK\xa5\xce\xb3o\xa9\xdb\xa8\xd3\xe4\xfa\xa4\xdf\xa6o\xaa\xba\xa1C"),
					},
				},
			},
			MD5:        "MlwwZY4bzjaIUZAotWmOPw",
			TheDate:    "03/07/2020 00:24:18",
			IP:         "58.152.169.221",
			Host:       "香港",
			EditNanoTS: 1583511858000000000,
			DBCS:       []byte("\r\n\xa7\xda\xb3\xcc\xaa\xec\xac\xdd\xa8\xec\xacO3.8\xa4\xc0\xa1A\xb5M\xab\xe1\xb2{\xa6b\xa6n\xb9\xb3\xa4\xc9\xa8\xec\xa4F4.5\xa1C\r\n\xac\xdd\xa8\xd3\xa4j\xaea\xa4]\xb8\xf2\xa7\xda\xa4@\xbc\xcb\xac\xdd\xa4F\xab\xdc\xae\xf0\xa1A\xa7\xda\xb2q\xaa\xc5\xc3\xf8\xa8\xba\xb6\xb0\xa4]\xa8S\xa8\xba\xbb\xf2\xb8\xd8\xb1i\xa1C\r\nBTW\xa7\xda\xbb{\xafu\xc4\xb1\xb1o\xb5\xb4\xb9\xef\xacO\xa4j\xb6\xfd\xa4\xb4\xa6b\xab\xeb\xbatIzzie\xaa\xba\xbat\xad\xfb\xa1A\xb6\xb6\xabK\xa5\xce\xb3o\xa9\xdb\xa8\xd3\xe4\xfa\xa4\xdf\xa6o\xaa\xba\xa1C\r"),
			CommentID:  "FfmqzVzXvKA:MlwwZY4bzjaIUZAotWmOPw",
			SortTime:   1583484543000100000,
		},
		{ //1
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("meowmeow516"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這個離場我不能接受QQ",
						Big5:   []byte("\xb3o\xad\xd3\xc2\xf7\xb3\xf5\xa7\xda\xa4\xa3\xaf\xe0\xb1\xb5\xa8\xfcQQ                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xad\xd3\xc2\xf7\xb3\xf5\xa7\xda\xa4\xa3\xaf\xe0\xb1\xb5\xa8\xfcQQ                              "),
					},
				},
			},
			MD5:        "Ircl0IPEos24qAKyPbg6ZQ",
			TheDate:    "03/07 00:20",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mmeowmeow516\x1b[m\x1b[33m: \xb3o\xad\xd3\xc2\xf7\xb3\xf5\xa7\xda\xa4\xa3\xaf\xe0\xb1\xb5\xa8\xfcQQ                              \x1b[m 03/07 00:20\r"),
			CommentID:  "FfnDaQ-14AA:Ircl0IPEos24qAKyPbg6ZQ",
			CreateTime: 1583511600000000000,
			SortTime:   1583511600000000000,
		},
		{ //2
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "其實可以就讓他像剛消失的那幾集活在演員台詞中就好了 Ale",
						Big5:   []byte("\xa8\xe4\xb9\xea\xa5i\xa5H\xb4N\xc5\xfd\xa5L\xb9\xb3\xad\xe8\xae\xf8\xa5\xa2\xaa\xba\xa8\xba\xb4X\xb6\xb0\xac\xa1\xa6b\xbat\xad\xfb\xa5x\xb5\xfc\xa4\xa4\xb4N\xa6n\xa4F Ale "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa8\xe4\xb9\xea\xa5i\xa5H\xb4N\xc5\xfd\xa5L\xb9\xb3\xad\xe8\xae\xf8\xa5\xa2\xaa\xba\xa8\xba\xb4X\xb6\xb0\xac\xa1\xa6b\xbat\xad\xfb\xa5x\xb5\xfc\xa4\xa4\xb4N\xa6n\xa4F Ale "),
					},
				},
			},
			MD5:        "-hljJY3kGIN-cQtsNwf5Cw",
			TheDate:    "03/07 00:22",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33msounan\x1b[m\x1b[33m: \xa8\xe4\xb9\xea\xa5i\xa5H\xb4N\xc5\xfd\xa5L\xb9\xb3\xad\xe8\xae\xf8\xa5\xa2\xaa\xba\xa8\xba\xb4X\xb6\xb0\xac\xa1\xa6b\xbat\xad\xfb\xa5x\xb5\xfc\xa4\xa4\xb4N\xa6n\xa4F Ale \x1b[m 03/07 00:22\r"),
			CommentID:  "FfnDhQBEkAA:-hljJY3kGIN-cQtsNwf5Cw",
			CreateTime: 1583511720000000000,
			SortTime:   1583511720000000000,
		},
		{ //3
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "x去找他媽媽了 他去了哪裡之類的 為什麼沒事跟Izzie扯上關",
						Big5:   []byte("x\xa5h\xa7\xe4\xa5L\xb6\xfd\xb6\xfd\xa4F \xa5L\xa5h\xa4F\xad\xfe\xb8\xcc\xa4\xa7\xc3\xfe\xaa\xba \xac\xb0\xa4\xb0\xbb\xf2\xa8S\xa8\xc6\xb8\xf2Izzie\xa7\xe8\xa4W\xc3\xf6 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("x\xa5h\xa7\xe4\xa5L\xb6\xfd\xb6\xfd\xa4F \xa5L\xa5h\xa4F\xad\xfe\xb8\xcc\xa4\xa7\xc3\xfe\xaa\xba \xac\xb0\xa4\xb0\xbb\xf2\xa8S\xa8\xc6\xb8\xf2Izzie\xa7\xe8\xa4W\xc3\xf6 "),
					},
				},
			},
			MD5:        "SRgpd8HvsFuCyCHPbBmaZw",
			TheDate:    "03/07 00:22",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: x\xa5h\xa7\xe4\xa5L\xb6\xfd\xb6\xfd\xa4F \xa5L\xa5h\xa4F\xad\xfe\xb8\xcc\xa4\xa7\xc3\xfe\xaa\xba \xac\xb0\xa4\xb0\xbb\xf2\xa8S\xa8\xc6\xb8\xf2Izzie\xa7\xe8\xa4W\xc3\xf6 \x1b[m 03/07 00:22\r"),
			CommentID:  "FfnDhQBT0kA:SRgpd8HvsFuCyCHPbBmaZw",
			CreateTime: 1583511720000000000,
			SortTime:   1583511720001000000,
		},
		{ //4
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "係= =",
						Big5:   []byte("\xabY= =                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xabY= =                                                  "),
					},
				},
			},
			MD5:        "U3Ik65lHMolfna0kbPCVBQ",
			TheDate:    "03/07 00:22",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xabY= =                                                  \x1b[m 03/07 00:22\r"),
			CommentID:  "FfnDhQBjFIA:U3Ik65lHMolfna0kbPCVBQ",
			CreateTime: 1583511720000000000,
			SortTime:   1583511720002000000,
		},
		{ //5
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_EDIT,
			Owner:     bbs.UUserID("mapinkie"),

			MD5:                "iNGAzCTcEvC5J-jBrqW9bQ",
			TheDate:            "03/07/2020 00:24:18",
			IP:                 "58.152.169.221",
			Host:               "香港",
			DBCS:               []byte("\xa1\xb0 \xbds\xbf\xe8: mapinkie (58.152.169.221 \xad\xbb\xb4\xe4), 03/07/2020 00:24:18\r"),
			CreateTime:         1583511858000000000,
			InferredCreateTime: 1583511858000000000,
			SortTime:           1583511858000000000,
			CommentID:          "FfnDpSG1dAA:iNGAzCTcEvC5J-jBrqW9bQ",
		},
		{ //6
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("artandcounse"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噁心她幹嘛遷怒我們的Alex！怒",
						Big5:   []byte("\xe4\xfa\xa4\xdf\xa6o\xb7F\xb9\xc0\xbeE\xab\xe3\xa7\xda\xad\xcc\xaa\xbaAlex\xa1I\xab\xe3                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xe4\xfa\xa4\xdf\xa6o\xb7F\xb9\xc0\xbeE\xab\xe3\xa7\xda\xad\xcc\xaa\xbaAlex\xa1I\xab\xe3                     "),
					},
				},
			},
			MD5:        "otKsr88jcjEFpdW-t4DrqQ",
			TheDate:    "03/07 00:24",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33martandcounse\x1b[m\x1b[33m: \xe4\xfa\xa4\xdf\xa6o\xb7F\xb9\xc0\xbeE\xab\xe3\xa7\xda\xad\xcc\xaa\xbaAlex\xa1I\xab\xe3                     \x1b[m 03/07 00:24\r"),
			CommentID:  "FfnDpSHEtkA:otKsr88jcjEFpdW-t4DrqQ",
			CreateTime: 1583511840000000000,
			SortTime:   1583511858001000000,
		},
		{ //7
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這幾集被寫爛的還有Andrew 要讓Meredith有新對象就不能讓",
						Big5:   []byte("\xb3o\xb4X\xb6\xb0\xb3Q\xbcg\xc4\xea\xaa\xba\xc1\xd9\xa6\xb3Andrew \xadn\xc5\xfdMeredith\xa6\xb3\xb7s\xb9\xef\xb6H\xb4N\xa4\xa3\xaf\xe0\xc5\xfd  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3o\xb4X\xb6\xb0\xb3Q\xbcg\xc4\xea\xaa\xba\xc1\xd9\xa6\xb3Andrew \xadn\xc5\xfdMeredith\xa6\xb3\xb7s\xb9\xef\xb6H\xb4N\xa4\xa3\xaf\xe0\xc5\xfd  "),
					},
				},
			},
			MD5:        "hBi-IucpCpvhokjuT5tI4g",
			TheDate:    "03/07 00:27",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xb3o\xb4X\xb6\xb0\xb3Q\xbcg\xc4\xea\xaa\xba\xc1\xd9\xa6\xb3Andrew \xadn\xc5\xfdMeredith\xa6\xb3\xb7s\xb9\xef\xb6H\xb4N\xa4\xa3\xaf\xe0\xc5\xfd  \x1b[m 03/07 00:27\r"),
			CommentID:  "FfnDytmpSAA:hBi-IucpCpvhokjuT5tI4g",
			CreateTime: 1583512020000000000,
			SortTime:   1583512020000000000,
		},
		{ //8
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "她跟Andrew好聚好散嗎 非要把他寫成這樣 雖然整體來說我反",
						Big5:   []byte("\xa6o\xb8\xf2Andrew\xa6n\xbbE\xa6n\xb4\xb2\xb6\xdc \xabD\xadn\xa7\xe2\xa5L\xbcg\xa6\xa8\xb3o\xbc\xcb \xc1\xf6\xb5M\xbe\xe3\xc5\xe9\xa8\xd3\xbb\xa1\xa7\xda\xa4\xcf "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6o\xb8\xf2Andrew\xa6n\xbbE\xa6n\xb4\xb2\xb6\xdc \xabD\xadn\xa7\xe2\xa5L\xbcg\xa6\xa8\xb3o\xbc\xcb \xc1\xf6\xb5M\xbe\xe3\xc5\xe9\xa8\xd3\xbb\xa1\xa7\xda\xa4\xcf "),
					},
				},
			},
			MD5:        "Why7Wi-d6DBBM0qpl7qw0Q",
			TheDate:    "03/07 00:27",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xa6o\xb8\xf2Andrew\xa6n\xbbE\xa6n\xb4\xb2\xb6\xdc \xabD\xadn\xa7\xe2\xa5L\xbcg\xa6\xa8\xb3o\xbc\xcb \xc1\xf6\xb5M\xbe\xe3\xc5\xe9\xa8\xd3\xbb\xa1\xa7\xda\xa4\xcf \x1b[m 03/07 00:27\r"),
			CommentID:  "FfnDytm4ikA:Why7Wi-d6DBBM0qpl7qw0Q",
			CreateTime: 1583512020000000000,
			SortTime:   1583512020001000000,
		},
		{ //9
			BBoardID:  "test",
			ArticleID: "test16",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("sounan"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "而覺得Andrew跟Maggie配對的化學反應是最好的",
						Big5:   []byte("\xa6\xd3\xc4\xb1\xb1oAndrew\xb8\xf2Maggie\xb0t\xb9\xef\xaa\xba\xa4\xc6\xbe\xc7\xa4\xcf\xc0\xb3\xacO\xb3\xcc\xa6n\xaa\xba             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6\xd3\xc4\xb1\xb1oAndrew\xb8\xf2Maggie\xb0t\xb9\xef\xaa\xba\xa4\xc6\xbe\xc7\xa4\xcf\xc0\xb3\xacO\xb3\xcc\xa6n\xaa\xba             "),
					},
				},
			},
			MD5:        "KFYDlvlUcGYwnxMH-3YMxA",
			TheDate:    "03/07 00:27",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33msounan\x1b[m\x1b[33m: \xa6\xd3\xc4\xb1\xb1oAndrew\xb8\xf2Maggie\xb0t\xb9\xef\xaa\xba\xa4\xc6\xbe\xc7\xa4\xcf\xc0\xb3\xacO\xb3\xcc\xa6n\xaa\xba             \x1b[m 03/07 00:27\r"),
			CommentID:  "FfnDytnHzIA:KFYDlvlUcGYwnxMH-3YMxA",
			CreateTime: 1583512020000000000,
			SortTime:   1583512020002000000,
		},
		{ //10
			BBoardID:   "test",
			ArticleID:  "test16",
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("mapinkie"),
			CreateTime: 1583512190000000000,
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "我真的希望不要再讓Mere搞對象了，就讓她普通的當醫生和做單親媽媽吧！",
						Big5:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xa4\xa3\xadn\xa6A\xc5\xfdMere\xb7d\xb9\xef\xb6H\xa4F\xa1A\xb4N\xc5\xfd\xa6o\xb4\xb6\xb3q\xaa\xba\xb7\xed\xc2\xe5\xa5\xcd\xa9M\xb0\xb5\xb3\xe6\xbf\xcb\xb6\xfd\xb6\xfd\xa7a\xa1I"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xa4\xa3\xadn\xa6A\xc5\xfdMere\xb7d\xb9\xef\xb6H\xa4F\xa1A\xb4N\xc5\xfd\xa6o\xb4\xb6\xb3q\xaa\xba\xb7\xed\xc2\xe5\xa5\xcd\xa9M\xb0\xb5\xb3\xe6\xbf\xcb\xb6\xfd\xb6\xfd\xa7a\xa1I"),
					},
				},
			},
			MD5:        "7tkvXPK6iHO5IzFBfJCOwA",
			TheDate:    "03/07/2020 00:29:50",
			IP:         "58.152.169.221",
			EditNanoTS: 1583512190000000000,
			Host:       "香港",
			DBCS:       []byte("\xa7\xda\xafu\xaa\xba\xa7\xc6\xb1\xe6\xa4\xa3\xadn\xa6A\xc5\xfdMere\xb7d\xb9\xef\xb6H\xa4F\xa1A\xb4N\xc5\xfd\xa6o\xb4\xb6\xb3q\xaa\xba\xb7\xed\xc2\xe5\xa5\xcd\xa9M\xb0\xb5\xb3\xe6\xbf\xcb\xb6\xfd\xb6\xfd\xa7a\xa1I\r"),
			CommentID:  "FfnDytnJUyA:7tkvXPK6iHO5IzFBfJCOwA",
			SortTime:   1583512020002100000,
		},
	}

}
