package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestDeleteRanks(t *testing.T) {
	setupTest()
	defer teardownTest()
	defer Rank_c.Drop()
	// boardID: "test",
	// articleID: "test",
	// ownerID: "test",
	// rank: 1,
	// updateNanoTS: 1234567890000000000
	_, _ = UpdateRank("test", "test", "test", 1, 1234567890000000000)

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
			name: "test deleting ranks by articles",
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
			if err := DeleteRanks(tt.args.boardID, tt.args.articleIDs, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRanks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
