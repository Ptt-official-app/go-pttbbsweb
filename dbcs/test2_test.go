package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename2            = "M.1607937174.A.081"
	testContentAll2          []byte
	testContent2             []byte
	testSignature2           []byte
	testComment2             []byte
	testFirstCommentsDBCS2   []byte
	testTheRestCommentsDBCS2 []byte
	testContent2Big5         [][]*types.Rune
	testContent2Utf8         [][]*types.Rune

	testFirstComments2 []*schema.Comment
)

func initTest2() {
	testContentAll2, testContent2, testSignature2, testComment2, testFirstCommentsDBCS2, testTheRestCommentsDBCS2 = loadTest(testFilename2)

	testContent2Big5 = [][]*types.Rune{
		{
			{
				Big5:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
			},
		},
		{
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xb0Q\xbd\xd7] \xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1\xe3\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb0Q\xbd\xd7] \xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1\xe3\xa1H"),
			},
		},
		{
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Dec 14 17:12:52 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Mon Dec 14 17:12:52 2020"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xb4\xaa\xa1\xe3\xb9\xd2\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb4\xaa\xa1\xe3\xb9\xd2\xa1\xe3"),
			},
		},
		{},
		{},
		{
			{
				Big5:   []byte("\xadn\xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xadn\xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xacO\xa4\xa3\xacO\xadn\xb5\xa5\xab\xdd 1 min \xa9O\xa1\xe3\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xacO\xa4\xa3\xacO\xadn\xb5\xa5\xab\xdd 1 min \xa9O\xa1\xe3\xa1H"),
			},
		},
		{},
		{
			{
				Big5:   []byte("p \xb9\xf4\xc3\xf8\xc1\xc8\xb0\xda\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("p \xb9\xf4\xc3\xf8\xc1\xc8\xb0\xda\xa1\xe3"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xacO\xa4\xa3\xacO\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xacO\xa4\xa3\xacO\xa9O\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xa7\xda\xa5u\xacO\xb7Q\xadn\xb4\xfa\xb8\xd5\xa4@\xa4U p \xb9\xf4\xacO\xab\xe7\xbc\xcb\xa4l\xb3Q\xa6s\xa6b .DIR \xb8\xcc\xa6\xd3\xa4w XD."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xa5u\xacO\xb7Q\xadn\xb4\xfa\xb8\xd5\xa4@\xa4U p \xb9\xf4\xacO\xab\xe7\xbc\xcb\xa4l\xb3Q\xa6s\xa6b .DIR \xb8\xcc\xa6\xd3\xa4w XD."),
			},
		},
		{},
		{
			{
				Big5:   []byte("\xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
			},
		},
	}

	testContent2Utf8 = [][]*types.Rune{
		{
			{
				Utf8:   "作者: SYSOP () 看板: WhoAmI",
				Big5:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: SYSOP () \xac\xdd\xaaO: WhoAmI"),
			},
		},
		{
			{
				Utf8:   "標題: [討論] 怎樣子才可以賺 p 幣呢～？",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb0Q\xbd\xd7] \xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1\xe3\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb0Q\xbd\xd7] \xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1\xe3\xa1H"),
			},
		},
		{
			{
				Utf8:   "時間: Mon Dec 14 17:12:52 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Dec 14 17:12:52 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Mon Dec 14 17:12:52 2020"),
			},
		},
		{},
		{
			{
				Utf8:   "揪～境～",
				Big5:   []byte("\xb4\xaa\xa1\xe3\xb9\xd2\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb4\xaa\xa1\xe3\xb9\xd2\xa1\xe3"),
			},
		},
		{},
		{},
		{
			{
				Utf8:   "要怎樣子才可以賺 p 幣呢？～",
				Big5:   []byte("\xadn\xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xadn\xab\xe7\xbc\xcb\xa4l\xa4~\xa5i\xa5H\xc1\xc8 p \xb9\xf4\xa9O\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Utf8:   "是不是要等待 1 min 呢～？",
				Big5:   []byte("\xacO\xa4\xa3\xacO\xadn\xb5\xa5\xab\xdd 1 min \xa9O\xa1\xe3\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xacO\xa4\xa3\xacO\xadn\xb5\xa5\xab\xdd 1 min \xa9O\xa1\xe3\xa1H"),
			},
		},
		{},
		{
			{
				Utf8:   "p 幣難賺啊～",
				Big5:   []byte("p \xb9\xf4\xc3\xf8\xc1\xc8\xb0\xda\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("p \xb9\xf4\xc3\xf8\xc1\xc8\xb0\xda\xa1\xe3"),
			},
		},
		{},
		{
			{
				Utf8:   "是不是呢？～",
				Big5:   []byte("\xacO\xa4\xa3\xacO\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xacO\xa4\xa3\xacO\xa9O\xa1H\xa1\xe3"),
			},
		},
		{},
		{
			{
				Utf8:   "我只是想要測試一下 p 幣是怎樣子被存在 .DIR 裡而已 XD.",
				Big5:   []byte("\xa7\xda\xa5u\xacO\xb7Q\xadn\xb4\xfa\xb8\xd5\xa4@\xa4U p \xb9\xf4\xacO\xab\xe7\xbc\xcb\xa4l\xb3Q\xa6s\xa6b .DIR \xb8\xcc\xa6\xd3\xa4w XD."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7\xda\xa5u\xacO\xb7Q\xadn\xb4\xfa\xb8\xd5\xa4@\xa4U p \xb9\xf4\xacO\xab\xe7\xbc\xcb\xa4l\xb3Q\xa6s\xa6b .DIR \xb8\xcc\xa6\xd3\xa4w XD."),
			},
		},
		{},
		{
			{
				Utf8:   "然後呢？～",
				Big5:   []byte("\xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb5M\xab\xe1\xa9O\xa1H\xa1\xe3"),
			},
		},
	}
}
