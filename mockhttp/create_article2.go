package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func CreateArticle2(params *api.CreateArticleParams) (ret api.CreateArticleResult) {
	ret = api.CreateArticleResult(&bbs.ArticleSummary{
		BBoardID:   bbs.BBoardID("1_SYSOP"),
		ArticleID:  bbs.ArticleID("1VrooM23"),
		IsDeleted:  false,
		Filename:   "M.1607937174.A.083",
		CreateTime: types.Time4(1607937174),
		MTime:      types.Time4(1607937100),
		Recommend:  0,
		Owner:      bbs.UUserID("SYSOP"),
		Class:      []byte{0xb4, 0xfa, 0xb8, 0xd5}, // 測試
		FullTitle:  []byte("[\xb4\xfa\xb8\xd5]this is a test"),
		Money:      0,
		Filemode:   0,
		Read:       false,
		RealTitle:  []byte("this is a test"),
	})

	return ret
}
