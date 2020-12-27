package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleComments struct {
	BBoardID  bbs.BBoardID  `bson:"bid"` //
	ArticleID bbs.ArticleID `bson:"aid"` //

	CommentsUpdateNanoTS types.NanoTS `bson:"comments_update_nano_ts"`
}

var (
	EMPTY_ARTICLE_COMMENTS = &ArticleComments{}
	articleCommentsFields  = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_COMMENTS)
)

func UpdateArticleComments(articleComments *ArticleComments) (err error) {

	query := bson.M{
		"$or": bson.A{
			bson.M{
				ARTICLE_BBOARD_ID_b:  articleComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: articleComments.ArticleID,
				ARTICLE_COMMENTS_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				ARTICLE_BBOARD_ID_b:  articleComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: articleComments.ArticleID,
				ARTICLE_COMMENTS_UPDATE_NANO_TS_b: bson.M{
					"$lt": articleComments.CommentsUpdateNanoTS,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	r, err := Article_c.UpdateOneOnly(query, articleComments)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
