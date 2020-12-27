package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename7            = "temp2"
	testContentAll7          []byte
	testContent7             []byte
	testSignature7           []byte
	testComment7             []byte
	testFirstCommentsDBCS7   []byte
	testTheRestCommentsDBCS7 []byte
	testContent7Big5         [][]*types.Rune
	testContent7Utf8         [][]*types.Rune

	testFirstComments7 []*schema.Comment
)

func initTest7() {
	testContentAll7, testContent7, testSignature7, testComment7, testFirstCommentsDBCS7, testTheRestCommentsDBCS7 = loadTest(testFilename7)

	testContent7Big5 = [][]*types.Rune{
		{ //0
			{
				Big5:   []byte("\xa7@\xaa\xcc: Psycap (\xb1a\xa4p\xaaB\xa4\xcd~~) \xac\xdd\xaaO: b885060xx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //1
			{
				Big5:   []byte("\xbc\xd0\xc3D: \xae\xd1\xa8\xf7\xa5X\xc4l\xc5o^100%"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //2
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Mar 31 14:29:44 2003"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //3
		{ //4
			{
				Big5:   []byte("\xa5\xbb\xafZ\xae\xd1\xa8\xf7\xa6W\xb3\xe6 :"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //5
		{ //6
			{
				Big5:   []byte("aaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //7
			{
				Big5:   []byte("(\xa5\xfe\xb9p\xa5\xb4\xa4\xfd)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //8
		{ //9
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //10
		{ //11
			{
				Big5:   []byte("http://reg.aca.ntu.edu.tw/win911.htm"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //12
		{ //13
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	testContent7Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: Psycap (帶小朋友~~) 看板: b885060xx",
				Big5:   []byte("\xa7@\xaa\xcc: Psycap (\xb1a\xa4p\xaaB\xa4\xcd~~) \xac\xdd\xaaO: b885060xx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //1
			{
				Utf8:   "標題: 書卷出爐囉^100%",
				Big5:   []byte("\xbc\xd0\xc3D: \xae\xd1\xa8\xf7\xa5X\xc4l\xc5o^100%"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //2
			{
				Utf8:   "時間: Mon Mar 31 14:29:44 2003",
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Mar 31 14:29:44 2003"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //3
		{ //4
			{
				Utf8:   "本班書卷名單 :",
				Big5:   []byte("\xa5\xbb\xafZ\xae\xd1\xa8\xf7\xa6W\xb3\xe6 :"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //5
		{ //6
			{
				Utf8:   "aaaaaaaa",
				Big5:   []byte("aaaaaaaa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //7
			{
				Utf8:   "(全雷打王)",
				Big5:   []byte("(\xa5\xfe\xb9p\xa5\xb4\xa4\xfd)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //8
		{ //9
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //10
		{ //11
			{
				Utf8:   "http://reg.aca.ntu.edu.tw/win911.htm",
				Big5:   []byte("http://reg.aca.ntu.edu.tw/win911.htm"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //12
		{ //13
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}
}
