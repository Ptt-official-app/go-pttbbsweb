package api

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

func Test_bidToBoardID(t *testing.T) {
	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name            string
		args            args
		expectedBoardID bbs.BBoardID
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoardID, err := bidToBoardID(tt.args.bid)
			if (err != nil) != tt.wantErr {
				t.Errorf("bidToBoardID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoardID, tt.expectedBoardID) {
				t.Errorf("bidToBoardID() = %v, want %v", gotBoardID, tt.expectedBoardID)
			}
		})
	}
}
