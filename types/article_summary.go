package types

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type ArticleSummary struct {
	BBoardID   bbs.BBoardID     `json:"bid"`
	ArticleID  bbs.ArticleID    `json:"aid"`
	IsDeleted  bool             `json:"deleted"`
	Filename   string           `json:"filename"`
	CreateTime Time8            `json:"create_time"`
	MTime      Time8            `json:"modified"`
	Recommend  int              `json:"recommend"`
	Owner      string           `json:"owner"`
	Date       string           `json:"date"`
	Title      string           `json:"title"`
	Money      int              `json:"money"`
	Filemode   ptttype.FileMode `json:"mode"`
	URL        string           `json:"url"`
	Read       bool             `json:"read"`
}

func (a *ArticleSummary) Deserialize(a_b *bbs.ArticleSummary) {
	a.BBoardID = a_b.BBoardID
	a.ArticleID = a_b.ArticleID
	a.IsDeleted = a_b.IsDeleted
	a.Filename = a_b.Filename
	a.CreateTime = Time8(a_b.CreateTime)
	a.MTime = Time8(a_b.MTime)
	a.Recommend = int(a_b.Recommend)
	a.Owner = a_b.Owner
	a.Date = a_b.Date
	a.Title = a_b.Title
	a.Money = int(a_b.Money)
	a.Filemode = a_b.Filemode
	a.URL = a_b.URL
	a.Read = a_b.Read
}
