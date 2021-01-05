package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const GET_ARTICLE_R = "/board/:bid/article/:aid"

type GetArticleDetailParams struct {
}

type GetArticleDetailPath struct {
	BBoardID  bbs.BBoardID  `uri:"bid"`
	ArticleID bbs.ArticleID `uri:"aid"`
}

type GetArticleDetailResult struct {
	BBoardID   bbs.BBoardID     `json:"bid"`         //0
	ArticleID  bbs.ArticleID    `json:"aid"`         //1
	IsDeleted  bool             `json:"deleted"`     //2
	CreateTime types.Time8      `json:"create_time"` //3
	MTime      types.Time8      `json:"modified"`    //4
	Recommend  int              `json:"recommend"`   //5
	Owner      bbs.UUserID      `json:"owner"`       //6
	Title      string           `json:"title"`       //7
	Money      int              `json:"money"`       //8
	Class      string           `json:"class"`       //can be: R: 轉, [class]
	Filemode   ptttype.FileMode `json:"mode"`        //10

	URL  string `json:"url"`  //11
	Read bool   `json:"read"` //12

	Brdname string          `json:"brdname"`
	Content [][]*types.Rune `json:"content"`
	IP      string          `json:"ip"`
	Host    string          `json:"host"` //ip 的中文呈現, 外國則為國家.
	BBS     string          `json:"bbs"`
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

	// ensure that we do have the article.
	content, ip, host, bbs, articleDetailSummary, err := tryGetArticleContentInfo(userID, thePath.BBoardID, thePath.ArticleID, c)
	if err != nil {
		return nil, 400, err
	}

	url := apitypes.ToURL(articleDetailSummary.BBoardID, articleDetailSummary.ArticleID)

	result = &GetArticleDetailResult{
		BBoardID:   articleDetailSummary.BBoardID,
		ArticleID:  articleDetailSummary.ArticleID,
		CreateTime: articleDetailSummary.CreateTime.ToTime8(),
		MTime:      articleDetailSummary.MTime.ToTime8(),
		Recommend:  articleDetailSummary.Recommend,
		Owner:      articleDetailSummary.Owner,
		Title:      articleDetailSummary.Title,
		Money:      articleDetailSummary.Money,
		Class:      articleDetailSummary.Class,
		Filemode:   articleDetailSummary.Filemode,

		URL:  url,
		Read: true,

		Brdname: thePath.BBoardID.ToBrdname(),

		Content: content,
		IP:      ip,
		Host:    host,
		BBS:     bbs,
	}

	return result, 200, nil
}

func tryGetArticleContentInfo(userID bbs.UUserID, bboardID bbs.BBoardID, articleID bbs.ArticleID, c *gin.Context) (content [][]*types.Rune, ip string, host string, bbs string, articleDetailSummary *schema.ArticleDetailSummary, err error) {

	updateNanoTS := types.NanoTS(0)

	// set user-read-article-id
	defer func() {
		setUserReadArticle(content, userID, articleID, updateNanoTS)
	}()

	//get article-info (ensuring valid article-id)
	articleFilename, ownerID := articleID.ToRaw()
	articleCreateTime, err := articleFilename.CreateTime()
	if err != nil {
		return nil, "", "", "", nil, err
	}

	articleCreateTimeNanoTS := types.Time4ToNanoTS(articleCreateTime)

	//get from backend with content-mtime
	//estimated max 500ms + 3 seconds
	articleDetailSummary_db, err := schema.GetArticleDetailSummary(bboardID, articleID)
	if err != nil { //something went wrong with db.
		return nil, "", "", "", nil, err
	}
	if articleDetailSummary_db != nil {
		if articleDetailSummary_db.IsDeleted {
			return nil, "", "", "", nil, ErrAlreadyDeleted
		}

		if len(articleDetailSummary_db.Owner) == 0 && articleDetailSummary_db.CreateTime == 0 {
			articleDetailSummary_db.Owner = ownerID
			articleDetailSummary_db.CreateTime = articleCreateTimeNanoTS
		}

		if tryGetArticleContentInfoTooSoon(articleDetailSummary_db.ContentUpdateNanoTS) {
			contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
			if err != nil {
				return nil, "", "", "", nil, err
			}
			return contentInfo.Content, contentInfo.IP, contentInfo.Host, contentInfo.BBS, articleDetailSummary_db, nil
		}

	} else {
		articleDetailSummary_db = &schema.ArticleDetailSummary{
			BBoardID:   bboardID,
			ArticleID:  articleID,
			CreateTime: articleCreateTimeNanoTS,
			Owner:      ownerID,
		}
	}

	//do lock. if failed, return the data in db.
	lockKey := string(bboardID) + ":" + string(articleID)
	err = schema.TryLock(lockKey, ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
		if err != nil {
			return nil, "", "", "", nil, err
		}
		updateNanoTS = types.NowNanoTS()
		return contentInfo.Content, contentInfo.IP, contentInfo.Host, contentInfo.BBS, articleDetailSummary_db, nil
	}
	defer schema.Unlock(lockKey)

	theParams_b := &pttbbsapi.GetArticleParams{
		RetrieveTS: articleDetailSummary_db.ContentMTime.ToTime4(),
	}
	var result_b *pttbbsapi.GetArticleResult

	urlMap := make(map[string]string)
	urlMap["bid"] = string(bboardID)
	urlMap["aid"] = string(articleID)
	url := utils.MergeURL(urlMap, pttbbsapi.GET_ARTICLE_R)

	statusCode, err := utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, "", "", "", nil, err
	}
	if statusCode != 200 {
		return nil, "", "", "", nil, ErrInvalidBackendStatusCode
	}

	//check content-mtime (no modify from backend, no need to parse again)

	contentMTime := types.Time4ToNanoTS(result_b.MTime)
	if articleDetailSummary_db.ContentMTime >= contentMTime {
		contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
		if err != nil {
			return nil, "", "", "", nil, err
		}
		return contentInfo.Content, contentInfo.IP, contentInfo.Host, contentInfo.BBS, articleDetailSummary_db, nil
	}

	if result_b.Content == nil { //XXX possibly the article is deleted. Need to check error-code and mark the article as deleted.
		return nil, "", "", "", nil, ErrNoArticle
	}

	//parse content

	content, contentMD5, ip, host, bbs, commentsDBCS := dbcs.ParseContent(result_b.Content, articleDetailSummary_db.ContentMD5)

	//update article
	//we need update-article-content be the 1st to upload,
	//because it's possible that there is no first-comments.
	//only article-content is guaranteed.
	contentInfo := &schema.ArticleContentInfo{
		ContentMD5: contentMD5,

		Content: content,
		IP:      ip,
		Host:    host,
		BBS:     bbs,

		ContentUpdateNanoTS: updateNanoTS,
	}

	err = schema.UpdateArticleContentInfo(bboardID, articleID, contentInfo)

	//parse comments
	updateNanoTS = types.NowNanoTS()

	firstComments, firstCommentsMD5, firstCommentsLastTime, theRestComments := dbcs.ParseFirstComments(
		bboardID,
		articleID,
		ownerID,
		articleCreateTimeNanoTS,
		commentsDBCS,
		articleDetailSummary_db.FirstCommentsMD5,
		articleDetailSummary_db.FirstCommentsLastTime,
		updateNanoTS,
	)

	//update first-comments
	//possibly err because the data is too old.
	//we don't need to queue and update content-mtime if the data is too old.
	err = tryUpdateFirstComments(firstComments, firstCommentsMD5, firstCommentsLastTime, updateNanoTS, articleDetailSummary_db)
	if err != nil {
		return content, ip, host, bbs, articleDetailSummary_db, nil
	}

	//if failed update: we still send the content back.
	//(no updating the content in db,
	// so the data will be re-processed again next time).
	if err != nil {
		return content, ip, host, bbs, articleDetailSummary_db, nil
	}

	//enqueue.
	if theRestComments != nil {
		err = queue.QueueCommentDBCS(bboardID, articleID, ownerID, theRestComments, firstCommentsLastTime, updateNanoTS)
		if err != nil {
			return content, ip, host, bbs, articleDetailSummary_db, nil
		}
	}

	//everything is good, update content-mtime
	_ = schema.UpdateArticleContentMTime(bboardID, articleID, contentMTime)

	return content, ip, host, bbs, articleDetailSummary_db, nil
}

func tryGetArticleContentInfoTooSoon(updateNanoTS types.NanoTS) bool {
	nowNanoTS := types.NowNanoTS()
	return nowNanoTS-updateNanoTS < GET_ARTICLE_CONTENT_INFO_TOO_SOON_NANO_TS
}

func setUserReadArticle(content [][]*types.Rune, userID bbs.UUserID, articleID bbs.ArticleID, updateNanoTS types.NanoTS) {
	if content == nil {
		return
	}

	//user read article
	userReadArticle := &schema.UserReadArticle{
		UserID:       userID,
		ArticleID:    articleID,
		UpdateNanoTS: updateNanoTS,
	}
	_ = schema.UpdateUserReadArticle(userReadArticle)
}
