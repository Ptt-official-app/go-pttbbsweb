package api

import (
	"bytes"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/cmsys"
	"github.com/gin-gonic/gin"
)

const EDIT_ARTICLE_R = "/board/:bid/article/:aid/edit"

type EditArticleParams struct {
	PostType string          `json:"class" form:"class" url:"class"`
	Title    string          `json:"title" form:"title" url:"title"`
	Content  [][]*types.Rune `json:"content" form:"content" url:"content"`
}

type EditArticlePath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type EditArticleResult struct {
	MTime   types.Time8     `json:"modified"` //
	Content [][]*types.Rune `json:"content"`
	Title   string          `json:"title"` //
	Class   string          `json:"class"` // can be: R: 轉, [class]
}

func EditArticleDetailWrapper(c *gin.Context) {
	params := &EditArticleParams{}
	path := &EditArticlePath{}
	LoginRequiredPathJSON(EditArticleDetail, params, path, c)
}

func EditArticleDetail(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*EditArticleParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*EditArticlePath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}
	articleID := thePath.FArticleID.ToArticleID()

	_, oldSignatureDBCS, _, oldSZ, oldsum, statusCode, err := editArticleGetArticleContentInfo(userID, boardID, articleID, c)
	if err != nil {
		return nil, statusCode, err
	}

	allContentDBCS, err := editArticleCompileContent(boardID, articleID, theParams.Content, oldSignatureDBCS)
	if err != nil {
		return nil, 500, err
	}

	// 4. do lock. if failed, return the data in db.
	lockKey := ArticleLockKey(boardID, articleID)
	err = schema.TryLock(lockKey, ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		return nil, 500, err
	}
	defer func() { _ = schema.Unlock(lockKey) }()

	// edit article
	var theType []byte
	var theTitle []byte
	if theParams.Title != "" {
		theType = types.Utf8ToBig5(theParams.PostType)
		theTitle = types.Utf8ToBig5(theParams.Title)
	}

	theParams_b := &pttbbsapi.EditArticleParams{
		OldSZ:    oldSZ,
		OldSum:   oldsum,
		PostType: theType,
		Title:    theTitle,
		Content:  allContentDBCS,
	}
	var result_b *pttbbsapi.EditArticleResult

	urlMap := map[string]string{
		"bid": string(boardID),
		"aid": string(articleID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.EDIT_ARTICLE_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	updateNanoTS := types.NowNanoTS()
	content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, _ := dbcs.ParseContent(result_b.Content, "")

	// update article
	// we need update-article-content be the 1st to upload,
	// because it's possible that there is no first-comments.
	// only article-content is guaranteed.

	err = UpdateArticleContentInfo(boardID, articleID, content, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	contentMTime := types.Time4ToNanoTS(result_b.MTime)
	_ = schema.UpdateArticleContentMTime(boardID, articleID, contentMTime)

	// update article-edit-summary
	articleEditSummary := schema.NewArticleEditSummary(contentMTime, result_b.RealTitle, result_b.FullTitle, result_b.Class, updateNanoTS)
	err = schema.UpdateArticleEditSummary(boardID, articleID, articleEditSummary, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	result = &EditArticleResult{
		MTime:   contentMTime.ToTime8(),
		Content: theParams.Content,
		Title:   articleEditSummary.Title,
		Class:   articleEditSummary.Class,
	}
	return result, 200, nil
}

func editArticleGetArticleContentInfo(userID bbs.UUserID, boardID bbs.BBoardID, articleID bbs.ArticleID, c *gin.Context) (oldContent [][]*types.Rune, signatureDBCS []byte, articleDetailSummary_db *schema.ArticleDetailSummary, sz int, hash cmsys.Fnv64_t, statusCode int, err error) {
	oldContent, _, _, _, _, _, signatureDBCS, articleDetailSummary_db, sz, hash, statusCode, err = TryGetArticleContentInfo(userID, boardID, articleID, c, false, true)

	return oldContent, signatureDBCS, articleDetailSummary_db, sz, hash, statusCode, err
}

func editArticleCompileContent(boardID bbs.BBoardID, articleID bbs.ArticleID, content [][]*types.Rune, signatureDBCS []byte) (allContentDBCS [][]byte, err error) {
	commentDBCSs, err := schema.GetAllCommentDBCSs(boardID, articleID)
	if err != nil {
		return nil, err
	}

	commentsDBCS := commentDBCSsToCommentsDBCS(commentDBCSs)

	contentDBCS := dbcs.Utf8ToDBCS(content)

	signaturesDBCS := bytes.Split(signatureDBCS, []byte{'\n'})

	lenAllContentDBCS := len(contentDBCS) + len(commentsDBCS) + len(signaturesDBCS)
	allContentDBCS = make([][]byte, 0, lenAllContentDBCS)
	allContentDBCS = append(allContentDBCS, contentDBCS...)
	allContentDBCS = append(allContentDBCS, signaturesDBCS...)
	allContentDBCS = append(allContentDBCS, commentsDBCS...)

	return allContentDBCS, nil
}
