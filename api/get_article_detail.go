package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
)

const GET_ARTICLE_R = "/board/:bid/article/:aid"

type GetArticleDetailParams struct{}

type GetArticleDetailPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type GetArticleDetailResult struct {
	BBoardID   apitypes.FBoardID   `json:"bid"`         //
	ArticleID  apitypes.FArticleID `json:"aid"`         //
	IsDeleted  bool                `json:"deleted"`     //
	CreateTime types.Time8         `json:"create_time"` //
	MTime      types.Time8         `json:"modified"`    //
	Recommend  int                 `json:"recommend"`   //
	NComments  int                 `json:"n_comments"`  //
	Owner      bbs.UUserID         `json:"owner"`       //
	Nickname   string              `json:"nickname"`
	Title      string              `json:"title"` //
	Money      int                 `json:"money"` //
	Class      string              `json:"class"` // can be: R: 轉, [class]
	Filemode   ptttype.FileMode    `json:"mode"`  //

	URL  string `json:"url"`  //
	Read bool   `json:"read"` //

	Brdname string `json:"brdname"`

	Content       [][]*types.Rune `json:"content"`
	ContentPrefix [][]*types.Rune `json:"prefix"`

	IP   string `json:"ip"`
	Host string `json:"host"` // ip 的中文呈現, 外國則為國家.
	BBS  string `json:"bbs"`

	Rank int `json:"rank"` // 評價

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func GetArticleDetailWrapper(c *gin.Context) {
	params := &GetArticleDetailParams{}
	path := &GetArticleDetailPath{}
	LoginRequiredPathQuery(GetArticleDetail, params, path, c)
}

func GetArticleDetail(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	thePath, ok := path.(*GetArticleDetailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}
	articleID := thePath.FArticleID.ToArticleID()

	// validate user
	_, statusCode, err = isBoardValidUser(boardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	// ensure that we do have the article.
	content, contentPrefix, _, ip, host, bbs, _, _, articleDetailSummary, _, _, statusCode, err := TryGetArticleContentInfo(userID, boardID, articleID, c, false, false, true)
	if err != nil {
		return nil, statusCode, err
	}

	url := apitypes.ToURL(thePath.FBoardID, thePath.FArticleID)

	nickname, err := schema.GetUserNickname(articleDetailSummary.Owner)
	if err != nil {
		return nil, statusCode, err
	}

	result = &GetArticleDetailResult{
		BBoardID:      apitypes.ToFBoardID(articleDetailSummary.BBoardID),
		ArticleID:     apitypes.ToFArticleID(articleDetailSummary.ArticleID),
		IsDeleted:     false,
		CreateTime:    articleDetailSummary.CreateTime.ToTime8(),
		MTime:         articleDetailSummary.MTime.ToTime8(),
		Recommend:     articleDetailSummary.Recommend,
		NComments:     articleDetailSummary.NComments,
		Owner:         articleDetailSummary.Owner,
		Nickname:      nickname,
		Title:         apitypes.ToFTitle(articleDetailSummary.Title),
		Money:         articleDetailSummary.Money,
		Class:         articleDetailSummary.Class,
		Filemode:      articleDetailSummary.Filemode,
		URL:           url,
		Read:          true,
		Brdname:       boardID.ToBrdname(),
		Content:       content,
		ContentPrefix: contentPrefix,
		IP:            ip,
		Host:          host,
		BBS:           bbs,
		Rank:          articleDetailSummary.Rank,

		TokenUser: userID,
	}

	return result, 200, nil
}
