package boardd

import "github.com/Ptt-official-app/go-pttbbs/ptttype"

var (
	testBoard6 = &Board{ // nolint
		Parent:        5,
		Bid:           6,
		Attributes:    uint32(ptttype.BRD_POSTMASK),
		Name:          "ALLPOST",
		Title:         "跨板式LOCAL新文章",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testBoard7 = &Board{ // nolint
		Parent:        5,
		Bid:           7,
		Name:          "deleted",
		Title:         "資源回收筒",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testBoard11 = &Board{ // nolint
		Parent:        5,
		Bid:           11,
		Name:          "EditExp",
		Title:         "範本精靈投稿區",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testBoard8 = &Board{
		Parent:        5,
		Bid:           8,
		Name:          "Note",
		Title:         "動態看板及歌曲投稿",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testBoard1 = &Board{
		Parent:        2,
		Bid:           1,
		Attributes:    uint32(ptttype.BRD_POSTMASK),
		Name:          "SYSOP",
		Title:         "站長好!",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testBoard9 = &Board{ // nolint
		Parent:        5,
		Bid:           9,
		Attributes:    uint32(ptttype.BRD_POSTMASK),
		Name:          "Record",
		Title:         "我們的成果",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testBoard10 = &Board{
		Parent:        5,
		Bid:           10,
		Name:          "WhoAmI",
		Title:         "呵呵，猜猜我是誰！",
		Bclass:        "嘰哩",
		RawModerators: "",
	}

	testArticle0 = &Post{
		Index:         0,
		Filename:      "M.1584665022.A.ED0",
		RawDate:       "2020/03/20 09:43:42",
		NumRecommends: 17,
		Owner:         "hellohiro",
		Title:         "[問卦] 為何打麻將叫賭博但買股票叫投資？",
		ModifiedNsec:  1644506386000000000,
	}

	testArticle1 = &Post{
		Index:         1,
		Filename:      "M.1608386280.A.BC9",
		RawDate:       "2020/12/19 21:58:00",
		NumRecommends: 9,
		Owner:         "SYSOP",
		Title:         "[心得] 測試一下特殊字～",
		ModifiedNsec:  1608386280000000000,
	}

	testArticle2 = &Post{
		Index:         2,
		Filename:      "M.1608388506.A.85D",
		RawDate:       "2020/12/19 22:35:06",
		NumRecommends: 9,
		Owner:         "SYSOP",
		Title:         "[閒聊] 所以特殊字真的是有綠色的∼",
		ModifiedNsec:  1608386280000000000,
	}

	testArticle3 = &Post{
		Index:         3,
		Filename:      "M.1607202240.A.30D",
		RawDate:       "2020/12/06 05:04:00",
		NumRecommends: 23,
		Owner:         "cheinshin",
		Title:         "[新聞] TVBS六都民調 侯奪冠、盧升第四、柯墊底",
		ModifiedNsec:  1607202240000000000,
	}

	testArticle4 = &Post{
		Index:         4,
		Filename:      "M.1607937174.A.081",
		RawDate:       "2020/12/14 17:12:54",
		NumRecommends: 3,
		Owner:         "teemo",
		Title:         "[閒聊] 新書的情報",
		ModifiedNsec:  1607937100000000000,
	}

	testArticle5 = &Post{
		Index:         5,
		Filename:      "M.1234567890.A.123",
		RawDate:       "2009/02/14 07:30:00",
		NumRecommends: 8,
		Owner:         "okcool",
		Title:         "[問題]然後呢？～",
		ModifiedNsec:  1234567889000000000,
	}

	testArticle6 = &Post{
		Index:         6,
		Filename:      "M.1234560000.A.081",
		RawDate:       "2009/02/14 05:20:00",
		NumRecommends: 13,
		Owner:         "SYSOP",
		Title:         "[問題]這是 SYSOP",
		ModifiedNsec:  1234560000000000000,
	}

	testArticle7 = &Post{
		Index:         7,
		Filename:      "M.1607937176.A.081",
		RawDate:       "2020/12/14 17:12:54",
		NumRecommends: 13,
		Owner:         "SYSOP",
		Title:         "[心得] 測試一下特殊字～",
		ModifiedNsec:  1607937100000000000,
	}

	testArticle8 = &Post{
		Index:         8,
		Filename:      "M.1234567892.A.123",
		RawDate:       "2009/02/14 07:30:02",
		NumRecommends: 24,
		Owner:         "teemo",
		Title:         "[問題]然後呢？～",
		ModifiedNsec:  1234567889000000000,
	}
)
