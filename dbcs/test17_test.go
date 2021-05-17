package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	testFilename17            = "temp12"
	testContentAll17          []byte
	testContent17             []byte
	testSignature17           []byte
	testComment17             []byte
	testFirstCommentsDBCS17   []byte
	testTheRestCommentsDBCS17 []byte
	testContent17Big5         [][]*types.Rune
	testContent17Utf8         [][]*types.Rune

	testFirstComments17     []*schema.Comment
	testFullFirstComments17 []*schema.Comment
)

func initTest17() {
	testContentAll17, testContent17, testSignature17, testComment17, testFirstCommentsDBCS17, testTheRestCommentsDBCS17 = loadTest(testFilename17)

	testContent17Big5 = [][]*types.Rune{
		{ //0
			{

				Big5:   []byte("\xa7@\xaa\xcc: teemocogs (teemo) \xac\xdd\xaaO: Publish"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: teemocogs (teemo) \xac\xdd\xaaO: Publish\r"),
			},
		},
		{ //1
			{

				Big5:   []byte("\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xb7s\xae\xd1\xaa\xba\xb1\xa1\xb3\xf8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xb7s\xae\xd1\xaa\xba\xb1\xa1\xb3\xf8\r"),
			},
		},
		{ //2
			{

				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Apr 10 22:18:58 2021"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sat Apr 10 22:18:58 2021\r"),
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

				Big5:   []byte("  109\xa6~\xb9\xcf\xae\xd1\xb3\xf8\xa7i\xa5X\xc4l \xa5x\xc6W\xa5X\xaa\xa9\xc1`\xb6q\xb3\xd020\xa6~\xa8\xd3\xb7s\xa7C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  109\xa6~\xb9\xcf\xae\xd1\xb3\xf8\xa7i\xa5X\xc4l \xa5x\xc6W\xa5X\xaa\xa9\xc1`\xb6q\xb3\xd020\xa6~\xa8\xd3\xb7s\xa7C\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //8
			{

				Big5:   []byte("  \xae\xda\xbe\xda\xa6~\xab\xd7\xb3\xf8\xa7i\xa4\xba\xaee\xab\xfc\xa5X\xa1A\xa5\xc1\xb0\xea109\xa6~\xa5\xd3\xbd\xd0\xae\xd1\xb8\xb9\xaa\xba\xa5X\xaa\xa9\xaa\xc0\xa6@\xa6\xb34694\xaea\xa1A\xb8\xfb108\xa6~\xb4\xee\xa4\xd6258\xaea\xa1A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xae\xda\xbe\xda\xa6~\xab\xd7\xb3\xf8\xa7i\xa4\xba\xaee\xab\xfc\xa5X\xa1A\xa5\xc1\xb0\xea109\xa6~\xa5\xd3\xbd\xd0\xae\xd1\xb8\xb9\xaa\xba\xa5X\xaa\xa9\xaa\xc0\xa6@\xa6\xb34694\xaea\xa1A\xb8\xfb108\xa6~\xb4\xee\xa4\xd6258\xaea\xa1A\r"),
			},
		},
		{ //9
			{

				Big5:   []byte("\xa5X\xaa\xa9\xbc\xc6\xb6q\xabh\xb4\xee\xa4\xd61769\xba\xd8\xa1A\xa8\xd3\xa8\xec3\xb8U5041\xba\xd8\xa1A\xb7s\xae\xd1\xa5X\xaa\xa9\xb6q\xb5\xa5\xa9\xf3\xb3s3\xa6~\xa7e\xb2{\xa4U\xb6^\xc1\xcd\xb6\xd5\xa1A\xb1q107\xa6~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5X\xaa\xa9\xbc\xc6\xb6q\xabh\xb4\xee\xa4\xd61769\xba\xd8\xa1A\xa8\xd3\xa8\xec3\xb8U5041\xba\xd8\xa1A\xb7s\xae\xd1\xa5X\xaa\xa9\xb6q\xb5\xa5\xa9\xf3\xb3s3\xa6~\xa7e\xb2{\xa4U\xb6^\xc1\xcd\xb6\xd5\xa1A\xb1q107\xa6~\r"),
			},
		},
		{ //10
			{

				Big5:   []byte("\xb0_\xa1A\xa8C\xa6~\xb6^\xb4T\xa4\xc0\xa7O\xac\xb03.19%\xa1B5.89%\xa1B4.81%\xa1A\xa5B\xb6^\xaf}90\xa6~\xaa\xba3\xb8U6353\xba\xd8\xa1A\xa8\xd3\xa8\xec\xb7s\xa7C\xc2I\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb0_\xa1A\xa8C\xa6~\xb6^\xb4T\xa4\xc0\xa7O\xac\xb03.19%\xa1B5.89%\xa1B4.81%\xa1A\xa5B\xb6^\xaf}90\xa6~\xaa\xba3\xb8U6353\xba\xd8\xa1A\xa8\xd3\xa8\xec\xb7s\xa7C\xc2I\xa1C\r"),
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

				Big5:   []byte("  https://www.cna.com.tw/news/firstnews/202104010163.aspx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  https://www.cna.com.tw/news/firstnews/202104010163.aspx\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
	}

	testContent17Utf8 = [][]*types.Rune{
		{ //0
			{
				Utf8:   "作者: teemocogs (teemo) 看板: Publish",
				Big5:   []byte("\xa7@\xaa\xcc: teemocogs (teemo) \xac\xdd\xaaO: Publish"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: teemocogs (teemo) \xac\xdd\xaaO: Publish\r"),
			},
		},
		{ //1
			{
				Utf8:   "標題: [閒聊] 新書的情報",
				Big5:   []byte("\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xb7s\xae\xd1\xaa\xba\xb1\xa1\xb3\xf8"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xb6\xa2\xb2\xe1] \xb7s\xae\xd1\xaa\xba\xb1\xa1\xb3\xf8\r"),
			},
		},
		{ //2
			{
				Utf8:   "時間: Sat Apr 10 22:18:58 2021",
				Big5:   []byte("\xae\xc9\xb6\xa1: Sat Apr 10 22:18:58 2021"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Sat Apr 10 22:18:58 2021\r"),
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
				Utf8:   "  109年圖書報告出爐 台灣出版總量創20年來新低",
				Big5:   []byte("  109\xa6~\xb9\xcf\xae\xd1\xb3\xf8\xa7i\xa5X\xc4l \xa5x\xc6W\xa5X\xaa\xa9\xc1`\xb6q\xb3\xd020\xa6~\xa8\xd3\xb7s\xa7C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  109\xa6~\xb9\xcf\xae\xd1\xb3\xf8\xa7i\xa5X\xc4l \xa5x\xc6W\xa5X\xaa\xa9\xc1`\xb6q\xb3\xd020\xa6~\xa8\xd3\xb7s\xa7C\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //8
			{
				Utf8:   "  根據年度報告內容指出，民國109年申請書號的出版社共有4694家，較108年減少258家，",
				Big5:   []byte("  \xae\xda\xbe\xda\xa6~\xab\xd7\xb3\xf8\xa7i\xa4\xba\xaee\xab\xfc\xa5X\xa1A\xa5\xc1\xb0\xea109\xa6~\xa5\xd3\xbd\xd0\xae\xd1\xb8\xb9\xaa\xba\xa5X\xaa\xa9\xaa\xc0\xa6@\xa6\xb34694\xaea\xa1A\xb8\xfb108\xa6~\xb4\xee\xa4\xd6258\xaea\xa1A"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  \xae\xda\xbe\xda\xa6~\xab\xd7\xb3\xf8\xa7i\xa4\xba\xaee\xab\xfc\xa5X\xa1A\xa5\xc1\xb0\xea109\xa6~\xa5\xd3\xbd\xd0\xae\xd1\xb8\xb9\xaa\xba\xa5X\xaa\xa9\xaa\xc0\xa6@\xa6\xb34694\xaea\xa1A\xb8\xfb108\xa6~\xb4\xee\xa4\xd6258\xaea\xa1A\r"),
			},
		},
		{ //9
			{
				Utf8:   "出版數量則減少1769種，來到3萬5041種，新書出版量等於連3年呈現下跌趨勢，從107年",
				Big5:   []byte("\xa5X\xaa\xa9\xbc\xc6\xb6q\xabh\xb4\xee\xa4\xd61769\xba\xd8\xa1A\xa8\xd3\xa8\xec3\xb8U5041\xba\xd8\xa1A\xb7s\xae\xd1\xa5X\xaa\xa9\xb6q\xb5\xa5\xa9\xf3\xb3s3\xa6~\xa7e\xb2{\xa4U\xb6^\xc1\xcd\xb6\xd5\xa1A\xb1q107\xa6~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa5X\xaa\xa9\xbc\xc6\xb6q\xabh\xb4\xee\xa4\xd61769\xba\xd8\xa1A\xa8\xd3\xa8\xec3\xb8U5041\xba\xd8\xa1A\xb7s\xae\xd1\xa5X\xaa\xa9\xb6q\xb5\xa5\xa9\xf3\xb3s3\xa6~\xa7e\xb2{\xa4U\xb6^\xc1\xcd\xb6\xd5\xa1A\xb1q107\xa6~\r"),
			},
		},
		{ //10
			{
				Utf8:   "起，每年跌幅分別為3.19%、5.89%、4.81%，且跌破90年的3萬6353種，來到新低點。",
				Big5:   []byte("\xb0_\xa1A\xa8C\xa6~\xb6^\xb4T\xa4\xc0\xa7O\xac\xb03.19%\xa1B5.89%\xa1B4.81%\xa1A\xa5B\xb6^\xaf}90\xa6~\xaa\xba3\xb8U6353\xba\xd8\xa1A\xa8\xd3\xa8\xec\xb7s\xa7C\xc2I\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb0_\xa1A\xa8C\xa6~\xb6^\xb4T\xa4\xc0\xa7O\xac\xb03.19%\xa1B5.89%\xa1B4.81%\xa1A\xa5B\xb6^\xaf}90\xa6~\xaa\xba3\xb8U6353\xba\xd8\xa1A\xa8\xd3\xa8\xec\xb7s\xa7C\xc2I\xa1C\r"),
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
				Utf8:   "  https://www.cna.com.tw/news/firstnews/202104010163.aspx",
				Big5:   []byte("  https://www.cna.com.tw/news/firstnews/202104010163.aspx"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("  https://www.cna.com.tw/news/firstnews/202104010163.aspx\r"),
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
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ //15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
	}

	testFirstComments17 = []*schema.Comment{
		{ //0
			TheType: types.COMMENT_TYPE_DELETED,
			Owner:   bbs.UUserID("teemocogs"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "teemocogs 的推文: 誤植",
						Big5:   []byte("teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3"),
					},
				},
			},
			MD5:  "k1p8LprlcctQ5v2JmITfsA",
			DBCS: []byte("\x1b[1;30m(teemocogs \xa7R\xb0\xa3 teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m\r"),
		},
	}

	testFullFirstComments17 = []*schema.Comment{
		{ //0
			BBoardID:  "test",
			ArticleID: "test17",
			TheType:   types.COMMENT_TYPE_DELETED,
			Owner:     bbs.UUserID("teemocogs"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "teemocogs 的推文: 誤植",
						Big5:   []byte("teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3"),
					},
				},
			},
			MD5:        "k1p8LprlcctQ5v2JmITfsA",
			DBCS:       []byte("\x1b[1;30m(teemocogs \xa7R\xb0\xa3 teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m\r"),
			CommentID:  "FnSE8UpG9kA:k1p8LprlcctQ5v2JmITfsA",
			CreateTime: 1618064338001000000,
			SortTime:   1618064338001000000,
		},
	}
}
