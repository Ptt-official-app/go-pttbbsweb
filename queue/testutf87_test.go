package queue

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var (
	testUtf8Filename7            = "M.1644506392.A.03C.utf8"
	testUtf8ContentAll7          []byte
	testUtf8Content7             []byte
	testUtf8Signature7           []byte
	testUtf8Comment7             []byte
	testUtf8FirstCommentsDBCS7   []byte
	testUtf8TheRestCommentsDBCS7 []byte
	testUtf8Content7Utf8         [][]*types.Rune

	testUtf8FirstComments7       []*schema.Comment
	testUtf8FullFirstComments7   []*schema.Comment
	testUtf8TheRestComments7     []*schema.Comment
	testUtf8FullTheRestComments7 []*schema.Comment
	testUtf8FullComments7        []*schema.Comment
)

func initTestUtf87() {
	testUtf8ContentAll7, testUtf8Content7, testUtf8Signature7, testUtf8Comment7, testUtf8FirstCommentsDBCS7, testUtf8TheRestCommentsDBCS7 = loadTest(testUtf8Filename7)

	testUtf8Content7Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:    "作者: PttACT (PTT活動部專用帳號) 看板: SYSOP",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: PttACT (PTT活動部專用帳號) 看板: SYSOP",
			},
		},
		{ // 1
			{
				Utf8:    "標題: [情報] PTT自創曲大賽 決選投票抽獎品",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [情報] PTT自創曲大賽 決選投票抽獎品",
			},
		},
		{ // 7
			{
				Utf8:    "時間: Fri Jan 26 17:18:22 2018",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Fri Jan 26 17:18:22 2018",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "PTT自創曲大賽決選終於開始囉！即起進入決選試聽階段！",
				DBCSStr: "PTT自創曲大賽決選終於開始囉！即起進入決選試聽階段！",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "參賽者十強的第二首曲子已於 ",
				DBCSStr: "參賽者十強的第二首曲子已於 ",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "OriginalSong板",
				DBCSStr: "\x1b[36;1mOriginalSong板",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_CYAN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    " 釋出",
				DBCSStr: "\x1b[m 釋出",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "投票日期預計為：1/29 ~ 2/9",
				DBCSStr: "投票日期預計為：1/29 ~ 2/9",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "決選推文票選除了",
				DBCSStr: "決選推文票選除了",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
			{
				Utf8:    "隨機選五位鄉民發放 3000 批幣",
				DBCSStr: "\x1b[33;1m隨機選五位鄉民發放 3000 批幣",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "另外再加碼，隨機抽出一位參與決選票選的鄉民",
				DBCSStr: "另外再加碼，隨機抽出一位參與決選票選的鄉民",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "贈送紀念品「純手工印刷．PTT 特別版 LOGO 束口袋」一個！！",
				DBCSStr: "\x1b[33;1m贈送紀念品「純手工印刷．PTT 特別版 LOGO 束口袋」一個！！",
				Color0:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1:  types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
			},
			{
				Utf8:    "",
				DBCSStr: "\x1b[m",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "超級限量款，保證絕版停產，除了比賽前五名，就只有你有喔！",
				DBCSStr: "超級限量款，保證絕版停產，除了比賽前五名，就只有你有喔！",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "如圖：https://imgur.com/a/yeRfy",
				DBCSStr: "如圖：https://imgur.com/a/yeRfy",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "心動了嗎？詳情快到 OriginalSong板 先聽為快！",
				DBCSStr: "心動了嗎？詳情快到 OriginalSong板 先聽為快！",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "PTT活動部",
				DBCSStr: "PTT活動部",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 5
			{
				Utf8:    "--",
				DBCSStr: "--",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "市民廣場     報告站長 PTT咬我",
				DBCSStr: "市民廣場     報告站長 PTT咬我",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "   PttAct        本站 Σ批踢踢活動部",
				DBCSStr: "   PttAct        本站 Σ批踢踢活動部",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "     OriginalSong  原創 ◎2017 PTT自創曲大賽",
				DBCSStr: "     OriginalSong  原創 ◎2017 PTT自創曲大賽",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
	}

	testUtf8FirstComments7 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "OriginalSong",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "OriginalSong",
					},
				},
			},
			MD5:     "URNg5eSVtEPmB2nG_5sLdQ",
			TheDate: "01/26 17:19",
			DBCSStr: "※ \x1b[1;32mPttACT\x1b[0;32m:轉錄至看板 OriginalSong\x1b[m                                  01/26 17:19",
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "PttEarnMoney",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "PttEarnMoney",
					},
				},
			},
			TheDate: "01/26 17:50",
			MD5:     "Fd3ySk4FHQSmAr3wLzf9yw",
			DBCSStr: "※ \x1b[1;32mPttACT\x1b[0;32m:轉錄至看板 PttEarnMoney\x1b[m                                  01/26 17:50",
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Gossiping",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Gossiping",
					},
				},
			},
			MD5:     "FItlEmTBYMrTQ9IcXLP9iA",
			TheDate: "01/26 22:25",
			DBCSStr: "※ \x1b[1;32msimons\x1b[0;32m:轉錄至看板 Gossiping\x1b[m                                     01/26 22:25",
		},
		{ // 3
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "WomenTalk",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "WomenTalk",
					},
				},
			},
			MD5:     "lTyGGzulrnfmw976XK-XyQ",
			TheDate: "01/26 22:34",
			DBCSStr: "※ \x1b[1;32msimons\x1b[0;32m:轉錄至看板 WomenTalk\x1b[m                                     01/26 22:34",
		},
		{ // 4
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("mono5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "talk",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "talk",
					},
				},
			},
			MD5:     "eVZblSu5odix2sJ0jvmbWQ",
			TheDate: "01/26 22:39",
			DBCSStr: "※ \x1b[1;32mmono5566\x1b[0;32m:轉錄至看板 talk\x1b[m                                        01/26 22:39",
		},
		{ // 5
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("a6234709"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "marvel",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "marvel",
					},
				},
			},
			MD5:     "WSpDkjp4VI6We4sEvFyjtg",
			TheDate: "01/27 00:43",
			DBCSStr: "※ \x1b[1;32ma6234709\x1b[0;32m:轉錄至看板 marvel\x1b[m                                      01/27 00:43",
		},
		{ // 6
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "KMU",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "KMU",
					},
				},
			},
			MD5:     "1PEozd6yqQDjYDjPpbDMjA",
			TheDate: "01/28 12:10",
			DBCSStr: "※ \x1b[1;32mgghost1002\x1b[0;32m:轉錄至看板 KMU\x1b[m                                       01/28 12:10",
		},
		{ // 7
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Kaohsiung",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Kaohsiung",
					},
				},
			},
			MD5:     "VS-pLk73v09FJZG1B9Z7HA",
			TheDate: "01/28 12:11",
			DBCSStr: "※ \x1b[1;32mgghost1002\x1b[0;32m:轉錄至看板 Kaohsiung\x1b[m                                 01/28 12:11",
		},
		{ // 8
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("fennyccc"),
			MD5:     "C6N7KYHhMqgCdwwpmZmQqw",
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "讚",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "讚",
					},
				},
			},
			IP:      "115.82.149.161",
			TheDate: "01/28 12:34",
			DBCSStr: "\x1b[1;37m推 \x1b[33mfennyccc\x1b[m\x1b[33m: 讚                                    \x1b[m 115.82.149.161 01/28 12:34",
		},
		{ // 9
			TheType: ptttype.COMMENT_TYPE_FORWARD,
			Owner:   bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "某隱形看板",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "某隱形看板",
					},
				},
			},
			MD5:     "mw4YMh6JZFC9-gdXfv_jJw",
			TheDate: "01/29 02:39",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至某隱形看板\x1b[m                                         01/29 02:39",
		},
	}
	testUtf8FullFirstComments7 = []*schema.Comment{
		{ // 0
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "OriginalSong",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "OriginalSong",
					},
				},
			},
			MD5:     "URNg5eSVtEPmB2nG_5sLdQ",
			TheDate: "01/26 17:19",
			DBCSStr: "※ \x1b[1;32mPttACT\x1b[0;32m:轉錄至看板 OriginalSong\x1b[m                                  01/26 17:19",

			CommentID:  "FQ1RkrLEKAA:URNg5eSVtEPmB2nG_5sLdQ",
			CreateTime: 1516958340000000000,
			SortTime:   1516958340000000000,
		},
		{ // 1
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("PttACT"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "PttEarnMoney",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "PttEarnMoney",
					},
				},
			},
			TheDate: "01/26 17:50",
			MD5:     "Fd3ySk4FHQSmAr3wLzf9yw",
			DBCSStr: "※ \x1b[1;32mPttACT\x1b[0;32m:轉錄至看板 PttEarnMoney\x1b[m                                  01/26 17:50",

			CommentID:  "FQ1TQ8Nn0AA:Fd3ySk4FHQSmAr3wLzf9yw",
			CreateTime: 1516960200000000000,
			SortTime:   1516960200000000000,
		},
		{ // 2
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Gossiping",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Gossiping",
					},
				},
			},
			MD5:     "FItlEmTBYMrTQ9IcXLP9iA",
			TheDate: "01/26 22:25",
			DBCSStr: "※ \x1b[1;32msimons\x1b[0;32m:轉錄至看板 Gossiping\x1b[m                                     01/26 22:25",

			CommentID:  "FQ1iRXgLWAA:FItlEmTBYMrTQ9IcXLP9iA",
			CreateTime: 1516976700000000000,
			SortTime:   1516976700000000000,
		},
		{ // 3
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("simons"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "WomenTalk",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "WomenTalk",
					},
				},
			},
			MD5:     "lTyGGzulrnfmw976XK-XyQ",
			TheDate: "01/26 22:34",
			DBCSStr: "※ \x1b[1;32msimons\x1b[0;32m:轉錄至看板 WomenTalk\x1b[m                                     01/26 22:34",

			CommentID:  "FQ1iwzKNcAA:lTyGGzulrnfmw976XK-XyQ",
			CreateTime: 1516977240000000000,
			SortTime:   1516977240000000000,
		},
		{ // 4
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("mono5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "talk",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "talk",
					},
				},
			},
			MD5:     "eVZblSu5odix2sJ0jvmbWQ",
			TheDate: "01/26 22:39",
			DBCSStr: "※ \x1b[1;32mmono5566\x1b[0;32m:轉錄至看板 talk\x1b[m                                        01/26 22:39",

			CommentID:  "FQ1jCQvyKAA:eVZblSu5odix2sJ0jvmbWQ",
			CreateTime: 1516977540000000000,
			SortTime:   1516977540000000000,
		},
		{ // 5
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("a6234709"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "marvel",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "marvel",
					},
				},
			},
			MD5:     "WSpDkjp4VI6We4sEvFyjtg",
			TheDate: "01/27 00:43",
			DBCSStr: "※ \x1b[1;32ma6234709\x1b[0;32m:轉錄至看板 marvel\x1b[m                                      01/27 00:43",

			CommentID:  "FQ1pzU6AyAA:WSpDkjp4VI6We4sEvFyjtg",
			CreateTime: 1516984980000000000,
			SortTime:   1516984980000000000,
		},
		{ // 6
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "KMU",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "KMU",
					},
				},
			},
			MD5:     "1PEozd6yqQDjYDjPpbDMjA",
			TheDate: "01/28 12:10",
			DBCSStr: "※ \x1b[1;32mgghost1002\x1b[0;32m:轉錄至看板 KMU\x1b[m                                       01/28 12:10",

			CommentID:  "FQ3d3ydE8AA:1PEozd6yqQDjYDjPpbDMjA",
			CreateTime: 1517112600000000000,
			SortTime:   1517112600000000000,
		},
		{ // 7
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("gghost1002"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Kaohsiung",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Kaohsiung",
					},
				},
			},
			MD5:     "VS-pLk73v09FJZG1B9Z7HA",
			TheDate: "01/28 12:11",
			DBCSStr: "※ \x1b[1;32mgghost1002\x1b[0;32m:轉錄至看板 Kaohsiung\x1b[m                                 01/28 12:11",

			CommentID:  "FQ3d7R-MSAA:VS-pLk73v09FJZG1B9Z7HA",
			CreateTime: 1517112660000000000,
			SortTime:   1517112660000000000,
		},
		{ // 8
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("fennyccc"),
			MD5:       "C6N7KYHhMqgCdwwpmZmQqw",
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "讚",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "讚",
					},
				},
			},
			IP:      "115.82.149.161",
			TheDate: "01/28 12:34",
			DBCSStr: "\x1b[1;37m推 \x1b[33mfennyccc\x1b[m\x1b[33m: 讚                                    \x1b[m 115.82.149.161 01/28 12:34",

			CommentID:  "FQ3fLm31MAA:C6N7KYHhMqgCdwwpmZmQqw",
			CreateTime: 1517114040000000000,
			SortTime:   1517114040000000000,
		},
		{ // 9
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "某隱形看板",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "某隱形看板",
					},
				},
			},
			MD5:     "mw4YMh6JZFC9-gdXfv_jJw",
			TheDate: "01/29 02:39",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至某隱形看板\x1b[m                                         01/29 02:39",

			CommentID:  "FQ4NSvFyqAA:mw4YMh6JZFC9-gdXfv_jJw",
			CreateTime: 1517164740000000000,
			SortTime:   1517164740000000000,
		},
	}
	testUtf8TheRestComments7 = []*schema.Comment{}
	testUtf8FullTheRestComments7 = []*schema.Comment{
		{ // 0
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Lifeismoney",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Lifeismoney",
					},
				},
			},
			MD5:     "qzphU6MBgHYz-gjw1qm6-A",
			TheDate: "01/29 02:44",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至看板 Lifeismoney\x1b[m                                   01/29 02:44",

			CommentID:  "FQ4NkMrXYAA:qzphU6MBgHYz-gjw1qm6-A",
			CreateTime: 1517165040000000000,
			SortTime:   1517165040000000000,
		},
		{ // 1
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "MenTalk",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "MenTalk",
					},
				},
			},
			MD5:     "eWwz_EGc-UA2cK-c23FHqw",
			TheDate: "01/29 02:47",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至看板 MenTalk\x1b[m                                       01/29 02:47",

			CommentID:  "FQ4NurOtaAA:eWwz_EGc-UA2cK-c23FHqw",
			CreateTime: 1517165220000000000,
			SortTime:   1517165220000000000,
		},
		{ // 2
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "SENIORHIGH",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "SENIORHIGH",
					},
				},
			},
			MD5:     "4S0V66Vhgg3qK0Ciw820aQ",
			TheDate: "01/29 02:53",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至看板 SENIORHIGH\x1b[m                                    01/29 02:53",

			CommentID:  "FQ4ODoVZeAA:4S0V66Vhgg3qK0Ciw820aQ",
			CreateTime: 1517165580000000000,
			SortTime:   1517165580000000000,
		},
		{ // 3
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "某隱形看板",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "某隱形看板",
					},
				},
			},
			MD5:     "RDrYJM4wpMrPnDCiKKD1xQ",
			TheDate: "01/29 02:58",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至某隱形看板\x1b[m                                         01/29 02:58",

			CommentID:  "FQ4OVF6-MAA:RDrYJM4wpMrPnDCiKKD1xQ",
			CreateTime: 1517165880000000000,
			SortTime:   1517165880000000000,
		},
		{ // 4
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("frojet"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "ONE_PIECE",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "ONE_PIECE",
					},
				},
			},
			MD5:     "osAeoHL3ZxxionQ6O674gw",
			TheDate: "01/29 04:14",
			DBCSStr: "※ \x1b[1;32mfrojet\x1b[0;32m:轉錄至看板 ONE_PIECE\x1b[m                                     01/29 04:14",

			CommentID:  "FQ4SehPsUAA:osAeoHL3ZxxionQ6O674gw",
			CreateTime: 1517170440000000000,
			SortTime:   1517170440000000000,
		},
		{ // 5
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("a6234709"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "movie",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "movie",
					},
				},
			},
			MD5:     "n0ZUYVJ4WbB4XtMrq1RhNA",
			TheDate: "01/29 10:53",
			DBCSStr: "※ \x1b[1;32ma6234709\x1b[0;32m:轉錄至看板 movie\x1b[m                                       01/29 10:53",

			CommentID:  "FQ4oQAseeAA:n0ZUYVJ4WbB4XtMrq1RhNA",
			CreateTime: 1517194380000000000,
			SortTime:   1517194380000000000,
		},
		{ // 6
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("jasome"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Gossiping",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Gossiping",
					},
				},
			},
			MD5:     "H5n3DA1GnmpSQS4Iu-zK9w",
			TheDate: "01/29 17:07",
			DBCSStr: "※ \x1b[1;32mjasome\x1b[0;32m:轉錄至看板 Gossiping\x1b[m                                     01/29 17:07",

			CommentID:  "FQ48qMNZCAA:H5n3DA1GnmpSQS4Iu-zK9w",
			CreateTime: 1517216820000000000,
			SortTime:   1517216820000000000,
		},
		{ // 7
			BBoardID:  "test",
			ArticleID: "test7",
			TheType:   ptttype.COMMENT_TYPE_FORWARD,
			Owner:     bbs.UUserID("gogin"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "Wanted",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "Wanted",
					},
				},
			},
			MD5:     "RLzKgitYkrjOFGqJwCrrBQ",
			TheDate: "01/30 08:56",
			DBCSStr: "※ \x1b[1;32mgogin\x1b[0;32m:轉錄至看板 Wanted\x1b[m                                         01/30 08:56",

			CommentID:  "Fe6EGH8YQAA:RLzKgitYkrjOFGqJwCrrBQ",
			CreateTime: 1580345760000000000,
			SortTime:   1580345760000000000,
		},
	}

	testUtf8FullComments7 = append(testUtf8FullFirstComments7, testUtf8FullTheRestComments7...)
}
