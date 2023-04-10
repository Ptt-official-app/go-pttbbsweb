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

	testTitle301 = "籃球類"

	testSubFav301 = &Fav{
		FavNum:   3,
		NBoards:  3,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Depth:    2,
		Favh: []*FavType{
			{15, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13830, 0, 0}},
			{16, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13650, 0, 0}},
			{17, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14332, 0, 0}},
		},
	}

	testSubFav30 = &Fav{
		FavNum:   9,
		NBoards:  4,
		NLines:   1,
		NFolders: 1,
		LineID:   1,
		FolderID: 1,
		Depth:    1,
		Favh: []*FavType{
			{9, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14332, 0, 0}},
			{10, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14326, 0, 0}},
			{11, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{12, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13703, 0, 0}},
			{13, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle301, testSubFav301}},
			{14, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13650, 0, 0}},
		},
	}

	testSubFav31 = &Fav{
		FavNum:   3,
		NBoards:  3,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Depth:    1,
		Favh: []*FavType{
			{18, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13810, 0, 0}},
			{19, pttbbsfav.FAVT_BOARD, 1, &FavBoard{10, 0, 0}},
			{20, pttbbsfav.FAVT_BOARD, 1, &FavBoard{326, 0, 0}},
		},
	}

	testTitle30 = "運動類"
	testTitle31 = "這是我的新的目錄"

	testFav3 = &Fav{
		FavNum:   21,
		NBoards:  5,
		NLines:   2,
		NFolders: 2,
		LineID:   2,
		FolderID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{10, 0, 0}},
			{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{2569, 0, 0}},
			{2, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{3, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14468, 0, 0}},
			{4, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
			{5, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
			{6, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle30, testSubFav30}},
			{7, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{2, testTitle31, testSubFav31}},
			{8, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13703, 0, 0}},
		},
	}

	testSubFav401 = &Fav{
		FavNum:   3,
		NBoards:  3,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Depth:    2,
		Favh: []*FavType{
			{17, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13830, 0, 0}},
			{18, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13650, 0, 0}},
			{19, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14332, 0, 0}},
		},
	}

	testSubFav40 = &Fav{
		FavNum:   9,
		NBoards:  4,
		NLines:   1,
		NFolders: 1,
		LineID:   1,
		FolderID: 1,
		Depth:    1,
		Favh: []*FavType{
			{11, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14332, 0, 0}},
			{12, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14326, 0, 0}},
			{13, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{14, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13703, 0, 0}},
			{15, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle301, testSubFav401}},
			{16, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13650, 0, 0}},
		},
	}

	testSubFav41 = &Fav{
		FavNum:   3,
		NBoards:  3,
		NLines:   0,
		NFolders: 0,
		LineID:   0,
		FolderID: 0,
		Depth:    1,
		Favh: []*FavType{
			{20, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13810, 0, 0}},
			{21, pttbbsfav.FAVT_BOARD, 1, &FavBoard{10, 0, 0}},
			{22, pttbbsfav.FAVT_BOARD, 1, &FavBoard{326, 0, 0}},
		},
	}

	testFav4 = &Fav{
		FavNum:   23,
		NBoards:  7,
		NLines:   2,
		NFolders: 2,
		LineID:   2,
		FolderID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{10, 0, 0}},
			{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{2569, 0, 0}},
			{2, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{3, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14468, 0, 0}},
			{4, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
			{5, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
			{6, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle30, testSubFav40}},
			{7, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{2, testTitle31, testSubFav41}},
			{8, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13703, 0, 0}},
			{9, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14326, 0, 0}},
			{10, pttbbsfav.FAVT_BOARD, 1, &FavBoard{12783, 0, 0}},
		},
	}

	testFav5 = &Fav{
		FavNum:   23,
		NBoards:  7,
		NLines:   2,
		NFolders: 2,
		LineID:   2,
		FolderID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{10, 0, 0}},
			{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{2569, 0, 0}},
			{2, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{3, pttbbsfav.FAVT_BOARD, 1, &FavBoard{14468, 0, 0}},
			{4, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
			{5, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}},
			{6, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{1, testTitle30, testSubFav40}},
			{7, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{2, testTitle31, testSubFav41}},
			{8, pttbbsfav.FAVT_BOARD, 1, &FavBoard{13703, 0, 0}},
			{9, pttbbsfav.FAVT_BOARD, 1, &FavBoard{21, 0, 0}},
			{10, pttbbsfav.FAVT_BOARD, 1, &FavBoard{23, 0, 0}},
		},
	}
)
