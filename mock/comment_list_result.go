package mock

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type CommentListResult_t struct {
	List    []*types.Comment `json:"list"`
	NextIdx string           `json:"next_idx"`
}

var (
	CommentListResult = &CommentListResult_t{
		List: []*types.Comment{
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("WEFSDHASD"),
				CommentID:  "alasdjfksj",
				TheType:    0,
				RefID:      "bbasdkfjs",
				IsDeleted:  false,
				CreateTime: types.Time8(1345678901),
				Owner:      "ckoool",
				Date:       "12/16",
				Content:    "怎麼了嗎？～",
				IP:         "localhost",
				Country:    "我家",
			},
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("WEFSDHASD"),
				CommentID:  "bcdasdjfksj",
				TheType:    0,
				RefID:      "bdfsdkfjs",
				IsDeleted:  false,
				CreateTime: types.Time8(1345678902),
				Owner:      "teamore",
				Date:       "12/16",
				Content:    "真的嗎？～",
				IP:         "localhost",
				Country:    "全家",
			},
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("WEFSDHASD"),
				CommentID:  "bcdDFsdjfksj",
				TheType:    0,
				RefID:      "bdfsdkfjs",
				IsDeleted:  true,
				CreateTime: types.Time8(1345688902),
				Owner:      "tealess",
				Date:       "12/17",
				Content:    "真的假的？～",
				IP:         "localhost",
				Country:    "妳家",
			},
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("WEFSDHASD"),
				CommentID:  "bcdDsadFsdjfksj",
				TheType:    0,
				RefID:      "bdfsdkfjs",
				IsDeleted:  true,
				CreateTime: types.Time8(1345700902),
				Owner:      "ok2",
				Date:       "12/18",
				Content:    "然後呢？～",
				IP:         "localhost",
				Country:    "大家",
			},
		},
		NextIdx: "10",
	}
)
