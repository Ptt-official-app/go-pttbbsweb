package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/cmsys"
	"github.com/gin-gonic/gin"
)

func UpdateArticleContentInfo(
	boardID bbs.BBoardID,
	articleID bbs.ArticleID,
	content [][]*types.Rune,
	contentMD5 string,
	ip string,
	host string,
	bbs string,
	signatureMD5 string,
	signatureDBCS []byte,
	updateNanoTS types.NanoTS) (err error) {

	if contentMD5 == "" {
		return nil
	}

	contentInfo := &schema.ArticleContentInfo{
		ContentMD5: contentMD5,

		Content: content,
		IP:      ip,
		Host:    host,
		BBS:     bbs,

		SignatureDBCS: signatureDBCS,
		SignatureMD5:  signatureMD5,

		ContentUpdateNanoTS: updateNanoTS,
	}

	err = schema.UpdateArticleContentInfo(boardID, articleID, contentInfo)

	return err
}

func DeserializeArticlesAndUpdateDB(articleSummaries_b []*bbs.ArticleSummary, updateNanoTS types.NanoTS) (articleSummaries []*schema.ArticleSummaryWithRegex, err error) {
	if len(articleSummaries_b) == 0 {
		return nil, nil
	}
	articleSummaries = make([]*schema.ArticleSummaryWithRegex, len(articleSummaries_b))
	for idx, each_b := range articleSummaries_b {
		articleSummaries[idx] = schema.NewArticleSummaryWithRegex(each_b, updateNanoTS)
	}

	err = schema.UpdateArticleSummaryWithRegexes(articleSummaries, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return articleSummaries, nil
}

func deserializeArticlesAndUpdateDB(userID bbs.UUserID, bboardID bbs.BBoardID, articleSummaries_b []*bbs.ArticleSummary, updateNanoTS types.NanoTS) (articleSummaries []*schema.ArticleSummaryWithRegex, userReadArticleMap map[bbs.ArticleID]bool, err error) {
	if len(articleSummaries_b) == 0 {
		return nil, nil, nil
	}

	articleSummaries, err = DeserializeArticlesAndUpdateDB(articleSummaries_b, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadArticles := make([]*schema.UserReadArticle, 0, len(articleSummaries_b))
	userReadArticleMap = make(map[bbs.ArticleID]bool)
	for _, each_b := range articleSummaries_b {
		if each_b.Read {
			each_db := &schema.UserReadArticle{
				UserID:       userID,
				ArticleID:    each_b.ArticleID,
				UpdateNanoTS: updateNanoTS,
			}
			userReadArticles = append(userReadArticles, each_db)

			userReadArticleMap[each_db.ArticleID] = true
		}
	}

	err = schema.UpdateUserReadArticles(userReadArticles, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	// get n-comments
	updateArticleNComments(bboardID, articleSummaries)

	return articleSummaries, userReadArticleMap, err
}

func updateArticleNComments(bboardID bbs.BBoardID, articleSummaries []*schema.ArticleSummaryWithRegex) {
	if len(articleSummaries) == 0 {
		return
	}

	articleIDs := make([]bbs.ArticleID, len(articleSummaries))
	for idx, each := range articleSummaries {
		articleIDs[idx] = each.ArticleID
	}

	articleNComments, err := schema.GetArticleNCommentsByArticleIDs(bboardID, articleIDs)
	if err != nil {
		return
	}

	nCommentsByArticleIDMap := make(map[bbs.ArticleID]*schema.ArticleNComments)
	for _, each := range articleNComments {
		nCommentsByArticleIDMap[each.ArticleID] = each
	}

	for _, each := range articleSummaries {
		eachArticleNComments := nCommentsByArticleIDMap[each.ArticleID]
		if eachArticleNComments == nil {
			continue
		}

		each.NComments = eachArticleNComments.NComments
		each.Rank = eachArticleNComments.Rank
	}
}

func ArticleLockKey(boardID bbs.BBoardID, articleID bbs.ArticleID) (key string) {
	return string(boardID) + ":" + string(articleID)
}

//TryGetArticleContentInfo
//
//嘗試拿到 article-content
//
//1. 根據 article-id 得到相對應的 filename, ownerid, create-time.
//2. 嘗試從 schema 拿到 db summary 資訊. (create-time)
//3. 如果可以從 schema 拿到 db 資訊:
//   3.1. 如果已經 deleted: return deleted.
//   3.2. 如果距離上次跟 pttbbs 拿的時間太近: 從 schema 拿到 content, return schema-content.
//4. 嘗試做 lock.
//   4.1. 如果 lock 失敗: 從 schema 拿到 content, return schema-content.
//5. 從 pttbbs 拿到 article
//6. 如果從 pttbbs 拿到的時間比 schema 裡拿到的時間舊的話: return schema-content.
//7. parse article 為 content / comments.
//8. 將 comments parse 為 firstComments / theRestComments.
//9. 將 theRestComments 丟進 queue 裡.
func TryGetArticleContentInfo(userID bbs.UUserID, bboardID bbs.BBoardID, articleID bbs.ArticleID, c *gin.Context, isSystem bool, isHash bool) (content [][]*types.Rune, contentMD5 string, ip string, host string, bbs string, signatureMD5 string, signatureDBCS []byte, articleDetailSummary *schema.ArticleDetailSummary, fileSize int, hash cmsys.Fnv64_t, statusCode int, err error) {
	updateNanoTS := types.NanoTS(0)
	// set user-read-article-id
	defer func() {
		if err == nil {
			setUserReadArticle(content, userID, articleID, updateNanoTS)
		}
	}()

	isForce := false
	isQueue := true

	// if not isForce => not isHash
	if isHash {
		isForce = true
		isQueue = false
	}

	if isSystem {
		isQueue = false
	}

	// get article-info (ensuring valid article-id)
	articleFilename := articleID.ToRaw()
	articleCreateTime, err := articleFilename.CreateTime()
	if err != nil {
		return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
	}

	articleCreateTimeNanoTS := types.Time4ToNanoTS(articleCreateTime)

	// get from backend with content-mtime
	// estimated max 500ms + 3 seconds

	articleDetailSummary_db, statusCode, err := TryGetArticleDetailSummary(userID, bboardID, articleID, articleCreateTime, c, isSystem)
	if err != nil {
		return nil, "", "", "", "", "", nil, nil, 0, 0, statusCode, err
	}

	// preliminarily checking article-detail-summary.
	if articleDetailSummary_db.IsDeleted {
		return nil, "", "", "", "", "", nil, nil, 500, 0, 0, ErrAlreadyDeleted
	}

	if articleDetailSummary_db.CreateTime == 0 {
		articleDetailSummary_db.CreateTime = articleCreateTimeNanoTS
	}

	if !isForce && TryGetArticleContentInfoTooSoon(articleDetailSummary_db.ContentUpdateNanoTS) {

		contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
		if err != nil {
			return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
		}
		return contentInfo.Content, contentInfo.ContentMD5, contentInfo.IP, contentInfo.Host, contentInfo.BBS, contentInfo.SignatureMD5, contentInfo.SignatureDBCS, articleDetailSummary_db, 0, 0, 200, nil
	}

	ownerID := articleDetailSummary_db.Owner

	// 4. do lock. if failed, return the data in db.
	lockKey := ArticleLockKey(bboardID, articleID)
	err = schema.TryLock(lockKey, ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		if isForce {
			return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
		}

		contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
		if err != nil {
			return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
		}
		updateNanoTS = types.NowNanoTS()
		return contentInfo.Content, contentInfo.ContentMD5, contentInfo.IP, contentInfo.Host, contentInfo.BBS, contentInfo.SignatureMD5, contentInfo.SignatureDBCS, articleDetailSummary_db, 0, 0, 200, nil
	}
	defer func() { _ = schema.Unlock(lockKey) }()

	// 5. get article from pttbbs
	theParams_b := &pttbbsapi.GetArticleParams{
		RetrieveTS: articleDetailSummary_db.ContentMTime.ToTime4(),
		IsSystem:   isSystem,
		IsHash:     isHash,
	}
	var result_b *pttbbsapi.GetArticleResult

	urlMap := map[string]string{
		"bid": string(bboardID),
		"aid": string(articleID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.GET_ARTICLE_R)
	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, "", "", "", "", "", nil, nil, 0, 0, statusCode, err
	}

	fileSize = len(result_b.Content)
	hash = result_b.Hash

	// 6. check content-mtime (no modify from backend, no need to parse again)
	contentMTime := types.Time4ToNanoTS(result_b.MTime)
	if articleDetailSummary_db.ContentMTime >= contentMTime {
		contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
		if err != nil {
			return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
		}
		return contentInfo.Content, contentInfo.ContentMD5, contentInfo.IP, contentInfo.Host, contentInfo.BBS, contentInfo.SignatureMD5, contentInfo.SignatureDBCS, articleDetailSummary_db, fileSize, hash, 200, nil
	}

	if result_b.Content == nil { // XXX possibly the article is deleted. Need to check error-code and mark the article as deleted.
		return nil, "", "", "", "", "", nil, nil, 0, 0, 500, ErrNoArticle
	}

	// 7. parse article as content / commentsDBCS
	updateNanoTS = types.NowNanoTS()

	content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, commentsDBCS := dbcs.ParseContent(result_b.Content, articleDetailSummary_db.ContentMD5)

	// update article
	// we need update-article-content be the 1st to upload,
	// because it's possible that there is no first-comments.
	// only article-content is guaranteed.

	err = UpdateArticleContentInfo(bboardID, articleID, content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, updateNanoTS)

	if err != nil {
		return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
	}

	if contentMD5 == "" {
		contentInfo, err := schema.GetArticleContentInfo(bboardID, articleID)
		if err != nil {
			return nil, "", "", "", "", "", nil, nil, 0, 0, 500, err
		}
		content = contentInfo.Content
		ip = contentInfo.IP
		host = contentInfo.Host
		bbs = contentInfo.BBS
		signatureMD5 = contentInfo.SignatureMD5
		signatureDBCS = contentInfo.SignatureDBCS
	}

	if isQueue {
		// 8. parse comments as firstComments and theRestComments
		firstComments, firstCommentsMD5, _, err := dbcs.ParseFirstComments(
			bboardID,
			articleID,
			ownerID,
			articleCreateTimeNanoTS,
			contentMTime,
			commentsDBCS,
			articleDetailSummary_db.FirstCommentsMD5,
		)

		// update first-comments
		// possibly err because the data is too old.
		// we don't need to queue and update content-mtime if the data is too old.
		err = tryUpdateFirstComments(firstComments, firstCommentsMD5, updateNanoTS, articleDetailSummary_db)
		if err != nil {
			//if failed update: we still send the content back.
			//(no updating the content in db,
			// so the data will be re-processed again next time).
			return content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, articleDetailSummary_db, fileSize, hash, 200, nil
		}

		// 9. enqueue and n_comments
		err = queue.QueueCommentDBCS(bboardID, articleID, ownerID, commentsDBCS, articleCreateTimeNanoTS, contentMTime, updateNanoTS)
		if err != nil {
			return content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, articleDetailSummary_db, fileSize, hash, 200, nil
		}

		if articleDetailSummary_db.NComments == 0 {
			articleDetailSummary_db.NComments = len(firstComments)
		}
	} else {
		commentQueue := &queue.CommentQueue{
			BBoardID:          bboardID,
			ArticleID:         articleID,
			OwnerID:           ownerID,
			CommentDBCS:       commentsDBCS,
			ArticleCreateTime: articleCreateTimeNanoTS,
			ArticleMTime:      contentMTime,
			UpdateNanoTS:      updateNanoTS,
		}

		_ = queue.ProcessCommentQueue(commentQueue)
	}

	// everything is good, update content-mtime
	_ = schema.UpdateArticleContentMTime(bboardID, articleID, contentMTime)

	return content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, articleDetailSummary_db, fileSize, hash, 200, nil
}
