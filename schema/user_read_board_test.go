package schema

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateUserReadBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserReadBoard_c.Drop()

	userReadBoard := &UserReadBoard{
		UserID:       bbs.UUserID("testuser0"),
		BBoardID:     bbs.BBoardID("testboard0"),
		UpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	type args struct {
		userReadBoard *UserReadBoard
	}
	tests := []struct {
		name     string
		args     args
		expected *UserReadBoard
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{userReadBoard},
			expected: userReadBoard,
		},
		{
			args:     args{userReadBoard},
			expected: userReadBoard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadBoard(tt.args.userReadBoard); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadBoard() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := &UserReadBoard{}

			query := bson.M{
				USER_READ_BOARD_USER_ID_b:   tt.args.userReadBoard.UserID,
				USER_READ_BOARD_BBOARD_ID_b: tt.args.userReadBoard.BBoardID,
			}

			_ = UserReadBoard_c.FindOne(query, got, nil)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}

func TestUpdateUserReadBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserReadBoard_c.Drop()

	userReadBoards := []*UserReadBoard{
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard0"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard1"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard2"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}
	userReadBoards1 := []*UserReadBoard{
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard2"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard3"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	type args struct {
		userReadBoards []*UserReadBoard
		updateNanoTS   types.NanoTS
	}
	tests := []struct {
		name         string
		args         args
		findUserID   bbs.UUserID
		findBoardIDs []bbs.BBoardID
		expected     []*UserReadBoard
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			args:         args{userReadBoards: userReadBoards, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:   bbs.UUserID("testuser0"),
			findBoardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			expected:     userReadBoards,
		},
		{
			args:         args{userReadBoards: userReadBoards, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:   bbs.UUserID("testuser0"),
			findBoardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			expected:     userReadBoards,
		},
		{
			args:         args{userReadBoards: userReadBoards1, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:   bbs.UUserID("testuser0"),
			findBoardIDs: []bbs.BBoardID{"testboard2", "testboard3"},
			expected:     userReadBoards1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadBoards(tt.args.userReadBoards, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadBoards() error = %v, wantErr %v", err, tt.wantErr)
				got, err := FindUserReadBoards(tt.findUserID, tt.findBoardIDs)
				if err != nil {
					t.Errorf("UpdateUserReadBoards: e: %v", err)
				}

				sort.SliceStable(got, func(i, j int) bool {
					return strings.Compare(string(got[i].BBoardID), string(got[j].BBoardID)) <= 0
				})
				testutil.TDeepEqual(t, "got", got, tt.expected)

				t.Errorf("false")

			}
		})
	}
}

func TestFindUserReadBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserReadBoard_c.Drop()

	userReadBoards := []*UserReadBoard{
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard0"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard1"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:       bbs.UUserID("testuser0"),
			BBoardID:     bbs.BBoardID("testboard2"),
			UpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	_ = UpdateUserReadBoards(userReadBoards, types.NanoTS(1234567890000000000))

	type args struct {
		userID   bbs.UUserID
		boardIDs []bbs.BBoardID
	}
	tests := []struct {
		name     string
		args     args
		expected []*UserReadBoard
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				userID:   bbs.UUserID("testuser0"),
				boardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			},
			expected: userReadBoards,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserReadBoards(tt.args.userID, tt.args.boardIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserReadBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindUserReadBoards() = %v, want %v", got, tt.expected)
			}

			sort.SliceStable(got, func(i, j int) bool {
				return strings.Compare(string(got[i].BBoardID), string(got[j].BBoardID)) <= 0
			})
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}
