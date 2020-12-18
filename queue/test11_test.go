package queue

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename11            = "temp6"
	testContentAll11          []byte
	testContent11             []byte
	testSignature11           []byte
	testComment11             []byte
	testFirstCommentsDBCS11   []byte
	testTheRestCommentsDBCS11 []byte

	testContent11Big5 [][]*types.Rune
	testContent11Utf8 [][]*types.Rune

	testFirstComments11 []*schema.Comment

	testTheRestComments11 []*schema.Comment
)

func initTest11() {
	testContentAll11, testContent11, testSignature11, testComment11, testFirstCommentsDBCS11, testTheRestCommentsDBCS11 = loadTest(testFilename11)

	testContent11Big5 = [][]*types.Rune{
		{ //0
			{
				Big5:   []byte("\xa7@\xaa\xcc: cheinshin (\xa8\xba\xb4N\xb3o\xbc\xcb\xa7a) \xac\xdd\xaaO: Gossiping"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //1
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xb7s\xbbD] TVBS\xa4\xbb\xb3\xa3\xa5\xc1\xbd\xd5 \xabJ\xb9\xdc\xaba\xa1B\xbfc\xa4\xc9\xb2\xc4\xa5|\xa1B\xac_\xb9\xd4\xa9\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //2
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Dec 21 19:45:20 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //3
		{}, //4
		{}, //5
		{}, //6
		{ //7
			{
				Big5:   []byte("1.\xb4C\xc5\xe9\xa8\xd3\xb7\xbd:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //8
		{ //9
			{
				Big5:   []byte("TVBS"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //10
		{ //11
			{
				Big5:   []byte("2.\xb0O\xaa\xcc\xb8p\xa6W:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //12
		{ //13
			{
				Big5:   []byte("\xad\xb3\xaea\xbb\xf4"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //14
		{ //15
			{
				Big5:   []byte("3.\xa7\xb9\xbe\xe3\xb7s\xbbD\xbc\xd0\xc3D:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //16
		{ //17
			{
				Big5:   []byte("TVBS\xa4\xbb\xb3\xa3\xa5\xc1\xbd\xd5 \xabJ\xb9\xdc\xaba\xa1B\xbfc\xa4\xc9\xb2\xc4\xa5|\xa1B\xac_\xb9\xd4\xa9\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //18
		{ //19
			{
				Big5:   []byte("4.\xa7\xb9\xbe\xe3\xb7s\xbbD\xa4\xba\xa4\xe5:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //20
		{}, //21
		{ //22
			{
				Big5:   []byte("2022\xbf\xef\xbe\xd4\xa7Y\xb1N\xb6}\xa5\xb4\xa1A\xa5\xc1\xb2\xb3\xa4]\xb9\xef\xa4\xbb\xb3\xa3\xa5\xab\xaa\xf8\xaa\xba\xacI\xacF\xaa\xed\xb2{\xa5\xb4\xa4F\xa4\xc0\xbc\xc6\xa1A\xae\xda\xbe\xdaTVBS\xb3\xcc\xb7s\xa5\xc1\xbd\xd5\xc5\xe3\xa5\xdc\xa1A\xa5x"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //23
			{
				Big5:   []byte("\xa4\xa4\xa5\xab\xaa\xf8\xbfc\xa8q\xbfP\xac\xf0\xc5\xa7AIT\xb3B\xaa\xf8\xaa\xed\xb9F\xa4\xcf\xb5\xdc\xbd\xde\xa5\xdf\xb3\xf5\xab\xe1\xa1A\xa4H\xae\xf0\xa4W\xa4\xc9\xa1A\xa4\xbb\xb3\xa3\xb1\xc6\xa6W\xb2\xc4\xa5|\xa1A\xa6\xdc\xa9\xf3\xa5h\xa6~\xa4~\xb8\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //24
			{
				Big5:   []byte("\xbf\xef\xa4W\xaa\xba\xb0\xaa\xb6\xaf\xa5\xab\xaa\xf8\xb3\xaf\xa8\xe4\xc1\xda\xa1A\xa5u\xae\xb3\xa4U\xb2\xc4\xa4\xad\xa6W\xa1A\xa5u\xc4\xb9\xa4F\xb9\xd4\xa9\xb3\xaa\xba\xa5x\xa5_\xa5\xab\xaa\xf8\xac_\xa4\xe5\xad\xf5\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //25
		{ //26
			{
				Big5:   []byte("\xa5x\xa4\xa4\xa5\xab\xaa\xf8\xbfc\xa8q\xbfP\xa1G\xa1u\xa4\xa3\xa7\xc6\xb1\xe6\xa7\xda\xad\xcc\xb8\xd1\xb8T\xa6\xb3\xc3\xf6\xa9\xf3\xa1A\xa7t\xa6\xb3\xb5\xdc\xa7J\xa6h\xa4\xda\xd3i\xa1A\xa9\xce\xbdG\xa6\xd7\xba\xeb\xaa\xba\xbd\xde\xa6\xd7\xb6i\xa4f\xa1C\xa1v"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //27
		{ //28
			{
				Big5:   []byte("\xb4N\xacO\xb3o\xbb\xf2\xb5L\xb9w\xc4\xb5\xac\xf0\xc5\xa7AIT\xb3B\xaa\xf8\xa1A\xaa\xed\xb9F\xa4\xcf\xb5\xdc\xbd\xde\xa5\xdf\xb3\xf5\xa1A\xc5\xfd\xa5x\xa4\xa4\xa5\xab\xaa\xf8\xbfc\xa8q\xbfP\xa4H\xae\xf0\xf6t\xa4\xc9\xa1A\xa4]\xa4\xcf\xc0\xb3\xa8\xec"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //29
			{
				Big5:   []byte("\xa6o\xaa\xba\xa5\xc1\xbd\xd5\xa4W\xad\xb1\xa1A\xae\xda\xbe\xdaTVBS\xb3\xcc\xb7s\xa5\xc1\xbd\xd5\xc5\xe3\xa5\xdc\xa1A\xbfc\xa8q\xbfP\xaa\xba\xacI\xacF\xba\xa1\xb7N\xab\xd7\xf6t\xa8\xec56%\xa1A\xa6\xec\xa9~\xa4\xbb\xb3\xa3\xb2\xc4\xa5|\xa6W"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //30
			{
				Big5:   []byte("\xa1A\xbb\xb7\xbb\xb7\xbb\xe2\xa5\xfd\xb3\xcc\xab\xe1\xa4@\xa6W\xaa\xba\xa1A\xa5x\xa5_\xa5\xab\xaa\xf8\xac_\xa4\xe5\xad\xf535%\xaa\xba\xa4\xe4\xab\xf9\xab\xd7\xa1A\xb8\xf2\xb0\xaa\xb6\xaf\xa5\xab\xaa\xf8\xb3\xaf\xa8\xe4\xc1\xda\xaa\xba44%\xa1A\xa5x\xabn\xa5\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //31
			{
				Big5:   []byte("\xaa\xf8\xb6\xc0\xb0\xb6\xad\xf5\xba\xa1\xb7N\xab\xd7\xa4\xf1\xa5h\xa6~\xa1A\xa6h\xa4F3\xad\xd3\xa6\xca\xa4\xc0\xc2I\xa8\xd3\xa8\xec58%\xa1A\xae\xb3\xa4U\xb2\xc4\xa4T\xa6W\xa1A\xb2\xc4\xa4G\xa6W\xabh\xacO\xae\xe7\xb6\xe9\xa5\xab\xaa\xf8\xbeG\xa4\xe5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //32
			{
				Big5:   []byte("\xc0\xe9\xa1A\xae\xb3\xa4U\xa4C\xa6\xa8\xaa\xba\xba\xa1\xb7N\xab\xd7\xa1A\xaba\xadx\xabh\xacO\xa5\xd1\xb1`\xb3\xd3\xadx\xaa\xba\xa1A\xb7s\xa5_\xa5\xab\xaa\xf8\xabJ\xa4\xcd\xa9y\xa5H77%\xa6A\xab\xd7\xc2\xcd\xc1p\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //33
		{ //34
			{
				Big5:   []byte("\xb7s\xa5_\xa5\xab\xaa\xf8\xabJ\xa4\xcd\xa9y\xa1G\xa1u\xa7\xda\xb3\xcc\xa6b\xa5G\xaa\xba\xacO\xa1A\xa7\xda\xc1\xd9\xa6\xb3\xab\xdc\xa6h\xa8\xc6\xa8S\xa6\xb3\xb3\xcc\xa6n\xaa\xba\xb3\xa1\xa4\xc0\xa1A\xadn\xa7\xf3\xa5[\xa7\xe2\xabl\xa1A\xa7\xe2\xa5\xa6\xb0\xb5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //35
			{
				Big5:   []byte("\xb1o\xa7\xf3\xa6n\xa1A\xb3o\xa4~\xacO\xa5\xab\xaa\xf8\xaa\xba\xa5\xbb\xb3d\xa1C\xa1v"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //36
		{ //37
			{
				Big5:   []byte("\xb1j\xbd\xd5\xa6\xdb\xa4v\xb4N\xacO\xaeI\xc0Y\xbb{\xafu\xb0\xb5\xa1A\xbe\xa8\xba\xde\xab\xb0\xa5\xab\xa5\xfa\xbaa\xb7P\xb3\xa1\xa4\xc0\xa1A\xac\xbd\xa5\xcf\xb0\xaa\xb6\xaf\xb8\xf2\xb9\xd4\xa9\xb3\xaa\xba\xa5x\xa5_\xa1A\xc5\xfd\xa4C\xa6\xa8\xaa\xba\xb7s\xa5_"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //38
			{
				Big5:   []byte("\xa5\xab\xa5\xc1\xb7P\xa8\xec\xa5\xfa\xbaa\xa1A\xa6\xfd\xae\xe7\xb6\xe9\xb8\xf2\xa5x\xabn\xaa\xba\xa5\xab\xa5\xc1\xa1A\xb5w\xacO\xa6h\xa5X\xa4F2\xad\xd3\xa6\xca\xa4\xc0\xc2I\xa1A\xa6\xd3\xa6\xb3\xa4K\xa6\xa8\xaa\xba\xa5x\xa4\xa4\xa5\xab\xa5\xc1\xbb{\xa6P"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //39
			{
				Big5:   []byte("\xa5x\xa4\xa4\xa1A\xbaa\xb5n\xa4\xbb\xb3\xa3\xab\xb0\xa5\xab\xa5\xfa\xbaa\xb7P\xb3\xcc\xb0\xaa\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //40
		{ //41
			{
				Big5:   []byte("\xa5x\xabn\xa5\xab\xaa\xf8\xb6\xc0\xb0\xb6\xad\xf5\xa1G\xa1u\xa6b\xbb\xdd\xadn\xc0\xcb\xb0Q\xa7\xef\xb6i\xaa\xba\xa6a\xa4\xe8\xa1A\xa7\xda\xad\xcc\xa4]\xb7|\xb5\xea\xa4\xdf\xa6a\xc0\xcb\xb0Q\xa5H\xa4\xce\xa7\xef\xb6i\xa1A\xa7\xc6\xb1\xe6\xb0\xb5\xb1o\xa7\xf3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //42
			{
				Big5:   []byte("\xa6n\xa1A\xc5\xfd\xa5\xab\xa5\xc1\xaaB\xa4\xcd\xad\xcc\xb9L\xb1o\xa7\xf3\xa6n\xa1C\xa1v"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //43
		{}, //44
		{ //45
			{
				Big5:   []byte("5.\xa7\xb9\xbe\xe3\xb7s\xbbD\xb3s\xb5\xb2 (\xa9\xce\xb5u\xba\xf4\xa7}):"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //46
		{ //47
			{
				Big5:   []byte("https://reurl.cc/8nK7eo"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //48
		{ //49
			{
				Big5:   []byte("6.\xb3\xc6\xb5\xf9:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //50
		{ //51
			{
				Big5:   []byte("1.\xabJ 2.\xbeG 3.\xb6\xc0 4.\xbfc 5.\xb3\xaf 6.\xac_"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //52
		{ //53
			{
				Big5:   []byte("\xaf\xf3\xa5]\xbau\xa4F \xac_\xaaG\xb5M\xb9\xd4\xa9\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //54
		{}, //55
		{}, //56
		{ //51
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	testContent11Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: cheinshin (那就這樣吧) 看板: Gossiping",
				Big5:   []byte("\xa7@\xaa\xcc: cheinshin (\xa8\xba\xb4N\xb3o\xbc\xcb\xa7a) \xac\xdd\xaaO: Gossiping"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //1
			{
				Utf8:   "標題: [新聞] TVBS六都民調 侯奪冠、盧升第四、柯墊底",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb7s\xbbD] TVBS\xa4\xbb\xb3\xa3\xa5\xc1\xbd\xd5 \xabJ\xb9\xdc\xaba\xa1B\xbfc\xa4\xc9\xb2\xc4\xa5|\xa1B\xac_\xb9\xd4\xa9\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //2
			{
				Utf8:   "時間: Mon Dec 21 19:45:20 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Dec 21 19:45:20 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //3
		{}, //4
		{}, //5
		{}, //6
		{ //7
			{
				Utf8:   "1.媒體來源:",
				Big5:   []byte("1.\xb4C\xc5\xe9\xa8\xd3\xb7\xbd:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //8
		{ //9
			{
				Utf8:   "TVBS",
				Big5:   []byte("TVBS"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //10
		{ //11
			{
				Utf8:   "2.記者署名:",
				Big5:   []byte("2.\xb0O\xaa\xcc\xb8p\xa6W:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //12
		{ //13
			{
				Utf8:   "韋家齊",
				Big5:   []byte("\xad\xb3\xaea\xbb\xf4"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //14
		{ //15
			{
				Utf8:   "3.完整新聞標題:",
				Big5:   []byte("3.\xa7\xb9\xbe\xe3\xb7s\xbbD\xbc\xd0\xc3D:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //16
		{ //17
			{
				Utf8:   "TVBS六都民調 侯奪冠、盧升第四、柯墊底",
				Big5:   []byte("TVBS\xa4\xbb\xb3\xa3\xa5\xc1\xbd\xd5 \xabJ\xb9\xdc\xaba\xa1B\xbfc\xa4\xc9\xb2\xc4\xa5|\xa1B\xac_\xb9\xd4\xa9\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //18
		{ //19
			{
				Utf8:   "4.完整新聞內文:",
				Big5:   []byte("4.\xa7\xb9\xbe\xe3\xb7s\xbbD\xa4\xba\xa4\xe5:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //20
		{}, //21
		{ //22
			{
				Utf8:   "2022選戰即將開打，民眾也對六都市長的施政表現打了分數，根據TVBS最新民調顯示，台",
				Big5:   []byte("2022\xbf\xef\xbe\xd4\xa7Y\xb1N\xb6}\xa5\xb4\xa1A\xa5\xc1\xb2\xb3\xa4]\xb9\xef\xa4\xbb\xb3\xa3\xa5\xab\xaa\xf8\xaa\xba\xacI\xacF\xaa\xed\xb2{\xa5\xb4\xa4F\xa4\xc0\xbc\xc6\xa1A\xae\xda\xbe\xdaTVBS\xb3\xcc\xb7s\xa5\xc1\xbd\xd5\xc5\xe3\xa5\xdc\xa1A\xa5x"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //23
			{
				Utf8:   "中市長盧秀燕突襲AIT處長表達反萊豬立場後，人氣上升，六都排名第四，至於去年才補",
				Big5:   []byte("\xa4\xa4\xa5\xab\xaa\xf8\xbfc\xa8q\xbfP\xac\xf0\xc5\xa7AIT\xb3B\xaa\xf8\xaa\xed\xb9F\xa4\xcf\xb5\xdc\xbd\xde\xa5\xdf\xb3\xf5\xab\xe1\xa1A\xa4H\xae\xf0\xa4W\xa4\xc9\xa1A\xa4\xbb\xb3\xa3\xb1\xc6\xa6W\xb2\xc4\xa5|\xa1A\xa6\xdc\xa9\xf3\xa5h\xa6~\xa4~\xb8\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //24
			{
				Utf8:   "選上的高雄市長陳其邁，只拿下第五名，只贏了墊底的台北市長柯文哲。",
				Big5:   []byte("\xbf\xef\xa4W\xaa\xba\xb0\xaa\xb6\xaf\xa5\xab\xaa\xf8\xb3\xaf\xa8\xe4\xc1\xda\xa1A\xa5u\xae\xb3\xa4U\xb2\xc4\xa4\xad\xa6W\xa1A\xa5u\xc4\xb9\xa4F\xb9\xd4\xa9\xb3\xaa\xba\xa5x\xa5_\xa5\xab\xaa\xf8\xac_\xa4\xe5\xad\xf5\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //25
		{ //26
			{
				Utf8:   "台中市長盧秀燕：「不希望我們解禁有關於，含有萊克多巴胺，或瘦肉精的豬肉進口。」",
				Big5:   []byte("\xa5x\xa4\xa4\xa5\xab\xaa\xf8\xbfc\xa8q\xbfP\xa1G\xa1u\xa4\xa3\xa7\xc6\xb1\xe6\xa7\xda\xad\xcc\xb8\xd1\xb8T\xa6\xb3\xc3\xf6\xa9\xf3\xa1A\xa7t\xa6\xb3\xb5\xdc\xa7J\xa6h\xa4\xda\xd3i\xa1A\xa9\xce\xbdG\xa6\xd7\xba\xeb\xaa\xba\xbd\xde\xa6\xd7\xb6i\xa4f\xa1C\xa1v"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //27
		{ //28
			{
				Utf8:   "就是這麼無預警突襲AIT處長，表達反萊豬立場，讓台中市長盧秀燕人氣飆升，也反應到",
				Big5:   []byte("\xb4N\xacO\xb3o\xbb\xf2\xb5L\xb9w\xc4\xb5\xac\xf0\xc5\xa7AIT\xb3B\xaa\xf8\xa1A\xaa\xed\xb9F\xa4\xcf\xb5\xdc\xbd\xde\xa5\xdf\xb3\xf5\xa1A\xc5\xfd\xa5x\xa4\xa4\xa5\xab\xaa\xf8\xbfc\xa8q\xbfP\xa4H\xae\xf0\xf6t\xa4\xc9\xa1A\xa4]\xa4\xcf\xc0\xb3\xa8\xec"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //29
			{
				Utf8:   "她的民調上面，根據TVBS最新民調顯示，盧秀燕的施政滿意度飆到56%，位居六都第四名",
				Big5:   []byte("\xa6o\xaa\xba\xa5\xc1\xbd\xd5\xa4W\xad\xb1\xa1A\xae\xda\xbe\xdaTVBS\xb3\xcc\xb7s\xa5\xc1\xbd\xd5\xc5\xe3\xa5\xdc\xa1A\xbfc\xa8q\xbfP\xaa\xba\xacI\xacF\xba\xa1\xb7N\xab\xd7\xf6t\xa8\xec56%\xa1A\xa6\xec\xa9~\xa4\xbb\xb3\xa3\xb2\xc4\xa5|\xa6W"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //30
			{
				Utf8:   "，遠遠領先最後一名的，台北市長柯文哲35%的支持度，跟高雄市長陳其邁的44%，台南市",
				Big5:   []byte("\xa1A\xbb\xb7\xbb\xb7\xbb\xe2\xa5\xfd\xb3\xcc\xab\xe1\xa4@\xa6W\xaa\xba\xa1A\xa5x\xa5_\xa5\xab\xaa\xf8\xac_\xa4\xe5\xad\xf535%\xaa\xba\xa4\xe4\xab\xf9\xab\xd7\xa1A\xb8\xf2\xb0\xaa\xb6\xaf\xa5\xab\xaa\xf8\xb3\xaf\xa8\xe4\xc1\xda\xaa\xba44%\xa1A\xa5x\xabn\xa5\xab"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //31
			{
				Utf8:   "長黃偉哲滿意度比去年，多了3個百分點來到58%，拿下第三名，第二名則是桃園市長鄭文",
				Big5:   []byte("\xaa\xf8\xb6\xc0\xb0\xb6\xad\xf5\xba\xa1\xb7N\xab\xd7\xa4\xf1\xa5h\xa6~\xa1A\xa6h\xa4F3\xad\xd3\xa6\xca\xa4\xc0\xc2I\xa8\xd3\xa8\xec58%\xa1A\xae\xb3\xa4U\xb2\xc4\xa4T\xa6W\xa1A\xb2\xc4\xa4G\xa6W\xabh\xacO\xae\xe7\xb6\xe9\xa5\xab\xaa\xf8\xbeG\xa4\xe5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //32
			{
				Utf8:   "燦，拿下七成的滿意度，冠軍則是由常勝軍的，新北市長侯友宜以77%再度蟬聯。",
				Big5:   []byte("\xc0\xe9\xa1A\xae\xb3\xa4U\xa4C\xa6\xa8\xaa\xba\xba\xa1\xb7N\xab\xd7\xa1A\xaba\xadx\xabh\xacO\xa5\xd1\xb1`\xb3\xd3\xadx\xaa\xba\xa1A\xb7s\xa5_\xa5\xab\xaa\xf8\xabJ\xa4\xcd\xa9y\xa5H77%\xa6A\xab\xd7\xc2\xcd\xc1p\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //33
		{ //34
			{
				Utf8:   "新北市長侯友宜：「我最在乎的是，我還有很多事沒有最好的部分，要更加把勁，把它做",
				Big5:   []byte("\xb7s\xa5_\xa5\xab\xaa\xf8\xabJ\xa4\xcd\xa9y\xa1G\xa1u\xa7\xda\xb3\xcc\xa6b\xa5G\xaa\xba\xacO\xa1A\xa7\xda\xc1\xd9\xa6\xb3\xab\xdc\xa6h\xa8\xc6\xa8S\xa6\xb3\xb3\xcc\xa6n\xaa\xba\xb3\xa1\xa4\xc0\xa1A\xadn\xa7\xf3\xa5[\xa7\xe2\xabl\xa1A\xa7\xe2\xa5\xa6\xb0\xb5"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //35
			{
				Utf8:   "得更好，這才是市長的本責。」",
				Big5:   []byte("\xb1o\xa7\xf3\xa6n\xa1A\xb3o\xa4~\xacO\xa5\xab\xaa\xf8\xaa\xba\xa5\xbb\xb3d\xa1C\xa1v"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //36
		{ //37
			{
				Utf8:   "強調自己就是埋頭認真做，儘管城市光榮感部分，狠甩高雄跟墊底的台北，讓七成的新北",
				Big5:   []byte("\xb1j\xbd\xd5\xa6\xdb\xa4v\xb4N\xacO\xaeI\xc0Y\xbb{\xafu\xb0\xb5\xa1A\xbe\xa8\xba\xde\xab\xb0\xa5\xab\xa5\xfa\xbaa\xb7P\xb3\xa1\xa4\xc0\xa1A\xac\xbd\xa5\xcf\xb0\xaa\xb6\xaf\xb8\xf2\xb9\xd4\xa9\xb3\xaa\xba\xa5x\xa5_\xa1A\xc5\xfd\xa4C\xa6\xa8\xaa\xba\xb7s\xa5_"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //38
			{
				Utf8:   "市民感到光榮，但桃園跟台南的市民，硬是多出了2個百分點，而有八成的台中市民認同",
				Big5:   []byte("\xa5\xab\xa5\xc1\xb7P\xa8\xec\xa5\xfa\xbaa\xa1A\xa6\xfd\xae\xe7\xb6\xe9\xb8\xf2\xa5x\xabn\xaa\xba\xa5\xab\xa5\xc1\xa1A\xb5w\xacO\xa6h\xa5X\xa4F2\xad\xd3\xa6\xca\xa4\xc0\xc2I\xa1A\xa6\xd3\xa6\xb3\xa4K\xa6\xa8\xaa\xba\xa5x\xa4\xa4\xa5\xab\xa5\xc1\xbb{\xa6P"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //39
			{
				Utf8:   "台中，榮登六都城市光榮感最高。",
				Big5:   []byte("\xa5x\xa4\xa4\xa1A\xbaa\xb5n\xa4\xbb\xb3\xa3\xab\xb0\xa5\xab\xa5\xfa\xbaa\xb7P\xb3\xcc\xb0\xaa\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //40
		{ //41
			{
				Utf8:   "台南市長黃偉哲：「在需要檢討改進的地方，我們也會虛心地檢討以及改進，希望做得更",
				Big5:   []byte("\xa5x\xabn\xa5\xab\xaa\xf8\xb6\xc0\xb0\xb6\xad\xf5\xa1G\xa1u\xa6b\xbb\xdd\xadn\xc0\xcb\xb0Q\xa7\xef\xb6i\xaa\xba\xa6a\xa4\xe8\xa1A\xa7\xda\xad\xcc\xa4]\xb7|\xb5\xea\xa4\xdf\xa6a\xc0\xcb\xb0Q\xa5H\xa4\xce\xa7\xef\xb6i\xa1A\xa7\xc6\xb1\xe6\xb0\xb5\xb1o\xa7\xf3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{ //42
			{
				Utf8:   "好，讓市民朋友們過得更好。」",
				Big5:   []byte("\xa6n\xa1A\xc5\xfd\xa5\xab\xa5\xc1\xaaB\xa4\xcd\xad\xcc\xb9L\xb1o\xa7\xf3\xa6n\xa1C\xa1v"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //43
		{}, //44
		{ //45
			{
				Utf8:   "5.完整新聞連結 (或短網址):",
				Big5:   []byte("5.\xa7\xb9\xbe\xe3\xb7s\xbbD\xb3s\xb5\xb2 (\xa9\xce\xb5u\xba\xf4\xa7}):"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //46
		{ //47
			{
				Utf8:   "https://reurl.cc/8nK7eo",
				Big5:   []byte("https://reurl.cc/8nK7eo"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //48
		{ //49
			{
				Utf8:   "6.備註:",
				Big5:   []byte("6.\xb3\xc6\xb5\xf9:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //50
		{ //51
			{
				Utf8:   "1.侯 2.鄭 3.黃 4.盧 5.陳 6.柯",
				Big5:   []byte("1.\xabJ 2.\xbeG 3.\xb6\xc0 4.\xbfc 5.\xb3\xaf 6.\xac_"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //52
		{ //53
			{
				Utf8:   "草包滾了 柯果然墊底",
				Big5:   []byte("\xaf\xf3\xa5]\xbau\xa4F \xac_\xaaG\xb5M\xb9\xd4\xa9\xb3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{}, //54
		{}, //55
		{}, //56
		{ //51
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	testFirstComments11 = []*schema.Comment{
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg8UTpsAA:dNsw21NtPwVNZNl_fVeZew"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("s555666"),
			CreateTime: types.NanoTS(1261395960000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "https://i.imgur.com/CrYi0Ns.jpg",
						Big5:   []byte("https://i.imgur.com/CrYi0Ns.jpg                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "dNsw21NtPwVNZNl_fVeZew",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg8UT48kA:8xhYGp0JR7bWPuXgeieXJw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("s555666"),
			CreateTime: types.NanoTS(1261395960001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "柯糞崩潰啦",
						Big5:   []byte("\xac_\xc1T\xb1Y\xbc\xec\xb0\xd5                                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "8xhYGp0JR7bWPuXgeieXJw",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg8UUINIA:4ZAxuEaFj97_uIwLWUwQXw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("MeowDeLay"),
			CreateTime: types.NanoTS(1261395960002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "上任三個月就這麼強，真猛",
						Big5:   []byte("\xa4W\xa5\xf4\xa4T\xad\xd3\xa4\xeb\xb4N\xb3o\xbb\xf2\xb1j\xa1A\xafu\xb2r                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "4ZAxuEaFj97_uIwLWUwQXw",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg8UUXdsA:2SbMNPMpel6LHu5z9loYEg"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("milk250"),
			CreateTime: types.NanoTS(1261395960003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "看吧  台南人喜歡大張支票",
						Big5:   []byte("\xac\xdd\xa7a  \xa5x\xabn\xa4H\xb3\xdf\xc5w\xa4j\xb1i\xa4\xe4\xb2\xbc                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "2SbMNPMpel6LHu5z9loYEg",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg8UUmuQA:mLDBwhh8d0GLH7fluO-Vvg"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("JC910"),
			CreateTime: types.NanoTS(1261395960004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "爛名調，民進黨應該是第一名才對",
						Big5:   []byte("\xc4\xea\xa6W\xbd\xd5\xa1A\xa5\xc1\xb6i\xc4\xd2\xc0\xb3\xb8\xd3\xacO\xb2\xc4\xa4@\xa6W\xa4~\xb9\xef                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "mLDBwhh8d0GLH7fluO-Vvg",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z0xCAA:OauX-VCfv9NlKVTznlBVpw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("greensaru"),
			CreateTime: types.NanoTS(1261396020000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "還好啦！又不是第一次墊底了，習慣了啦！",
						Big5:   []byte("\xc1\xd9\xa6n\xb0\xd5\xa1I\xa4S\xa4\xa3\xacO\xb2\xc4\xa4@\xa6\xb8\xb9\xd4\xa9\xb3\xa4F\xa1A\xb2\xdf\xbaD\xa4F\xb0\xd5\xa1I               "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "OauX-VCfv9NlKVTznlBVpw",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z1ASkA:TKZBkWPYxqj27AnYM0LebA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("s359999"),
			CreateTime: types.NanoTS(1261396020001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "慶記市民就愛黑色道路這味嗎？",
						Big5:   []byte("\xbcy\xb0O\xa5\xab\xa5\xc1\xb4N\xb7R\xb6\xc2\xa6\xe2\xb9D\xb8\xf4\xb3o\xa8\xfd\xb6\xdc\xa1H                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "TKZBkWPYxqj27AnYM0LebA",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z1PjIA:kyVnoF9yRMlQP0cD0fif3w"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("yuchihsu"),
			CreateTime: types.NanoTS(1261396020002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "垃圾柯：台北人不懂感恩",
						Big5:   []byte("\xa9U\xa7\xa3\xac_\xa1G\xa5x\xa5_\xa4H\xa4\xa3\xc0\xb4\xb7P\xae\xa6                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "kyVnoF9yRMlQP0cD0fif3w",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z1ezsA:Q2BPLds3FiPA0H3MXrC8OQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Qoo20811"),
			CreateTime: types.NanoTS(1261396020003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "黃這樣還有第三 笑死",
						Big5:   []byte("\xb6\xc0\xb3o\xbc\xcb\xc1\xd9\xa6\xb3\xb2\xc4\xa4T \xaf\xba\xa6\xba                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "Q2BPLds3FiPA0H3MXrC8OQ",
		},
		{
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z1uEQA:N-j9rNFkWUj4nuYGq1ixjw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("linceass"),
			CreateTime: types.NanoTS(1261396020004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "藍綠賤畜一貫作風 把非藍綠往死裡打",
						Big5:   []byte("\xc2\xc5\xba\xf1\xbd\xe2\xafb\xa4@\xb3e\xa7@\xad\xb7 \xa7\xe2\xabD\xc2\xc5\xba\xf1\xa9\xb9\xa6\xba\xb8\xcc\xa5\xb4                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "N-j9rNFkWUj4nuYGq1ixjw",
		},
	}

	testTheRestComments11 = []*schema.Comment{
		{ //0
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z19U0A:7MK-wwy_g72bIylNku9cew"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("s359999"),
			CreateTime: types.NanoTS(1261396020005000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "https://i.imgur.com/LUaQ9yP.jpg",
						Big5:   []byte("https://i.imgur.com/LUaQ9yP.jpg                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "7MK-wwy_g72bIylNku9cew",
		},
		{ //1
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z2MlYA:0rTBC1_9pgXCCWFH_I6b2A"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("akway"),
			CreateTime: types.NanoTS(1261396020006000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "35%比當年韓草包還低了吧",
						Big5:   []byte("35%\xa4\xf1\xb7\xed\xa6~\xc1\xfa\xaf\xf3\xa5]\xc1\xd9\xa7C\xa4F\xa7a                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "0rTBC1_9pgXCCWFH_I6b2A",
		},
		{ //2
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z2b18A:kn6-2_1PZsvngRHQu38y8Q"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("b777787"),
			CreateTime: types.NanoTS(1261396020007000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "4%舔支仔：不可能",
						Big5:   []byte("4%\xbbQ\xa4\xe4\xa5J\xa1G\xa4\xa3\xa5i\xaf\xe0                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "kn6-2_1PZsvngRHQu38y8Q",
		},
		{ //3
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z2rGgA:LWCD5fMJhSo3iwACW8FGfQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("pieceiori"),
			CreateTime: types.NanoTS(1261396020008000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "假民調! 我看高雄板天天都歌舞昇平 幸福美滿",
						Big5:   []byte("\xb0\xb2\xa5\xc1\xbd\xd5! \xa7\xda\xac\xdd\xb0\xaa\xb6\xaf\xaaO\xa4\xd1\xa4\xd1\xb3\xa3\xbaq\xbbR\xaa@\xa5\xad \xa9\xaf\xba\xd6\xac\xfc\xba\xa1            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "LWCD5fMJhSo3iwACW8FGfQ",
		},
		{ //4
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z26XEA:X4Pa-FYKG05udfJ6SnQLow"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("proprome"),
			CreateTime: types.NanoTS(1261396020009000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "哈 笑死 連TVBS都墊底",
						Big5:   []byte("\xab\xa2 \xaf\xba\xa6\xba \xb3sTVBS\xb3\xa3\xb9\xd4\xa9\xb3                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "X4Pa-FYKG05udfJ6SnQLow",
		},
		{ //5
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFg_z3JnoA:ic7tXU1fc7SI_B8Fo4pK6g"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("fuhaho"),
			CreateTime: types.NanoTS(1261396020010000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "柯粉趕快怒洗一波政績取暖",
						Big5:   []byte("\xac_\xaf\xbb\xbb\xb0\xa7\xd6\xab\xe3\xac~\xa4@\xaai\xacF\xc1Z\xa8\xfa\xb7x                                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "ic7tXU1fc7SI_B8Fo4pK6g",
		},
		{ //6
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTV4YAA:NhPch5oUn9yd7aQAMgWWxQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("yulis"),
			CreateTime: types.NanoTS(1261396080000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "唬爛吧 誰有被民調問過的",
						Big5:   []byte("\xb0\xe4\xc4\xea\xa7a \xbd\xd6\xa6\xb3\xb3Q\xa5\xc1\xbd\xd5\xb0\xdd\xb9L\xaa\xba                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "NhPch5oUn9yd7aQAMgWWxQ",
		},
		{ //7
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTWHokA:FjJ28CxZc_-Q97pBvLvHYg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("mrlinwng"),
			CreateTime: types.NanoTS(1261396080001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "@.@",
						Big5:   []byte("@.@                                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "FjJ28CxZc_-Q97pBvLvHYg",
		},
		{ //8
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTWW5IA:YIIviyBSu4GFkqUTlcvnTQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("yoshiringo"),
			CreateTime: types.NanoTS(1261396080002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "\\我邁威武/\\我邁威武/\\我邁威武/\\我邁威武/\\我邁威武/",
						Big5:   []byte("\\\xa7\xda\xc1\xda\xab\xc2\xaaZ/\\\xa7\xda\xc1\xda\xab\xc2\xaaZ/\\\xa7\xda\xc1\xda\xab\xc2\xaaZ/\\\xa7\xda\xc1\xda\xab\xc2\xaaZ/\\\xa7\xda\xc1\xda\xab\xc2\xaaZ/  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "YIIviyBSu4GFkqUTlcvnTQ",
		},
		{ //9
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTWmJsA:oRh-k0pDZqxJs6PScht86A"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("Aidrux"),
			CreateTime: types.NanoTS(1261396080003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "白蠊垃圾崩潰",
						Big5:   []byte("\xa5\xd5\xf2\xe7\xa9U\xa7\xa3\xb1Y\xbc\xec                                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "oRh-k0pDZqxJs6PScht86A",
		},
		{ //10
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTW1aQA:mf6Iet5iaIlS1cg8-Fs-Sg"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("StarLeauge"),
			CreateTime: types.NanoTS(1261396080004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "陳其邁丟人",
						Big5:   []byte("\xb3\xaf\xa8\xe4\xc1\xda\xa5\xe1\xa4H                                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "mf6Iet5iaIlS1cg8-Fs-Sg",
		},
		{ //11
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTXEq0A:KITMtX7T0ZqdaCZt51AHLw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("KEYSOLIDER"),
			CreateTime: types.NanoTS(1261396080005000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "連任的阿北竟然輸給韓粉最討厭 的人尤其最近有美豬效",
						Big5:   []byte("\xb3s\xa5\xf4\xaa\xba\xaa\xfc\xa5_\xb3\xba\xb5M\xbf\xe9\xb5\xb9\xc1\xfa\xaf\xbb\xb3\xcc\xb0Q\xb9\xbd \xaa\xba\xa4H\xa4\xd7\xa8\xe4\xb3\xcc\xaa\xf1\xa6\xb3\xac\xfc\xbd\xde\xae\xc4  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "KITMtX7T0ZqdaCZt51AHLw",
		},
		{ //12
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTXT7YA:hVx6dxtjwHjIukYpqyVXOg"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("pieceiori"),
			CreateTime: types.NanoTS(1261396080006000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "家人被酒駕 市長給大張支票 真好",
						Big5:   []byte("\xaea\xa4H\xb3Q\xb0s\xber \xa5\xab\xaa\xf8\xb5\xb9\xa4j\xb1i\xa4\xe4\xb2\xbc \xafu\xa6n                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "hVx6dxtjwHjIukYpqyVXOg",
		},
		{ //13
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTXjL8A:S5yrznntUVoPzdC71azDxA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("giggleboy"),
			CreateTime: types.NanoTS(1261396080007000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "笑了 爐渣市第三名",
						Big5:   []byte("\xaf\xba\xa4F \xc4l\xb4\xed\xa5\xab\xb2\xc4\xa4T\xa6W                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "S5yrznntUVoPzdC71azDxA",
		},
		{ //14
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTXycgA:RNOKtUq-Ww-9lW26TwIvFw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("KEYSOLIDER"),
			CreateTime: types.NanoTS(1261396080008000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "應",
						Big5:   []byte("\xc0\xb3                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "RNOKtUq-Ww-9lW26TwIvFw",
		},
		{ //15
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhDTYBtEA:5qbzNAAjjU1D3LhUNmMWaA"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("OmegaWind"),
			CreateTime: types.NanoTS(1261396080009000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "綠畜有陳機掰   還好意思笑柯",
						Big5:   []byte("\xba\xf1\xafb\xa6\xb3\xb3\xaf\xbe\xf7\xd9T   \xc1\xd9\xa6n\xb7N\xab\xe4\xaf\xba\xac_                         "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "5qbzNAAjjU1D3LhUNmMWaA",
		},
		{ //16
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy2_uAA:ymz0HRfGNJIX9cdasLCyEw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("gunng"),
			CreateTime: types.NanoTS(1261396140000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "笑死 八卦爐渣洗了上百篇 結果黃還上升XD",
						Big5:   []byte("\xaf\xba\xa6\xba \xa4K\xa8\xf6\xc4l\xb4\xed\xac~\xa4F\xa4W\xa6\xca\xbdg \xb5\xb2\xaaG\xb6\xc0\xc1\xd9\xa4W\xa4\xc9XD                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "ymz0HRfGNJIX9cdasLCyEw",
		},
		{ //17
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy3O-kA:qvaVu3REh0fJL7ZBH7m6PQ"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("gogobar"),
			CreateTime: types.NanoTS(1261396140001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "笑死",
						Big5:   []byte("\xaf\xba\xa6\xba                                                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "qvaVu3REh0fJL7ZBH7m6PQ",
		},
		{ //18
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy3O-kA:qvaVu3REh0fJL7ZBH7m6PQ:R"),
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("cheinshin"),
			CreateTime: types.NanoTS(1261396140001100000),
			RefIDs:     []types.CommentID{"EYFhGy3O-kA:qvaVu3REh0fJL7ZBH7m6PQ"},
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "看狗屁崩潰 的確蠻好笑的",
						Big5:   []byte("\xac\xdd\xaa\xaf\xa7\xbe\xb1Y\xbc\xec \xaa\xba\xbdT\xc6Z\xa6n\xaf\xba\xaa\xba"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			EditNanoTS: types.NanoTS(1608551574000000000),
			IP:         "49.216.65.39",
			Host:       "臺灣",
			MD5:        "a5pFPRdkfCEDY_1FOMBWVg",
		},
		{ //19 (20)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy3ePIA:AsusUdGgu_XVhE9jh5QMAg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("bluezero000"),
			CreateTime: types.NanoTS(1261396140002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "吃爐渣感謝市長",
						Big5:   []byte("\xa6Y\xc4l\xb4\xed\xb7P\xc1\xc2\xa5\xab\xaa\xf8                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "AsusUdGgu_XVhE9jh5QMAg",
		},
		{ //20 (21)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy3tfsA:xENP4FFXDXO6amoUcs038g"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("chenchuhao"),
			CreateTime: types.NanoTS(1261396140003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "柯民調贏才是奇蹟，藍綠洗他洗成這樣哪有可能高XD",
						Big5:   []byte("\xac_\xa5\xc1\xbd\xd5\xc4\xb9\xa4~\xacO\xa9_\xc2\xdd\xa1A\xc2\xc5\xba\xf1\xac~\xa5L\xac~\xa6\xa8\xb3o\xbc\xcb\xad\xfe\xa6\xb3\xa5i\xaf\xe0\xb0\xaaXD      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "xENP4FFXDXO6amoUcs038g",
		},
		{ //21 (22)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy38wQA:bd8rQDmuhL0uon7EKO37YA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("b777787"),
			CreateTime: types.NanoTS(1261396140004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "柯畜氣到發抖",
						Big5:   []byte("\xac_\xafb\xae\xf0\xa8\xec\xb5o\xa7\xdd                                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "bd8rQDmuhL0uon7EKO37YA",
		},
		{ //22 (23)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhGy4MA0A:fTnc-D1T2UogLI79V7-gVQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("qweertyui891"),
			CreateTime: types.NanoTS(1261396140005000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "證明反不反萊豬跟民調無關",
						Big5:   []byte("\xc3\xd2\xa9\xfa\xa4\xcf\xa4\xa3\xa4\xcf\xb5\xdc\xbd\xde\xb8\xf2\xa5\xc1\xbd\xd5\xb5L\xc3\xf6                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "fTnc-D1T2UogLI79V7-gVQ",
		},
		{ //23 (24)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhKSYHEAA:6o0zsr2Rmw_IyWb5gI4hzw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("indium111"),
			CreateTime: types.NanoTS(1261396200000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "TVBS偏藍，給綠營縣市長的數字本來就偏低了",
						Big5:   []byte("TVBS\xb0\xbe\xc2\xc5\xa1A\xb5\xb9\xba\xf1\xc0\xe7\xbf\xa4\xa5\xab\xaa\xf8\xaa\xba\xbc\xc6\xa6r\xa5\xbb\xa8\xd3\xb4N\xb0\xbe\xa7C\xa4F             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "6o0zsr2Rmw_IyWb5gI4hzw",
		},
		{ //24 (25)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhKSYWUkA:HfV9utwJvAm5UlC8hCMCAQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("chenchuhao"),
			CreateTime: types.NanoTS(1261396200001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不過西瓜哲這麼高真的蠻可笑的...",
						Big5:   []byte("\xa4\xa3\xb9L\xa6\xe8\xa5\xca\xad\xf5\xb3o\xbb\xf2\xb0\xaa\xafu\xaa\xba\xc6Z\xa5i\xaf\xba\xaa\xba...                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "HfV9utwJvAm5UlC8hCMCAQ",
		},
		{ //25 (26)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhKSYllIA:VoAV787V5UmqE3rWfjP6zQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("Aidrux"),
			CreateTime: types.NanoTS(1261396200002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "好了啦柯粉 去找小勝文取暖啦",
						Big5:   []byte("\xa6n\xa4F\xb0\xd5\xac_\xaf\xbb \xa5h\xa7\xe4\xa4p\xb3\xd3\xa4\xe5\xa8\xfa\xb7x\xb0\xd5                             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "VoAV787V5UmqE3rWfjP6zQ",
		},
		{ //26 (27)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhKSY01sA:BK_DZPCFozgILSlXLXydfw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("andy3580"),
			CreateTime: types.NanoTS(1261396200003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "現實就是4% 可憐",
						Big5:   []byte("\xb2{\xb9\xea\xb4N\xacO4% \xa5i\xbc\xa6                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "BK_DZPCFozgILSlXLXydfw",
		},
		{ //27 (28)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhKSZEGQA:ETDoEPI4tu7gydfdU0VXDw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("foolfighter"),
			CreateTime: types.NanoTS(1261396200004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "南部柯糞整天幻想阿北好棒XDDD",
						Big5:   []byte("\xabn\xb3\xa1\xac_\xc1T\xbe\xe3\xa4\xd1\xa4\xdb\xb7Q\xaa\xfc\xa5_\xa6n\xb4\xceXDDD                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "ETDoEPI4tu7gydfdU0VXDw",
		},
		{ //28 (29)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhKSZTW0A:q024J-pzNIzR1CJv2v614w"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("berkeley5566"),
			CreateTime: types.NanoTS(1261396200005000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "臺南人真好",
						Big5:   []byte("\xbbO\xabn\xa4H\xafu\xa6n                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "q024J-pzNIzR1CJv2v614w",
		},
		{ //29 (30)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx5OaAA:mwnGopEbEQjAB0QDa03HwA"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("indium111"),
			CreateTime: types.NanoTS(1261396260000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "考慮到機構效應，鄭黃陳三人的民調數字應該會更高",
						Big5:   []byte("\xa6\xd2\xbc{\xa8\xec\xbe\xf7\xbac\xae\xc4\xc0\xb3\xa1A\xbeG\xb6\xc0\xb3\xaf\xa4T\xa4H\xaa\xba\xa5\xc1\xbd\xd5\xbc\xc6\xa6r\xc0\xb3\xb8\xd3\xb7|\xa7\xf3\xb0\xaa       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "mwnGopEbEQjAB0QDa03HwA",
		},
		{ //30 (31)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx5dqkA:8S9fqyPjGvWH-uASU8-dgw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("johnwu"),
			CreateTime: types.NanoTS(1261396260001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "台北市長不是黃33嗎? 柯屁誰啊?",
						Big5:   []byte("\xa5x\xa5_\xa5\xab\xaa\xf8\xa4\xa3\xacO\xb6\xc033\xb6\xdc? \xac_\xa7\xbe\xbd\xd6\xb0\xda?                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "8S9fqyPjGvWH-uASU8-dgw",
		},
		{ //31 (32)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx5s7IA:cyCktxdV8xZ76mWofGMjaQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("KEYSOLIDER"),
			CreateTime: types.NanoTS(1261396260002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "應該是邁失言次數不算多 不然罷免韓總成功才過幾個月",
						Big5:   []byte("\xc0\xb3\xb8\xd3\xacO\xc1\xda\xa5\xa2\xa8\xa5\xa6\xb8\xbc\xc6\xa4\xa3\xba\xe2\xa6h \xa4\xa3\xb5M\xbd}\xa7K\xc1\xfa\xc1`\xa6\xa8\xa5\\\xa4~\xb9L\xb4X\xad\xd3\xa4\xeb   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "cyCktxdV8xZ76mWofGMjaQ",
		},
		{ //32 (33)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx58LsA:FpLsrv8smx-gy5lkG4S7fg"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("KEYSOLIDER"),
			CreateTime: types.NanoTS(1261396260003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "韓粉一定超恨",
						Big5:   []byte("\xc1\xfa\xaf\xbb\xa4@\xa9w\xb6W\xab\xeb                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "FpLsrv8smx-gy5lkG4S7fg",
		},
		{ //33 (34)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx6LcQA:q6fRVVC5SPnnsEHByC7p4w"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("azeroth"),
			CreateTime: types.NanoTS(1261396260004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "爐渣市民真的很光榮XDDDDD",
						Big5:   []byte("\xc4l\xb4\xed\xa5\xab\xa5\xc1\xafu\xaa\xba\xab\xdc\xa5\xfa\xbaaXDDDDD                               "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "q6fRVVC5SPnnsEHByC7p4w",
		},
		{ //34 (35)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx6as0A:11qiZDDyJEkFN_lFleo5Dw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("newstarisme"),
			CreateTime: types.NanoTS(1261396260005000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這我信",
						Big5:   []byte("\xb3o\xa7\xda\xabH                                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "11qiZDDyJEkFN_lFleo5Dw",
		},
		{ //35 (36)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhNx6p9YA:FzLbMSDglY8_K-jw2HeUOg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("whitecow"),
			CreateTime: types.NanoTS(1261396260006000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "盧也能第四..爛透",
						Big5:   []byte("\xbfc\xa4]\xaf\xe0\xb2\xc4\xa5|..\xc4\xea\xb3z                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "FzLbMSDglY8_K-jw2HeUOg",
		},
		{ //36 (37)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhRRaVwAA:JqSIETfjIWvw1-sQoJGJVg"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("indium111"),
			CreateTime: types.NanoTS(1261396320000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "三立做民調如果盧只有40%，我相信藍粉也會喊機構效應的",
						Big5:   []byte("\xa4T\xa5\xdf\xb0\xb5\xa5\xc1\xbd\xd5\xa6p\xaaG\xbfc\xa5u\xa6\xb340%\xa1A\xa7\xda\xac\xdb\xabH\xc2\xc5\xaf\xbb\xa4]\xb7|\xb3\xdb\xbe\xf7\xbac\xae\xc4\xc0\xb3\xaa\xba  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "JqSIETfjIWvw1-sQoJGJVg",
		},
		{ //37 (39)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhRRalAkA:VotHH7dDQucqxEj2FOC0ig"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("d86506"),
			CreateTime: types.NanoTS(1261396320001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "陳水扁是民調第一但連任選輸才有選總統的資格，不知道柯連",
						Big5:   []byte("\xb3\xaf\xa4\xf4\xab\xf3\xacO\xa5\xc1\xbd\xd5\xb2\xc4\xa4@\xa6\xfd\xb3s\xa5\xf4\xbf\xef\xbf\xe9\xa4~\xa6\xb3\xbf\xef\xc1`\xb2\xce\xaa\xba\xb8\xea\xae\xe6\xa1A\xa4\xa3\xaa\xbe\xb9D\xac_\xb3s "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "VotHH7dDQucqxEj2FOC0ig",
		},
		{ //38 (40)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhRRa0RIA:9hKRfrW5maV-p0kyiZaoQg"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("d86506"),
			CreateTime: types.NanoTS(1261396320002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "任選贏，但民調吊車尾，有沒有資格選總統？",
						Big5:   []byte("\xa5\xf4\xbf\xef\xc4\xb9\xa1A\xa6\xfd\xa5\xc1\xbd\xd5\xa6Q\xa8\xae\xa7\xc0\xa1A\xa6\xb3\xa8S\xa6\xb3\xb8\xea\xae\xe6\xbf\xef\xc1`\xb2\xce\xa1H                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "9hKRfrW5maV-p0kyiZaoQg",
		},
		{ //39 (41)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhRRbDhsA:48BL2Dd9HWBE0KFyTkx3iQ"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("BlackBass"),
			CreateTime: types.NanoTS(1261396320003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "TVB狗屎",
						Big5:   []byte("TVB\xaa\xaf\xab\xcb                                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "48BL2Dd9HWBE0KFyTkx3iQ",
		},
		{ //40 (42)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("EraKing"),
			CreateTime: types.NanoTS(1261396380000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "台南高雄兩個這樣的滿意度根本超爛",
						Big5:   []byte("\xa5x\xabn\xb0\xaa\xb6\xaf\xa8\xe2\xad\xd3\xb3o\xbc\xcb\xaa\xba\xba\xa1\xb7N\xab\xd7\xae\xda\xa5\xbb\xb6W\xc4\xea                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "yBPgyUXJMLN6p6EYeODktQ",
		},
		{ //41 (43)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ:R"),
			RefIDs:     []types.CommentID{"EYFhUw7dGAA:yBPgyUXJMLN6p6EYeODktQ"},
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("cheinshin"),
			CreateTime: types.NanoTS(1261396380000100000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "黃在遠見有78 平均一下 大約68喔",
						Big5:   []byte("\xb6\xc0\xa6b\xbb\xb7\xa8\xa3\xa6\xb378 \xa5\xad\xa7\xa1\xa4@\xa4U \xa4j\xac\xf968\xb3\xe1"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			IP:         "49.216.65.39",
			Host:       "臺灣",
			MD5:        "m4zgWiP6HQDKfLxk-6bAhw",
			EditNanoTS: types.NanoTS(1608551574000000000),
		},
		{ //42 (45)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhUw7sWkA:3z12a9ih3kgrtdLxUACH1Q"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("lakeisland"),
			CreateTime: types.NanoTS(1261396380001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "垃圾綠媒假民調",
						Big5:   []byte("\xa9U\xa7\xa3\xba\xf1\xb4C\xb0\xb2\xa5\xc1\xbd\xd5                                     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "3z12a9ih3kgrtdLxUACH1Q",
		},
		{ //43 (46)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhUw77nIA:smfcjOjPJVDSyoTWjyENnw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("KEYSOLIDER"),
			CreateTime: types.NanoTS(1261396380002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "八卦熱呼呼 民調冷冰冰",
						Big5:   []byte("\xa4K\xa8\xf6\xbc\xf6\xa9I\xa9I \xa5\xc1\xbd\xd5\xa7N\xa6B\xa6B                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "smfcjOjPJVDSyoTWjyENnw",
		},
		{ //44 (48)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhYQckcAA:qjTVjQQ7Vy7K3EJTNXro_w"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("indium111"),
			CreateTime: types.NanoTS(1261396440000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "TVBS什麼時候變成綠媒了XD",
						Big5:   []byte("TVBS\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xc5\xdc\xa6\xa8\xba\xf1\xb4C\xa4FXD                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "qjTVjQQ7Vy7K3EJTNXro_w",
		},
		{ //45 (49)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhYQczskA:r7TpBHcfWBjUjeMypcTUIA"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("bibiwei"),
			CreateTime: types.NanoTS(1261396440001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "今年才補選上的啦,什麼時候去年補選上了?",
						Big5:   []byte("\xa4\xb5\xa6~\xa4~\xb8\xc9\xbf\xef\xa4W\xaa\xba\xb0\xd5,\xa4\xb0\xbb\xf2\xae\xc9\xad\xd4\xa5h\xa6~\xb8\xc9\xbf\xef\xa4W\xa4F?                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "r7TpBHcfWBjUjeMypcTUIA",
		},
		{ //46 (50)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhYQdC9IA:XwyfHMVLp1uw_UgD5HX_yg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("AndyZer"),
			CreateTime: types.NanoTS(1261396440002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "高雄人覺得又選錯人了 (╯°□°）╯︵ ┴━┴",
						Big5:   []byte("\xb0\xaa\xb6\xaf\xa4H\xc4\xb1\xb1o\xa4S\xbf\xef\xbf\xf9\xa4H\xa4F (\xa2\xa3\xa2X\xa1\xbc\xa2X\xa1^\xa2\xa3\xa1_ \xa2r\x9d}\xa2r           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "XwyfHMVLp1uw_UgD5HX_yg",
		},
		{ //47 (51)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhYQdSNsA:_6wV6MlZUZWb0qIqx5a7tA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("billybbb"),
			CreateTime: types.NanoTS(1261396440003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "綠智障們不是罵盧破壞臺美關係嗎？可憐喔被民調洗臉",
						Big5:   []byte("\xba\xf1\xb4\xbc\xbb\xd9\xad\xcc\xa4\xa3\xacO\xbd|\xbfc\xaf}\xc3a\xbbO\xac\xfc\xc3\xf6\xabY\xb6\xdc\xa1H\xa5i\xbc\xa6\xb3\xe1\xb3Q\xa5\xc1\xbd\xd5\xac~\xc1y     "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "_6wV6MlZUZWb0qIqx5a7tA",
		},
		{ //48 (52)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhYQdheQA:a1nbh9m8KYnC0QQevDRyqA"),
			TheType:    types.COMMENT_TYPE_BOO,
			Owner:      bbs.UUserID("carrey8"),
			CreateTime: types.NanoTS(1261396440004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "搞個錘子 連其邁都贏不了?",
						Big5:   []byte("\xb7d\xad\xd3\xc1\xe8\xa4l \xb3s\xa8\xe4\xc1\xda\xb3\xa3\xc4\xb9\xa4\xa3\xa4F?                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "a1nbh9m8KYnC0QQevDRyqA",
		},
		{ //49 (53)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv9ryAA:u0m4ezxFyxG8CX56gLwr4Q"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("hyde711034"),
			CreateTime: types.NanoTS(1261396500000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "台南人可憐還開心指數最高",
						Big5:   []byte("\xa5x\xabn\xa4H\xa5i\xbc\xa6\xc1\xd9\xb6}\xa4\xdf\xab\xfc\xbc\xc6\xb3\xcc\xb0\xaa                           "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "u0m4ezxFyxG8CX56gLwr4Q",
		},
		{ //50 (54)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv97CkA:P0aJ_OzsnWCHCW4E1dszLw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("icelaw"),
			CreateTime: types.NanoTS(1261396500001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "柯在綠媒偏低 正常",
						Big5:   []byte("\xac_\xa6b\xba\xf1\xb4C\xb0\xbe\xa7C \xa5\xbf\xb1`                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "P0aJ_OzsnWCHCW4E1dszLw",
		},
		{ //51 (55)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv-KTIA:lXzclvbNv1xXqIYNFr_6fg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("touchbird"),
			CreateTime: types.NanoTS(1261396500002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "阿北沒有輸!!!",
						Big5:   []byte("\xaa\xfc\xa5_\xa8S\xa6\xb3\xbf\xe9!!!                                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "lXzclvbNv1xXqIYNFr_6fg",
		},
		{ //52 (56)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv-ZjsA:qNlvQ8zw4_MVcjzrLYZ9gw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("bradpete"),
			CreateTime: types.NanoTS(1261396500003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "笑死 哈哈哈 身在福中不知福.jpg",
						Big5:   []byte("\xaf\xba\xa6\xba \xab\xa2\xab\xa2\xab\xa2 \xa8\xad\xa6b\xba\xd6\xa4\xa4\xa4\xa3\xaa\xbe\xba\xd6.jpg                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "qNlvQ8zw4_MVcjzrLYZ9gw",
		},
		{ //53 (57)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv-o0QA:yFk_pNMxLmG5oVw-7Eyl5w"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("bradpete"),
			CreateTime: types.NanoTS(1261396500004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "不來八卦版做一下民調嗎 柯肯定第一名",
						Big5:   []byte("\xa4\xa3\xa8\xd3\xa4K\xa8\xf6\xaa\xa9\xb0\xb5\xa4@\xa4U\xa5\xc1\xbd\xd5\xb6\xdc \xac_\xaa\xd6\xa9w\xb2\xc4\xa4@\xa6W                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "yFk_pNMxLmG5oVw-7Eyl5w",
		},
		{ //54 (58)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhbv-4E0A:6_xY0__4tPC8xB1k7Wt5NQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("amos30627"),
			CreateTime: types.NanoTS(1261396500005000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "阿北又被綠媒抹黑",
						Big5:   []byte("\xaa\xfc\xa5_\xa4S\xb3Q\xba\xf1\xb4C\xa9\xd9\xb6\xc2                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "6_xY0__4tPC8xB1k7Wt5NQ",
		},
		{ //55 (59)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhfPezIAA:HKKMDWLDUaUAwCydMIXChQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Niro"),
			CreateTime: types.NanoTS(1261396560000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "一般通常會有蜜月期  想不到侯黃還那麼高",
						Big5:   []byte("\xa4@\xaf\xeb\xb3q\xb1`\xb7|\xa6\xb3\xbbe\xa4\xeb\xb4\xc1  \xb7Q\xa4\xa3\xa8\xec\xabJ\xb6\xc0\xc1\xd9\xa8\xba\xbb\xf2\xb0\xaa                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "HKKMDWLDUaUAwCydMIXChQ",
		},
		{ //56 (60)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhiu_6eAA:1TuCV38ll-rSzuvRJ97moQ"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("VVizZ"),
			CreateTime: types.NanoTS(1261396620000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "4%笑死",
						Big5:   []byte("4%\xaf\xba\xa6\xba                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "1TuCV38ll-rSzuvRJ97moQ",
		},
		{ //57 (61)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhivAJukA:qyNopZVr9jmK-p7AOqXPpA"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("indium111"),
			CreateTime: types.NanoTS(1261396620001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這民調再度顯示現在的八卦版根本社會反指標了",
						Big5:   []byte("\xb3o\xa5\xc1\xbd\xd5\xa6A\xab\xd7\xc5\xe3\xa5\xdc\xb2{\xa6b\xaa\xba\xa4K\xa8\xf6\xaa\xa9\xae\xda\xa5\xbb\xaa\xc0\xb7|\xa4\xcf\xab\xfc\xbc\xd0\xa4F          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "qyNopZVr9jmK-p7AOqXPpA",
		},
		{ //58 (62)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhivAY_IA:QfUeb-goUPOrC5n7PSCKBw"),
			TheType:    types.COMMENT_TYPE_COMMENT,
			Owner:      bbs.UUserID("takeda3234"),
			CreateTime: types.NanoTS(1261396620002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "這樣藍營2022很穩阿...台北拿回機率大",
						Big5:   []byte("\xb3o\xbc\xcb\xc2\xc5\xc0\xe72022\xab\xdc\xc3\xad\xaa\xfc...\xa5x\xa5_\xae\xb3\xa6^\xbe\xf7\xb2v\xa4j                "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "QfUeb-goUPOrC5n7PSCKBw",
		},
		{ //59 (63)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhivAoPsA:KLkw5UsBLsrn1DAUKxbe5g"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("Niubert"),
			CreateTime: types.NanoTS(1261396620003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "爐渣市長怎麼可能上升，連我深綠台南同學都發文抱怨了",
						Big5:   []byte("\xc4l\xb4\xed\xa5\xab\xaa\xf8\xab\xe7\xbb\xf2\xa5i\xaf\xe0\xa4W\xa4\xc9\xa1A\xb3s\xa7\xda\xb2`\xba\xf1\xa5x\xabn\xa6P\xbe\xc7\xb3\xa3\xb5o\xa4\xe5\xa9\xea\xab\xe8\xa4F    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "KLkw5UsBLsrn1DAUKxbe5g",
		},
		{ //60 (64)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhivA3gQA:bwLysalLRJbnWlE1MH9xXw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("qweertyui891"),
			CreateTime: types.NanoTS(1261396620004000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "民調第一有總統資格，看來侯友宜穩了",
						Big5:   []byte("\xa5\xc1\xbd\xd5\xb2\xc4\xa4@\xa6\xb3\xc1`\xb2\xce\xb8\xea\xae\xe6\xa1A\xac\xdd\xa8\xd3\xabJ\xa4\xcd\xa9y\xc3\xad\xa4F               "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "bwLysalLRJbnWlE1MH9xXw",
		},
		{ //61 (65)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhB0AA:klMsUdLuqHMayZ9fLs7RSg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("irosehead"),
			CreateTime: types.NanoTS(1261396680000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "柯屁連吸盡韓糞仇恨的陳其邁都打不贏 真廢",
						Big5:   []byte("\xac_\xa7\xbe\xb3s\xa7l\xba\xc9\xc1\xfa\xc1T\xa4\xb3\xab\xeb\xaa\xba\xb3\xaf\xa8\xe4\xc1\xda\xb3\xa3\xa5\xb4\xa4\xa3\xc4\xb9 \xafu\xbco             "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "klMsUdLuqHMayZ9fLs7RSg",
		},
		{ //62 (66)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("foolfighter"),
			CreateTime: types.NanoTS(1261396680001000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "土柯糞幾成有台北市戶籍啊？",
						Big5:   []byte("\xa4g\xac_\xc1T\xb4X\xa6\xa8\xa6\xb3\xa5x\xa5_\xa5\xab\xa4\xe1\xc4y\xb0\xda\xa1H                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "Es26f7U0EXdr7Gp4a9N8pQ",
		},
		{ //63 (67)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("yehpi"),
			CreateTime: types.NanoTS(1261396680002000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "黃有超大支票（物理）",
						Big5:   []byte("\xb6\xc0\xa6\xb3\xb6W\xa4j\xa4\xe4\xb2\xbc\xa1]\xaa\xab\xb2z\xa1^                                    "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "gmrKWXE7BjV-1U89GcPqHg",
		},
		{ //64 (68)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("lockeyman"),
			CreateTime: types.NanoTS(1261396680003000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噗",
						Big5:   []byte("\xbcP                                                  "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "cpqbGyLoF_jIyITF4bv-rQ",
		},
		{ //65 (69)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R"),
			RefIDs:     []types.CommentID{"EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ"},
			TheType:    types.COMMENT_TYPE_REPLY,
			Owner:      bbs.UUserID("cheinshin"),
			CreateTime: types.NanoTS(1261396680003100000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "test123123",
						Big5:   []byte("test123123"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
				{},
				{
					{
						Utf8:   "test124124",
						Big5:   []byte("test124124"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
				{},
				{
					{
						Utf8:   "test125125",
						Big5:   []byte("test125125"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5:        "YdlPEuzP2CXzn8nA-n92Ow",
			IP:         "49.216.65.39",
			Host:       "臺灣",
			EditNanoTS: types.NanoTS(1608551574000000000),
		},
		{ //66 (76)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFuT-Ew6AA:ALE6XIa5ARhXunryJTB3xg"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("deathdancer"),
			CreateTime: types.NanoTS(1261410660000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "阿北才沒有輸",
						Big5:   []byte("\xaa\xfc\xa5_\xa4~\xa8S\xa6\xb3\xbf\xe9                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "ALE6XIa5ARhXunryJTB3xg",
		},
		{ //67 (77)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFu6YxBsAA:xKzE3ZYx3C20_i3o3OMHbQ"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("joshualiu"),
			CreateTime: types.NanoTS(1261411320000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "黃偉哲是怎麼",
						Big5:   []byte("\xb6\xc0\xb0\xb6\xad\xf5\xacO\xab\xe7\xbb\xf2                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "xKzE3ZYx3C20_i3o3OMHbQ",
		},
		{ //68 (78)
			BBoardID:   bbs.BBoardID("test"),
			ArticleID:  bbs.ArticleID("test"),
			CommentID:  types.CommentID("EYFvZ0bDyAA:E_omXaZAX9dffXNX2NAZYw"),
			TheType:    types.COMMENT_TYPE_RECOMMEND,
			Owner:      bbs.UUserID("ken0062"),
			CreateTime: types.NanoTS(1261411860000000000),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "做越多越被嫌ww",
						Big5:   []byte("\xb0\xb5\xb6V\xa6h\xb6V\xb3Q\xb6\xfbww                                        "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
			MD5: "E_omXaZAX9dffXNX2NAZYw",
		},
	}
}
