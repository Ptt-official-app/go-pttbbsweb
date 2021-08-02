package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

var UserReadArticle_c *db.Collection

type UserReadArticle struct {
	// 已讀文章紀錄

	UserID       bbs.UUserID   `bson:"user_id"`
	BoardID      bbs.BBoardID  `bson:"bid"`
	ArticleID    bbs.ArticleID `bson:"aid"`
	UpdateNanoTS types.NanoTS  `bson:"update_nano_ts"`
}

var EMPTY_USER_READ_ARTICLE = &UserReadArticle{}

var (
	USER_READ_ARTICLE_USER_ID_b        = getBSONName(EMPTY_USER_READ_ARTICLE, "UserID")
	USER_READ_ARTICLE_BOARD_ID_b       = getBSONName(EMPTY_USER_READ_ARTICLE, "BoardID")
	USER_READ_ARTICLE_ARTICLE_ID_b     = getBSONName(EMPTY_USER_READ_ARTICLE, "ArticleID")
	USER_READ_ARTICLE_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_READ_ARTICLE, "UpdateNanoTS")
)

func assertUserReadArticleFields() error {
	if err := assertFields(EMPTY_USER_READ_ARTICLE, EMPTY_USER_READ_ARTICLE_QUERY); err != nil {
		return err
	}

	return nil
}

type UserReadArticleQuery struct {
	UserID    bbs.UUserID   `bson:"user_id"`
	BoardID   bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`
}

var EMPTY_USER_READ_ARTICLE_QUERY = &UserReadArticleQuery{}

func UpdateUserReadArticle(userReadArticle *UserReadArticle) (err error) {
	query := bson.M{
		USER_READ_ARTICLE_USER_ID_b:    userReadArticle.UserID,
		USER_READ_ARTICLE_BOARD_ID_b:   userReadArticle.BoardID,
		USER_READ_ARTICLE_ARTICLE_ID_b: userReadArticle.ArticleID,
	}

	r, err := UserReadArticle_c.CreateOnly(query, userReadArticle)
	if err != nil {
		return err
	}
	if r.UpsertedCount == 1 { // created, no need to update.
		return nil
	}

	query[USER_READ_ARTICLE_UPDATE_NANO_TS_b] = bson.M{
		"$lt": userReadArticle.UpdateNanoTS,
	}

	_, err = UserReadArticle_c.UpdateOneOnly(query, userReadArticle)

	return err
}

func UpdateUserReadArticles(userReadArticles []*UserReadArticle, updateNanoTS types.NanoTS) (err error) {
	if len(userReadArticles) == 0 {
		return nil
	}

	theList := make([]*db.UpdatePair, len(userReadArticles))
	for idx, each := range userReadArticles {
		query := &UserReadArticleQuery{
			UserID:    each.UserID,
			BoardID:   each.BoardID,
			ArticleID: each.ArticleID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := UserReadArticle_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(userReadArticles)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateUserReadArticles := make([]*db.UpdatePair, 0, len(userReadArticles))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*UserReadArticleQuery)
		filter := bson.M{
			USER_READ_ARTICLE_USER_ID_b:    origFilter.UserID,
			USER_READ_ARTICLE_ARTICLE_ID_b: origFilter.ArticleID,
			USER_READ_ARTICLE_UPDATE_NANO_TS_b: bson.M{
				"$lt": updateNanoTS,
			},
		}
		each.Filter = filter
		updateUserReadArticles = append(updateUserReadArticles, each)
	}

	_, err = UserReadArticle_c.BulkUpdateOneOnly(updateUserReadArticles)

	return err
}

func FindUserReadArticles(userID bbs.UUserID, boardID bbs.BBoardID, articleIDs []bbs.ArticleID) ([]*UserReadArticle, error) {
	// query
	query := bson.M{
		USER_READ_ARTICLE_USER_ID_b:  userID,
		USER_READ_ARTICLE_BOARD_ID_b: boardID,
		USER_READ_ARTICLE_ARTICLE_ID_b: bson.M{
			"$in": articleIDs,
		},
	}

	var dbResults []*UserReadArticle
	err := UserReadArticle_c.Find(query, 0, &dbResults, nil, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}

func FindUserReadArticlesByArticleIDs(userID bbs.UUserID, articleIDs []bbs.ArticleID) ([]*UserReadArticle, error) {
	// query
	query := bson.M{
		USER_READ_ARTICLE_USER_ID_b: userID,
		USER_READ_ARTICLE_ARTICLE_ID_b: bson.M{
			"$in": articleIDs,
		},
	}

	var dbResults []*UserReadArticle
	err := UserReadArticle_c.Find(query, 0, &dbResults, nil, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}
