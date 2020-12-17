package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types/proto"
)

type Article struct {
	BBoardID   bbs.BBoardID     `bson:"bid"`
	ArticleID  bbs.ArticleID    `bson:"aid"`
	IsDeleted  bool             `bson:"deleted"`
	Filename   string           `bson:"filename"`
	CreateTime types.Time8      `bson:"create_time_ts"`
	MTime      types.Time8      `bson:"mtime_ts"`
	Recommend  int              `bson:"recommend"`
	Owner      bbs.UUserID      `bson:"owner"`
	Date       string           `bson:"date"`
	Title      string           `bson:"title"`
	Money      int              `bson:"money"`
	Type       string           `bson:"type"`
	Class      string           `bson:"class"`
	Filemode   ptttype.FileMode `bson:"mode"`
	URL        string           `bson:"url"`
	Read       bool             `bson:"read"`

	Content proto.Content `bson:"content"`
	IP      string        `bson:"ip"`
	Country string        `bson:"country"`
	BBS     string        `bson:"bbs"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
