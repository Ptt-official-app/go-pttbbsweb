package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleNComments struct {
	// ArticleSummary
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`

	NComments int `bson:"n_comments,omitempty"` // n_comments is read-only in article-summary.

	Rank int `bson:"rank,omitempty"` // 評價
}

var (
	EMPTY_ARTICLE_N_COMMENTS = &ArticleNComments{}
	articleNCommentsFields   = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_N_COMMENTS)
)

// GetArticleNCommentsByArticleIDs
func GetArticleNCommentsByArticleIDs(bboardID bbs.BBoardID, articleIDs []bbs.ArticleID) (articleNComments []*ArticleNComments, err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b: bboardID,
		ARTICLE_ARTICLE_ID_b: bson.M{
			"$in": articleIDs,
		},
	}

	err = Article_c.Find(query, 0, &articleNComments, articleNCommentsFields, nil)
	if err != nil {
		return nil, err
	}

	return articleNComments, nil
}
