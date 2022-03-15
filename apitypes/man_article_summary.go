package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
)

type ManArticleSummary struct {
	FBoardID   FBoardID    `json:"bid"`         //
	ArticleID  FArticleID  `json:"aid"`         //
	LevelIdx   FArticleID  `json:"level_idx"`   //
	CreateTime types.Time8 `json:"create_time"` //
	MTime      types.Time8 `json:"modified"`    //

	Title string `json:"title"` //

	IsDir bool `json:"is_dir"`
}

func NewManArticleSummary(articleSummary_db *schema.ManArticleSummary) (articleSummary *ManArticleSummary) {
	return &ManArticleSummary{
		FBoardID:  ToFBoardID(articleSummary_db.BBoardID),
		ArticleID: ToFArticleIDFromManArticleID(articleSummary_db.ArticleID),

		LevelIdx: ToFArticleIDFromManArticleID(articleSummary_db.LevelIdx),

		CreateTime: articleSummary_db.CreateTime.ToTime8(),
		MTime:      articleSummary_db.MTime.ToTime8(),

		Title: articleSummary_db.Title,
		IsDir: articleSummary_db.IsDir,
	}
}
