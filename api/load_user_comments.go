package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_USER_COMMENTS_R = "/user/:user_id/comments"

type LoadUserCommentsParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitempty"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"max,omitempty" form:"max,omitempty" url:"max,omitempty"`
}

type LoadUserCommentsPath struct {
	UserID bbs.UUserID `json:"user_id"`
}

type LoadUserCommentsResult struct {
	List    []*apitypes.ArticleComment `json:"list"`
	NextIdx string                     `json:"next_idx"`
}

func NewLoadUserCommentsParams() *LoadUserCommentsParams {
	return &LoadUserCommentsParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadUserCommentsWrapper(c *gin.Context) {
	params := NewLoadUserCommentsParams()
	path := &LoadUserCommentsPath{}
	LoginRequiredPathQuery(LoadUserComments, params, path, c)
}

func LoadUserComments(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadUserCommentsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadUserCommentsPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	commentSummaries_db, nextIdx, err := loadUserComments(thePath.UserID, theParams.StartIdx, theParams.Descending, theParams.Max)
	if err != nil {
		return nil, 500, err
	}

	articleSummaryMap, userReadArticleMap, err := getArticleMapFromCommentSummaries(userID, commentSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadUserCommentsResult(commentSummaries_db, articleSummaryMap, userReadArticleMap, nextIdx)

	return r, 200, nil
}

func loadUserComments(ownerID bbs.UUserID, startIdx string, descending bool, max int) (commentSummaries_db []*schema.CommentSummary, nextIdx string, err error) {
	commentSummaries_db = make([]*schema.CommentSummary, 0, max+1)

	var nextSortTime types.NanoTS
	if startIdx != "" {
		_, startIdx, err = apitypes.DeserializeArticleCommentIdx(startIdx)
		if err != nil {
			return nil, "", err
		}
		nextSortTime, _ = apitypes.DeserializeCommentIdx(startIdx)
	}

	isEndLoop := false
	remaining := max
	for !isEndLoop && remaining > 0 {
		eachCommentSummaries_db, err := schema.GetBasicCommentSummariesByOwnerID(ownerID, nextSortTime, descending, max+1)
		if err != nil {
			return nil, "", err
		}

		// check is-last query
		if len(eachCommentSummaries_db) < max+1 {
			isEndLoop = true
			nextSortTime = 0
		} else {
			// setup next
			nextCommentSummary := eachCommentSummaries_db[len(eachCommentSummaries_db)-1]
			eachCommentSummaries_db = eachCommentSummaries_db[:len(eachCommentSummaries_db)-1]

			nextSortTime = nextCommentSummary.SortTime
			nextIdx = apitypes.SerializeCommentIdx(nextSortTime, nextCommentSummary.CommentID)
		}

		// is-valid
		validCommentSummaries_db, err := isValidCommentSummaries(eachCommentSummaries_db)
		if err != nil {
			return nil, "", err
		}

		// append
		if len(validCommentSummaries_db) > remaining {
			nextCommentSummary := validCommentSummaries_db[remaining]
			validCommentSummaries_db = validCommentSummaries_db[:remaining]

			nextSortTime = nextCommentSummary.SortTime
			nextIdx = apitypes.SerializeCommentIdx(nextSortTime, nextCommentSummary.CommentID)
		}

		commentSummaries_db = append(commentSummaries_db, validCommentSummaries_db...)
		remaining -= len(validCommentSummaries_db)
	}

	return commentSummaries_db, nextIdx, nil
}

// isValidCommentSummaries
// XXX TODO
func isValidCommentSummaries(commentSummaries_db []*schema.CommentSummary) ([]*schema.CommentSummary, error) {
	return commentSummaries_db, nil
}

func getArticleMapFromCommentSummaries(userID bbs.UUserID, commentSummaries_db []*schema.CommentSummary) (articleSummaryMap map[bbs.ArticleID]*schema.ArticleSummary, userReadArticleMap map[bbs.ArticleID]types.NanoTS, err error) {
	articleIDs := make([]bbs.ArticleID, 0, len(commentSummaries_db))
	articleIDMap := make(map[bbs.ArticleID]bool)
	for _, each := range commentSummaries_db {
		_, ok := articleIDMap[each.ArticleID]
		if ok {
			continue
		}

		articleIDMap[each.ArticleID] = true
		articleIDs = append(articleIDs, each.ArticleID)
	}

	// article summaries
	articleSummaries, err := schema.GetArticleSummariesByArticleIDs(articleIDs)
	if err != nil {
		return nil, nil, err
	}

	articleSummaryMap = make(map[bbs.ArticleID]*schema.ArticleSummary)
	for _, each := range articleSummaries {
		articleSummaryMap[each.ArticleID] = each
	}

	// user read articles
	userReadArticles, err := schema.FindUserReadArticlesByArticleIDs(userID, articleIDs)
	if err != nil {
		return nil, nil, err
	}

	userReadArticleMap = make(map[bbs.ArticleID]types.NanoTS)
	for _, each := range userReadArticles {
		userReadArticleMap[each.ArticleID] = each.UpdateNanoTS
	}

	return articleSummaryMap, userReadArticleMap, nil
}

func NewLoadUserCommentsResult(
	commentSummaries_db []*schema.CommentSummary,
	articleSummaryMap map[bbs.ArticleID]*schema.ArticleSummary,
	userReadArticleMap map[bbs.ArticleID]types.NanoTS,
	nextIdx string) (result *LoadUserCommentsResult) {
	comments := make([]*apitypes.ArticleComment, len(commentSummaries_db))
	for idx, each := range commentSummaries_db {
		articleSummary, ok := articleSummaryMap[each.ArticleID]
		if !ok {
			continue
		}
		comments[idx] = apitypes.NewArticleCommentFromComment(articleSummary, each)

		// read
		readNanoTS, ok := userReadArticleMap[each.ArticleID]
		if !ok {
			continue
		}

		if readNanoTS > articleSummary.MTime {
			comments[idx].Read = types.READ_STATUS_MTIME
		} else if readNanoTS > each.SortTime {
			comments[idx].Read = types.READ_STATUS_COMMENT_TIME
		} else if readNanoTS > each.CreateTime {
			comments[idx].Read = types.READ_STATUS_CREATE_TIME
		} else {
			comments[idx].Read = types.READ_STATUS_UNREAD
		}
	}

	result = &LoadUserCommentsResult{
		List:    comments,
		NextIdx: nextIdx,
	}
	return result
}
