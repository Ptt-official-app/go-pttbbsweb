package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var testComments0 = []*schema.Comment{
	{
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ"),
		TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
		Owner:      bbs.UUserID("foolfighter"),
		CreateTime: types.NanoTS(1261396680001000000),
		SortTime:   types.NanoTS(1261396680001000000),
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "土柯糞幾成有台北市戶籍啊？",
					Big5:   []byte("\xa4g\xac_\xc1T\xb4X\xa6\xa8\xa6\xb3\xa5x\xa5_\xa5\xab\xa4\xe1\xc4y\xb0\xda\xa1H                        "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("\xa4g\xac_\xc1T\xb4X\xa6\xa8\xa6\xb3\xa5x\xa5_\xa5\xab\xa4\xe1\xc4y\xb0\xda\xa1H                        "),
				},
			},
		},
		MD5:          "Es26f7U0EXdr7Gp4a9N8pQ",
		UpdateNanoTS: types.NanoTS(1334567890000000000),
		TheDate:      "12/21 19:58",
		DBCS:         []byte("\x1b[1;37m\xb1\xc0 \x1b[33mfoolfighter\x1b[m\x1b[33m: \xa4g\xac_\xc1T\xb4X\xa6\xa8\xa6\xb3\xa5x\xa5_\xa5\xab\xa4\xe1\xc4y\xb0\xda\xa1H                        \x1b[m 12/21 19:58\r"),
	},
	{
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg"),
		TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
		Owner:      bbs.UUserID("yehpi"),
		CreateTime: types.NanoTS(1261396680002000000),
		SortTime:   types.NanoTS(1261396680002000000),
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "黃有超大支票（物理）",
					Big5:   []byte("\xb6\xc0\xa6\xb3\xb6W\xa4j\xa4\xe4\xb2\xbc\xa1]\xaa\xab\xb2z\xa1^                                    "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("\xb6\xc0\xa6\xb3\xb6W\xa4j\xa4\xe4\xb2\xbc\xa1]\xaa\xab\xb2z\xa1^                                    "),
				},
			},
		},
		MD5:          "gmrKWXE7BjV-1U89GcPqHg",
		UpdateNanoTS: types.NanoTS(1334567890000000000),
		TheDate:      "12/21 19:58",
		DBCS:         []byte("\x1b[1;37m\xb1\xc0 \x1b[33myehpi\x1b[m\x1b[33m: \xb6\xc0\xa6\xb3\xb6W\xa4j\xa4\xe4\xb2\xbc\xa1]\xaa\xab\xb2z\xa1^                                    \x1b[m 12/21 19:58\r"),
	},
	{
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"),
		TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
		Owner:      bbs.UUserID("lockeyman"),
		CreateTime: types.NanoTS(1261396680003000000),
		SortTime:   types.NanoTS(1261396680003000000),
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "噗",
					Big5:   []byte("\xbcP                                                  "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("\xbcP                                                  "),
				},
			},
		},
		MD5:          "cpqbGyLoF_jIyITF4bv-rQ",
		UpdateNanoTS: types.NanoTS(1334567890000000000),
		TheDate:      "12/21 19:58",
		DBCS:         []byte("\x1b[1;37m\xb1\xc0 \x1b[33mlockeyman\x1b[m\x1b[33m: \xbcP                                                  \x1b[m 12/21 19:58\r"),
	},
	{
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R"),
		RefIDs:     []types.CommentID{"EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"},
		TheType:    ptttype.COMMENT_TYPE_REPLY,
		Owner:      bbs.UUserID("cheinshin"),
		CreateTime: types.NanoTS(1261396680003100000),
		SortTime:   types.NanoTS(1261396680003100000),
		Content: [][]*types.Rune{
			{
				{
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("\r"),
				},
			},
			{
				{
					Utf8:   "test123123",
					Big5:   []byte("test123123"),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("test123123\r"),
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
					Utf8:   "test124124",
					Big5:   []byte("test124124"),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("test124124\r"),
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
					Utf8:   "test125125",
					Big5:   []byte("test125125"),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("test125125\r"),
				},
			},
		},
		MD5:          "VMu8YlVFJ4k06pYnUILy4w",
		IP:           "49.216.65.39",
		Host:         "臺灣",
		EditNanoTS:   types.NanoTS(1608551574000000000),
		UpdateNanoTS: types.NanoTS(1334567890000000000),
		DBCS:         []byte("\r\ntest123123\r\n\r\ntest124124\r\n\r\ntest125125\r\n\r"),
	},
	{
		BBoardID:   bbs.BBoardID("test"),
		ArticleID:  bbs.ArticleID("test"),
		CommentID:  types.CommentID("EYFuT-Ew6AA:ALE6XIa5ARhXunryJTB3xg"),
		TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
		Owner:      bbs.UUserID("deathdancer"),
		CreateTime: types.NanoTS(1261410660000000000),
		SortTime:   types.NanoTS(1261410660000000000),
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "阿北才沒有輸",
					Big5:   []byte("\xaa\xfc\xa5_\xa4~\xa8S\xa6\xb3\xbf\xe9                                      "),
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("\xaa\xfc\xa5_\xa4~\xa8S\xa6\xb3\xbf\xe9                                      "),
				},
			},
		},
		MD5:          "ALE6XIa5ARhXunryJTB3xg",
		UpdateNanoTS: types.NanoTS(1334567890000000000),
		TheDate:      "12/21 23:51",
		DBCS:         []byte("\x1b[1;37m\xb1\xc0 \x1b[33mdeathdancer\x1b[m\x1b[33m: \xaa\xfc\xa5_\xa4~\xa8S\xa6\xb3\xbf\xe9                                      \x1b[m 12/21 23:51\r"),
	},
}
