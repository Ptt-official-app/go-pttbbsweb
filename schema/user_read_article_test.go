package schema

import (
	"sort"
	"strings"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateUserReadArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserReadArticle_c.Drop()

	userReadArticles := []*UserReadArticle{
		{
			UserID:       bbs.UUserID("testuser0"),
			BoardID:      bbs.BBoardID("10_WhoAmI"),
			ArticleID:    bbs.ArticleID("testarticle0"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BoardID:      bbs.BBoardID("10_WhoAmI"),
			ArticleID:    bbs.ArticleID("testarticle1"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BoardID:      bbs.BBoardID("10_WhoAmI"),
			ArticleID:    bbs.ArticleID("testarticle2"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	type args struct {
		userReadArticles []*UserReadArticle
		updateNanoTS     types.NanoTS
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		findUserID     bbs.UUserID
		findBoardID    bbs.BBoardID
		findArticleIDs []bbs.ArticleID
		expected       []*UserReadArticle
	}{
		// TODO: Add test cases.
		{
			args:           args{userReadArticles: userReadArticles, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:     bbs.UUserID("testuser0"),
			findBoardID:    bbs.BBoardID("10_WhoAmI"),
			findArticleIDs: []bbs.ArticleID{"testarticle0", "testarticle1", "testarticle2"},
			expected:       userReadArticles,
		},
		{
			args:           args{userReadArticles: userReadArticles, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:     bbs.UUserID("testuser0"),
			findBoardID:    bbs.BBoardID("10_WhoAmI"),
			findArticleIDs: []bbs.ArticleID{"testarticle0", "testarticle1", "testarticle2"},
			expected:       userReadArticles,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadArticles(tt.args.userReadArticles, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadArticles() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := FindUserReadArticles(tt.findUserID, tt.findBoardID, tt.findArticleIDs)
			if err != nil {
				t.Errorf("UpdateUserReadArticles: e: %v", err)
			}

			sort.SliceStable(got, func(i, j int) bool {
				return strings.Compare(string(got[i].ArticleID), string(got[j].ArticleID)) <= 0
			})
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}

func TestUpdateUserReadArticle(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserReadArticle_c.Drop()

	userReadArticle := &UserReadArticle{
		UserID:       bbs.UUserID("testuser0"),
		BoardID:      bbs.BBoardID("10_WhoAmI"),
		ArticleID:    bbs.ArticleID("testarticle0"),
		UpdateNanoTS: types.NanoTS(1234567890000000000),
	}
	type args struct {
		userReadArticle *UserReadArticle
	}
	tests := []struct {
		name     string
		args     args
		expected *UserReadArticle
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{userReadArticle: userReadArticle},
			expected: userReadArticle,
		},
		{
			args:     args{userReadArticle: userReadArticle},
			expected: userReadArticle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadArticle(tt.args.userReadArticle); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadArticle() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := &UserReadArticle{}

			query := bson.M{
				USER_READ_ARTICLE_USER_ID_b:    tt.args.userReadArticle.UserID,
				USER_READ_ARTICLE_BOARD_ID_b:   tt.args.userReadArticle.BoardID,
				USER_READ_ARTICLE_ARTICLE_ID_b: tt.args.userReadArticle.ArticleID,
			}
			_ = UserReadArticle_c.FindOne(query, got, nil)

			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}
