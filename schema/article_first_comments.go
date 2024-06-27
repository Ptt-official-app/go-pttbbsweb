package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleFirstComments struct {
	BBoardID  bbs.BBoardID  `bson:"bid"` //
	ArticleID bbs.ArticleID `bson:"aid"` //

	FirstCommentsMD5          string       `bson:"first_comments_md5"`
	FirstCommentsLastTime     types.NanoTS `bson:"first_comments_last_time_nano_ts"`
	FirstCommentsUpdateNanoTS types.NanoTS `bson:"first_comments_update_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"` //
}

var (
	EMPTY_ARTICLE_FIRST_COMMENTS = &ArticleFirstComments{}

	articleFirstCommentsFields = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_FIRST_COMMENTS) //nolint // consistent with programming pattern
)

func UpdateArticleFirstComments(articleFirstComments *ArticleFirstComments) (err error) {
	query := bson.M{
		"$or": bson.A{
			bson.M{
				ARTICLE_BBOARD_ID_b:  articleFirstComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: articleFirstComments.ArticleID,
				ARTICLE_FIRST_COMMENTS_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				ARTICLE_BBOARD_ID_b:  articleFirstComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: articleFirstComments.ArticleID,
				ARTICLE_FIRST_COMMENTS_UPDATE_NANO_TS_b: bson.M{
					"$lt": articleFirstComments.FirstCommentsUpdateNanoTS,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	r, err := Article_c.UpdateOneOnly(query, articleFirstComments)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
