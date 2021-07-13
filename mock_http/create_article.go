package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func CreateArticle(params *api.CreateArticleParams) (ret api.CreateArticleResult) {
	ret = api.CreateArticleResult(&bbs.ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("1VrooM21"),
		IsDeleted:  false,
		Filename:   "M.1607937174.A.081",
		CreateTime: types.Time4(1607937174),
		MTime:      types.Time4(1607937100),
		Recommend:  0,
		Owner:      bbs.UUserID("SYSOP"),
		Class:      []byte{0xb4, 0xfa, 0xb8, 0xd5}, // 測試
		Title:      []byte("[\xb4\xfa\xb8\xd5]this is a test"),
		Money:      0,
		Filemode:   0,
		Read:       false,
	})

	return ret
}
