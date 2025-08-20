package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

var UserArticle_c *db.Collection

type UserArticle struct {
	// user-文章

	UserID           bbs.UUserID          `bson:"user_id"`
	BoardID          bbs.BBoardID         `bson:"bid"`
	ArticleID        bbs.ArticleID        `bson:"aid"`
	BoardArticleID   types.BoardArticleID `bson:"baid"`
	ReadUpdateNanoTS types.NanoTS         `bson:"update_nano_ts"`

	ArticleBlocked             bool         `bson:"article_blocked"` // 不看這篇文章
	ArticleBlockedReason       string       `bson:"article_blocked_reason"`
	ArticleBlockedUpdateNanoTS types.NanoTS `bson:"article_blocked_update_nano_ts"`

	ArticleReported             bool         `bson:"article_reported"` // 檢舉這篇文章
	ArticleReportedReason       string       `bson:"article_reported_reason"`
	ArticleReportedUpdateNanoTS types.NanoTS `bson:"article_reported_update_nano_ts"`
}

var EMPTY_USER_ARTICLE = &UserArticle{}

var (
	USER_ARTICLE_USER_ID_b             = getBSONName(EMPTY_USER_ARTICLE, "UserID")
	USER_ARTICLE_BOARD_ID_b            = getBSONName(EMPTY_USER_ARTICLE, "BoardID")
	USER_ARTICLE_ARTICLE_ID_b          = getBSONName(EMPTY_USER_ARTICLE, "ArticleID")
	USER_ARTICLE_BOARD_ARTICLE_ID_b    = getBSONName(EMPTY_USER_ARTICLE, "BoardArticleID")
	USER_ARTICLE_READ_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_ARTICLE, "ReadUpdateNanoTS")

	USER_ARTICLE_ARTICLE_BLOCKED_b                = getBSONName(EMPTY_USER_ARTICLE, "ArticleBlocked")
	USER_ARTICLE_ARTICLE_BLOCKED_REASON_b         = getBSONName(EMPTY_USER_ARTICLE, "ArticleBlockedReason")
	USER_ARTICLE_ARTICLE_BLOCKED_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_ARTICLE, "ArticleBlockedUpdateNanoTS")

	USER_ARTICLE_ARTICLE_REPORTED_b                = getBSONName(EMPTY_USER_ARTICLE, "ArticleReported")
	USER_ARTICLE_ARTICLE_REPORTED_REASON_b         = getBSONName(EMPTY_USER_ARTICLE, "ArticleReportedReason")
	USER_ARTICLE_ARTICLE_REPORTED_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_ARTICLE, "ArticleReportedUpdateNanoTS")
)

func assertUserArticleFields() error {
	if err := assertFields(EMPTY_USER_ARTICLE, EMPTY_USER_READ_ARTICLE); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER_ARTICLE, EMPTY_USER_READ_ARTICLE_QUERY); err != nil {
		return err
	}

	return nil
}

func FindUserArticles(userID bbs.UUserID, boardID bbs.BBoardID, articleIDs []bbs.ArticleID) ([]*UserArticle, error) {
	// query
	query := bson.M{
		USER_ARTICLE_USER_ID_b:  userID,
		USER_ARTICLE_BOARD_ID_b: boardID,
		USER_ARTICLE_ARTICLE_ID_b: bson.M{
			"$in": articleIDs,
		},
	}

	var dbResults []*UserArticle
	err := UserBoard_c.Find(query, 0, &dbResults, nil, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}
