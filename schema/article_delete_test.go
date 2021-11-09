package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestDeleteArticles(t *testing.T) {
	setupTest()
	defer teardownTest()
	defer Article_c.Drop()

	articleContent := &ArticleContentInfo{
		ContentMD5:          "test1",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	_ = UpdateArticleContentInfo(bbs.BBoardID("board0"), bbs.ArticleID("article0"), articleContent)

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
			name: "test deleting articles",
			args: args{
				bbs.BBoardID("board0"),
				[]bbs.ArticleID{bbs.ArticleID("article1")},
				types.NowNanoTS(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteArticles(tt.args.boardID, tt.args.articleIDs, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("DeleteArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
