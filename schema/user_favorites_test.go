package schema

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/fav"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestFavToUserFavorites(t *testing.T) {
	setupTest()
	defer teardownTest()

	/*
		filename1 := "./testcase/home1/t/testUser2/.fav"
		theBytes1, _ := os.ReadFile(filename1)
		buf := bytes.NewReader(theBytes1)
		_ = binary.Read(buf, binary.LittleEndian, &version)
		f1, _ := fav.ReadFavrec(buf, nil, nil, 0)
	*/

	type args struct {
		f               *fav.Fav
		userID          bbs.UUserID
		doubleBufferIdx int
		updateNanoTS    types.NanoTS
		mTime           types.NanoTS
	}
	tests := []struct {
		name                  string
		args                  args
		expectedMeta          *UserFavoritesMeta
		expectedUserFavorites []*UserFavorites
	}{
		// TODO: Add test cases.
		{
			args:                  args{f: testFav0, userID: "SYSOP", doubleBufferIdx: 0, updateNanoTS: 1234567890000000000, mTime: 1234567890000000000},
			expectedMeta:          testFavMeta0,
			expectedUserFavorites: testUserFavorites0,
		},
		{
			args:                  args{f: testFav1, userID: "SYSOP", doubleBufferIdx: 0, updateNanoTS: 1234567890000000001, mTime: 1234567890000000001},
			expectedMeta:          testFavMeta1,
			expectedUserFavorites: testUserFavorites1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMeta, gotUserFavorites := FavToUserFavorites(tt.args.f, tt.args.userID, tt.args.doubleBufferIdx, tt.args.updateNanoTS, tt.args.mTime)

			testutil.TDeepEqual(t, "meta", gotMeta, tt.expectedMeta)
			testutil.TDeepEqual(t, "favorites", gotUserFavorites, tt.expectedUserFavorites)
		})
	}
}

func TestUserFavoritesToFav(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		meta          *FolderMeta
		userFavorites []*UserFavorites
		depth         int
	}
	tests := []struct {
		name      string
		args      args
		expectedF *fav.Fav
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				meta:          &testFavMeta0.FolderMeta,
				userFavorites: testUserFavorites0,
			},
			expectedF: testFav0,
		},
		{
			args: args{
				meta:          &testFavMeta1.FolderMeta,
				userFavorites: testUserFavorites1,
			},
			expectedF: testFav1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, err := UserFavoritesToFav(tt.args.meta, tt.args.userFavorites, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserFavoritesToFav() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotF.SetFavTypeFavIdx(0)
			gotF.CleanParentAndRoot()
			testutil.TDeepEqual(t, "fav", gotF, tt.expectedF)
		})
	}
}

func TestUpdateUserFavorites(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserFavorites_c.Drop()

	type args struct {
		userID          bbs.UUserID
		doubleBufferIdx int
		userFavorites   []*UserFavorites
		mTime           types.NanoTS
		updateNanoTS    types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*UserFavorites
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{userID: "SYSOP", userFavorites: testUserFavorites0, mTime: 1234567890000000000, updateNanoTS: 1234567890000000000},
			expected: testUserFavorites0[:6],
		},
		{
			args:     args{userID: "SYSOP", userFavorites: testUserFavorites1, mTime: 1234567890000000001, updateNanoTS: 1234567890000000001},
			expected: testUserFavorites1[:6],
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			if err := UpdateUserFavorites(tt.args.userID, tt.args.doubleBufferIdx, tt.args.userFavorites, tt.args.mTime, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserFavorites() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotUserFavorites, _ := GetUserFavorites(tt.args.userID, 0, "", 0, true, 100, tt.args.mTime)

			gotUserFavorites = SortUserFavoritesByFavIdx(gotUserFavorites, true)

			testutil.TDeepEqual(t, "userFavorites", gotUserFavorites, tt.expected)
		})
		wg.Wait()
	}
}

func TestGetUserFavorites(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserFavorites_c.Drop()

	UpdateUserFavorites("SYSOP", 0, testUserFavorites1, 1234567890000000001, 1234567890000000001)

	type args struct {
		userID          bbs.UUserID
		doubleBufferIdx int
		levelIdx        LevelIdx
		startIdx        int
		ascending       bool
		limit           int
		mTime           types.NanoTS
	}
	tests := []struct {
		name                  string
		args                  args
		expectedUserFavorites []*UserFavorites
		wantErr               bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				userID:    "SYSOP",
				startIdx:  0,
				ascending: true,
				limit:     3,
				mTime:     1234567890000000001,
			},
			expectedUserFavorites: testUserFavorites1[:3],
		},
		{
			args: args{
				userID:    "SYSOP",
				startIdx:  2,
				ascending: true,
				limit:     3,
				mTime:     1234567890000000001,
			},
			expectedUserFavorites: testUserFavorites1[2:5],
		},
		{
			args: args{
				userID:    "SYSOP",
				levelIdx:  ":1",
				startIdx:  0,
				ascending: true,
				limit:     3,
				mTime:     1234567890000000001,
			},
			expectedUserFavorites: testUserFavorites1[6:7],
		},
		{
			args: args{
				userID:    "SYSOP",
				levelIdx:  ":4",
				startIdx:  0,
				ascending: true,
				limit:     3,
				mTime:     1234567890000000001,
			},
			expectedUserFavorites: testUserFavorites1[7:8],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserFavorites, err := GetUserFavorites(tt.args.userID, tt.args.doubleBufferIdx, tt.args.levelIdx, tt.args.startIdx, tt.args.ascending, tt.args.limit, tt.args.mTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFavorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotUserFavorites = SortUserFavoritesByFavIdx(gotUserFavorites, true)
			if !reflect.DeepEqual(gotUserFavorites, tt.expectedUserFavorites) {
				t.Errorf("GetUserFavorites() = %v, want %v", gotUserFavorites, tt.expectedUserFavorites)
			}
		})
	}
}
