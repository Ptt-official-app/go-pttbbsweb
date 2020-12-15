package mock_http

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadGeneralArticles(params *backend.LoadGeneralArticlesParams) (ret *backend.LoadGeneralArticlesResult) {
	ret = &backend.LoadGeneralArticlesResult{
		Articles: []*bbs.ArticleSummary{
			{
				BBoardID:   bbs.BBoardID("1_test1"),
				ArticleID:  bbs.ArticleID("1_123124"),
				IsDeleted:  false,
				Filename:   "M.1234567890.A.324",
				CreateTime: types.Time4(1234567890),
				MTime:      types.Time4(1234567889),
				Recommend:  8,
				Owner:      "okcool",
				Date:       "12/04",
				Title:      "[問題]然後呢？～",
				Money:      3,
				Filemode:   0,
				URL:        "http://localhost/bbs/test1/M.1234567890.A.324.html",
				Read:       false,
			},
			{
				BBoardID:   bbs.BBoardID("1_test1"),
				ArticleID:  bbs.ArticleID("2_123125"),
				IsDeleted:  false,
				Filename:   "M.1234567890.A.325",
				CreateTime: types.Time4(1234567900),
				MTime:      types.Time4(1234567890),
				Recommend:  3,
				Owner:      "teemo",
				Date:       "12/05",
				Title:      "[問題]再來呢？～",
				Money:      12,
				Filemode:   0,
				URL:        "http://localhost/bbs/test1/M.1234567890.A.325.html",
				Read:       false,
			},
		},
		IsNewest: true,
		NextIdx:  "testNextIdx",
	}

	return ret
}
