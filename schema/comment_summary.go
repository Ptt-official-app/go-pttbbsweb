package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentSummary struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`

	CommentID types.CommentID     `bson:"cid"`
	TheType   ptttype.CommentType `bson:"type"`
	IsDeleted bool                `bson:"deleted,omitempty"`

	CreateTime   types.NanoTS `bson:"create_time_nano_ts"`
	Owner        bbs.UUserID  `bson:"owner"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	Content  [][]*types.Rune `bson:"content"` // content in comment is
	SortTime types.NanoTS    `bson:"sort_time_nano_ts"`
}

var (
	EMPTY_COMMENT_SUMMARY = &CommentSummary{}
	commentSummaryFields  = getFields(EMPTY_COMMENT, EMPTY_COMMENT_SUMMARY)
)

// GetCommentSummaries
//
// get comment summaries with startNanoTS <= createTime  < endNanoTS (excluding endNanoTS)
func GetCommentSummaries(bboardID bbs.BBoardID, articleID bbs.ArticleID, startNanoTS types.NanoTS, endNanoTS types.NanoTS) (commentSummaries []*CommentSummary, err error) {
	query := bson.M{
		COMMENT_BBOARD_ID_b:  bboardID,
		COMMENT_ARTICLE_ID_b: articleID,
		COMMENT_CREATE_TIME_b: bson.M{
			"$gte": startNanoTS,
			"$lt":  endNanoTS,
		},
	}

	err = Comment_c.Find(query, 0, &commentSummaries, commentSummaryFields, nil)
	if err != nil {
		return nil, err
	}

	return commentSummaries, nil
}

func GetBasicCommentSummariesByOwnerID(ownerID bbs.UUserID, startSortTime types.NanoTS, descending bool, limit int) (result []*CommentSummary, err error) {
	// setup query
	var query bson.M
	if startSortTime == 0 {
		query = bson.M{
			COMMENT_OWNER_b: ownerID,
			COMMENT_THE_TYPE_b: bson.M{
				"$lte": ptttype.COMMENT_TYPE_BASIC,
			},
			COMMENT_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	} else {
		theDirCommentID := "$gte"
		if descending {
			theDirCommentID = "$lte"
		}

		query = bson.M{
			COMMENT_OWNER_b: ownerID,
			COMMENT_THE_TYPE_b: bson.M{
				"$lte": ptttype.COMMENT_TYPE_BASIC,
			},
			COMMENT_SORT_TIME_b: bson.M{
				theDirCommentID: startSortTime,
			},
			COMMENT_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	}
	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: COMMENT_SORT_TIME_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: COMMENT_SORT_TIME_b, Value: 1},
		}
	}

	// find
	err = Comment_c.Find(query, int64(limit), &result, nil, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateCommentSummaries(bboardID bbs.BBoardID, articleID bbs.ArticleID, commentSummaries []*CommentSummary, updateNanoTS types.NanoTS) (err error) {
	if len(commentSummaries) == 0 {
		return nil
	}

	p_commentSummaries := commentSummaries

	var first []*CommentSummary
	for block := getBlock(len(p_commentSummaries), MAX_COMMENT_SUMMARY_BLOCK); len(p_commentSummaries) > 0; block = getBlock(len(p_commentSummaries), MAX_COMMENT_SUMMARY_BLOCK) {
		first, p_commentSummaries = p_commentSummaries[:block], p_commentSummaries[block:]

		err = updateCommentSummariesCore(bboardID, articleID, first, updateNanoTS)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateCommentSummariesCore(bboardID bbs.BBoardID, articleID bbs.ArticleID, commentSummaries []*CommentSummary, updateNanoTS types.NanoTS) (err error) {
	// this is to update only, must be with updateNanoTS.
	theList := make([]*db.UpdatePair, len(commentSummaries))
	for idx, each := range commentSummaries {
		query := bson.M{
			COMMENT_BBOARD_ID_b:  bboardID,
			COMMENT_ARTICLE_ID_b: articleID,
			COMMENT_COMMENT_ID_b: each.CommentID,

			COMMENT_UPDATE_NANO_TS_b: bson.M{
				"$lt": updateNanoTS,
			},
		}
		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: bson.M{
				"$set": each,
				"$unset": bson.M{
					COMMENT_IS_DELETED_b:    true,
					COMMENT_DELETE_REASON_b: true,
				},
			},
		}
	}

	_, err = Comment_c.BulkUpdateOneOnlyNoSet(theList)

	return err
}
