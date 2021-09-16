package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentDBCS struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	ArticleID bbs.ArticleID   `bson:"aid"`
	CommentID types.CommentID `bson:"cid"`

	SortTime types.NanoTS `bson:"sort_time_nano_ts"`
	DBCS     []byte       `bson:"dbcs"`
}

var (
	EMPTY_COMMENT_DBCS = &CommentDBCS{}
	commentDBCSFields  = getFields(EMPTY_COMMENT, EMPTY_COMMENT_DBCS)
)

func GetAllCommentDBCSs(boardID bbs.BBoardID, articleID bbs.ArticleID) (commentDBCS []*CommentDBCS, err error) {
	query := bson.M{
		COMMENT_BBOARD_ID_b:  boardID,
		COMMENT_ARTICLE_ID_b: articleID,
		COMMENT_IS_DELETED_b: bson.M{
			"$exists": false,
		},
	}

	sortOpts := bson.D{
		{Key: COMMENT_SORT_TIME_b, Value: 1},
		{Key: COMMENT_COMMENT_ID_b, Value: 1},
	}
	// find
	err = Comment_c.Find(query, 0, &commentDBCS, commentDBCSFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return commentDBCS, nil
}
