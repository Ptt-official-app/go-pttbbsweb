package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/sirupsen/logrus"
)

func LoadGeneralArticles3(params *api.LoadGeneralArticlesParams) (ret *api.LoadGeneralArticlesResult) {
	logrus.Infof("mockhttp.LoadGeneralArticles3: params: %v", params)
	if params.StartIdx == "" {
		return &api.LoadGeneralArticlesResult{
			Articles: []*bbs.ArticleSummary{
				{
					BBoardID:   bbs.BBoardID("1_SYSOP"),
					ArticleID:  bbs.ArticleID("1VrooO21"),
					IsDeleted:  false,
					Filename:   "M.1607937176.A.081",
					CreateTime: types.Time4(1607937176),
					MTime:      types.Time4(1607937100),
					Recommend:  10,
					Owner:      bbs.UUserID("teemo"),
					Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
					FullTitle:  []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]再來呢？～
					Money:      12,
					Filemode:   0,
					Read:       false,
					Idx:        "1607937176@1VrooO21",
					RealTitle:  []byte{0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3},
				},
				{
					BBoardID:   bbs.BBoardID("1_SYSOP"),
					ArticleID:  bbs.ArticleID("19bWBK4Z"),
					IsDeleted:  false,
					Filename:   "M.1234567892.A.123",
					CreateTime: types.Time4(1234567892),
					MTime:      types.Time4(1234567889),
					Recommend:  24,
					Owner:      bbs.UUserID("teemo"),
					Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
					FullTitle:  []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]然後呢？～
					Money:      3,
					Filemode:   0,
					Read:       false,
					Idx:        "1234567892@19bWBK4Z",
					RealTitle:  []byte{0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3},
				},
			},
			IsNewest:       true,
			NextIdx:        "1234560000@19bUG021",
			NextCreateTime: 1234560000,
		}
	}

	var articleSummary *bbs.ArticleSummary

	if params.StartIdx == "1607937174@1VrooM21" {
		articleSummary = &bbs.ArticleSummary{
			BBoardID:   bbs.BBoardID("1_SYSOP"),
			ArticleID:  bbs.ArticleID("1VrooM21"),
			IsDeleted:  false,
			Filename:   "M.1607937174.A.081",
			CreateTime: types.Time4(1607937174),
			MTime:      types.Time4(1607937100),
			Recommend:  3,
			Owner:      bbs.UUserID("SYSOP"),
			Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
			FullTitle:  []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]再來呢？～
			Money:      12,
			Filemode:   0,
			Read:       false,
			Idx:        "1607937174@1VrooM21",
			RealTitle:  []byte{0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3},
		}
	} else if params.StartIdx == "1608386280@1VtWRel9" {
		articleSummary = &bbs.ArticleSummary{
			BBoardID:   bbs.BBoardID("1_SYSOP"),
			ArticleID:  bbs.ArticleID("1VtWRel9"),
			IsDeleted:  false,
			Filename:   "M.1234567890.A.123",
			CreateTime: types.Time4(1608386280),
			MTime:      types.Time4(1608386280),
			Recommend:  8,
			Owner:      bbs.UUserID("SYSOP"),
			Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
			FullTitle:  []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]然後呢？～
			Money:      3,
			Filemode:   0,
			Read:       false,
			Idx:        "1234567890@1VtWRel9",
			RealTitle:  []byte{0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3},
		}
	}
	ret = &api.LoadGeneralArticlesResult{
		Articles: []*bbs.ArticleSummary{
			articleSummary,
		},
		IsNewest:       true,
		NextIdx:        "1234560000@19bUG021",
		NextCreateTime: 1234560000,
	}

	return ret
}
