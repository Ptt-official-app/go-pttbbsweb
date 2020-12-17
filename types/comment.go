package types

import "github.com/Ptt-official-app/go-pttbbs/bbs"

type Comment struct {
	BBoardID   bbs.BBoardID  `json:"bid"`
	ArticleID  bbs.ArticleID `json:"aid"`
	CommentID  CommentID     `json:"cid"`
	TheType    int           `json:"type"`
	RefID      CommentID     `json:"refid"`
	IsDeleted  bool          `json:"deleted"`
	CreateTime Time8         `json:"create_time"`
	Owner      bbs.UUserID   `json:"owner"`
	Date       string        `json:"date"`
	Content    string        `json:"content"`
	IP         string        `json:"ip"`
	Country    string        `json:"country"`
}
