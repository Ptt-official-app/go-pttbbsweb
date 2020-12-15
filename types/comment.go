package types

import "github.com/Ptt-official-app/go-pttbbs/bbs"

type Comment struct {
	BBoardID   bbs.BBoardID  `json:"bid"`
	ArticleID  bbs.ArticleID `json:"aid"`
	CommentID  string        `json:"cid"`
	TheType    int           `json:"type"`
	RefID      string        `json:"refid"`
	IsDeleted  bool          `json:"deleted"`
	CreateTime Time8         `json:"create_time"`
	Owner      string        `json:"owner"`
	Date       string        `json:"date"`
	Content    string        `json:"content"`
	IP         string        `json:"ip"`
	Country    string        `json:"country"`
}
