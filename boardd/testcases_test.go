package boardd

import (
	"runtime"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/sirupsen/logrus"
)

var (
	testUserecRaw1 = &ptttype.UserecRaw{
		Version:    4194,
		UserID:     ptttype.UserID_t{83, 89, 83, 79, 80},
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   ptttype.Nickname_t{175, 171},
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    65535,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
	}

	testUserecRaw2 = &ptttype.UserecRaw{
		Version:      ptttype.PASSWD_VERSION,
		UserID:       ptttype.UserID_t{'C', 'o', 'd', 'i', 'n', 'g', 'M', 'a', 'n'},
		UFlag:        33557216,
		UserLevel:    7,
		NumLoginDays: 1,
		NumPosts:     0,
		FirstLogin:   1600737659,
		LastLogin:    1600737960,
		LastHost:     ptttype.IPv4_t{'5', '9', '.', '1', '2', '4', '.', '1', '6', '7', '.', '2', '2', '6'},
	}

	testUserecRaw3 = &ptttype.UserecRaw{
		Version:      ptttype.PASSWD_VERSION,
		UserID:       ptttype.UserID_t{'t', 'e', 's', 't'},
		UFlag:        33557216,
		UserLevel:    7 | ptttype.PERM_BOARD | ptttype.PERM_POST | ptttype.PERM_LOGINOK,
		NumLoginDays: 1,
		NumPosts:     0,
		FirstLogin:   1600737659,
		LastLogin:    1600737960,
		LastHost:     ptttype.IPv4_t{'5', '9', '.', '1', '2', '4', '.', '1', '6', '7', '.', '2', '2', '6'},
	}

	testUserecRaw4 = &ptttype.UserecRaw{
		Version:      ptttype.PASSWD_VERSION,
		UserID:       ptttype.UserID_t{'S', 'Y', 'S', 'O', 'P', '3'},
		Nickname:     ptttype.Nickname_t{'s', 't', 'r', 'i', 'n', 'g'},
		PasswdHash:   ptttype.Passwd_t{0x2c, 0x52, 0x69, 0x36, 0x53, 0x6b, 0x55, 0x33, 0x34, 0x72, 0x65, 0x74, 0x41},
		UFlag:        33557088,
		UserLevel:    7,
		NumLoginDays: 23,
		NumPosts:     0,
		FirstLogin:   1608226066,
		LastLogin:    1628305138,
		LastHost:     ptttype.IPv4_t{'1', '7', '2', '.', '1', '9', '.', '0', '.', '1'},
		Email:        ptttype.Email_t{'t', 'e', 's', 't', '@', 't', 'e', 's', 't', '.', 'c', 'o', 'm', '.', 't', 'w'},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		LastSeen:     1628305138,
	}

	testNewPostUser1 = &ptttype.UserecRaw{
		Version:    4194,
		UserID:     ptttype.UserID_t{65, 49}, // A1
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   ptttype.Nickname_t{175, 171},
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    7 | ptttype.PERM_LOGINOK | ptttype.PERM_POST,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
	}

	testNewPostUser2 = &ptttype.UserecRaw{
		Version:    4194,
		UserID:     ptttype.UserID_t{65, 50}, // A2
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   ptttype.Nickname_t{175, 171},
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    7 | ptttype.PERM_LOGINOK | ptttype.PERM_POST,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
	}

	testSetupNewUser1 = &ptttype.UserecRaw{
		Version:    4194,
		UserID:     ptttype.UserID_t{65, 48}, // A0
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   ptttype.Nickname_t{175, 171},
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    536871943,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
	}

	testNewRegister1 = &ptttype.UserecRaw{
		Version:    4194,
		UserID:     ptttype.UserID_t{66, 49}, // B1
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   ptttype.Nickname_t{175, 171},
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    7,
		NumLoginDays: 1,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
		UaVersion:    128,
	}

	testNewUserInfoRawUserecRaw = &ptttype.UserecRaw{
		Version:    4194,
		UserID:     ptttype.UserID_t{66, 49}, // B1
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   ptttype.Nickname_t{0x1e, 0x80, 0x30, 0x40, 0x80, 0x40},
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    7,
		NumLoginDays: 1,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
		UaVersion:    128,
	}

	testNewUserInfoRawUserInfoRaw = &ptttype.UserInfoRaw{
		Pid:  types.Pid_t(types.DEFAULT_PID_MAX + 10),
		UID:  10,
		Mode: ptttype.USER_OP_LOGIN,

		UserID:   ptttype.UserID_t{66, 49},
		Nickname: ptttype.Nickname_t{0x30, 0x40, 0x80, 0x40, 0x00, 0x40},

		UserLevel: 7,
		FromIP:    0xc0a80001, // 192.168.0.1
		From:      ptttype.From_t{'1', '9', '2', '.', '1', '6', '8', '.', '0', '.', '1'},
		Pager:     ptttype.PAGER_ON,
	}

	testGetNewUtmpEnt0 = &ptttype.UserInfoRaw{
		Pid:  types.Pid_t(types.DEFAULT_PID_MAX + 10),
		UID:  10,
		Mode: ptttype.USER_OP_LOGIN,

		Nickname: ptttype.Nickname_t{0x30, 0x40, 0x80, 0x40, 0x00, 0x40},

		UserLevel: 7,
		FromIP:    0xc0a80001, // 192.168.0.1
		From:      ptttype.From_t{'1', '9', '2', '.', '1', '6', '8', '.', '0', '.', '1'},
		Pager:     ptttype.PAGER_ON,
	}

	testUserInfo1 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'S', 'Y', 'S', 'O', 'P'},
		UID:    1,
		From:   ptttype.From_t{'D'},
		Pid:    3,
	}

	testUserInfo2 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'A', '1'},
		UID:    2,
		From:   ptttype.From_t{'B'},
		Pid:    2,
	}

	testUserInfo3 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'A', '0'},
		UID:    3,
		From:   ptttype.From_t{'S'},
		Pid:    1,
	}
	testUserInfo4 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'Z', '2'},
		UID:    5,
		From:   ptttype.From_t{'K'},
		Pid:    5,
	}
	testUserInfo5 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'Z', '1'},
		UID:    4,
		From:   ptttype.From_t{'H'},
		Pid:    4,
	}

	testUserInfo6 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'Z', '3'},
		UID:    6,
		From:   ptttype.From_t{'K'},
		Pid:    6,
	}

	testNewUserInfoRawNickname = ptttype.Nickname_t{0x1e, 0x80, 0x30, 0x40, 0x80, 0x40}

	testNewRegister1Passwd = []byte("!@Ab86")

	testBoardSummary6 = &ptttype.BoardSummaryRaw{
		Gid:     5,
		Bid:     6,
		BrdAttr: ptttype.BRD_POSTMASK,
		Brdname: &ptttype.BoardID_t{'A', 'L', 'L', 'P', 'O', 'S', 'T', 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xf3, 0xaa,
			0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0xb7, 0x73,
			0xa4, 0xe5, 0xb3, 0xb9, 0x00, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testBoardSummary7 = &ptttype.BoardSummaryRaw{
		Gid:     5,
		Bid:     7,
		Brdname: &ptttype.BoardID_t{0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xea, 0xb7,
			0xbd, 0xa6, 0x5e, 0xa6, 0xac, 0xb5, 0xa9, 0x00, 0xb7, 0x73,
			0xa4, 0xe5, 0xb3, 0xb9, 0x00, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testBoardSummary11 = &ptttype.BoardSummaryRaw{
		Gid:     5,
		Bid:     11,
		Brdname: &ptttype.BoardID_t{0x45, 0x64, 0x69, 0x74, 0x45, 0x78, 0x70, 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xbd, 0x64, 0xa5,
			0xbb, 0xba, 0xeb, 0xc6, 0x46, 0xa7, 0xeb, 0xbd, 0x5a, 0xb0,
			0xcf, 0x00, 0xd6, 0xa1, 0x49, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testBoardSummary8 = &ptttype.BoardSummaryRaw{
		Gid:     5,
		Bid:     8,
		Brdname: &ptttype.BoardID_t{0x4e, 0x6f, 0x74, 0x65, 0x00, 0x65, 0x64, 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb0, 0xca, 0xba,
			0x41, 0xac, 0xdd, 0xaa, 0x4f, 0xa4, 0xce, 0xba, 0x71, 0xa6,
			0xb1, 0xa7, 0xeb, 0xbd, 0x5a, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testBoardSummary1 = &ptttype.BoardSummaryRaw{
		Gid:     2,
		Bid:     1,
		BrdAttr: ptttype.BRD_POSTMASK,
		Brdname: &ptttype.BoardID_t{'S', 'Y', 'S', 'O', 'P'},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		},
		BM: []*ptttype.UserID_t{},
	}

	testBoardSummary9 = &ptttype.BoardSummaryRaw{
		Gid:     5,
		Bid:     9,
		BrdAttr: ptttype.BRD_POSTMASK,
		Brdname: &ptttype.BoardID_t{'R', 'e', 'c', 'o', 'r', 'd', 0x00, 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa7, 0xda, 0xad,
			0xcc, 0xaa, 0xba, 0xa6, 0xa8, 0xaa, 0x47, 0x00, 0x71, 0xa6,
			0xb1, 0xa7, 0xeb, 0xbd, 0x5a, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
		BM: []*ptttype.UserID_t{},
	}

	testBoardSummary10 = &ptttype.BoardSummaryRaw{
		Gid:     5,
		Bid:     10,
		Brdname: &ptttype.BoardID_t{'W', 'h', 'o', 'A', 'm', 'I', 0x00, 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa8, 0xfe, 0xa8,
			0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71, 0xa7, 0xda, 0xac,
			0x4f, 0xbd, 0xd6, 0xa1, 0x49, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x0, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testTitle13        *ptttype.BoardTitle_t
	testBoardSummary13 *ptttype.BoardSummaryRaw

	testBoardHeaderRaw1 = &ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'t', 'e', 's', 't'},
		BrdAttr: ptttype.BRD_HIDE,
	}

	testBoardHeaderRaw2 = &ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'W', 'h', 'o', 'A', 'm', 'I'},
		BrdAttr: ptttype.BRD_HIDE,
	}

	testBoardHeaderRaw10 = &ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'W', 'h', 'o', 'A', 'm', 'I', 0, 0, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa8, 0xfe, 0xa8,
			0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71, 0xa7, 0xda, 0xac,
			0x4f, 0xbd, 0xd6, 0xa1, 0x49, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		Gid: 5,
	}

	testBoardDetailRaw10 = &ptttype.BoardDetailRaw{
		Bid:            10,
		BoardHeaderRaw: testBoardHeaderRaw10,
	}

	testClassSummary2 = &ptttype.BoardSummaryRaw{
		Gid:     1,
		Bid:     2,
		BrdAttr: ptttype.BRD_GROUPBOARD,
		Brdname: &ptttype.BoardID_t{
			'1', '.', '.', '.', '.', '.', '.', '.', '.', '.',
			'.', '.',
		},
		Title: &ptttype.BoardTitle_t{
			0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0xa3, 0x55, 0xa4, 0xa4, 0xa5,
			0xa1, 0xac, 0x46, 0xa9, 0xb2, 0x20, 0x20, 0xa1, 0x6d, 0xb0,
			0xaa, 0xc0, 0xa3, 0xa6, 0x4d, 0xc0, 0x49, 0x2c, 0xab, 0x44,
			0xa4, 0x48, 0xa5, 0x69, 0xbc, 0xc4, 0xa1, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testClassSummary5 = &ptttype.BoardSummaryRaw{
		Gid:     1,
		Bid:     5,
		BrdAttr: ptttype.BRD_GROUPBOARD,
		Brdname: &ptttype.BoardID_t{
			'2', '.', '.', '.', '.', '.', '.', '.', '.', '.',
			'.', '.',
		},
		Title: &ptttype.BoardTitle_t{
			0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0xa3, 0x55, 0xa5, 0xab, 0xa5,
			0xc1, 0xbc, 0x73, 0xb3, 0xf5, 0x20, 0x20, 0x20, 0x20, 0x20,
			0xb3, 0xf8, 0xa7, 0x69, 0x20, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		BM: []*ptttype.UserID_t{},
	}

	testArticleSummary0 = &ptttype.ArticleSummaryRaw{
		Aid: 1,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607202239.A.30D
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x32, 0x32, 0x33, 0x39, 0x2e, 0x41, 0x2e, 0x33,
				0x30, 0x44,
			},
			Modified: 1607202238,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x50}, // SYSOP
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[問題] 我是誰？～
				0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7,
				0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x48, 0xa1,
				0xe3,
			},
		},
	}
	testArticleSummary1 = &ptttype.ArticleSummaryRaw{
		Aid: 2,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607203395.A.00D
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x33, 0x33, 0x39, 0x35, 0x2e, 0x41, 0x2e, 0x30,
				0x30, 0x44,
			},
			Modified: 1607203394,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x50}, // SYSOP
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[心得] 然後呢？～
				0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
				0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
				0xe3,
			},
			Filemode: ptttype.FILE_MARKED,
		},
	}

	testBottomSummary1 = &ptttype.ArticleSummaryRaw{
		Aid: 1,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607203395.A.00D
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x33, 0x33, 0x39, 0x35, 0x2e, 0x41, 0x2e, 0x30,
				0x30, 0x44,
			},
			Modified: 1607203394,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x50}, // SYSOP
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[心得] 然後呢？～
				0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
				0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
				0xe3,
			},
			Filemode: ptttype.FILE_MULTI,
		},
	}

	testContent1 = []byte{
		0xa7, 0x40, 0xaa, 0xcc, 0x3a, 0x20, 0x53, 0x59, 0x53,
		0x4f, 0x50, 0x20, 0x28, 0x29, 0x20, 0xac, 0xdd, 0xaa,
		0x4f, 0x3a, 0x20, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49,
		0x0a, 0xbc, 0xd0, 0xc3, 0x44, 0x3a, 0x20, 0x5b, 0xb0,
		0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7, 0xda, 0xac, 0x4f,
		0xbd, 0xd6, 0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0xae, 0xc9,
		0xb6, 0xa1, 0x3a, 0x20, 0x53, 0x75, 0x6e, 0x20, 0x44,
		0x65, 0x63, 0x20, 0x20, 0x36, 0x20, 0x30, 0x35, 0x3a,
		0x30, 0x33, 0x3a, 0x35, 0x37, 0x20, 0x32, 0x30, 0x32,
		0x30, 0x0a, 0x0a, 0xa7, 0xda, 0xac, 0x4f, 0xbd, 0xd6,
		0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0x0a, 0xa7, 0xda, 0xa6,
		0x62, 0xad, 0xfe, 0xb8, 0xcc, 0xa1, 0x48, 0xa1, 0xe3,
		0x0a, 0x0a, 0xa7, 0xda, 0xac, 0xb0, 0xa4, 0xb0, 0xbb,
		0xf2, 0xb7, 0x7c, 0xa6, 0x62, 0xb3, 0x6f, 0xb8, 0xcc,
		0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0x0a, 0x2d,
		0x2d, 0x0a, 0xa1, 0xb0, 0x20, 0xb5, 0x6f, 0xab, 0x48,
		0xaf, 0xb8, 0x3a, 0x20, 0xa7, 0xe5, 0xbd, 0xf0, 0xbd,
		0xf0, 0x20, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x28,
		0x70, 0x74, 0x74, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
		0x2e, 0x74, 0x65, 0x73, 0x74, 0x29, 0x2c, 0x20, 0xa8,
		0xd3, 0xa6, 0xdb, 0x3a, 0x20, 0x31, 0x37, 0x32, 0x2e,
		0x31, 0x38, 0x2e, 0x30, 0x2e, 0x31, 0x0a,
	}
)

func initVars() {
	if testTitle13 != nil {
		if testBoardSummary13.Brdname == nil || testBoardSummary13.Brdname[0] == 0 {
			logrus.Errorf("initVars: invalid testBoardSummary13: %v", testBoardSummary13)
		}
		return
	}

	testTitle13 = &ptttype.BoardTitle_t{}
	copy(testTitle13[:], []byte("CPBL \xa1\xb7new-board"))

	testBoardSummary13 = &ptttype.BoardSummaryRaw{
		Gid:      2,
		Bid:      13,
		Brdname:  &ptttype.BoardID_t{'m', 'n', 'e', 'w', 'b', 'o', 'a', 'r', 'd', '0'},
		Title:    testTitle13,
		BM:       []*ptttype.UserID_t{},
		StatAttr: ptttype.NBRD_FAV,
		BrdAttr:  0x200000,
	}
}

func freeTestVars() {
	testTitle13 = nil
	testBoardSummary13 = nil

	runtime.GC()
}
