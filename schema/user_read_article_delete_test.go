package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestDeleteUserReadArticles(t *testing.T) {
	setupTest()
	defer teardownTest()
	defer UserReadArticle_c.Drop()

	userReadArticles := []*UserReadArticle{
		{
			UserID:       bbs.UUserID("testuser0"),
			BoardID:      bbs.BBoardID("test"),
			ArticleID:    bbs.ArticleID("test"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	_ = UpdateUserReadArticles(userReadArticles, types.NanoTS(1234567890000000000))
	type args struct {
		boardID      bbs.BBoardID
		articleIDs   []bbs.ArticleID
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test deleting user_read_articles",
			args: args{
				bbs.BBoardID("test"),
				[]bbs.ArticleID{bbs.ArticleID("test")},
				types.NowNanoTS(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUserReadArticles(tt.args.boardID, tt.args.articleIDs, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserReadArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
