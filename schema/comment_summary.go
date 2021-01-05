package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentSummary struct {
	CommentID    types.CommentID `bson:"cid"`
	CreateTime   types.NanoTS    `bson:"create_time_ts"`
	UpdateNanoTS types.NanoTS    `bson:"update_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"`
}

var (
	EMPTY_COMMENT_SUMMARY = &CommentSummary{}
	commentSummaryFields  = getFields(EMPTY_COMMENT, EMPTY_COMMENT_SUMMARY)
)

//GetCommentSummaries
//
//get comment summaries with startNanoTS <= createTime  < endNanoTS (excluding endNanoTS)
func GetCommentSummaries(bboardID bbs.BBoardID, articleID bbs.ArticleID, startNanoTS types.NanoTS, endNanoTS types.NanoTS) (commentSummaries []*CommentSummary, err error) {
	query := bson.M{
		COMMENT_BBOARD_ID_b:  bboardID,
		COMMENT_ARTICLE_ID_b: articleID,
		COMMENT_CREATE_TIME_b: bson.M{
			"$gte": startNanoTS,
			"$lt":  endNanoTS,
		},
	}

	err = Comment_c.Find(query, 0, &commentSummaries, commentSummaryFields)
	if err != nil {
		return nil, err
	}

	return commentSummaries, nil
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
	//this is to update only, must be with updateNanoTS.
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
