package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var (
	testUtf8Filename6            = "M.1644506386.A.ED0.utf8"
	testUtf8ContentAll6          []byte
	testUtf8Content6             []byte
	testUtf8Signature6           []byte
	testUtf8Comment6             []byte
	testUtf8FirstCommentsDBCS6   []byte
	testUtf8TheRestCommentsDBCS6 []byte
	testUtf8Content6Utf8         [][]*types.Rune

	testUtf8FirstComments6     []*schema.Comment
	testUtf8FullFirstComments6 []*schema.Comment
	testUtf8TheRestComments6   []*schema.Comment
)

func initTestUtf86() {
	testUtf8ContentAll6, testUtf8Content6, testUtf8Signature6, testUtf8Comment6, testUtf8FirstCommentsDBCS6, testUtf8TheRestCommentsDBCS6 = loadTest(testUtf8Filename6)

	testUtf8Content6Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:    "作者: hellohiro (☁☁☁☂) 看板: Gossiping",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: hellohiro (☁☁☁☂) 看板: Gossiping",
			},
		},
		{ // 1
			{
				Utf8:    "標題: [問卦] 為何打麻將叫賭博但買股票叫投資？",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [問卦] 為何打麻將叫賭博但買股票叫投資？",
			},
		},
		{ // 6
			{
				Utf8:    "時間: Fri Mar 20 09:43:42 2020",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Fri Mar 20 09:43:42 2020",
			},
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "打麻將摸幾圈",
				DBCSStr: "打麻將摸幾圈",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "-可能賠錢可能多賺",
				DBCSStr: "-可能賠錢可能多賺",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "買股票存股操盤",
				DBCSStr: "買股票存股操盤",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "-可能賠錢可能多賺",
				DBCSStr: "-可能賠錢可能多賺",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "目前看來打麻將賠的還比買股市少",
				DBCSStr: "目前看來打麻將賠的還比買股市少",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "為何打麻將就要被污名化叫賭博",
				DBCSStr: "為何打麻將就要被污名化叫賭博",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "買股票就可以叫投資吶～",
				DBCSStr: "買股票就可以叫投資吶～",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
	}

	testUtf8FirstComments6 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("lats"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "政府抽到的稅叫投資，政府抽不到稅叫賭博",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "政府抽到的稅叫投資，政府抽不到稅叫賭博",
					},
				},
			},
			MD5:     "M-lz7qoajllkXbuzxNL2Ww",
			IP:      "118.160.112.18",
			TheDate: "03/20 09:44",
			DBCSStr: "\x1b[1;37m推 \x1b[33mlats\x1b[m\x1b[33m: 政府抽到的稅叫投資，政府抽不到稅叫賭博    \x1b[m 118.160.112.18 03/20 09:44",
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("lats"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "不然你認為大樂透不算賭博嗎？",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "不然你認為大樂透不算賭博嗎？",
					},
				},
			},
			IP:      "118.160.112.18",
			TheDate: "03/20 09:45",
			MD5:     "uqvxsUOAumZKM7Ote65O5g",
			DBCSStr: "\x1b[1;31m→ \x1b[33mlats\x1b[m\x1b[33m: 不然你認為大樂透不算賭博嗎？              \x1b[m 118.160.112.18 03/20 09:45",
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("D1"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "打麻將有股利嗎？",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "打麻將有股利嗎？",
					},
				},
			},
			MD5:     "8M52DhmjUJDOUbeiKRdKwg",
			IP:      "101.137.108.233",
			TheDate: "03/20 09:45",
			DBCSStr: "\x1b[1;37m推 \x1b[33mD1\x1b[m\x1b[33m: 打麻將有股利嗎？                            \x1b[m101.137.108.233 03/20 09:45",
		},
		{ // 3
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ss1130"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "股票不賣就不賠啊，賭博輸了就是沒了",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "股票不賣就不賠啊，賭博輸了就是沒了",
					},
				},
			},
			MD5:     "bneaAP3dx10yrX_pqGTX9A",
			IP:      "114.37.83.201",
			TheDate: "03/20 09:45",
			DBCSStr: "\x1b[1;37m推 \x1b[33mss1130\x1b[m\x1b[33m: 股票不賣就不賠啊，賭博輸了就是沒了      \x1b[m  114.37.83.201 03/20 09:45",
		},
		{ // 4
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("jin956"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "開公司做生意有賺有賠 是賭博還是投資?",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "開公司做生意有賺有賠 是賭博還是投資?",
					},
				},
			},
			MD5:     "6LgHTCPxb3k7REJNlhf6GA",
			IP:      "110.26.190.156",
			TheDate: "03/20 09:45",
			DBCSStr: "\x1b[1;31m→ \x1b[33mjin956\x1b[m\x1b[33m: 開公司做生意有賺有賠 是賭博還是投資?    \x1b[m 110.26.190.156 03/20 09:45",
		},
		{ // 5
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "股票的意義是 一間公司想生產東西現在沒錢",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "股票的意義是 一間公司想生產東西現在沒錢",
					},
				},
			},
			MD5:     "gPPKXXc-6zGdsceRFTufkg",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:48",
			DBCSStr: "\x1b[1;31m→ \x1b[33mkumori\x1b[m\x1b[33m: 股票的意義是 一間公司想生產東西現在沒錢 \x1b[m122.118.131.149 03/20 09:48",
		},
		{ // 6
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("rindesu"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "四個人打麻將可以，你們一群人打麻將？",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "四個人打麻將可以，你們一群人打麻將？",
					},
				},
			},
			MD5:     "vo8LsdHalm5QR0jkoHsACg",
			IP:      "110.28.108.240",
			TheDate: "03/20 09:49",
			DBCSStr: "\x1b[1;37m推 \x1b[33mrindesu\x1b[m\x1b[33m: 四個人打麻將可以，你們一群人打麻將？   \x1b[m 110.28.108.240 03/20 09:49",
		},
		{ // 7
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "發行一張紙跟你借錢 你認為這間公司以後會",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "發行一張紙跟你借錢 你認為這間公司以後會",
					},
				},
			},
			MD5:     "vgSdv2gWw4H7hFaJF_An7A",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:49",
			DBCSStr: "\x1b[1;31m→ \x1b[33mkumori\x1b[m\x1b[33m: 發行一張紙跟你借錢 你認為這間公司以後會 \x1b[m122.118.131.149 03/20 09:49",
		},
		{ // 8
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("vowpool"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "應該用賽馬跟股市來做類比 比較像",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "應該用賽馬跟股市來做類比 比較像",
					},
				},
			},
			MD5:     "Ma3obNmlqjqpYfRYJ76vYg",
			IP:      "125.227.40.62",
			TheDate: "03/20 09:50",
			DBCSStr: "\x1b[1;31m→ \x1b[33mvowpool\x1b[m\x1b[33m: 應該用賽馬跟股市來做類比 比較像        \x1b[m  125.227.40.62 03/20 09:50",
		},
		{ // 9
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("vowpool"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "賽馬是賭博 股市是金融",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "賽馬是賭博 股市是金融",
					},
				},
			},
			MD5:     "vMnnNm6xdtuPOi0ztl4sZA",
			IP:      "125.227.40.62",
			TheDate: "03/20 09:50",
			DBCSStr: "\x1b[1;31m→ \x1b[33mvowpool\x1b[m\x1b[33m: 賽馬是賭博 股市是金融                  \x1b[m  125.227.40.62 03/20 09:50",
		},
	}
	testUtf8FullFirstComments6 = []*schema.Comment{}
	testUtf8TheRestComments6 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "賺 於是你借他錢 後來他真的賺了 你就可以",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "賺 於是你借他錢 後來他真的賺了 你就可以",
					},
				},
			},
			MD5:     "am4iq8LtZTp8p_GYd_G9-Q",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:50",
			DBCSStr: "\x1b[1;31m→ \x1b[33mkumori\x1b[m\x1b[33m: 賺 於是你借他錢 後來他真的賺了 你就可以 \x1b[m122.118.131.149 03/20 09:50",
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Besorgen"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "叔叔我看過打麻將，好像也有募資行為。",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "叔叔我看過打麻將，好像也有募資行為。",
					},
				},
			},
			MD5:     "GS_QPa9vSmLcP5xJDzod2Q",
			IP:      "125.230.35.42",
			TheDate: "03/20 09:51",
			DBCSStr: "\x1b[1;37m推 \x1b[33mBesorgen\x1b[m\x1b[33m: 叔叔我看過打麻將，好像也有募資行為。  \x1b[m  125.230.35.42 03/20 09:51",
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "得到本金 以及利息",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "得到本金 以及利息",
					},
				},
			},
			MD5:     "pSRX0Ox7KFM1nip7dmdIeg",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:51",
			DBCSStr: "\x1b[1;31m→ \x1b[33mkumori\x1b[m\x1b[33m: 得到本金 以及利息                       \x1b[m122.118.131.149 03/20 09:51",
		},
		{ // 3
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "拿回",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "拿回",
					},
				},
			},
			MD5:     "2wXn35Dn9hRYtaN-RZF3GA",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:52",
			DBCSStr: "\x1b[1;31m→ \x1b[33mkumori\x1b[m\x1b[33m: 拿回                                    \x1b[m122.118.131.149 03/20 09:52",
		},
		{ // 4
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("cc03233"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "重點 看誰做莊",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "重點 看誰做莊",
					},
				},
			},
			MD5:     "S05sF49K3wBZ_1RBE5SvlQ",
			IP:      "1.161.225.69",
			TheDate: "03/20 09:53",
			DBCSStr: "\x1b[1;31m→ \x1b[33mcc03233\x1b[m\x1b[33m: 重點 看誰做莊                          \x1b[m   1.161.225.69 03/20 09:53",
		},
		{ // 5
			TheType: ptttype.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("RuleAllWorld"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "你先找出一副有成長性的麻將給我看",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "你先找出一副有成長性的麻將給我看",
					},
				},
			},
			MD5:     "aTMqK0yOEFnkptsBJ8IIdw",
			IP:      "49.216.4.199",
			TheDate: "03/20 09:54",
			DBCSStr: "\x1b[1;31m噓 \x1b[33mRuleAllWorld\x1b[m\x1b[33m: 你先找出一副有成長性的麻將給我看  \x1b[m   49.216.4.199 03/20 09:54",
		},
		{ // 6
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("DinoZavolta"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "台灣人多的是把股市當賭場玩的",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "台灣人多的是把股市當賭場玩的",
					},
				},
			},
			MD5:     "S9ApOqExHyA7qPkaQbgv7Q",
			IP:      "112.105.78.63",
			TheDate: "03/20 10:27",
			DBCSStr: "\x1b[1;37m推 \x1b[33mDinoZavolta\x1b[m\x1b[33m: 台灣人多的是把股市當賭場玩的       \x1b[m  112.105.78.63 03/20 10:27",
		},
		{ // 7
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("oyui111"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "因為大家把股票當麻將玩",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "因為大家把股票當麻將玩",
					},
				},
			},
			MD5:     "jkKatMV8HPQxUhENf5wVpw",
			IP:      "223.141.120.153",
			TheDate: "03/20 10:29",
			DBCSStr: "\x1b[1;37m推 \x1b[33moyui111\x1b[m\x1b[33m: 因為大家把股票當麻將玩                 \x1b[m223.141.120.153 03/20 10:29",
		},
		{ // 8
			TheType: ptttype.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("jackeman"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "稅稅稅稅稅",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "稅稅稅稅稅",
					},
				},
			},
			MD5:     "LopnABllsd8x2w5MEenj3w",
			IP:      "61.223.69.223",
			TheDate: "03/20 10:33",
			DBCSStr: "\x1b[1;31m噓 \x1b[33mjackeman\x1b[m\x1b[33m: 稅稅稅稅稅                            \x1b[m  61.223.69.223 03/20 10:33",
		},
		{ // 9
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("tomxyz"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "症腐做莊的都合法喔 科科",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "症腐做莊的都合法喔 科科",
					},
				},
			},
			MD5:     "2uS06Ou67Y5VfwpStxv_Og",
			IP:      "42.72.59.33",
			TheDate: "03/20 10:57",
			DBCSStr: "\x1b[1;31m→ \x1b[33mtomxyz\x1b[m\x1b[33m: 症腐做莊的都合法喔 科科                 \x1b[m    42.72.59.33 03/20 10:57",
		},
		{ // 10
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("bryanma"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "國家開的啊",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "國家開的啊",
					},
				},
			},
			MD5:     "4oMXnXegENBL9Dn6-SddrA",
			IP:      "223.136.131.5",
			TheDate: "03/21 18:36",
			DBCSStr: "\x1b[1;37m推 \x1b[33mbryanma\x1b[m\x1b[33m: 國家開的啊                             \x1b[m  223.136.131.5 03/21 18:36",
		},
	}
}
