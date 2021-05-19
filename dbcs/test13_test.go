package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename13            = "temp8"
	testContentAll13          []byte
	testContent13             []byte
	testSignature13           []byte
	testComment13             []byte
	testFirstCommentsDBCS13   []byte
	testTheRestCommentsDBCS13 []byte
	testContent13Big5         [][]*types.Rune
	testContent13Utf8         [][]*types.Rune

	testFirstComments13     []*schema.Comment
	testFullFirstComments13 []*schema.Comment

	testTheRestComments13 []*schema.Comment
)

func initTest13() {
	testContentAll13, testContent13, testSignature13, testComment13, testFirstCommentsDBCS13, testTheRestCommentsDBCS13 = loadTest(testFilename13)

	testContent13Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: ledia (totally defeated) \xac\xdd\xaaO: b885060xx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: ledia (totally defeated) \xac\xdd\xaaO: b885060xx\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: \xa4j\xaea\xaa`\xb7N\xa4@\xa4U"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: \xa4j\xaea\xaa`\xb7N\xa4@\xa4U\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Tue Nov 19 12:21:41 2002"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Tue Nov 19 12:21:41 2002\r"),
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

				Big5:   []byte("  \xa6\xb3\xc5\xb2\xa9\xf3\xc2I\xc2I\xc2I\xa6]\xac\xb0\xbf\xef\xbf\xf9\xbd\xd2\xa5i\xaf\xe0\xadn\xa9\xb5\xb2\xa6\xa4F.."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa6\xb3\xc5\xb2\xa9\xf3\xc2I\xc2I\xc2I\xa6]\xac\xb0\xbf\xef\xbf\xf9\xbd\xd2\xa5i\xaf\xe0\xadn\xa9\xb5\xb2\xa6\xa4F..\r"),
			},
		},
		{ //5
			{

				Big5:   []byte("  \xa4j\xaea\xb3\xcc\xa6n\xa6\xdb\xa4v check \xa4@\xa4U\xa6\xdb\xa4v\xaa\xba\xb1M\xc3D\xac\xe3\xa8s\xacO\xb4X\xbe\xc7\xa4\xc0\xaa\xba"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa4j\xaea\xb3\xcc\xa6n\xa6\xdb\xa4v check \xa4@\xa4U\xa6\xdb\xa4v\xaa\xba\xb1M\xc3D\xac\xe3\xa8s\xacO\xb4X\xbe\xc7\xa4\xc0\xaa\xba\r"),
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

				Big5:   []byte("  \xc2I\xc2I\xc2I\xbf\xef\xa8\xec\xaa\xba\xacO\xa4@\xbe\xc7\xa4\xc0\xaa\xba, \xac\xe3\xa8s\xa9\xd2\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\""),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xc2I\xc2I\xc2I\xbf\xef\xa8\xec\xaa\xba\xacO\xa4@\xbe\xc7\xa4\xc0\xaa\xba, \xac\xe3\xa8s\xa9\xd2\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\"\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("  \xa5\xbf\xbdT\xaa\xba\xc0\xb3\xb8\xd3\xacO\xa8\xe2\xbe\xc7\xa4\xc0\xaa\xba, \xa4j\xbe\xc7\xb3\xa1\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\""),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa5\xbf\xbdT\xaa\xba\xc0\xb3\xb8\xd3\xacO\xa8\xe2\xbe\xc7\xa4\xc0\xaa\xba, \xa4j\xbe\xc7\xb3\xa1\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\"\r"),
			},
		},
		{ //9
			{

				Big5:   []byte("  \xa4\xa3\xadn\xc3h\xba\xc3, \xbd\xd2\xa6W\xa4@\xbc\xcb ^^;"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa4\xa3\xadn\xc3h\xba\xc3, \xbd\xd2\xa6W\xa4@\xbc\xcb ^^;\r"),
			},
		},
		{ //10
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //11
			{

				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
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

				Big5:   []byte("  \xc3\xf8\xa5H\xa7\xdc\xa9\xda"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xc3\xf8\xa5H\xa7\xdc\xa9\xda\r"),
			},
		},
		{ //14
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
	}

	testContent13Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: ledia (totally defeated) 看板: b885060xx",
				Big5:   []byte("\xa7@\xaa\xcc: ledia (totally defeated) \xac\xdd\xaaO: b885060xx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: ledia (totally defeated) \xac\xdd\xaaO: b885060xx\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: 大家注意一下",
				Big5:   []byte("\xbc\xd0\xc3D: \xa4j\xaea\xaa`\xb7N\xa4@\xa4U"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: \xa4j\xaea\xaa`\xb7N\xa4@\xa4U\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Tue Nov 19 12:21:41 2002",
				Big5:   []byte("\xae\xc9\xb6\xa1: Tue Nov 19 12:21:41 2002"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Tue Nov 19 12:21:41 2002\r"),
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
				Utf8:   "  有鑑於點點點因為選錯課可能要延畢了..",
				Big5:   []byte("  \xa6\xb3\xc5\xb2\xa9\xf3\xc2I\xc2I\xc2I\xa6]\xac\xb0\xbf\xef\xbf\xf9\xbd\xd2\xa5i\xaf\xe0\xadn\xa9\xb5\xb2\xa6\xa4F.."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa6\xb3\xc5\xb2\xa9\xf3\xc2I\xc2I\xc2I\xa6]\xac\xb0\xbf\xef\xbf\xf9\xbd\xd2\xa5i\xaf\xe0\xadn\xa9\xb5\xb2\xa6\xa4F..\r"),
			},
		},
		{ //5
			{
				Utf8:   "  大家最好自己 check 一下自己的專題研究是幾學分的",
				Big5:   []byte("  \xa4j\xaea\xb3\xcc\xa6n\xa6\xdb\xa4v check \xa4@\xa4U\xa6\xdb\xa4v\xaa\xba\xb1M\xc3D\xac\xe3\xa8s\xacO\xb4X\xbe\xc7\xa4\xc0\xaa\xba"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa4j\xaea\xb3\xcc\xa6n\xa6\xdb\xa4v check \xa4@\xa4U\xa6\xdb\xa4v\xaa\xba\xb1M\xc3D\xac\xe3\xa8s\xacO\xb4X\xbe\xc7\xa4\xc0\xaa\xba\r"),
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
				Utf8:   "  點點點選到的是一學分的, 研究所的 \"專題研究\"",
				Big5:   []byte("  \xc2I\xc2I\xc2I\xbf\xef\xa8\xec\xaa\xba\xacO\xa4@\xbe\xc7\xa4\xc0\xaa\xba, \xac\xe3\xa8s\xa9\xd2\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\""),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xc2I\xc2I\xc2I\xbf\xef\xa8\xec\xaa\xba\xacO\xa4@\xbe\xc7\xa4\xc0\xaa\xba, \xac\xe3\xa8s\xa9\xd2\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\"\r"),
			},
		},
		{ //8
			{
				Utf8:   "  正確的應該是兩學分的, 大學部的 \"專題研究\"",
				Big5:   []byte("  \xa5\xbf\xbdT\xaa\xba\xc0\xb3\xb8\xd3\xacO\xa8\xe2\xbe\xc7\xa4\xc0\xaa\xba, \xa4j\xbe\xc7\xb3\xa1\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\""),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa5\xbf\xbdT\xaa\xba\xc0\xb3\xb8\xd3\xacO\xa8\xe2\xbe\xc7\xa4\xc0\xaa\xba, \xa4j\xbe\xc7\xb3\xa1\xaa\xba \"\xb1M\xc3D\xac\xe3\xa8s\"\r"),
			},
		},
		{ //9
			{
				Utf8:   "  不要懷疑, 課名一樣 ^^;",
				Big5:   []byte("  \xa4\xa3\xadn\xc3h\xba\xc3, \xbd\xd2\xa6W\xa4@\xbc\xcb ^^;"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xa4\xa3\xadn\xc3h\xba\xc3, \xbd\xd2\xa6W\xa4@\xbc\xcb ^^;\r"),
			},
		},
		{ //10
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //11
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
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
				Utf8:   "  難以抗拒",
				Big5:   []byte("  \xc3\xf8\xa5H\xa7\xdc\xa9\xda"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xc3\xf8\xa5H\xa7\xdc\xa9\xda\r"),
			},
		},
		{ //14
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
	}

	testFirstComments13 = []*schema.Comment{
		{ //0
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("easy1"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "下學期 一次修兩個專題?",
						Big5:   []byte("\xa4U\xbe\xc7\xb4\xc1 \xa4@\xa6\xb8\xad\xd7\xa8\xe2\xad\xd3\xb1M\xc3D?"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4U\xbe\xc7\xb4\xc1 \xa4@\xa6\xb8\xad\xd7\xa8\xe2\xad\xd3\xb1M\xc3D?"),
					},
				},
			},
			MD5:     "5m3AB1JdQSO8lw6ygnPVNg",
			TheDate: "11/19",
			IP:      "140.112.28.31",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33measy1\x1b[m\x1b[33m:\xa4U\xbe\xc7\xb4\xc1 \xa4@\xa6\xb8\xad\xd7\xa8\xe2\xad\xd3\xb1M\xc3D?\x1b[m                        \xb1\xc0  140.112.28.31 11/19\r"),
		},
		{ //1
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("Foxwall"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "學校沒有開...除非有教授要開專題一..",
						Big5:   []byte("\xbe\xc7\xae\xd5\xa8S\xa6\xb3\xb6}...\xb0\xa3\xabD\xa6\xb3\xb1\xd0\xb1\xc2\xadn\xb6}\xb1M\xc3D\xa4@.."),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbe\xc7\xae\xd5\xa8S\xa6\xb3\xb6}...\xb0\xa3\xabD\xa6\xb3\xb1\xd0\xb1\xc2\xadn\xb6}\xb1M\xc3D\xa4@.."),
					},
				},
			},
			MD5:     "POPKa8tQwVnBsvb31Zjtnw",
			TheDate: "11/19",
			IP:      "ntucst",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mFoxwall\x1b[m\x1b[33m:\xbe\xc7\xae\xd5\xa8S\xa6\xb3\xb6}...\xb0\xa3\xabD\xa6\xb3\xb1\xd0\xb1\xc2\xadn\xb6}\xb1M\xc3D\xa4@..\x1b[m         \xb1\xc0         ntucst 11/19\r"),
		},
	}

	testFullFirstComments13 = []*schema.Comment{
		{ //0
			BBoardID:  "test",
			ArticleID: "test13",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("easy1"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "下學期 一次修兩個專題?",
						Big5:   []byte("\xa4U\xbe\xc7\xb4\xc1 \xa4@\xa6\xb8\xad\xd7\xa8\xe2\xad\xd3\xb1M\xc3D?"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4U\xbe\xc7\xb4\xc1 \xa4@\xa6\xb8\xad\xd7\xa8\xe2\xad\xd3\xb1M\xc3D?"),
					},
				},
			},
			MD5:        "5m3AB1JdQSO8lw6ygnPVNg",
			TheDate:    "11/19",
			IP:         "140.112.28.31",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33measy1\x1b[m\x1b[33m:\xa4U\xbe\xc7\xb4\xc1 \xa4@\xa6\xb8\xad\xd7\xa8\xe2\xad\xd3\xb1M\xc3D?\x1b[m                        \xb1\xc0  140.112.28.31 11/19\r"),
			CommentID:  "DmaUMLDMVEA:5m3AB1JdQSO8lw6ygnPVNg",
			CreateTime: 1037635200000000000,
			SortTime:   1037679701001000000,
		},
		{ //1
			BBoardID:  "test",
			ArticleID: "test13",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("Foxwall"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "學校沒有開...除非有教授要開專題一..",
						Big5:   []byte("\xbe\xc7\xae\xd5\xa8S\xa6\xb3\xb6}...\xb0\xa3\xabD\xa6\xb3\xb1\xd0\xb1\xc2\xadn\xb6}\xb1M\xc3D\xa4@.."),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbe\xc7\xae\xd5\xa8S\xa6\xb3\xb6}...\xb0\xa3\xabD\xa6\xb3\xb1\xd0\xb1\xc2\xadn\xb6}\xb1M\xc3D\xa4@.."),
					},
				},
			},
			MD5:        "POPKa8tQwVnBsvb31Zjtnw",
			TheDate:    "11/19",
			IP:         "ntucst",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mFoxwall\x1b[m\x1b[33m:\xbe\xc7\xae\xd5\xa8S\xa6\xb3\xb6}...\xb0\xa3\xabD\xa6\xb3\xb1\xd0\xb1\xc2\xadn\xb6}\xb1M\xc3D\xa4@..\x1b[m         \xb1\xc0         ntucst 11/19\r"),
			CommentID:  "DmaUMOxX3AA:POPKa8tQwVnBsvb31Zjtnw",
			CreateTime: 1037679702000000000,
			SortTime:   1037679702000000000,
		},
	}
}
