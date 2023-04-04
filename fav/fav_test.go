package fav

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"reflect"
	"sync"
	"testing"

	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFavRaw_AddBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0, _ := NewFav(nil, nil, 0)
	logrus.Infof("TestAddBoard: f0: %v", f0)
	ft0 := &FavType{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}}

	f1 := &Fav{
		FavNum:  1,
		NBoards: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}
	f1.Root = f1

	f1_1 := &Fav{
		FavNum:  1,
		NBoards: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}
	f1_1.Root = f1_1
	ft1 := &FavType{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{2, 0, 0}}

	f1_2 := &Fav{
		FavNum:  1,
		NBoards: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}
	f1_2.Root = f1_2

	f2 := &Fav{
		FavNum:  2,
		NBoards: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
			{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 2}},
		},
	}
	f2.Root = f1_2

	f3, _ := NewFav(nil, nil, 0)
	_, _ = f3.AddBoard(1)
	favType3, _ := f3.AddFolder()
	favFolder3 := favType3.Fp.(*FavFolder)

	ft3 := &FavType{2, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}}

	sub3 := &Fav{
		NBoards: 1,
		FavNum:  1,
		Depth:   1,
		Favh: []*FavType{
			{2, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}

	expected3 := &Fav{
		FavNum:   3,
		NBoards:  1,
		NFolders: 1,
		FolderID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
			{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: sub3}},
		},
	}

	filename3 := "./testcase/home1/t/testUser3/.fav"

	theBytes3, _ := ioutil.ReadFile(filename3)

	buf3 := bytes.NewReader(theBytes3)

	filename4 := "./testcase/home1/t/testUser4/.fav"

	theBytes4, _ := ioutil.ReadFile(filename4)

	buf4 := bytes.NewReader(theBytes4)

	// version
	version := int16(0)
	_ = binary.Read(buf3, binary.LittleEndian, &version)

	_ = binary.Read(buf4, binary.LittleEndian, &version)

	fav3, _ := ReadFavrec(buf3, nil, nil, 0)

	fav3Ft10 := &FavType{10, pttbbsfav.FAVT_BOARD, 1, &FavBoard{23, 0, 0}}

	fav3.AddBoard(21)

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name            string
		f               *Fav
		args            args
		expectedFavType *FavType
		cmp             *Fav
		expected        *Fav
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			name:            "add bid = 1",
			f:               f0,
			args:            args{bid: 1},
			expectedFavType: ft0,
			cmp:             f0,
			expected:        f1,
		},
		{
			name:            "add bid = 1 (again)",
			f:               f1_1,
			args:            args{bid: 1},
			expectedFavType: ft0,
			cmp:             f1_1,
			expected:        f1_1,
		},
		{
			name:            "add bid = 2",
			f:               f1_2,
			args:            args{bid: 2},
			expectedFavType: ft1,
			cmp:             f1_2,
			expected:        f2,
		},
		{
			name:            "add bid = 1 (sub-folder)",
			f:               favFolder3.ThisFolder,
			args:            args{bid: 1},
			expectedFavType: ft3,
			cmp:             f3,
			expected:        expected3,
		},

		{
			name:            "add bid = 23 (after setting up sub-folder)",
			f:               fav3,
			args:            args{bid: 23},
			expectedFavType: fav3Ft10,
			cmp:             fav3,
			expected:        testFav5,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			f := tt.f
			gotFavType, err := f.AddBoard(tt.args.bid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fav.AddBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.cmp.SetFavTypeFavIdx(0)

			if !reflect.DeepEqual(gotFavType, tt.expectedFavType) {
				t.Errorf("Fav.AddBoard() = %v, want %v", gotFavType, tt.expectedFavType)
			}

			tt.cmp.CleanParentAndRoot()
			tt.expected.CleanParentAndRoot()

			testutil.TDeepEqual(t, "expected", tt.cmp, tt.expected)
		})
		wg.Wait()
	}
}

func TestFavRaw_AddLine(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0, _ := NewFav(nil, nil, 0)
	ft0 := &FavType{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}}

	expected0 := &Fav{
		FavNum: 1,
		NLines: 1,
		LineID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
		},
	}

	f1 := &Fav{
		FavNum: 1,
		NLines: 1,
		LineID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
		},
	}
	f1.Root = f1

	ft1 := &FavType{1, pttbbsfav.FAVT_LINE, 1, &FavLine{2}}

	expected1 := &Fav{
		Parent: f1,
		FavNum: 2,
		NLines: 2,
		LineID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{1, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
		},
	}

	tests := []struct {
		name            string
		f               *Fav
		expectedFavType *FavType
		expected        *Fav
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			f:               f0,
			expectedFavType: ft0,
			expected:        expected0,
		},
		{
			name:            "add 2nd line",
			f:               f1,
			expectedFavType: ft1,
			expected:        expected1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.f
			gotFavType, err := f.AddLine()
			if (err != nil) != tt.wantErr {
				t.Errorf("FavRaw.AddLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			f.SetFavTypeFavIdx(0)

			f.CleanParentAndRoot()
			tt.expected.CleanParentAndRoot()

			assert.Equal(t, gotFavType, tt.expectedFavType)

			testutil.TDeepEqual(t, "expected", f, tt.expected)
		})
	}
}

func TestFavRaw_AddFolder(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0, _ := NewFav(nil, nil, 0)

	folder0, _ := NewFav(nil, nil, 1)
	ft0 := &FavType{0, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: folder0}}
	expected0 := &Fav{
		FavNum:   1,
		NFolders: 1,
		FolderID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: folder0}},
		},
	}

	f1, _ := NewFav(nil, nil, 0)
	_, _ = f1.AddFolder()

	folder1, _ := NewFav(nil, nil, 1)
	ft1 := &FavType{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 2, ThisFolder: folder1}}
	expected1 := &Fav{
		FavNum:   2,
		NFolders: 2,
		FolderID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: folder1}},
			{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 2, ThisFolder: folder1}},
		},
	}

	tests := []struct {
		name            string
		f               *Fav
		expectedFavType *FavType
		expected        *Fav
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			f:               f0,
			expectedFavType: ft0,
			expected:        expected0,
		},
		{
			f:               f1,
			expectedFavType: ft1,
			expected:        expected1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.f
			gotFavType, err := f.AddFolder()
			if (err != nil) != tt.wantErr {
				t.Errorf("FavRaw.AddFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			f.SetFavTypeFavIdx(0)
			f.CleanParentAndRoot()
			tt.expected.CleanParentAndRoot()

			gotFolder := gotFavType.Fp.(*FavFolder)
			gotFolder.ThisFolder.CleanParentAndRoot()

			expectedFolder := tt.expectedFavType.Fp.(*FavFolder)
			expectedFolder.ThisFolder.CleanParentAndRoot()

			assert.Equal(t, gotFavType, tt.expectedFavType)

			testutil.TDeepEqual(t, "expected", f, tt.expected)
		})
	}
}

func TestReadFavrec(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename0 := "./testcase/home1/t/testUser/.fav"

	theBytes0, _ := ioutil.ReadFile(filename0)

	buf0 := bytes.NewReader(theBytes0)

	filename1 := "./testcase/home1/t/testUser2/.fav"

	theBytes1, _ := ioutil.ReadFile(filename1)

	buf1 := bytes.NewReader(theBytes1)

	filename3 := "./testcase/home1/t/testUser3/.fav"

	theBytes3, _ := ioutil.ReadFile(filename3)

	buf3 := bytes.NewReader(theBytes3)

	filename4 := "./testcase/home1/t/testUser4/.fav"

	theBytes4, _ := ioutil.ReadFile(filename4)

	buf4 := bytes.NewReader(theBytes4)

	// version
	version := int16(0)
	_ = binary.Read(buf0, binary.LittleEndian, &version)

	_ = binary.Read(buf1, binary.LittleEndian, &version)

	_ = binary.Read(buf3, binary.LittleEndian, &version)

	_ = binary.Read(buf4, binary.LittleEndian, &version)

	type args struct {
		file io.ReadSeeker
	}
	tests := []struct {
		name           string
		args           args
		expectedFavrec *Fav
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{file: buf0},
			expectedFavrec: testFav0,
		},
		{
			args:           args{file: buf1},
			expectedFavrec: testFav1,
		},
		{
			args:           args{file: buf3},
			expectedFavrec: testFav3,
		},
		{
			args:           args{file: buf4},
			expectedFavrec: testFav4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFavrec, err := ReadFavrec(tt.args.file, nil, nil, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFavrec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotFavrec.SetFavTypeFavIdx(0)
			gotFavrec.CleanParentAndRoot()

			testutil.TDeepEqual(t, "got", gotFavrec, tt.expectedFavrec)
		})
	}
}

func TestFav_WriteFavrec(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename := "./testcase/home1/t/testUser/.fav"

	theBytes0, _ := ioutil.ReadFile(filename)

	filename1 := "./testcase/home1/t/testUser2/.fav"

	theBytes1, _ := ioutil.ReadFile(filename1)

	tests := []struct {
		name         string
		favrec       *Fav
		expectedFile []byte
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			favrec:       testFav0,
			expectedFile: theBytes0,
		},
		{
			favrec:       testFav1,
			expectedFile: theBytes1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.favrec
			file := &bytes.Buffer{}
			if err := f.WriteFavrec(file); (err != nil) != tt.wantErr {
				t.Errorf("Fav.WriteFavrec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotFile := file.Bytes()

			testutil.TDeepEqual(t, "got", gotFile, tt.expectedFile)
		})
	}
}
