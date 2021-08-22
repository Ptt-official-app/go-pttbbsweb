package cron

import (
	"fmt"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/queue"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TryGetArticleContentInfo(boardID bbs.BBoardID, articleID bbs.ArticleID) (err error) {
	// get article-info (ensuring valid article-id)
	articleFilename := articleID.ToRaw()
	articleCreateTime, err := articleFilename.CreateTime()
	if err != nil {
		return err
	}

	articleCreateTimeNanoTS := types.Time4ToNanoTS(articleCreateTime)

	articleDetailSummary, err := schema.GetArticleDetailSummary(boardID, articleID)
	if err != nil { // something went wrong with db.
		return err
	}
	if articleDetailSummary == nil {
		return api.ErrNoArticle
	}
	if api.TryGetArticleContentInfoTooSoon(articleDetailSummary.ContentUpdateNanoTS) {
		return nil
	}

	// 4. do lock. if failed, return the data in db.
	lockKey := api.ArticleLockKey(boardID, articleID)
	err = schema.TryLock(lockKey, api.ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		return err
	}
	defer func() { _ = schema.Unlock(lockKey) }()

	// 5. get article from pttbbs
	theParams_b := &pttbbsapi.GetArticleParams{
		RetrieveTS: articleDetailSummary.ContentMTime.ToTime4(),
		IsSystem:   true,
	}
	var result_b *pttbbsapi.GetArticleResult

	urlMap := map[string]string{
		"bid": string(boardID),
		"aid": string(articleID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.GET_ARTICLE_R)

	statusCode, err := utils.BackendGet(nil, url, theParams_b, nil, &result_b)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return fmt.Errorf("TryGetArticleContentInfo: BackendGet statusCode err: bid: %v aid: %v statusCode: %v", boardID, articleID, statusCode)
	}

	// 6. check content-mtime (no modify from backend, no need to parse again)
	contentMTime := types.Time4ToNanoTS(result_b.MTime)
	if articleDetailSummary.ContentMTime >= contentMTime {
		return nil
	}

	if result_b.Content == nil {
		return api.ErrNoArticle
	}

	// 7. parse article as content / commentsDBCS
	updateNanoTS := types.NowNanoTS()

	content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, commentsDBCS := dbcs.ParseContent(result_b.Content, articleDetailSummary.ContentMD5)

	err = api.UpdateArticleContentInfo(boardID, articleID, content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, updateNanoTS)
	if err != nil {
		return err
	}

	// 8. comments
	commentQueue := &queue.CommentQueue{
		BBoardID:          boardID,
		ArticleID:         articleID,
		OwnerID:           articleDetailSummary.Owner,
		CommentDBCS:       commentsDBCS,
		ArticleCreateTime: articleCreateTimeNanoTS,
		ArticleMTime:      contentMTime,
		UpdateNanoTS:      updateNanoTS,
	}
	return queue.ProcessCommentQueue(commentQueue)
}
