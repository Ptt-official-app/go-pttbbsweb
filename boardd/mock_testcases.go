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
		Filename:      "M.1607937174.A.081",
		RawDate:       "2020/12/14 17:12:54",
		NumRecommends: 3,
		Owner:         "teemo",
		Title:         "[問題]再來呢？～",
		ModifiedNsec:  1607937100000000000,
	}

	testArticle1 = &Post{
		Index:         1,
		Filename:      "M.1234567890.A.123",
		RawDate:       "2009/02/14 07:30:00",
		NumRecommends: 8,
		Owner:         "okcool",
		Title:         "[問題]然後呢？～",
		ModifiedNsec:  1234567889000000000,
	}

	testArticle2 = &Post{
		Index:         2,
		Filename:      "M.1234560000.A.081",
		RawDate:       "2009/02/14 05:20:00",
		NumRecommends: 13,
		Owner:         "SYSOP",
		Title:         "[問題]這是 SYSOP",
		ModifiedNsec:  1234560000000000000,
	}

	testArticle3 = &Post{
		Index:         2,
		Filename:      "M.1607937176.A.081",
		RawDate:       "2020/12/14 17:12:54",
		NumRecommends: 13,
		Owner:         "teemo",
		Title:         "[問題]再來呢？～",
		ModifiedNsec:  1607937100000000000,
	}

	testArticle4 = &Post{
		Index:         2,
		Filename:      "M.1234567892.A.123",
		RawDate:       "2009/02/14 07:30:02",
		NumRecommends: 24,
		Owner:         "teemo",
		Title:         "[問題]然後呢？～",
		ModifiedNsec:  1234567889000000000,
	}
)
