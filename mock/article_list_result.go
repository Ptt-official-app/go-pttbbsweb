package mock

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

type ArticleListResult_t struct {
	List    []*apitypes.ArticleSummary `json:"list"`
	NextIdx string                     `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

var ArticleListResult = &ArticleListResult_t{
	List: []*apitypes.ArticleSummary{
		{
			FBoardID:   apitypes.FBoardID("WhoAmI"),
			ArticleID:  apitypes.FArticleID("M.1234567890.A.123"),
			IsDeleted:  false,
			CreateTime: types.Time8(1234567890),
			MTime:      types.Time8(1234567888),
			Recommend:  8,
			Owner:      "okcool",
			Title:      "[問題]然後呢？～",
			Class:      "問題",
			Money:      5,
			Filemode:   0,
			URL:        "http://localhost/bbs/test/M.1234567890.A.123",
		},
		{
			FBoardID:   apitypes.FBoardID("WhoAmI"),
			ArticleID:  apitypes.FArticleID("M.1234567900.A.125"),
			IsDeleted:  true,
			CreateTime: types.Time8(1234567900),
			MTime:      types.Time8(1234567898),
			Recommend:  -20,
			Owner:      "somebody",
			Title:      "再來呢？～",
			Money:      123,
			Filemode:   0,
			URL:        "http://localhost/bbs/test/M.1234567900.A.125",
		},
		{
			FBoardID:   apitypes.FBoardID("WhoAmI"),
			ArticleID:  apitypes.FArticleID("M.1234568900.A.125"),
			IsDeleted:  false,
			CreateTime: types.Time8(1234568900),
			MTime:      types.Time8(1234568898),
			Recommend:  -20,
			Owner:      "somebody2",
			Title:      "還有呢？～",
			Money:      123,
			Filemode:   0,
			URL:        "http://localhost/bbs/test/M.1234568900.A.125",
		},
	},
	NextIdx: "5",
}
