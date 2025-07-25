package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/Ptt-official-app/go-pttbbsweb/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TryLoadPttWebPopularBoards(c *gin.Context) (boards []*schema.BoardSummary, err error) {
	req, err := http.NewRequest("GET", types.PTTWEB_HOTBOARD_URL, nil)
	if err != nil {
		logrus.Errorf("unable to NewRequest: e: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(req.Context(), time.Duration(types.EXPIRE_HTTP_REQUEST_TS)*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		logrus.Errorf("unable to client.Do: e: %v", err)
		return nil, err
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		logrus.Errorf("unable to Parse body: e: %v", err)
		return nil, err
	}

	boards = make([]*schema.BoardSummary, 0, types.MAX_POPULAR_BOARDS)
	for node := range doc.Descendants() {
		if node.DataAtom == atom.Div {
			attrMap := attrsToMap(node.Attr)
			theClass, ok := attrMap["class"]
			if !ok {
				continue
			}
			if theClass.Val == "b-ent" {
				board := loadPttWebPopularBoardsParseBoard(node, c)
				if board == nil {
					continue
				}
				boards = append(boards, board)
			}
		}
	}

	updateNanoTS := types.NowNanoTS()
	for _, board := range boards {
		board.UpdateNanoTS = updateNanoTS
	}

	_ = schema.ResetBoardIsPopular()
	err = schema.UpdateBoardSummaries(boards, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func attrsToMap(attrs []html.Attribute) (ret map[string]html.Attribute) {
	ret = make(map[string]html.Attribute)
	for _, each := range attrs {
		ret[each.Key] = each
	}

	return ret
}

func loadPttWebPopularBoardsParseBoard(node *html.Node, c *gin.Context) (board *schema.BoardSummary) {
	boardName := ""
	nUser := 0
	boardClass := ""
	boardTitle := ""
	boardType := ""

	var err error
	for childNode := range node.Descendants() {
		attrMap := attrsToMap(childNode.Attr)
		theClass, ok := attrMap["class"]
		if !ok {
			continue
		}

		switch theClass.Val {
		case "board-name":
			boardName = childNode.FirstChild.Data
		case "board-nuser":
			for nUserChildNode := range childNode.Descendants() {
				if nUserChildNode.Type != html.TextNode {
					continue
				}
				nUserStr := nUserChildNode.Data
				nUser, err = strconv.Atoi(nUserStr)
				if err != nil {
					nUser = 0
				}
			}
		case "board-class":
			boardClass = childNode.FirstChild.Data
		case "board-title":
			// boardType = childNode.FirstChild.Data[:1]
			boardType = "◎"
			boardTitle = childNode.FirstChild.Data[3:]
		}
	}

	if boardName == "" {
		return nil
	}

	total, lastPostTime, err := loadPttWebPopularBoardsGetBoardInfo(boardName, c)
	if err != nil {
		logrus.Warnf("unable to get total/lastPosTime from go-pttbbs: boardName: %v e: %v", boardName, err)
	}

	board = &schema.BoardSummary{
		BBoardID:  bbs.BBoardID(boardName),
		Brdname:   boardName,
		Title:     boardTitle,
		BoardType: boardType,
		Category:  boardClass,

		NUser:        nUser,
		Total:        total,
		LastPostTime: lastPostTime,

		IsPopular: true,
	}

	return board
}

func loadPttWebPopularBoardsGetBoardInfo(boardName string, c *gin.Context) (total int, lastPostTime types.NanoTS, err error) {
	theParams_b := &pttbbsapi.LoadBoardSummaryParams{}
	urlMap := map[string]string{
		"bid": boardName,
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_BOARD_SUMMARY_R)

	var result_b pttbbsapi.LoadBoardSummaryResult
	_, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return 0, 0, err
	}

	total = int(result_b.Total)
	lastPostTime = types.Time4ToNanoTS(result_b.LastPostTime)

	return total, lastPostTime, nil
}
