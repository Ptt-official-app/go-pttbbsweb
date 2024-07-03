package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/gin-gonic/gin"
)

const LOAD_ARTICLE_COMMENTS_R = "/board/:bid/article/:aid/comments"

type LoadArticleCommentsParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadArticleCommentsPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type LoadArticleCommentsResult struct {
	List    []*apitypes.Comment `json:"list"`
	NextIdx string              `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
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

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}
	articleID := thePath.FArticleID.ToArticleID()

	// is board-valid-user
	_, statusCode, err = isBoardValidUser(boardID, c)
	if err != nil {
		return nil, statusCode, err
	}

	// get comments
	querySortNanoTS, queryCommentID := apitypes.DeserializeCommentIdx(theParams.StartIdx)

	comments_db, err := schema.GetComments(boardID, articleID, querySortNanoTS, queryCommentID, theParams.Descending, theParams.Max+1)
	if err != nil {
		return nil, 500, err
	}

	var nextComment *schema.Comment
	if len(comments_db) == theParams.Max+1 {
		nextComment = comments_db[theParams.Max]
		comments_db = comments_db[:theParams.Max]
	}

	result = NewLoadArticleCommentsResult(comments_db, nextComment, userID)
	return result, 200, nil
}

func NewLoadArticleCommentsResult(comments_db []*schema.Comment, nextComment *schema.Comment, userID bbs.UUserID) (result *LoadArticleCommentsResult) {
	nextIdx := ""
	if nextComment != nil {
		nextIdx = apitypes.SerializeCommentIdx(nextComment.SortTime, nextComment.CommentID)
	}

	comments := make([]*apitypes.Comment, len(comments_db))
	for idx, each := range comments_db {
		comments[idx] = apitypes.NewComment(each, "")
	}

	result = &LoadArticleCommentsResult{
		List:      comments,
		NextIdx:   nextIdx,
		TokenUser: userID,
	}
	return result
}
