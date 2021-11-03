package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleIsDeleted struct {
	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}


func DeleteArticles(boardID bbs.BBoardID, articleIDs []bbs.ArticleID, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:  boardID,
		ARTICLE_ARTICLE_ID_b:     bson.M{"$in": articleIDs},
		ARTICLE_UPDATE_NANO_TS_b: bson.M{"$lt": updateNanoTS},
	}
	update := &ArticleIsDeleted{
		IsDeleted:    true,
		UpdateNanoTS: updateNanoTS,
	}
	_, err = Article_c.UpdateManyOnly(query, update)
	_, err = Comment_c.UpdateManyOnly(query, update)
	_, err = Rank_c.UpdateManyOnly(query, update)
	_, err = UserReadArticle_c.UpdateManyOnly(query, update)
	return nil
}