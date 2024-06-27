package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentSortTime struct {
	CommentID types.CommentID `bson:"cid"`
	SortTime  types.NanoTS    `bson:"sort_time_nano_ts"`
}

var (
	EMPTY_COMMENT_SORT_TIME = &CommentSortTime{}
	commentSortTimeFields   = getFields(EMPTY_COMMENT, EMPTY_COMMENT_SORT_TIME)
)

func GetCommentSortTimeMapByCommentIDs(boardID bbs.BBoardID, articleID bbs.ArticleID, commentIDs []types.CommentID) (sortTimeMap map[types.CommentID]types.NanoTS, err error) {
	query := bson.M{
		COMMENT_BBOARD_ID_b:  boardID,
		COMMENT_ARTICLE_ID_b: articleID,
		COMMENT_COMMENT_ID_b: bson.M{
			"$in": commentIDs,
		},
	}

	// find
	var commentSortTimes []*CommentSortTime
	err = Comment_c.Find(query, 0, &commentSortTimes, commentSortTimeFields, nil)
	if err != nil {
		return nil, err
	}

	sortTimeMap = make(map[types.CommentID]types.NanoTS)
	for _, each := range commentSortTimes {
		sortTimeMap[each.CommentID] = each.SortTime
	}

	return sortTimeMap, nil
}
