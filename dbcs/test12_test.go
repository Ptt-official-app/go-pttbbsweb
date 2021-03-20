package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

var (
	testFilename12            = "temp7"
	testContentAll12          []byte
	testContent12             []byte
	testSignature12           []byte
	testComment12             []byte
	testFirstCommentsDBCS12   []byte
	testTheRestCommentsDBCS12 []byte
	testContent12Big5         [][]*types.Rune
	testContent12Utf8         [][]*types.Rune

	testFirstComments12 []*schema.Comment
)

func initTest12() {
	testContentAll12, testContent12, testSignature12, testComment12, testFirstCommentsDBCS12, testTheRestCommentsDBCS12 = loadTest(testFilename12)

	testContent12Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: thouloveme (\xbb\xae\xbb\xae) \xac\xdd\xaaO: Gossiping"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: thouloveme (\xbb\xae\xbb\xae) \xac\xdd\xaaO: Gossiping\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb7s\xbbD] \xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb7s\xbbD] \xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Dec 21 18:54:24 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Mon Dec 21 18:54:24 2020\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //5
			{

				Big5:   []byte("1.\xb4C\xc5\xe9\xa8\xd3\xb7\xbd:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("1.\xb4C\xc5\xe9\xa8\xd3\xb7\xbd:\r"),
			},
		},
		{ //6
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //7
			{

				Big5:   []byte("\xa4T\xa5\xdf\xb7s\xbbD"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4T\xa5\xdf\xb7s\xbbD\r"),
			},
		},
		{ //8
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //9
			{

				Big5:   []byte("2.\xb0O\xaa\xcc\xb8p\xa6W:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("2.\xb0O\xaa\xcc\xb8p\xa6W:\r"),
			},
		},
		{ //10
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //11
			{

				Big5:   []byte("\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9\r"),
			},
		},
		{ //12
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //13
			{

				Big5:   []byte("2020/12/21 18:45:00"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("2020/12/21 18:45:00\r"),
			},
		},
		{ //14
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //15
			{

				Big5:   []byte("3.\xa7\xb9\xbe\xe3\xb7s\xbbD\xbc\xd0\xc3D:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("3.\xa7\xb9\xbe\xe3\xb7s\xbbD\xbc\xd0\xc3D:\r"),
			},
		},
		{ //16
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //17
			{

				Big5:   []byte("\xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa\r"),
			},
		},
		{ //18
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //19
			{

				Big5:   []byte("4.\xa7\xb9\xbe\xe3\xb7s\xbbD\xa4\xba\xa4\xe5:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("4.\xa7\xb9\xbe\xe3\xb7s\xbbD\xa4\xba\xa4\xe5:\r"),
			},
		},
		{ //20
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //21
			{

				Big5:   []byte("\xa5\xcd\xac\xa1\xa4\xa4\xa4\xdf\xa1\xfe\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xcd\xac\xa1\xa4\xa4\xa4\xdf\xa1\xfe\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9\r"),
			},
		},
		{ //22
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //23
			{

				Big5:   []byte("\xaa\xf1\xa4\xe9\xc0W\xb6\xc7\xb6\xa7\xa9\xfa\xa4s\xb0\xea\xaea\xa4\xbd\xb6\xe9\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xbc\xc9\xc0\xc5\xa8\xc6\xa5\xf3\xa1A\xa6\xdb\xa4\xb5\xa6~\xa5H\xa8\xd3\xa4w\xa6\xb324\xb0\xa6\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xb3\xb0\xc4\xf2\xa6\xba\xa4`\xa1A\xb9\xef"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaa\xf1\xa4\xe9\xc0W\xb6\xc7\xb6\xa7\xa9\xfa\xa4s\xb0\xea\xaea\xa4\xbd\xb6\xe9\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xbc\xc9\xc0\xc5\xa8\xc6\xa5\xf3\xa1A\xa6\xdb\xa4\xb5\xa6~\xa5H\xa8\xd3\xa4w\xa6\xb324\xb0\xa6\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xb3\xb0\xc4\xf2\xa6\xba\xa4`\xa1A\xb9\xef\r"),
			},
		},
		{ //24
			{

				Big5:   []byte("\xa6\xb9\xa5_\xa5\xab\xb0\xca\xabO\xb3B\xa4]\xbd\xd0\xa8\xd3\xb1M\xb7~\xc3~\xc2\xe5\xaev\xad\xe5\xc0\xcb\xb1\xc4\xbc\xcb\xa1A\xbe\xda\xb1\xc4\xc0\xcb\xb5\xb2\xaaG\xc5\xe3\xa5\xdc\xa1A\xa6\xba\xa4`\xad\xec\xa6]\xa5i\xaf\xe0\xa9M\xad\xb9\xaa\xab\xab~\xbd\xe8\xa4\xa3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6\xb9\xa5_\xa5\xab\xb0\xca\xabO\xb3B\xa4]\xbd\xd0\xa8\xd3\xb1M\xb7~\xc3~\xc2\xe5\xaev\xad\xe5\xc0\xcb\xb1\xc4\xbc\xcb\xa1A\xbe\xda\xb1\xc4\xc0\xcb\xb5\xb2\xaaG\xc5\xe3\xa5\xdc\xa1A\xa6\xba\xa4`\xad\xec\xa6]\xa5i\xaf\xe0\xa9M\xad\xb9\xaa\xab\xab~\xbd\xe8\xa4\xa3\r"),
			},
		},
		{ //25
			{

				Big5:   []byte("\xa8\xce\xa6\xb3\xc3\xf6\xa1A\xbe\xc9\xadP\xa4\xfb\xb0\xa6\xad\xcc\xaa\xf8\xb4\xc1\xc0\xe7\xbei\xc0\xf2\xa8\xfa\xa4\xa3\xa8\xac\xa1A\xa4]\xaa\xec\xa8B\xb1\xc6\xb0\xa3\xacO\xb0\xca\xaa\xab\xb6\xa1\xb6\xc7\xacV\xaff\xa9\xce\xacO\xa4H\xafb\xa6@\xb3q\xb6\xc7\xacV\xaff"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa8\xce\xa6\xb3\xc3\xf6\xa1A\xbe\xc9\xadP\xa4\xfb\xb0\xa6\xad\xcc\xaa\xf8\xb4\xc1\xc0\xe7\xbei\xc0\xf2\xa8\xfa\xa4\xa3\xa8\xac\xa1A\xa4]\xaa\xec\xa8B\xb1\xc6\xb0\xa3\xacO\xb0\xca\xaa\xab\xb6\xa1\xb6\xc7\xacV\xaff\xa9\xce\xacO\xa4H\xafb\xa6@\xb3q\xb6\xc7\xacV\xaff\r"),
			},
		},
		{ //26
			{

				Big5:   []byte("\xaa\xba\xa5i\xaf\xe0\xa9\xca"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaa\xba\xa5i\xaf\xe0\xa9\xca\r"),
			},
		},
		{ //27
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //28
			{

				Big5:   []byte("https://attach.setn.com/newsimages/2020/12/08/2924507-PH.jpg"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("https://attach.setn.com/newsimages/2020/12/08/2924507-PH.jpg\r"),
			},
		},
		{ //29
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //30
			{

				Big5:   []byte("5.\xa7\xb9\xbe\xe3\xb7s\xbbD\xb3s\xb5\xb2 (\xa9\xce\xb5u\xba\xf4\xa7}):"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("5.\xa7\xb9\xbe\xe3\xb7s\xbbD\xb3s\xb5\xb2 (\xa9\xce\xb5u\xba\xf4\xa7}):\r"),
			},
		},
		{ //31
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //32
			{

				Big5:   []byte("https://www.setn.com/News.aspx?NewsID=869094"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("https://www.setn.com/News.aspx?NewsID=869094\r"),
			},
		},
		{ //33
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //34
			{

				Big5:   []byte("6.\xb3\xc6\xb5\xf9:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("6.\xb3\xc6\xb5\xf9:\r"),
			},
		},
		{ //35
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //36
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //37
			{

				Big5:   []byte("\xc0\xe7\xbei\xa4\xa3\xa8}?????????????????????"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xc0\xe7\xbei\xa4\xa3\xa8}?????????????????????\r"),
			},
		},
		{ //38
			{

				Big5:   []byte("\xb3o\xa6\xb3\xc2I\xa7\xe8..."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3o\xa6\xb3\xc2I\xa7\xe8...\r"),
			},
		}}

	testContent12Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: thouloveme (赫赫) 看板: Gossiping",
				Big5:   []byte("\xa7@\xaa\xcc: thouloveme (\xbb\xae\xbb\xae) \xac\xdd\xaaO: Gossiping"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: thouloveme (\xbb\xae\xbb\xae) \xac\xdd\xaaO: Gossiping\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: [新聞] 陽明山擎天崗24頭牛暴斃　死因曝光",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb7s\xbbD] \xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb7s\xbbD] \xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Mon Dec 21 18:54:24 2020",
				Big5:   []byte("\xae\xc9\xb6\xa1: Mon Dec 21 18:54:24 2020"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Mon Dec 21 18:54:24 2020\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //5
			{
				Utf8:   "1.媒體來源:",
				Big5:   []byte("1.\xb4C\xc5\xe9\xa8\xd3\xb7\xbd:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("1.\xb4C\xc5\xe9\xa8\xd3\xb7\xbd:\r"),
			},
		},
		{ //6
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //7
			{
				Utf8:   "三立新聞",
				Big5:   []byte("\xa4T\xa5\xdf\xb7s\xbbD"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4T\xa5\xdf\xb7s\xbbD\r"),
			},
		},
		{ //8
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //9
			{
				Utf8:   "2.記者署名:",
				Big5:   []byte("2.\xb0O\xaa\xcc\xb8p\xa6W:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("2.\xb0O\xaa\xcc\xb8p\xa6W:\r"),
			},
		},
		{ //10
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //11
			{
				Utf8:   "林奇樺報導",
				Big5:   []byte("\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9\r"),
			},
		},
		{ //12
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //13
			{
				Utf8:   "2020/12/21 18:45:00",
				Big5:   []byte("2020/12/21 18:45:00"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("2020/12/21 18:45:00\r"),
			},
		},
		{ //14
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //15
			{
				Utf8:   "3.完整新聞標題:",
				Big5:   []byte("3.\xa7\xb9\xbe\xe3\xb7s\xbbD\xbc\xd0\xc3D:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("3.\xa7\xb9\xbe\xe3\xb7s\xbbD\xbc\xd0\xc3D:\r"),
			},
		},
		{ //16
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //17
			{
				Utf8:   "陽明山擎天崗24頭牛暴斃　死因曝光",
				Big5:   []byte("\xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb6\xa7\xa9\xfa\xa4s\xc0\xba\xa4\xd1\xb1^24\xc0Y\xa4\xfb\xbc\xc9\xc0\xc5\xa1@\xa6\xba\xa6]\xc3n\xa5\xfa\r"),
			},
		},
		{ //18
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //19
			{
				Utf8:   "4.完整新聞內文:",
				Big5:   []byte("4.\xa7\xb9\xbe\xe3\xb7s\xbbD\xa4\xba\xa4\xe5:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("4.\xa7\xb9\xbe\xe3\xb7s\xbbD\xa4\xba\xa4\xe5:\r"),
			},
		},
		{ //20
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //21
			{
				Utf8:   "生活中心／林奇樺報導",
				Big5:   []byte("\xa5\xcd\xac\xa1\xa4\xa4\xa4\xdf\xa1\xfe\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5\xcd\xac\xa1\xa4\xa4\xa4\xdf\xa1\xfe\xaaL\xa9_\xbe\xec\xb3\xf8\xbe\xc9\r"),
			},
		},
		{ //22
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //23
			{
				Utf8:   "近日頻傳陽明山國家公園野放牛隻暴斃事件，自今年以來已有24隻野放牛隻陸續死亡，對",
				Big5:   []byte("\xaa\xf1\xa4\xe9\xc0W\xb6\xc7\xb6\xa7\xa9\xfa\xa4s\xb0\xea\xaea\xa4\xbd\xb6\xe9\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xbc\xc9\xc0\xc5\xa8\xc6\xa5\xf3\xa1A\xa6\xdb\xa4\xb5\xa6~\xa5H\xa8\xd3\xa4w\xa6\xb324\xb0\xa6\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xb3\xb0\xc4\xf2\xa6\xba\xa4`\xa1A\xb9\xef"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaa\xf1\xa4\xe9\xc0W\xb6\xc7\xb6\xa7\xa9\xfa\xa4s\xb0\xea\xaea\xa4\xbd\xb6\xe9\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xbc\xc9\xc0\xc5\xa8\xc6\xa5\xf3\xa1A\xa6\xdb\xa4\xb5\xa6~\xa5H\xa8\xd3\xa4w\xa6\xb324\xb0\xa6\xb3\xa5\xa9\xf1\xa4\xfb\xb0\xa6\xb3\xb0\xc4\xf2\xa6\xba\xa4`\xa1A\xb9\xef\r"),
			},
		},
		{ //24
			{
				Utf8:   "此北市動保處也請來專業獸醫師剖檢採樣，據採檢結果顯示，死亡原因可能和食物品質不",
				Big5:   []byte("\xa6\xb9\xa5_\xa5\xab\xb0\xca\xabO\xb3B\xa4]\xbd\xd0\xa8\xd3\xb1M\xb7~\xc3~\xc2\xe5\xaev\xad\xe5\xc0\xcb\xb1\xc4\xbc\xcb\xa1A\xbe\xda\xb1\xc4\xc0\xcb\xb5\xb2\xaaG\xc5\xe3\xa5\xdc\xa1A\xa6\xba\xa4`\xad\xec\xa6]\xa5i\xaf\xe0\xa9M\xad\xb9\xaa\xab\xab~\xbd\xe8\xa4\xa3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa6\xb9\xa5_\xa5\xab\xb0\xca\xabO\xb3B\xa4]\xbd\xd0\xa8\xd3\xb1M\xb7~\xc3~\xc2\xe5\xaev\xad\xe5\xc0\xcb\xb1\xc4\xbc\xcb\xa1A\xbe\xda\xb1\xc4\xc0\xcb\xb5\xb2\xaaG\xc5\xe3\xa5\xdc\xa1A\xa6\xba\xa4`\xad\xec\xa6]\xa5i\xaf\xe0\xa9M\xad\xb9\xaa\xab\xab~\xbd\xe8\xa4\xa3\r"),
			},
		},
		{ //25
			{
				Utf8:   "佳有關，導致牛隻們長期營養獲取不足，也初步排除是動物間傳染病或是人畜共通傳染病",
				Big5:   []byte("\xa8\xce\xa6\xb3\xc3\xf6\xa1A\xbe\xc9\xadP\xa4\xfb\xb0\xa6\xad\xcc\xaa\xf8\xb4\xc1\xc0\xe7\xbei\xc0\xf2\xa8\xfa\xa4\xa3\xa8\xac\xa1A\xa4]\xaa\xec\xa8B\xb1\xc6\xb0\xa3\xacO\xb0\xca\xaa\xab\xb6\xa1\xb6\xc7\xacV\xaff\xa9\xce\xacO\xa4H\xafb\xa6@\xb3q\xb6\xc7\xacV\xaff"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa8\xce\xa6\xb3\xc3\xf6\xa1A\xbe\xc9\xadP\xa4\xfb\xb0\xa6\xad\xcc\xaa\xf8\xb4\xc1\xc0\xe7\xbei\xc0\xf2\xa8\xfa\xa4\xa3\xa8\xac\xa1A\xa4]\xaa\xec\xa8B\xb1\xc6\xb0\xa3\xacO\xb0\xca\xaa\xab\xb6\xa1\xb6\xc7\xacV\xaff\xa9\xce\xacO\xa4H\xafb\xa6@\xb3q\xb6\xc7\xacV\xaff\r"),
			},
		},
		{ //26
			{
				Utf8:   "的可能性",
				Big5:   []byte("\xaa\xba\xa5i\xaf\xe0\xa9\xca"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xaa\xba\xa5i\xaf\xe0\xa9\xca\r"),
			},
		},
		{ //27
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //28
			{
				Utf8:   "https://attach.setn.com/newsimages/2020/12/08/2924507-PH.jpg",
				Big5:   []byte("https://attach.setn.com/newsimages/2020/12/08/2924507-PH.jpg"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("https://attach.setn.com/newsimages/2020/12/08/2924507-PH.jpg\r"),
			},
		},
		{ //29
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //30
			{
				Utf8:   "5.完整新聞連結 (或短網址):",
				Big5:   []byte("5.\xa7\xb9\xbe\xe3\xb7s\xbbD\xb3s\xb5\xb2 (\xa9\xce\xb5u\xba\xf4\xa7}):"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("5.\xa7\xb9\xbe\xe3\xb7s\xbbD\xb3s\xb5\xb2 (\xa9\xce\xb5u\xba\xf4\xa7}):\r"),
			},
		},
		{ //31
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //32
			{
				Utf8:   "https://www.setn.com/News.aspx?NewsID=869094",
				Big5:   []byte("https://www.setn.com/News.aspx?NewsID=869094"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("https://www.setn.com/News.aspx?NewsID=869094\r"),
			},
		},
		{ //33
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //34
			{
				Utf8:   "6.備註:",
				Big5:   []byte("6.\xb3\xc6\xb5\xf9:"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("6.\xb3\xc6\xb5\xf9:\r"),
			},
		},
		{ //35
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //36
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //37
			{
				Utf8:   "營養不良?????????????????????",
				Big5:   []byte("\xc0\xe7\xbei\xa4\xa3\xa8}?????????????????????"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xc0\xe7\xbei\xa4\xa3\xa8}?????????????????????\r"),
			},
		},
		{ //38
			{
				Utf8:   "這有點扯...",
				Big5:   []byte("\xb3o\xa6\xb3\xc2I\xa7\xe8..."),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb3o\xa6\xb3\xc2I\xa7\xe8...\r"),
			},
		},
	}
}
