package mock

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type CommentListResult_t struct {
	List    []*apitypes.Comment `json:"list"`
	NextIdx string              `json:"next_idx"`
}

var (
	CommentListResult = &CommentListResult_t{
		List: []*apitypes.Comment{
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("WEFSDHASD"),
				CommentID:  "alasdjfksj",
				TheType:    0,
				RefID:      "bbasdkfjs",
				IsDeleted:  false,
				CreateTime: types.Time8(1345678901),
				Owner:      "ckoool",
				Content: [][]*types.Rune{
					{
						{Utf8: "怎麼了嗎？～"},
					},
				},
				IP:   "localhost",
				Host: "我家",
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
				Content: [][]*types.Rune{
					{
						{Utf8: "真的嗎？～"},
					},
				},
				IP:   "localhost",
				Host: "全家",
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
				Content: [][]*types.Rune{
					{
						{Utf8: "真的假的？～"},
					},
				},
				IP:   "localhost",
				Host: "妳家",
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
				Content: [][]*types.Rune{
					{
						{Utf8: "然後呢？～"},
					},
				},
				IP:   "localhost",
				Host: "大家",
			},
		},
		NextIdx: "10",
	}
)
