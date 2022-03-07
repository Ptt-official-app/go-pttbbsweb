package dbcs

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUtf8Filename5            = "M.1644506379.A.AE6.utf8"
	testUtf8ContentAll5          []byte
	testUtf8Content5             []byte
	testUtf8Signature5           []byte
	testUtf8Comment5             []byte
	testUtf8FirstCommentsDBCS5   []byte
	testUtf8TheRestCommentsDBCS5 []byte
	testUtf8Content5Utf8         [][]*types.Rune

	testUtf8FirstComments5     []*schema.Comment
	testUtf8FullFirstComments5 []*schema.Comment
	testUtf8TheRestComments5   []*schema.Comment
)

func initTestUtf85() {
	testUtf8ContentAll5, testUtf8Content5, testUtf8Signature5, testUtf8Comment5, testUtf8FirstCommentsDBCS5, testUtf8TheRestCommentsDBCS5 = loadTest(testUtf8Filename5)

	testUtf8Content5Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:    "作者: teemocogs (teemo) 看板: Publish",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "作者: teemocogs (teemo) 看板: Publish",
			},
		},
		{ // 1
			{
				Utf8:    "標題: [閒聊] 新書的情報",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "標題: [閒聊] 新書的情報",
			},
		},
		{ // 5
			{
				Utf8:    "時間: Sat Apr 10 22:18:58 2021",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
				DBCSStr: "時間: Sat Apr 10 22:18:58 2021",
			},
		},
		{ // 3
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "  109年圖書報告出爐 台灣出版總量創20年來新低",
				DBCSStr: "  109年圖書報告出爐 台灣出版總量創20年來新低",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 3
		},
		{ // 3
		},
		{ // 4
			{
				Utf8:    "  根據年度報告內容指出，民國109年申請書號的出版社共有4694家，較108年減少258家，",
				DBCSStr: "  根據年度報告內容指出，民國109年申請書號的出版社共有4694家，較108年減少258家，",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "出版數量則減少1769種，來到3萬5041種，新書出版量等於連3年呈現下跌趨勢，從107年",
				DBCSStr: "出版數量則減少1769種，來到3萬5041種，新書出版量等於連3年呈現下跌趨勢，從107年",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 4
			{
				Utf8:    "起，每年跌幅分別為3.19%、5.89%、4.81%，且跌破90年的3萬6353種，來到新低點。",
				DBCSStr: "起，每年跌幅分別為3.19%、5.89%、4.81%，且跌破90年的3萬6353種，來到新低點。",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 4
			{
				Utf8:    "  https://www.cna.com.tw/news/firstnews/202104010163.aspx",
				DBCSStr: "  https://www.cna.com.tw/news/firstnews/202104010163.aspx",
				Color0:  types.DefaultColor,
				Color1:  types.DefaultColor,
			},
		},
		{ // 5
		},
		{ // 5
		},
		{ // 5
		},
	}

	testUtf8FirstComments5 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_DELETED,
			Owner:   bbs.UUserID("teemocogs"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:    "teemocogs 刪除 teemocogs 的推文: 誤植",
						Color0:  types.DefaultColor,
						Color1:  types.DefaultColor,
						DBCSStr: "teemocogs 刪除 teemocogs 的推文: 誤植",
					},
				},
			},
			MD5:     "01IlJdl-9TPqoasvW0iEQg",
			DBCSStr: "\x1b[1;30m(teemocogs 刪除 teemocogs 的推文: 誤植)\x1b[m",
		},
	}
}
