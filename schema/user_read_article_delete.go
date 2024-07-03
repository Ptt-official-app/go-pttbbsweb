package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserReadArticleIsDeleted struct {
	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

// DeleteUserReadArticles deletes user_read_articles in Database
func DeleteUserReadArticles(boardID bbs.BBoardID, articleIDs []bbs.ArticleID, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:      boardID,
		ARTICLE_ARTICLE_ID_b:     bson.M{"$in": articleIDs},
		ARTICLE_UPDATE_NANO_TS_b: bson.M{"$lt": updateNanoTS},
	}
	update := &UserReadArticleIsDeleted{
		IsDeleted:    true,
		UpdateNanoTS: updateNanoTS,
	}
	_, err = UserReadArticle_c.UpdateManyOnly(query, update)
	return err
}
