package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"

	"github.com/Ptt-official-app/go-pttbbs/types/proto"
)

const GET_ARTICLE_R = "/board/:bid/article/:aid"

type GetArticleDetailParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" url:"fields,omitempty"`
}

type GetArticleDetailPath struct {
	BBoardID  bbs.BBoardID  `uri:"bid"`
	ArticleID bbs.ArticleID `uri:"aid"`
}

type GetArticleDetailResult struct {
	BBoardID   bbs.BBoardID     `json:"bid"`
	ArticleID  bbs.ArticleID    `json:"aid"`
	Brdname    string           `json:"brdname"`
	IsDeleted  bool             `json:"deleted"`
	Filename   string           `json:"filename"`
	CreateTime types.Time8      `json:"create_time"`
	MTime      types.Time8      `json:"modified"`
	Recommend  int              `json:"recommend"`
	Owner      string           `json:"owner"`
	Date       string           `json:"date"`
	Title      string           `json:"title"`
	Money      int              `json:"money"`
	Filemode   ptttype.FileMode `json:"mode"`
	URL        string           `json:"url"`
	Read       bool             `json:"read"`
	Content    proto.Content    `json:"content"`

	Class string `json:"class"`

	IP      string `json:"ip"`
	Country string `json:"country"`
	BBS     string `json:"bbs"`
}

func GetArticleDetail(remoteAddr string, userID string, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	_, ok := path.(*GetArticleDetailPath)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	result = &GetArticleDetailResult{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("1_123124"),
		Brdname:    "test1",
		IsDeleted:  false,
		Filename:   "M.1234567890.A.324",
		CreateTime: types.Time8(1234567890),
		MTime:      types.Time8(1234567889),
		Recommend:  8,
		Owner:      "okcool",
		Date:       "12/04",
		Title:      "[問題]然後呢？～",
		Class:      "問題",
		Money:      3,
		Filemode:   0,
		URL:        "http://localhost/bbs/test1/M.1234567890.A.324.html",
		Read:       false,
	}

	return result, 200, nil
}
