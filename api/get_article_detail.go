package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
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
	content, contentPrefix, _, ip, host, bbs, _, _, articleDetailSummary, _, _, statusCode, err := TryGetArticleContentInfo(userID, boardID, articleID, c, false, false)
	if err != nil {
		return nil, statusCode, err
	}

	url := apitypes.ToURL(thePath.FBoardID, thePath.FArticleID)

	nickname, err := schema.GetUserNickname(articleDetailSummary.Owner)
	if err != nil {
		return nil, statusCode, err
	}

	if articleDetailSummary.IsDeleted && userID != articleDetailSummary.Owner {
		return nil, 404, ErrArticleDeleted
	}

	result = &GetArticleDetailResult{
		BBoardID:   apitypes.ToFBoardID(articleDetailSummary.BBoardID),
		ArticleID:  apitypes.ToFArticleID(articleDetailSummary.ArticleID),
		CreateTime: articleDetailSummary.CreateTime.ToTime8(),
		MTime:      articleDetailSummary.MTime.ToTime8(),
		Recommend:  articleDetailSummary.Recommend,
		NComments:  articleDetailSummary.NComments,
		Owner:      articleDetailSummary.Owner,
		Nickname:   nickname,
		Title:      apitypes.ToFTitle(articleDetailSummary.Title),
		Money:      articleDetailSummary.Money,
		Class:      articleDetailSummary.Class,
		Filemode:   articleDetailSummary.Filemode,
		Rank:       articleDetailSummary.Rank,

		URL:  url,
		Read: true,

		Brdname: boardID.ToBrdname(),

		Content:       content,
		ContentPrefix: contentPrefix,
		IP:            ip,
		Host:          host,
		BBS:           bbs,
	}

	return result, 200, nil
}

func TryGetArticleContentInfoTooSoon(updateNanoTS types.NanoTS) bool {
	nowNanoTS := types.NowNanoTS()
	return nowNanoTS-updateNanoTS < GET_ARTICLE_CONTENT_INFO_TOO_SOON_NANO_TS
}

func TryGetArticleDetailSummary(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID, articleCreateTime pttbbstypes.Time4, c *gin.Context, isSystem bool) (articleDetailSummary *schema.ArticleDetailSummary, statusCode int, err error) {
	articleDetailSummary, err = schema.GetArticleDetailSummary(boardID, articleID)
	if err != nil { // something went wrong with db.
		return nil, 500, err
	}
	if articleDetailSummary != nil {
		return articleDetailSummary, 200, nil
	}

	// init startIdx
	articleSummary := &bbs.ArticleSummary{ArticleID: articleID, CreateTime: articleCreateTime}
	startIdx := bbs.SerializeArticleIdxStr(articleSummary)

	// backend load-general-articles
	theParams_b := &pttbbsapi.LoadGeneralArticlesParams{
		StartIdx:  startIdx,
		NArticles: 1,
		Desc:      false,
		IsSystem:  isSystem,
	}
	var result_b *pttbbsapi.LoadGeneralArticlesResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_GENERAL_ARTICLES_R)
	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}
	if len(result_b.Articles) == 0 {
		return nil, 500, ErrNoArticle
	}

	article_b := result_b.Articles[0]
	if article_b.ArticleID != articleID {
		return nil, 500, ErrNoArticle
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	articleSummaries_db, _, err := deserializeArticlesAndUpdateDB(userID, boardID, result_b.Articles, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	articleSummary_db := articleSummaries_db[0]

	articleDetailSummary = &schema.ArticleDetailSummary{
		BBoardID:     boardID,
		ArticleID:    articleID,
		CreateTime:   articleSummary_db.CreateTime,
		MTime:        articleSummary_db.MTime,
		Recommend:    articleSummary_db.Recommend,
		Owner:        articleSummary_db.Owner,
		Title:        articleSummary_db.Title,
		Money:        articleSummary_db.Money,
		Class:        articleSummary_db.Class,
		Filemode:     articleSummary_db.Filemode,
		UpdateNanoTS: articleSummary_db.UpdateNanoTS,
	}

	return articleDetailSummary, 200, nil
}

func setUserReadArticle(content [][]*types.Rune, userID bbs.UUserID, articleID bbs.ArticleID, updateNanoTS types.NanoTS) {
	if content == nil {
		return
	}

	// user read article
	userReadArticle := &schema.UserReadArticle{
		UserID:       userID,
		ArticleID:    articleID,
		UpdateNanoTS: updateNanoTS,
	}
	_ = schema.UpdateUserReadArticle(userReadArticle)
}
