package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_ARTICLE_COMMENTS_R = "/board/:bid/article/:aid/comments"

type LoadArticleCommentsParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadArticleCommentsPath struct {
	BBoardID  bbs.BBoardID  `uri:"bid"`
	ArticleID bbs.ArticleID `uri:"aid"`
}

type LoadArticleCommentsResult struct {
	List    []*apitypes.Comment `json:"list"`
	NextIdx string              `json:"next_idx"`
}

func NewLoadArticleCommentsParams() *LoadArticleCommentsParams {
	return &LoadArticleCommentsParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadArticleCommentsWrapper(c *gin.Context) {
	params := NewLoadArticleCommentsParams()
	path := &LoadArticleCommentsPath{}
	LoginRequiredPathQuery(LoadArticleComments, params, path, c)
}

func LoadArticleComments(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadArticleCommentsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadArticleCommentsPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	//is board-valid-user
	_, statusCode, err = isBoardValidUser(thePath.BBoardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	//get comments
	querySortNanoTS, queryCommentID := loadArticleCommentsDeserializeIdx(theParams.StartIdx)

	comments_db, err := schema.GetComments(thePath.BBoardID, thePath.ArticleID, querySortNanoTS, queryCommentID, theParams.Descending, theParams.Max+1)
	if err != nil {
		return nil, 500, err
	}

	var nextComment *schema.Comment
	if len(comments_db) == theParams.Max+1 {
		nextComment = comments_db[theParams.Max]
		comments_db = comments_db[:theParams.Max]
	}

	result = NewLoadArticleCommentsResult(comments_db, nextComment)
	return result, 200, nil
}

func loadArticleCommentsDeserializeIdx(startIdx string) (sortNanoTS types.NanoTS, commentID types.CommentID) {
	theList := strings.Split(startIdx, "@")
	if len(theList) != 2 {
		return 0, ""
	}

	nanoTSInt, err := strconv.Atoi(theList[0])
	if err != nil {
		return 0, ""
	}

	return types.NanoTS(nanoTSInt), types.CommentID(theList[1])
}

func loadArticleCommentsSerializeIdx(sortNanoTS types.NanoTS, commentID types.CommentID) string {

	return fmt.Sprintf("%v@%v", sortNanoTS, commentID)
}

func NewLoadArticleCommentsResult(comments_db []*schema.Comment, nextComment *schema.Comment) (result *LoadArticleCommentsResult) {
	nextIdx := ""
	if nextComment != nil {
		nextIdx = loadArticleCommentsSerializeIdx(nextComment.SortTime, nextComment.CommentID)
	}

	comments := make([]*apitypes.Comment, len(comments_db))
	for idx, each := range comments_db {
		comments[idx] = apitypes.NewComment(each)
	}

	result = &LoadArticleCommentsResult{
		List:    comments,
		NextIdx: nextIdx,
	}
	return result
}
