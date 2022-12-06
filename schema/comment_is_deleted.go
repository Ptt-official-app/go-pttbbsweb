package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentIsDeleted struct {
	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var (
	EMPTY_COMMENT_IS_DELETED = &CommentIsDeleted{}
	commentIsDeletedFields   = getFields(EMPTY_COMMENT, EMPTY_COMMENT_IS_DELETED) //nolint // consistent with programming pattern
)

// RemoveCommentIDs
//
// Assuming that the data already exists,
// no need to check that update-nano-ts does not exists.
func RemoveCommentIDs(bboardID bbs.BBoardID, articleID bbs.ArticleID, toRemoveCommentIDs []types.CommentID, updateNanoTS types.NanoTS, reason string) (err error) {
	if len(toRemoveCommentIDs) == 0 {
		return nil
	}

	query := bson.M{
		COMMENT_BBOARD_ID_b:      bboardID,
		COMMENT_ARTICLE_ID_b:     articleID,
		COMMENT_COMMENT_ID_b:     bson.M{"$in": toRemoveCommentIDs},
		COMMENT_UPDATE_NANO_TS_b: bson.M{"$lt": updateNanoTS},
	}

	update := &CommentIsDeleted{
		IsDeleted:    true,
		UpdateNanoTS: updateNanoTS,
	}

	_, err = Comment_c.UpdateManyOnly(query, update)

	return err
}

// RecoverCommentIDs
//
// We need to use $unset to remove IS_DELETED
// IS_DELETED should never be false. (we use omitempty to avoid this issue.)
func RecoverCommentIDs(commentQueries []*CommentQuery) (err error) {
	theList := make([]*db.UpdatePair, len(commentQueries))

	update := bson.M{
		"$unset": bson.M{
			COMMENT_IS_DELETED_b:    true,
			COMMENT_DELETE_REASON_b: true,
		},
	}
	for idx, each := range commentQueries {
		theList[idx] = &db.UpdatePair{Filter: each, Update: update}
	}

	_, err = Comment_c.BulkUpdateOneOnlyNoSet(theList)

	return err
}

// DeleteCommentsByArticles deletes all comments in deleted articles
func DeleteCommentsByArticles(boardID bbs.BBoardID, articleIDs []bbs.ArticleID, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:      boardID,
		ARTICLE_ARTICLE_ID_b:     bson.M{"$in": articleIDs},
		ARTICLE_UPDATE_NANO_TS_b: bson.M{"$lt": updateNanoTS},
	}
	update := &CommentIsDeleted{
		IsDeleted:    true,
		UpdateNanoTS: updateNanoTS,
	}
	_, err = Comment_c.UpdateManyOnly(query, update)
	return err
}
