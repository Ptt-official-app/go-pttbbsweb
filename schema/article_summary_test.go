package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mockhttp"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

	ret := mockhttp.LoadGeneralArticles(nil)

	updateNanoTS := types.NowNanoTS() - 200

	articleSummaries0 := make([]*ArticleSummary, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaries0[idx] = NewArticleSummary(each_b, updateNanoTS)
	}

	query0 := &ArticleQuery{BBoardID: bbs.BBoardID("10_WhoAmI"), ArticleID: bbs.ArticleID("19bWBI4Z")}
	articleSummary0 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("19bWBI4Z"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		FullTitle:    "[問題]然後呢？～",
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1234567890@19bWBI4Z",
	}

	articleSummary1 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("1VrooM21"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    3,
		Owner:        bbs.UUserID("teemo"),
		FullTitle:    "[問題]再來呢？～",
		Title:        "再來呢？～",
		Money:        12,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1607937174@1VrooM21",
	}

	updateNanoTS = types.NowNanoTS() - 100

	articleSummary2 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("1VrooM21"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    4,
		Owner:        bbs.UUserID("teem2"),
		FullTitle:    "[問題]再來呢2？～",
		Title:        "再來呢2？～",
		Money:        15,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1607937174@1VrooM21",
	}

	updateNanoTS1 := types.NowNanoTS()

	articleSummary3 := &ArticleSummary{
		BBoardID:     bbs.BBoardID("10_WhoAmI"),
		ArticleID:    bbs.ArticleID("1VrooM21"),
		IsDeleted:    false,
		CreateTime:   types.NanoTS(1607937174000000000),
		MTime:        types.NanoTS(1607937100000000000),
		FullTitle:    "[問題]再來呢3？～",
		Recommend:    12,
		Owner:        bbs.UUserID("teemo"),
		Title:        "再來呢3？～",
		Money:        20,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS1,
		Idx:          "1607937174@1VrooM21",
	}

	articleSummaries1 := []*ArticleSummary{articleSummary2}

	query1 := &ArticleQuery{BBoardID: bbs.BBoardID("10_WhoAmI"), ArticleID: bbs.ArticleID("1VrooM21")}

	query2 := &ArticleQuery{BBoardID: bbs.BBoardID("10_WhoAmI"), ArticleID: bbs.ArticleID("1VrooM21")}

	articleSummaries2 := []*ArticleSummary{articleSummary3}

	type args struct {
		articleSummaries []*ArticleSummary
		updateNanoTS     types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		query    *ArticleQuery
		expected *ArticleSummary
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{articleSummaries: articleSummaries0, updateNanoTS: updateNanoTS},
			query:    query0,
			expected: articleSummary0,
		},
		{
			args:     args{articleSummaries: articleSummaries0, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: articleSummary1,
		},
		{
			args:     args{articleSummaries: articleSummaries1, updateNanoTS: updateNanoTS},
			query:    query2,
			expected: articleSummary2,
		},
		{
			args:     args{articleSummaries: articleSummaries2, updateNanoTS: updateNanoTS1},
			query:    query1,
			expected: articleSummary3,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateArticleSummaries(tt.args.articleSummaries, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := GetArticleSummary(tt.query.BBoardID, tt.query.ArticleID)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
