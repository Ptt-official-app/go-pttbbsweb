package fav

import pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"

/*
test0:
[35, 13,
 3, 0, 2, 1,
 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 3, 1, 1,
 2, 1, 1, 183, 115, 170, 186, 165, 216, 191, 253, 0, 0,
 	   0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 	   0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 	   0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
       0, 0, 0, 0, 0, 0, 0, 0, 0,
 3, 1, 2,
 1, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 1, 1, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

 1, 0, 0, 0,
 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
]
*/

var (
	//nolint:unused
	testTitle0 = "新的目錄"
	//nolint:unused
	testSubFav0 = &Fav{
		FavNum:   1,
		Depth:    1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*FavType{
			{6, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
		},
	}

	testFav0 = &Fav{ //nolint // possible use in the future
		NBoards:  3,
		NLines:   2,
		NFolders: 1,
		LineID:   2,
		FolderID: 1,
		FavNum:   7,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
			{1, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{2, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle0, testSubFav0}},
			{3, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
			{4, pttbbsfav.FAVT_BOARD, 1, &FavBoard{9, 0, 0}},
			{5, pttbbsfav.FAVT_BOARD, 1, &FavBoard{8, 0, 0}},
		},
	}

	//nolint:unused
	testSubFav1 = &Fav{
		FavNum:   1,
		Depth:    1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*FavType{
			{6, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
		},
	}

	//nolint:unused
	testSubFav2 = &Fav{
		FavNum:   1,
		Depth:    1,
		NBoards:  1,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Favh: []*FavType{
			{7, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
		},
	}

	testFav1 = &Fav{ //nolint // possible use in the future.
		FavNum:   8,
		NBoards:  2,
		NLines:   2,
		NFolders: 2,
		LineID:   2,
		FolderID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle0, testSubFav1}},
			{2, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
			{3, pttbbsfav.FAVT_BOARD, 1, &FavBoard{9, 0, 0}},
			{4, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{2, testTitle0, testSubFav2}},
			{5, pttbbsfav.FAVT_BOARD, 1, &FavBoard{8, 0, 0}},
		},
	}
)
