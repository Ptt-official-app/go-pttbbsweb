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
)
