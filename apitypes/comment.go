package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type Comment struct {
	BBoardID   bbs.BBoardID    `json:"bid"`
	ArticleID  bbs.ArticleID   `json:"aid"`
	CommentID  types.CommentID `json:"cid"`
	TheType    int             `json:"type"`
	RefID      types.CommentID `json:"refid"`
	IsDeleted  bool            `json:"deleted"`
	CreateTime types.Time8     `json:"create_time"`
	Owner      bbs.UUserID     `json:"owner"`
	Date       string          `json:"date"`
	Content    string          `json:"content"`
	IP         string          `json:"ip"`
	Host       string          `json:"host"` //ip 的中文呈現, 外國則為國家.
}
