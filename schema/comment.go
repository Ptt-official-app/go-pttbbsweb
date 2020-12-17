package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type Comment struct {
	BBoardID   bbs.BBoardID    `bson:"bid"`
	ArticleID  bbs.ArticleID   `bson:"aid"`
	CommentID  types.CommentID `bson:"cid"`
	TheType    int             `bson:"type"`
	RefID      types.CommentID `bson:"refid"`
	IsDeleted  bool            `bson:"deleted"`
	CreateTime types.Time8     `bson:"create_time_ts"`
	Owner      bbs.UUserID     `bson:"owner"`
	Date       string          `bson:"date"`
	Content    string          `bson:"content"` //content in comment is colorless.
	IP         string          `bson:"ip"`
	Country    string          `bson:"country"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
