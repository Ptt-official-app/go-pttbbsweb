package schema

import (
	"sort"
	"strings"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentMD5 struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`

	CommentID types.CommentID   `bson:"cid"`
	TheType   types.CommentType `bson:"type"`

	CreateTime types.NanoTS `bson:"create_time_nano_ts"`
	MD5        string       `bson:"md5"`

	FirstCreateTime    types.NanoTS `bson:"first_create_time_nano_ts,omitempty"`    //create-time from first-comments.
	InferredCreateTime types.NanoTS `bson:"inferred_create_time_nano_ts,omitempty"` //create-time from inferred.
	NewCreateTime      types.NanoTS `bson:"new_create_time_nano_ts,omitempty"`      //create-time from new comment.

	SortTime types.NanoTS `bson:"sort_time_nano_ts"`
	TheDate  string       `bson:"the_date"`

	EditNanoTS types.NanoTS `bson:"edit_nano_ts"` //for reply.
}

var (
	EMPTY_COMMENT_MD5 = &CommentMD5{}
	commentMD5Fields  = getFields(EMPTY_COMMENT, EMPTY_COMMENT_MD5)
)

func GetAllCommentMD5s(boardID bbs.BBoardID, articleID bbs.ArticleID) (commentMD5s []*CommentMD5, err error) {

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
	//find
	err = Comment_c.Find(query, 0, &commentMD5s, commentMD5Fields, sortOpts)
	if err != nil {
		return nil, err
	}

	commentMD5s = SortCommentMD5sBySortTime(commentMD5s)

	return commentMD5s, nil
}

func SortCommentMD5sBySortTime(commentMD5s []*CommentMD5) (newCommentMD5s []*CommentMD5) {
	sort.SliceStable(commentMD5s, func(i, j int) bool {
		if commentMD5s[i].SortTime == commentMD5s[j].SortTime {
			return strings.Compare(string(commentMD5s[i].CommentID), string(commentMD5s[j].CommentID)) < 0
		} else {
			return commentMD5s[i].SortTime-commentMD5s[j].SortTime < 0
		}
	})

	return commentMD5s
}
