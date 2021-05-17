package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename18            = "temp13"
	testContentAll18          []byte
	testContent18             []byte
	testSignature18           []byte
	testComment18             []byte
	testFirstCommentsDBCS18   []byte
	testTheRestCommentsDBCS18 []byte
	testContent18Big5         [][]*types.Rune
	testContent18Utf8         [][]*types.Rune

	testFirstComments18     []*schema.Comment
	testFullFirstComments18 []*schema.Comment
	testTheRestComments18   []*schema.Comment
)

func initTest18() {
	testContentAll18, testContent18, testSignature18, testComment18, testFirstCommentsDBCS18, testTheRestCommentsDBCS18 = loadTest(testFilename18)

	testContent18Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: hellohiro (\x9d\xba\x9d\xba\x9d\xba\x9d\xbc) \xac\xdd\xaaO: Gossiping"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: hellohiro (\x9d\xba\x9d\xba\x9d\xba\x9d\xbc) \xac\xdd\xaaO: Gossiping\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xa8\xf6] \xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xa5s\xbd\xe4\xb3\xd5\xa6\xfd\xb6R\xaa\xd1\xb2\xbc\xa5s\xa7\xeb\xb8\xea\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xa8\xf6] \xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xa5s\xbd\xe4\xb3\xd5\xa6\xfd\xb6R\xaa\xd1\xb2\xbc\xa5s\xa7\xeb\xb8\xea\xa1H\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Mar 20 09:43:42 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Mar 20 09:43:42 2020\r"),
			},
		},
		{ //3
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //4
			{
				Big5:   []byte("\xa5\xb4\xb3\xc2\xb1N\xbaN\xb4X\xb0\xe9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xb4\xb3\xc2\xb1N\xbaN\xb4X\xb0\xe9\r"),
			},
		},
		{ //5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //6
			{
				Big5:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8\r"),
			},
		},
		{ //7
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("\xb6R\xaa\xd1\xb2\xbc\xa6s\xaa\xd1\xbe\xde\xbdL"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6R\xaa\xd1\xb2\xbc\xa6s\xaa\xd1\xbe\xde\xbdL\r"),
			},
		},
		{ //9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //10
			{

				Big5:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8\r"),
			},
		},
		{ //11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //12
			{

				Big5:   []byte("\xa5\xd8\xabe\xac\xdd\xa8\xd3\xa5\xb4\xb3\xc2\xb1N\xbd\xdf\xaa\xba\xc1\xd9\xa4\xf1\xb6R\xaa\xd1\xa5\xab\xa4\xd6"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xd8\xabe\xac\xdd\xa8\xd3\xa5\xb4\xb3\xc2\xb1N\xbd\xdf\xaa\xba\xc1\xd9\xa4\xf1\xb6R\xaa\xd1\xa5\xab\xa4\xd6\r"),
			},
		},
		{ //13
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //14
			{
				Big5:   []byte("\xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xb4N\xadn\xb3Q\xa6\xc3\xa6W\xa4\xc6\xa5s\xbd\xe4\xb3\xd5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xb4N\xadn\xb3Q\xa6\xc3\xa6W\xa4\xc6\xa5s\xbd\xe4\xb3\xd5\r"),
			},
		},
		{ //15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //16
			{
				Big5:   []byte("\xb6R\xaa\xd1\xb2\xbc\xb4N\xa5i\xa5H\xa5s\xa7\xeb\xb8\xea\xa7o\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6R\xaa\xd1\xb2\xbc\xb4N\xa5i\xa5H\xa5s\xa7\xeb\xb8\xea\xa7o\xa1\xe3\r"),
			},
		},
	}

	testContent18Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: hellohiro (☁☁☁☂) 看板: Gossiping",
				Big5:   []byte("\xa7@\xaa\xcc: hellohiro (\x9d\xba\x9d\xba\x9d\xba\x9d\xbc) \xac\xdd\xaaO: Gossiping"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: hellohiro (\x9d\xba\x9d\xba\x9d\xba\x9d\xbc) \xac\xdd\xaaO: Gossiping\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: [問卦] 為何打麻將叫賭博但買股票叫投資？",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xa8\xf6] \xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xa5s\xbd\xe4\xb3\xd5\xa6\xfd\xb6R\xaa\xd1\xb2\xbc\xa5s\xa7\xeb\xb8\xea\xa1H"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb0\xdd\xa8\xf6] \xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xa5s\xbd\xe4\xb3\xd5\xa6\xfd\xb6R\xaa\xd1\xb2\xbc\xa5s\xa7\xeb\xb8\xea\xa1H\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Fri Mar 20 09:43:42 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Mar 20 09:43:42 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Mar 20 09:43:42 2020\r"),
			},
		},
		{ //3
			{

				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //4
			{
				Utf8:   "打麻將摸幾圈",
				Big5:   []byte("\xa5\xb4\xb3\xc2\xb1N\xbaN\xb4X\xb0\xe9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xb4\xb3\xc2\xb1N\xbaN\xb4X\xb0\xe9\r"),
			},
		},
		{ //5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //6
			{
				Utf8:   "-可能賠錢可能多賺",
				Big5:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8\r"),
			},
		},
		{ //7
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //8
			{
				Utf8:   "買股票存股操盤",
				Big5:   []byte("\xb6R\xaa\xd1\xb2\xbc\xa6s\xaa\xd1\xbe\xde\xbdL"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6R\xaa\xd1\xb2\xbc\xa6s\xaa\xd1\xbe\xde\xbdL\r"),
			},
		},
		{ //9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //10
			{
				Utf8:   "-可能賠錢可能多賺",
				Big5:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("-\xa5i\xaf\xe0\xbd\xdf\xbf\xfa\xa5i\xaf\xe0\xa6h\xc1\xc8\r"),
			},
		},
		{ //11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //12
			{
				Utf8:   "目前看來打麻將賠的還比買股市少",
				Big5:   []byte("\xa5\xd8\xabe\xac\xdd\xa8\xd3\xa5\xb4\xb3\xc2\xb1N\xbd\xdf\xaa\xba\xc1\xd9\xa4\xf1\xb6R\xaa\xd1\xa5\xab\xa4\xd6"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xd8\xabe\xac\xdd\xa8\xd3\xa5\xb4\xb3\xc2\xb1N\xbd\xdf\xaa\xba\xc1\xd9\xa4\xf1\xb6R\xaa\xd1\xa5\xab\xa4\xd6\r"),
			},
		},
		{ //13
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //14
			{
				Utf8:   "為何打麻將就要被污名化叫賭博",
				Big5:   []byte("\xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xb4N\xadn\xb3Q\xa6\xc3\xa6W\xa4\xc6\xa5s\xbd\xe4\xb3\xd5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xac\xb0\xa6\xf3\xa5\xb4\xb3\xc2\xb1N\xb4N\xadn\xb3Q\xa6\xc3\xa6W\xa4\xc6\xa5s\xbd\xe4\xb3\xd5\r"),
			},
		},
		{ //15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //16
			{
				Utf8:   "買股票就可以叫投資吶～",
				Big5:   []byte("\xb6R\xaa\xd1\xb2\xbc\xb4N\xa5i\xa5H\xa5s\xa7\xeb\xb8\xea\xa7o\xa1\xe3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6R\xaa\xd1\xb2\xbc\xb4N\xa5i\xa5H\xa5s\xa7\xeb\xb8\xea\xa7o\xa1\xe3\r"),
			},
		},
	}

	testFirstComments18 = []*schema.Comment{
		{ //0
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("lats"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "政府抽到的稅叫投資，政府抽不到稅叫賭博",
						Big5:   []byte("\xacF\xa9\xb2\xa9\xe2\xa8\xec\xaa\xba\xb5|\xa5s\xa7\xeb\xb8\xea\xa1A\xacF\xa9\xb2\xa9\xe2\xa4\xa3\xa8\xec\xb5|\xa5s\xbd\xe4\xb3\xd5    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xacF\xa9\xb2\xa9\xe2\xa8\xec\xaa\xba\xb5|\xa5s\xa7\xeb\xb8\xea\xa1A\xacF\xa9\xb2\xa9\xe2\xa4\xa3\xa8\xec\xb5|\xa5s\xbd\xe4\xb3\xd5    "),
					},
				},
			},
			MD5:     "re1zjGmRX94EC6xhYp4stQ",
			IP:      "118.160.112.18",
			TheDate: "03/20 09:44",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mlats\x1b[m\x1b[33m: \xacF\xa9\xb2\xa9\xe2\xa8\xec\xaa\xba\xb5|\xa5s\xa7\xeb\xb8\xea\xa1A\xacF\xa9\xb2\xa9\xe2\xa4\xa3\xa8\xec\xb5|\xa5s\xbd\xe4\xb3\xd5    \x1b[m 118.160.112.18 03/20 09:44\r"),
		},
		{ //1
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("lats"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不然你認為大樂透不算賭博嗎？",
						Big5:   []byte("\xa4\xa3\xb5M\xa7A\xbb{\xac\xb0\xa4j\xbc\xd6\xb3z\xa4\xa3\xba\xe2\xbd\xe4\xb3\xd5\xb6\xdc\xa1H              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xa3\xb5M\xa7A\xbb{\xac\xb0\xa4j\xbc\xd6\xb3z\xa4\xa3\xba\xe2\xbd\xe4\xb3\xd5\xb6\xdc\xa1H              "),
					},
				},
			},
			IP:      "118.160.112.18",
			TheDate: "03/20 09:45",
			MD5:     "u5v6PWoN_O3ri9oBsiOqFw",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mlats\x1b[m\x1b[33m: \xa4\xa3\xb5M\xa7A\xbb{\xac\xb0\xa4j\xbc\xd6\xb3z\xa4\xa3\xba\xe2\xbd\xe4\xb3\xd5\xb6\xdc\xa1H              \x1b[m 118.160.112.18 03/20 09:45\r"),
		},
		{ //2
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("D1"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "打麻將有股利嗎？",
						Big5:   []byte("\xa5\xb4\xb3\xc2\xb1N\xa6\xb3\xaa\xd1\xa7Q\xb6\xdc\xa1H                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5\xb4\xb3\xc2\xb1N\xa6\xb3\xaa\xd1\xa7Q\xb6\xdc\xa1H                            "),
					},
				},
			},
			MD5:     "nAnF91Ukrs9IefrybKz_HQ",
			IP:      "101.137.108.233",
			TheDate: "03/20 09:45",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mD1\x1b[m\x1b[33m: \xa5\xb4\xb3\xc2\xb1N\xa6\xb3\xaa\xd1\xa7Q\xb6\xdc\xa1H                            \x1b[m101.137.108.233 03/20 09:45\r"),
		},
		{ //3
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("ss1130"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "股票不賣就不賠啊，賭博輸了就是沒了",
						Big5:   []byte("\xaa\xd1\xb2\xbc\xa4\xa3\xbd\xe6\xb4N\xa4\xa3\xbd\xdf\xb0\xda\xa1A\xbd\xe4\xb3\xd5\xbf\xe9\xa4F\xb4N\xacO\xa8S\xa4F      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaa\xd1\xb2\xbc\xa4\xa3\xbd\xe6\xb4N\xa4\xa3\xbd\xdf\xb0\xda\xa1A\xbd\xe4\xb3\xd5\xbf\xe9\xa4F\xb4N\xacO\xa8S\xa4F      "),
					},
				},
			},
			MD5:     "ZxvPgcGyIk2nVe7ZzxdzoA",
			IP:      "114.37.83.201",
			TheDate: "03/20 09:45",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mss1130\x1b[m\x1b[33m: \xaa\xd1\xb2\xbc\xa4\xa3\xbd\xe6\xb4N\xa4\xa3\xbd\xdf\xb0\xda\xa1A\xbd\xe4\xb3\xd5\xbf\xe9\xa4F\xb4N\xacO\xa8S\xa4F      \x1b[m  114.37.83.201 03/20 09:45\r"),
		},
		{ //4
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("jin956"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "開公司做生意有賺有賠 是賭博還是投資?",
						Big5:   []byte("\xb6}\xa4\xbd\xa5q\xb0\xb5\xa5\xcd\xb7N\xa6\xb3\xc1\xc8\xa6\xb3\xbd\xdf \xacO\xbd\xe4\xb3\xd5\xc1\xd9\xacO\xa7\xeb\xb8\xea?    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb6}\xa4\xbd\xa5q\xb0\xb5\xa5\xcd\xb7N\xa6\xb3\xc1\xc8\xa6\xb3\xbd\xdf \xacO\xbd\xe4\xb3\xd5\xc1\xd9\xacO\xa7\xeb\xb8\xea?    "),
					},
				},
			},
			MD5:     "v7aPdDoGIeABPHYvkfe2TA",
			IP:      "110.26.190.156",
			TheDate: "03/20 09:45",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mjin956\x1b[m\x1b[33m: \xb6}\xa4\xbd\xa5q\xb0\xb5\xa5\xcd\xb7N\xa6\xb3\xc1\xc8\xa6\xb3\xbd\xdf \xacO\xbd\xe4\xb3\xd5\xc1\xd9\xacO\xa7\xeb\xb8\xea?    \x1b[m 110.26.190.156 03/20 09:45\r"),
		},
		{ //5
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "股票的意義是 一間公司想生產東西現在沒錢",
						Big5:   []byte("\xaa\xd1\xb2\xbc\xaa\xba\xb7N\xb8q\xacO \xa4@\xb6\xa1\xa4\xbd\xa5q\xb7Q\xa5\xcd\xb2\xa3\xaaF\xa6\xe8\xb2{\xa6b\xa8S\xbf\xfa "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaa\xd1\xb2\xbc\xaa\xba\xb7N\xb8q\xacO \xa4@\xb6\xa1\xa4\xbd\xa5q\xb7Q\xa5\xcd\xb2\xa3\xaaF\xa6\xe8\xb2{\xa6b\xa8S\xbf\xfa "),
					},
				},
			},
			MD5:     "0Z_CmWQ03_875tDRT_AH4Q",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:48",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xaa\xd1\xb2\xbc\xaa\xba\xb7N\xb8q\xacO \xa4@\xb6\xa1\xa4\xbd\xa5q\xb7Q\xa5\xcd\xb2\xa3\xaaF\xa6\xe8\xb2{\xa6b\xa8S\xbf\xfa \x1b[m122.118.131.149 03/20 09:48\r"),
		},
		{ //6
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("rindesu"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "四個人打麻將可以，你們一群人打麻將？",
						Big5:   []byte("\xa5|\xad\xd3\xa4H\xa5\xb4\xb3\xc2\xb1N\xa5i\xa5H\xa1A\xa7A\xad\xcc\xa4@\xb8s\xa4H\xa5\xb4\xb3\xc2\xb1N\xa1H   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5|\xad\xd3\xa4H\xa5\xb4\xb3\xc2\xb1N\xa5i\xa5H\xa1A\xa7A\xad\xcc\xa4@\xb8s\xa4H\xa5\xb4\xb3\xc2\xb1N\xa1H   "),
					},
				},
			},
			MD5:     "jd3J52KJCF5KQzt2GErBJg",
			IP:      "110.28.108.240",
			TheDate: "03/20 09:49",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mrindesu\x1b[m\x1b[33m: \xa5|\xad\xd3\xa4H\xa5\xb4\xb3\xc2\xb1N\xa5i\xa5H\xa1A\xa7A\xad\xcc\xa4@\xb8s\xa4H\xa5\xb4\xb3\xc2\xb1N\xa1H   \x1b[m 110.28.108.240 03/20 09:49\r"),
		},
		{ //7
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "發行一張紙跟你借錢 你認為這間公司以後會",
						Big5:   []byte("\xb5o\xa6\xe6\xa4@\xb1i\xaf\xc8\xb8\xf2\xa7A\xad\xc9\xbf\xfa \xa7A\xbb{\xac\xb0\xb3o\xb6\xa1\xa4\xbd\xa5q\xa5H\xab\xe1\xb7| "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb5o\xa6\xe6\xa4@\xb1i\xaf\xc8\xb8\xf2\xa7A\xad\xc9\xbf\xfa \xa7A\xbb{\xac\xb0\xb3o\xb6\xa1\xa4\xbd\xa5q\xa5H\xab\xe1\xb7| "),
					},
				},
			},
			MD5:     "I6z7DF7JHmn_RveuhTM9kg",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:49",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xb5o\xa6\xe6\xa4@\xb1i\xaf\xc8\xb8\xf2\xa7A\xad\xc9\xbf\xfa \xa7A\xbb{\xac\xb0\xb3o\xb6\xa1\xa4\xbd\xa5q\xa5H\xab\xe1\xb7| \x1b[m122.118.131.149 03/20 09:49\r"),
		},
		{ //8
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("vowpool"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "應該用賽馬跟股市來做類比 比較像",
						Big5:   []byte("\xc0\xb3\xb8\xd3\xa5\xce\xc1\xc9\xb0\xa8\xb8\xf2\xaa\xd1\xa5\xab\xa8\xd3\xb0\xb5\xc3\xfe\xa4\xf1 \xa4\xf1\xb8\xfb\xb9\xb3        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc0\xb3\xb8\xd3\xa5\xce\xc1\xc9\xb0\xa8\xb8\xf2\xaa\xd1\xa5\xab\xa8\xd3\xb0\xb5\xc3\xfe\xa4\xf1 \xa4\xf1\xb8\xfb\xb9\xb3        "),
					},
				},
			},
			MD5:     "HxFf1L0eaMiqE3wyJ0Nyow",
			IP:      "125.227.40.62",
			TheDate: "03/20 09:50",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mvowpool\x1b[m\x1b[33m: \xc0\xb3\xb8\xd3\xa5\xce\xc1\xc9\xb0\xa8\xb8\xf2\xaa\xd1\xa5\xab\xa8\xd3\xb0\xb5\xc3\xfe\xa4\xf1 \xa4\xf1\xb8\xfb\xb9\xb3        \x1b[m  125.227.40.62 03/20 09:50\r"),
		},
		{ //9
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("vowpool"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "賽馬是賭博 股市是金融",
						Big5:   []byte("\xc1\xc9\xb0\xa8\xacO\xbd\xe4\xb3\xd5 \xaa\xd1\xa5\xab\xacO\xaa\xf7\xbf\xc4                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc9\xb0\xa8\xacO\xbd\xe4\xb3\xd5 \xaa\xd1\xa5\xab\xacO\xaa\xf7\xbf\xc4                  "),
					},
				},
			},
			MD5:     "KZhYeiENlhq_U1YasuvFlQ",
			IP:      "125.227.40.62",
			TheDate: "03/20 09:50",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mvowpool\x1b[m\x1b[33m: \xc1\xc9\xb0\xa8\xacO\xbd\xe4\xb3\xd5 \xaa\xd1\xa5\xab\xacO\xaa\xf7\xbf\xc4                  \x1b[m  125.227.40.62 03/20 09:50\r"),
		},
	}

	testFullFirstComments18 = []*schema.Comment{
		{ //0
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("lats"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "政府抽到的稅叫投資，政府抽不到稅叫賭博",
						Big5:   []byte("\xacF\xa9\xb2\xa9\xe2\xa8\xec\xaa\xba\xb5|\xa5s\xa7\xeb\xb8\xea\xa1A\xacF\xa9\xb2\xa9\xe2\xa4\xa3\xa8\xec\xb5|\xa5s\xbd\xe4\xb3\xd5    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xacF\xa9\xb2\xa9\xe2\xa8\xec\xaa\xba\xb5|\xa5s\xa7\xeb\xb8\xea\xa1A\xacF\xa9\xb2\xa9\xe2\xa4\xa3\xa8\xec\xb5|\xa5s\xbd\xe4\xb3\xd5    "),
					},
				},
			},
			MD5:        "re1zjGmRX94EC6xhYp4stQ",
			IP:         "118.160.112.18",
			TheDate:    "03/20 09:44",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mlats\x1b[m\x1b[33m: \xacF\xa9\xb2\xa9\xe2\xa8\xec\xaa\xba\xb5|\xa5s\xa7\xeb\xb8\xea\xa1A\xacF\xa9\xb2\xa9\xe2\xa4\xa3\xa8\xec\xb5|\xa5s\xbd\xe4\xb3\xd5    \x1b[m 118.160.112.18 03/20 09:44\r"),
			CommentID:  "Ff3fu23mwAA:re1zjGmRX94EC6xhYp4stQ",
			CreateTime: 1584668640000000000,
			SortTime:   1584668640000000000,
		},
		{ //1
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("lats"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不然你認為大樂透不算賭博嗎？",
						Big5:   []byte("\xa4\xa3\xb5M\xa7A\xbb{\xac\xb0\xa4j\xbc\xd6\xb3z\xa4\xa3\xba\xe2\xbd\xe4\xb3\xd5\xb6\xdc\xa1H              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa4\xa3\xb5M\xa7A\xbb{\xac\xb0\xa4j\xbc\xd6\xb3z\xa4\xa3\xba\xe2\xbd\xe4\xb3\xd5\xb6\xdc\xa1H              "),
					},
				},
			},
			IP:         "118.160.112.18",
			TheDate:    "03/20 09:45",
			MD5:        "u5v6PWoN_O3ri9oBsiOqFw",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mlats\x1b[m\x1b[33m: \xa4\xa3\xb5M\xa7A\xbb{\xac\xb0\xa4j\xbc\xd6\xb3z\xa4\xa3\xba\xe2\xbd\xe4\xb3\xd5\xb6\xdc\xa1H              \x1b[m 118.160.112.18 03/20 09:45\r"),
			CommentID:  "Ff3fyWYuGAA:u5v6PWoN_O3ri9oBsiOqFw",
			CreateTime: 1584668700000000000,
			SortTime:   1584668700000000000,
		},
		{ //2
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("D1"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "打麻將有股利嗎？",
						Big5:   []byte("\xa5\xb4\xb3\xc2\xb1N\xa6\xb3\xaa\xd1\xa7Q\xb6\xdc\xa1H                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5\xb4\xb3\xc2\xb1N\xa6\xb3\xaa\xd1\xa7Q\xb6\xdc\xa1H                            "),
					},
				},
			},
			MD5:        "nAnF91Ukrs9IefrybKz_HQ",
			IP:         "101.137.108.233",
			TheDate:    "03/20 09:45",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mD1\x1b[m\x1b[33m: \xa5\xb4\xb3\xc2\xb1N\xa6\xb3\xaa\xd1\xa7Q\xb6\xdc\xa1H                            \x1b[m101.137.108.233 03/20 09:45\r"),
			CommentID:  "Ff3fyWY9WkA:nAnF91Ukrs9IefrybKz_HQ",
			CreateTime: 1584668700000000000,
			SortTime:   1584668700001000000,
		},
		{ //3
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("ss1130"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "股票不賣就不賠啊，賭博輸了就是沒了",
						Big5:   []byte("\xaa\xd1\xb2\xbc\xa4\xa3\xbd\xe6\xb4N\xa4\xa3\xbd\xdf\xb0\xda\xa1A\xbd\xe4\xb3\xd5\xbf\xe9\xa4F\xb4N\xacO\xa8S\xa4F      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaa\xd1\xb2\xbc\xa4\xa3\xbd\xe6\xb4N\xa4\xa3\xbd\xdf\xb0\xda\xa1A\xbd\xe4\xb3\xd5\xbf\xe9\xa4F\xb4N\xacO\xa8S\xa4F      "),
					},
				},
			},
			MD5:        "ZxvPgcGyIk2nVe7ZzxdzoA",
			IP:         "114.37.83.201",
			TheDate:    "03/20 09:45",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mss1130\x1b[m\x1b[33m: \xaa\xd1\xb2\xbc\xa4\xa3\xbd\xe6\xb4N\xa4\xa3\xbd\xdf\xb0\xda\xa1A\xbd\xe4\xb3\xd5\xbf\xe9\xa4F\xb4N\xacO\xa8S\xa4F      \x1b[m  114.37.83.201 03/20 09:45\r"),
			CommentID:  "Ff3fyWZMnIA:ZxvPgcGyIk2nVe7ZzxdzoA",
			CreateTime: 1584668700000000000,
			SortTime:   1584668700002000000,
		},
		{ //4
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("jin956"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "開公司做生意有賺有賠 是賭博還是投資?",
						Big5:   []byte("\xb6}\xa4\xbd\xa5q\xb0\xb5\xa5\xcd\xb7N\xa6\xb3\xc1\xc8\xa6\xb3\xbd\xdf \xacO\xbd\xe4\xb3\xd5\xc1\xd9\xacO\xa7\xeb\xb8\xea?    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb6}\xa4\xbd\xa5q\xb0\xb5\xa5\xcd\xb7N\xa6\xb3\xc1\xc8\xa6\xb3\xbd\xdf \xacO\xbd\xe4\xb3\xd5\xc1\xd9\xacO\xa7\xeb\xb8\xea?    "),
					},
				},
			},
			MD5:        "v7aPdDoGIeABPHYvkfe2TA",
			IP:         "110.26.190.156",
			TheDate:    "03/20 09:45",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mjin956\x1b[m\x1b[33m: \xb6}\xa4\xbd\xa5q\xb0\xb5\xa5\xcd\xb7N\xa6\xb3\xc1\xc8\xa6\xb3\xbd\xdf \xacO\xbd\xe4\xb3\xd5\xc1\xd9\xacO\xa7\xeb\xb8\xea?    \x1b[m 110.26.190.156 03/20 09:45\r"),
			CommentID:  "Ff3fyWZb3sA:v7aPdDoGIeABPHYvkfe2TA",
			CreateTime: 1584668700000000000,
			SortTime:   1584668700003000000,
		},
		{ //5
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "股票的意義是 一間公司想生產東西現在沒錢",
						Big5:   []byte("\xaa\xd1\xb2\xbc\xaa\xba\xb7N\xb8q\xacO \xa4@\xb6\xa1\xa4\xbd\xa5q\xb7Q\xa5\xcd\xb2\xa3\xaaF\xa6\xe8\xb2{\xa6b\xa8S\xbf\xfa "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaa\xd1\xb2\xbc\xaa\xba\xb7N\xb8q\xacO \xa4@\xb6\xa1\xa4\xbd\xa5q\xb7Q\xa5\xcd\xb2\xa3\xaaF\xa6\xe8\xb2{\xa6b\xa8S\xbf\xfa "),
					},
				},
			},
			MD5:        "0Z_CmWQ03_875tDRT_AH4Q",
			IP:         "122.118.131.149",
			TheDate:    "03/20 09:48",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xaa\xd1\xb2\xbc\xaa\xba\xb7N\xb8q\xacO \xa4@\xb6\xa1\xa4\xbd\xa5q\xb7Q\xa5\xcd\xb2\xa3\xaaF\xa6\xe8\xb2{\xa6b\xa8S\xbf\xfa \x1b[m122.118.131.149 03/20 09:48\r"),
			CommentID:  "Ff3f808EIAA:0Z_CmWQ03_875tDRT_AH4Q",
			CreateTime: 1584668880000000000,
			SortTime:   1584668880000000000,
		},
		{ //6
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("rindesu"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "四個人打麻將可以，你們一群人打麻將？",
						Big5:   []byte("\xa5|\xad\xd3\xa4H\xa5\xb4\xb3\xc2\xb1N\xa5i\xa5H\xa1A\xa7A\xad\xcc\xa4@\xb8s\xa4H\xa5\xb4\xb3\xc2\xb1N\xa1H   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5|\xad\xd3\xa4H\xa5\xb4\xb3\xc2\xb1N\xa5i\xa5H\xa1A\xa7A\xad\xcc\xa4@\xb8s\xa4H\xa5\xb4\xb3\xc2\xb1N\xa1H   "),
					},
				},
			},
			MD5:        "jd3J52KJCF5KQzt2GErBJg",
			IP:         "110.28.108.240",
			TheDate:    "03/20 09:49",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mrindesu\x1b[m\x1b[33m: \xa5|\xad\xd3\xa4H\xa5\xb4\xb3\xc2\xb1N\xa5i\xa5H\xa1A\xa7A\xad\xcc\xa4@\xb8s\xa4H\xa5\xb4\xb3\xc2\xb1N\xa1H   \x1b[m 110.28.108.240 03/20 09:49\r"),
			CommentID:  "Ff3gAUdLeAA:jd3J52KJCF5KQzt2GErBJg",
			CreateTime: 1584668940000000000,
			SortTime:   1584668940000000000,
		},
		{ //7
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "發行一張紙跟你借錢 你認為這間公司以後會",
						Big5:   []byte("\xb5o\xa6\xe6\xa4@\xb1i\xaf\xc8\xb8\xf2\xa7A\xad\xc9\xbf\xfa \xa7A\xbb{\xac\xb0\xb3o\xb6\xa1\xa4\xbd\xa5q\xa5H\xab\xe1\xb7| "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb5o\xa6\xe6\xa4@\xb1i\xaf\xc8\xb8\xf2\xa7A\xad\xc9\xbf\xfa \xa7A\xbb{\xac\xb0\xb3o\xb6\xa1\xa4\xbd\xa5q\xa5H\xab\xe1\xb7| "),
					},
				},
			},
			MD5:        "I6z7DF7JHmn_RveuhTM9kg",
			IP:         "122.118.131.149",
			TheDate:    "03/20 09:49",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xb5o\xa6\xe6\xa4@\xb1i\xaf\xc8\xb8\xf2\xa7A\xad\xc9\xbf\xfa \xa7A\xbb{\xac\xb0\xb3o\xb6\xa1\xa4\xbd\xa5q\xa5H\xab\xe1\xb7| \x1b[m122.118.131.149 03/20 09:49\r"),
			CommentID:  "Ff3gAUdaukA:I6z7DF7JHmn_RveuhTM9kg",
			CreateTime: 1584668940000000000,
			SortTime:   1584668940001000000,
		},
		{ //8
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("vowpool"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "應該用賽馬跟股市來做類比 比較像",
						Big5:   []byte("\xc0\xb3\xb8\xd3\xa5\xce\xc1\xc9\xb0\xa8\xb8\xf2\xaa\xd1\xa5\xab\xa8\xd3\xb0\xb5\xc3\xfe\xa4\xf1 \xa4\xf1\xb8\xfb\xb9\xb3        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc0\xb3\xb8\xd3\xa5\xce\xc1\xc9\xb0\xa8\xb8\xf2\xaa\xd1\xa5\xab\xa8\xd3\xb0\xb5\xc3\xfe\xa4\xf1 \xa4\xf1\xb8\xfb\xb9\xb3        "),
					},
				},
			},
			MD5:        "HxFf1L0eaMiqE3wyJ0Nyow",
			IP:         "125.227.40.62",
			TheDate:    "03/20 09:50",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mvowpool\x1b[m\x1b[33m: \xc0\xb3\xb8\xd3\xa5\xce\xc1\xc9\xb0\xa8\xb8\xf2\xaa\xd1\xa5\xab\xa8\xd3\xb0\xb5\xc3\xfe\xa4\xf1 \xa4\xf1\xb8\xfb\xb9\xb3        \x1b[m  125.227.40.62 03/20 09:50\r"),
			CommentID:  "Ff3gDz-S0AA:HxFf1L0eaMiqE3wyJ0Nyow",
			CreateTime: 1584669000000000000,
			SortTime:   1584669000000000000,
		},
		{ //9
			BBoardID:  "test",
			ArticleID: "test18",
			TheType:   types.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("vowpool"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "賽馬是賭博 股市是金融",
						Big5:   []byte("\xc1\xc9\xb0\xa8\xacO\xbd\xe4\xb3\xd5 \xaa\xd1\xa5\xab\xacO\xaa\xf7\xbf\xc4                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc9\xb0\xa8\xacO\xbd\xe4\xb3\xd5 \xaa\xd1\xa5\xab\xacO\xaa\xf7\xbf\xc4                  "),
					},
				},
			},
			MD5:        "KZhYeiENlhq_U1YasuvFlQ",
			IP:         "125.227.40.62",
			TheDate:    "03/20 09:50",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mvowpool\x1b[m\x1b[33m: \xc1\xc9\xb0\xa8\xacO\xbd\xe4\xb3\xd5 \xaa\xd1\xa5\xab\xacO\xaa\xf7\xbf\xc4                  \x1b[m  125.227.40.62 03/20 09:50\r"),
			CommentID:  "Ff3gDz-iEkA:KZhYeiENlhq_U1YasuvFlQ",
			CreateTime: 1584669000000000000,
			SortTime:   1584669000001000000,
		},
	}

	testTheRestComments18 = []*schema.Comment{
		{ //0
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "賺 於是你借他錢 後來他真的賺了 你就可以",
						Big5:   []byte("\xc1\xc8 \xa9\xf3\xacO\xa7A\xad\xc9\xa5L\xbf\xfa \xab\xe1\xa8\xd3\xa5L\xafu\xaa\xba\xc1\xc8\xa4F \xa7A\xb4N\xa5i\xa5H "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xc1\xc8 \xa9\xf3\xacO\xa7A\xad\xc9\xa5L\xbf\xfa \xab\xe1\xa8\xd3\xa5L\xafu\xaa\xba\xc1\xc8\xa4F \xa7A\xb4N\xa5i\xa5H "),
					},
				},
			},
			MD5:     "hxSTT_5bOL5jCmWV5vXs2g",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:50",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xc1\xc8 \xa9\xf3\xacO\xa7A\xad\xc9\xa5L\xbf\xfa \xab\xe1\xa8\xd3\xa5L\xafu\xaa\xba\xc1\xc8\xa4F \xa7A\xb4N\xa5i\xa5H \x1b[m122.118.131.149 03/20 09:50\r"),
		},
		{ //1
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Besorgen"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "叔叔我看過打麻將，好像也有募資行為。",
						Big5:   []byte("\xa8\xfb\xa8\xfb\xa7\xda\xac\xdd\xb9L\xa5\xb4\xb3\xc2\xb1N\xa1A\xa6n\xb9\xb3\xa4]\xa6\xb3\xb6\xd2\xb8\xea\xa6\xe6\xac\xb0\xa1C  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa8\xfb\xa8\xfb\xa7\xda\xac\xdd\xb9L\xa5\xb4\xb3\xc2\xb1N\xa1A\xa6n\xb9\xb3\xa4]\xa6\xb3\xb6\xd2\xb8\xea\xa6\xe6\xac\xb0\xa1C  "),
					},
				},
			},
			MD5:     "hx_riWhj_HuLg0UIlBSeuQ",
			IP:      "125.230.35.42",
			TheDate: "03/20 09:51",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mBesorgen\x1b[m\x1b[33m: \xa8\xfb\xa8\xfb\xa7\xda\xac\xdd\xb9L\xa5\xb4\xb3\xc2\xb1N\xa1A\xa6n\xb9\xb3\xa4]\xa6\xb3\xb6\xd2\xb8\xea\xa6\xe6\xac\xb0\xa1C  \x1b[m  125.230.35.42 03/20 09:51\r"),
		},
		{ //2
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "得到本金 以及利息",
						Big5:   []byte("\xb1o\xa8\xec\xa5\xbb\xaa\xf7 \xa5H\xa4\xce\xa7Q\xae\xa7                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1o\xa8\xec\xa5\xbb\xaa\xf7 \xa5H\xa4\xce\xa7Q\xae\xa7                       "),
					},
				},
			},
			MD5:     "9_hUpOAIUE1HQcxhmr5YcQ",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:51",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xb1o\xa8\xec\xa5\xbb\xaa\xf7 \xa5H\xa4\xce\xa7Q\xae\xa7                       \x1b[m122.118.131.149 03/20 09:51\r"),
		},
		{ //3
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("kumori"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "拿回",
						Big5:   []byte("\xae\xb3\xa6^                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xae\xb3\xa6^                                    "),
					},
				},
			},
			MD5:     "Bkl1I9Dexr0692-6nBt7PQ",
			IP:      "122.118.131.149",
			TheDate: "03/20 09:52",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mkumori\x1b[m\x1b[33m: \xae\xb3\xa6^                                    \x1b[m122.118.131.149 03/20 09:52\r"),
		},
		{ //4
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("cc03233"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "重點 看誰做莊",
						Big5:   []byte("\xad\xab\xc2I \xac\xdd\xbd\xd6\xb0\xb5\xb2\xf8                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xad\xab\xc2I \xac\xdd\xbd\xd6\xb0\xb5\xb2\xf8                          "),
					},
				},
			},
			MD5:     "my7ntqlgU9SpujIvhhpV9g",
			IP:      "1.161.225.69",
			TheDate: "03/20 09:53",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mcc03233\x1b[m\x1b[33m: \xad\xab\xc2I \xac\xdd\xbd\xd6\xb0\xb5\xb2\xf8                          \x1b[m   1.161.225.69 03/20 09:53\r"),
		},
		{ //5
			TheType: types.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("RuleAllWorld"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "你先找出一副有成長性的麻將給我看",
						Big5:   []byte("\xa7A\xa5\xfd\xa7\xe4\xa5X\xa4@\xb0\xc6\xa6\xb3\xa6\xa8\xaa\xf8\xa9\xca\xaa\xba\xb3\xc2\xb1N\xb5\xb9\xa7\xda\xac\xdd  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa7A\xa5\xfd\xa7\xe4\xa5X\xa4@\xb0\xc6\xa6\xb3\xa6\xa8\xaa\xf8\xa9\xca\xaa\xba\xb3\xc2\xb1N\xb5\xb9\xa7\xda\xac\xdd  "),
					},
				},
			},
			MD5:     "JHZlo-XfqbkQjxtH4jwk4g",
			IP:      "49.216.4.199",
			TheDate: "03/20 09:54",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mRuleAllWorld\x1b[m\x1b[33m: \xa7A\xa5\xfd\xa7\xe4\xa5X\xa4@\xb0\xc6\xa6\xb3\xa6\xa8\xaa\xf8\xa9\xca\xaa\xba\xb3\xc2\xb1N\xb5\xb9\xa7\xda\xac\xdd  \x1b[m   49.216.4.199 03/20 09:54\r"),
		},
		{ //6
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("DinoZavolta"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "台灣人多的是把股市當賭場玩的",
						Big5:   []byte("\xa5x\xc6W\xa4H\xa6h\xaa\xba\xacO\xa7\xe2\xaa\xd1\xa5\xab\xb7\xed\xbd\xe4\xb3\xf5\xaa\xb1\xaa\xba       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5x\xc6W\xa4H\xa6h\xaa\xba\xacO\xa7\xe2\xaa\xd1\xa5\xab\xb7\xed\xbd\xe4\xb3\xf5\xaa\xb1\xaa\xba       "),
					},
				},
			},
			MD5:     "wSG9KCFXJmDKZqu4nH962Q",
			IP:      "112.105.78.63",
			TheDate: "03/20 10:27",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mDinoZavolta\x1b[m\x1b[33m: \xa5x\xc6W\xa4H\xa6h\xaa\xba\xacO\xa7\xe2\xaa\xd1\xa5\xab\xb7\xed\xbd\xe4\xb3\xf5\xaa\xb1\xaa\xba       \x1b[m  112.105.78.63 03/20 10:27\r"),
		},
		{ //7
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("oyui111"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "因為大家把股票當麻將玩",
						Big5:   []byte("\xa6]\xac\xb0\xa4j\xaea\xa7\xe2\xaa\xd1\xb2\xbc\xb7\xed\xb3\xc2\xb1N\xaa\xb1                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa6]\xac\xb0\xa4j\xaea\xa7\xe2\xaa\xd1\xb2\xbc\xb7\xed\xb3\xc2\xb1N\xaa\xb1                 "),
					},
				},
			},
			MD5:     "Boqz1QHBlDpR3u9t2EUMag",
			IP:      "223.141.120.153",
			TheDate: "03/20 10:29",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33moyui111\x1b[m\x1b[33m: \xa6]\xac\xb0\xa4j\xaea\xa7\xe2\xaa\xd1\xb2\xbc\xb7\xed\xb3\xc2\xb1N\xaa\xb1                 \x1b[m223.141.120.153 03/20 10:29\r"),
		},
		{ //8
			TheType: types.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("jackeman"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "稅稅稅稅稅",
						Big5:   []byte("\xb5|\xb5|\xb5|\xb5|\xb5|                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb5|\xb5|\xb5|\xb5|\xb5|                            "),
					},
				},
			},
			MD5:     "YGqJ97m7yvJxtp5p7ylxPA",
			IP:      "61.223.69.223",
			TheDate: "03/20 10:33",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mjackeman\x1b[m\x1b[33m: \xb5|\xb5|\xb5|\xb5|\xb5|                            \x1b[m  61.223.69.223 03/20 10:33\r"),
		},
		{ //9
			TheType: types.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("tomxyz"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "症腐做莊的都合法喔 科科",
						Big5:   []byte("\xafg\xbbG\xb0\xb5\xb2\xf8\xaa\xba\xb3\xa3\xa6X\xaak\xb3\xe1 \xac\xec\xac\xec                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xafg\xbbG\xb0\xb5\xb2\xf8\xaa\xba\xb3\xa3\xa6X\xaak\xb3\xe1 \xac\xec\xac\xec                 "),
					},
				},
			},
			MD5:     "-_eTj8tTDmfX7iSxnqvURA",
			IP:      "42.72.59.33",
			TheDate: "03/20 10:57",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mtomxyz\x1b[m\x1b[33m: \xafg\xbbG\xb0\xb5\xb2\xf8\xaa\xba\xb3\xa3\xa6X\xaak\xb3\xe1 \xac\xec\xac\xec                 \x1b[m    42.72.59.33 03/20 10:57\r"),
		},
		{ //10
			TheType: types.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("bryanma"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "國家開的啊",
						Big5:   []byte("\xb0\xea\xaea\xb6}\xaa\xba\xb0\xda                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb0\xea\xaea\xb6}\xaa\xba\xb0\xda                             "),
					},
				},
			},
			MD5:     "-vyYfu3vDRxZT2HYvjERFw",
			IP:      "223.136.131.5",
			TheDate: "03/21 18:36",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mbryanma\x1b[m\x1b[33m: \xb0\xea\xaea\xb6}\xaa\xba\xb0\xda                             \x1b[m  223.136.131.5 03/21 18:36\r"),
		},
	}
}
