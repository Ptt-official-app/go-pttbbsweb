package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadGeneralArticles(params *api.LoadGeneralArticlesParams) (ret *api.LoadGeneralArticlesResult) {
	ret = &api.LoadGeneralArticlesResult{
		Articles: []*bbs.ArticleSummary{
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("1VrooM21"),
				IsDeleted:  false,
				Filename:   "M.1607937174.A.081",
				CreateTime: types.Time4(1607937174),
				MTime:      types.Time4(1607937100),
				Recommend:  3,
				Owner:      bbs.UUserID("teemo"),
				Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
				Title:      []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]再來呢？～
				Money:      12,
				Filemode:   0,
				Read:       false,
			},
			{
				BBoardID:   bbs.BBoardID("10_WhoAmI"),
				ArticleID:  bbs.ArticleID("19bWBI4Z"),
				IsDeleted:  false,
				Filename:   "M.1234567890.A.123",
				CreateTime: types.Time4(1234567890),
				MTime:      types.Time4(1234567889),
				Recommend:  8,
				Owner:      bbs.UUserID("okcool"),
				Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
				Title:      []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]然後呢？～
				Money:      3,
				Filemode:   0,
				Read:       false,
			},
		},
		IsNewest:       true,
		NextIdx:        "1234560000@19bUG021",
		NextCreateTime: 1234560000,
	}

	return ret
}
