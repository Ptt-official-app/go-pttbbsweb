package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename5            = "temp0"
	testContentAll5          []byte
	testContent5             []byte
	testSignature5           []byte
	testComment5             []byte
	testFirstCommentsDBCS5   []byte
	testTheRestCommentsDBCS5 []byte
	testContent5Big5         [][]*types.Rune
	testContent5Utf8         [][]*types.Rune

	testFirstComments5     []*schema.Comment
	testFullFirstComments5 []*schema.Comment
)

func initTest5() {
	testContentAll5, testContent5, testSignature5, testComment5, testFirstCommentsDBCS5, testTheRestCommentsDBCS5 = loadTest(testFilename5)

	testContent5Big5 = [][]*types.Rune{
		{
			{
				Big5:   []byte("\xa7@\xaa\xcc: generalife (\xb8\xf2\xc0H Vijay \xa4~\xacO\xa4\xfd\xb9D!) \xac\xdd\xaaO: b885060xx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: generalife (\xb8\xf2\xc0H Vijay \xa4~\xacO\xa4\xfd\xb9D!) \xac\xdd\xaaO: b885060xx\r"),
			},
		},
		{
			{
				Big5:   []byte("\xbc\xd0\xc3D: Re: \xb9w\xa9x"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: Re: \xb9w\xa9x\r"),
			},
		},
		{
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Mar 29 16:09:40 2003"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sat Mar 29 16:09:40 2003\r"),
			},
		},
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{
			{
				Big5:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mmarch20 (Luc)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mmarch20 (Luc)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G\r"),
			},
		},
		{
			{
				Big5:   []byte(": \xab\xe7\xbb\xf2\xa5i\xaf\xe0\xb7|\xa6\xb3\xaaF\xa8F\xa1B\xa4\xd3\xa5\xad, \xb3o\xa8\xc7\xa6a\xa4\xe8\xad\xfe\xb7|\xa6\xb3\xb8\xea\xb0T\xa9x\xbds\xa8\xee..."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(": \xab\xe7\xbb\xf2\xa5i\xaf\xe0\xb7|\xa6\xb3\xaaF\xa8F\xa1B\xa4\xd3\xa5\xad, \xb3o\xa8\xc7\xa6a\xa4\xe8\xad\xfe\xb7|\xa6\xb3\xb8\xea\xb0T\xa9x\xbds\xa8\xee...\r"),
			},
		},
		{
			{
				Big5:   []byte(": \xc0\xb3\xb8\xd3\xacO\xaaF\xa4\xde\xb0\xd5."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(": \xc0\xb3\xb8\xd3\xacO\xaaF\xa4\xde\xb0\xd5.\r"),
			},
		},
		{
			{
				Big5:   []byte("\xa8t\xb2\xce\xa4u\xb5{\xa6n\xb9\xb3\xa4\xa3\xabO\xc3\xd2\xacO\xb8\xea\xb0T\xa9x...."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa8t\xb2\xce\xa4u\xb5{\xa6n\xb9\xb3\xa4\xa3\xabO\xc3\xd2\xacO\xb8\xea\xb0T\xa9x....\r"),
			},
		},
		{
			{
				Big5:   []byte("\xa4\xa3\xb9L\xadn\xb3Q\xa4\xc0\xa8\xec\xaaF\xa8F\xa5h\xaa\xba\xb8\xdc\xa4]\xba\xe2\xacO\xbe\xf7\xb7|\xc3\xf8\xb1o :P"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4\xa3\xb9L\xadn\xb3Q\xa4\xc0\xa8\xec\xaaF\xa8F\xa5h\xaa\xba\xb8\xdc\xa4]\xba\xe2\xacO\xbe\xf7\xb7|\xc3\xf8\xb1o :P\r"),
			},
		},
	}

	testContent5Utf8 = [][]*types.Rune{
		{
			{
				Utf8:   "作者: generalife (跟隨 Vijay 才是王道!) 看板: b885060xx",
				Big5:   []byte("\xa7@\xaa\xcc: generalife (\xb8\xf2\xc0H Vijay \xa4~\xacO\xa4\xfd\xb9D!) \xac\xdd\xaaO: b885060xx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: generalife (\xb8\xf2\xc0H Vijay \xa4~\xacO\xa4\xfd\xb9D!) \xac\xdd\xaaO: b885060xx\r"),
			},
		},
		{
			{
				Utf8:   "標題: Re: 預官",
				Big5:   []byte("\xbc\xd0\xc3D: Re: \xb9w\xa9x"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: Re: \xb9w\xa9x\r"),
			},
		},
		{
			{
				Utf8:   "時間: Sat Mar 29 16:09:40 2003",
				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Mar 29 16:09:40 2003"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sat Mar 29 16:09:40 2003\r"),
			},
		},
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{
			{
				Utf8:   "※ 引述《march20 (Luc)》之銘言：",
				Big5:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mmarch20 (Luc)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mmarch20 (Luc)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G\r"),
			},
		},
		{
			{
				Utf8:   ": 怎麼可能會有東沙、太平, 這些地方哪會有資訊官編制...",
				Big5:   []byte(": \xab\xe7\xbb\xf2\xa5i\xaf\xe0\xb7|\xa6\xb3\xaaF\xa8F\xa1B\xa4\xd3\xa5\xad, \xb3o\xa8\xc7\xa6a\xa4\xe8\xad\xfe\xb7|\xa6\xb3\xb8\xea\xb0T\xa9x\xbds\xa8\xee..."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(": \xab\xe7\xbb\xf2\xa5i\xaf\xe0\xb7|\xa6\xb3\xaaF\xa8F\xa1B\xa4\xd3\xa5\xad, \xb3o\xa8\xc7\xa6a\xa4\xe8\xad\xfe\xb7|\xa6\xb3\xb8\xea\xb0T\xa9x\xbds\xa8\xee...\r"),
			},
		},
		{
			{
				Utf8:   ": 應該是東引啦.",
				Big5:   []byte(": \xc0\xb3\xb8\xd3\xacO\xaaF\xa4\xde\xb0\xd5."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(": \xc0\xb3\xb8\xd3\xacO\xaaF\xa4\xde\xb0\xd5.\r"),
			},
		},
		{
			{
				Utf8:   "系統工程好像不保證是資訊官....",
				Big5:   []byte("\xa8t\xb2\xce\xa4u\xb5{\xa6n\xb9\xb3\xa4\xa3\xabO\xc3\xd2\xacO\xb8\xea\xb0T\xa9x...."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa8t\xb2\xce\xa4u\xb5{\xa6n\xb9\xb3\xa4\xa3\xabO\xc3\xd2\xacO\xb8\xea\xb0T\xa9x....\r"),
			},
		},
		{
			{
				Utf8:   "不過要被分到東沙去的話也算是機會難得 :P",
				Big5:   []byte("\xa4\xa3\xb9L\xadn\xb3Q\xa4\xc0\xa8\xec\xaaF\xa8F\xa5h\xaa\xba\xb8\xdc\xa4]\xba\xe2\xacO\xbe\xf7\xb7|\xc3\xf8\xb1o :P"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4\xa3\xb9L\xadn\xb3Q\xa4\xc0\xa8\xec\xaaF\xa8F\xa5h\xaa\xba\xb8\xdc\xa4]\xba\xe2\xacO\xbe\xf7\xb7|\xc3\xf8\xb1o :P\r"),
			},
		},
	}
}
